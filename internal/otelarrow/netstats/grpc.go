// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package netstats // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/netstats"

import "google.golang.org/grpc"

// GRPCStreamMethodName applies the logic gRPC uses but does not expose to construct
// method names.  This allows direct calling of the netstats interface
// from outside a gRPC stats handler.
func GRPCStreamMethodName(desc grpc.ServiceDesc, stream grpc.StreamDesc) string {
	_ = "STUB: not implemented"
	return ""
}

// GRPCUnaryMethodName applies the logic gRPC uses but does not expose to construct
// method names.  This allows direct calling of the netstats interface
// from outside a gRPC stats handler.
func GRPCUnaryMethodName(desc grpc.ServiceDesc, method grpc.MethodDesc) string {
	_ = "STUB: not implemented"
	return ""
}
