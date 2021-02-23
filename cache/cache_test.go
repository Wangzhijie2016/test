package cache

import (
	"testing"
)

func TestCache(t *testing.T) {
	if Cache() != "v2.0.12: Cache()" {
		t.Error("Cache() error")
	}
}
