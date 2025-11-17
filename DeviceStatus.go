package main

import "github.com/nasa9084/go-switchbot/v5"

type DeviceStatus struct {
	ID                     string                         `json:"deviceId"`
	Type                   switchbot.PhysicalDeviceType   `json:"deviceType"`
	Hub                    string                         `json:"hubDeviceId"`
	Power                  switchbot.PowerState           `json:"power,omitempty"`
	Humidity               int                            `json:"humidity,omitempty"`
	Temperature            float64                        `json:"temperature,omitempty"`
	NebulizationEfficiency int                            `json:"nebulizationEfficiency,omitempty"`
	IsAuto                 bool                           `json:"auto,omitempty"`
	IsChildLock            bool                           `json:"childLock,omitempty"`
	IsSound                bool                           `json:"sound,omitempty"`
	IsCalibrated           bool                           `json:"calibrate,omitempty"`
	IsGrouped              bool                           `json:"group,omitempty"`
	IsMoving               bool                           `json:"moving,omitempty"`
	SlidePosition          int                            `json:"slidePosition,omitempty"`
	FanSpeed               int                            `json:"fanSpeed,omitempty"`
	IsMoveDetected         bool                           `json:"moveDetected,omitempty"`
	Brightness             switchbot.BrightnessState      `json:"brightness,omitempty"`
	LightLevel             int                            `json:"lightLevel,omitempty"`
	OpenState              switchbot.OpenState            `json:"openState,omitempty"`
	Color                  string                         `json:"color,omitempty"`
	ColorTemperature       int                            `json:"colorTemperature,omitempty"`
	IsLackWater            bool                           `json:"lackWater,omitempty"`
	Voltage                float64                        `json:"voltage,omitempty"`
	Weight                 float64                        `json:"weight,omitempty"`
	ElectricityOfDay       int                            `json:"electricityOfDay,omitempty"`
	ElectricCurrent        float64                        `json:"electricCurrent,omitempty"`
	LockState              string                         `json:"lockState,omitempty"`
	DoorState              string                         `json:"doorState,omitempty"`
	WorkingStatus          switchbot.CleanerWorkingStatus `json:"workingStatus,omitempty"`
	OnlineStatus           switchbot.CleanerOnlineStatus  `json:"onlineStatus,omitempty"`
	Battery                int                            `json:"battery,omitempty"`
	WaterBaseBattery       int                            `json:"waterBaseBattery,omitempty"`
	TaskType               switchbot.CleanerTaskType      `json:"taskType,omitempty"`
	Version                switchbot.DeviceVersion        `json:"version,omitempty"`
	Direction              string                         `json:"direction,omitempty"`
	CO2                    int                            `json:"CO2,omitempty"`
	Mode                   switchbot.Mode                 `json:"mode,omitempty"`
	NightStatus            switchbot.NightStatus          `json:"nightStatus,omitempty"`
	Oscillation            switchbot.OscillationStatus    `json:"oscillation,omitempty"`
	VerticalOscillation    switchbot.OscillationStatus    `json:"verticalOscillation,omitempty"`
	ChargingStatus         switchbot.ChargingStatus       `json:"chargingStatus,omitempty"`
	DeviceMode             switchbot.BotDeviceMode        `json:"deviceMode,omitempty"`
	LeakStatus             switchbot.WaterLeakStatus      `json:"status,omitempty"`
	UsedElectricity        int                            `json:"usedElectricity,omitempty"`
}

func ParseDeviceStatus(status *switchbot.DeviceStatus) *DeviceStatus {
	return &DeviceStatus{
		ID:                     status.ID,
		Type:                   status.Type,
		Hub:                    status.Hub,
		Power:                  status.Power,
		Humidity:               status.Humidity,
		Temperature:            status.Temperature,
		NebulizationEfficiency: status.NebulizationEfficiency,
		IsAuto:                 status.IsAuto,
		IsChildLock:            status.IsChildLock,
		IsSound:                status.IsSound,
		IsCalibrated:           status.IsCalibrated,
		IsGrouped:              status.IsGrouped,
		IsMoving:               status.IsMoving,
		SlidePosition:          status.SlidePosition,
		FanSpeed:               status.FanSpeed,
		IsMoveDetected:         status.IsMoveDetected,
		Brightness:             status.Brightness,
		LightLevel:             status.LightLevel,
		OpenState:              status.OpenState,
		Color:                  status.Color,
		ColorTemperature:       status.ColorTemperature,
		IsLackWater:            status.IsLackWater,
		Voltage:                status.Voltage,
		Weight:                 status.Weight,
		ElectricityOfDay:       status.ElectricityOfDay,
		ElectricCurrent:        status.ElectricCurrent,
		LockState:              status.LockState,
		DoorState:              status.DoorState,
		WorkingStatus:          status.WorkingStatus,
		OnlineStatus:           status.OnlineStatus,
		Battery:                status.Battery,
		WaterBaseBattery:       status.WaterBaseBattery,
		TaskType:               status.TaskType,
		Version:                status.Version,
		Direction:              status.Direction,
		CO2:                    status.CO2,
		Mode:                   status.Mode,
		NightStatus:            status.NightStatus,
		Oscillation:            status.Oscillation,
		VerticalOscillation:    status.VerticalOscillation,
		ChargingStatus:         status.ChargingStatus,
		DeviceMode:             status.DeviceMode,
		LeakStatus:             status.LeakStatus,
		UsedElectricity:        status.UsedElectricity,
	}
}
