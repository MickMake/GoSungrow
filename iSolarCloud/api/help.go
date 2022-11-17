package api

import "fmt"

func HelpDataType() string {
	ret := fmt.Sprintln("date_type:1 = Day")
	ret += fmt.Sprintln("\tdate_id: Format YYYYmmdd")
	ret += fmt.Sprintln("date_type:2 = Month")
	ret += fmt.Sprintln("\tdate_id: Format YYYYmm")
	ret += fmt.Sprintln("date_type:3 = Year")
	ret += fmt.Sprintln("\tdate_id: Format YYYY")
	return ret
}

func HelpDateId() string {
	ret := fmt.Sprintln("date_type:1 = Day")
	ret += fmt.Sprintln("\tstart_time_stamp:YYYYmmddHHMM00")
	ret += fmt.Sprintln("\tend_time_stamp:YYYYmmddHHMM00")
	ret += fmt.Sprintln("date_type:2 = Month")
	ret += fmt.Sprintln("\tstart_time_stamp:YYYYmm")
	ret += fmt.Sprintln("\tend_time_stamp:YYYYmm")
	ret += fmt.Sprintln("date_type:3 = Year")
	ret += fmt.Sprintln("\tstart_time_stamp:YYYY")
	ret += fmt.Sprintln("\tend_time_stamp:YYYY")
	return ret
}

func HelpReportType() string {
	ret := fmt.Sprintf("report_type is from 1-4\n")
	ret += fmt.Sprintf("1: Day Report.\n")
	ret += fmt.Sprintf("2: Month Report.\n")
	ret += fmt.Sprintf("3: Year Report.\n")
	ret += fmt.Sprintf("4: Total Report.\n")
	return ret
}
