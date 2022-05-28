// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Transaction is the golang structure for table transaction.
type Transaction struct {
	Id            int         `json:"id"             description:""`
	ProductName   string      `json:"product_name"   description:""`
	Payer         string      `json:"payer"          description:""`
	Payee         string      `json:"payee"          description:""`
	Money         int         `json:"money"          description:"单位为分"`
	IsReceive     int         `json:"is_receive"     description:"0未收货，1已收货"`
	BusinessLevel int         `json:"business_level" description:""`
	CreatedAt     *gtime.Time `json:"created_at"     description:""`
	UpdatedAt     *gtime.Time `json:"updated_at"     description:""`
}
