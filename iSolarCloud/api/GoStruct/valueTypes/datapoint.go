package valueTypes

import (
	"encoding/json"
	"strings"

	"github.com/MickMake/GoUnify/Only"
)

type DataPoint struct {
	endPoint string  `json:"end_point,omitempty"`
	pointId  PointId `json:"point_id,omitempty"`
	Valid    bool    `json:"valid"`
	Error    error   `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *DataPoint) UnmarshalJSON(data []byte) error {
	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		var r string
		// Store result from string
		t.Error = json.Unmarshal(data, &r)
		if t.Error == nil {
			t.SetEndPoint(r)
			break
		}

		t.SetEndPoint(string(data))
	}

	return t.Error
}

// MarshalJSON - Convert value to JSON
func (t DataPoint) MarshalJSON() ([]byte, error) {
	var data []byte

	for range Only.Once {
		t.Valid = false

		data, t.Error = json.Marshal(t.String())
		if t.Error != nil {
			break
		}
		t.Valid = true
	}

	return data, t.Error
}

func (t DataPoint) EndPoint() string {
	return t.endPoint
}

func (t DataPoint) PointId() PointId {
	return t.pointId
}

func (t DataPoint) String() string {
	return t.endPoint + "." + t.pointId.String()
}

func (t DataPoint) Split() []string {
	return strings.Split(t.String(), ".")
}

func (t DataPoint) last() string {
	a := strings.Split(t.String(), ".")
	return a[len(a)-1]
}

func (t *DataPoint) Set(endPoint string, pointId string) DataPoint {
	for range Only.Once {
		t.endPoint = endPoint
		if pointId == "" {
			pointId = t.last()
		}
		t.pointId = SetPointIdString(pointId)
		t.Valid = true
	}

	return *t
}

func (t *DataPoint) SetPointId(pointId string) DataPoint {
	for range Only.Once {
		if pointId == "" {
			t.Valid = false
			break
		}

		t.pointId = SetPointIdString(pointId)
		t.Valid = true
	}

	return *t
}

func (t *DataPoint) SetEndPoint(endPoint string) DataPoint {
	for range Only.Once {
		t.endPoint = endPoint
		if t.pointId.String() == "" {
			t.SetPointId(t.last())
		}
		t.Valid = true
	}

	return *t
}

func SetDataPoint(endPoint string, pointId string) DataPoint {
	var t DataPoint
	return t.Set(endPoint, pointId)
}

func JoinDataPoint(value ...string) DataPoint {
	var t DataPoint
	v := strings.Join(value, ".")
	return t.SetEndPoint(v)
}

func JoinStrings(value ...string) string {
	return strings.Join(value, ".")
}
