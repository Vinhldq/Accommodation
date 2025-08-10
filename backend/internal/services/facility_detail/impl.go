package facilitydetail

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
)

type serviceImpl struct {
	sqlc *database.Queries
}

func New(sqlc *database.Queries) Service {
	return &serviceImpl{sqlc: sqlc}
}

func (f *serviceImpl) DeleteFacilityDetail(ctx *gin.Context, in *vo.DeleteFacilityDetailInput) (codeStatus int, err error) {
	// TODO: get userID from gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, fmt.Errorf("userID not found in context")
	}

	// TODO: check user is admin
	isExists, err := f.sqlc.CheckUserAdminExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("get admin failed")
	}

	if !isExists {
		return response.ErrCodeForbidden, fmt.Errorf("user not admin")
	}

	err = f.sqlc.DeleteFacilityDetail(ctx, in.ID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("delete facility detail failed: %s", err)
	}
	return response.ErrCodeDeleteFacilityDetailSuccess, nil
}

func (f *serviceImpl) UpdateFacilityDetail(ctx *gin.Context, in *vo.UpdateFacilityDetailInput) (codeStatus int, out *vo.UpdateFacilityDetailOutput, err error) {
	out = &vo.UpdateFacilityDetailOutput{}

	// TODO: get userID from gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user is admin
	isExists, err := f.sqlc.CheckUserAdminExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get admin failed")
	}

	if !isExists {
		return response.ErrCodeForbidden, nil, fmt.Errorf("user not admin")
	}

	// TODO: update facility
	now := utiltime.GetTimeNow()
	err = f.sqlc.UpdateFacilityDetail(ctx, database.UpdateFacilityDetailParams{
		Name:      in.Name,
		UpdatedAt: now,
		ID:        in.ID,
	})
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("update facility detail failed: %s", err)
	}

	// TODO: get facility
	facility, err := f.sqlc.GetAccommodationFacilityDetailById(ctx, in.ID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get facility failed: %s", err)
	}

	out.ID = facility.ID
	out.Name = facility.Name
	return response.ErrCodeUpdateFacilityDetailSuccess, out, nil
}

func (f *serviceImpl) CreateFacilityDetail(ctx *gin.Context, in *vo.CreateFacilityDetailInput) (codeStatus int, out *vo.CreateFacilityDetailOutput, err error) {
	out = &vo.CreateFacilityDetailOutput{}

	// TODO: get userID from gin context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user is admin
	isExists, err := f.sqlc.CheckUserAdminExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get admin failed")
	}

	if !isExists {
		return response.ErrCodeForbidden, nil, fmt.Errorf("user not admin")
	}

	// TODO: save facility
	id := uuid.New().String()
	now := utiltime.GetTimeNow()
	err = f.sqlc.CreateAccommodationFacilityDetail(ctx, database.CreateAccommodationFacilityDetailParams{
		ID:        id,
		Name:      in.Name,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("create facility failed: %s", err)
	}

	// TODO: get facility
	facility, err := f.sqlc.GetAccommodationFacilityDetailById(ctx, id)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get facility failed: %s", err)
	}

	out.ID = facility.ID
	out.Name = facility.Name
	return response.ErrCodeCreateFacilityDetailSuccess, out, nil
}

func (f *serviceImpl) GetFacilityDetail(ctx *gin.Context) (codeStatus int, out []*vo.GetFacilityDetailOutput, err error) {
	// TODO: get facilities
	facilities, err := f.sqlc.GetAccommodationFacilityDetail(ctx)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get facility failed: %s", err)
	}

	for _, facility := range facilities {
		out = append(out, &vo.GetFacilityDetailOutput{
			ID:   facility.ID,
			Name: facility.Name,
		})
	}

	return response.ErrCodeGetFacilityDetailSuccess, out, nil
}
