// Code generated by sqlgen. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// UserModel represents a user model.
type UserModel struct {
	db gorm.DB
}

// User represents a user struct data.
type User struct {
	Id         uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Password   string    `gorm:"column:password" json:"password"`
	Mobile     string    `gorm:"column:mobile" json:"mobile"`
	Gender     string    `gorm:"column:gender" json:"gender"`
	Nickname   string    `gorm:"column:nickname" json:"nickname"`
	Type       int8      `gorm:"column:type" json:"type"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

// DeleteWhereParameter is a where parameter structure.
type DeleteWhereParameter struct {
	IdEqual uint64
}

// DeleteByNameWhereParameter is a where parameter structure.
type DeleteByNameWhereParameter struct {
	NameEqual string
}

// DeleteByNameAndMobileWhereParameter is a where parameter structure.
type DeleteByNameAndMobileWhereParameter struct {
	NameEqual   string
	MobileEqual string
}

// DeleteOrderByIDWhereParameter is a where parameter structure.
type DeleteOrderByIDWhereParameter struct {
	IdEqual uint64
}

// DeleteOrderByIDLimitWhereParameter is a where parameter structure.
type DeleteOrderByIDLimitWhereParameter struct {
	IdEqual uint64
}

// DeleteOrderByIDLimitLimitParameter is a limit parameter structure.
type DeleteOrderByIDLimitLimitParameter struct {
	Count int
}

// TableName returns the table name. it implemented by gorm.Tabler.
func (User) TableName() string {
	return "user"
}

// NewUserModel returns a new user model.
func NewUserModel(db gorm.DB) *UserModel {
	return &UserModel{db: db}
}

// Create creates  user data.
func (m *UserModel) Create(ctx context.Context, data ...*User) error {
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}

	db := m.db.WithContext(ctx)
	var list []User
	for _, v := range data {
		list = append(list, *v)
	}

	return db.Create(&list).Error
}

// Delete is generated from sql:
// delete from user where id = ?;
func (m *UserModel) Delete(ctx context.Context, where DeleteWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db.Where(`id = ?`, where.IdEqual)
	db.Delete(&User{})
	return db.Error
}

// DeleteByName is generated from sql:
// delete from user where name = ?;
func (m *UserModel) DeleteByName(ctx context.Context, where DeleteByNameWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db.Where(`name = ?`, where.NameEqual)
	db.Delete(&User{})
	return db.Error
}

// DeleteByNameAndMobile is generated from sql:
// delete from user where name = ? and mobile = ?;
func (m *UserModel) DeleteByNameAndMobile(ctx context.Context, where DeleteByNameAndMobileWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db.Where(`name = ? AND mobile = ?`, where.NameEqual, where.MobileEqual)
	db.Delete(&User{})
	return db.Error
}

// DeleteOrderByID is generated from sql:
// delete from user where id = ? order by id desc;
func (m *UserModel) DeleteOrderByID(ctx context.Context, where DeleteOrderByIDWhereParameter) error {
	var db = m.db.WithContext(ctx)
	db.Where(`id = ?`, where.IdEqual)
	db.Order(`id desc`)
	db.Delete(&User{})
	return db.Error
}

// DeleteOrderByIDLimit is generated from sql:
// delete from user where id = ? order by id desc limit 10;
func (m *UserModel) DeleteOrderByIDLimit(ctx context.Context, where DeleteOrderByIDLimitWhereParameter, limit DeleteOrderByIDLimitLimitParameter) error {
	var db = m.db.WithContext(ctx)
	db.Where(`id = ?`, where.IdEqual)
	db.Order(`id desc`)
	db.Limit(limit.Count)
	db.Delete(&User{})
	return db.Error
}
