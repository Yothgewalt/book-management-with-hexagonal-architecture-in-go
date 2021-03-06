package port

import (
	"github.com/google/uuid"
	"time"
)

type Member struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `gorm:"not null;unique;type:varchar(24)"`
	Firstname string `gorm:"not null;type:varchar(64)"`
	Lastname  string `gorm:"not null;type:varchar(64)"`
	Password  string `gorm:"not null"`
}

type MemberRepository interface {
	LoginMember(member Member) (*Member, error)
	CreateMember(member Member) (*Member, error)
	GetAllMember() ([]*Member, error)
	GetMemberById(uuid uuid.UUID) (*Member, error)
	GetMemberByNameWithPassword(name string) (*Member, error)
	DropMemberById(uuid uuid.UUID) error
}
