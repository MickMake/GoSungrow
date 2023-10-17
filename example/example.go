package example

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/login"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/queryMutiPointDataList"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/output"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

// Example1 - GoSungrow API example
func Example1(startDate string, endDate string, interval string, points []string) error {
	var err error
	for range Only.Once {

		// -------------------------------------------------------------------------------- //
		// Initial setup.
		url := "https://augateway.isolarcloud.com"
		appKey := iSolarCloud.DefaultApiAppKey
		user := "mick"
		password := "password"
		token := "tokenfile.json"

		sg := iSolarCloud.NewSunGro(url, "./cache")
		if sg.Error != nil {
			err = sg.Error
			break
		}

		err = sg.Init()
		if err != nil {
			break
		}

		auth := login.SunGrowAuth{
			AppKey:       appKey,
			UserAccount:  user,
			UserPassword: password,
			TokenFile:    token,
			Force:        true,
		}
		err = sg.Login(auth)
		if err != nil {
			break
		}

		// -------------------------------------------------------------------------------- //
		// Produce output.
		data := sg.NewSunGrowData()

		startDate = valueTypes.NewDateTime(startDate).Format(valueTypes.DateTimeLayoutSecond)
		endDate = valueTypes.NewDateTime(endDate).Format(valueTypes.DateTimeLayoutSecond)

		fmt.Printf("Points: %s\n", strings.Join(points, ","))
		data.SetArgs(
			"StartTimeStamp:"+startDate,
			"EndTimeStamp:"+endDate,
			"MinuteInterval:"+interval,
			"Points:"+strings.Join(points, ","),
		)
		data.SetEndpoints(queryMutiPointDataList.EndPointName)

		sg.Error = data.GetData()
		if sg.Error != nil {
			// break
		}

		sg.SetOutputType(output.StringTypeTable)
		sg.SaveAsFile = false

		sg.Error = data.OutputDataTables()
		if sg.Error != nil {
			break
		}
	}

	return err
}

