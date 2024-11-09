// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"golodge/deploy/script/mysql/genModel"
	"strings"

	"time"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"golodge/common/globalkey"
)

var (
	homestayFieldNames          = builder.RawFieldNames(&Homestay{})
	homestayRows                = strings.Join(homestayFieldNames, ",")
	homestayRowsExpectAutoSet   = strings.Join(stringx.Remove(homestayFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	homestayRowsWithPlaceHolder = strings.Join(stringx.Remove(homestayFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheLooklookTravelHomestayIdPrefix = "cache:looklookTravel:homestay:id:"
)

type (
	homestayModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *Homestay) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Homestay, error)
		Update(ctx context.Context, session sqlx.Session, data *Homestay) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *Homestay) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *Homestay) error
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Homestay, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Homestay, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Homestay, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Homestay, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Homestay, error)
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultHomestayModel struct {
		sqlc.CachedConn
		table string
	}

	//Homestay struct {
	//	Id                  int64     `db:"id"`
	//	CreateTime          time.Time `db:"create_time"`
	//	UpdateTime          time.Time `db:"update_time"`
	//	DeleteTime          time.Time `db:"delete_time"`
	//	DelState            int64     `db:"del_state"`
	//	Version             int64     `db:"version"`               // 版本号
	//	Title               string    `db:"title"`                 // 标题
	//	SubTitle            string    `db:"sub_title"`             // 副标题
	//	Banner              string    `db:"banner"`                // 轮播图，第一张封面
	//	Info                string    `db:"info"`                  // 介绍
	//	PeopleNum           int64     `db:"people_num"`            // 容纳人的数量
	//	HomestayBusinessId  int64     `db:"homestay_business_id"`  // 民宿店铺id
	//	UserId              int64     `db:"user_id"`               // 房东id，冗余字段
	//	RowState            int64     `db:"row_state"`             // 0:下架 1:上架
	//	RowType             int64     `db:"row_type"`              // 售卖类型0：按房间出售 1:按人次出售
	//	FoodInfo            string    `db:"food_info"`             // 餐食标准
	//	FoodPrice           int64     `db:"food_price"`            // 餐食价格（分）
	//	HomestayPrice       int64     `db:"homestay_price"`        // 民宿价格（分）
	//	MarketHomestayPrice int64     `db:"market_homestay_price"` // 民宿市场价格（分）
	//}
	Homestay struct {
		Id                 int64     `db:"id"`
		CreateTime         time.Time `db:"create_time"`
		UpdateTime         time.Time `db:"update_time"`
		DeleteTime         time.Time `db:"delete_time"`
		DelState           int64     `db:"del_state"`
		Version            int64     `db:"version"`      // 版本号
		Title              string    `db:"title"`        // 标题
		RatingStars        float64   `db:"rating_stars"` // 评分
		BannerUrls         string    `db:"banner_urls"`
		TitleTags          string    `db:"title_tags"`
		CommentCount       int64     `db:"comment_count"`
		Location           string    `db:"location"` // 位置
		TypeId             int64     `db:"type_id"`
		Longitude          float64   `db:"longitude"`
		Latitude           float64   `db:"latitude"`
		Facilities         string    `db:"facilities"`
		Cover              string    `db:"cover"` // 轮播图，第一张封面
		Area               string    `db:"area"`
		RoomConfig         string    `db:"room_config"` // 介绍
		CleanVideo         string    `db:"clean_video"`
		HomestayBusinessId int64     `db:"homestay_business_id"` // 民宿店铺id
		HostId             int64     `db:"host_id"`              // 房东id，冗余字段
		HostAvatar         string    `db:"host_avatar"`
		HostNickname       string    `db:"host_nickname"`
		RowState           int64     `db:"row_state"`    // 0:下架 1:上架
		PriceBefore        int64     `db:"price_before"` // 民宿价格（分）
		PriceAfter         int64     `db:"price_after"`
	}
)

func newHomestayModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultHomestayModel {
	return &defaultHomestayModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`homestay`",
	}
}

