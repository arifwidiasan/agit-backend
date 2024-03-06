package helper

import "time"

// get time now (UTC +7)/Jakarta
func GetTimeNow() time.Time {
	//return time now with jakarta timezone and format time.Time
	return time.Now().In(time.FixedZone("Jakarta", 7*60*60))
}
