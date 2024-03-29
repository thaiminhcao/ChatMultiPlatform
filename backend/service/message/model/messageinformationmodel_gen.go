// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	messageInformationFieldNames          = builder.RawFieldNames(&MessageInformation{})
	messageInformationRows                = strings.Join(messageInformationFieldNames, ",")
	messageInformationRowsExpectAutoSet   = strings.Join(stringx.Remove(messageInformationFieldNames, "`message_id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	messageInformationRowsWithPlaceHolder = strings.Join(stringx.Remove(messageInformationFieldNames, "`message_id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"
)

type (
	messageInformationModel interface {
		Insert(ctx context.Context, data *MessageInformation) (sql.Result, error)
		FindOne(ctx context.Context, messageId int64) (*MessageInformation, error)
		Update(ctx context.Context, data *MessageInformation) error
		Delete(ctx context.Context, messageId int64) error
	}

	defaultMessageInformationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	MessageInformation struct {
		MessageId int64          `db:"message_id"`
		Title     sql.NullString `db:"title"`
		UserId    sql.NullInt64  `db:"user_id"`
		Broker    sql.NullString `db:"broker"`
		CreatedAt sql.NullInt64  `db:"created_at"`
	}
)

func newMessageInformationModel(conn sqlx.SqlConn) *defaultMessageInformationModel {
	return &defaultMessageInformationModel{
		conn:  conn,
		table: "`message_information`",
	}
}

func (m *defaultMessageInformationModel) Delete(ctx context.Context, messageId int64) error {
	query := fmt.Sprintf("delete from %s where `message_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, messageId)
	return err
}

func (m *defaultMessageInformationModel) FindOne(ctx context.Context, messageId int64) (*MessageInformation, error) {
	query := fmt.Sprintf("select %s from %s where `message_id` = ? limit 1", messageInformationRows, m.table)
	var resp MessageInformation
	err := m.conn.QueryRowCtx(ctx, &resp, query, messageId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMessageInformationModel) Insert(ctx context.Context, data *MessageInformation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, messageInformationRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Title, data.UserId, data.Broker)
	return ret, err
}

func (m *defaultMessageInformationModel) Update(ctx context.Context, data *MessageInformation) error {
	query := fmt.Sprintf("update %s set %s where `message_id` = ?", m.table, messageInformationRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Title, data.UserId, data.Broker, data.MessageId)
	return err
}

func (m *defaultMessageInformationModel) tableName() string {
	return m.table
}