func (m *defaultHomestayModel) Insert(ctx context.Context, session sqlx.Session, data *Homestay) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	data.DelState = globalkey.DelStateNo
	looklookTravelHomestayIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?)", m.table, homestayRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Title, data.RatingStars, data.BannerUrls, data.TitleTags, data.CommentCount, data.Location, data.TypeId, data.Latitude, data.Longitude, data.Facilities, data.Cover, data.Area, data.RoomConfig, data.CleanVideo, data.HomestayBusinessId, data.HostId, data.HostAvatar, data.HostNickname, data.RowState, data.PriceBefore, data.PriceAfter)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Title, data.RatingStars, data.BannerUrls, data.TitleTags, data.CommentCount, data.Location, data.Latitude, data.Longitude, data.Facilities, data.Cover, data.Area, data.RoomConfig, data.CleanVideo, data.HomestayBusinessId, data.HostId, data.HostAvatar, data.HostNickname, data.RowState, data.PriceBefore, data.PriceAfter)
	}, looklookTravelHomestayIdKey)
}

func (m *defaultHomestayModel) FindOne(ctx context.Context, id int64) (*Homestay, error) {
	looklookTravelHomestayIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayIdPrefix, id)
	var resp Homestay
	err := m.QueryRowCtx(ctx, &resp, looklookTravelHomestayIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", homestayRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, genModel.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHomestayModel) Update(ctx context.Context, session sqlx.Session, data *Homestay) (sql.Result, error) {
	looklookTravelHomestayIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, homestayRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Title, data.RatingStars, data.BannerUrls, data.TitleTags, data.CommentCount, data.Location, data.TypeId, data.Latitude, data.Longitude, data.Facilities, data.Cover, data.Area, data.RoomConfig, data.CleanVideo, data.HomestayBusinessId, data.HostId, data.HostAvatar, data.HostNickname, data.RowState, data.PriceBefore, data.PriceAfter, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Title, data.RatingStars, data.BannerUrls, data.TitleTags, data.CommentCount, data.Location, data.TypeId, data.Latitude, data.Longitude, data.Facilities, data.Cover, data.Area, data.RoomConfig, data.CleanVideo, data.HomestayBusinessId, data.HostId, data.HostAvatar, data.HostNickname, data.RowState, data.PriceBefore, data.PriceAfter, data.Id)
	}, looklookTravelHomestayIdKey)
}

func (m *defaultHomestayModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *Homestay) error {

	// bug 这里误注释了
	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	looklookTravelHomestayIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayIdPrefix, data.Id)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, homestayRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Title, data.RatingStars, data.BannerUrls, data.TitleTags, data.CommentCount, data.Location, data.TypeId, data.Latitude, data.Longitude, data.Facilities, data.Cover, data.Area, data.RoomConfig, data.CleanVideo, data.HomestayBusinessId, data.HostId, data.HostAvatar, data.HostNickname, data.RowState, data.PriceBefore, data.PriceAfter, data.Id, oldVersion)
		}
		// bug最后得带data.id和oldVersion，不然会报参数无法满足
		// ctx, query, data.DeleteTime, data.DelState, data.Version, data.Title, data.SubTitle, data.Banner, data.Info, data.PeopleNum, data.HomestayBusinessId, data.UserId, data.RowState, data.RowType, data.FoodInfo, data.FoodPrice, data.HomestayPrice, data.MarketHomestayPrice, data.Id, oldVersion
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Title, data.RatingStars, data.BannerUrls, data.TitleTags, data.CommentCount, data.Location, data.TypeId, data.Latitude, data.Longitude, data.Facilities, data.Cover, data.Area, data.RoomConfig, data.CleanVideo, data.HomestayBusinessId, data.HostId, data.HostAvatar, data.HostNickname, data.RowState, data.PriceBefore, data.PriceAfter, data.Id, oldVersion)
	}, looklookTravelHomestayIdKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return genModel.ErrNoRowsUpdate
	}

	return nil
}

func (m *defaultHomestayModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *Homestay) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "HomestayModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultHomestayModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultHomestayModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultHomestayModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*Homestay, error) {

	builder = builder.Columns(homestayRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Homestay
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Homestay, error) {

	builder = builder.Columns(homestayRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Homestay
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Homestay, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(homestayRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*Homestay
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultHomestayModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Homestay, error) {

	builder = builder.Columns(homestayRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Homestay
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Homestay, error) {

	builder = builder.Columns(homestayRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Homestay
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultHomestayModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
func (m *defaultHomestayModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	looklookTravelHomestayIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHomestayIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, looklookTravelHomestayIdKey)
	return err
}
func (m *defaultHomestayModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheLooklookTravelHomestayIdPrefix, primary)
}
func (m *defaultHomestayModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", homestayRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultHomestayModel) tableName() string {
	return m.table
}
