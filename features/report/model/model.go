package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Report struct {
	//all
	Id                   string `gorm:"primary key"`
	ReportType           string `gorm:"type:enum('tumpukan sampah','pelanggaran sampah')"`
	UsersId              string `gorm:"type:varchar(191);index"`
	Longitude            float64
	Latitude             float64
	Location             string
	Description          string
	Images               []Image `gorm:"foreignKey:ReportId"`
	AddressPoint         string
	Status               string `gorm:"type:enum('perlu ditinjau','diterima','ditolak');default:'perlu ditinjau'"`
	RejectionDescription string

	//rubbish only
	TrashType string `gorm:"type:enum('sampah kering','sampah basah');default:Null"`

	//littering only
	ScaleType      string `gorm:"type:enum('skala besar', 'skala Kecil');default:Null"`
	InsidentDate   string `gorm:"null"`
	InsidentTime   string `gorm:"null"`
	DangerousWaste bool
	CompanyName    string

	//all
	CreatedAt time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Image struct {
	ID        string `gorm:"primaryKey"`
	ReportId  string `gorm:"index;foreignKey:Id"`
	Image     string
	CreatedAt time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (r *Report) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	if r.ReportType == "tumpukan sampah" {
		trimmedUuid := strings.ReplaceAll(newUuid.String(), "-", "")[:15]
		uppercasedUUID := strings.ToUpper(trimmedUuid)
		r.Id = "TS-" + uppercasedUUID
	}
	if r.ReportType == "pelanggaran sampah" {
		trimmedUuid := strings.ReplaceAll(newUuid.String(), "-", "")[:15]
		uppercasedUUID := strings.ToUpper(trimmedUuid)
		r.Id = "PS-" + uppercasedUUID
	}
	return nil
}

func (i *Image) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	i.ID = newUuid.String()
	return nil
}
