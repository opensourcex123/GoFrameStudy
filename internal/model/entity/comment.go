// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Comment is the golang structure for table comment.
type Comment struct {
	Id        int         `json:"id"         description:""`
	UserId    int         `json:"user_id"    description:""`
	ProductId int         `json:"product_id" description:""`
	Comment   string      `json:"comment"    description:""`
	Type      int         `json:"type"       description:"0是差评，1是好评"`
	CreatedAt *gtime.Time `json:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updated_at" description:""`
}