package controller

import (
	"chk/models"
	"chk/services"
	"chk/utils"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	_ "github.com/joho/godotenv/autoload"
	"go/types"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

type HomeController struct {
}

func (receiver HomeController) Upload(c *gin.Context) {
	mediaType, params, err := mime.ParseMediaType(c.GetHeader("Content-Type"))
	if err != nil {
		utils.HttpErrorResponse(c, http.StatusBadRequest, err, "parse media type error")
		return
	}

	if !strings.HasPrefix(mediaType, "multipart/") {
		utils.HttpErrorResponse(c, http.StatusBadRequest, err, "content type not support")
		return
	}

	var priceHistories []models.PriceHistory
	var parts []byte
	mr := multipart.NewReader(c.Request.Body, params["boundary"])

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			utils.HttpErrorResponse(c, http.StatusBadRequest, err, "reading next part error")
			return
		}
		parts, err = ioutil.ReadAll(p)
		if err != nil {
			utils.HttpErrorResponse(c, http.StatusBadRequest, err, "reading part error")
			return
		}
	}

	err = gocsv.UnmarshalBytes(parts, &priceHistories)
	if err != nil {
		utils.HttpErrorResponse(c, http.StatusUnprocessableEntity, err, "parse csv error")
		return
	}

	priceHistoryService := new(services.PriceHistoryService)
	result := priceHistoryService.Create(priceHistories)

	if result.Error != nil {
		utils.HttpErrorResponse(c, http.StatusServiceUnavailable, result.Error, "service error")
		return
	}

	utils.HttpSuccessResponse(c, http.StatusOK, types.Interface{}, "ok")
}

func (receiver HomeController) Get(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}

	getParam := models.GetParam{
		Page:  page,
		Limit: limit,
	}

	priceHistoryService := new(services.PriceHistoryService)
	result := priceHistoryService.Get(getParam)
	if result.Error != nil {
		utils.HttpErrorResponse(c, http.StatusServiceUnavailable, result.Error, "service error")
	}

	utils.HttpSuccessResponse(c, http.StatusOK, result.Body, "ok")
}
