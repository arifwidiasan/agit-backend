package helper

import "time"

// get time now (UTC +7)/Jakarta
func GetTimeNow() time.Time {
	//return time now with jakarta timezone and format time.Time
	return time.Now().In(time.FixedZone("Jakarta", 7*60*60))
}

// calculate umur / age based on tanggal lahir / date of birth
func CalculateUmur(tanggalLahir *time.Time) uint {
	now := time.Now()
	years := now.Year() - tanggalLahir.Year()

	// Adjust years based on the month and day
	if now.Month() < tanggalLahir.Month() ||
		(now.Month() == tanggalLahir.Month() && now.Day() < tanggalLahir.Day()) {
		years--
	}

	return uint(years)
}
