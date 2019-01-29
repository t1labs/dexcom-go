package dexcom_test

import (
	"os"
	"testing"
)

func setupTest(t *testing.T) string {
	accessToken := os.Getenv("TEST_ACCESS_TOKEN")
	if accessToken == "" {
		t.Skip("TEST_ACCESS_TOKEN not set")
	}

	return accessToken
}