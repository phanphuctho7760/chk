package repository

import (
	"chk/models"
	"chk/utils"
	"reflect"
)

type PriceHistoryRepository struct {
}

const tableName = "price_histories"

func (receiver PriceHistoryRepository) Create(data interface{}) models.Result {
	priceHistoryModel := new(models.PriceHistory)
	resultCreate := priceHistoryModel.Conn().Table(tableName).Create(data)

	return models.Result{
		StatusCode: "success",
		StatusMsg:  "ok",
		Error:      nil,
		Body:       resultCreate,
	}
}

func (receiver PriceHistoryRepository) Get(getParam models.GetParam) models.ResultGet {
	var priceHistories []models.PriceHistory
	var total int
	var totalInt64 int64
	var count int
	conn := new(models.PriceHistory).Conn()

	query := conn.Table(tableName)
	//query := conn.Table(tableName).Where("id", 1)
	query.Count(&totalInt64)
	total = int(totalInt64)

	resultQuery := query.Scopes(utils.Paginate(getParam.Page, getParam.Limit)).Find(&priceHistories)
	count = reflect.ValueOf(priceHistories).Len()

	if resultQuery.Error != nil {
		return models.ResultGet{
			StatusCode: "error",
			StatusMsg:  resultQuery.Error.Error(),
			Error:      resultQuery.Error,
			Body:       models.Paginate{},
		}
	}

	return models.ResultGet{
		StatusCode: "success",
		StatusMsg:  "ok",
		Error:      nil,
		Body: models.Paginate{
			Count:  count,
			Total:  total,
			Prices: priceHistories,
		},
	}
}
