// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Transaction is the golang structure of table transaction for DAO operations like Where/Data.
type Transaction struct {
	g.Meta        `orm:"table:transaction, do:true"`
	Id            interface{} //
	ProductName   interface{} //
	Payer         interface{} //
	Payee         interface{} //
	Money         interface{} // 单位为分
	IsReceive     interface{} // 0未收货，1已收货
	BusinessLevel interface{} //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
