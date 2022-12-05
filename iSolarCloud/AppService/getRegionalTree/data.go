package getRegionalTree

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/orgService/getRegionalTree"
const Disabled = false
const EndPointName = "AppService.getRegionalTree"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData   struct {
	ResultList []struct {
		PsId       valueTypes.PsId    `json:"ps_id"`
		Id         valueTypes.String  `json:"id"`
		Name       valueTypes.String  `json:"name"`
		OrgId      valueTypes.Integer `json:"org_id"`
		ParentId   valueTypes.Integer `json:"pId" PointId:"pid"`
		Checked    valueTypes.Bool    `json:"checked"`
		IsParent   valueTypes.Bool    `json:"isParent" PointId:"is_parent"`
		IsFirstOrg valueTypes.Bool    `json:"isFirstOrg" PointId:"is_first_org"`
		Open       valueTypes.Bool    `json:"open"`
		ShareType  valueTypes.Integer `json:"share_type"`
	} `json:"resultList" PointId:"results_list" DataTable:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
