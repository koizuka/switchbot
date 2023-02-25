package main

import (
	"context"
	"github.com/nasa9084/go-switchbot/v2"
)

type Commands struct {
	client *switchbot.Client
}

func NewCommands(client *switchbot.Client) *Commands {
	return &Commands{
		client: client,
	}
}

func (c *Commands) ListDevices(ctx context.Context) ([]switchbot.Device, error) {
	pdev, _, err := c.client.Device().List(ctx)
	if err != nil {
		return nil, err
	}
	return pdev, nil
}

func (c *Commands) Status(ctx context.Context, id string) (*switchbot.DeviceStatus, error) {
	status, err := c.client.Device().Status(ctx, id)
	if err != nil {
		return nil, err
	}
	return &status, nil
}
