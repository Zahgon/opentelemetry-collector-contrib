// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package container // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/container"

import (
	"context"
	"regexp"
	"sync"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/attrs"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const (
	dockerFormat        = "docker"
	crioFormat          = "crio"
	containerdFormat    = "containerd"
	recombineInternalID = "recombine_container_internal"
	dockerPattern       = "^\\{"
	crioPattern         = "^(?P<time>[^ Z]+) (?P<stream>stdout|stderr) (?P<logtag>[^ ]*) ?(?P<log>.*)$"
	containerdPattern   = "^(?P<time>[^ ^Z]+Z) (?P<stream>stdout|stderr) (?P<logtag>[^ ]*) ?(?P<log>.*)$"
	logpathPattern      = "^.*(\\/|\\\\)(?P<namespace>[^_]+)_(?P<pod_name>[^_]+)_(?P<uid>[a-f0-9\\-]+)(\\/|\\\\)(?P<container_name>[^\\._]+)(\\/|\\\\)(?P<restart_count>\\d+)\\.log(\\.\\d{8}-\\d{6})?$"
	logPathField        = attrs.LogFilePath
	crioTimeLayout      = "2006-01-02T15:04:05.999999999Z07:00"
	goTimeLayout        = "2006-01-02T15:04:05.999Z"
)

var (
	dockerMatcher     = regexp.MustCompile(dockerPattern)
	crioMatcher       = regexp.MustCompile(crioPattern)
	containerdMatcher = regexp.MustCompile(containerdPattern)
	pathMatcher       = regexp.MustCompile(logpathPattern)
)

var (
	logFieldsMapping = map[string]string{
		"stream": "log.iostream",
	}
	k8sMetadataMapping = map[string]string{
		"container_name": "k8s.container.name",
		"namespace":      "k8s.namespace.name",
		"pod_name":       "k8s.pod.name",
		"restart_count":  "k8s.container.restart_count",
		"uid":            "k8s.pod.uid",
	}
)

// Parser is an operator that parses Container logs.
type Parser struct {
	helper.ParserOperator
	recombineParser         operator.Operator
	format                  string
	addMetadataFromFilepath bool
	criLogEmitter           *helper.BatchingLogEmitter
	recombineStarted        bool
	recombineStartOnce      sync.Once
	timeLayout              string
}

func (p *Parser) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Send CRI entries as a batch to recombine

// Write all docker/skipped entries as a batch

// Process will parse an entry of Container logs
func (p *Parser) Process(ctx context.Context, entry *entry.Entry) (err error) {
	_ = "STUB: not implemented"
	// Short circuit if the "if" condition does not match
	return nil
}

// parse the message

// parse the message

// send it to the recombine operator

// Stop ensures that the internal recombineParser and criLogEmitter are stopped
// in the proper order without being affected by any possible race conditions.
func (p *Parser) Stop() error { _ = "STUB: not implemented"; return nil }

// nothing is started return

// detectFormat will detect the container log format
func (p *Parser) detectFormat(e *entry.Entry) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// parseCRIO will parse a crio log value based on a fixed regexp
func (*Parser) parseCRIO(value any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// parseContainerd will parse a containerd log value based on a fixed regexp
func (*Parser) parseContainerd(value any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// parseDocker will parse a docker log value as JSON
func (*Parser) parseDocker(value any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// handleTimeAndAttributeMappings handles fields' mappings and k8s meta extraction
func (p *Parser) handleTimeAndAttributeMappings(e *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// handleMoveAttributes moves fields to final attributes
func (*Parser) handleMoveAttributes(e *entry.Entry) error {
	_ = "STUB: not implemented"
	// move `log` to `body` explicitly first to avoid
	// moving after more attributes have been added under the `log.*` key
	return nil
}

// then move the rest of the fields

// extractk8sMetaFromFilePath extracts metadata attributes from logfilePath
func (p *Parser) extractk8sMetaFromFilePath(e *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) consumeEntries(ctx context.Context, entries []*entry.Entry) {
	_ = "STUB: not implemented"
	return
}

func moveField(e *entry.Entry, originalKey, mappedKey string) error {
	_ = "STUB: not implemented"
	return nil
}

func moveFieldToBody(e *entry.Entry, originalKey, mappedKey string) error {
	_ = "STUB: not implemented"
	return nil
}

func parseTime(e *entry.Entry, layout string) error { _ = "STUB: not implemented"; return nil }

// If a timestamp ends with 'Z', it should be interpreted at Zulu (UTC) time

// timeutils.ParseGotime calls timeutils.SetTimestampYear before returning the timeValue
