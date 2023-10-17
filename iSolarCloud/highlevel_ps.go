package iSolarCloud

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/findPsType"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getChnnlListByPsId"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getDeviceList"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getHouseholdStoragePsReport"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getIncomeSettingInfos"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPListinfoFromMysql"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPowerChargeSettingInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPowerStationBasicInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPowerStationData"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPowerStationForHousehold"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPowerStationInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPowerStatistics"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPsDetail"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPsDetailForSinglePage"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPsDetailWithPsType"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPsHealthState"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPsInstallerByPsId"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPsInstallerOrgInfoByPsId"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPsList"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getPsWeatherList"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getRemoteUpgradeTaskList"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getReportData"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/psForcastInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/psHourPointsValue"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/queryAllPsIdAndName"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/queryDeviceListForApp"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/queryPowerStationInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/queryPsIdList"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/queryPsProfit"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/queryPsStructureList"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/querySysAdvancedParam"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/reportList"
	"github.com/anicoll/gosungrow/iSolarCloud/Common"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/getPsDeviceListValue"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/getPsKpiForHoursByPsId"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDevicePointAttrs"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsIdState"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getReportPsTree"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/showPSView"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getMaxDeviceIdByPsId"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPsTreeMenu"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/queryDeviceListForBackSys"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	datatable "go.pennock.tech/tabular/auto"
)

// PsList - Return all ps_ids.
func (sg *SunGrow) PsList(psIds ...string) (string, error) {
	var ret string

	for range Only.Once {
		var devices getDeviceList.Devices
		devices, sg.Error = sg.GetDeviceList()
		if sg.Error != nil {
			break
		}

		ret += fmt.Sprintf("# Devices on ps_id %s:\n", strings.Join(psIds, ", "))
		table := datatable.New("utf8-heavy")
		table.AddHeaders("Ps Key", "Ps Id", "Device Type", "Device Code", "Channel Id", "Serial #", "Factory Name", "Device Model")
		for _, device := range devices {
			table.AddRowItems(device.PsKey, device.PsId, device.DeviceType, device.DeviceCode, device.ChannelId, device.Sn, device.FactoryName, device.DeviceModel)
		}

		ret, sg.Error = table.Render()
		if sg.Error != nil {
			break
		}
	}

	return ret, sg.Error
}

// PsPoints - Return all points associated with psIds and device_type filter.
func (sg *SunGrow) PsPoints(psIds []string, deviceType string) (string, error) {
	var ret string

	for range Only.Once {
		var points getDevicePointAttrs.Points
		points, sg.Error = sg.DevicePointAttrs(deviceType, psIds...)
		if sg.Error != nil {
			break
		}

		// Sort table based on PsId + DeviceType + Id
		pn := map[string]int{}
		for index, point := range points {
			pn[point.PsId.String()+"."+point.DeviceType.String()+"."+point.Id.String()] = index
		}
		var names []string
		for point := range pn {
			names = append(names, point)
		}
		sort.Strings(names)

		table := datatable.New("utf8-heavy")
		table.AddHeaders("Id", "Name", "Unit", "Unit Type", "Ps Id", "Device Type", "Device Name")
		for _, name := range names {
			index := pn[name]
			point := points[index]
			table.AddRowItems(point.Id, point.Name, point.Unit, point.UnitType, point.PsId, point.DeviceType, point.DeviceName)
		}

		// PsKey
		// PsId
		// DeviceType
		// DeviceCode
		// ChannelId

		var r string
		r, sg.Error = table.Render()
		if sg.Error != nil {
			break
		}
		ret += fmt.Sprintln("# Available points:")
		ret += r

		// var pids valueTypes.PsIds
		// pids = sg.SetPsIds(psIds...)
		// if sg.Error != nil {
		// 	break
		// }
		//
		// for _, pid := range pids {
		// 	var points []getDevicePointAttrs.Points
		// 	points, sg.Error = sg.GetDevicePointAttrs(pid)
		// 	if sg.Error != nil {
		// 		break
		// 	}
		//
		// 	if len(points) == 0 {
		// 		continue
		// 	}
		//
		// 	ret += fmt.Sprintf("# Available points for ps_id %s:\n", pid.String())
		// 	table := datatable.New("utf8-heavy")
		// 	table.AddHeaders("Id", "Name", "Unit", "UnitType", "PsId", "DeviceType", "DeviceName")
		// 	for _, point := range points {
		// 		if (deviceType != "") && point.DeviceType.MatchString(deviceType) {
		// 			continue
		// 		}
		// 		table.AddRowItems(point.Id, point.Name, point.Unit, point.UnitType, point.PsId, point.DeviceType, point.DeviceName)
		// 	}
		//
		// 	var r string
		// 	r, sg.Error = table.Render()
		// 	if sg.Error != nil {
		// 		break
		// 	}
		// 	ret += r
		//
		// 	// @TODO - Include AppService.getPowerDevicePointNames
		// 	// points2 := cmds.Api.SunGrow.GetDevicePointNames(pid)
		// 	// if c.Error != nil {
		// 	// 	break
		// 	// }
		// 	// if len(points) == 0 {
		// 	// 	continue
		// 	// }
		//
		// }
	}

	return ret, sg.Error
}

