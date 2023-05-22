package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MessageInformationModel = (*customMessageInformationModel)(nil)

type (
	// MessageInformationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageInformationModel.
	MessageInformationModel interface {
		messageInformationModel
	}

	customMessageInformationModel struct {
		*defaultMessageInformationModel
	}
)

// NewMessageInformationModel returns a model for the database table.
func NewMessageInformationModel(conn sqlx.SqlConn) MessageInformationModel {
	return &customMessageInformationModel{
		defaultMessageInformationModel: newMessageInformationModel(conn),
	}
}
