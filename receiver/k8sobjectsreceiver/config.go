// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sobjectsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sobjectsreceiver"

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/filter"
	"k8s.io/apimachinery/pkg/runtime/schema"
	apiWatch "k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sinventory"
)

const (
	defaultPullInterval    time.Duration     = time.Hour
	defaultMode            k8sinventory.Mode = k8sinventory.PullMode
	defaultResourceVersion                   = "1"
)

var modeMap = map[k8sinventory.Mode]bool{
	k8sinventory.PullMode:  true,
	k8sinventory.WatchMode: true,
}

type ErrorMode string

const (
	PropagateError ErrorMode = "propagate"
	IgnoreError    ErrorMode = "ignore"
	SilentError    ErrorMode = "silent"
)

type K8sObjectsConfig struct {
	Name              string               `mapstructure:"name"`
	Group             string               `mapstructure:"group"`
	Namespaces        []string             `mapstructure:"namespaces"`
	ExcludeNamespaces []filter.Config      `mapstructure:"exclude_namespaces"`
	Mode              k8sinventory.Mode    `mapstructure:"mode"`
	LabelSelector     string               `mapstructure:"label_selector"`
	FieldSelector     string               `mapstructure:"field_selector"`
	Interval          time.Duration        `mapstructure:"interval"`
	ResourceVersion   string               `mapstructure:"resource_version"`
	ExcludeWatchType  []apiWatch.EventType `mapstructure:"exclude_watch_type"`
	exclude           map[apiWatch.EventType]bool
	gvr               *schema.GroupVersionResource
}

type Config struct {
	k8sconfig.APIConfig `mapstructure:",squash"`

	Objects             []*K8sObjectsConfig `mapstructure:"objects"`
	Storage             *component.ID       `mapstructure:"storage"`
	ErrorMode           ErrorMode           `mapstructure:"error_mode"`
	IncludeInitialState bool                `mapstructure:"include_initial_state"`

	K8sLeaderElector *component.ID `mapstructure:"k8s_leader_elector"`

	// For mocking purposes only.
	makeDiscoveryClient func() (discovery.ServerResourcesInterface, error)
	makeDynamicClient   func() (dynamic.Interface, error)
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *Config) getDiscoveryClient() (discovery.ServerResourcesInterface, error) {
	_ = "STUB: not implemented"
	return *new(discovery.ServerResourcesInterface), nil
}

func (c *Config) getDynamicClient() (dynamic.Interface, error) {
	_ = "STUB: not implemented"
	return *new(dynamic.Interface), nil
}

func (c *Config) getValidObjects() (map[string][]*schema.GroupVersionResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if Partial result is returned from discovery client, that means some API servers have issues,
// but we can still continue, as we check for the needed groups later in Validate function.

func (k *K8sObjectsConfig) DeepCopy() *K8sObjectsConfig { _ = "STUB: not implemented"; return nil }
