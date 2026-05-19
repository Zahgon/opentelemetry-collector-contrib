// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package activedirectorydsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/activedirectorydsreceiver"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters"
)

const (
	draInboundBytesCompressed              = "DRA Inbound Bytes Compressed (Between Sites, After Compression) Since Boot"
	draInboundBytesNotCompressed           = "DRA Inbound Bytes Not Compressed (Within Site) Since Boot"
	draOutboundBytesCompressed             = "DRA Outbound Bytes Compressed (Between Sites, After Compression) Since Boot"
	draOutboundBytesNotCompressed          = "DRA Outbound Bytes Not Compressed (Within Site) Since Boot"
	draInboundFullSyncObjectsRemaining     = "DRA Inbound Full Sync Objects Remaining"
	draInboundObjects                      = "DRA Inbound Objects/sec"
	draOutboundObjects                     = "DRA Outbound Objects/sec"
	draInboundProperties                   = "DRA Inbound Properties Total/sec"
	draOutboundProperties                  = "DRA Outbound Properties/sec"
	draInboundValuesDNs                    = "DRA Inbound Values (DNs only)/sec" //revive:disable-line:var-naming
	draInboundValuesTotal                  = "DRA Inbound Values Total/sec"
	draOutboundValuesDNs                   = "DRA Outbound Values (DNs only)/sec" //revive:disable-line:var-naming
	draOutboundValuesTotal                 = "DRA Outbound Values Total/sec"
	draPendingReplicationOperations        = "DRA Pending Replication Operations"
	draSyncFailuresSchemaMismatch          = "DRA Sync Failures on Schema Mismatch"
	draSyncRequestsSuccessful              = "DRA Sync Requests Successful"
	draSyncRequestsMade                    = "DRA Sync Requests Made"
	dsDirectoryReads                       = "DS Directory Reads/sec"
	dsDirectoryWrites                      = "DS Directory Writes/sec"
	dsDirectorySearches                    = "DS Directory Searches/sec"
	dsClientBinds                          = "DS Client Binds/sec"
	dsServerBinds                          = "DS Server Binds/sec"
	dsNameCacheHitRate                     = "DS Name Cache hit rate"
	dsNotifyQueueSize                      = "DS Notify Queue Size"
	dsSecurityDescriptorPropagationsEvents = "DS Security Descriptor Propagations Events"
	dsSearchSubOperations                  = "DS Search sub-operations/sec"
	dsSecurityDescriptorSubOperations      = "DS Security Descriptor sub-operations/sec"
	dsThreadsInUse                         = "DS Threads in Use"
	ldapClientSessions                     = "LDAP Client Sessions"
	ldapBindTime                           = "LDAP Bind Time"
	ldapSuccessfulBinds                    = "LDAP Successful Binds/sec"
	ldapSearches                           = "LDAP Searches/sec"
)

type watchers struct {
	closed bool

	counterNameToWatcher map[string]winperfcounters.PerfCounterWatcher
}

func (w *watchers) Scrape(name string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *watchers) Close() error { _ = "STUB: not implemented"; return nil }

func getWatchers(wc watcherCreator) (*watchers, error) { _ = "STUB: not implemented"; return nil, nil }

type watcherCreator interface {
	Create(counterName string) (winperfcounters.PerfCounterWatcher, error)
}

const (
	instanceName = "NTDS"
	object       = "DirectoryServices"
)

type defaultWatcherCreator struct{}

func (defaultWatcherCreator) Create(counterName string) (winperfcounters.PerfCounterWatcher, error) {
	_ = "STUB: not implemented"
	return *new(winperfcounters.PerfCounterWatcher), nil
}
