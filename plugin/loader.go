package plugin

import (
	"fmt"
	"go-twing/core/config"
	"plugin"
	"runtime"
)

var preloadedPlugins []plugin.Plugin

func Preload(plugins ...plugin.Plugin) {
	preloadedPlugins = append(preloadedPlugins, plugins...)
}

var loadPluginFunc = func(string) ([]plugin.Plugin, error) {
	return nil, fmt.Errorf("unsupported %s", runtime.GOOS)
}

type loaderState int

const (
	loaderLoading loaderState = iota
	loaderInitializing
	loaderInitialized
	loaderInjecting
	loaderInjected
	loaderStarting
	loaderStarted
	loaderClosing
	loaderClosed
	loaderFailed
)

func (ls loaderState) String() string {
	switch ls {
	case loaderLoading:
		return "Loading"
	case loaderInitializing:
		return "Initializing"
	case loaderInitialized:
		return "Initialized"
	case loaderInjecting:
		return "Injecting"
	case loaderInjected:
		return "Injected"
	case loaderStarting:
		return "Starting"
	case loaderStarted:
		return "Started"
	case loaderClosing:
		return "Closing"
	case loaderClosed:
		return "Closed"
	case loaderFailed:
		return "Failed"
	default:
		return "Unknown"
	}
}

type PluginLoader struct {
	state   loaderState
	plugins map[string]plugin.Plugin
	started []plugin.Plugin
	config  config.Plugins
	repo    string
}

func NewPluginLoader(repo string) (*PluginLoader, error) {
	return nil, nil
}
