package iSolarCloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/output"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

// ****************************************************** //

func (sg *SunGrow) NewSunGrowData() SunGrowData {
	var data SunGrowData
	for range Only.Once {
		data.New(sg)

		data.SetOutputType(sg.OutputType)
		data.SetSaveAsFile(sg.SaveAsFile)
	}
	return data
}

func SplitArg(arg string) []string {
	var ret []string
	for range Only.Once {
		ret = []string{arg}
		for _, s := range []string{",", "/", " "} {
			if strings.Contains(arg, s) {
				ret = strings.Split(arg, s)
				break
			}
		}
	}
	return ret
}

type (
	EndPoints map[string]EndPoint
	EndPoint  struct {
		Func    SunGrowDataFunction
		HasArgs bool
	}
)

type SunGrowData struct {
	Args      []string
	endPoints []string
	Request   SunGrowDataRequest
	Results   SunGrowDataResults

	sunGrow *SunGrow
	// outputType   output.OutputType
	// saveAsFile   bool
	cacheTimeout time.Duration

	Debug bool
	Error error
}

func (sgd *SunGrowData) PrintDebug(format string, args ...interface{}) {
	if sgd.Debug {
		_, _ = fmt.Fprintf(os.Stderr, format, args...)
	}
}

func (sgd *SunGrowData) New(ref *SunGrow) {
	for range Only.Once {
		sgd.sunGrow = ref
		sgd.Results = make(SunGrowDataResults)
		sgd.cacheTimeout = time.Minute * 5
	}
}

func (sgd *SunGrowData) SetCacheTimeout(t time.Duration) {
	sgd.cacheTimeout = t
}

func (sgd *SunGrowData) SetOutput(t string) {
	sgd.sunGrow.OutputType.Set(t)
}

func (sgd *SunGrowData) SetOutputType(t output.OutputType) {
	sgd.sunGrow.OutputType = t
}

func (sgd *SunGrowData) SetSaveAsFile(yes bool) {
	sgd.sunGrow.SaveAsFile = yes
}

func (sgd *SunGrowData) SetEndpoints(endpoints ...string) {
	sgd.endPoints = endpoints
}

func (sgd *SunGrowData) SetArgs(args ...string) {
	sgd.Args = args
}

func (sgd *SunGrowData) SetPsIds(psids ...string) {
	for range Only.Once {
		var pids valueTypes.PsIds
		pids = sgd.sunGrow.SetPsIds(psids...)
		if sgd.Error != nil {
			break
		}

		sgd.Request.SetPSIDs(pids.Strings())
	}
}

func (sgd *SunGrowData) CallEndpoint(endpoint api.EndPoint, request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		if !request.Validate(endpoint) {
			request.Help(endpoint)
			sgd.Error = errors.New("missing argument")
			break
		}

		var req []byte
		req, sgd.Error = json.Marshal(request)
		if sgd.Error != nil {
			fmt.Printf("GetEndPoint - Error: %s\n", sgd.Error)
			break
		}
		// fmt.Printf("Request: %s\n", req)

		if string(req) != "" {
			endpoint = endpoint.SetRequestByJson(output.Json(req))
			sgd.Error = endpoint.GetError()
			if sgd.Error != nil {
				fmt.Println(endpoint.Help())
				break
			}
		}
		sgd.PrintDebug("Request: %s\n", endpoint.GetRequestJson())

		// @TODO - Make this a config option.
		endpoint = endpoint.SetCacheTimeout(sgd.cacheTimeout)

		endpoint = endpoint.Call()
		sgd.Error = endpoint.GetError()
		if sgd.Error != nil {
			if strings.Contains(sgd.Error.Error(), "er_token_login_invalid") {
				sgd.sunGrow.Logout()
				break
			}
			fmt.Println(endpoint.Help())
			sgd.PrintDebug("Error response[%s]: %s\n", sgd.Error, endpoint.GetResponseJson())
			break
		}
		sgd.PrintDebug("Response: %s\n", endpoint.GetResponseJson())

		response.Data = endpoint.GetEndPointData()
		args := request.GetArgs(response.Data.EndPoint)
		hash := request.GetArgsHash(response.Data.EndPoint)
		name := endpoint.GetArea().String() + "." + endpoint.GetName().String()
		var title string
		var file string // + " - " + request.RequestAsFilePrefix(),
		key := request.GetPrimaryArg()
		if key != "" {
			title = key
			file = key
		}

		response.Options = OutputOptions{
			Name:        name,
			OutputType:  sgd.sunGrow.OutputType,
			PrimaryKey:  key,
			Directory:   sgd.sunGrow.Directory,
			FilePrefix:  hash,
			FileSuffix:  file,
			SaveAsFile:  sgd.sunGrow.SaveAsFile,
			TitleSuffix: args,
			GraphRequest: output.GraphRequest{
				Title:       title,
				SubTitle:    args,
				TimeColumn:  nil,
				DataColumn:  nil,
				UnitsColumn: nil,
				NameColumn:  nil,
				DataMin:     nil,
				DataMax:     nil,
				Width:       nil,
				Height:      nil,
				Error:       nil,
			},
		}
		sgd.PrintDebug("OutputOptions: %v\n", response.Options)
	}

	return response
}