func IsInArray(find string, array ...string) bool {
	var yes bool
	for range Only.Once {
		if len(array) == 0 {
			yes = true
			break
		}

		if find == "" {
			break
		}

		for _, s := range array {
			if s == find {
				yes = true
				break
			}
		}
	}
	return yes
}

// PsPointsData - Return all points associated with psIds and device_type filter.
func (sg *SunGrow) PsPointsData(psIds []string, deviceType string, startDate string, endDate string, interval string) error {
	for range Only.Once {
		var tp []string
		for _, i := range psIds {
			if i != "" {
				tp = append(tp, i)
			}
		}
		psIds = tp

		var pskeys valueTypes.PsKeys
		pskeys, sg.Error = sg.GetPsKeys()
		if sg.Error != nil {
			break
		}
		if len(psIds) == 0 {
			psIds = pskeys.PsIds()
		}
		_, _ = fmt.Fprintf(os.Stderr, "Using ps_ids: %s\n", psIds)
		_, _ = fmt.Fprintf(os.Stderr, "Found ps_keys: %s\n", pskeys)

		for _, psId := range psIds {
			var points getDevicePointAttrs.Points
			points, sg.Error = sg.DevicePointAttrs(deviceType, psId)
			if sg.Error != nil {
				break
			}

			var ps []string
			for _, pid := range points {
				match := pskeys.MatchPsIdDeviceType(pid.PsId.String(), pid.DeviceType.String())
				if match.Valid {
					ps = append(ps, fmt.Sprintf("%s.%s", match, pid.Id))
				}
				_, _ = fmt.Fprintf(os.Stderr, "Found points: %s\n", strings.Join(ps, " "))

				sg.Error = sg.PointData(startDate, endDate, interval, ps...)
				if sg.Error != nil {
					break
				}
			}
		}
	}

	return sg.Error
}

