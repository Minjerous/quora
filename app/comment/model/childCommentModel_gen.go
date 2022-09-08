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
	childCommentFieldNames          = builder.RawFieldNames(&ChildComment{})
	childCommentRows                = strings.Join(childCommentFieldNames, ",")
	childCommentRowsExpectAutoSet   = strings.Join(stringx.Remove(childCommentFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	childCommentRowsWithPlaceHolder = strings.Join(stringx.Remove(childCommentFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheQuoraChildCommentIdPrefix = "cache:quora:childComment:id:"
)

type (
	childCommentModel interface {
		Insert(ctx context.Context, data *ChildComment) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ChildComment, error)
		FindByPid(ctx context.Context, page, pageSize int64, orderBy string, qid int64) ([]*ChildComment, error)
		Update(ctx context.Context, data *ChildComment) error
		Delete(ctx context.Context, id int64) error
	}

	defaultChildCommentModel struct {
		sqlc.CachedConn
		table string
	}

	ChildComment struct {
		Id       int64     `db:"id"`
		Pid      int64     `db:"pid"`
		Content  string    `db:"content"`
		Uid      int64     `db:"uid"`
		Likes    int64     `db:"likes"`
		CreateAt time.Time `db:"create_at"`
		ReplyUid int64     `db:"reply_uid"`
		Ip       string    `db:"ip"`
		ReplyCid int64     `db:"reply_cid"`
	}
)

func newChildCommentModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultChildCommentModel {
	return &defaultChildCommentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`child_comment`",
	}
}

func (m *defaultChildCommentModel) Insert(ctx context.Context, data *ChildComment) (sql.Result, error) {
	quoraChildCommentIdKey := fmt.Sprintf("%s%v", cacheQuoraChildCommentIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, childCommentRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Pid, data.Content, data.Uid, data.Likes, data.CreateAt, data.ReplyUid, data.Ip, data.ReplyCid)
	}, quoraChildCommentIdKey)
	return ret, err
}

func (m *defaultChildCommentModel) FindOne(ctx context.Context, id int64) (*ChildComment, error) {
	quoraChildCommentIdKey := fmt.Sprintf("%s%v", cacheQuoraChildCommentIdPrefix, id)
	var resp ChildComment
	err := m.QueryRowCtx(ctx, &resp, quoraChildCommentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", childCommentRows, m.table)
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

func (m *defaultChildCommentModel) FindByPid(ctx context.Context, page, pageSize int64, orderBy string, qid int64) ([]*ChildComment, error) {
	quoraAnswerIdKey := fmt.Sprintf("%s%v", cacheCommentQidPrefix, qid)

	if orderBy == "" {
		orderBy = "likes DESC"
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	var resp []*ChildComment
	err := m.QueryRowCtx(ctx, &resp, quoraAnswerIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `pid` = ? ORDER BY ? limit ? OFFSET ?", childCommentRows, m.table)
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

func (m *defaultChildCommentModel) Update(ctx context.Context, data *ChildComment) error {
	quoraChildCommentIdKey := fmt.Sprintf("%s%v", cacheQuoraChildCommentIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, childCommentRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Pid, data.Content, data.Uid, data.Likes, data.CreateAt, data.ReplyUid, data.Ip, data.ReplyCid, data.Id)
	}, quoraChildCommentIdKey)
	return err
}

func (m *defaultChildCommentModel) Delete(ctx context.Context, id int64) error {
	quoraChildCommentIdKey := fmt.Sprintf("%s%v", cacheQuoraChildCommentIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, quoraChildCommentIdKey)
	return err
}

func (m *defaultChildCommentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheQuoraChildCommentIdPrefix, primary)
}

func (m *defaultChildCommentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", childCommentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultChildCommentModel) tableName() string {
	return m.table
}
