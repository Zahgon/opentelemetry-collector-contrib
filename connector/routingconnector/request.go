// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package routingconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector"

import (
	"context"
	"regexp"
)

// This file defines an extremely simple request condition grammar. The goal is to provide a similar feel to OTTL,
// but it's not clear that anything more than a simple comparison is needed.  We can expand this grammar in the
// future if needed. For now, it expects the condition to be in exactly the format:
// 'request["<name>"] <comparator> <value>' where <comparator> is either '==' or '!='.

var (
	requestFieldRegex = regexp.MustCompile(`request\[".*"\]`)
	valueFieldRegex   = regexp.MustCompile(`".*"`)
	comparatorRegex   = regexp.MustCompile(`==|!=`)
)

type requestCondition struct {
	compareFunc   func(string) bool
	attributeName string
}

func parseRequestCondition(condition string) (*requestCondition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rc *requestCondition) matchRequest(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (rc *requestCondition) matchGRPC(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (rc *requestCondition) matchHTTP(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}
