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
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type SmartBinServiceImpl struct {
	DB           *gorm.DB
	Validator    *validator.Validate
	SmartBinRepo repository.SmartBinRepository
}

func NewSmartBinService(db *gorm.DB, validator *validator.Validate, smartBinRepo repository.SmartBinRepository) SmartBinService {
	return &SmartBinServiceImpl{DB: db, Validator: validator, SmartBinRepo: smartBinRepo}
}

func (serv *SmartBinServiceImpl) AddSmartBin(ctx context.Context, request web.SmartBinCreateRequest, userId string) web.SmartBinCreateResponse {
	configRepo := repository.NewConfigRepository()
	configServ := NewConfigService(serv.DB, serv.Validator, configRepo)

	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	rand.NewSource(time.Now().UnixNano())

	var smartBin model.SmartBin
	loadCellValue := map[string]float64{
		"organic":     0,
		"non_organic": 0,
	}

	ultraSonicValue := map[string]float64{
		"organic":     0,
		"non_organic": 0,
	}

	jsonLoadCell, _ := json.Marshal(loadCellValue)
	jsonUltraSonic, _ := json.Marshal(ultraSonicValue)

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		smartBin = serv.SmartBinRepo.CreateSmartBin(ctx, tx, model.SmartBin{
			ID:              request.BinId,
			UserID:          userId,
			LoadCellValue:   datatypes.JSON(jsonLoadCell),
			UltraSonicValue: datatypes.JSON(jsonUltraSonic),
			Name:            fmt.Sprintf("SmartBin_%s", helper.GenerateRandomString(10)),
			Location:        "location01",
		})
		return nil
	})
	helper.Err(txErr)

	configServ.AddConfig(ctx, model.Config{
		BinID: request.BinId,
	})

	return web.SmartBinCreateResponse{
		ID:      smartBin.ID,
		AddedAt: smartBin.CreatedAt,
	}
}

func (serv *SmartBinServiceImpl) IsSmartBinExsist(ctx context.Context, binId string) {

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		_, isExsist := serv.SmartBinRepo.GetSmartBinById(ctx, tx, binId)
		if isExsist {
			panic(exception.NewBadRequestError(fmt.Sprintf("smart bin with id %s is owned by another user", binId)))
		}
		return nil
	})
	helper.Err(txErr)

}

func (serv *SmartBinServiceImpl) UpdateSmartBinById(ctx context.Context, request web.SmartBinUpdateRequest, binId string, userId string) web.SmartBinUpdateResponse {
	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	var smartBin model.SmartBin

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		smartBin, isExsist := serv.SmartBinRepo.GetSmartBinById(ctx, tx, binId)
		if !isExsist {
			panic(exception.NewNotFoundError("smart bin not found"))
		} else {
			if smartBin.UserID != userId {
				panic(exception.NewBadRequestError("user doesn't own this smart bin"))
			}
		}
		smartBin = serv.SmartBinRepo.UpdateSmartBinById(ctx, tx, model.SmartBin{
			ID:       binId,
			Name:     request.Name,
			Location: request.Location,
		})
		return nil
	})
	helper.Err(txErr)

	txErr = serv.DB.Transaction(func(tx *gorm.DB) error {
		smartBin = serv.SmartBinRepo.UpdateSmartBinById(ctx, tx, model.SmartBin{
			ID:       binId,
			Name:     request.Name,
			Location: request.Location,
		})
		return nil
	})
	helper.Err(txErr)

	return web.SmartBinUpdateResponse{
		ID:        smartBin.ID,
		UpdatedAt: smartBin.UpdatedAt,
	}
}

