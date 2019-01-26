package pvars

import (
	"testing"
)

func TestConstants(t *testing.T) {
	v := TimestampFormat
	if v == "" {
		t.Error("TimestampFormat is not set")
	}

	if v == "null" {
		t.Error("TimestampFormat has wrong value")
	}

	v = PackageName
	if v != "pvars" {
		t.Error("Package name is wrong:", v)
	}

}

