package util

import "time"

const (
	Standard = "2006-01-02"
	HS = "20060102"
)


func StandardFmt(currentTime time.Time) (standardTime string) {
	standardTime = currentTime.Format(Standard)
	return
}