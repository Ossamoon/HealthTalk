package model

import (
	"time"

	"gorm.io/gorm"
)

type HealthRecord struct {
	gorm.Model
	UserID				uint		`json:"user_id" gorm:"index;not null;"`
	Date				time.Time	`json:"date" gorm:"index;not null;type:date;"`
	BodyTemperature		float32		`json:"body_temperature"`
	BloodPressureMax	int			`json:"blood_pressure_max"`
	BloodPressureMin	int			`json:"blood_pressure_min"`
	HeartRate			int			`json:"heart_rate"`
	CoughOrSoreThroat	float32		`json:"have_cough"`
	Headache			float32		`json:"have_headache"`
	Stomachache			float32		`json:"have_stomachache"`
	FeelTired			float32		`json:"feel_tired"`
	Memo				string		`json:"memo"`
}

func CreateHealthRecord(healthRecord *HealthRecord) {
    db.Create(healthRecord)
}

func FindHealthRecords(record *HealthRecord) []HealthRecord {
    var healthRecords []HealthRecord
    db.Where(record).Find(&healthRecords)
    return healthRecords
}