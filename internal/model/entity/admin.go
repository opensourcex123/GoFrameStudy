// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure for table admin.
type Admin struct {
	Id        int         `json:"id"         description:""`
	UserName  string      `json:"user_name"  description:""`
	Password  string      `json:"password"   description:""`
	Avatar    string      `json:"avatar"     description:""`
	Income    int         `json:"income"     description:""`
	CreatedAt *gtime.Time `json:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updated_at" description:""`
}
