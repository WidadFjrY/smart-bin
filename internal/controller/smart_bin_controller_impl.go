package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"smart-trash-bin/domain/web"
	"smart-trash-bin/internal/service"
	"smart-trash-bin/pkg/exception"
	"smart-trash-bin/pkg/helper"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type SmartBinControllerImpl struct {
	SmartBinServ service.SmartBinService
}

func NewSmartBinController(smartBinServ service.SmartBinService) SmartBinController {
	return &SmartBinControllerImpl{SmartBinServ: smartBinServ}
}

func (cntrl *SmartBinControllerImpl) AddSmartBin(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	var request web.SmartBinCreateRequest

	if ctx.Request.ContentLength == 0 {
		panic(exception.NewBadRequestError("request body is empty"))
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	helper.Err(err)

	cntrl.SmartBinServ.IsSmartBinExsist(ctx.Request.Context(), request.BinId)

	mqttConf := web.MQTTRequest{
		ClientId: "server",
		Topic:    fmt.Sprintf("smartBin/add/%s", request.BinId),
		Payload:  "ADD_SMART_BIN_TO_SERVER",
		MsgResp:  "OK",
	}

	mqttResp := helper.NewMQTT(mqttConf)
	if mqttResp {
		response := cntrl.SmartBinServ.AddSmartBin(ctx.Request.Context(), request, userId.(string))
		helper.Response(ctx, http.StatusOK, "Ok", response)
	} else {
		panic(exception.NewNotFoundError(fmt.Sprintf("smart bin with id %s not found", request.BinId)))
	}
}

func (cntrl *SmartBinControllerImpl) UpdateSmartBinById(ctx *gin.Context) {
	var request web.SmartBinUpdateRequest
	userId, _ := ctx.Get("user_id")
	binId := ctx.Params.ByName("bin_id")

	if len(binId) == 0 {
		panic(exception.NewBadRequestError("bin_id required"))
	}

	if ctx.Request.ContentLength == 0 {
		panic(exception.NewBadRequestError("request body is empty"))
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	helper.Err(err)

	response := cntrl.SmartBinServ.UpdateSmartBinById(ctx.Request.Context(), request, binId, userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (cntrl *SmartBinControllerImpl) GetSmartBinById(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	binId := ctx.Params.ByName("bin_id")

	if len(binId) == 0 {
		panic(exception.NewBadRequestError("bin_id required"))
	}

	response := cntrl.SmartBinServ.GetSmartBinById(ctx.Request.Context(), binId, userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (cntrl *SmartBinControllerImpl) DeleteSmartBinById(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	binId := ctx.Params.ByName("bin_id")

	if len(binId) == 0 {
		panic(exception.NewBadRequestError("bin_id required"))
	}

	cntrl.SmartBinServ.GetSmartBinById(ctx.Request.Context(), binId, userId.(string))

	mqttConf := web.MQTTRequest{
		ClientId: "server",
		Topic:    fmt.Sprintf("smartBin/remove/%s", binId),
		Payload:  "REMOVE_SMART_BIN_FROM_SERVER",
		MsgResp:  "OK",
	}

	mqttResp := helper.NewMQTT(mqttConf)
	if mqttResp {
		response := cntrl.SmartBinServ.DeleteSmartBinById(ctx.Request.Context(), binId, userId.(string))
		helper.Response(ctx, http.StatusOK, "Ok", response)
	} else {
		panic(exception.NewNotFoundError(fmt.Sprintf("failed to remove smart bin with id %s, 'cause smart bin not found", binId)))
	}
}

func (cntrl *SmartBinControllerImpl) GetSmartBins(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	page := ctx.Params.ByName("page")

	if len(page) == 0 {
		panic(exception.NewBadRequestError("page required"))
	}

	pageInt, _ := strconv.Atoi(page)

	data, totalItems, totalPages := cntrl.SmartBinServ.GetSmartBins(ctx, pageInt, userId.(string))
	helper.ResponseWithPage(ctx, http.StatusOK, "Ok", data, pageInt, totalPages, totalItems)
}

func (cntrl *SmartBinControllerImpl) LockSmartBin(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	binId := ctx.Params.ByName("bin_id")

	if len(binId) == 0 {
		panic(exception.NewBadRequestError("bin_id required"))
	}

	smartBin := cntrl.SmartBinServ.GetSmartBinById(ctx, binId, userId.(string))

	if smartBin.SmartBin.IsLocked {
		panic(exception.NewBadRequestError(fmt.Sprintf("failed to lock smart bin with id %s, 'cause smart bin already locked", binId)))
	}

	mqttConf := web.MQTTRequest{
		ClientId: "server",
		Topic:    fmt.Sprintf("smartBin/lock/%s", binId),
		Payload:  "LOCK_SMART_BIN",
		MsgResp:  "OK",
	}

	mqttResp := helper.NewMQTT(mqttConf)
	if mqttResp {
		response := cntrl.SmartBinServ.LockAndUnlockSmartBin(ctx, true, binId)
		helper.Response(ctx, http.StatusOK, "Ok", response)
	} else {
		panic(exception.NewBadRequestError(fmt.Sprintf("failed to lock smart bin with id %s, 'cause smart bin not found", binId)))
	}
}

func (cntrl *SmartBinControllerImpl) UnlockSmartBin(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	binId := ctx.Params.ByName("bin_id")

	if len(binId) == 0 {
		panic(exception.NewBadRequestError("bin_id required"))
	}

	smartBin := cntrl.SmartBinServ.GetSmartBinById(ctx, binId, userId.(string))
	cntrl.SmartBinServ.IsSmartBinFull(ctx, false, binId)

	if !smartBin.SmartBin.IsLocked {
		panic(exception.NewBadRequestError(fmt.Sprintf("failed to unlock smart bin with id %s, 'cause smart bin already unlocked", binId)))
	}

	mqttConf := web.MQTTRequest{
		ClientId: "server",
		Topic:    fmt.Sprintf("smartBin/unlock/%s", binId),
		Payload:  "UNLOCK_SMART_BIN",
		MsgResp:  "OK",
	}

	mqttResp := helper.NewMQTT(mqttConf)
	if mqttResp {
		response := cntrl.SmartBinServ.LockAndUnlockSmartBin(ctx, false, binId)
		helper.Response(ctx, http.StatusOK, "Ok", response)
	} else {
		panic(exception.NewBadRequestError(fmt.Sprintf("failed to unlock smart bin with id %s, 'cause smart bin not found", binId)))
	}
}

func (cntrl *SmartBinControllerImpl) ClassifyImage(ctx *gin.Context) {
	godotenv.Load()
	binId := ctx.GetHeader("BinID")
	if binId == "" {
		panic(exception.NewBadRequestError("header with key BinID required"))
	}

	file, err := ctx.FormFile("file")

	fmt.Println(err)

	allowedExtensionImage := map[string]bool{
		"jpg":  true,
		"png":  true,
		"jpeg": true,
	}

	extension := strings.Split(file.Filename, ".")[1]
	filePath := fmt.Sprintf("waste_model/images/%s.%s", helper.GenerateRandomString(15), extension)

	if !allowedExtensionImage[extension] {
		panic(exception.NewBadRequestError("file not supported! file must be jpg/png/jpeg"))
	}

	err = ctx.SaveUploadedFile(file, filePath)
	helper.Err(err)

	apiURL := os.Getenv("API_URL_CLASSIFY")
	requestData := web.ClassifyRequest{
		PathImage: filePath,
	}

	requestBody, err := json.Marshal(requestData)
	helper.Err(err)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	helper.Err(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("AUTH", os.Getenv("SECRET_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	helper.Err(err)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	helper.Err(err)

	var classifyResponse web.ClassifyResponse
	err = json.Unmarshal(body, &classifyResponse)
	helper.Err(err)

	response := cntrl.SmartBinServ.ClassifyImage(ctx, binId, classifyResponse)

	os.Remove(filePath)
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (cntrl *SmartBinControllerImpl) UpdateSmartBinValue(ctx *gin.Context) {
	var request web.UpdateValueRequest
	binId := ctx.GetHeader("BinId")

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	helper.Err(err)

	response := cntrl.SmartBinServ.UpdateDataSmartBin(ctx.Request.Context(), binId, request)
	if response.Loked {
		mqttConf := web.MQTTRequest{
			ClientId: "server",
			Topic:    fmt.Sprintf("user/notificaton/%s", response.UserId),
			Payload:  fmt.Sprintf("Smart bin with ID %s automatically locked 'cause %s", response.ID, response.LokedDesc),
			MsgResp:  "OK",
		}
		helper.NewMQTT(mqttConf)
	}
	helper.Response(ctx, http.StatusOK, "Ok", response)
}