func (sgd *SunGrowData) GetData() error {
	for range Only.Once {
		if len(sgd.endPoints) == 0 {
			sgd.Error = errors.New("need an endpoint")
			break
		}

		for _, endpoint := range sgd.endPoints {
			sgd.Error = sgd.GetDataSingle(endpoint)
			if sgd.Error != nil {
				break
			}
		}
	}

	return sgd.Error
}

func (sgd *SunGrowData) GetDataSingle(endpoint string) error {
	for range Only.Once {
		// Lookup endpoint interface from string.
		ep := sgd.sunGrow.GetEndpoint(endpoint)
		if sgd.sunGrow.IsError() {
			sgd.Error = sgd.sunGrow.Error
			break
		}
		sgd.Request.SetRequired(ep.GetRequestArgNames())
		sgd.Request.SetArgs(sgd.Args...)

		// PsId not required.
		if sgd.Request.IsPsIdNotRequired() {
			sgd.Error = sgd.getDataSinglePsIdNotRequired(ep)
			break
		}

		// PsId required.
		if sgd.Request.IsPsIdRequired() {
			sgd.Error = sgd.getDataSinglePsIdRequired(ep)
			break
		}
	}

	return sgd.Error
}

func (sgd *SunGrowData) getDataSinglePsIdNotRequired(ep api.EndPoint) error {
	for range Only.Once {
		var result SunGrowDataResult
		result.EndPointArea = ep.GetArea()
		result.EndPointName = ep.GetName()
		result.EndPoint = ep
		result.Request = sgd.Request

		result.Response = sgd.CallEndpoint(ep, result.Request)
		if sgd.Error != nil {
			break
		}

		sgd.Results[result.EndPointName.String()] = result
		sgd.Error = sgd.Process()
		if sgd.Error != nil {
			break
		}
	}

	return sgd.Error
}

func (sgd *SunGrowData) getDataSinglePsIdRequired(ep api.EndPoint) error {
	for range Only.Once {
		if sgd.Request.aPsId == nil {
			sgd.SetPsIds()
		}

		for _, psId := range sgd.Request.aPsId {
			var result SunGrowDataResult
			result.Request = sgd.Request
			result.Request.SetPsId(psId.String())

			result.EndPointArea = ep.GetArea()
			result.EndPointName = ep.GetName()
			result.EndPoint = ep
			result.Response = sgd.CallEndpoint(ep, result.Request)
			if sgd.Error != nil {
				break
			}
			sgd.Results[result.EndPointName.String()+"/"+psId.String()] = result
		}
		if sgd.Error != nil {
			break
		}

		sgd.Error = sgd.Process()
		if sgd.Error != nil {
			break
		}
	}

	return sgd.Error
}

func (sgd *SunGrowData) Process() error {
	for range Only.Once {
		if len(sgd.Results) == 0 {
			fmt.Println("No results found.")
			break
		}

		for _, result := range sgd.Results {
			result.Response.Options.OutputType = sgd.sunGrow.OutputType
			result.Response.Options.SaveAsFile = sgd.sunGrow.SaveAsFile
			result.Response.Options.Directory = sgd.sunGrow.Directory
			result.Response.Data.ProcessMap()
			if sgd.Error != nil {
				break
			}
		}
	}

	return sgd.Error
}

