// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadbalancingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

const (
	maxPositions     uint32 = 36000 // 360 degrees with two decimal places
	defaultWeight    int    = 100   // the number of points in the ring for each entry. For better results, it should be greater than 100.
	linearProbeLimit int    = 10    // The number of times to probe ahead in the hash ring if there is a collision while constructing the hash ring
)

// position represents a specific angle in the ring.
// Each entry in the ring is positioned at an angle in a hypothetical circle, meaning that it ranges from 0 to 360.
type position uint32

// ringItem connects a specific angle in the ring with a specific endpoint.
type ringItem struct {
	pos      position
	endpoint string
}

// hashRing is a consistent hash ring following Karger et al.
type hashRing struct {
	// ringItems holds all the positions, used for the lookup the position for the closest next ring item
	items []ringItem
}

// newHashRing builds a new immutable consistent hash ring based on the given endpoints.
func newHashRing(endpoints []string) *hashRing { _ = "STUB: not implemented"; return nil }

// endpointFor calculates which backend is responsible for the given traceID
func (h *hashRing) endpointFor(identifier []byte) string {
	_ = "STUB: not implemented"

	// perhaps the ring itself couldn't get initialized yet?
	return ""
}

// findEndpoint returns the "next" endpoint starting from the given position, or an empty string in case no endpoints are available
func (h *hashRing) findEndpoint(pos position) string { _ = "STUB: not implemented"; return "" }

// bsearch is a binary search-like algorithm, returning the closest "next" item instead of an exact match
func bsearch(pos position, left, right []ringItem) ringItem {
	_ = "STUB: not implemented"
	// if it's the last item of the left side, return it
	return *new(ringItem)
}

// if it's the first item of the right side, return it

// if we want a higher angle than the highest from the ring, the first angle is the right one

// if the requested position is greater than the highest in the left, the item is in the right side

// not on the right side, has to be on the left side

// positionFor calculates all the positions in the ring based. The numPoints indicates how many positions to calculate.
// The slice length of the result matches the numPoints.
func positionsFor(endpoint string, numPoints int) []position { _ = "STUB: not implemented"; return nil }

// positionsForEndpoints calculates all the positions for all the given endpoints
func positionsForEndpoints(endpoints []string, weight int) []ringItem {
	_ = "STUB: not implemented"
	return nil
}

// tracking the used positions

// for this initial implementation, we don't allow endpoints to have custom weights

// if this position is occupied already, look ahead in the array for a free position

// Not able to find a free spot; skip this item

func (h *hashRing) equal(candidate *hashRing) bool { _ = "STUB: not implemented"; return false }
