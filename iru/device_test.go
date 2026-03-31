package iru

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func newTestClient(t *testing.T) *Client {
	t.Helper()
	_ = godotenv.Load("../.env")
	apiKey := os.Getenv("IRU_API_KEY")
	subdomain := os.Getenv("IRU_SUBDOMAIN")
	if apiKey == "" || subdomain == "" {
		t.Skip("skipping integration test: IRU_API_KEY and IRU_SUBDOMAIN must be set")
	}
	return NewClient(Config{
		APIKey:    apiKey,
		Subdomain: subdomain,
	})
}

func TestListDevices(t *testing.T) {
	c := newTestClient(t)
	ctx := context.Background()

	devices, err := c.ListDevices(ctx)
	if err != nil {
		t.Fatalf("ListDevices: %v", err)
	}
	t.Logf("got %d devices", len(devices))
}

func TestGetDevice(t *testing.T) {
	c := newTestClient(t)
	ctx := context.Background()

	deviceID := os.Getenv("IRU_DEVICE_ID")
	if deviceID == "" {
		devices, err := c.ListDevices(ctx)
		if err != nil {
			t.Fatalf("ListDevices: %v", err)
		}
		if len(devices) == 0 {
			t.Skip("skipping: no devices available")
		}
		deviceID = devices[0].DeviceID
	}

	device, err := c.GetDevice(ctx, deviceID)
	if err != nil {
		t.Fatalf("GetDevice: %v", err)
	}
	if device.DeviceID != deviceID {
		t.Errorf("GetDevice: got device_id %q, want %q", device.DeviceID, deviceID)
	}
	t.Logf("got device: %s (%s)", device.DeviceName, device.SerialNumber)
}

func TestGetDeviceDetails(t *testing.T) {
	c := newTestClient(t)
	ctx := context.Background()

	deviceID := os.Getenv("IRU_DEVICE_ID")
	if deviceID == "" {
		devices, err := c.ListDevices(ctx)
		if err != nil {
			t.Fatalf("ListDevices: %v", err)
		}
		if len(devices) == 0 {
			t.Skip("skipping: no devices available")
		}
		deviceID = devices[0].DeviceID
	}

	details, err := c.GetDeviceDetails(ctx, deviceID)
	if err != nil {
		t.Fatalf("GetDeviceDetails: %v", err)
	}
	if details.General.DeviceID != deviceID {
		t.Errorf("GetDeviceDetails: got device_id %q, want %q", details.General.DeviceID, deviceID)
	}
	t.Logf("got details for: %s (%s)", details.General.DeviceName, details.HardwareOverview.SerialNumber)
}
