package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MerchantsModel = (*customMerchantsModel)(nil)

type (
	// MerchantsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMerchantsModel.
	MerchantsModel interface {
		merchantsModel
		withSession(session sqlx.Session) MerchantsModel

		// 自定义方法
		FindList(ctx context.Context, page, pageSize int32, verifiedStatus, merchantStatus string) ([]*Merchants, error)
		Count(ctx context.Context, verifiedStatus, merchantStatus string) (int64, error)
	}

	customMerchantsModel struct {
		*defaultMerchantsModel
	}
)

// NewMerchantsModel returns a model for the database table.
func NewMerchantsModel(conn sqlx.SqlConn) MerchantsModel {
	return &customMerchantsModel{
		defaultMerchantsModel: newMerchantsModel(conn),
	}
}

func (m *customMerchantsModel) withSession(session sqlx.Session) MerchantsModel {
	return NewMerchantsModel(sqlx.NewSqlConnFromSession(session))
}

// FindList 查询商户列表
func (m *customMerchantsModel) FindList(ctx context.Context, page, pageSize int32, verifiedStatus, merchantStatus string) ([]*Merchants, error) {
	// 构建查询条件
	var conditions []string
	var args []interface{}

	if verifiedStatus != "" {
		conditions = append(conditions, "`verified_status` = ?")
		args = append(args, verifiedStatus)
	}

	if merchantStatus != "" {
		conditions = append(conditions, "`merchant_status` = ?")
		args = append(args, merchantStatus)
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 构建查询SQL
	query := fmt.Sprintf("SELECT %s FROM %s %s ORDER BY `created_at` DESC LIMIT ? OFFSET ?",
		merchantsRows, m.table, whereClause)

	args = append(args, pageSize, offset)

	var merchants []*Merchants
	err := m.conn.QueryRowsCtx(ctx, &merchants, query, args...)
	if err != nil {
		return nil, err
	}

	return merchants, nil
}

// Count 统计商户数量
func (m *customMerchantsModel) Count(ctx context.Context, verifiedStatus, merchantStatus string) (int64, error) {
	// 构建查询条件
	var conditions []string
	var args []interface{}

	if verifiedStatus != "" {
		conditions = append(conditions, "`verified_status` = ?")
		args = append(args, verifiedStatus)
	}

	if merchantStatus != "" {
		conditions = append(conditions, "`merchant_status` = ?")
		args = append(args, merchantStatus)
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// 构建统计SQL
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s %s", m.table, whereClause)

	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query, args...)
	if err != nil {
		return 0, err
	}

	return count, nil
}
