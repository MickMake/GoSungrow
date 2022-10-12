package mmHa

import (
	"regexp"
	"strings"
)


func JoinStrings(args ...string) string {
	return strings.TrimSpace(strings.Join(args, " "))
}

func JoinStringsForName(sep string, args ...string) string {
	var newargs []string
	var re = regexp.MustCompile(`(\.)+`)
	for _, a := range args {
		if a == "" {
			continue
		}
		a = strings.TrimSpace(a)
		a = re.ReplaceAllString(a, ` `)
		newargs = append(newargs, a)
	}

	ret := strings.TrimSpace(strings.Join(newargs, sep))
	return ret
}

func JoinStringsForId(args ...string) string {
	var newargs []string
	var re = regexp.MustCompile(`(/| |:|\.)+`)
	for _, a := range args {
		if a == "" {
			continue
		}
		a = strings.TrimSpace(a)
		a = re.ReplaceAllString(a, `-`)
		newargs = append(newargs, a)
	}
	// return strings.ReplaceAll(strings.TrimSpace(strings.Join(args, ".")), ".", "_")
	return strings.Join(newargs, "-")
}

// func (c *Config) JoinStringsForId() string {
// 	return JoinStringsForId(m.Device.FullName, c.ParentName, c.FullName)
// }

func JoinStringsForTopic(args ...string) string {
	var newargs []string
	var re = regexp.MustCompile(`( |:)+`)
	for _, a := range args {
		if a == "" {
			continue
		}
		a = strings.TrimSpace(a)
		a = re.ReplaceAllString(a, `_`)
		newargs = append(newargs, a)
	}
	// return strings.ReplaceAll(strings.TrimSpace(strings.Join(args, ".")), ".", "_")
	return strings.Join(newargs, "/")

	// ret := strings.ReplaceAll(strings.Join(args, "/"), "//", "/")
	// return ret
}


