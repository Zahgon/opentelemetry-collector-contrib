// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"container/list"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/packages"
)

type Parser struct {
	config       *Config
	schema       *Schema
	types        map[string]*TypeInfo
	processQueue *list.List
	pkg          *packages.Package
	current      *TypeInfo
}

type TypeInfo struct {
	spec      *ast.TypeSpec
	comms     []*ast.CommentGroup
	imports   map[string]string
	typeName  string
	processed bool
}

func NewParser(cfg *Config) *Parser { _ = "STUB: not implemented"; return nil }

func (p *Parser) Parse() (*Schema, error) { _ = "STUB: not implemented"; return nil, nil }

// select types to process

// process types

func (p *Parser) processPackages(set *token.FileSet, pkgs []*packages.Package) {
	_ = "STUB: not implemented"
	return
}

func (p *Parser) collectTypesAndImports(file *ast.File, pkgPath string, cmap ast.CommentMap) {
	_ = "STUB: not implemented"

	// collect imports from current file, distinguish internal vs external
	return
}

// collect exported type specs

func (p *Parser) feedProcessQueue() error { _ = "STUB: not implemented"; return nil }

// in component mode process only the config type initially

// in package mode process all exported types

func (p *Parser) processTypes() error { _ = "STUB: not implemented"; return nil }

// pick next type to process from the queue

// skip already processed types

// skip struct types with no exported fields

// add parsed type to schema

func (p *Parser) isConfigType(typeInfo *TypeInfo) bool { _ = "STUB: not implemented"; return false }

func (p *Parser) parseType(typeInfo *TypeInfo) (SchemaElement, error) {
	_ = "STUB: not implemented"
	return *

	// skip non-struct types at the top level
	new(SchemaElement), nil
}

// skip these types

func (p *Parser) parseExpr(expr ast.Expr) (SchemaElement, error) {
	_ = "STUB: not implemented"
	return *new(SchemaElement), nil
}

func (p *Parser) parseStruct(structType *ast.StructType) (SchemaElement, error) {
	_ = "STUB: not implemented"
	return *new(SchemaElement), nil
}

func (p *Parser) addEmbeddedField(field *ast.Field, schemaObject SchemaObject) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) addNamedFields(field *ast.Field, schemaObject SchemaObject) {
	_ = "STUB: not implemented"
	return
}

func (p *Parser) addNamedField(fieldName string, field *ast.Field, schemaObject SchemaObject) {
	_ = "STUB: not implemented"
	return
}

func (p *Parser) parseArray(array *ast.ArrayType) (SchemaElement, error) {
	_ = "STUB: not implemented"
	return *new(SchemaElement), nil
}

func (p *Parser) parseIdent(ident *ast.Ident) (SchemaElement, error) {
	_ = "STUB: not implemented"
	return *new(SchemaElement), nil
}

func (p *Parser) parseMap(m *ast.MapType) (SchemaElement, error) {
	_ = "STUB: not implemented"
	return *new(SchemaElement), nil
}

func (p *Parser) parsePointer(pointer *ast.StarExpr) (SchemaElement, error) {
	_ = "STUB: not implemented"
	return *new(SchemaElement), nil
}

func (p *Parser) parseSelector(selector *ast.SelectorExpr) (SchemaElement, error) {
	_ = "STUB: not implemented"
	return *new(SchemaElement), nil
}

// always allow internal packages

// otherwise check allowed refs

// if allowed - create ref, else create any with custom type

// if ref is in the same namespace/repository

func (p *Parser) parseOptional(indexExpr *ast.IndexExpr) (SchemaElement, error) {
	_ = "STUB: not implemented"
	return *new(SchemaElement), nil
}
