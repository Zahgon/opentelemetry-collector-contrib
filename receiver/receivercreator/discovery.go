// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package receivercreator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator"

import (
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

const (
	// hints prefix
	otelHints = "io.opentelemetry.discovery"

	// hint suffix for metrics
	otelMetricsHints = otelHints + ".metrics"
	otelLogsHints    = otelHints + ".logs"

	// hints definitions
	discoveryEnabledHint = "enabled"
	scraperHint          = "scraper"
	configHint           = "config"

	logsReceiver          = "file_log"
	defaultLogPathPattern = "/var/log/pods/%s_%s_%s/%s/*.log"
)

// k8sHintsBuilder creates configurations from hints provided as Pod's annotations.
type k8sHintsBuilder struct {
	logger               *zap.Logger
	ignoreReceivers      map[string]bool
	defaultAnnotations   map[string]string
	defaultFileLogConfig userConfigMap
}

func createK8sHintsBuilder(config DiscoveryConfig, logger *zap.Logger) k8sHintsBuilder {
	_ = "STUB: not implemented"
	return *new(k8sHintsBuilder)
}

// createReceiverTemplateFromHints creates a receiver configuration based on the provided hints.
// Hints are extracted from Pod's annotations.
// Scraper configurations are only created for Port Endpoints.
// Log receiver configurations are only created for Pod Container Endpoints.
func (builder *k8sHintsBuilder) createReceiverTemplateFromHints(env observer.EndpointEnv) (*receiverTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (builder *k8sHintsBuilder) createScraper(
	annotations map[string]string,
	env observer.EndpointEnv,
) (*receiverTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no scraper hint detected

// scraper is ignored

func (builder *k8sHintsBuilder) createLogsReceiver(
	annotations map[string]string,
	env observer.EndpointEnv,
) (*receiverTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// receiver is ignored

func getScraperConfFromAnnotations(
	annotations map[string]string,
	defaultEndpoint, scopeSuffix string,
	logger *zap.Logger,
) (userConfigMap, error) {
	_ = "STUB: not implemented"
	return *new(userConfigMap), nil
}

// defaultEndpoint will be added properly later in observerHandler.startReceiver method

// skip endpoint's validation if there is no user provided endpoint
// defaultEndpoint will be added properly later in observerHandler.startReceiver method

func createLogsConfig(
	annotations map[string]string,
	containerName, podUID, podName, namespace string,
	defaultConfMap userConfigMap,
	logger *zap.Logger,
) userConfigMap {
	_ = "STUB: not implemented"
	return *new(userConfigMap)
}

// path cannot be other than the one of the target container

func getHintAnnotation(annotations map[string]string, hintBase, hintKey, suffix string) (string, bool) {
	_ = "STUB: not implemented"
	// try to scope the hint more on container level by suffixing
	// with .<port> in case of Port event or .<container_name> in case of Pod Container event
	return "", false
}

// if there is no container level hint defined try to use the Pod level hint

func discoveryEnabled(annotations map[string]string, hintBase, scopeSuffix string) bool {
	_ = "STUB: not implemented"
	return false
}

func getStringEnv(env observer.EndpointEnv, key string) string {
	_ = "STUB: not implemented"
	return ""
}

func validateEndpoint(endpoint, defaultEndpoint string) error {
	_ = "STUB: not implemented"
	// replace temporarily the dynamic reference to ease the url parsing
	return nil
}

// target endpoint can come in form ip:port. In that case we fix the uri
// temporarily with adding http scheme

// configured endpoint should include the target Pod's endpoint

func mergeAnnotations(podAnnotations, defaultAnnotations map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Start with defaultAnnotations (lower priority)

// Overwrite with podAnnotations (higher priority)
