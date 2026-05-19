// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"go/ast"
	"regexp"
)

func ExtractDescriptionFromComment(group *ast.CommentGroup) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func goPrimitiveToSchemaType(typeName string) (SchemaType, bool) {
	_ = "STUB: not implemented"
	return *new(SchemaType), false
}

var importRegExp = regexp.MustCompile(`^(.+?)(?:/([^/"]+))?$`)

func ParseImport(imp *ast.ImportSpec) (string, string) { _ = "STUB: not implemented"; return "", "" }

type TagInfo struct {
	Name      string
	OmitEmpty bool
	Squash    bool
}

func ParseTag(tag *ast.BasicLit) (*TagInfo, bool) { _ = "STUB: not implemented"; return nil, false }