func (sgd *SunGrowData) Output() error {
	for range Only.Once {
		if len(sgd.Results) == 0 {
			fmt.Println("No results found.")
			break
		}

		for _, result := range sgd.Results {
			result.Response.Options.OutputType = sgd.sunGrow.OutputType
			result.Response.Options.SaveAsFile = sgd.sunGrow.SaveAsFile
			result.Response.Options.Directory = sgd.sunGrow.Directory
			sgd.Error = result.Response.Output()
			if sgd.Error != nil {
				break
			}
		}
	}

	return sgd.Error
}

func (sgd *SunGrowData) OutputDataTables() error {
	for range Only.Once {
		if len(sgd.Results) == 0 {
			fmt.Println("No results found.")
			break
		}

		for _, result := range sgd.Results {
			result.Response.Options.OutputType = sgd.sunGrow.OutputType
			result.Response.Options.SaveAsFile = sgd.sunGrow.SaveAsFile
			result.Response.Options.Directory = sgd.sunGrow.Directory
			sgd.Error = result.Response.OutputDataTables()
			if sgd.Error != nil {
				break
			}
		}
	}

	return sgd.Error
}

type (
	SunGrowDataResults map[string]SunGrowDataResult
	SunGrowDataResult  struct {
		EndPointArea api.AreaName
		EndPointName api.EndPointName
		EndPoint     api.EndPoint
		Request      SunGrowDataRequest
		Response     SunGrowDataResponse

		Error error
	}
)

func (sgd *SunGrowDataResult) Process() error {
	sgd.Response.Data.ProcessMap()
	sgd.Error = sgd.Response.Data.Error
	return sgd.Error
}

func (sgd *SunGrowDataResult) Sort() []string {
	return sgd.Response.Data.Sort()
}

func (sgd *SunGrowDataResult) Print() {
	fmt.Println(sgd.Response.Data.String())
}

type OutputOptions struct {
	Name        string
	TitleSuffix string
	PrimaryKey  string

	OutputType output.OutputType
	Directory  string
	SaveAsFile bool
	FileSuffix string
	FilePrefix string

	GraphRequest output.GraphRequest
}

type (
	SunGrowDataResponses map[string]SunGrowDataResponse
	SunGrowDataFunction  func(request SunGrowDataRequest) SunGrowDataResponse
	SunGrowDataResponse  struct {
		Data    api.DataMap
		Options OutputOptions
		Error   error
	}
)

func (sgd *SunGrowDataResponse) Output() error {
	for range Only.Once {
		// Outputs that don't drop through.
		if sgd.Options.OutputType.IsStruct() || sgd.Options.OutputType.IsList() || sgd.Options.OutputType.IsRaw() || sgd.Options.OutputType.IsJson() {
			table := sgd.Data.CreateResultTable(true)
			table.OutputType = sgd.Options.OutputType
			table.SetSaveFile(sgd.Options.SaveAsFile)
			table.AppendTitle(" - %s", sgd.Options.TitleSuffix)
			table.SetDirectory(sgd.Options.Directory)
			table.PrependFilePrefix(sgd.Options.FilePrefix)
			table.AppendFilePrefix(sgd.Options.FileSuffix)
			sgd.Error = table.Output()
			break
		}

		// Outputs that can drop through to DataTables.
		if sgd.Options.OutputType.IsTable() || sgd.Options.OutputType.IsXLSX() || sgd.Options.OutputType.IsCsv() || sgd.Options.OutputType.IsXML() {
			table := sgd.Data.CreateResultTable(false)
			table.OutputType = sgd.Options.OutputType
			table.SetSaveFile(sgd.Options.SaveAsFile)
			table.AppendTitle(" - %s", sgd.Options.TitleSuffix)
			table.SetDirectory(sgd.Options.Directory)
			table.PrependFilePrefix(sgd.Options.FilePrefix)
			table.AppendFilePrefix(sgd.Options.FileSuffix)
			sgd.Error = table.Output()
			if sgd.Error != nil {
				break
			}
			// break
		}

		sgd.Error = sgd.OutputDataTables()
	}

	return sgd.Error
}

