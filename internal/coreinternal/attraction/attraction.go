// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package attraction // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/attraction"

import (
	"context"
	"regexp"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"
)

// Settings specifies the processor settings.
type Settings struct {
	// Actions specifies the list of attributes to act on.
	// The set of actions are {INSERT, UPDATE, UPSERT, DELETE, HASH, EXTRACT, CONVERT}.
	// This is a required field.
	Actions []ActionKeyValue `mapstructure:"actions"`
}

// ActionKeyValue specifies the attribute key to act upon.
type ActionKeyValue struct {
	// Key specifies the attribute to act upon.
	// This is a required field.
	Key string `mapstructure:"key"`

	// Value specifies the value to populate for the key.
	// The type of the value is inferred from the configuration.
	Value any `mapstructure:"value"`

	// DefaultValue specifies the value to use if Value/FromAttribute/FromContext
	// doesn't provide a value (e.g., environment variable not set, attribute doesn't exist).
	// Only used with INSERT, UPDATE, and UPSERT actions.
	DefaultValue any `mapstructure:"default_value"`

	// A regex pattern must be specified for the action EXTRACT.
	// It uses the attribute specified by `key' to extract values from
	// The target keys are inferred based on the names of the matcher groups
	// provided and the names will be inferred based on the values of the
	// matcher group.
	// Note: All subexpressions must have a name.
	// Note: The value type of the source key must be a string. If it isn't,
	// no extraction will occur.
	RegexPattern string `mapstructure:"pattern"`

	// FromAttribute specifies the attribute to use to populate
	// the value. If the attribute doesn't exist, no action is performed.
	FromAttribute string `mapstructure:"from_attribute"`

	// FromContext specifies the context value to use to populate
	// the value. The values would be searched in client.Info.Metadata.
	// If the key doesn't exist, no action is performed.
	// If the key has multiple values the values will be joined with `;` separator.
	FromContext string `mapstructure:"from_context"`

	// ConvertedType specifies the target type of an attribute to be converted
	// If the key doesn't exist, no action is performed.
	// If the value cannot be converted, the original value will be left as-is
	ConvertedType string `mapstructure:"converted_type"`

	// Action specifies the type of action to perform.
	// The set of values are {INSERT, UPDATE, UPSERT, DELETE, HASH}.
	// Both lower case and upper case are supported.
	// INSERT -  Inserts the key/value to attributes when the key does not exist.
	//           No action is applied to attributes where the key already exists.
	//           Either Value, FromAttribute or FromContext must be set.
	// UPDATE -  Updates an existing key with a value. No action is applied
	//           to attributes where the key does not exist.
	//           Either Value, FromAttribute or FromContext must be set.
	// UPSERT -  Performs insert or update action depending on the attributes
	//           containing the key. The key/value is inserted to attributes
	//           that did not originally have the key. The key/value is updated
	//           for attributes where the key already existed.
	//           Either Value, FromAttribute or FromContext must be set.
	// DELETE  - Deletes the attribute. If the key doesn't exist,
	//           no action is performed.
	// HASH    - Calculates the SHA-1 hash of an existing value and overwrites the
	//           value with its SHA-1 hash result. If the feature gate
	//           `coreinternal.attraction.hash.sha256` is enabled, it uses SHA2-256
	//           instead.
	// EXTRACT - Extracts values using a regular expression rule from the input
	//           'key' to target keys specified in the 'rule'. If a target key
	//           already exists, it will be overridden.
	// CONVERT  - converts the type of an existing attribute, if convertable
	// This is a required field.
	Action Action `mapstructure:"action"`
}

func (a *ActionKeyValue) valueSourceCount() int { _ = "STUB: not implemented"; return 0 }

// Action is the enum to capture the four types of actions to perform on an
// attribute.
type Action string

const (
	// INSERT adds the key/value to attributes when the key does not exist.
	// No action is applied to attributes where the key already exists.
	INSERT Action = "insert"

	// UPDATE updates an existing key with a value. No action is applied
	// to attributes where the key does not exist.
	UPDATE Action = "update"

	// UPSERT performs the INSERT or UPDATE action. The key/value is
	// inserted to attributes that did not originally have the key. The key/value is
	// updated for attributes where the key already existed.
	UPSERT Action = "upsert"

	// DELETE deletes the attribute. If the key doesn't exist, no action is performed.
	// Supports pattern which is matched against attribute key.
	DELETE Action = "delete"

	// HASH calculates the SHA-256 hash of an existing value and overwrites the
	// value with it's SHA-256 hash result.
	// Supports pattern which is matched against attribute key.
	HASH Action = "hash"

	// EXTRACT extracts values using a regular expression rule from the input
	// 'key' to target keys specified in the 'rule'. If a target key already
	// exists, it will be overridden.
	EXTRACT Action = "extract"

	// CONVERT converts the type of an existing attribute, if convertable
	CONVERT Action = "convert"
)

type attributeAction struct {
	Key           string
	FromAttribute string
	FromContext   string
	ConvertedType string
	DefaultValue  *pcommon.Value
	// Compiled regex if provided
	Regex *regexp.Regexp
	// Attribute names extracted from the regexp's subexpressions.
	AttrNames []string
	// Number of non empty strings in above array

	// TODO https://go.opentelemetry.io/collector/issues/296
	// Do benchmark testing between having action be of type string vs integer.
	// The reason is attributes processor will most likely be commonly used
	// and could impact performance.
	Action         Action
	AttributeValue *pcommon.Value
}

// AttrProc is an attribute processor.
type AttrProc struct {
	actions []attributeAction
}

// NewAttrProc validates that the input configuration has all of the required fields for the processor
// and returns a AttrProc to be used to process attributes.
// An error is returned if there are any invalid inputs.
func NewAttrProc(settings *Settings) (*AttrProc, error) { _ = "STUB: not implemented"; return nil, nil }

// Convert `action` to lowercase for comparison.

// requires `key` and/or `pattern`

// `key` is a required field

// Convert the raw value from the configuration to the internal trace representation of the value.

// Handle default_value

// Process applies the AttrProc to an attribute map.
func (ap *AttrProc) Process(ctx context.Context, logger *zap.Logger, attrs pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// TODO https://go.opentelemetry.io/collector/issues/296
// Do benchmark testing between having action be of type string vs integer.
// The reason is attributes processor will most likely be commonly used
// and could impact performance.

func getAttributeValueFromContext(ctx context.Context, key string) (pcommon.Value, bool) {
	_ = "STUB: not implemented"
	return *new(pcommon.Value), false
}

// TODO: Warn about unexpected attribute types.

// Fallback to metadata for backwards compatibility.

func getSourceAttributeValue(ctx context.Context, action attributeAction, attrs pcommon.Map) (pcommon.Value, bool) {
	_ = "STUB: not implemented"
	// Set the key with a value from the configuration.
	return *new(pcommon.Value), false
}

// If FromContext didn't find the value, fall through to try FromAttribute and DefaultValue

// If FromAttribute didn't find the value, fall through to DefaultValue

// Use default value if no other source provided a value

func convertAttribute(logger *zap.Logger, action attributeAction, attrs pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func extractAttributes(action attributeAction, attrs pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// Extracting values only functions on strings.

// Note: The number of matches will always be equal to number of
// subexpressions.

// Start from index 1, which is the first submatch (index 0 is the entire
// match).
