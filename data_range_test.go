package dexcom_test

import (
	"github.com/t1labs/dexcom-go"
	"testing"
)

func TestGetDataRange(t *testing.T) {
	accessToken := setupTest(t)

	c := dexcom.New(accessToken)
	cr, egvR, eventR, err := c.GetDataRange()
	if err != nil {
		t.Fatal(err.Error())
	}

	if cr.Start.SystemTime.IsZero() {
		t.Fatal("expected calibration range to have start date")
	}
	if cr.End.SystemTime.IsZero() {
		t.Fatal("expected calibration range to have start date")
	}

	if egvR.Start.SystemTime.IsZero() {
		t.Fatal("expected egv range to have start date")
	}
	if egvR.End.SystemTime.IsZero() {
		t.Fatal("expected egv range to have start date")
	}

	if eventR.Start.SystemTime.IsZero() {
		t.Fatal("expected event range to have start date")
	}
	if eventR.End.SystemTime.IsZero() {
		t.Fatal("expected event range to have start date")
	}
}