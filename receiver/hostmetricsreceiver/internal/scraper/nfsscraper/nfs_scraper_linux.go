// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package nfsscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/nfsscraper"

import (
	"io"
)

const (
	nfsProcFile  = "/proc/net/rpc/nfs"
	nfsdProcFile = "/proc/net/rpc/nfsd"
)

// from linux/fs/nfs/nfs3xdr.c:nfs3_procedures v6.12
// note that this array starts at element 1, with a NULL at element 0
var nfsV3Procedures = []string{
	"NULL",
	"GETATTR",
	"SETATTR",
	"LOOKUP",
	"ACCESS",
	"READLINK",
	"READ",
	"WRITE",
	"CREATE",
	"MKDIR",
	"SYMLINK",
	"MKNOD",
	"REMOVE",
	"RMDIR",
	"RENAME",
	"LINK",
	"READDIR",
	"READDIRPLUS",
	"FSSTAT",
	"FSINFO",
	"PATHCONF",
	"COMMIT",
}

// from linux/fs/nfs/nfs4xdr.c:nfs4_procedures v6.12
// note that these are technically NFSv4 operations, part of a COMPOUND RPC procedure
// note that this array starts at element 1, with a NULL at element 0
var nfsV4Procedures = []string{
	"NULL",
	"READ",
	"WRITE",
	"COMMIT",
	"OPEN",
	"OPEN_CONFIRM",
	"OPEN_NOATTR",
	"OPEN_DOWNGRADE",
	"CLOSE",
	"SETATTR",
	"FSINFO",
	"RENEW",
	"SETCLIENTID",
	"SETCLIENTID_CONFIRM",
	"LOCK",
	"LOCKT",
	"LOCKU",
	"ACCESS",
	"GETATTR",
	"LOOKUP",
	"LOOKUP_ROOT",
	"REMOVE",
	"RENAME",
	"LINK",
	"SYMLINK",
	"CREATE",
	"PATHCONF",
	"STATFS",
	"READLINK",
	"READDIR",
	"SERVER_CAPS",
	"DELEGRETURN",
	"GETACL",
	"SETACL",
	"FS_LOCATIONS",
	"RELEASE_LOCKOWNER",
	"SECINFO",
	"FSID_PRESENT",
	"EXCHANGE_ID",
	"CREATE_SESSION",
	"DESTROY_SESSION",
	"SEQUENCE",
	"GET_LEASE_TIME",
	"RECLAIM_COMPLETE",
	"GETDEVICEINFO",
	"LAYOUTGET",
	"LAYOUTCOMMIT",
	"LAYOUTRETURN",
	"SECINFO_NO_NAME",
	"TEST_STATEID",
	"FREE_STATEID",
	"GETDEVICELIST",
	"BIND_CONN_TO_SESSION",
	"DESTROY_CLIENTID",
	"SEEK",
	"ALLOCATE",
	"DEALLOCATE",
	"LAYOUTSTATS",
	"CLONE",
	"COPY",
	"OFFLOAD_CANCEL",
	"COPY_NOTIFY",
	"LOOKUPP",
	"LAYOUTERROR",
	"GETXATTR",
	"SETXATTR",
	"LISTXATTRS",
	"REMOVEXATTR",
	"READ_PLUS",
}

// from linux/fs/nfsd/nfs3proc.c:nfsd_procedures3 v6.12
var nfsdV3Procedures = []string{
	"NULL",
	"GETATTR",
	"SETATTR",
	"LOOKUP",
	"ACCESS",
	"READLINK",
	"READ",
	"WRITE",
	"CREATE",
	"MKDIR",
	"SYMLINK",
	"MKNOD",
	"REMOVE",
	"RMDIR",
	"RENAME",
	"LINK",
	"READDIR",
	"READDIRPLUS",
	"FSSTAT",
	"FSINFO",
	"PATHCONF",
	"COMMIT",
}

// from linux/fs/nfsd/nfs4proc.c:nfsd_procedures4 v6.12
var nfsdV4Procedures = []string{
	"NULL",
	"COMPOUND",
}

