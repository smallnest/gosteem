package api

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	resp, err := GetConfig(client)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetConfig, resp)
}

func TestGetDynamicGlobalProperties(t *testing.T) {
	resp, err := GetDynamicGlobalProperties(client)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetDynamicGlobalProperties, resp)
}

func TestGetWitnessSchedule(t *testing.T) {
	resp, err := GetWitnessSchedule(client)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetWitnessSchedule, resp)
}

func TestGetFeedHistory(t *testing.T) {
	resp, err := GetFeedHistory(client)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetFeedHistory, resp)
}

func TestGetActiveWitnesses(t *testing.T) {
	resp, err := GetActiveWitnesses(client)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %+v", MethodGetFeedHistory, resp)
}
