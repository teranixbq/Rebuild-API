package validation

import (
	"errors"

	"recything/utils/constanta"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-sql-driver/mysql"
)

func CheckDataEmpty(data ...any) error {
	for _, value := range data {
		if value == "" {
			return errors.New(constanta.ERROR_EMPTY)
		}
		if value == 0 {
			return errors.New(constanta.ERROR_EMPTY)
		}
	}
	return nil
}

func CheckEqualData(data string, validData []string) (string, error) {
	inputData := strings.ToLower(data)

	isValidData := false
	for _, data := range validData {
		if inputData == strings.ToLower(data) {
			isValidData = true
			break
		}
	}

	if !isValidData {
		return "", errors.New(constanta.ERROR_INVALID_INPUT)
	}

	return inputData, nil
}

func EmailFormat(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if emailRegex.MatchString(email) {
		return nil
	}
	return errors.New(constanta.ERROR_FORMAT_EMAIL)
}

func PhoneNumber(phone string) error {
	if len(phone) < 10 || len(phone) > 16 {
		return errors.New("panjang nomor telepon harus antara 10 dan 16 karakter")
	}

	phoneRegex := `^(0811|0812|0813|0821|0822|0823|0851|0852|0853|0814|0815|0816|0855|0856|0857|0858|0895|0896|0897|0898|0899|0817|0818|0819|0859|0877|0878|0879|0881|0882|0883|0885|0886|0887|0888|0889|0810|0854|0880|0884|0889|0891|0892|0893|0894|0896|0897|0899|62811|62812|62813|62821|62822|62823|62851|62852|62853|62814|62815|62816|62855|62856|62857|62858|62895|62896|62897|62898|62899|62817|62818|62819|62859|62877|62878|62879|62881|62882|62883|62885|62886|62887|62888|62889|62810|62854|62880|62884|62889|62891|62892|62893|62894|62896|62897|62899)\d{8}$`
	regex := regexp.MustCompile(phoneRegex)

	if regex.MatchString(phone) {
		return nil
	}

	return errors.New("format nomor telepon tidak valid")
}

func MinLength(data string, minLength int) error {
	if len(data) < minLength {
		return errors.New("minimal " + strconv.Itoa(minLength) + " karakter, ulangi kembali!")
	}
	return nil
}

func ValidateTime(openTime, closeTime string) error {
	open, err := time.Parse("15:04", openTime)
	if err != nil {
		return errors.New("format waktu buka tidak valid")
	}

	close, err := time.Parse("15:04", closeTime)
	if err != nil {
		return errors.New("format waktu tutup tidak valid")
	}

	if close.Before(open) || close.Equal(open) {
		return errors.New("waktu penutupan harus setelah waktu pembukaan")
	}

	return nil
}

func ValidateDate(startDate, endDate string) error {
	layout := "2006-01-02"
	currentTime := time.Now().Truncate(24 * time.Hour)

	start, err := time.Parse(layout, startDate)
	if err != nil {
		return errors.New("tanggal harus dalam format 'yyyy-mm-dd'")
	}

	if start.Before(currentTime) {
		return errors.New("tanggal mulai harus hari ini atau setelahnya")
	}

	end, err := time.Parse(layout, endDate)
	if err != nil {
		return errors.New("tanggal harus dalam format 'yyyy-mm-dd'")
	}

	if end.Before(start) {
		return errors.New("tanggal selesai harus setelah tanggal mulai")
	}

	if end.Equal(start) {
		return errors.New("tanggal mulai harus berbeda dari tanggal selesai")
	}

	return nil
}

func ValidateDateForUpdate(startDate, endDate string) error {
	layout := "2006-01-02"

	start, err := time.Parse(layout, startDate)
	if err != nil {
		return errors.New("tanggal harus dalam format 'yyyy-mm-dd'")
	}

	end, err := time.Parse(layout, endDate)
	if err != nil {
		return errors.New("tanggal harus dalam format 'yyyy-mm-dd'")
	}

	if end.Before(start) {
		return errors.New("tanggal selesai harus setelah tanggal mulai")
	}

	if end.Equal(start) {
		return errors.New("tanggal mulai harus berbeda dari tanggal selesai")
	}

	return nil
}

// for repository
func IsDuplicateError(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		return mysqlErr.Number == 1062
	}
	return false
}

func ValidateParamsPagination(page, limit string) (int, int, error) {
	var limitInt int
	var pageInt int
	var err error
	if limit == "" {
		limitInt = 10
	}
	if limit != "" {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			return 0, 0, errors.New("limit harus berupa angka")
		}

		if limitInt > 10 {
			return 0, 0, errors.New("limit tidak boleh lebih dari 10")
		}
	}

	if page == "" {
		pageInt = 1
	}
	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			return 0, 0, errors.New("page harus berupa angka")
		}
	}

	pageInt, limitInt = ValidateCountLimitAndPage(pageInt, limitInt)
	return pageInt, limitInt, nil

}

func ValidateCountLimitAndPage(page, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}

	maxLimit := 10
	if limit <= 0 || limit > maxLimit {
		limit = maxLimit
	}

	return page, limit
}

func ValidateTypePaginationParameter(limit, page string) (int, int, error) {
	var limitInt, pageInt int
	var limitErr, pageErr error
	var limitError, pageError bool

	// Fungsi bantu untuk memeriksa apakah string berupa angka
	isNumeric := func(s string) bool {
		_, err := strconv.Atoi(s)
		return err == nil
	}

	// Validasi untuk limit
	if limit != "" {
		limitInt, limitErr = strconv.Atoi(limit)
		if limitErr != nil || !isNumeric(limit) {
			limitError = true
		} else if limitInt > 10 {
			limitError = true
			return 0, 0, errors.New("limit tidak boleh lebih dari 10")
		}
	}

	// Validasi untuk page
	if page != "" {
		pageInt, pageErr = strconv.Atoi(page)
		if pageErr != nil || !isNumeric(page) {
			pageError = true
		}
	}

	// Menambahkan validasi kedua parameter bersamaan
	if limitError && pageError {
		return 0, 0, errors.New("limit dan page harus berupa angka")
	}

	// Menambahkan validasi jika hanya limit yang error
	if limitError {
		return 0, 0, errors.New("limit harus berupa angka")
	}

	// Menambahkan validasi jika hanya page yang error
	if pageError {
		return 0, 0, errors.New("page harus berupa angka")
	}

	return pageInt, limitInt, nil
}

func CheckLatLong(latitude, longitude float64) error {
	dataLatitude := govalidator.ToString(latitude)
	dataLongitude := govalidator.ToString(longitude)

	errLatLong := govalidator.IsLatitude(dataLatitude)
	if !errLatLong {
		return errors.New("bukan latitude")
	}

	errLongitude := govalidator.IsLongitude(dataLongitude)
	if !errLongitude {
		return errors.New("bukan longitude")
	}

	return nil
}

func ValidateMissionStatus(filter string) string {
	if filter == "aktif" {
		return constanta.ACTIVE
	}
	if filter == "melewati tenggat" {
		return constanta.OVERDUE
	}
	return filter
}