func (serv *SmartBinServiceImpl) GetSmartBinById(ctx context.Context, binId string, userId string) web.SmartBinGetResponse {
	var smartBinModel model.SmartBin

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		smartBin, isExsist := serv.SmartBinRepo.GetSmartBinById(ctx, tx, binId)
		if !isExsist {
			panic(exception.NewNotFoundError("smart bin not found"))
		} else {
			if smartBin.UserID != userId {
				panic(exception.NewBadRequestError("user doesn't own this smart bin"))
			}
			smartBinModel = smartBin
		}
		return nil
	})
	helper.Err(txErr)

	if smartBinModel.GroupID == nil {
		groupId := "not in any group"
		smartBinModel.GroupID = &groupId
	}

	var loadCellValue map[string]interface{}
	var ultraSonicValue map[string]interface{}

	json.Unmarshal(smartBinModel.LoadCellValue, &loadCellValue)
	json.Unmarshal(smartBinModel.UltraSonicValue, &ultraSonicValue)

	return web.SmartBinGetResponse{
		Id: smartBinModel.ID,
		SmartBin: web.SmartBin{
			UserID:               smartBinModel.UserID,
			Name:                 smartBinModel.Name,
			GroupID:              *smartBinModel.GroupID,
			OrganicWeight:        loadCellValue["organic"],
			NonOrganicWeight:     loadCellValue["non_organic"],
			OrganicHeight:        ultraSonicValue["organic"],
			NonOrganicHeight:     ultraSonicValue["non_organic"],
			TotalOrganicWaste:    smartBinModel.TotalOrganicWaste,
			TotalNonOrganicWaste: smartBinModel.TotalNonOrganicWaste,
			IsLocked:             smartBinModel.IsLocked,
			Location:             smartBinModel.Location,
			Config: web.ConfigGetResponse{
				ConfigId:  smartBinModel.Config.ID,
				MaxHeight: smartBinModel.Config.MaxHeight,
				MaxWeight: smartBinModel.Config.MaxWeight,
			},
			CreatedAt: smartBinModel.CreatedAt,
			UpdatedAt: smartBinModel.UpdatedAt,
		},
	}
}

func (serv *SmartBinServiceImpl) DeleteSmartBinById(ctx context.Context, binId string, userId string) web.SmartBinDeleteResponse {
	var smartBinModel model.SmartBin

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		smartBinModel = serv.SmartBinRepo.DeleteSmartBinById(ctx, tx, binId)
		return nil
	})
	helper.Err(txErr)

	return web.SmartBinDeleteResponse{
		ID:        smartBinModel.ID,
		DeletedAt: smartBinModel.DeletedAt.Time,
	}
}

func (serv *SmartBinServiceImpl) GetSmartBins(ctx context.Context, page int, userId string) ([]web.SmartBinGetResponse, int64, int) {
	limit := 10
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	var smartBins []model.SmartBin
	var totalItems int64

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		smartBins = serv.SmartBinRepo.GetSmartBins(ctx, tx, limit, offset, userId)
		totalItems = serv.SmartBinRepo.TotalSmartBin(ctx, tx, userId)
		return nil
	})
	helper.Err(txErr)

	var smartBinResponses []web.SmartBinGetResponse
	for _, smartBin := range smartBins {
		if smartBin.GroupID == nil {
			groupId := "not in any group"
			smartBin.GroupID = &groupId
		}

		var loadCellValue map[string]interface{}
		var ultraSonicValue map[string]interface{}

		json.Unmarshal(smartBin.LoadCellValue, &loadCellValue)
		json.Unmarshal(smartBin.UltraSonicValue, &ultraSonicValue)

		smartBinResponse := web.SmartBinGetResponse{
			Id: smartBin.ID,
			SmartBin: web.SmartBin{
				UserID:               smartBin.UserID,
				GroupID:              *smartBin.GroupID,
				Name:                 smartBin.Name,
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
		smartBinResponses = append(smartBinResponses, smartBinResponse)
	}

	totalPages := math.Ceil(float64(totalItems) / float64(limit))
	if page > int(totalPages) && totalItems != 0 {
		panic(exception.NewBadRequestError(fmt.Sprintf("only have %v page", totalPages)))
	}
	if totalItems == 0 {
		panic(exception.NewBadRequestError("user doesn't have any smart bin"))
	}
	return smartBinResponses, totalItems, int(totalPages)
}

func (serv *SmartBinServiceImpl) LockAndUnlockSmartBin(ctx context.Context, status bool, binIds []string) web.LockAndUnlockResponse {
	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		for _, binId := range binIds {
			serv.SmartBinRepo.LockAndUnlockSmartBin(ctx, tx, binId, status)

		}
		return nil
	})
	helper.Err(txErr)

	historyRepo := repository.NewHistoryRepository()
	historyServ := NewHistoryService(serv.DB, historyRepo)

	for _, binId := range binIds {
		if status {
			historyServ.CreateHistory(ctx, web.HistoryCreateRequest{
				BinID:   binId,
				Status:  "Success",
				Message: "Smart Bin has been locked by user",
			})
		} else {
			historyServ.CreateHistory(ctx, web.HistoryCreateRequest{
				BinID:   binId,
				Status:  "Success",
				Message: "Smart Bin has been unlocked by user",
			})
		}
	}

	return web.LockAndUnlockResponse{
		BinId:     binIds,
		UpdatedAt: time.Now(),
	}
}

