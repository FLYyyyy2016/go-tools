package api

import "testing"

func TestGetMyExtranetIP(t *testing.T) {
	ipStatus, err := GetMyExtranetIP("")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", ipStatus)
}

func TestGetOtherExtranetIP(t *testing.T) {
	ipStatus, err := GetMyExtranetIP("114.114.114.114")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", ipStatus)
}
