// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver"

import (
	"context"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/monitor/query/azmetrics"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources/v3"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver/internal/metadata"
)

type azureType struct {
	name                      *string
	resourceIDs               []string
	metricsByCompositeKey     map[metricsCompositeKey]*azureResourceMetrics
	metricsDefinitionsUpdated time.Time
}

func newBatchScraper(conf *Config, settings receiver.Settings) *azureBatchScraper {
	_ = "STUB: not implemented"
	return nil
}

type azureBatchScraper struct {
	cred             azcore.TokenCredential
	cfg              *Config
	receiverSettings receiver.Settings
	settings         component.TelemetrySettings
	// resources on which we'll get attributes. Stored by resource id and subscription id.
	resources map[string]map[string]*azureResource
	// resourceTypes on which we'll collect metrics. Stored by resource type and subscription id.
	resourceTypes map[string]map[string]*azureType
	// subscriptions on which we'll look up resources. Stored by subscription id.
	subscriptions        map[string]*azureSubscription
	subscriptionsUpdated time.Time
	// regions on which we'll collect metrics. Stored by subscription id.
	regions map[string]map[string]struct{}
	mbs     concurrentMetricsBuilderMap[*metadata.MetricsBuilder]

	mutex                        *sync.Mutex
	time                         timeNowIface
	clientOptionsResolver        ClientOptionsResolver
	storageAccountSpecificConfig storageAccountSpecificConfig
}

func (s *azureBatchScraper) GetMetricsBatchValuesClient(region string) (*azmetrics.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *azureBatchScraper) start(_ context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *azureBatchScraper) loadSubscription(sub azureSubscription) {
	_ = "STUB: not implemented"
	return
}

func (s *azureBatchScraper) unloadSubscription(id string) { _ = "STUB: not implemented"; return }

func (s *azureBatchScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Once all metrics has been collected for one subscription, we save them in the associated metrics builder.
// Having a map of metrics builders, one per subscription, allows us to collect each subscription concurrently.
// We'll be able to emit them all at once at the end of the scrape, once all subscriptions have been processed.

// TODO: duplicate
func (s *azureBatchScraper) loadSubscriptions(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// Subscriptions discovery enabled or not, we'll need a client.
// - If enabled, to get the subscription list
// - If not, to get more info about the subscription
// The only case where it won't be needed is when we don't want the subscription name in resource attributes.

// Make a special case for when we only have subscription ids configured (discovery disabled)

// we don't need additional info,
// => It simply load the subscription id

// We need additional info,
// => It makes some get requests

// Prepare a map of existing subscriptions to detect removed ones later

// Unload subscriptions that are no longer present

// TODO: partially duplicate
func (s *azureBatchScraper) loadResourcesAndTypes(ctx context.Context, subscriptionID string) {
	_ = "STUB: not implemented"
	return
}

// Prepare a map of existing resources to detect removed ones later

// Prepare the tags filter to apply on resources later on.
// TODO: We don't need to do it per subscription. It can be done upper in the code to improve performances.

// Prepare the options to get the resource list.

// Unload resources that are no longer present

// processResources is a workaround specially done for the storageAccount metrics.
// Every StorageAccount resources have some implicit sub resources (/blobServices/default, fileServices/default, etc...) that are not returned by the API.
// We need to add them manually to the resources and resourceTypes map.
// Note that we add these virtual sub resources only if the user has asked the sub resource types explicitly in the services config.
// Example:
// For each resource with id .../Microsoft.Storage/storageAccount/myResource of type Microsoft.Storage/storageAccount,
// It will create a virtual resource with id .../Microsoft.Storage/storageAccount/myResource/blobServices/default of type Microsoft.Storage/storageAccounts/blobServices.
// TODO: duplicate
func (s *azureBatchScraper) processResources(resources []*armresources.GenericResourceExpanded) []*armresources.GenericResourceExpanded {
	_ = "STUB: not implemented"
	return nil
}

// TODO: duplicate
func (s *azureBatchScraper) getResourcesFilter() string {
	_ = "STUB: not implemented"
	// TODO: switch to parsing services from
	// https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/metrics-supported
	return ""
}

// TODO: Partially duplicate
func (s *azureBatchScraper) loadResourceMetricsDefinitionsByType(ctx context.Context, subscriptionID, resourceType string) {
	_ = "STUB: not implemented"
	return
}

// Prepare the map of metrics by composite key.

// TODO: duplicate
func (s *azureBatchScraper) loadMetricsDefinitionByType(subscriptionID, resourceType, metricName string, compositeKey metricsCompositeKey) {
	_ = "STUB: not implemented"
	return
}

func (s *azureBatchScraper) loadBatchMetricsValues(ctx context.Context, subscriptionID, resourceType string) {
	_ = "STUB: not implemented"
	return
}

// Anti-duplicate guard: do not re-scrape a (resourceType, compositeKey) pair more often
// than its Azure timeGrain. The batch scraper emits points using the original Azure
// timestamp (cf. processQueryTimeseriesData), so re-scraping the same window would
// republish identical (labels, timestamp) tuples and trigger 409 "duplicate sample
// for timestamp" errors on Prometheus-compatible backends (Thanos, Mimir, Cortex…).
// This mirrors the guard already present in the non-batch scraper (scraper.go).
// Note: the guard is best-effort. metricsByCompositeKey is reset whenever the
// metrics definitions cache (CacheResourcesDefinitions, default 24h) expires,
// so at most one duplicating scrape may happen right after each expiration and on restarts.

// times 4 because for some resources, data are missing for the very latest timestamp. The processing will keep only the latest timestamp with data.

// reverse for loop because newest timestamp is at the end of the slice

// newQueryResourcesOptions builds the options to make the QueryResources request.
func newQueryResourcesOptions(
	dimensionsStr string,
	timeGrain string,
	aggregationsStr string,
	start time.Time,
	end time.Time,
	top int32,
) azmetrics.QueryResourcesOptions {
	_ = "STUB: not implemented"
	return *new(azmetrics.QueryResourcesOptions)
}

// Defaults to 10 (may be limiting results)

// metricValueIsNotEmpty checks if the metric value is empty.
// This is necessary to compensate for the fact that Azure Monitor sometimes returns empty values.
func metricValueIsNotEmpty(metricValue azmetrics.MetricValue) bool {
	_ = "STUB: not implemented"
	// Using an "or" chain is a bet on performance improvement. Assuming that it's not checking others if one is not nil. Not strictly verified though.
	return false
}

func (s *azureBatchScraper) processQueryTimeseriesData(
	mb *metadata.MetricsBuilder,
	resourceID string,
	metric azmetrics.Metric,
	metricValue azmetrics.MetricValue,
	attributes map[string]*string,
) {
	_ = "STUB: not implemented"
	return
}
