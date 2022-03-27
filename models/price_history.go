package models

import (
	database "chk/infrastructure"
	"gorm.io/gorm"
)

type PriceHistory struct {
	Unix   int64   `gorm:"unix:milli" json:"unix" csv:"UNIX"`
	Symbol string  `json:"symbol" csv:"SYMBOL"`
	Open   float64 `json:"open" csv:"OPEN"`
	High   float64 `json:"high" csv:"HIGH"`
	Low    float64 `json:"low" csv:"LOW"`
	Close  float64 `json:"close" csv:"CLOSE"`
}

func (receiver PriceHistory) Conn() *gorm.DB {
	return database.DbGorm
}
