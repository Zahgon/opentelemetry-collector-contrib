// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver"

import (
	"context"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources/v3"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver/internal/metadata"
)

var (
	timeGrains = map[string]int64{
		"PT1M":  60,
		"PT5M":  300,
		"PT15M": 900,
		"PT30M": 1800,
		"PT1H":  3600,
		"PT6H":  21600,
		"PT12H": 43200,
		"P1D":   86400,
	}
	aggregations = []string{
		"Average",
		"Count",
		"Maximum",
		"Minimum",
		"Total",
	}
)

const (
	attributeLocation      = "location"
	attributeName          = "name"
	attributeResourceGroup = "resource_group"
	attributeResourceType  = "type"
	metadataPrefix         = "metadata_"
	tagPrefix              = "tags_"
	truncateTimeGrain      = time.Minute
	filterAllAggregations  = "*"
	storageAccountType     = "Microsoft.Storage/storageAccounts"
)

// azureSubscription is an extract of armsubscriptions.Subscription.
// It designates a common structure between complex structures retrieved from the AP
// and simple subscriptions ids that you can find in config.
type azureSubscription struct {
	SubscriptionID   string
	DisplayName      string
	resourcesUpdated time.Time
}

type azureResource struct {
	attributes                map[string]*string
	metricsByCompositeKey     map[metricsCompositeKey]*azureResourceMetrics
	metricsDefinitionsUpdated time.Time
	tags                      map[string]*string
	resourceType              *string
}

type metricsCompositeKey struct {
	dimensions   string // comma separated sorted dimensions
	aggregations string // comma separated sorted aggregations
	timeGrain    string
}

type azureResourceMetrics struct {
	metrics              []string
	metricsValuesUpdated time.Time
}

type void struct{}

type timeNowIface interface {
	Now() time.Time
}

type timeWrapper struct{}

func (*timeWrapper) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

type storageAccountSpecificConfig struct {
	askedBlobServices  bool
	askedFileServices  bool
	askedQueueServices bool
	askedTableServices bool
}

func newStorageAccountSpecificConfig(services []string) storageAccountSpecificConfig {
	_ = "STUB: not implemented"
	return *new(storageAccountSpecificConfig)
}

func newScraper(conf *Config, settings receiver.Settings) *azureScraper {
	_ = "STUB: not implemented"
	return nil
}

type azureScraper struct {
	cred azcore.TokenCredential

	cfg      *Config
	settings component.TelemetrySettings
	// resources on which we'll collect metrics. Stored by resource id and subscription id.
	resources map[string]map[string]*azureResource
	// subscriptions on which we'll look up resources. Stored by subscription id.
	subscriptions        map[string]*azureSubscription
	subscriptionsUpdated time.Time
	mb                   *metadata.MetricsBuilder

	mutex                        *sync.Mutex
	time                         timeNowIface
	clientOptionsResolver        ClientOptionsResolver
	storageAccountSpecificConfig storageAccountSpecificConfig
}

func (s *azureScraper) start(_ context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *azureScraper) loadSubscription(sub azureSubscription) { _ = "STUB: not implemented"; return }

func (s *azureScraper) unloadSubscription(id string) { _ = "STUB: not implemented"; return }

func (s *azureScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Once all metrics has been collected for one subscription, we move to the next.
// We need to keep it synchronous to have the subscription id in resource attributes and not metrics attributes.
// It can be revamped later if we need to parallelize more, but currently, resource emit is not thread safe.

func (s *azureScraper) loadSubscriptions(ctx context.Context) { _ = "STUB: not implemented"; return }

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

func (s *azureScraper) loadResources(ctx context.Context, subscriptionID string) {
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
// Note that we do that hack only if the user has asked the sub resource types explicitly in the services config.
// Example:
// For each resource with id .../Microsoft.Storage/storageAccount/myResource of type Microsoft.Storage/storageAccount,
// It will create a fake resource with id .../Microsoft.Storage/storageAccount/myResource/blobServices/default of type Microsoft.Storage/storageAccounts/blobServices.
// TODO: duplicate
func (s *azureScraper) processResources(resources []*armresources.GenericResourceExpanded) []*armresources.GenericResourceExpanded {
	_ = "STUB: not implemented"
	return nil
}

// buildSubTypeResource creates a virtual new resource with given type and ID.
// The rest of the attributes (location, tags, etc...) are copied from the original resource.
func buildSubTypeResource(orig armresources.GenericResourceExpanded, newType, newID string) *armresources.GenericResourceExpanded {
	_ = "STUB: not implemented"
	return nil
}

func getResourceGroupFromID(id string) string { _ = "STUB: not implemented"; return "" }

func (s *azureScraper) getResourcesFilter() string {
	_ = "STUB: not implemented"
	// TODO: switch to parsing services from
	// https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/metrics-supported
	return ""
}

func (s *azureScraper) loadMetricsDefinitions(ctx context.Context, subscriptionID, resourceID string) {
	_ = "STUB: not implemented"
	return
}

// Prepare the map of metrics by composite key.

func (s *azureScraper) loadMetricsDefinition(subscriptionID, resourceID, metricName string, compositeKey metricsCompositeKey) {
	_ = "STUB: not implemented"
	return
}

func (s *azureScraper) loadMetricsValues(ctx context.Context, subscriptionID, resourceID string) {
	_ = "STUB: not implemented"
	return
}

func newResourceMetricsValuesRequestOptions(
	metrics []string,
	dimensionsStr string,
	timeGrain string,
	aggregationsStr string,
	start int,
	end int,
	top int32,
) armmonitor.MetricsClientListOptions {
	_ = "STUB: not implemented"
	return *new(armmonitor.MetricsClientListOptions)
}

func (s *azureScraper) processTimeseriesData(
	resourceID string,
	metric *armmonitor.Metric,
	metricValue *armmonitor.MetricValue,
	attributes map[string]*string,
) {
	_ = "STUB: not implemented"
	return
}

// getMetricAggregations returns a list of aggregations for a given namespace/metric.
// Two parameters are considered to know the aggregation to choose
// - a filter (given in configuration)
// - a list of supported aggregations (given by the API)
// If one namespace/metrics combination matches a provided filter,
// > Then it returns the aggregations in the filter
// > Otherwise it returns all supported aggregations.
// Note that a special filter * is supported to return all supported aggregations explicitly.
// /!\ It does not control the aggregations in the filters. If it's not in the supported list, it still lets it pass.
func getMetricAggregations(metricNamespace, metricName string, filters NestedListAlias, supportedAggregations []string) []string {
	_ = "STUB: not implemented"
	// default behavior when no metric filters specified: pass all metrics with all aggregations
	return nil
}

// metric namespace isn't found, or it's empty: pass all metrics from the namespace

// if the target metric is absent in the metrics map: filter out metric

// allow all aggregations if others are not specified

// collect known aggregations without filtering on supported

func convertAggregationsToStr(aggregations []*armmonitor.AggregationType) []string {
	_ = "STUB: not implemented"
	return nil
}

func mapFindInsensitive[T any](m map[string]T, key string) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// getTagsFilterMap returns a map used to filter tags.
// Each user-configured tag key is normalized to lowercase and added to the map for case-insensitive lookup.
func getTagsFilterMap(appendTagsAsAttributes []string) (tagsFilterMap map[string]struct{}) {
	_ = "STUB: not implemented"
	return nil
}

// filterResourceTags filter out resource tags according to configured tag list (append_tags_as_attributes)
func filterResourceTags(tagFilterList map[string]struct{}, resourceTags map[string]*string) map[string]*string {
	_ = "STUB: not implemented"
	return nil
}

// wildcard not found. include only configured tags
