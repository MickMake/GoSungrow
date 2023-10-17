package showTjReport

import (
	"fmt"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/reportService/showTjReport"
	Disabled     = false
	EndPointName = "WebAppService.showTjReport"
)

type RequestData struct {
	PsId      valueTypes.PsId     `json:"ps_id" required:"true"`
	MonthDate valueTypes.DateTime `json:"month_date" required:"true"`
}

// ./goraw.sh WebAppService.showTjReport '{"ps_id":1171348,"month_date":"202210"}'

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct{}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
