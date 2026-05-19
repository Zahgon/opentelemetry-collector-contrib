// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package db // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor/internal/db"

import (
	"github.com/DataDog/datadog-agent/pkg/obfuscate"
	"go.uber.org/zap"
)

type databaseObfuscator interface {
	Obfuscate(string) (string, error)
	ObfuscateWithSystem(string, string) (string, error)
	ObfuscateAttribute(string, string) (string, error)
	ShouldProcessAttribute(string) bool
	SupportsSystem(string) bool
}

type dbAttributes struct {
	attributes map[string]bool
	dbSystems  map[string]bool
}

func (d *dbAttributes) ShouldProcessAttribute(attributeKey string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *dbAttributes) SupportsSystem(dbSystem string) bool {
	_ = "STUB: not implemented"
	return false
}

type sqlObfuscator struct {
	dbAttributes
	obfuscator *obfuscate.Obfuscator
}

var _ databaseObfuscator = &sqlObfuscator{}

func (o *sqlObfuscator) Obfuscate(s string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *sqlObfuscator) ObfuscateAttribute(attributeValue, attributeKey string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *sqlObfuscator) ObfuscateWithSystem(s, dbSystem string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type redisObfuscator struct {
	dbAttributes
	obfuscator *obfuscate.Obfuscator
}

var _ databaseObfuscator = &redisObfuscator{}

func (o *redisObfuscator) Obfuscate(s string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *redisObfuscator) ObfuscateAttribute(attributeValue, attributeKey string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *redisObfuscator) ObfuscateWithSystem(s, _ string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type memcachedObfuscator struct {
	dbAttributes
	obfuscator *obfuscate.Obfuscator
}

var _ databaseObfuscator = &memcachedObfuscator{}

func (o *memcachedObfuscator) Obfuscate(s string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *memcachedObfuscator) ObfuscateAttribute(attributeValue, attributeKey string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *memcachedObfuscator) ObfuscateWithSystem(s, _ string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type mongoObfuscator struct {
	dbAttributes
	obfuscator *obfuscate.Obfuscator
	logger     *zap.Logger
}

var _ databaseObfuscator = &mongoObfuscator{}

func (o *mongoObfuscator) Obfuscate(s string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *mongoObfuscator) ObfuscateAttribute(attributeValue, attributeKey string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *mongoObfuscator) ObfuscateWithSystem(s, _ string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type opensearchObfuscator struct {
	dbAttributes
	obfuscator *obfuscate.Obfuscator
	logger     *zap.Logger
}

var _ databaseObfuscator = &opensearchObfuscator{}

func (o *opensearchObfuscator) Obfuscate(s string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *opensearchObfuscator) ObfuscateAttribute(attributeValue, attributeKey string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *opensearchObfuscator) ObfuscateWithSystem(s, _ string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type esObfuscator struct {
	dbAttributes
	obfuscator *obfuscate.Obfuscator
	logger     *zap.Logger
}

var _ databaseObfuscator = &esObfuscator{}

func (o *esObfuscator) Obfuscate(s string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *esObfuscator) ObfuscateAttribute(attributeValue, attributeKey string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *esObfuscator) ObfuscateWithSystem(s, _ string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func isValidJSON(value string) bool { _ = "STUB: not implemented"; return false }
