// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReturnRequestDao is the data access object for table return_request.
type ReturnRequestDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns ReturnRequestColumns // columns contains all the column names of Table for convenient usage.
}

// ReturnRequestColumns defines and stores column names for table return_request.
type ReturnRequestColumns struct {
	Id             string //
	UserName       string //
	BusinessMan    string //
	ProductName    string //
	Money          string //
	RequestMessage string //
	RejectMessage  string //
	Status         string // -1不通过，0审核中，1通过
	CreatedAt      string //
	UpdatedAt      string //
}

//  returnRequestColumns holds the columns for table return_request.
var returnRequestColumns = ReturnRequestColumns{
	Id:             "id",
	UserName:       "user_name",
	BusinessMan:    "business_man",
	ProductName:    "product_name",
	Money:          "money",
	RequestMessage: "request_message",
	RejectMessage:  "reject_message",
	Status:         "status",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewReturnRequestDao creates and returns a new DAO object for table data access.
func NewReturnRequestDao() *ReturnRequestDao {
	return &ReturnRequestDao{
		group:   "default",
		table:   "return_request",
		columns: returnRequestColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReturnRequestDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReturnRequestDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReturnRequestDao) Columns() ReturnRequestColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReturnRequestDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReturnRequestDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReturnRequestDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
