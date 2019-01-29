package dexcom

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type getCalibrationsResponse struct {
	Calibrations     []Calibration  `json:"calibrations"`
}

type Calibration struct {
	SystemTime    Time `json:"systemTime"`
	DisplayTime   Time `json:"displayTime"`
	Value         int       `json:"value"`
	Unit string `json:"unit"`
}

// GetCalibrations will return the calibrations made in the given range
func (c *Client) GetCalibrations(start, end time.Time) ([]Calibration, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/v2/users/self/calibrations", c.Endpoint))
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

	var calibrations getCalibrationsResponse
	err = json.NewDecoder(resp.Body).Decode(&calibrations)
	if err != nil {
		c.Logger.Log("level", "error", "msg", "could not json decode response", "err", err.Error())
		return nil, err
	}

	c.Logger.Log("level", "info", "msg", "received EGVs", "n", len(calibrations.Calibrations))

	return calibrations.Calibrations, nil
}