func (serv *SmartBinServiceImpl) ClassifyImage(ctx context.Context, binId string, classify web.ClassifyResponse) web.SmartBinClassification {
	classifyResult := map[string]float64{
		"organic":     0,
		"non_organic": 0,
	}

	historyRepo := repository.NewHistoryRepository()
	historyServ := NewHistoryService(serv.DB, historyRepo)

	if classify.Prediction == "organic" {
		classifyResult["organic"] = 1
		historyServ.CreateHistory(ctx, web.HistoryCreateRequest{
			BinID:   binId,
			Status:  "Success",
			Message: "Storing waste in the organic bin",
		})
	} else if classify.Prediction == "non-organic" {
		classifyResult["non_organic"] = 1
		historyServ.CreateHistory(ctx, web.HistoryCreateRequest{
			BinID:   binId,
			Status:  "Success",
			Message: "Storing waste in the non-organic bin",
		})
	}

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.SmartBinRepo.UpdateClassifyValue(ctx, tx, classifyResult, binId)
		return nil
	})
	helper.Err(txErr)

	return web.SmartBinClassification{
		BinId:      binId,
		Waste:      classify.Prediction,
		ClassifyAt: time.Now(),
	}
}

func (serv *SmartBinServiceImpl) UpdateDataSmartBin(ctx context.Context, binId string, request web.UpdateValueRequest) web.UpdateValueResponse {
	configRepo := repository.NewConfigRepository()
	configServ := NewConfigService(serv.DB, serv.Validator, configRepo)

	historyRepo := repository.NewHistoryRepository()
	historyServ := NewHistoryService(serv.DB, historyRepo)

	notifRepo := repository.NewNotificationRepository()
	notifServ := NewNotificationService(serv.DB, notifRepo)

	sensorValues := map[string]float64{
		"load_cell_organic":       request.LoadCellOrganic,
		"load_cell_non_organic":   request.LoadCellNonOrganic,
		"ultra_sonic_organic":     request.UltraSonicOrganic,
		"ultra_sonic_non_organic": request.UltraSonicNonOrganic,
	}

	config := configServ.GetConfigByBinId(ctx, binId)
	exceeded := []bool{
		request.LoadCellOrganic > config.MaxWeight,
		request.UltraSonicOrganic > config.MaxHeight,
		request.LoadCellNonOrganic > config.MaxWeight,
		request.UltraSonicNonOrganic > config.MaxHeight,
	}

	isFull := false
	isAlmostFull := false
	almostFull := 90
	var userId string
	var warn string

	serv.DB.Transaction(func(tx *gorm.DB) error {
		smartBin, _ := serv.SmartBinRepo.GetSmartBinById(ctx, tx, binId)
		userId = smartBin.UserID
		return nil
	})

	for _, condition := range exceeded {
		if condition {
			isFull = true
			serv.DB.Transaction(func(tx *gorm.DB) error {
				serv.SmartBinRepo.LockAndUnlockSmartBin(ctx, tx, binId, true)
				return nil
			})
			if condition == exceeded[0] {
				warn = "the weight of organic trash bin exceeds the limit"
			} else if condition == exceeded[1] {
				warn = "the height of organic trash bin exceeds the limit"
			} else if condition == exceeded[2] {
				warn = "the weight of non-organic trash bin exceeds the limit"
			} else if condition == exceeded[3] {
				warn = "the height of non-organic trash bin exceeds the limit"
			}

			historyServ.CreateHistory(ctx, web.HistoryCreateRequest{
				BinID:   binId,
				Status:  "Success",
				Message: fmt.Sprintf("Smart Bin has been locked automatically, 'cause %s", warn),
			})

			notifServ.CreateNotification(ctx, web.NotificationCreateRequest{
				Title: "Smart Bin Locked Automatically",
				Desc:  fmt.Sprintf("Smart Bin with id %s has been locked automatically, 'cause %s", binId, warn),
			}, userId)
			break
		}

	}

	if (config.MaxHeight*float64(almostFull))/100 < request.UltraSonicOrganic && !isFull {
		warn = fmt.Sprintf("the organic bin with id %s is almost full", binId)
		isAlmostFull = true
	}
	if (config.MaxHeight*float64(almostFull))/100 < request.UltraSonicNonOrganic && !isFull {
		warn = fmt.Sprintf("the non-organic bin with id %s is almost full", binId)
		isAlmostFull = true
	}

	if isAlmostFull && !isFull {
		notifServ.CreateNotification(ctx, web.NotificationCreateRequest{
			Title: "Smart Bin Almost Full",
			Desc:  warn,
		}, userId)
	}

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.SmartBinRepo.UpdateSensorValue(ctx, tx, sensorValues, binId)
		return nil
	})
	helper.Err(txErr)

	return web.UpdateValueResponse{
		ID:        binId,
		UserId:    userId,
		Loked:     isFull,
		LokedDesc: warn,
		UpdatedAt: time.Now(),
	}
}

