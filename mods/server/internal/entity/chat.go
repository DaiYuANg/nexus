package entity

import "time"

type Chat struct {
	BaseModel
	ChatType string `gorm:"not null"` // "private" 或 "group"
	Name     string `gorm:"size:255"`
}

func (Chat) TableName() string {
	return "chat"
}

type ChatMember struct {
	BaseModel
	ChatID   uint      `gorm:"not null"`
	UserID   uint      `gorm:"not null"`
	JoinedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Chat     Chat      `gorm:"foreignKey:ChatID"`
	User     User      `gorm:"foreignKey:UserID"`
}

type Message struct {
	BaseModel
	ChatID      uint   `gorm:"not null"`
	SenderID    uint   `gorm:"not null"`
	MessageType string `gorm:"type:enum('text', 'image', 'file', 'audio');not null"`
	Content     string `gorm:"type:text"`
	MediaURL    string `gorm:"size:255"`
	Status      string `gorm:"type:enum('sent', 'delivered', 'read');default:'sent'"`
	Chat        Chat   `gorm:"foreignKey:ChatID"`
	Sender      User   `gorm:"foreignKey:SenderID"`
}
