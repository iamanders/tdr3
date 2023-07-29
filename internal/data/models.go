package data

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string         `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Set UUID when inserting a BaseModel model
func (model *BaseModel) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", strings.ToLower(uuid.New().String()))
	return nil
}

type Time struct {
	BaseModel
	ClientID string `json:"-"`
	Client   Client `json:"client"`

	ProjectID string  `json:"-"`
	Project   Project `json:"project"`

	Code      *string    `json:"code"`
	Comment   *string    `json:"comment"`
	StartedAt *time.Time `json:"startedAt"`
	StoppedAt *time.Time `json:"stoppedAt"`
}

type Client struct {
	BaseModel
	Title       string  `json:"title"`
	Description *string `json:"description"`
}

type Project struct {
	BaseModel
	Title       string  `json:"title"`
	Description *string `json:"description"`
}
