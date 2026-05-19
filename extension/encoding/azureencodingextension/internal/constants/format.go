// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package constants // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/constants"

// Format is the encoding format for Azure logs.
type Format string

const (
	// FormatIdentificationTag is the attribute key used to identify the encoding format of the log.
	FormatIdentificationTag = "encoding.format"

	// Format values at log family level
	FormatActivity           Format = "azure.activity"
	FormatApplicationGateway Format = "azure.application_gateway"
	FormatAudit              Format = "azure.audit"
	FormatAppService         Format = "azure.appservice"
	FormatCdn                Format = "azure.cdn"
	FormatDataFactory        Format = "azure.datafactory"
	FormatFrontDoor          Format = "azure.frontdoor"
	FormatFunctionApp        Format = "azure.function_app"
	FormatMessaging          Format = "azure.messaging"
	FormatStorage            Format = "azure.storage"

	// FormatGeneric is the value when a generic parser is applied: the log has an Azure category
	// but no dedicated parser yet.
	FormatGeneric Format = "azure.generic"
)

// FormatForCategory returns the encoding.format value (log family) for the given Azure log category.
// Unsupported categories return FormatGeneric (generic parser applied).
func FormatForCategory(category string) Format { _ = "STUB: not implemented"; return *new(Format) }

// ProviderFromResourceID extracts the Azure resource provider from a resource ID
// and returns an encoding format string. For Microsoft providers, the "microsoft."
// prefix is stripped (e.g., "Microsoft.Web" becomes "azure.web"). For non-Microsoft
// providers, the full lowercased name is used (e.g., "Sendgrid.Email" becomes
// "azure.sendgrid.email"). Returns "azure.generic" if the resource ID is empty
// or the provider cannot be extracted.
func ProviderFromResourceID(resourceID string) string { _ = "STUB: not implemented"; return "" }