// Example2 - GoSungrow API example
func Example2(startDate string, endDate string, interval string, points []string) error {
	var err error
	for range Only.Once {

		// -------------------------------------------------------------------------------- //
		// Setup values.
		p := valueTypes.SetPointIdsString(points...)

		if len(p.PointIds) == 0 {
			break
		}

		psids := p.PsIds()
		if len(psids) == 0 {
			break
		}
		psId := psids[0]

		i := valueTypes.SetIntegerString(interval)
		if !i.Valid {
			i.SetValue(5)
		}

		sd := valueTypes.SetDateTimeString(startDate)
		if sd.IsZero() {
			sd.SetValue(time.Now())
			sd.SetDayStart()
		}
		sd.SetDateType(valueTypes.DateTimeLayoutSecond)

		ed := valueTypes.SetDateTimeString(endDate)
		if ed.IsZero() {
			ed.SetValue(sd.Value())
			ed.SetDayEnd()
		}
		ed.SetDateType(valueTypes.DateTimeLayoutSecond)

		// -------------------------------------------------------------------------------- //
		// Initial setup.
		url := "https://augateway.isolarcloud.com"
		appKey := iSolarCloud.DefaultApiAppKey
		user := "mick"
		password := "password"
		token := "tokenfile.json"

		sg := iSolarCloud.NewSunGro(url, "./cache")
		if sg.Error != nil {
			err = sg.Error
			break
		}

		err = sg.Init()
		if err != nil {
			break
		}

		auth := login.SunGrowAuth{
			AppKey:       appKey,
			UserAccount:  user,
			UserPassword: password,
			TokenFile:    token,
			Force:        true,
		}
		err = sg.Login(auth)
		if err != nil {
			break
		}

		// -------------------------------------------------------------------------------- //
		// Produce output.
		ep := sg.GetByStruct(
			"AppService.queryMutiPointDataList",
			queryMutiPointDataList.RequestData{
				PsId:           psId,
				StartTimeStamp: sd,
				EndTimeStamp:   ed,
				MinuteInterval: i,
				PsKeys:         *p.PsKeys(),
				Points:         p,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		var response iSolarCloud.SunGrowDataResponse
		response.Data = ep.GetEndPointData()

		response.Data.ProcessMap()
		if response.Data.Error != nil {
			sg.Error = response.Data.Error
			break
		}

		table := response.Data.CreateResultTable(false, "")
		table.OutputType = output.TypeTable
		table.SetSaveFile(false)
		sg.Error = table.Output()
		if sg.Error != nil {
			break
		}

		response.Options = iSolarCloud.OutputOptions{
			OutputType:   output.TypeTable,
			SaveAsFile:   false,
			GraphRequest: output.GraphRequest{},
		}

		sg.Error = response.OutputDataTables("")
		if sg.IsError() {
			break
		}
	}

	return err
}

func Example3(startDate valueTypes.DateTime, endDate valueTypes.DateTime, interval valueTypes.Integer, points valueTypes.PointIds) error {
	var err error
	for range Only.Once {

		// -------------------------------------------------------------------------------- //
		// Setup values.
		if len(points.PointIds) == 0 {
			break
		}

		psids := points.PsIds()
		if len(psids) == 0 {
			break
		}
		psId := psids[0]

		if !interval.Valid {
			interval.SetValue(5)
		}

		if startDate.IsZero() {
			startDate.SetValue(time.Now())
			startDate.SetDayStart()
		}
		startDate.SetDateType(valueTypes.DateTimeLayoutSecond)

		if endDate.IsZero() {
			endDate.SetValue(startDate.Value())
			endDate.SetDayEnd()
		}
		endDate.SetDateType(valueTypes.DateTimeLayoutSecond)

		if !interval.Valid {
			interval.SetValue(5)
		}

		// -------------------------------------------------------------------------------- //
		// Initial setup.
		url := "https://augateway.isolarcloud.com"
		appKey := iSolarCloud.DefaultApiAppKey
		user := "mick"
		password := "password"
		token := "tokenfile.json"

		sg := iSolarCloud.NewSunGro(url, "./cache")
		if sg.Error != nil {
			err = sg.Error
			break
		}

		err = sg.Init()
		if err != nil {
			break
		}

		auth := login.SunGrowAuth{
			AppKey:       appKey,
			UserAccount:  user,
			UserPassword: password,
			TokenFile:    token,
			Force:        true,
		}
		err = sg.Login(auth)
		if err != nil {
			break
		}

		// -------------------------------------------------------------------------------- //
		// Produce output.
		ep := sg.GetByStruct(
			"AppService.queryMutiPointDataList",
			queryMutiPointDataList.RequestData{
				PsId:           psId,
				StartTimeStamp: startDate,
				EndTimeStamp:   endDate,
				MinuteInterval: interval,
				PsKeys:         *points.PsKeys(),
				Points:         points,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		var response iSolarCloud.SunGrowDataResponse
		response.Data = ep.GetEndPointData()

		response.Data.ProcessMap()
		if response.Data.Error != nil {
			sg.Error = response.Data.Error
			break
		}

		response.Options = iSolarCloud.OutputOptions{
			OutputType:   output.TypeGraph,
			SaveAsFile:   false,
			GraphRequest: output.GraphRequest{},
		}

		sg.Error = response.OutputDataTables()
		if sg.IsError() {
			break
		}
	}

	return err
}

func main() {
	var err error

	startDate := "20220701"
	endDate := "20220701"
	interval := "5"
	points := []string{"Points:1171348_43_2_2.p58603", "1171348_43_2_2.p58604", "1171348_43_2_2.p58605", "1171348_43_2_2.p58606"}

	err = Example1(startDate, endDate, interval, points)
	if err != nil {
		log.Fatalf("Error: %s", err)
		os.Exit(1)
	}

	err = Example2(startDate, endDate, interval, points)
	if err != nil {
		log.Fatalf("Error: %s", err)
		os.Exit(1)
	}

	err = Example3(valueTypes.SetDateTimeString(startDate), valueTypes.SetDateTimeString(endDate),
		valueTypes.SetIntegerString(interval), valueTypes.SetPointIdsString(points...))
	if err != nil {
		log.Fatalf("Error: %s", err)
		os.Exit(1)
	}

	os.Exit(0)
}
