package login

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/consts"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/auth"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/crypto"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/random"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/sendto"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
	"go.uber.org/zap"
)

type serviceImpl struct {
	sqlc *database.Queries
}

func New(sqlc *database.Queries) Service {
	return &serviceImpl{
		sqlc: sqlc,
	}
}

func (u *serviceImpl) Register(ctx *gin.Context, in *vo.RegisterInput) (codeStatus int, err error) {
	// TODO: check user base exists
	userFound, err := u.sqlc.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("error for check user already exists: %s", err)
	}

	if userFound {
		return response.ErrCodeAccountAlreadyExists, fmt.Errorf("user already exists")
	}

	// TODO: check user verified otp
	isVerified, err := u.sqlc.CheckUserVerifiedOTP(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("error for otp verified: %s", err)
	}

	if isVerified {
		return response.ErrCodeOTPAlreadyVerified, fmt.Errorf("error for otp verified")
	}

	// TODO: check user spam / rate limiting by ip

	// TODO: hash email
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	// TODO: check user already has active otp
	userKey := utils.GetUserKey(hashKey)
	otpFound, err := global.Redis.Get(ctx, userKey).Result()
	switch {
	// key not exists
	case err == redis.Nil:
	case err != nil:
		return response.ErrCodeInternalServerError, fmt.Errorf("error for get otp: %s", err)
	case otpFound != "":
		return response.ErrCodeOTPAlreadyExists, fmt.Errorf("otp already exists")
	}

	// TODO: generate otp
	otpNew := random.GenerateOTP()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}

	// TODO: save otp in redis
	err = global.Redis.SetEx(ctx, userKey, otpNew, time.Duration(consts.TIME_OTP_REGISTER*time.Minute)).Err()
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("save otp in redis failed: %s", err)
	}

	// TODO: send otp to verify type (email, sms)
	switch in.VerifyType {
	case consts.EMAIL:
		// TODO: Send to kafka -> send email service ->
		// body := make(map[string]interface{})
		// body["otp"] = otpNew
		// body["email"] = in.VerifyKey

		// bodyRequest, _ := json.Marshal(body)

		// msg := kafka.Message{
		// 	Key:   []byte("otp-auth"),
		// 	Value: []byte(bodyRequest),
		// 	Time:  time.Now(),
		// }

		// err = global.Kafka.WriteMessages(global.Ctx, msg)
		// if err != nil {
		// 	return response.ErrCodeSendEmailOTP, err
		// }

		// TODO: send to email
		err = sendto.SendEmail([]string{in.VerifyKey}, "otp_auth.html", map[string]interface{}{
			"otp": otpNew,
		}, consts.OTP_VERIFY)
		if err != nil {
			return response.ErrCodeInternalServerError, fmt.Errorf("send email failed: %s", err)
		}
	}

	// TODO: save otp to mysql
	id := uuid.New().String()
	err = u.sqlc.CreateUserVerify(ctx, database.CreateUserVerifyParams{
		ID:        id,
		Otp:       strconv.Itoa(otpNew),
		VerifyKey: in.VerifyKey,
		KeyHash:   hashKey,
		Type:      in.VerifyType,
		CreatedAt: utiltime.GetTimeNow(),
		UpdatedAt: utiltime.GetTimeNow(),
	})

	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("save otp to database failed: %s", err)
	}

	// TODO: return
	return response.ErrCodeRegisterSuccess, nil
}

func (u *serviceImpl) VerifyOTP(ctx *gin.Context, in *vo.VerifyOTPInput) (codeStatus int, out *vo.VerifyOTPOutput, err error) {
	out = &vo.VerifyOTPOutput{}

	// TODO: hash email
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	// TODO: check otp already exists in redis
	otpFound, err := global.Redis.Get(ctx, utils.GetUserKey(hashKey)).Result()
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for check otp in redis: %s", err)
	}

	// TODO: check otp match
	if in.VerifyCode != otpFound {
		return response.ErrCodeOTPNotMatch, nil, fmt.Errorf("OTP not match")
	}

	// TODO: if match, get info otp
	infoOTP, err := u.sqlc.GetUserUnverify(ctx, hashKey)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get info otp failed: %s", err)
	}

	if infoOTP.IsVerified != 0 {
		return response.ErrCodeOTPAlreadyVerified, nil, fmt.Errorf("otp is verified")
	}

	// TODO: update user verify status and delete otp
	err = u.sqlc.UpdateUserVerifyStatus(ctx, database.UpdateUserVerifyStatusParams{
		UpdatedAt: utiltime.GetTimeNow(),
		KeyHash:   hashKey,
	})

	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("update user verify failed: %s", err)
	}

	// TODO: delete otp in redis
	err = global.Redis.Del(ctx, utils.GetUserKey(hashKey)).Err()
	if err != nil {
		fmt.Printf("VerifyOTP error: %s\n", err.Error())
		global.Logger.Error("VerifyOTP error: ", zap.String("error", err.Error()))
	}

	// TODO: return
	out.Token = infoOTP.KeyHash
	return response.ErrCodeVerifyOTPSuccess, out, nil
}

