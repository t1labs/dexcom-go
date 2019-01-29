package dexcom

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type getEGVsResponse struct {
	Unit     string `json:"unit"`
	rateUnit string `json:"rateUnit"`
	EGVs     []EGV  `json:"egvs"`
}

type EGV struct {
	SystemTime    Time    `json:"systemTime"`
	DisplayTime   Time    `json:"displayTime"`
	Value         int     `json:"value"`
	RealtimeValue int     `json:"realtimeValue"`
	SmoothedValue int     `json:"smoothedValue"`
	Status        string  `json:"status"`
	Trend         string  `json:"trend"`
	TrendRate     float64 `json:"trendRate"`
}

// GetEGVs will return all of the EGVs in the given time range or an error describing the failed network call
func (c *Client) GetEGVs(start, end time.Time) ([]EGV, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/v2/users/self/egvs", c.Endpoint))
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

	var egvs getEGVsResponse
	err = json.NewDecoder(resp.Body).Decode(&egvs)
	if err != nil {
		c.Logger.Log("level", "error", "msg", "could not json decode response", "err", err.Error())
		return nil, err
	}

	c.Logger.Log("level", "info", "msg", "received EGVs", "n", len(egvs.EGVs))

	return egvs.EGVs, nil
}
