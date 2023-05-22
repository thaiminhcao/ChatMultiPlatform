package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TourModel = (*customTourModel)(nil)

type (
	// TourModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTourModel.
	TourModel interface {
		tourModel
	}

	customTourModel struct {
		*defaultTourModel
	}
)

// NewTourModel returns a model for the database table.
func NewTourModel(conn sqlx.SqlConn) TourModel {
	return &customTourModel{
		defaultTourModel: newTourModel(conn),
	}
}
