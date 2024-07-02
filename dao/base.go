package dao

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"

	"blogs/lib/infra"
)

type Base struct {
	ID        int64     `json:"id" gorm:"column:id"`
	DeletedID int64     `json:"deleted_id" gorm:"column:deleted_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

var db *gorm.DB

const (
	CallbackQueryNotDeleted = "not_deleted"

	CtxKeyTransID = "ctx_trans_id"
)

func defaultDB(ctx context.Context) *gorm.DB {
	dbItf := ctx.Value(CtxKeyTransID)
	existDB, ok := dbItf.(*gorm.DB)
	if ok && existDB != nil {
		return existDB
	}

	return db.WithContext(ctx)
}

func Trans(parentCtx context.Context, fn func(ctx context.Context) error) error {
	var transDB *gorm.DB
	dbItf := parentCtx.Value(CtxKeyTransID)
	existDB, ok := dbItf.(*gorm.DB)
	if ok && existDB != nil {
		transDB = existDB
	} else {
		transDB = db.WithContext(parentCtx)
		parentCtx = context.WithValue(parentCtx, CtxKeyTransID, transDB)
	}

	return transDB.Transaction(func(tx *gorm.DB) error {
		return fn(parentCtx)
	})
}

func InitDao() error {
	db = infra.MysqlClient
	err := db.Callback().Query().Register(CallbackQueryNotDeleted, func(d *gorm.DB) {
		d.Where("deleted_id = 0")
	})
	return err
}

// Strings 是 []string 的自定义类型
type Strings []string

// Scan 实现 Scanner 接口，用于从数据库读取数据
func (c *Strings) Scan(value interface{}) error {
	if bytes, ok := value.([]byte); ok {
		var ss []string
		err := json.Unmarshal(bytes, &ss)
		if err != nil {
			return err
		}
		*c = ss
		return nil
	}

	// 如果不是字节切片，则可能是 NULL 或其他类型，这里返回错误
	return fmt.Errorf("failed to scan CustomStringSlice")
}

// Value 实现 Valuer 接口，用于将数据写入数据库
func (c Strings) Value() (driver.Value, error) {
	// 将 []string 转换为 JSON 格式的字节切片
	bytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
	// 注意：这里返回 string 而不是 []byte，因为某些数据库驱动可能需要 string
	// 如果你的数据库驱动需要 []byte，则直接返回 bytes 即可
}