func (serv *SmartBinServiceImpl) IsSmartBinFull(ctx context.Context, binId string) {
	configRepo := repository.NewConfigRepository()
	configServ := NewConfigService(serv.DB, serv.Validator, configRepo)
	config := configServ.GetConfigByBinId(ctx, binId)

	var smartBin model.SmartBin
	errTx := serv.DB.Transaction(func(tx *gorm.DB) error {
		smartBin, _ = serv.SmartBinRepo.GetSmartBinById(ctx, tx, binId)
		return nil
	})
	helper.Err(errTx)

	var loadCellValue map[string]interface{}
	var ultraSonicValue map[string]interface{}

	json.Unmarshal(smartBin.LoadCellValue, &loadCellValue)
	json.Unmarshal(smartBin.UltraSonicValue, &ultraSonicValue)

	exceeded := []bool{
		loadCellValue["organic"].(float64) > config.MaxWeight,
		ultraSonicValue["organic"].(float64) > config.MaxHeight,
		loadCellValue["non_organic"].(float64) > config.MaxWeight,
		ultraSonicValue["non_organic"].(float64) > config.MaxHeight,
	}

	for _, condition := range exceeded {
		if condition {
			panic(exception.NewBadRequestError(fmt.Sprintf("the trash with id %s must be emptied before it can be opened", binId)))
		}
	}
}

func (serv *SmartBinServiceImpl) AddSmartBinToGroup(ctx context.Context, request web.SmartBinCreateRequest, userId string, groupId string) web.SmartBinUpdateResponse {
	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		smartBin, isExsist := serv.SmartBinRepo.GetSmartBinById(ctx, tx, request.BinId)
		if !isExsist {
			panic(exception.NewNotFoundError("smart bin not found"))
		} else {
			if smartBin.UserID != userId {
				panic(exception.NewBadRequestError("user doesn't own this smart bin"))
			}
		}
		if smartBin.GroupID != nil {
			panic(exception.NewBadRequestError("smart bin already in group"))
		}
		serv.SmartBinRepo.AddSmartBinToGroup(ctx, tx, groupId, request.BinId)
		return nil
	})
	helper.Err(txErr)
	return web.SmartBinUpdateResponse{
		ID:        request.BinId,
		UpdatedAt: time.Now(),
	}
}

func (serv *SmartBinServiceImpl) RemoveSmartBinFromGroup(ctx context.Context, request web.SmartBinCreateRequest, userId string) web.SmartBinUpdateResponse {
	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		smartBin, isExsist := serv.SmartBinRepo.GetSmartBinById(ctx, tx, request.BinId)
		if !isExsist {
			panic(exception.NewNotFoundError("smart bin not found"))
		} else {
			if smartBin.UserID != userId {
				panic(exception.NewBadRequestError("user doesn't own this smart bin"))
			}
		}

		serv.SmartBinRepo.RemoveSmartBinFromGroup(ctx, tx, request.BinId)
		return nil
	})
	helper.Err(txErr)
	return web.SmartBinUpdateResponse{
		ID:        request.BinId,
		UpdatedAt: time.Now(),
	}
}
