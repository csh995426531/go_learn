package remote_package_test

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestRemotePackage(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
