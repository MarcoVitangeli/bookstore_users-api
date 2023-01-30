package date

import "time"

/*
*
Golang uses a sample date as reference, we could play with the
values, and Go is going to parse them as a placeholder for
the value that they represent (2006 is a year, and so on)
NOTE: "Z" means the standard time zone. If we wanted to use argentina
time zone it would be like this:
2006-01-02T15:04:05-03:00
to say that we are `-3` hours of difference with UTC
*/
const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNow() time.Time {
	return time.Now().UTC() // we usa the universal time zone
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
