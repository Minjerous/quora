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
	questionFieldNames          = builder.RawFieldNames(&Question{})
	questionRows                = strings.Join(questionFieldNames, ",")
	questionRowsExpectAutoSet   = strings.Join(stringx.Remove(questionFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	questionRowsWithPlaceHolder = strings.Join(stringx.Remove(questionFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheQuoraQuestionIdPrefix = "cache:quora:question:id:"
)

type (
	questionModel interface {
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		Insert(ctx context.Context, data *Question) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Question, error)
		Update(ctx context.Context, data *Question) error
		AddQuestionNumber(ctx context.Context, qid int64, Type string, session sqlx.Session) error
		SubQuestionNumber(ctx context.Context, qid int64, Type string, session sqlx.Session) error
		Delete(ctx context.Context, id int64) error
	}

	defaultQuestionModel struct {
		sqlc.CachedConn
		table string
	}

	Question struct {
		Id          int64     `db:"id"`
		Uid         int64     `db:"uid"`
		Topic       string    `db:"topic"`
		Title       string    `db:"title"`
		Value       string    `db:"value"`
		Comments    int64     `db:"comments"`
		Collections int64     `db:"collections"`
		Likes       int64     `db:"likes"`
		Pageviews   int64     `db:"pageviews"`
		Followers   int64     `db:"followers"`
		CreatedAt   time.Time `db:"created_at"`
		DeletedAt   time.Time `db:"deleted_at"`
		UpdateAt    time.Time `db:"update_at"`
	}
)

func newQuestionModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultQuestionModel {
	return &defaultQuestionModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`question`",
	}
}

func (m *defaultQuestionModel) Insert(ctx context.Context, data *Question) (sql.Result, error) {
	quoraQuestionIdKey := fmt.Sprintf("%s%v", cacheQuoraQuestionIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, questionRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uid, data.Topic, data.Title, data.Value, data.Comments, data.Collections, data.Likes, data.Pageviews, data.Followers, data.CreatedAt, data.DeletedAt, data.UpdateAt)
	}, quoraQuestionIdKey)
	return ret, err
}

func (m *defaultQuestionModel) FindOne(ctx context.Context, id int64) (*Question, error) {
	quoraQuestionIdKey := fmt.Sprintf("%s%v", cacheQuoraQuestionIdPrefix, id)
	var resp Question
	err := m.QueryRowCtx(ctx, &resp, quoraQuestionIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", questionRows, m.table)
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

func (m *defaultQuestionModel) Update(ctx context.Context, data *Question) error {
	quoraQuestionIdKey := fmt.Sprintf("%s%v", cacheQuoraQuestionIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, questionRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uid, data.Topic, data.Title, data.Value, data.Comments, data.Collections, data.Likes, data.Pageviews, data.Followers, data.CreatedAt, data.DeletedAt, data.UpdateAt, data.Id)
	}, quoraQuestionIdKey)
	return err
}

func (m *defaultQuestionModel) Delete(ctx context.Context, id int64) error {
	quoraQuestionIdKey := fmt.Sprintf("%s%v", cacheQuoraQuestionIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, quoraQuestionIdKey)
	return err
}

func (m *defaultQuestionModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheQuoraQuestionIdPrefix, primary)
}

func (m *defaultQuestionModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", questionRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultQuestionModel) tableName() string {
	return m.table
}

func (m *defaultQuestionModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultQuestionModel) AddQuestionNumber(ctx context.Context, qid int64, Type string, session sqlx.Session) error {
	quoraQuestionIdKey := fmt.Sprintf("%s%v", cacheQuoraQuestionIdPrefix, qid)
	questionUpdateRows := fmt.Sprintf("%s=%s+1", Type, Type)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, questionUpdateRows)
		return conn.ExecCtx(ctx, query)
	}, quoraQuestionIdKey)
	return err
}

func (m *defaultQuestionModel) SubQuestionNumber(ctx context.Context, qid int64, Type string, session sqlx.Session) error {
	quoraQuestionIdKey := fmt.Sprintf("%s%v", cacheQuoraQuestionIdPrefix, qid)
	questionUpdateRows := fmt.Sprintf("%s=%s-1", Type, Type)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, questionUpdateRows)
		return conn.ExecCtx(ctx, query)
	}, quoraQuestionIdKey)
	return err
}