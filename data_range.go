package dexcom

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type getDataRangeResponse struct {
	CalibrationRange CalibrationRange `json:"calibrations"`
	EGVRange         EGVRange         `json:"egvs"`
	EventRange       EventRange       `json:"events"`
}

type Range struct {
	DisplayTime Time `json:"displayTime"`
	SystemTime  Time `json:"systemTime"`
}

type CalibrationRange struct {
	End   Range `json:"end"`
	Start Range `json:"start"`
}

type EGVRange struct {
	End   Range `json:"end"`
	Start Range `json:"start"`
}

type EventRange struct {
	End   Range `json:"end"`
	Start Range `json:"start"`
}

// GetCalibrations will return the calibrations made in the given range
func (c *Client) GetDataRange() (CalibrationRange, EGVRange, EventRange, error) {
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
