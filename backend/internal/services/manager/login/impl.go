package login

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/consts"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/auth"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/crypto"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
)

type serviceImpl struct {
	sqlc *database.Queries
}

func New(sqlc *database.Queries) Service {
	return &serviceImpl{sqlc: sqlc}
}

func (m *serviceImpl) Login(ctx *gin.Context, in *vo.ManagerLoginInput) (codeStatus int, out *vo.ManagerLoginOutput, err error) {
	out = &vo.ManagerLoginOutput{}

	// TODO: get manager info
	userManager, err := m.sqlc.GetUserManager(ctx, in.UserAccount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeLoginFailed, nil, fmt.Errorf("manager not found")
		}
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get user manager failed: %s", err)
	}

	// TODO: check password match
	if !crypto.CheckPasswordHash(in.UserPassword, userManager.Password) {
		return response.ErrCodeLoginFailed, nil, fmt.Errorf("dose not match password")
	}

	// TODO: check two-factor authentication

	// TODO: update login
	go m.sqlc.UpdateUserManagerLogin(ctx, database.UpdateUserManagerLoginParams{
		LoginTime: utiltime.GetTimeNow(),
		Account:   in.UserAccount,
	})

	// TODO: create uuid user
	subToken := utils.GenerateCliTokenUUID(userManager.ID)

	userManagerInfor := vo.ManagerInfor{
		Account:  userManager.Account,
		UserName: userManager.UserName,
	}

	userManagerInforJson, err := json.Marshal(userManagerInfor)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("convert to json failed: %v", err)
	}

	// TODO: save manager info to redis
	err = global.Redis.SetEx(ctx, subToken, userManagerInforJson, time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("save manager info to redis failed: %s", err)
	}

	out.Token, err = auth.CreateToken(userManager.ID, consts.MANAGER)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("error for create token failed: %s", err)
	}

	out.Account = userManagerInfor.Account
	out.UserName = userManagerInfor.UserName

	return response.ErrCodeLoginSuccess, out, nil
}

func (m *serviceImpl) Register(ctx *gin.Context, in *vo.ManagerRegisterInput) (codeStatus int, err error) {
	// TODO: check user is admin
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, fmt.Errorf("userID not found in context")
	}

	// TODO: check user exists
	exists, err := m.sqlc.CheckUserAdminExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("get user admin failed: %s", err)
	}

	if !exists {
		return response.ErrCodeForbidden, fmt.Errorf("user admin not found")
	}

	// TODO: check email exists in user manager
	managerFound, err := m.sqlc.CheckUserManagerExistsByEmail(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("error for check manager already exists: %s", err)
	}

	if managerFound {
		return response.ErrCodeAccountAlreadyExists, fmt.Errorf("manager already exists")
	}

	// TODO: check user spam / rate limiting by ip

	// TODO: create manager
	id := uuid.New().String()
	now := utiltime.GetTimeNow()
	hashPassword, err := crypto.HashPassword(in.UserPassword)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("hash password failed: %s", err)
	}

	err = m.sqlc.CreateUserManage(ctx, database.CreateUserManageParams{
		ID:        id,
		Account:   in.UserAccount,
		UserName:  in.Username,
		Password:  hashPassword,
		CreatedAt: now,
		UpdatedAt: now,
	})

	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("error for register manager failed: %s", err)
	}

	return response.ErrCodeRegisterSuccess, nil
}
