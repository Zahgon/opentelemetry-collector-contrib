// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package pagingscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/pagingscraper"

import (
	"sync"
	"syscall"

	"golang.org/x/sys/windows"
)

var (
	modKernel32 = windows.NewLazySystemDLL("kernel32.dll")
	modPsapi    = windows.NewLazySystemDLL("psapi.dll")

	procGetNativeSystemInfo = modKernel32.NewProc("GetNativeSystemInfo")
	procEnumPageFilesW      = modPsapi.NewProc("EnumPageFilesW")
)

var (
	pageSize     uint64
	pageSizeOnce sync.Once
)

type systemInfo struct {
	wProcessorArchitecture      uint16
	wReserved                   uint16
	dwPageSize                  uint32
	lpMinimumApplicationAddress uintptr
	lpMaximumApplicationAddress uintptr
	dwActiveProcessorMask       uintptr
	dwNumberOfProcessors        uint32
	dwProcessorType             uint32
	dwAllocationGranularity     uint32
	wProcessorLevel             uint16
	wProcessorRevision          uint16
}

func getPageSize() uint64 { _ = "STUB: not implemented"; return 0 }

//nolint:errcheck

// system type as defined in https://docs.microsoft.com/en-us/windows/win32/api/psapi/ns-psapi-enum_page_file_information
type enumPageFileInformation struct {
	cb         uint32 //nolint:unused
	reserved   uint32 //nolint:unused
	totalSize  uint64
	totalInUse uint64
	peakUsage  uint64 //nolint:unused
}

func getPageFileStats() ([]*pageFileStats, error) { _ = "STUB: not implemented"; return nil, nil }

// the following system call invokes the supplied callback function once for each page file before returning
// see https://docs.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-enumpagefilesw

// system callback as defined in https://docs.microsoft.com/en-us/windows/win32/api/psapi/nc-psapi-penum_page_file_callbackw
func pEnumPageFileCallbackW(pageFiles *[]*pageFileStats, enumPageFileInfo *enumPageFileInformation, lpFilenamePtr *[syscall.MAX_LONG_PATH]uint16) *bool {
	_ = "STUB: not implemented"
	return nil
}

// return true to continue enumerating page files
