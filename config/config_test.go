package config

import (
	"testing"

	"encoding/json"

	"fmt"

	"os"

	"github.com/xiaonanln/goworld/gwlog"
)

func TestLoad(t *testing.T) {
	config := Get()
	gwlog.Debug("goworld config: \n%s", config)
	if config == nil {
		t.FailNow()
	}
	if config.Dispatcher.Ip == "" {
		t.Errorf("dispatch ip not found")
	}
	if config.Dispatcher.Port == 0 {
		t.Errorf("dispatcher port not found")
	}
	for serverName, serverConfig := range config.Servers {
		if serverConfig.Ip == "" {
			t.Errorf("server %s ip not found", serverName)
		}
		if serverConfig.Port == 0 {
			t.Errorf("server %s port not found", serverName)
		}
	}

	gwlog.Info("read goworld config: %v", config)
}

func TestReload(t *testing.T) {
	config := Get()
	config = Reload()
	gwlog.Debug("goworld config: \n%s", config)
}

func TestGetDispatcher(t *testing.T) {
	cfg := GetDispatcher()
	cfgStr, _ := json.Marshal(cfg)
	fmt.Printf("dispatcher config: %s", string(cfgStr))
}

func TestGetServer(t *testing.T) {
	for id := 1; id <= 10; id++ {
		cfg := GetServer(id)
		if cfg == nil {
			gwlog.Info("Server %d not found", id)
		} else {
			gwlog.Info("Server %d config: %v", id, cfg)
		}
	}
}

func TestGetStorage(t *testing.T) {
	cfg := GetStorage()
	if cfg == nil {
		t.Errorf("storage config not found")
	}
	gwlog.Info("storage config:")
	fmt.Fprintf(os.Stderr, "%s\n", DumpPretty(cfg))
}
