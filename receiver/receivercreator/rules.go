// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package receivercreator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator"

import (
	"fmt"
	"regexp"

	"github.com/expr-lang/expr/vm"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

// rule wraps expr rule for later evaluation.
type rule struct {
	program *vm.Program
}

// ruleRe is used to verify the rule starts type check.
var ruleRe = regexp.MustCompile(
	fmt.Sprintf(`^type\s*==\s*(%q|%q|%q|%q|%q|%q|%q|%q|%q)`, observer.PodType, observer.K8sServiceType, observer.K8sIngressType, observer.PortType, observer.PodContainerType, observer.HostPortType, observer.ContainerType, observer.K8sNodeType, observer.KafkaTopicType),
)

// newRule creates a new rule instance.
func newRule(ruleStr string) (rule, error) { _ = "STUB: not implemented"; return *new(rule), nil }

// TODO: Try validating against bytecode instead.

// TODO: Maybe use https://godoc.org/github.com/expr-lang/expr#Env in type checking
// depending on type == specified.

// expr v1.14.1 introduced a `type` builtin whose implementation we relocate to `typeOf`
// to avoid collision

// eval the rule against the given endpoint.
func (r *rule) eval(env observer.EndpointEnv) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
