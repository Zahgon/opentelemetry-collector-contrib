// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sclusterreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver"

// transformObject transforms the k8s object by removing the data that is not utilized by the receiver.
// Only highly utilized objects are transformed here while others are kept as is.
func transformObject(object any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }
