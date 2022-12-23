package getDeploymentList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/faultService/getDeploymentList"
const Disabled = false
const EndPointName = "WebIscmAppService.getDeploymentList"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStructParent GoStruct.GoStructParent  `json:"-" DataTable:"true" DataTableSortOn:"DeploymentTime"`

	DeploymentTime valueTypes.DateTime `json:"deploymentTime" PointId:"deployment_time" PointNameDateFormat:"DateTimeLayout"`
	Id             valueTypes.String   `json:"id"`
	Name           valueTypes.String   `json:"name"`
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
