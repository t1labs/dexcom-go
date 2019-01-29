package dexcom

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type getDevicesRequest struct {
	Devices []Device `json:"devices"`
}

type Device struct {
	TransmitterGeneration string          `json:"transmitterGeneration"`
	DisplayDevice         string          `json:"displayDevice"`
	LastUploadDate        Time            `json:"lastUploadDate"`
	AlertSchedules        []AlertSchedule `json:"alertScheduleList"`
}

type AlertSchedule struct {
	Settings AlertScheduleSettings `json:"alertScheduleSettings"`
	Alerts   []Alert
}

type AlertScheduleSettings struct {
	Name              string   `json:"alertScheduleName"`
	IsEnabled         bool     `json:"isEnabled"`
	IsDefaultSchedule bool     `json:"isDefaultSchedule"`
	StartTime         string   `json:"startTime"`
	EndTime           string   `json:"endTime"`
	DaysOfWeek        []string `json:"daysOfWeek"`
}

type Alert struct {
	Name        string `json:"alertName"`
	Value       int    `json:"value"`
	Unit        string `json:"unit"`
	Snooze      int    `json:"snooze"`
	SystemTime  Time   `json:"systemTime"`
	DisplayTime Time   `json:"displayTime"`
	Enabled     bool   `json:"enabled"`
}

// GetCalibrations will return the calibrations made in the given range
func (c *Client) GetDevices(start, end time.Time) ([]Device, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/v2/users/self/devices", c.Endpoint))
	if err != nil {
		return nil, err
	}

	vals := uri.Query()
	vals.Add("startDate", start.Format("2006-01-02T15:04:05"))
	vals.Add("endDate", end.Format("2006-01-02T15:04:05"))
	uri.RawQuery = vals.Encode()

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		c.Logger.Log("level", "error", "msg", "could make new request", "err", err.Error())
		return nil, err
	}

	resp, err := c.c.Do(req)
	if err != nil {
		c.Logger.Log("level", "error", "msg", "could not do request", "err", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	var devices getDevicesRequest
	err = json.NewDecoder(resp.Body).Decode(&devices)
	if err != nil {
		c.Logger.Log("level", "error", "msg", "could not json decode response", "err", err.Error())
		return nil, err
	}

	c.Logger.Log("level", "info", "msg", "received devices", "n", len(devices.Devices))

	return devices.Devices, nil
}
