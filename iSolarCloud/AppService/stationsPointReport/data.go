package stationsPointReport

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"
)

const Url = "/v1/powerStationService/stationsPointReport"
const Disabled = false
const EndPointName = "AppService.stationsPointReport"

type RequestData struct {
	Type valueTypes.String `json:"type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData   struct {
	List []struct {
		GoStruct.GoStructParent `json:"-" PointIdFromChild:"PsId" PointIdReplace:"true"`

		PsId        valueTypes.Integer `json:"ps_id"`
		PsName      valueTypes.String  `json:"ps_name"`
		PsShortName valueTypes.String  `json:"ps_short_name"`

		Co2            valueTypes.Float `json:"co2"`
		Coal           valueTypes.Float `json:"coal"`
		MaxPr          valueTypes.Float `json:"max_pr"`
		Meter          valueTypes.Float `json:"meter"`
		MinPr          valueTypes.Float `json:"min_pr"`
		P83018         valueTypes.Float `json:"p83018"`
		P83022         valueTypes.Float `json:"p83022"`
		P83023         valueTypes.Float `json:"p83023"`
		P83023Original valueTypes.Float `json:"p83023_original"`
		P83025         valueTypes.Float `json:"p83025"`
		P83033         valueTypes.Float `json:"p83033"`
		P83037         valueTypes.Float `json:"p83037"`
		P83038         valueTypes.Float `json:"p83038"`
		P83202         valueTypes.Float `json:"p83202"`
		ParamCo2       valueTypes.Float `json:"param_co2"`
		ParamCoal      valueTypes.Float `json:"param_coal"`
		ParamMeter     valueTypes.Float `json:"param_meter"`
		ParamNox       valueTypes.Float `json:"param_nox"`
		ParamPowder    valueTypes.Float `json:"param_powder"`
		ParamSo2       valueTypes.Float `json:"param_so2"`
		ParamTree      valueTypes.Float `json:"param_tree"`
		ParamWater     valueTypes.Float `json:"param_water"`
		Pr             valueTypes.Float `json:"pr"`
		PrramPr        valueTypes.Float `json:"prram_pr"`
		PrramPrValue   valueTypes.Float `json:"prram_pr_value"`
		So2            valueTypes.Float `json:"so2"`
		Tree           valueTypes.Float `json:"tree"`
	} `json:"list" PointId:"devices" PointIdReplace:"true" DataTable:"true" DataTableSortOn:"PsId"`
	MinDateId valueTypes.DateTime  `json:"min_date_id"`
	RowCount  valueTypes.Integer `json:"rowCount" PointId:"row_count"`
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
