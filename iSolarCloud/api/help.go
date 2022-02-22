package api

import "fmt"

func HelpDataType() string {
	ret := fmt.Sprintln("date_type: Day = 1")
	ret += fmt.Sprintln("\tdate_id: Format YYYYmmdd")
	ret += fmt.Sprintln("date_type: Month = 2")
	ret += fmt.Sprintln("\tdate_id: Format YYYYmm")
	ret += fmt.Sprintln("date_type: Year = 3")
	ret += fmt.Sprintln("\tdate_id: Format YYYY")
	return ret
}
