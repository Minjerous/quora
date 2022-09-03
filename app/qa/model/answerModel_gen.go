// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	answerFieldNames          = builder.RawFieldNames(&Answer{})
	answerRows                = strings.Join(answerFieldNames, ",")
	answerRowsExpectAutoSet   = strings.Join(stringx.Remove(answerFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	answerRowsWithPlaceHolder = strings.Join(stringx.Remove(answerFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheQuoraAnswerIdPrefix  = "cache:quora:answer:id:"
	cacheQuoraAnswerUidPrefix = "cache:quora:answer:uid:"
	cacheQuoraAnswerQidPrefix = "cache:quora:answer:qid:"
	ORDER                     = "order:"
	START                     = "start:"
	END                       = "end:"
)

type (
	answerModel interface {
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		Insert(ctx context.Context, data *Answer, session sqlx.Session) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Answer, error)
		FindByPid(ctx context.Context, page, pageSize int64, orderBy string, qid int64) ([]*Answer, error)
		Update(ctx context.Context, data *Answer) error
		Delete(ctx context.Context, id int64, session sqlx.Session) error
	}

	defaultAnswerModel struct {
		sqlc.CachedConn
		table string
	}

	Answer struct {
		Id          int64     `db:"id"`
		Uid         int64     `db:"uid"`
		Pid         int64     `db:"pid"`
		Type        string    `db:"type"`
		Value       string    `db:"value"`
		Likes       int64     `db:"likes"`
		Disagrees   int64     `db:"disagrees"`
		Agrees      int64     `db:"agrees"`
		Collections int64     `db:"collections"`
		Followers   int64     `db:"followers"`
		CreatedAt   time.Time `db:"created_at"`
		DeletedAt   time.Time `db:"deleted_at"`
		UpdateAt    time.Time `db:"update_at"`
	}
)

func newAnswerModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAnswerModel {
	return &defaultAnswerModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`answer`",
	}
}

func (m *defaultAnswerModel) Insert(ctx context.Context, data *Answer, session sqlx.Session) (sql.Result, error) {
	quoraAnswerIdKey := fmt.Sprintf("%s%v", cacheQuoraAnswerIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, answerRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Uid, data.Pid, data.Type, data.Value, data.Likes, data.Disagrees, data.Agrees, data.Collections, data.Followers, data.CreatedAt, data.DeletedAt, data.UpdateAt)
		}
		return conn.ExecCtx(ctx, query, data.Uid, data.Pid, data.Type, data.Value, data.Likes, data.Disagrees, data.Agrees, data.Collections, data.Followers, data.CreatedAt, data.DeletedAt, data.UpdateAt)
	}, quoraAnswerIdKey)
	return ret, err
}

func (m *defaultAnswerModel) FindOne(ctx context.Context, id int64) (*Answer, error) {
	quoraAnswerIdKey := fmt.Sprintf("%s%v", cacheQuoraAnswerIdPrefix, id)
	var resp Answer
	err := m.QueryRowCtx(ctx, &resp, quoraAnswerIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", answerRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAnswerModel) FindByPid(ctx context.Context, page, pageSize int64, orderBy string, qid int64) ([]*Answer, error) {
	quoraAnswerIdKey := fmt.Sprintf("%s%v", cacheQuoraAnswerQidPrefix, qid)

	if orderBy == "" {
		orderBy = "comments DESC"
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	var resp []*Answer
	err := m.QueryRowCtx(ctx, &resp, quoraAnswerIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `pid` = ? ORDER BY ? limit ? OFFSET ?", answerRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, qid, orderBy, pageSize, offset)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAnswerModel) Update(ctx context.Context, data *Answer) error {
	quoraAnswerIdKey := fmt.Sprintf("%s%v", cacheQuoraAnswerIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, answerRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uid, data.Pid, data.Type, data.Value, data.Likes, data.Disagrees, data.Agrees, data.Collections, data.Followers, data.CreatedAt, data.DeletedAt, data.UpdateAt, data.Id)
	}, quoraAnswerIdKey)
	return err
}

func (m *defaultAnswerModel) Delete(ctx context.Context, id int64, session sqlx.Session) error {
	quoraAnswerIdKey := fmt.Sprintf("%s%v", cacheQuoraAnswerIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session == nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, quoraAnswerIdKey)
	return err
}

func (m *defaultAnswerModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheQuoraAnswerIdPrefix, primary)
}

func (m *defaultAnswerModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", answerRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAnswerModel) tableName() string {
	return m.table
}

func (m *defaultAnswerModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}
