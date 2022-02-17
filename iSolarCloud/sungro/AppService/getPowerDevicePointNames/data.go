package getPowerDevicePointNames


var Url = "/v1/reportService/getPowerDevicePointNames"

type ResultData   []struct {
	PointCalType int64  `json:"point_cal_type"`
	PointID      int64  `json:"point_id"`
	PointName    string `json:"point_name"`
}
