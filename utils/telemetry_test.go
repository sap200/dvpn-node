package utils

import "testing"

func TestGetBandwidth(t *testing.T) {
	x := GetBandwidth()

	if x != "" {
		t.Logf("Expected %s, got %s", x, x)
	} else {
		t.Fatalf("Expected format (1,2,3), got %s", x)
	}
}
