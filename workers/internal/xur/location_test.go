package xur

import "testing"

func TestLocation(t *testing.T) {
	location, err := Location()
	if err != nil {
		t.Fatal(err)
	}

	if !location.Found {
		t.Logf("Xur was not available")
	} else {
		t.Logf("%#v", location)
	}
}
