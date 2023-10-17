package cmdHassio

import (
	"encoding/json"
	"fmt"

	"github.com/anicoll/gosungrow/pkg/only"
)

type MqttState struct {
	LastReset string `json:"last_reset,omitempty"`
	Value     string `json:"value"`
}

func (mq *MqttState) Json() string {
	var ret string
	for range only.Once {
		j, err := json.Marshal(*mq)
		if err != nil {
			ret = fmt.Sprintf("{ \"error\": \"%s\"", err)
			break
		}
		ret = string(j)
	}
	return ret
}

type SensorState string