// PsPointsDataSave - Return all points associated with psIds and device_type filter and save to files.
func (sg *SunGrow) PsPointsDataSave(psIds []string, deviceType string, startDate string, endDate string, interval string) error {
	for range Only.Once {
		var pskeys valueTypes.PsKeys
		pskeys, sg.Error = sg.GetPsKeys()
		if sg.Error != nil {
			break
		}
		_, _ = fmt.Fprintf(os.Stderr, "Found ps_keys: %s\n", pskeys)

		var points getDevicePointAttrs.Points
		points, sg.Error = sg.DevicePointAttrs(deviceType, psIds...)
		if sg.Error != nil {
			break
		}

		var ps []string
		for _, pid := range points {
			match := pskeys.MatchPsIdDeviceType(pid.PsId.String(), pid.DeviceType.String())
			if match.Valid {
				ps = append(ps, fmt.Sprintf("%s.%s", match, pid.Id))
			}
		}
		// _, _ = fmt.Fprintf(os.Stderr, "Found points: %s\n", strings.Join(ps, " "))

		sg.Error = sg.PointDataSave(startDate, endDate, interval, ps...)
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetPsKeys() (valueTypes.PsKeys, error) {
	var ret valueTypes.PsKeys

	for range Only.Once {
		//        PsKey                   valueTypes.PsKey   `json:"ps_key" PointId:"ps_key" PointUpdateFreq:"UpdateFreqBoot"`
		//        PsId                    valueTypes.PsId    `json:"ps_id" PointId:"ps_id" PointUpdateFreq:"UpdateFreqBoot"`
		//        DeviceType              valueTypes.Integer `json:"device_type" PointId:"device_type" PointUpdateFreq:"UpdateFreqBoot"`
		//        DeviceCode              valueTypes.Integer `json:"device_code" PointId:"device_code" PointUpdateFreq:"UpdateFreqBoot"`
		//        ChannelId               valueTypes.Integer `json:"chnnl_id" PointId:"channel_id" PointUpdateFreq:"UpdateFreqBoot"`

		var pskeys []string
		// Different methods to obtain all pskeys.
		// queryDeviceInfo - DEAD
		// AppService.queryDeviceInfoForApp DeviceSn:B2281302388 - Needs UUID.
		// getDeviceList - Doesn't return the parent device.

		// -------------------------------------------------------------------------------- //
		// Method 1: WebIscmAppService.queryDeviceListForBackSys - Missing values.
		// pids := sg.GetDevices()
		// if sg.Error != nil {
		// 	break
		// }
		// fmt.Printf("%v\n", pids)
		// for _, pid := range pids {
		// 	fmt.Printf("%s_%s_%s_%s\n", "UUID", pid.DeviceType, pid.DeviceCode, pid.ChannelId)
		// }

		// -------------------------------------------------------------------------------- //
		// Method 2: PsTreeMenu
		var trees PsTrees
		trees, sg.Error = sg.PsTreeMenu()
		if sg.Error != nil {
			break
		}
		for pid := range trees {
			for _, tree := range trees[pid].Devices {
				// pskey := fmt.Sprintf("%s_%s_%s_%s", pid.PsId, pid.DeviceType, pid.PsKey.DeviceCode, pid.PsKey.ChannelId)
				pskeys = append(pskeys, tree.PsKey.String())
			}
		}
		if len(pskeys) > 0 {
			ret.Set(pskeys...)
			break
		}

		// -------------------------------------------------------------------------------- //
		// Method 3: AppService.getDeviceList - Around 7kB
		pids2, _ := sg.GetDeviceList()
		if sg.Error != nil {
			break
		}
		for _, pid := range pids2 {
			pskey := fmt.Sprintf("%s_%s_%s_%s", pid.PsId, pid.DeviceType, pid.DeviceCode, pid.ChannelId)
			pskeys = append(pskeys, pskey)
		}
		if len(pskeys) > 0 {
			ret.Set(pskeys...)
			break
		}

		// -------------------------------------------------------------------------------- //
		// Method 4: AppService.getPsList - Around 8.5kB - Doesn't return any pid.DeviceType, pid.DeviceCode, pid.ChannelId
		// pids4, _ := sg.GetPsList()
		// if sg.Error != nil {
		// 	break
		// }
		// // fmt.Printf("%v\n", pids4)
		// for _, pid := range pids4 {
		// 	pskey := fmt.Sprintf("%s_%s_%s_%s\n", pid.PsId, pid.DeviceType, pid.DeviceCode, pid.ChannelId)
		// 	pskeys = append(pskeys, pskey)
		// 	fmt.Printf("(%s) %s\n", pid.PsKey, pskey)
		// }
		// if len(pskeys) > 0 {
		// 	ret.Set(pskeys...)
		// 	break
		// }

		// -------------------------------------------------------------------------------- //
		// Method 5: AppService.queryDeviceList - Around 118kB
		pids3, _ := sg.QueryDeviceList()
		if sg.Error != nil {
			break
		}
		for _, pid := range pids3 {
			pskey := fmt.Sprintf("%s_%s_%s_%s", pid.PsId, pid.DeviceType, pid.DeviceCode, pid.ChannelId)
			pskeys = append(pskeys, pskey)
		}
		if len(pskeys) > 0 {
			ret.Set(pskeys...)
			break
		}
	}

	return ret, sg.Error
}

func (sg *SunGrow) GetDevices() []queryDeviceListForBackSys.Device {
	var ret []queryDeviceListForBackSys.Device
	for range Only.Once {
		var pids valueTypes.PsIds
		pids, sg.Error = sg.GetPsIds()
		if sg.Error != nil {
			break
		}

		for _, pid := range pids {
			var devices []queryDeviceListForBackSys.Device
			devices, sg.Error = sg.QueryDeviceListForBackSys(pid.String())
			if sg.Error != nil {
				break
			}

			ret = append(ret, devices...)
		}
	}
	return ret
}

func (sg *SunGrow) SetPsIds(args ...string) valueTypes.PsIds {
	var pids valueTypes.PsIds
	for range Only.Once {
		if len(args) > 0 {
			pids = valueTypes.SetPsIdStrings(args)
			if len(pids) > 0 {
				break
			}
		}

		pids, sg.Error = sg.GetPsIds()
		if sg.Error != nil {
			break
		}
	}
	return pids
}

func (sg *SunGrow) GetPsIds() (valueTypes.PsIds, error) {
	var ret valueTypes.PsIds

	for range Only.Once {
		ep := sg.GetByStruct(getPsList.EndPointName, nil, DefaultCacheTimeout)
		if sg.IsError() {
			break
		}

		data := getPsList.AssertResultData(ep)
		ret = data.GetPsIds()
	}

	return ret, sg.Error
}

// func (sg *SunGrow) GetPsKeys() ([]string, error) {
// 	var ret []string
//
// 	for range Only.Once {
// 		var psIds valueTypes.PsIds
// 		psIds, sg.Error = sg.GetPsIds()
// 		if sg.Error != nil {
// 			break
// 		}
//
// 		for _, psId := range psIds {
// 			ep := sg.GetByStruct(getPsDetailWithPsType.EndPointName,
// 				// getPsDetailWithPsType.RequestData{PsId: strconv.FormatInt(psId, 10)},
// 				getPsDetailWithPsType.RequestData{PsId: psId},
// 				DefaultCacheTimeout)
// 			if sg.IsError() {
// 				break
// 			}
//
// 			data := getPsDetailWithPsType.Assert(ep)
// 			ret = append(ret, data.GetPsKeys()...)
// 		}
// 	}
//
// 	return ret, sg.Error
// }

// func (sg *SunGrow) GetPsId() (valueTypes.PsId, error) {
// 	var ret valueTypes.PsId
//
// 	for range Only.Once {
//
// 		ep := sg.GetByStruct(getPsList.EndPointName, nil, DefaultCacheTimeout)
// 		if sg.IsError() {
// 			break
// 		}
//
// 		_getPsList := getPsList.AssertResultData(ep)
// 		psIds := _getPsList.GetPsIds()
// 		if len(psIds) == 0 {
// 			break
// 		}
//
// 		ret = psIds[0]
// 	}
//
// 	return ret, sg.Error
// }

// PsTreeMenu - WebIscmAppService.getPsTreeMenu
func (sg *SunGrow) PsTreeMenu(psIds ...string) (PsTrees, error) {
	ret := make(PsTrees)

	for range Only.Once {
		pids := sg.SetPsIds(psIds...)
		if sg.Error != nil {
			break
		}

		for _, psId := range pids {
			ep := sg.GetByStruct(getPsTreeMenu.EndPointName,
				getPsTreeMenu.RequestData{PsId: psId},
				time.Hour,
			)
			if sg.IsError() {
				break
			}

			data := getPsTreeMenu.Assert(ep)

			var p PsTree
			p.Scan(data.Response.ResultData.List, false)
			ret[psId.String()] = p
		}
	}

	return ret, sg.Error
}

const Root = "0"

type Ps struct {
	Children []getPsTreeMenu.Ps
	Depth    int
	Print    bool
}
type PsTree struct {
	Devices  []getPsTreeMenu.Ps
	Map      map[string]*Ps
	Print    bool
	Stringer string
}

func (p PsTree) String() string {
	// p.loop(Root, 0, 0)
	return p.Stringer
}

func (p *PsTree) Root() *getPsTreeMenu.Ps {
	var ret *getPsTreeMenu.Ps
	for range Only.Once {
		if _, ok := p.Map[Root]; !ok {
			break
		}

		if len(p.Map[Root].Children) == 0 {
			break
		}

		ret = &(p.Map[Root].Children[0])
	}
	return ret
}

func (p *PsTree) Scan(devices []getPsTreeMenu.Ps, print bool) {
	p.Devices = devices
	p.Map = make(map[string]*Ps)
	for _, ps := range p.Devices {
		name := ps.UpUUID.String()
		// if name == "0" {
		// 	// name = ps.PsId.String()
		// }

		// fmt.Printf("[%s]\tParent:%s\tSelf: - %s\t%s\t%s\t%s\n", name, ps.UpUUID, ps.UUID, ps.PsId, ps.PsKey, ps.DeviceName)
		if _, ok := p.Map[name]; !ok {
			p.Map[name] = &Ps{
				Children: []getPsTreeMenu.Ps{ps},
			}
		} else {
			p.Map[name].Children = append(p.Map[name].Children, ps)
		}

		p.Print = print
		p.loop(Root, 0, 0)
	}
}

func (p *PsTree) loop(current string, count int, depth int) {
	if _, ok := p.Map[current]; !ok {
		return
	}
	if count >= len(p.Map) {
		return
	}
	count++
	p.Map[current].Depth = depth

	for _, child := range p.Map[current].Children {
		p.Stringer += fmt.Sprintf("%s\tPsId:%s\tPsName:%s\tPsKey:%s\tDeviceName:%s\tUuid:%s\n",
			"+"+strings.Repeat("--", depth),
			child.PsId, child.PsName,
			child.PsKey, child.DeviceName, child.UUID)

		name := child.UUID.String()
		depth++
		p.loop(name, count, depth)
		depth--
	}
}

type PsTrees map[string]PsTree

func (p PsTrees) String() string {
	var ret string
	for i := range p {
		ret += p[i].String()
	}
	return ret
}

// QueryDeviceListForBackSys - WebIscmAppService.queryDeviceListForBackSys
func (sg *SunGrow) QueryDeviceListForBackSys(psId string) ([]queryDeviceListForBackSys.Device, error) {
	var ret []queryDeviceListForBackSys.Device

	for range Only.Once {
		pid := valueTypes.SetPsIdString(psId)
		if !pid.Valid {
			sg.Error = pid.Error
			break
		}

		ep := sg.GetByStruct(queryDeviceListForBackSys.EndPointName,
			queryDeviceListForBackSys.RequestData{PsId: pid},
			DefaultCacheTimeout)
		if sg.IsError() {
			break
		}

		ret = queryDeviceListForBackSys.AssertResultData(ep)
	}

	return ret, sg.Error
}

// GetPsList - AppService.getPsList
func (sg *SunGrow) GetPsList() ([]Common.Device, error) {
	var ret []Common.Device

	for range Only.Once {
		ep := sg.GetByStruct(getPsList.EndPointName, nil, DefaultCacheTimeout)
		if sg.IsError() {
			break
		}

		data := getPsList.AssertResultData(ep)
		ret = data.PageList
	}

	return ret, sg.Error
}

// QueryPowerStationInfo - AppService.queryPowerStationInfo
func (sg *SunGrow) QueryPowerStationInfo(psId valueTypes.PsId, sn valueTypes.String) (queryPowerStationInfo.ResultData, error) {
	var ret queryPowerStationInfo.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(queryPowerStationInfo.EndPointName,
			queryPowerStationInfo.RequestData{
				PsId: psId,
				Sn:   sn,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := queryPowerStationInfo.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// QueryPsIdList - AppService.queryPsIdList
func (sg *SunGrow) QueryPsIdList() ([]valueTypes.String, error) {
	var ret []valueTypes.String
	for range Only.Once {
		ep := sg.GetByStruct(queryPsIdList.EndPointName,
			queryPsIdList.RequestData{},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := queryPsIdList.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// GetPsInstallerOrgInfoByPsId - AppService.getPsInstallerOrgInfoByPsId
func (sg *SunGrow) GetPsInstallerOrgInfoByPsId(psId valueTypes.PsId) (getPsInstallerOrgInfoByPsId.ResultData, error) {
	var ret getPsInstallerOrgInfoByPsId.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsInstallerOrgInfoByPsId.EndPointName,
			getPsInstallerOrgInfoByPsId.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPsInstallerOrgInfoByPsId.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// GetPListinfoFromMysql - AppService.getPListinfoFromMysql
func (sg *SunGrow) GetPListinfoFromMysql(psIds valueTypes.PsIds) (getPListinfoFromMysql.ResultData, error) {
	var ret getPListinfoFromMysql.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPListinfoFromMysql.EndPointName,
			getPListinfoFromMysql.RequestData{
				PsIds: psIds,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPListinfoFromMysql.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// QueryAllPsIdAndName - AppService.queryAllPsIdAndName
func (sg *SunGrow) QueryAllPsIdAndName() (queryAllPsIdAndName.ResultData, error) {
	var ret queryAllPsIdAndName.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(queryAllPsIdAndName.EndPointName,
			queryAllPsIdAndName.RequestData{},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := queryAllPsIdAndName.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// GetRemoteUpgradeTaskList - AppService.getRemoteUpgradeTaskList
func (sg *SunGrow) getRemoteUpgradeTaskList(psId valueTypes.PsIds) (getRemoteUpgradeTaskList.ResultData, error) {
	var ret getRemoteUpgradeTaskList.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(
			"WebAppService.getRemoteUpgradeTaskList",
			getRemoteUpgradeTaskList.RequestData{
				PsIdList: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getRemoteUpgradeTaskList.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// QuerySysAdvancedParam - AppService.querySysAdvancedParam
func (sg *SunGrow) QuerySysAdvancedParam(psId valueTypes.PsId, curPage valueTypes.Integer, size valueTypes.Integer) (querySysAdvancedParam.ResultData, error) {
	var ret querySysAdvancedParam.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(querySysAdvancedParam.EndPointName,
			querySysAdvancedParam.RequestData{
				PsId2:   psId,
				CurPage: curPage,
				Size:    size,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := querySysAdvancedParam.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// GetPsInstallerByPsId - AppService.getPsInstallerByPsId
func (sg *SunGrow) GetPsInstallerByPsId(psId valueTypes.PsId) (getPsInstallerByPsId.ResultData, error) {
	var ret getPsInstallerByPsId.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsInstallerByPsId.EndPointName,
			getPsInstallerByPsId.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPsInstallerByPsId.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// FindPsType - AppService.findPsType
func (sg *SunGrow) FindPsType(psId valueTypes.PsId) (findPsType.ResultData, error) {
	var ret findPsType.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(findPsType.EndPointName,
			findPsType.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := findPsType.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// GetChannelListByPsId - AppService.getChnnlListByPsId
func (sg *SunGrow) GetChannelListByPsId(psId valueTypes.PsId) (getChnnlListByPsId.ResultData, error) {
	var ret getChnnlListByPsId.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getChnnlListByPsId.EndPointName,
			getChnnlListByPsId.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getChnnlListByPsId.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetHouseholdStoragePsReport - AppService.getHouseholdStoragePsReport DateId:2022
func (sg *SunGrow) GetHouseholdStoragePsReport(psId valueTypes.PsId, dateId valueTypes.DateTime) (getHouseholdStoragePsReport.ResultData, error) {
	var ret getHouseholdStoragePsReport.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getHouseholdStoragePsReport.EndPointName,
			getHouseholdStoragePsReport.RequestData{
				PsId:     psId,
				DateId:   dateId,
				DateType: valueTypes.SetStringValue(dateId.DateType),
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getHouseholdStoragePsReport.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetIncomeSettingInfos - AppService.getIncomeSettingInfos
func (sg *SunGrow) GetIncomeSettingInfos(psId valueTypes.PsId) (getIncomeSettingInfos.ResultData, error) {
	var ret getIncomeSettingInfos.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getIncomeSettingInfos.EndPointName,
			getIncomeSettingInfos.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getIncomeSettingInfos.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPowerChargeSettingInfo - AppService.getPowerChargeSettingInfo
func (sg *SunGrow) GetPowerChargeSettingInfo(psId valueTypes.PsId) (getPowerChargeSettingInfo.ResultData, error) {
	var ret getPowerChargeSettingInfo.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPowerChargeSettingInfo.EndPointName,
			getPowerChargeSettingInfo.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPowerChargeSettingInfo.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPowerStationBasicInfo - AppService.getPowerStationBasicInfo
func (sg *SunGrow) GetPowerStationBasicInfo(psId valueTypes.PsId) (getPowerStationBasicInfo.ResultData, error) {
	var ret getPowerStationBasicInfo.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPowerStationBasicInfo.EndPointName,
			getPowerStationBasicInfo.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPowerStationBasicInfo.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPowerStationData - AppService.getPowerStationData DateId:2022
func (sg *SunGrow) GetPowerStationData(psId valueTypes.PsId, dateId valueTypes.DateTime) (getPowerStationData.ResultData, error) {
	var ret getPowerStationData.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPowerStationData.EndPointName,
			getPowerStationData.RequestData{
				PsId:     psId,
				DateId:   dateId,
				DateType: valueTypes.SetStringValue(dateId.DateType),
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPowerStationData.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPowerStationForHousehold - AppService.getPowerStationForHousehold
func (sg *SunGrow) GetPowerStationForHousehold(psId valueTypes.PsId) (getPowerStationForHousehold.ResultData, error) {
	var ret getPowerStationForHousehold.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPowerStationForHousehold.EndPointName,
			getPowerStationForHousehold.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPowerStationForHousehold.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPowerStationInfo - AppService.getPowerStationInfo
func (sg *SunGrow) GetPowerStationInfo(psId valueTypes.PsId) (getPowerStationInfo.ResultData, error) {
	var ret getPowerStationInfo.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPowerStationInfo.EndPointName,
			getPowerStationInfo.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPowerStationInfo.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPowerStatistics - AppService.getPowerStatistics
func (sg *SunGrow) GetPowerStatistics(psId valueTypes.PsId) (getPowerStatistics.ResultData, error) {
	var ret getPowerStatistics.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPowerStatistics.EndPointName,
			getPowerStatistics.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPowerStatistics.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsDetail - AppService.getPsDetail
func (sg *SunGrow) GetPsDetail(psId valueTypes.PsId) (getPsDetail.ResultData, error) {
	var ret getPsDetail.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsDetail.EndPointName,
			getPsDetail.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPsDetail.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsDetailForSinglePage - AppService.getPsDetailForSinglePage
func (sg *SunGrow) GetPsDetailForSinglePage(psId valueTypes.PsId) (getPsDetailForSinglePage.ResultData, error) {
	var ret getPsDetailForSinglePage.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsDetailForSinglePage.EndPointName,
			getPsDetailForSinglePage.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPsDetailForSinglePage.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsDetailWithPsType - AppService.getPsDetailWithPsType
func (sg *SunGrow) GetPsDetailWithPsType(psId valueTypes.PsId) (getPsDetailWithPsType.ResultData, error) {
	var ret getPsDetailWithPsType.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsDetailWithPsType.EndPointName,
			getPsDetailWithPsType.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPsDetailWithPsType.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsHealthState - AppService.getPsHealthState
func (sg *SunGrow) GetPsHealthState(psId valueTypes.PsId) (getPsHealthState.ResultData, error) {
	var ret getPsHealthState.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsHealthState.EndPointName,
			getPsHealthState.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPsHealthState.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsWeatherList - AppService.getPsWeatherList
func (sg *SunGrow) GetPsWeatherList(psId valueTypes.PsId) (getPsWeatherList.ResultData, error) {
	var ret getPsWeatherList.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsWeatherList.EndPointName,
			getPsWeatherList.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPsWeatherList.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetReportData - AppService.getReportData ReportType:1
func (sg *SunGrow) GetReportData(psId valueTypes.PsId, reportType int64) (getReportData.ResultData, error) {
	var ret getReportData.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getReportData.EndPointName,
			getReportData.RequestData{
				PsId:       psId,
				ReportType: valueTypes.SetIntegerValue(reportType),
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getReportData.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsForecastInfo - AppService.psForcastInfo
func (sg *SunGrow) GetPsForecastInfo(psId valueTypes.PsId) (psForcastInfo.ResultData, error) {
	var ret psForcastInfo.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(psForcastInfo.EndPointName,
			psForcastInfo.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := psForcastInfo.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsHourPointsValue - AppService.psHourPointsValue
func (sg *SunGrow) GetPsHourPointsValue(psId valueTypes.PsId) (psHourPointsValue.ResultData, error) {
	var ret psHourPointsValue.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(psHourPointsValue.EndPointName,
			psHourPointsValue.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := psHourPointsValue.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// QueryDeviceListForApp - AppService.queryDeviceListForApp
func (sg *SunGrow) QueryDeviceListForApp(psId valueTypes.PsId) (queryDeviceListForApp.ResultData, error) {
	var ret queryDeviceListForApp.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(queryDeviceListForApp.EndPointName,
			queryDeviceListForApp.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := queryDeviceListForApp.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// QueryPsProfit - AppService.queryPsProfit DateId:2022
func (sg *SunGrow) QueryPsProfit(psId valueTypes.PsId, dateId valueTypes.DateTime) (queryPsProfit.ResultData, error) {
	var ret queryPsProfit.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(queryPsProfit.EndPointName,
			queryPsProfit.RequestData{
				PsId:     psId,
				DateId:   dateId,
				DateType: valueTypes.SetStringValue(dateId.DateType),
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := queryPsProfit.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// QueryPsStructureList - AppService.queryPsStructureList
func (sg *SunGrow) QueryPsStructureList(psId valueTypes.PsId) (queryPsStructureList.ResultData, error) {
	var ret queryPsStructureList.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(queryPsStructureList.EndPointName,
			queryPsStructureList.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := queryPsStructureList.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// ReportList - AppService.reportList ReportType:1
func (sg *SunGrow) ReportList(psId valueTypes.PsId, reportType int64) (reportList.ResultData, error) {
	var ret reportList.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(reportList.EndPointName,
			reportList.RequestData{
				PsId:       psId,
				ReportType: valueTypes.SetIntegerValue(reportType),
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := reportList.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsDeviceListValue - MttvScreenService.getPsDeviceListValue
func (sg *SunGrow) GetPsDeviceListValue(psId valueTypes.PsId) (getPsDeviceListValue.ResultData, error) {
	var ret getPsDeviceListValue.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsDeviceListValue.EndPointName,
			getPsDeviceListValue.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPsDeviceListValue.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsKpiForHoursByPsId - MttvScreenService.getPsKpiForHoursByPsId Day:20221002
func (sg *SunGrow) GetPsKpiForHoursByPsId(psId valueTypes.PsId) (getPsKpiForHoursByPsId.ResultData, error) {
	var ret getPsKpiForHoursByPsId.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsKpiForHoursByPsId.EndPointName,
			getPsKpiForHoursByPsId.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPsKpiForHoursByPsId.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsIdState - WebAppService.getPsIdState
func (sg *SunGrow) GetPsIdState(psId valueTypes.PsId) (getPsIdState.ResultData, error) {
	var ret getPsIdState.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsIdState.EndPointName,
			getPsIdState.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPsIdState.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetReportPsTree - WebAppService.getReportPsTree DeviceType:14
func (sg *SunGrow) GetReportPsTree(psId valueTypes.PsId) (getReportPsTree.ResultData, error) {
	var ret getReportPsTree.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getReportPsTree.EndPointName,
			getReportPsTree.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getReportPsTree.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// ShowPSView - WebAppService.showPSView
func (sg *SunGrow) ShowPSView(psId valueTypes.PsId) (showPSView.ResultData, error) {
	var ret showPSView.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(showPSView.EndPointName,
			showPSView.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := showPSView.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetMaxDeviceIdByPsId - WebIscmAppService.getMaxDeviceIdByPsId
func (sg *SunGrow) GetMaxDeviceIdByPsId(psId valueTypes.PsId) (getMaxDeviceIdByPsId.ResultData, error) {
	var ret getMaxDeviceIdByPsId.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getMaxDeviceIdByPsId.EndPointName,
			getMaxDeviceIdByPsId.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getMaxDeviceIdByPsId.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsTreeMenu - WebIscmAppService.getPsTreeMenu
func (sg *SunGrow) GetPsTreeMenu(psId valueTypes.PsId) (getPsTreeMenu.ResultData, error) {
	var ret getPsTreeMenu.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsTreeMenu.EndPointName,
			getPsTreeMenu.RequestData{
				PsId: psId,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		data := getPsTreeMenu.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// // getRobotDynamicCleaningView - AppService.getRobotDynamicCleaningView
// func (sg *SunGrow) getRobotDynamicCleaningView(psId valueTypes.PsId) (getRobotDynamicCleaningView.ResultData, error) {
// 	var ret getRobotDynamicCleaningView.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetPowerRobotSweepAttrByPsId - AppService.getPowerRobotSweepAttrByPsId
// func (sg *SunGrow) GetPowerRobotSweepAttrByPsId(psId valueTypes.PsId) (getPowerRobotSweepAttrByPsId.ResultData, error) {
// 	var ret getPowerRobotSweepAttrByPsId.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // FindEnvironmentInfo - AppService.findEnvironmentInfo
// func (sg *SunGrow) FindEnvironmentInfo(psId valueTypes.PsId) (findEnvironmentInfo.ResultData, error) {
// 	var ret findEnvironmentInfo.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetAllDeviceByPsId - AppService.getAllDeviceByPsId
// func (sg *SunGrow) GetAllDeviceByPsId(psId valueTypes.PsId) (getAllDeviceByPsId.ResultData, error) {
// 	var ret getAllDeviceByPsId.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetAllPowerRobotViewInfoByPsId - AppService.getAllPowerRobotViewInfoByPsId
// func (sg *SunGrow) GetAllPowerRobotViewInfoByPsId(psId valueTypes.PsId) (getAllPowerRobotViewInfoByPsId.ResultData, error) {
// 	var ret getAllPowerRobotViewInfoByPsId.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // QueryPsNameByPsId - AppService.queryPsNameByPsId
// func (sg *SunGrow) QueryPsNameByPsId(psId valueTypes.PsId) (queryPsNameByPsId.ResultData, error) {
// 	var ret queryPsNameByPsId.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetTheoryAndActualPower - MttvScreenService.getTheoryAndActualPower
// func (sg *SunGrow) GetTheoryAndActualPower(psId valueTypes.PsId) (getTheoryAndActualPower.ResultData, error) {
// 	var ret getTheoryAndActualPower.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetInverterInfo - WebAppService.getInverterInfo
// func (sg *SunGrow) GetInverterInfo(psId valueTypes.PsId) (getInverterInfo.ResultData, error) {
// 	var ret getInverterInfo.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetMultiPowers - WebAppService.getMultiPowers
// func (sg *SunGrow) GetMultiPowers(psId valueTypes.PsId) (getMultiPowers.ResultData, error) {
// 	var ret getMultiPowers.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetPsCBoxDetail - WebAppService.getPsCBoxDetail
// func (sg *SunGrow) GetPsCBoxDetail(psId valueTypes.PsId) (getPsCBoxDetail.ResultData, error) {
// 	var ret getPsCBoxDetail.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // QueryBatteryBoardsList - WebAppService.queryBatteryBoardsList
// func (sg *SunGrow) QueryBatteryBoardsList(psId valueTypes.PsId) (queryBatteryBoardsList.ResultData, error) {
// 	var ret queryBatteryBoardsList.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // ShowTjReport - WebAppService.showTjReport
// func (sg *SunGrow) ShowTjReport(psId valueTypes.PsId) (showTjReport.ResultData, error) {
// 	var ret showTjReport.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetAuthKeyList - WebIscmAppService.getAuthKeyList
// func (sg *SunGrow) GetAuthKeyList(psId valueTypes.PsId) (getAuthKeyList.ResultData, error) {
// 	var ret getAuthKeyList.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }
