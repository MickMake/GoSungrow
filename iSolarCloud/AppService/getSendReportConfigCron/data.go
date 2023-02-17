package getSendReportConfigCron

import (
	"encoding/json"
	"fmt"

	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/reportService/getSendReportConfigCron"
const Disabled = false
const EndPointName = "AppService.getSendReportConfigCron"

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
	Schedule Schedule `json:"schedule"`
	// One       valueTypes.String `json:"1"`
	// One0      valueTypes.String `json:"10"`
	// One000    valueTypes.String `json:"1000"`
	// One00000  valueTypes.String `json:"100000"`
	// One00001  valueTypes.String `json:"100001"`
	// One00002  valueTypes.String `json:"100002"`
	// One001    valueTypes.String `json:"1001"`
	// One002    valueTypes.String `json:"1002"`
	// One003    valueTypes.String `json:"1003"`
	// One004    valueTypes.String `json:"1004"`
	// One005    valueTypes.String `json:"1005"`
	// One006    valueTypes.String `json:"1006"`
	// One007    valueTypes.String `json:"1007"`
	// One050    valueTypes.String `json:"1050"`
	// One051    valueTypes.String `json:"1051"`
	// One1      valueTypes.String `json:"11"`
	// One2      valueTypes.String `json:"12"`
	// Two       valueTypes.String `json:"2"`
	// Two000    valueTypes.String `json:"2000"`
	// Two001    valueTypes.String `json:"2001"`
	// Two002    valueTypes.String `json:"2002"`
	// Two003    valueTypes.String `json:"2003"`
	// Two004    valueTypes.String `json:"2004"`
	// Two005    valueTypes.String `json:"2005"`
	// Two006    valueTypes.String `json:"2006"`
	// Two007    valueTypes.String `json:"2007"`
	// Two010    valueTypes.String `json:"2010"`
	// Two011    valueTypes.String `json:"2011"`
	// Two012    valueTypes.String `json:"2012"`
	// Two013    valueTypes.String `json:"2013"`
	// Two014    valueTypes.String `json:"2014"`
	// Two015    valueTypes.String `json:"2015"`
	// Three     valueTypes.String `json:"3"`
	// Three001  valueTypes.String `json:"3001"`
	// Three002  valueTypes.String `json:"3002"`
	// Four      valueTypes.String `json:"4"`
	// Four001   valueTypes.String `json:"4001"`
	// Four002   valueTypes.String `json:"4002"`
	// Four101   valueTypes.String `json:"4101"`
	// Four201   valueTypes.String `json:"4201"`
	// Four301   valueTypes.String `json:"4301"`
	// Four501   valueTypes.String `json:"4501"`
	// Five      valueTypes.String `json:"5"`
	// Five00    valueTypes.String `json:"500"`
	// Five001   valueTypes.String `json:"5001"`
	// Five002   valueTypes.String `json:"5002"`
	// Five003   valueTypes.String `json:"5003"`
	// Six       valueTypes.String `json:"6"`
	// Seven     valueTypes.String `json:"7"`
	// Eight     valueTypes.String `json:"8"`
	// Eight00   valueTypes.String `json:"800"`
	// Nine      valueTypes.String `json:"9"`
	// Nine00    valueTypes.String `json:"900"`
	// Nine01    valueTypes.String `json:"901"`
	// Nine02    valueTypes.String `json:"902"`
	// Nine03    valueTypes.String `json:"903"`
	// Nine04    valueTypes.String `json:"904"`
	// Nine05    valueTypes.String `json:"905"`
	// Nine06    valueTypes.String `json:"906"`
	// Nine07    valueTypes.String `json:"907"`
	// Nine08    valueTypes.String `json:"908"`
	// Nine10    valueTypes.String `json:"910"`
	// Nine11    valueTypes.String `json:"911"`
	// Nine12    valueTypes.String `json:"912"`
	// Nine13    valueTypes.String `json:"913"`
	// Nine14    valueTypes.String `json:"914"`
	// Nine15    valueTypes.String `json:"915"`
	// Nine16    valueTypes.String `json:"916"`
	// Nine17    valueTypes.String `json:"917"`
	// Nine18    valueTypes.String `json:"918"`
	// Nine20    valueTypes.String `json:"920"`
	// Nine21    valueTypes.String `json:"921"`
	// Nine22    valueTypes.String `json:"922"`
	// Nine23    valueTypes.String `json:"923"`
	// Nine24    valueTypes.String `json:"924"`
	// Nine25    valueTypes.String `json:"925"`
	// Nine26    valueTypes.String `json:"926"`
	// Nine31    valueTypes.String `json:"931"`
	// Nine32    valueTypes.String `json:"932"`
	// Nine38    valueTypes.String `json:"938"`
	// Nine39    valueTypes.String `json:"939"`
	// Nine40    valueTypes.String `json:"940"`
	// Nine41    valueTypes.String `json:"941"`
	// Nine50    valueTypes.String `json:"950"`
	// Nine90    valueTypes.String `json:"990"`
	// Nine93    valueTypes.String `json:"993"`
	// Nine94    valueTypes.String `json:"994"`
	// Nine95    valueTypes.String `json:"995"`
	// Nine96    valueTypes.String `json:"996"`
	// Nine97    valueTypes.String `json:"997"`
	// Nine98    valueTypes.String `json:"998"`
	// Nine99    valueTypes.String `json:"999"`
	// Nine999   valueTypes.String `json:"9999"`
	// Nine99999 valueTypes.String `json:"999999"`
}
type Schedule map[string]valueTypes.String

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *ResultData) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}

		type scan Schedule
		var d scan
		err = json.Unmarshal(data, &d)
		if err != nil {
			break
		}

		e.Schedule = Schedule(d)
	}

	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
