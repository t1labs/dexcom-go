package dexcom_test

import (
	"testing"
	"time"

	"github.com/t1labs/dexcom-go"
)

func TestGetDevices(t *testing.T) {
	accessToken := setupTest(t)

	c := dexcom.New(accessToken)
	dec2018, _ := time.Parse(time.RFC3339, "2018-12-01T00:00:00Z")
	jan2019, _ := time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")
	egvs, err := c.GetDevices(dec2018, jan2019)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(egvs) < 1 {
		t.Fatal("expected at least one device to be returned")
	}
}