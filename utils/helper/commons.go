package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	drop_point "recything/features/drop-point/entity"
	"recything/utils/constanta"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
)

func DecodeJSON(e echo.Context, input interface{}) error {

	decoder := json.NewDecoder(e.Request().Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(input); err != nil {
		return errors.New("input salah, periksa kembali")
	}

	if _, err := govalidator.ValidateStruct(input); err != nil {
		return err
	}

	return nil
}

func BindFormData(c echo.Context, input interface{}) error {

	if err := c.Bind(input); err != nil {
		return err
	}
	// if err := c.Bind(input); err != nil {
	// 	return err
	// }

	decoder := schema.NewDecoder()
	if err := decoder.Decode(input, c.Request().Form); err != nil {
		return err
	}

	if _, err := govalidator.ValidateStruct(input); err != nil {
		return err
	}

	return nil
}

func HttpResponseCondition(err error, Messages ...string) bool {
	for _, Message := range Messages {
		if strings.Contains(err.Error(), Message) {
			return true
		}
	}
	return false
}

func FieldsEqual(a, b interface{}, fields ...string) bool {
	structA := reflect.ValueOf(a)
	structB := reflect.ValueOf(b)

	for _, fieldName := range fields {
		fieldA := structA.FieldByName(fieldName)
		fieldB := structB.FieldByName(fieldName)

		if !reflect.DeepEqual(fieldA.Interface(), fieldB.Interface()) {
			return false
		}
	}

	return true
}

func ConvertUnitToDecimal(unit string) (float64, error) {
	var numericChars []rune
	var decimalSeparatorFound bool

	unitLower := strings.ToLower(unit)

	// Jenis unit yang valid
	validUnits := []string{"kg", "ltr", "pcs"}
	var validUnitFound bool

	for _, validUnit := range validUnits {
		if strings.Contains(unitLower, validUnit) {
			validUnitFound = true
			break
		}
	}

	if !validUnitFound {
		return 0, errors.New("unit harus mengandung kata 'kg', 'ltr', atau 'pcs'")
	}

	for _, char := range unit {
		if unicode.IsDigit(char) {
			numericChars = append(numericChars, char)
		} else if char == '.' || char == ',' {
			if !decimalSeparatorFound {
				numericChars = append(numericChars, '.')
				decimalSeparatorFound = true
			}
		}
	}

	result, err := strconv.ParseFloat(string(numericChars), 64)
	if err != nil {
		return 0, errors.New("gagal mengonversi unit")
	}

	return result, nil
}

func GenerateRandomID(prefix string, length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)

	generatedID := prefix

	for i := 0; i < length; i++ {
		generatedID += fmt.Sprintf("%d", randomGenerator.Intn(10))
	}

	return generatedID
}

func ChangeStatusMission(endDate string) (string, error) {
	var status string
	endDateValid, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return status, err
	}
	currentTime := time.Now().Truncate(24 * time.Hour)
	if endDateValid.Before(currentTime) {
		status = constanta.OVERDUE
	} else {
		status = constanta.ACTIVE
	}
	return status, nil
}

func SortByDay(schedules []drop_point.ScheduleCore) []drop_point.ScheduleCore {
	sort.Slice(schedules, func(i, j int) bool {
		daysOrder := map[string]int{"senin": 1, "selasa": 2, "rabu": 3, "kamis": 4, "jumat": 5, "sabtu": 6, "minggu": 7}
		return daysOrder[schedules[i].Day] < daysOrder[schedules[j].Day]
	})

    return schedules
}

func CalculateBonus(badge string, missionPoint int) float64 {
	var bonusRate float64

	switch badge {
	case constanta.BRONZE:
		bonusRate = 0.1
	case constanta.SILVER:
		bonusRate = 0.12
	case constanta.GOLD:
		bonusRate = 0.15
	case constanta.PLATINUM:
		bonusRate = 0.2
	default:
		return 0
	}

	bonus := float64(missionPoint) * bonusRate
	return bonus + float64(missionPoint)
}

func GetWeeksInMonth(year int, month time.Month) int {
	// Menghitung jumlah hari dalam bulan ini
	daysInMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()

	// Menghitung jumlah minggu dalam bulan ini
	weeksInMonth := (daysInMonth + 6) / 7
	return int(weeksInMonth)
}

