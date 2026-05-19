// Go API over pdh syscalls
//go:build windows

package win_perf_counters // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters/internal/third_party/telegraf/win_perf_counters"

import (
	"time"
)

// CounterValue is abstraction for PDH_FMT_COUNTERVALUE_ITEM_DOUBLE
type CounterValue struct {
	InstanceName string
	Value        float64
}

// RawCounterValue is abstraction for PDH_RAW_COUNTER_ITEM
type RawCounterValue struct {
	InstanceName string
	RawValue     int64
}

// PerformanceQuery provides wrappers around Windows performance counters API for easy usage in GO
type PerformanceQuery interface {
	Open() error
	Close() error
	AddCounterToQuery(counterPath string) (PDH_HCOUNTER, error)
	AddEnglishCounterToQuery(counterPath string) (PDH_HCOUNTER, error)
	GetCounterPath(counterHandle PDH_HCOUNTER) (string, error)
	GetFormattedCounterValueDouble(hCounter PDH_HCOUNTER) (float64, error)
	GetFormattedCounterArrayDouble(hCounter PDH_HCOUNTER) ([]CounterValue, error)
	GetRawCounterValue(hCounter PDH_HCOUNTER) (int64, error)
	GetRawCounterArray(hCounter PDH_HCOUNTER) ([]RawCounterValue, error)
	CollectData() error
	CollectDataWithTime() (time.Time, error)
	IsVistaOrNewer() bool
}

// PdhError represents error returned from Performance Counters API
type PdhError struct {
	ErrorCode uint32
	errorText string
}

func (m *PdhError) Error() string { _ = "STUB: not implemented"; return "" }

func NewPdhError(code uint32) error { _ = "STUB: not implemented"; return nil }

// PerformanceQueryImpl is implementation of PerformanceQuery interface, which calls phd.dll functions
type PerformanceQueryImpl struct {
	query PDH_HQUERY
}

// Open creates a new counterPath that is used to manage the collection of performance data.
// It returns counterPath handle used for subsequent calls for adding counters and querying data
func (m *PerformanceQueryImpl) Open() error { _ = "STUB: not implemented"; return nil }

// Close closes the counterPath, releases associated counter handles and frees resources
func (m *PerformanceQueryImpl) Close() error { _ = "STUB: not implemented"; return nil }

func (m *PerformanceQueryImpl) AddCounterToQuery(counterPath string) (PDH_HCOUNTER, error) {
	_ = "STUB: not implemented"
	return *new(PDH_HCOUNTER), nil
}

func (m *PerformanceQueryImpl) AddEnglishCounterToQuery(counterPath string) (PDH_HCOUNTER, error) {
	_ = "STUB: not implemented"
	return *new(PDH_HCOUNTER), nil
}

// GetCounterPath return counter information for given handle
func (m *PerformanceQueryImpl) GetCounterPath(counterHandle PDH_HCOUNTER) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetFormattedCounterValueDouble computes a displayable value for the specified counter
func (m *PerformanceQueryImpl) GetFormattedCounterValueDouble(hCounter PDH_HCOUNTER) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *PerformanceQueryImpl) GetFormattedCounterArrayDouble(hCounter PDH_HCOUNTER) ([]CounterValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *PerformanceQueryImpl) GetRawCounterValue(hCounter PDH_HCOUNTER) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *PerformanceQueryImpl) GetRawCounterArray(hCounter PDH_HCOUNTER) ([]RawCounterValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *PerformanceQueryImpl) CollectData() error { _ = "STUB: not implemented"; return nil }

func (m *PerformanceQueryImpl) CollectDataWithTime() (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (m *PerformanceQueryImpl) IsVistaOrNewer() bool { _ = "STUB: not implemented"; return false }

// ExpandWildCardPath examines local computer and returns those counter paths that match the given counter path which contains wildcard characters.
func ExpandWildCardPath(counterPath string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UTF16PtrToString converts Windows API LPTSTR (pointer to string) to go string
func UTF16PtrToString(s *uint16) string { _ = "STUB: not implemented"; return "" }

// UTF16ToStringArray converts list of Windows API NULL terminated strings to go string array
func UTF16ToStringArray(buf []uint16) []string { _ = "STUB: not implemented"; return nil }
