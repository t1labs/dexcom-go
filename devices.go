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
	uri, err := url.Parse(fmt.Sprintf("%s/v2/users/self/dataRange", c.Endpoint))
	if err != nil {
		return CalibrationRange{}, EGVRange{}, EventRange{}, err
	}

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		c.Logger.Log("level", "error", "msg", "could make new request", "err", err.Error())
		return CalibrationRange{}, EGVRange{}, EventRange{}, err
	}

	resp, err := c.c.Do(req)
	if err != nil {
		c.Logger.Log("level", "error", "msg", "could not do request", "err", err.Error())
		return CalibrationRange{}, EGVRange{}, EventRange{}, err
	}
	defer resp.Body.Close()

	var dataRange getDataRangeResponse
	err = json.NewDecoder(resp.Body).Decode(&dataRange)
	if err != nil {
		c.Logger.Log("level", "error", "msg", "could not json decode response", "err", err.Error())
		return CalibrationRange{}, EGVRange{}, EventRange{}, err
	}

	c.Logger.Log("level", "info", "msg", "received data ranges")

	return dataRange.CalibrationRange, dataRange.EGVRange, dataRange.EventRange, nil
}
