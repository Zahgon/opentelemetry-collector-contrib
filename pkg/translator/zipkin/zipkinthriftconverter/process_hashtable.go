// Copyright The OpenTelemetry Authors
// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package zipkin // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin/zipkinthriftconverter"

import (
	"github.com/jaegertracing/jaeger-idl/model/v1"
)

type processHashtable struct {
	processes map[uint64][]*model.Process
	extHash   func(*model.Process) uint64
}

func newProcessHashtable() *processHashtable { _ = "STUB: not implemented"; return nil }

func (ph processHashtable) hash(process *model.Process) uint64 { _ = "STUB: not implemented"; return 0 }

// for testing collisions

// add checks if identical Process already exists in the hash table and returns it.
// Otherwise it adds process to the table and returns it.
func (ph processHashtable) add(process *model.Process) *model.Process {
	_ = "STUB: not implemented"
	return nil
}

// reuse existing Process object
