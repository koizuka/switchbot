package main

import (
	"github.com/nasa9084/go-switchbot/v2"
)

type DeviceItem struct {
	ID                   string                       `json:"deviceId"`
	Name                 string                       `json:"deviceName"`
	Type                 switchbot.PhysicalDeviceType `json:"deviceType"`
	IsEnableCloudService bool                         `json:"enableCloudService,omitempty"`
	Hub                  string                       `json:"hubDeviceId,omitempty"`
	Curtains             []string                     `json:"curtainDeviceesIds,omitempty"`
	IsCalibrated         bool                         `json:"calibrate,omitempty"`
	IsGrouped            bool                         `json:"group,omitempty"`
	IsMaster             bool                         `json:"master,omitempty"`
	OpenDirection        string                       `json:"openDirection,omitempty"`
	GroupName            bool                         `json:"groupName,omitempty"`
	LockDeviceIDs        []string                     `json:"lockDeviceIds,omitempty"`
	LockDeviceID         string                       `json:"lockDeviceId,omitempty"`
	KeyList              []switchbot.KeyListItem      `json:"keyList,omitempty"`
}

func ParseDeviceItem(item *switchbot.Device) *DeviceItem {
	return &DeviceItem{
		ID:                   item.ID,
		Name:                 item.Name,
		Type:                 item.Type,
		IsEnableCloudService: item.IsEnableCloudService,
		Hub:                  item.Hub,
		Curtains:             item.Curtains,
		IsCalibrated:         item.IsCalibrated,
		IsGrouped:            item.IsGrouped,
		IsMaster:             item.IsMaster,
		OpenDirection:        item.OpenDirection,
		GroupName:            item.GroupName,
		LockDeviceIDs:        item.LockDeviceIDs,
		LockDeviceID:         item.LockDeviceID,
		KeyList:              item.KeyList,
	}
}
