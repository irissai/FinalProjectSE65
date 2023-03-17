package orm

import (
	"time"

	"gorm.io/gorm"
)

// สร้าง structure เพื่อรองรับ json
type Booking struct {
	gorm.Model
	UserID string
	CarID  string
	Start  time.Time
	End    time.Time
}
