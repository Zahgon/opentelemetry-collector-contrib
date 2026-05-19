// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// tools package includes the list of tools available for the MCP extension. The contents of this
// file originated in https://github.com/pavolloffay/opentelemetry-mcp-server/blob/main/internal/tools/tools.go
package tools // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/mcp/internal/mcp/tools"

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/pavolloffay/opentelemetry-mcp-server/modules/collectorschema"
)

// Tool represents an MCP tool with its handler
type Tool struct {
	Tool    *mcp.Tool
	Handler mcp.ToolHandler
}

// GetAllTools returns a list of all available MCP tools
func GetAllTools() ([]Tool, error) { _ = "STUB: not implemented"; return nil, nil }

func parseArgs(req *mcp.CallToolRequest) map[string]any { _ = "STUB: not implemented"; return nil }

func textResult(text string) *mcp.CallToolResult { _ = "STUB: not implemented"; return nil }

func errResult(format string, args ...any) *mcp.CallToolResult {
	_ = "STUB: not implemented"
	return nil
}

func jsonResult(v any) *mcp.CallToolResult { _ = "STUB: not implemented"; return nil }

func stringArg(args map[string]any, key, defaultVal string) string {
	_ = "STUB: not implemented"
	return ""
}

func requireStringArg(args map[string]any, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func requireStringSliceArg(args map[string]any, key string) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func getCollectorVersionsTool(schemaManager *collectorschema.SchemaManager) Tool {
	_ = "STUB: not implemented"
	return *new(Tool)
}

func getCollectorComponentsTool(schemaManager *collectorschema.SchemaManager, latestCollectorVersion string) Tool {
	_ = "STUB: not implemented"
	return *new(Tool)
}

func getCollectorReadmeTool(schemaManager *collectorschema.SchemaManager, latestCollectorVersion string) Tool {
	_ = "STUB: not implemented"
	return *new(Tool)
}

func getCollectorChangelogTool(schemaManager *collectorschema.SchemaManager, latestCollectorVersion string) Tool {
	_ = "STUB: not implemented"
	return *new(Tool)
}

func getCollectorSchemaGetTool(schemaManager *collectorschema.SchemaManager, latestCollectorVersion string) Tool {
	_ = "STUB: not implemented"
	return *new(Tool)
}

func getCollectorSchemaValidationTool(schemaManager *collectorschema.SchemaManager, latestCollectorVersion string) Tool {
	_ = "STUB: not implemented"
	return *new(Tool)
}

type DeprecatedComponentFields struct {
	ComponentName    string                            `json:"componentName"`
	DeprecatedFields []collectorschema.DeprecatedField `json:"deprecatedFields"`
}

func getCollectorComponentDeprecatedTool(schemaManager *collectorschema.SchemaManager, latestCollectorVersion string) Tool {
	_ = "STUB: not implemented"
	return *new(Tool)
}

type DocumentationSearchResult struct {
	Results []collectorschema.DocumentSearchResult `json:"results"`
}

func getCollectorDocumentationRAG(schemaManager *collectorschema.SchemaManager, latestCollectorVersion string) Tool {
	_ = "STUB: not implemented"
	return *new(Tool)
}

func boolPtr(b bool) *bool { _ = "STUB: not implemented"; return nil }
