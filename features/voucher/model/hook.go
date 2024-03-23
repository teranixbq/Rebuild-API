package model

import (
	"crypto/sha256"
	"encoding/binary"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


func (r *Voucher) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	r.Id = newUuid.String()
	return nil
}

func (r *ExchangeVoucher) BeforeCreate(tx *gorm.DB) (err error) {
	myUUID := uuid.New()

	hash := sha256.Sum256(myUUID[:])
	Iduint := binary.BigEndian.Uint32(hash[:16])

	r.Id = strconv.FormatUint(uint64(Iduint), 10)
	return nil
}