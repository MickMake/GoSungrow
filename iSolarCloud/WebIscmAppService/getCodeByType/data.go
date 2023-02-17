package getCodeByType

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"
)

const Url = "/v1/powerStationService/getCodeByType"
const Disabled = false
const EndPointName = "WebIscmAppService.getCodeByType"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	Capital []struct {
		CodeValue  valueTypes.String `json:"code_value"`
		CodeName   valueTypes.String `json:"code_name"`
		CodeValue2 interface{}       `json:"code_value2"`
	} `json:"capital" DataTable:"true" DataTableSort:"CodeValue"`
	Scheduling []struct {
		CodeValue  valueTypes.String `json:"code_value"`
		CodeName   valueTypes.String `json:"code_name"`
		CodeValue2 interface{}       `json:"code_value2"`
	} `json:"scheduling" DataTable:"true" DataTableSort:"CodeValue"`
	Schema []struct {
		CodeValue  valueTypes.String `json:"code_value"`
		CodeName   valueTypes.String `json:"code_name"`
		CodeValue2 interface{}       `json:"code_value2"`
	} `json:"schema" DataTable:"true" DataTableSort:"CodeValue"`
	StationType []struct {
		CodeValue  valueTypes.String `json:"code_value"`
		CodeName   valueTypes.String `json:"code_name"`
		CodeValue2 interface{}       `json:"code_value2"`
	} `json:"stationtype" DataTable:"true" DataTableSort:"CodeValue"`
	Zone []struct {
		CodeValue  valueTypes.String `json:"code_value"`
		CodeName   valueTypes.String `json:"code_name"`
		CodeValue2 interface{}       `json:"code_value2"`
	} `json:"zone" DataTable:"true" DataTableSort:"CodeValue"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