func (sgd *SunGrowDataResponse) OutputDataTables() error {
	for range Only.Once {
		tables := sgd.Data.CreateDataTables()
		if sgd.Data.Error != nil {
			sgd.Error = sgd.Data.Error
			break
		}
		if len(tables) == 0 {
			break
		}

		// @iSolarCloud/api/struct_data.go:420
		for _, data := range tables {
			if sgd.Options.TitleSuffix == "" {
				sgd.Options.TitleSuffix = data.Table.GetTitle()
			}
			data.Table.OutputType = sgd.Options.OutputType
			data.Table.SetSaveFile(sgd.Options.SaveAsFile) // sgd.Options.SaveAsFile

			if sgd.Options.OutputType.IsGraph() {
				if !data.IsValid {
					fmt.Printf("# %s.%s - has no graphable data.\n", data.Area, data.Name)
					continue
				}

				data.Table.SetSaveFile(true)

				if sgd.Options.GraphRequest.TimeColumn == nil {
					for _, col := range data.Table.GetHeaders() {
						val := data.Values.GetCell(0, col)
						if val.IsTypeDateTime() {
							sgd.Options.GraphRequest.TimeColumn = &col
							break
						}
					}
				}
				if sgd.Options.GraphRequest.TimeColumn == nil {
					// No time column - abort.
					break
				}

				if sgd.Options.GraphRequest.UnitsColumn != nil {
					for _, col := range data.Table.GetHeaders() {
						if *sgd.Options.GraphRequest.UnitsColumn != col {
							continue
						}
						val := data.Values.GetCell(0, col)
						unit := val.Unit()
						if unit != "" {
							continue
						}
						sgd.Options.GraphRequest.UnitsColumn = &col
						sgd.Options.GraphRequest.DataUnit = &unit
						break
					}
				}

				if sgd.Options.GraphRequest.NameColumn == nil {
				}

				// if sgd.Options.GraphRequest.Width == nil {
				// }

				// if sgd.Options.GraphRequest.Height == nil {
				// }

				var values []string
				if sgd.Options.GraphRequest.DataColumn == nil {
					fmt.Println("Finding points to graph...")
					fmt.Printf("Table Headers: %s\n", strings.Join(data.Table.GetHeaders(), ", "))
					fmt.Printf("Table rows: %d\n", data.Rows)
					// We don't have any DataColumn defined - find them.
					for _, col := range data.Table.GetHeaders() {
						val := data.Values.GetCell(0, col)
						if !val.IsNumber() {
							continue
						}
						values = append(values, col)
						// if val.ValueKey() == "" {
						// 	values = append(values, col)
						// 	continue
						// }
						// if val.DeviceId() == "" {
						// 	values = append(values, val.ValueKey())
						// 	continue
						// }
						// values = append(values, val.DeviceId() + "." + val.ValueKey())
					}
					fmt.Printf("Found %d points.\n", len(values))
				}

				title := data.Table.GetTitle()
				file := data.Table.GetFilePrefix()

				sort.Strings(values)
				for _, value := range values {
					// @TODO - Lookup pointIds here.
					sgd.Options.GraphRequest.DataColumn = &value
					if sgd.Options.PrimaryKey != "" {
						data.Table.SetTitle("%s - %s - %s", title, sgd.Options.PrimaryKey, value)
						data.Table.SetFilePrefix("%s-%s-%s-%s", sgd.Options.FilePrefix, file, sgd.Options.PrimaryKey, value)
					} else {
						data.Table.SetTitle("%s - %s", title, value)
						data.Table.SetFilePrefix("%s-%s-%s", sgd.Options.FilePrefix, file, value)
					}
					sgd.Options.GraphRequest.Title = data.Table.GetTitle()

					sgd.Error = data.Table.SetGraph(sgd.Options.GraphRequest)
					if sgd.Error != nil {
						break
					}

					data.Table.SetDirectory(sgd.Options.Directory)
					sgd.Error = data.Table.Output()
					if sgd.Error != nil {
						break
					}
				}

				break
			}

			data.Table.AppendTitle(" - %s", sgd.Options.TitleSuffix)
			data.Table.SetDirectory(sgd.Options.Directory)
			data.Table.PrependFilePrefix(sgd.Options.FilePrefix)
			data.Table.AppendFilePrefix(sgd.Options.FileSuffix)
			fmt.Println()
			sgd.Error = data.Table.Output()
			if sgd.Error != nil {
				break
			}
		}
	}

	return sgd.Error
}

func (sgd *SunGrowDataResponse) LookUpPointId() {
}

func (sgd *SunGrowDataResponse) Print() {
	fmt.Println(sgd.Data.String())
}
