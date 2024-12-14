package db

import (
	"github.com/sony/sonyflake"
	"gorm.io/gorm"
)

func fillSnowflake(tx *gorm.DB, s *sonyflake.Sonyflake) error {
	// 获取当前操作的模型实例
	if tx.Statement.Schema == nil {
		return nil
	}
	// 只处理包含 `ID` 字段的结构体
	if _, ok := tx.Statement.Schema.FieldsByName["ID"]; ok {
		// 检查是否为零值，如果是则填充 ID
		idValue := tx.Statement.ReflectValue.FieldByName("ID")
		id, err := s.NextID()
		if err != nil {
			return err
		}
		// 生成一个 UUID 并填充到 ID 字段
		idValue.SetUint(id)
	}
	return nil
}
