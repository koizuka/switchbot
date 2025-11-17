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
	}
}
