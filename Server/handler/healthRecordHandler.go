package handler


import (
    "net/http"
    "strconv"

    "github.com/labstack/echo"
    "gorm.io/gorm"
    "github.com/Ossamoon/HealthTalk/Server/model"
)


type (
	UpdateHealthRecordRequest struct {
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
)


func AddHealthRecord(c echo.Context) error {
    healthRecord := new(model.HealthRecord)
    if err := c.Bind(healthRecord); err != nil {
        return err
    }

    userID := userIDFromToken(c)
    if user := model.FindUser(&model.User{Model: gorm.Model{ID: userID}}); user.ID == 0 {
        return echo.ErrNotFound
    }

    healthRecord.UserID = userID
    model.CreateHealthRecord(healthRecord)

    responce := CommonCreateResponce {
        ID: healthRecord.Model.ID,
        CreatedAt: healthRecord.Model.CreatedAt,
    }

    return c.JSON(http.StatusCreated, responce)
}


func GetHealthRecords(c echo.Context) error {
    userID := userIDFromToken(c)
    if user := model.FindUser(&model.User{Model: gorm.Model{ID: userID}}); user.ID == 0 {
        return echo.ErrNotFound
    }

    healthRecords := model.FindHealthRecords(&model.HealthRecord{UserID: userID})

    return c.JSON(http.StatusOK, healthRecords)
}


func UpdateHealthRecord(c echo.Context) error {
    tempUint64, _ := strconv.ParseUint(c.Param("record_id"), 10, 64)
    recordID := uint(tempUint64)
	record := model.FindHealthRecord(&model.HealthRecord{Model: gorm.Model{ID: recordID}})
    if record.ID == 0 {
        return echo.ErrNotFound
    }

    userID := userIDFromToken(c)
	if user := model.FindUser(&model.User{Model: gorm.Model{ID: userID}}); user.ID == 0 {
        return echo.ErrNotFound
    }

    if userID != record.UserID {
        return &echo.HTTPError{
            Code:    http.StatusBadRequest,
            Message: "invalid token to update this record",
        }
    }

    updating := new(UpdateHealthRecordRequest)
    if err := c.Bind(updating); err != nil {
        return err
    }

    record.BodyTemperature = updating.BodyTemperature
    record.BloodPressureMax = updating.BloodPressureMax
    record.BloodPressureMin = updating.BloodPressureMin
    record.HeartRate = updating.HeartRate
    record.CoughOrSoreThroat = updating.CoughOrSoreThroat
    record.Headache = updating.Headache
    record.Stomachache = updating.Stomachache
    record.Memo = updating.Memo
    model.UpdateRecord(&record)

    responce := CommonUpdateResponce {
        ID: record.Model.ID,
        UpdatedAt: record.Model.UpdatedAt,
    }

    return c.JSON(http.StatusOK, responce)
}