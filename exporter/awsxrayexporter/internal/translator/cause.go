// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter/internal/translator"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

// ExceptionEventName the name of the exception event.
// TODO: Remove this when collector defines this semantic convention.
const (
	ExceptionEventName              = "exception"
	AwsIndividualHTTPEventName      = "HTTP request failure"
	AwsIndividualHTTPErrorEventType = "aws.http.error.event"
	AwsIndividualHTTPErrorMsgAttr   = "aws.http.error_message"
)

func makeCause(span ptrace.Span, attributes map[string]pcommon.Value, resource pcommon.Resource) (isError, isFault, isThrottle bool,
	filtered map[string]pcommon.Value, cause *awsxray.CauseData,
) {
	_ = "STUB: not implemented"
	return false, false, false, nil, nil
}

// Use OpenCensus behavior if we didn't find any exception events to ease migration.

// The segment status for http spans will be based on their http.statuscode as we found some http
// spans does not fill with status.Code() but always filled with http.statuscode

// Default values

func parseException(exceptionType, message, stacktrace string, isRemote bool, language string) []awsxray.Exception {
	_ = "STUB: not implemented"
	return nil
}

// The PHP SDK formats stack traces exactly like Java would

func fillJavaStacktrace(stacktrace string, exceptions []awsxray.Exception) []awsxray.Exception {
	_ = "STUB: not implemented"
	return nil
}

// Skip first line containing top level message

// Class loader or Java module prefix, remove it

// Skip space after colon too.

// Need to peek lines since the message may have newlines.

// Stack frame (hopefully, user can masquerade since we only have a string), process above.

// String append overhead in this case, but multiline messages should be far less common than single
// line ones.

// when append causes `exceptions` to outgrow its existing
// capacity, re-allocation will happen so the place
// `exception` points to is no longer `exceptions[len(exceptions)-2]`,
// consequently, we cannot write `exception.Cause = newException.ID`
// below.

// We peeked to a line starting with "\tat", a stack frame, so continue straight to processing.

// We skip "..." (common frames) and Suppressed By exceptions.

func fillPythonStacktrace(stacktrace string, exceptions []awsxray.Exception) []awsxray.Exception {
	_ = "STUB: not implemented"
	// Need to read in reverse order so can't use a reader. Python formatted tracebacks always use '\n'
	// for newlines so we can just split on it without worrying about Windows newlines.
	return nil
}

// Skip last line containing top level exception / message

// Couldn't find a "  File ..." line before end of input, malformed stack trace.

// Join message which potentially has newlines. Message starts two lines from the next "File " line and ends
// two lines before the "During handling " line.

// Error not followed by a colon, malformed stack trace.

// when append causes `exceptions` to outgrow its existing
// capacity, re-allocation will happen so the place
// `exception` points to is no longer `exceptions[len(exceptions)-2]`,
// consequently, we cannot write `exception.Cause = newException.ID`
// below.

// lineIdx is set to the next File line so ready to process it.

func fillJavaScriptStacktrace(stacktrace string, exceptions []awsxray.Exception) []awsxray.Exception {
	_ = "STUB: not implemented"
	return nil
}

// Skip first line containing top level message

// only append the exception if at least one of the values is not default

func fillDotnetStacktrace(stacktrace string, exceptions []awsxray.Exception) []awsxray.Exception {
	_ = "STUB: not implemented"
	return nil
}

// Skip first line containing top level message

func fillGoStacktrace(stacktrace string, exceptions []awsxray.Exception) []awsxray.Exception {
	_ = "STUB: not implemented"
	return nil
}

// Skip first line containing top level message

// indexOf returns position of the first occurrence of a Byte in str starting at pos index.
func indexOf(str string, c byte, pos int) int { _ = "STUB: not implemented"; return 0 }