func (u *serviceImpl) UpdatePasswordRegister(ctx *gin.Context, in *vo.UpdatePasswordRegisterInput) (codeStatus int, err error) {
	// TODO: get info otp by key hash
	infoOTP, err := u.sqlc.GetUserVerified(ctx, in.Token)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("get info otp failed: %s", err)
	}

	// TODO: check otp is verified
	if infoOTP.IsVerified == 0 {
		return response.ErrCodeOTPNotVerified, fmt.Errorf("user OTP not verified")
	}

	// TODO: check user base exists
	userFound, err := u.sqlc.CheckUserBaseExists(ctx, infoOTP.VerifyKey)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("error for check user already exists: %s", err)
	}

	if userFound {
		return response.ErrCodeAccountAlreadyExists, fmt.Errorf("user already exists")
	}

	// TODO: update user base
	hashPassword, err := crypto.HashPassword(in.Password)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("hash password failed: %s", err)
	}
	now := utiltime.GetTimeNow()

	err = u.sqlc.AddUserBase(ctx, database.AddUserBaseParams{
		ID:         infoOTP.ID,
		Account:    infoOTP.VerifyKey,
		Password:   hashPassword,
		LoginTime:  now,
		LoginIp:    "",
		LogoutTime: 0,
		IsVerified: 1,
		CreatedAt:  now,
		UpdatedAt:  now,
	})
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("save user base failed: %s", err)
	}

	// TODO: create user info
	now = utiltime.GetTimeNow()
	err = u.sqlc.CreateUserInfo(ctx, database.CreateUserInfoParams{
		ID:               infoOTP.ID,
		Account:          infoOTP.VerifyKey,
		Status:           1,
		IsAuthentication: 1,
		CreatedAt:        now,
		UpdatedAt:        now,
	})

	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("update password register failed: %s", err)
	}

	return response.ErrCodeSuccessfully, nil
}

func (u *serviceImpl) Login(ctx *gin.Context, in *vo.LoginInput) (codeStatus int, out *vo.LoginOutput, err error) {
	out = &vo.LoginOutput{}

	// TODO: get user info
	userBase, err := u.sqlc.GetUserBaseByAccount(ctx, in.UserAccount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeLoginFailed, nil, fmt.Errorf("user not found")
		}
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get user base failed: %s", err)
	}

	// TODO: check password match
	if !crypto.CheckPasswordHash(in.UserPassword, userBase.Password) {
		return response.ErrCodeLoginFailed, nil, fmt.Errorf("dose not match password")
	}

	// TODO: check two-factor authentication

	// TODO: update login
	go u.sqlc.LoginUserBase(ctx, database.LoginUserBaseParams{
		LoginIp:   "",
		Account:   in.UserAccount,
		LoginTime: utiltime.GetTimeNow(),
	})

	// TODO: create uuid user
	subToken := utils.GenerateCliTokenUUID(userBase.ID)

	// TODO: get user info
	infoUser, err := u.sqlc.GetUserInfo(ctx, userBase.Account)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get user info failed: %s", err)
	}

	infoUserJson, err := json.Marshal(infoUser)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("convert to json failed: %v", err)
	}

	// TODO: give info user json to redis
	err = global.Redis.SetEx(ctx, subToken, infoUserJson, time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("save user info to redis failed: %s", err)
	}

	out.Token, err = auth.CreateToken(userBase.ID, consts.USER)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}

	return response.ErrCodeLoginSuccess, out, nil
}
