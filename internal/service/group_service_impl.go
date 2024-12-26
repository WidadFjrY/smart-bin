package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/domain/web"
	"smart-trash-bin/internal/repository"
	"smart-trash-bin/pkg/exception"
	"smart-trash-bin/pkg/helper"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type GroupServiceImpl struct {
	DB        *gorm.DB
	Validator *validator.Validate
	Repo      repository.GroupRepostory
}

func NewGroupService(db *gorm.DB, validator *validator.Validate, repo repository.GroupRepostory) GroupService {
	return &GroupServiceImpl{DB: db, Validator: validator, Repo: repo}
}

func (serv *GroupServiceImpl) Create(ctx context.Context, request web.GroupCreateRequest, userId string) web.GroupCreateResponse {
	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	rand.NewSource(time.Now().UnixNano())
	newGroup := model.Group{
		ID:       helper.GenerateRandomString(15),
		UserID:   userId,
		Name:     request.Name,
		Location: request.Location,
	}

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.Repo.Create(ctx, tx, newGroup)
		return nil
	})
	helper.Err(txErr)

	return web.GroupCreateResponse{
		GroupId:   newGroup.ID,
		CreatedAt: time.Now(),
	}
}

func (serv *GroupServiceImpl) GetGroupById(ctx context.Context, groupId string, userId string) web.GroupGetResponse {
	var group model.Group
	var smartBinResponse []web.SmartBins

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		group = serv.Repo.GetGroupById(ctx, tx, groupId)

		if group.UserID != userId {
			panic(exception.NewBadRequestError("user doesn't own this group"))
		}
		for _, smartBin := range group.SmartBin {
			var loadCellValue map[string]interface{}
			var ultraSonicValue map[string]interface{}

			json.Unmarshal(smartBin.LoadCellValue, &loadCellValue)
			json.Unmarshal(smartBin.UltraSonicValue, &ultraSonicValue)
			convSmartBin := web.SmartBins{
				Id: smartBin.ID,
				Data: web.Data{
					Name:                 smartBin.Name,
					UserID:               userId,
					OrganicWeight:        loadCellValue["organic"],
					NonOrganicWeight:     loadCellValue["non_organic"],
					OrganicHeight:        ultraSonicValue["organic"],
					NonOrganicHeight:     ultraSonicValue["non_organic"],
					TotalOrganicWaste:    smartBin.TotalOrganicWaste,
					TotalNonOrganicWaste: smartBin.TotalNonOrganicWaste,
					IsLocked:             smartBin.IsLocked,
					Location:             smartBin.Location,
					Config: web.ConfigGetResponse{
						ConfigId:  smartBin.Config.ID,
						MaxHeight: smartBin.Config.MaxHeight,
						MaxWeight: smartBin.Config.MaxWeight,
					},
					CreatedAt: smartBin.CreatedAt,
					UpdatedAt: smartBin.UpdatedAt,
				},
			}
			smartBinResponse = append(smartBinResponse, convSmartBin)
		}
		return nil
	})
	helper.Err(txErr)
	return web.GroupGetResponse{
		GroupId: group.ID,
		Group: web.Group{
			UserId:    group.UserID,
			Name:      group.Name,
			Location:  group.Location,
			Bins:      smartBinResponse,
			CreatedAt: group.CreatedAt,
			UpdatedAt: group.UpdatedAt,
		},
	}
}

func (serv *GroupServiceImpl) GetGroups(ctx context.Context, page int, userId string) ([]web.GroupGetResponses, int64, int) {
	limit := 10
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	var groups []model.Group
	var groupGetResponses []web.GroupGetResponses
	var totalItems int64

	errTx := serv.DB.Transaction(func(tx *gorm.DB) error {
		groups, totalItems = serv.Repo.GetGroups(ctx, tx, userId, offset, limit)
		return nil
	})
	helper.Err(errTx)

	for _, group := range groups {
		groupGetResponse := web.GroupGetResponses{
			Id: group.ID,
			Group: web.GroupResponses{
				UserId:    group.UserID,
				Name:      group.Name,
				Bins:      []web.SmartBinNames{},
				Location:  group.Location,
				CreatedAt: group.CreatedAt,
				UpdatedAt: group.UpdatedAt,
			},
		}
		for _, smartBin := range group.SmartBin {
			smartBinName := web.SmartBinNames{
				BinId:    smartBin.ID,
				Name:     smartBin.Name,
				IsLocked: smartBin.IsLocked,
			}
			groupGetResponse.Group.Bins = append(groupGetResponse.Group.Bins, smartBinName)
		}
		groupGetResponses = append(groupGetResponses, groupGetResponse)
	}

	totalPages := math.Ceil(float64(totalItems) / float64(limit))
	if page > int(totalPages) && totalItems != 0 {
		panic(exception.NewBadRequestError(fmt.Sprintf("only have %v page", totalPages)))
	}
	if totalItems == 0 {
		panic(exception.NewBadRequestError("user doesn't have any group"))
	}

	return groupGetResponses, totalItems, int(totalPages)
}
