package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var _ UsersModel = (*customUsersModel)(nil)
var (
	usersRowsExpectAutoSet2 = strings.Join(stringx.Remove(usersFieldNames), ",")
)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		FindByUserName(ctx context.Context, name string) (*Users, error)
		Insertvalues(ctx context.Context, data *Users) (sql.Result, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn),
	}
}

func (m *defaultUsersModel) FindByUserName(ctx context.Context, name string) (*Users, error) {
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", usersRows, m.table)
	var resp Users
	err := m.conn.QueryRowCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
func (m *defaultUsersModel) Insertvalues(ctx context.Context, data *Users) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet2)
	logx.Info(usersRows)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Username, data.Email, data.Password, data.Gender, data.Dob, data.CreatedAt)
	return ret, err
}
