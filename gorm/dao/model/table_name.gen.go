// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTableName = "table_name"

// TableName mapped from table <table_name>
type TableName struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	Name       string    `gorm:"column:name" json:"name"`
}

// TableName TableName's table name
func (*TableName) TableName() string {
	return TableNameTableName
}
