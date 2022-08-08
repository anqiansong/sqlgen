// Code generated by sqlgen. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"xorm.io/xorm"
)

// UserModel represents a user model.
type UserModel struct {
	engine *xorm.Engine
}

// User represents a user struct data.
type User struct {
	Id         uint64    `xorm:"pk autoincr 'id'" json:"id"`
	Name       string    `xorm:"'name'" json:"name"`
	Password   string    `xorm:"'password'" json:"password"`
	Mobile     string    `xorm:"'mobile'" json:"mobile"`
	Gender     string    `xorm:"'gender'" json:"gender"`
	Nickname   string    `xorm:"'nickname'" json:"nickname"`
	Type       int8      `xorm:"'type'" json:"type"`
	CreateTime time.Time `xorm:"'create_time'" json:"createTime"`
	UpdateTime time.Time `xorm:"'update_time'" json:"updateTime"`
}

// UpdateWhereParameter is a where parameter structure.
type UpdateWhereParameter struct {
	Id uint64
}

// UpdateByNameWhereParameter is a where parameter structure.
type UpdateByNameWhereParameter struct {
	Name string
}

// UpdatePartWhereParameter is a where parameter structure.
type UpdatePartWhereParameter struct {
	Id uint64
}

// UpdatePartByNameWhereParameter is a where parameter structure.
type UpdatePartByNameWhereParameter struct {
	Name string
}

// UpdateNameLimitWhereParameter is a where parameter structure.
type UpdateNameLimitWhereParameter struct {
	Id uint64
}

// UpdateNameLimitLimitParameter is a limit parameter structure.
type UpdateNameLimitLimitParameter struct {
	Count int
}

// UpdateNameLimitOrderWhereParameter is a where parameter structure.
type UpdateNameLimitOrderWhereParameter struct {
	Id uint64
}

// UpdateNameLimitOrderLimitParameter is a limit parameter structure.
type UpdateNameLimitOrderLimitParameter struct {
	Count int
}

func (User) TableName() string {
	return "user"
}

// Insert creates  user data.
func (m *UserModel) Insert(ctx context.Context, data ...*User) error {
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}

	var session = m.engine.Context(ctx)
	var list []User
	for _, v := range data {
		list = append(list, *v)
	}

	_, err := session.Insert(&list)
	return err
}

// Update is generated from sql:
// update user set name = ?, password = ?, mobile = ?, gender = ?, nickname = ?, type = ?, create_time = ?, update_time = ? where id = ?;
func (m *UserModel) Update(ctx context.Context, data *User, where UpdateWhereParameter) error {
	var session = m.engine.Context(ctx)
	session.Table(&User{})
	session.Where(`id = ?`, where.Id)
	_, err := session.Update(map[string]interface{}{
		"name":        data.Name,
		"password":    data.Password,
		"mobile":      data.Mobile,
		"gender":      data.Gender,
		"nickname":    data.Nickname,
		"type":        data.Type,
		"create_time": data.CreateTime,
		"update_time": data.UpdateTime,
	})
	return err
}

// UpdateByName is generated from sql:
// update user set password = ?, mobile = ?, gender = ?, nickname = ?, type = ?, create_time = ?, update_time = ? where name = ?;
func (m *UserModel) UpdateByName(ctx context.Context, data *User, where UpdateByNameWhereParameter) error {
	var session = m.engine.Context(ctx)
	session.Table(&User{})
	session.Where(`name = ?`, where.Name)
	_, err := session.Update(map[string]interface{}{
		"password":    data.Password,
		"mobile":      data.Mobile,
		"gender":      data.Gender,
		"nickname":    data.Nickname,
		"type":        data.Type,
		"create_time": data.CreateTime,
		"update_time": data.UpdateTime,
	})
	return err
}

// UpdatePart is generated from sql:
// update user set name = ?, nickname = ? where id = ?;
func (m *UserModel) UpdatePart(ctx context.Context, data *User, where UpdatePartWhereParameter) error {
	var session = m.engine.Context(ctx)
	session.Table(&User{})
	session.Where(`id = ?`, where.Id)
	_, err := session.Update(map[string]interface{}{
		"name":     data.Name,
		"nickname": data.Nickname,
	})
	return err
}

// UpdatePartByName is generated from sql:
// update user set name = ?, nickname = ? where name = ?;
func (m *UserModel) UpdatePartByName(ctx context.Context, data *User, where UpdatePartByNameWhereParameter) error {
	var session = m.engine.Context(ctx)
	session.Table(&User{})
	session.Where(`name = ?`, where.Name)
	_, err := session.Update(map[string]interface{}{
		"name":     data.Name,
		"nickname": data.Nickname,
	})
	return err
}

// UpdateNameLimit is generated from sql:
// update user set name = ? where id > ? limit ?;
func (m *UserModel) UpdateNameLimit(ctx context.Context, data *User, where UpdateNameLimitWhereParameter, limit UpdateNameLimitLimitParameter) error {
	var session = m.engine.Context(ctx)
	session.Table(&User{})
	session.Where(`id > ?`, where.Id)
	session.Limit(limit.Count)
	_, err := session.Update(map[string]interface{}{
		"name": data.Name,
	})
	return err
}

// UpdateNameLimitOrder is generated from sql:
// update user set name = ? where id > ? order by id desc limit ?;
func (m *UserModel) UpdateNameLimitOrder(ctx context.Context, data *User, where UpdateNameLimitOrderWhereParameter, limit UpdateNameLimitOrderLimitParameter) error {
	var session = m.engine.Context(ctx)
	session.Table(&User{})
	session.Where(`id > ?`, where.Id)
	session.OrderBy(`id desc`)
	session.Limit(limit.Count)
	_, err := session.Update(map[string]interface{}{
		"name": data.Name,
	})
	return err
}