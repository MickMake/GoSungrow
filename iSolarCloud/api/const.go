package api

import "time"

const (
	NullAreaName = "NullArea"
	NullEndPointName = "NullEndpoint"
	DefaultTimeout = time.Second * 60
)

const (
	DeviceType1 = "Inverter"
	DeviceType10 = "String"
	DeviceType11 = "Plant"
	DeviceType12 = "Circuit Protection"
	DeviceType13 = "Splitting Device"
	DeviceType14 = "Energy Storage System"
	DeviceType15 = "Sampling Device"
	DeviceType16 = "EMU"
	DeviceType17 = "Unit"
	DeviceType18 = "Temperature and Humidity Sensor"
	DeviceType19 = "Intelligent Power Distribution Cabinet"
	DeviceType20 = "Display Device"
	DeviceType21 = "AC Power Distributed Cabinet"
	DeviceType22 = "Communication Module"
	DeviceType23 = "System-BMS"
	DeviceType24 = "RackBMS"
	DeviceType25 = "DC-DC"
	DeviceType26 = "Energy Management System"
	DeviceType28 = "Wind Energy Converter"
	DeviceType29 = "SVG"
	DeviceType3 = "Grid-connection Point"
	DeviceType30 = "PT Cabinet"
	DeviceType31 = "Bus Protection"
	DeviceType32 = "Cleaning Robot"
	DeviceType33 = "Direct Current Cabinet"
	DeviceType34 = "Public Measurement and Control"
	DeviceType35 = "Anti-islanding Protection Device"
	DeviceType36 = "Frequency and Voltage Emergency Control Device"
	DeviceType37 = "PCS"
	DeviceType38 = "Cell BMS"
	DeviceType39 = "Power Quality"
	DeviceType4 = "Combiner Box"
	DeviceType40 = "Shuttle"
	DeviceType41 = "Optimizer"
	DeviceType42 = "Tracking axis communication box"
	DeviceType43 = "Battery"
	DeviceType44 = "Battery Cluster Management Unit"
	DeviceType45 = "Local Controller"
	DeviceType46 = "Networking Devices"
	DeviceType47 = "Energy Storage Unit"
	DeviceType48 = "DC Container"
	DeviceType5 = "Meteo Station"
	DeviceType50 = "IO Module"
	DeviceType6 = "Transformer"
	DeviceType7 = "Meter"
	DeviceType8 = "UPS"
	DeviceType9 = "Data Logger"
	DeviceType99 = "Others"
)

const (
	DeviceNameInverter=1
	DeviceNameString=10
	DeviceNamePlant=11
	DeviceNameCircuitProtection=12
	DeviceNameSplittingDevice=13
	DeviceNameEnergyStorageSystem=14
	DeviceNameSamplingDevice=15
	DeviceNameEMU=16
	DeviceNameUnit=17
	DeviceNameTemperatureAndHumiditySensor=18
	DeviceNameIntelligentPowerDistributionCabinet=19
	DeviceNameDisplayDevice=20
	DeviceNameACPowerDistributedCabinet=21
	DeviceNameCommunicationModule=22
	DeviceNameSystemBMS=23
	DeviceNameRackBMS=24
	DeviceNameDCToDC=25
	DeviceNameEnergyManagementSystem=26
	DeviceNameWindEnergyConverter=28
	DeviceNameSVG=29
	DeviceNameGridConnectionPoint=3
	DeviceNamePTCabinet=30
	DeviceNameBusProtection=31
	DeviceNameCleaningRobot=32
	DeviceNameDirectCurrentCabinet=33
	DeviceNamePublicMeasurementandControl=34
	DeviceNameAntiIslandingProtectionDevice=35
	DeviceNameFrequencyAndVoltageEmergencyControlDevice=36
	DeviceNamePCS=37
	DeviceNameCellBMS=38
	DeviceNamePowerQuality=39
	DeviceNameCombinerBox=4
	DeviceNameShuttle=40
	DeviceNameOptimizer=41
	DeviceNameTrackingAxisCommunicationBox=42
	DeviceNameBattery=43
	DeviceNameBatteryClusterManagementUnit=44
	DeviceNameLocalController=45
	DeviceNameNetworkingDevices=46
	DeviceNameEnergyStorageUnit=47
	DeviceNameDCContainer=48
	DeviceNameMeteoStation=5
	DeviceNameIOModule=50
	DeviceNameTransformer=6
	DeviceNameMeter=7
	DeviceNameUPS=8
	DeviceNameDataLogger=9
	DeviceNameOthers=99
)