// const DiscoveryPrefix = "homeassistant"
//
// var (
// 	retain bool = true
// 	qos    byte = 0
// )
//
// func getDevice() (d Device) {
//
// 	// id, err := machineid.ProtectedID(NodeID)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
//
// 	d.Identifiers = []string{"id"}
// 	d.Manufacturer = "MickMake"
// 	d.Model = "NodeID"
// 	d.Name = "InstanceName"
// 	d.SwVersion = "SWVersion"
//
// 	return
// }
//
// ///////////////////
//
// type store struct {
// 	BinarySensor map[string]string
// 	Light        struct {
// 		BrightnessState map[string]string
// 		ColorTempState  map[string]string
// 		EffectState     map[string]string
// 		HsState         map[string]string
// 		RgbState        map[string]string
// 		State           map[string]string
// 		WhiteValueState map[string]string
// 		XyState         map[string]string
// 	}
// 	Sensor map[string]string
// 	Switch map[string]string
// }
//
// func initStore() store {
// 	var s store
// 	s.BinarySensor = make(map[string]string)
// 	s.Light.BrightnessState = make(map[string]string)
// 	s.Light.ColorTempState = make(map[string]string)
// 	s.Light.EffectState = make(map[string]string)
// 	s.Light.HsState = make(map[string]string)
// 	s.Light.RgbState = make(map[string]string)
// 	s.Light.State = make(map[string]string)
// 	s.Light.WhiteValueState = make(map[string]string)
// 	s.Light.XyState = make(map[string]string)
// 	s.Sensor = make(map[string]string)
// 	s.Switch = make(map[string]string)
// 	return s
// }
//
// var topicStore = make(map[string]*func(mqtt.Message, mqtt.Client))
// var stateStore = initStore()
//
// ///////////////////
//
// // GetTopicPrefix gets the prefix for all state/command topics
// // This is for a light
// func (device Light) GetTopicPrefix() string {
// 	return NodeID + "/light/" + device.UniqueID + "/"
// }
//
// // GetTopicPrefix gets the prefix for all state/command topics
// // This is for a sensor
// func (device Sensor) GetTopicPrefix() string {
// 	return NodeID + "/sensor/" + device.UniqueID + "/"
// }
//
// // GetTopicPrefix gets the prefix for all state/command topics
// // This is for a switch
// func (device Switch) GetTopicPrefix() string {
// 	return NodeID + "/switch/" + device.UniqueID + "/"
// }
//
// // GetTopicPrefix gets the prefix for all state/command topics
// // This is for a binary sensor
// func (device BinarySensor) GetTopicPrefix() string {
// 	return NodeID + "/binary_sensor/" + device.UniqueID + "/"
// }
//
// // GetDiscoveryTopic gets the topic for a device's discovery topic.
// // This is for a light
// func (device Light) GetDiscoveryTopic() string {
// 	return DiscoveryPrefix + "/light/" + NodeID + "/" + device.UniqueID + "/" + "config"
// }
//
// // GetDiscoveryTopic gets the topic for a device's discovery topic.
// // This is for a sensor
// func (device Sensor) GetDiscoveryTopic() string {
// 	return DiscoveryPrefix + "/sensor/" + NodeID + "/" + device.UniqueID + "/" + "config"
// }
//
// // GetDiscoveryTopic gets the topic for a device's discovery topic.
// // This is for a switch
// func (device Switch) GetDiscoveryTopic() string {
// 	return DiscoveryPrefix + "/switch/" + NodeID + "/" + device.UniqueID + "/" + "config"
// }
//
// // GetDiscoveryTopic gets the topic for a device's discovery topic.
// // This is for a binary sensor
// func (device BinarySensor) GetDiscoveryTopic() string {
// 	return DiscoveryPrefix + "/binary_sensor/" + NodeID + "/" + device.UniqueID + "/" + "config"
// }
//
// // GetCommandTopic gets the command topic for a device
// // This is for a light
// func (device Light) GetCommandTopic() string {
// 	return device.GetTopicPrefix() + "command"
// }
//
// // GetCommandTopic gets the command topic for a device
// // This is for a switch
// func (device Switch) GetCommandTopic() string {
// 	return device.GetTopicPrefix() + "command"
// }
//
// // GetStateTopic gets the state topic for a device
// // This is for a light
// func (device Light) GetStateTopic() string {
// 	return device.GetTopicPrefix() + "state"
// }
//
// // GetStateTopic gets the state topic for a device
// // This is for a sensor
// func (device Sensor) GetStateTopic() string {
// 	return device.GetTopicPrefix() + "state"
// }
//
// // GetStateTopic gets the state topic for a device
// // This is for a switch
// func (device Switch) GetStateTopic() string {
// 	return device.GetTopicPrefix() + "state"
// }
//
// // GetStateTopic gets the state topic for a device
// // This is for a binary sensor
// func (device BinarySensor) GetStateTopic() string {
// 	return device.GetTopicPrefix() + "state"
// }
//
// // GetAvailabilityTopic gets the availability topic for a device
// // This is for a light
// func (device Light) GetAvailabilityTopic() string {
// 	return device.GetTopicPrefix() + "availability"
// }
//
// // GetAvailabilityTopic gets the availability topic for a device
// // This is for a sensor
// func (device Sensor) GetAvailabilityTopic() string {
// 	return device.GetTopicPrefix() + "availability"
// }
//
// // GetAvailabilityTopic gets the availability topic for a device
// // This is for a switch
// func (device Switch) GetAvailabilityTopic() string {
// 	return device.GetTopicPrefix() + "availability"
// }
//
// // GetAvailabilityTopic gets the availability topic for a device
// // This is for a binary sensor
// func (device BinarySensor) GetAvailabilityTopic() string {
// 	return device.GetTopicPrefix() + "availability"
// }
//
// // Initialize sets topics as needed on a Light
// func (device *Light) Initialize() {
// 	device.Retain = false
// 	device.Device = getDevice()
//
// 	device.AvailabilityTopic = device.GetAvailabilityTopic()
//
// 	// Brightness
// 	if device.BrightnessCommandFunc != nil {
// 		device.BrightnessCommandTopic = device.GetTopicPrefix() + "brightness/set"
// 		topicStore[device.BrightnessCommandTopic] = &device.BrightnessCommandFunc
// 	}
// 	if device.BrightnessStateFunc != nil {
// 		device.BrightnessStateTopic = device.GetTopicPrefix() + "brightness/get"
// 	}
//
// 	// ColorTemp
// 	if device.ColorTempCommandFunc != nil {
// 		device.ColorTempCommandTopic = device.GetTopicPrefix() + "color-temp/set"
// 		topicStore[device.ColorTempCommandTopic] = &device.ColorTempCommandFunc
// 	}
// 	if device.ColorTempStateFunc != nil {
// 		device.ColorTempStateTopic = device.GetTopicPrefix() + "color-temp/get"
// 	}
//
// 	// Command/State
// 	if device.CommandFunc != nil {
// 		device.CommandTopic = device.GetCommandTopic()
// 		topicStore[device.CommandTopic] = &device.CommandFunc
// 	}
// 	if device.StateFunc != nil {
// 		device.StateTopic = device.GetStateTopic()
// 	}
//
// 	// Effect
// 	if device.EffectCommandFunc != nil {
// 		device.EffectCommandTopic = device.GetTopicPrefix() + "effect/set"
// 		topicStore[device.EffectCommandTopic] = &device.EffectCommandFunc
// 	}
// 	if device.EffectStateFunc != nil {
// 		device.EffectStateTopic = device.GetTopicPrefix() + "effect/get"
// 	}
//
// 	// Hs
// 	if device.HsCommandFunc != nil {
// 		device.HsCommandTopic = device.GetTopicPrefix() + "hs/set"
// 		topicStore[device.HsCommandTopic] = &device.HsCommandFunc
// 	}
// 	if device.HsStateFunc != nil {
// 		device.HsStateTopic = device.GetTopicPrefix() + "hs/get"
// 	}
//
// 	// Rgb
// 	if device.RgbCommandFunc != nil {
// 		device.RgbCommandTopic = device.GetTopicPrefix() + "rgb/set"
// 		topicStore[device.RgbCommandTopic] = &device.RgbCommandFunc
// 	}
// 	if device.RgbStateFunc != nil {
// 		device.RgbStateTopic = device.GetTopicPrefix() + "rgb/get"
// 	}
//
// 	// WhiteValue
// 	if device.WhiteValueCommandFunc != nil {
// 		device.WhiteValueCommandTopic = device.GetTopicPrefix() + "white-value/set"
// 		topicStore[device.WhiteValueCommandTopic] = &device.WhiteValueCommandFunc
// 	}
// 	if device.WhiteValueStateFunc != nil {
// 		device.WhiteValueStateTopic = device.GetTopicPrefix() + "white-value/get"
// 	}
//
// 	// Xy
// 	if device.XyCommandFunc != nil {
// 		device.XyCommandTopic = device.GetTopicPrefix() + "xy/set"
// 		topicStore[device.XyCommandTopic] = &device.XyCommandFunc
// 	}
// 	if device.XyStateFunc != nil {
// 		device.XyStateTopic = device.GetTopicPrefix() + "xy/get"
// 	}
//
// 	device.messageHandler = func(client mqtt.Client, msg mqtt.Message) {
//
// 		topicFound := false
//
// 		for topic, f := range topicStore {
// 			if msg.Topic() == topic {
// 				topicFound = true
// 				(*f)(msg, client)
// 				device.UpdateState(client)
// 			}
// 		}
//
// 		if !topicFound {
// 			log.Println("Unknown Message on topic " + msg.Topic())
// 			log.Println(msg.Payload())
// 		}
// 	}
// }
//
// // Initialize sets topics as needed on a Switch
// func (device *Switch) Initialize() {
// 	device.CommandTopic = device.GetCommandTopic()
// 	if device.StateFunc != nil {
// 		device.StateTopic = device.GetStateTopic()
// 	}
// 	device.AvailabilityTopic = device.GetAvailabilityTopic()
// 	device.Device = getDevice()
// 	device.Retain = false
//
// 	topicStore[device.CommandTopic] = &device.CommandFunc
//
// 	device.messageHandler = func(client mqtt.Client, msg mqtt.Message) {
//
// 		topicFound := false
//
// 		for topic, f := range topicStore {
// 			if msg.Topic() == topic {
// 				topicFound = true
// 				(*f)(msg, client)
// 				device.UpdateState(client)
// 			}
// 		}
//
// 		if !topicFound {
// 			log.Println("Unknown Message on topic " + msg.Topic())
// 			log.Println(msg.Payload())
// 		}
// 	}
//
// }
//
// // Initialize sets topics as needed on a Sensor
// func (device *Sensor) Initialize() {
// 	device.StateTopic = device.GetStateTopic()
// 	device.AvailabilityTopic = device.GetAvailabilityTopic()
// 	device.Device = getDevice()
// }
//
// // Initialize sets topics as needed on a Binary Sensor
// func (device *BinarySensor) Initialize() {
// 	device.StateTopic = device.GetStateTopic()
// 	device.AvailabilityTopic = device.GetAvailabilityTopic()
// 	device.Device = getDevice()
// }
//
// // UpdateState publishes any new state for a device
// // This is for a light
// func (device Light) UpdateState(client mqtt.Client) {
// 	if device.BrightnessStateTopic != "" {
// 		state := device.BrightnessStateFunc()
// 		if state != stateStore.Light.BrightnessState[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.BrightnessStateTopic, qos, retain, state)
// 			stateStore.Light.BrightnessState[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// 	if device.ColorTempStateTopic != "" {
// 		state := device.ColorTempStateFunc()
// 		if state != stateStore.Light.ColorTempState[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.ColorTempStateTopic, qos, retain, state)
// 			stateStore.Light.ColorTempState[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// 	if device.EffectStateTopic != "" {
// 		state := device.EffectStateFunc()
// 		if state != stateStore.Light.EffectState[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.EffectStateTopic, qos, retain, state)
// 			stateStore.Light.EffectState[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// 	if device.HsStateTopic != "" {
// 		state := device.HsStateFunc()
// 		if state != stateStore.Light.HsState[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.HsStateTopic, qos, retain, state)
// 			stateStore.Light.HsState[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// 	if device.RgbStateTopic != "" {
// 		state := device.RgbStateFunc()
// 		if state != stateStore.Light.RgbState[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.RgbStateTopic, qos, retain, state)
// 			stateStore.Light.RgbState[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// 	if device.StateTopic != "" {
// 		state := device.StateFunc()
// 		if state != stateStore.Light.State[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.StateTopic, qos, retain, state)
// 			stateStore.Light.State[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// 	if device.WhiteValueStateTopic != "" {
// 		state := device.WhiteValueStateFunc()
// 		if state != stateStore.Light.WhiteValueState[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.WhiteValueStateTopic, qos, retain, state)
// 			stateStore.Light.WhiteValueState[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// 	if device.XyStateTopic != "" {
// 		state := device.XyStateFunc()
// 		if state != stateStore.Light.XyState[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.XyStateTopic, qos, retain, state)
// 			stateStore.Light.XyState[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// }
//
// // UpdateState publishes any new state for a device
// // This is for a switch
// func (device Switch) UpdateState(client mqtt.Client) {
// 	if device.StateFunc != nil {
// 		state := device.StateFunc()
// 		if state != stateStore.Switch[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.StateTopic, qos, retain, state)
// 			stateStore.Switch[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	} else {
// 		log.Println("No statefunc for " + device.UniqueID + strconv.FormatFloat(float64(device.UpdateInterval), 'g', 1, 64))
// 	}
// }
//
// // UpdateState publishes any new state for a device
// // This is for a sensor
// func (device Sensor) UpdateState(client mqtt.Client) {
// 	if device.StateFunc != nil {
// 		state := device.StateFunc()
// 		if state != stateStore.Sensor[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.StateTopic, qos, retain, state)
// 			token.Wait()
// 			stateStore.Sensor[device.UniqueID] = state
// 		}
// 	} else {
// 		log.Fatalln("No statefunc, this makes no sensor for a sensor!")
// 	}
// }
//
// // UpdateState publishes any new state for a device
// // This is for a binary sensor
// func (device BinarySensor) UpdateState(client mqtt.Client) {
// 	if device.StateFunc != nil {
// 		state := device.StateFunc()
// 		if state != stateStore.BinarySensor[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.StateTopic, qos, retain, state)
// 			token.Wait()
// 			stateStore.BinarySensor[device.UniqueID] = state
// 		}
// 	} else {
// 		log.Fatalln("No statefunc, this makes no sensor for a binary sensor!")
// 	}
// }
//
// // Subscribe announces and starts listening to MQTT topics appropriate for a device
// // This is for a light
// func (device Light) Subscribe(client mqtt.Client) {
//
// 	message, err := json.Marshal(device)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	if device.BrightnessCommandTopic != "" {
// 		if token := client.Subscribe(device.BrightnessCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
// 	if device.ColorTempCommandTopic != "" {
// 		if token := client.Subscribe(device.ColorTempCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
// 	if device.EffectCommandTopic != "" {
// 		if token := client.Subscribe(device.EffectCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
// 	if device.HsCommandTopic != "" {
// 		if token := client.Subscribe(device.HsCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
// 	if device.RgbCommandTopic != "" {
// 		if token := client.Subscribe(device.RgbCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
// 	if device.CommandTopic != "" {
// 		if token := client.Subscribe(device.CommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
// 	if device.WhiteValueCommandTopic != "" {
// 		if token := client.Subscribe(device.WhiteValueCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
// 	if device.XyCommandTopic != "" {
// 		if token := client.Subscribe(device.XyCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
//
// 	token := client.Publish(device.GetDiscoveryTopic(), 0, true, message)
// 	token.Wait()
//
// 	// HA needs time to process
// 	time.Sleep(500 * time.Millisecond)
//
// 	device.AnnounceAvailable(client)
//
// 	device.UpdateState(client)
//
// 	if token := client.Subscribe(device.GetCommandTopic(), 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 		log.Println(token.Error())
// 		os.Exit(1)
// 	}
//
// }
//
// // Subscribe announces and starts listening to MQTT topics appropriate for a device
// // This is for a switch
// func (device Switch) Subscribe(client mqtt.Client) {
//
// 	message, err := json.Marshal(device)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	token := client.Publish(device.GetDiscoveryTopic(), 0, true, message)
// 	token.Wait()
//
// 	// HA needs time to process
// 	time.Sleep(500 * time.Millisecond)
//
// 	device.AnnounceAvailable(client)
//
// 	if device.StateFunc != nil {
// 		device.UpdateState(client)
// 	}
//
// 	if token := client.Subscribe(device.GetCommandTopic(), 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 		log.Println(token.Error())
// 		os.Exit(1)
// 	}
//
// }
//
// // Subscribe announces and MQTT topics appropriate for a device
// // This is for a sensor
// func (device Sensor) Subscribe(client mqtt.Client) {
//
// 	message, err := json.Marshal(device)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	token := client.Publish(device.GetDiscoveryTopic(), 0, true, message)
// 	token.Wait()
//
// 	// HA needs time to process
// 	time.Sleep(500 * time.Millisecond)
//
// 	device.UpdateState(client)
// 	device.AnnounceAvailable(client)
//
// }
//
// // Subscribe announces and MQTT topics appropriate for a device
// // This is for a binary sensor
// func (device BinarySensor) Subscribe(client mqtt.Client) {
//
// 	message, err := json.Marshal(device)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	token := client.Publish(device.GetDiscoveryTopic(), 0, true, message)
// 	token.Wait()
//
// 	// HA needs time to process
// 	time.Sleep(500 * time.Millisecond)
//
// 	device.UpdateState(client)
// 	device.AnnounceAvailable(client)
//
// }
//
// // UnSubscribe from MQTT topics appropriate for a device, and publishes unavailability
// // This is for a light
// func (device Light) UnSubscribe(client mqtt.Client) {
// 	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "offline")
// 	token.Wait()
//
// 	if device.BrightnessCommandTopic != "" {
// 		if token := client.Unsubscribe(device.BrightnessCommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
//
// 	}
// 	if device.ColorTempCommandTopic != "" {
// 		if token := client.Unsubscribe(device.ColorTempCommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
//
// 	}
// 	if device.EffectCommandTopic != "" {
// 		if token := client.Unsubscribe(device.EffectCommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
//
// 	}
// 	if device.HsCommandTopic != "" {
// 		if token := client.Unsubscribe(device.HsCommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
//
// 	}
// 	if device.RgbCommandTopic != "" {
// 		if token := client.Unsubscribe(device.RgbCommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
//
// 	}
// 	if device.CommandTopic != "" {
// 		if token := client.Unsubscribe(device.CommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
//
// 	}
// 	if device.WhiteValueCommandTopic != "" {
// 		if token := client.Unsubscribe(device.WhiteValueCommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
//
// 	}
// 	if device.XyCommandTopic != "" {
// 		if token := client.Unsubscribe(device.XyCommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
//
// 	}
//
// }
//
// // UnSubscribe from MQTT topics appropriate for a device, and publishes unavailability
// // This is for a switch
// func (device Switch) UnSubscribe(client mqtt.Client) {
// 	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "offline")
// 	token.Wait()
//
// 	if token := client.Unsubscribe(device.GetCommandTopic()); token.Wait() && token.Error() != nil {
// 		log.Println(token.Error())
// 		os.Exit(1)
// 	}
// }
//
// // UnSubscribe publishes unavailability for a device
// // This is for a sensor
// func (device Sensor) UnSubscribe(client mqtt.Client) {
// 	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "offline")
// 	token.Wait()
// }
//
// // UnSubscribe publishes unavailability for a device
// // This is for a binary sensor
// func (device BinarySensor) UnSubscribe(client mqtt.Client) {
// 	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "offline")
// 	token.Wait()
// }
//
// // AnnounceAvailable publishes availability for a device
// // This is for a light
// func (device Light) AnnounceAvailable(client mqtt.Client) {
// 	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "online")
// 	token.Wait()
// }
//
// // AnnounceAvailable publishes availability for a device
// // This is for a switch
// func (device Switch) AnnounceAvailable(client mqtt.Client) {
// 	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "online")
// 	token.Wait()
// }
//
// // AnnounceAvailable publishes availability for a device
// // This is for a sensor
// func (device Sensor) AnnounceAvailable(client mqtt.Client) {
// 	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "online")
// 	token.Wait()
// }
//
// // AnnounceAvailable publishes availability for a device
// // This is for a binary sensor
// func (device BinarySensor) AnnounceAvailable(client mqtt.Client) {
// 	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "online")
// 	token.Wait()
// }
