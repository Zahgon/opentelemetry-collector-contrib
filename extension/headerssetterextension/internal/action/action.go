// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package action // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/headerssetterextension/internal/action"

import "net/http"

type Action interface {
	ApplyOnHeaders(http.Header, string)
	ApplyOnMetadata(map[string]string, string)
}

type Insert struct {
	Key string
}

func (a Insert) ApplyOnHeaders(header http.Header, value string) { _ = "STUB: not implemented"; return }

func (a Insert) ApplyOnMetadata(metadata map[string]string, value string) {
	_ = "STUB: not implemented"
	return
}

type Update struct {
	Key string
}

func (a Update) ApplyOnHeaders(header http.Header, value string) { _ = "STUB: not implemented"; return }

func (a Update) ApplyOnMetadata(metadata map[string]string, value string) {
	_ = "STUB: not implemented"
	return
}

type Delete struct {
	Key string
}

func (a Delete) ApplyOnHeaders(header http.Header, _ string) { _ = "STUB: not implemented"; return }

func (a Delete) ApplyOnMetadata(metadata map[string]string, _ string) {
	_ = "STUB: not implemented"
	return
}

type Upsert struct {
	Key string
}

func (a Upsert) ApplyOnHeaders(header http.Header, value string) { _ = "STUB: not implemented"; return }

func (a Upsert) ApplyOnMetadata(metadata map[string]string, value string) {
	_ = "STUB: not implemented"
	return
}
