// Code generated by sqlgen. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

// UserModel represents a user model.
type UserModel struct {
	db bun.IDB
}

// User represents a user struct data.
type User struct {
	bun.BaseModel `bun:"table:user"`
	Id            uint64    `bun:"id,pk,autoincrement;" json:"id"`
	Name          string    `bun:"name" json:"name"`
	Password      string    `bun:"password" json:"password"`
	Mobile        string    `bun:"mobile" json:"mobile"`
	Gender        string    `bun:"gender" json:"gender"`
	Nickname      string    `bun:"nickname" json:"nickname"`
	Type          int8      `bun:"type" json:"type"`
	CreateTime    time.Time `bun:"create_time" json:"createTime"`
	UpdateTime    time.Time `bun:"update_time" json:"updateTime"`
}

// UpdateWhereParameter is a where parameter structure.
type UpdateWhereParameter struct {
	IdEqual uint64
}

// UpdateByNameWhereParameter is a where parameter structure.
type UpdateByNameWhereParameter struct {
	NameEqual string
}

// UpdatePartWhereParameter is a where parameter structure.
type UpdatePartWhereParameter struct {
	IdEqual uint64
}

// UpdatePartByNameWhereParameter is a where parameter structure.
type UpdatePartByNameWhereParameter struct {
	NameEqual string
}

// UpdateNameLimitWhereParameter is a where parameter structure.
type UpdateNameLimitWhereParameter struct {
	IdGT uint64
}

// UpdateNameLimitOrderWhereParameter is a where parameter structure.
type UpdateNameLimitOrderWhereParameter struct {
	IdGT uint64
}

// NewUserModel creates a new user model.
func NewUserModel(db bun.IDB) *UserModel {
	return &UserModel{
		db: db,
	}
}

// Create creates  user data.
func (m *UserModel) Create(ctx context.Context, data ...*User) error {
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}

	var list []User
	for _, v := range data {
		list = append(list, *v)
	}

	_, err := m.db.NewInsert().Model(&list).Exec(ctx)
	return err
}

// Update is generated from sql:
// update user set name = ?, password = ?, mobile = ?, gender = ?, nickname = ?, type = ?, create_time = ?, update_time = ? where id = ?;
func (m *UserModel) Update(ctx context.Context, data *User, where UpdateWhereParameter) error {
	var db = m.db.NewUpdate()
	db.Model(map[string]interface{}{
		"name":        data.Name,
		"password":    data.Password,
		"mobile":      data.Mobile,
		"gender":      data.Gender,
		"nickname":    data.Nickname,
		"type":        data.Type,
		"create_time": data.CreateTime,
		"update_time": data.UpdateTime,
	})
	db.Where(`id = ?`, where.IdEqual)
	_, err := db.Exec(ctx)
	return err
}

// UpdateByName is generated from sql:
// update user set password = ?, mobile = ?, gender = ?, nickname = ?, type = ?, create_time = ?, update_time = ? where name = ?;
func (m *UserModel) UpdateByName(ctx context.Context, data *User, where UpdateByNameWhereParameter) error {
	var db = m.db.NewUpdate()
	db.Model(map[string]interface{}{
		"password":    data.Password,
		"mobile":      data.Mobile,
		"gender":      data.Gender,
		"nickname":    data.Nickname,
		"type":        data.Type,
		"create_time": data.CreateTime,
		"update_time": data.UpdateTime,
	})
	db.Where(`name = ?`, where.NameEqual)
	_, err := db.Exec(ctx)
	return err
}

// UpdatePart is generated from sql:
// update user set name = ?, nickname = ? where id = ?;
func (m *UserModel) UpdatePart(ctx context.Context, data *User, where UpdatePartWhereParameter) error {
	var db = m.db.NewUpdate()
	db.Model(map[string]interface{}{
		"name":     data.Name,
		"nickname": data.Nickname,
	})
	db.Where(`id = ?`, where.IdEqual)
	_, err := db.Exec(ctx)
	return err
}

// UpdatePartByName is generated from sql:
// update user set name = ?, nickname = ? where name = ?;
func (m *UserModel) UpdatePartByName(ctx context.Context, data *User, where UpdatePartByNameWhereParameter) error {
	var db = m.db.NewUpdate()
	db.Model(map[string]interface{}{
		"name":     data.Name,
		"nickname": data.Nickname,
	})
	db.Where(`name = ?`, where.NameEqual)
	_, err := db.Exec(ctx)
	return err
}

// UpdateNameLimit is generated from sql:
// update user set name = ? where id > ? limit ?;
func (m *UserModel) UpdateNameLimit(ctx context.Context, data *User, where UpdateNameLimitWhereParameter) error {
	var db = m.db.NewUpdate()
	db.Model(map[string]interface{}{
		"name": data.Name,
	})
	db.Where(`id > ?`, where.IdGT)
	_, err := db.Exec(ctx)
	return err
}

// UpdateNameLimitOrder is generated from sql:
// update user set name = ? where id > ? order by id desc limit ?;
func (m *UserModel) UpdateNameLimitOrder(ctx context.Context, data *User, where UpdateNameLimitOrderWhereParameter) error {
	var db = m.db.NewUpdate()
	db.Model(map[string]interface{}{
		"name": data.Name,
	})
	db.Where(`id > ?`, where.IdGT)
	_, err := db.Exec(ctx)
	return err
}