// from linux/include/linux/nfs4.h:nfs_opnum4 v6.12
var nfsdV4Operations = []string{
	"UNUSED_IGNORE0",
	"UNUSED_IGNORE1",
	"UNUSED_IGNORE2",
	"ACCESS",
	"CLOSE",
	"COMMIT",
	"CREATE",
	"DELEGPURGE",
	"DELEGRETURN",
	"GETATTR",
	"GETFH",
	"LINK",
	"LOCK",
	"LOCKT",
	"LOCKU",
	"LOOKUP",
	"LOOKUPP",
	"NVERIFY",
	"OPEN",
	"OPENATTR",
	"OPEN_CONFIRM",
	"OPEN_DOWNGRADE",
	"PUTFH",
	"PUTPUBFH",
	"PUTROOTFH",
	"READ",
	"READDIR",
	"READLINK",
	"REMOVE",
	"RENAME",
	"RENEW",
	"RESTOREFH",
	"SAVEFH",
	"SECINFO",
	"SETATTR",
	"SETCLIENTID",
	"SETCLIENTID_CONFIRM",
	"VERIFY",
	"WRITE",
	"RELEASE_LOCKOWNER",
	"BACKCHANNEL_CTL",
	"BIND_CONN_TO_SESSION",
	"EXCHANGE_ID",
	"CREATE_SESSION",
	"DESTROY_SESSION",
	"FREE_STATEID",
	"GET_DIR_DELEGATION",
	"GETDEVICEINFO",
	"GETDEVICELIST",
	"LAYOUTCOMMIT",
	"LAYOUTGET",
	"LAYOUTRETURN",
	"SECINFO_NO_NAME",
	"SEQUENCE",
	"SET_SSV",
	"TEST_STATEID",
	"WANT_DELEGATION",
	"DESTROY_CLIENTID",
	"RECLAIM_COMPLETE",
	"ALLOCATE",
	"COPY",
	"COPY_NOTIFY",
	"DEALLOCATE",
	"IO_ADVISE",
	"LAYOUTERROR",
	"LAYOUTSTATS",
	"OFFLOAD_CANCEL",
	"OFFLOAD_STATUS",
	"READ_PLUS",
	"SEEK",
	"WRITE_SAME",
	"CLONE",
	"GETXATTR",
	"SETXATTR",
	"LISTXATTRS",
	"REMOVEXATTR",
}

func getOSNfsStats() (*NfsStats, error) {
	_ = "STUB: not implemented"
	/* for testing: (nfsProcFileOut from nfs_scraper_linux_test.go)
	   data := strings.NewReader(nfsProcFileOut)

	   rv, err := parseNfsStats(data)
	   fmt.Fprintf(os.Stderr, "%#v\n", rv)
	   return rv, err
	*/return nil, nil
}

func getOSNfsdStats() (*nfsdStats, error) {
	_ = "STUB: not implemented"
	/* for testing: (nfsdProcFileOut from nfs_scraper_linux_test.go)
	   data := strings.NewReader(nfsdProcFileOut)

	   rv, err := parseNfsdStats(data)
	   fmt.Fprintf(os.Stderr, "%#v\n", rv)
	   return rv, err
	*/return nil, nil
}

func parseNfsNetStats(values []uint64) (*nfsNetStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseNfsRPCStats(values []uint64) (*nfsRPCStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseNfsdNetStats(values []uint64) (*nfsdNetStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseNfsdRepcacheStats(values []uint64) (*nfsdRepcacheStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseNfsdFhStats(values []uint64) (*nfsdFhStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseNfsdIoStats(values []uint64) (*nfsdIoStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseNfsdThreadStats(values []uint64) (*nfsdThreadStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseNfsdRPCStats(values []uint64) (*nfsdRPCStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseNfsCallStats(nfsVersion int64, names []string, values []uint64) ([]callStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// first element is numCalls

// found yet-to-be-supported procedures

func parseNfsStats(f io.Reader) (*NfsStats, error) { _ = "STUB: not implemented"; return nil, nil }

// Linux kernel calls NFSv4 client operations procedures, but they're actually
// operations of compound procedures, per RFC7530

func parseNfsdStats(f io.Reader) (*nfsdStats, error) { _ = "STUB: not implemented"; return nil, nil }

// th has a mix of uint64 and double. the double values are obsolete (always 0.000)

// parseStringsToUint64s parses a slice of strings into a slice of uint64
func parseStringsToUint64s(strSlice []string) ([]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Assuming decimal base and uint64

func CanScrapeAll() bool { _ = "STUB: not implemented"; return false }
