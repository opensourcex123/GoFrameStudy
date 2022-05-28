// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ReturnRequest is the golang structure for table return_request.
type ReturnRequest struct {
	Id             int         `json:"id"              description:""`
	UserName       string      `json:"user_name"       description:""`
	BusinessMan    string      `json:"business_man"    description:""`
	ProductName    string      `json:"product_name"    description:""`
	Money          int         `json:"money"           description:""`
	RequestMessage string      `json:"request_message" description:""`
	RejectMessage  string      `json:"reject_message"  description:""`
	Status         int         `json:"status"          description:"-1不通过，0审核中，1通过"`
	CreatedAt      *gtime.Time `json:"created_at"      description:""`
	UpdatedAt      *gtime.Time `json:"updated_at"      description:""`
}