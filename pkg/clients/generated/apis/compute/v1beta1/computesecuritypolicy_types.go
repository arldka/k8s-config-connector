// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SecuritypolicyAdaptiveProtectionConfig struct {
	/* Auto Deploy Config of this security policy. */
	// +optional
	AutoDeployConfig *SecuritypolicyAutoDeployConfig `json:"autoDeployConfig,omitempty"`

	/* Layer 7 DDoS Defense Config of this security policy. */
	// +optional
	Layer7DdosDefenseConfig *SecuritypolicyLayer7DdosDefenseConfig `json:"layer7DdosDefenseConfig,omitempty"`
}

type SecuritypolicyAdvancedOptionsConfig struct {
	/* Custom configuration to apply the JSON parsing. Only applicable when JSON parsing is set to STANDARD. */
	// +optional
	JsonCustomConfig *SecuritypolicyJsonCustomConfig `json:"jsonCustomConfig,omitempty"`

	/* JSON body parsing. Supported values include: "DISABLED", "STANDARD". */
	// +optional
	JsonParsing *string `json:"jsonParsing,omitempty"`

	/* Logging level. Supported values include: "NORMAL", "VERBOSE". */
	// +optional
	LogLevel *string `json:"logLevel,omitempty"`
}

type SecuritypolicyAutoDeployConfig struct {
	/* Rules are only automatically deployed for alerts on potential attacks with confidence scores greater than this threshold. */
	// +optional
	ConfidenceThreshold *float64 `json:"confidenceThreshold,omitempty"`

	/* Google Cloud Armor stops applying the action in the automatically deployed rule to an identified attacker after this duration. The rule continues to operate against new requests. */
	// +optional
	ExpirationSec *int `json:"expirationSec,omitempty"`

	/* Rules are only automatically deployed when the estimated impact to baseline traffic from the suggested mitigation is below this threshold. */
	// +optional
	ImpactedBaselineThreshold *float64 `json:"impactedBaselineThreshold,omitempty"`

	/* Identifies new attackers only when the load to the backend service that is under attack exceeds this threshold. */
	// +optional
	LoadThreshold *float64 `json:"loadThreshold,omitempty"`
}

type SecuritypolicyBanThreshold struct {
	/* Number of HTTP(S) requests for calculating the threshold. */
	Count int `json:"count"`

	/* Interval over which the threshold is computed. */
	IntervalSec int `json:"intervalSec"`
}

type SecuritypolicyConfig struct {
	/* Set of IP addresses or ranges (IPV4 or IPV6) in CIDR notation to match against inbound traffic. There is a limit of 10 IP ranges per rule. A value of '*' matches all IPs (can be used to override the default behavior). */
	SrcIpRanges []string `json:"srcIpRanges"`
}

type SecuritypolicyEnforceOnKeyConfigs struct {
	/* Rate limit key name applicable only for the following key types: HTTP_HEADER -- Name of the HTTP header whose value is taken as the key value. HTTP_COOKIE -- Name of the HTTP cookie whose value is taken as the key value. */
	// +optional
	EnforceOnKeyName *string `json:"enforceOnKeyName,omitempty"`

	/* Determines the key to enforce the rate_limit_threshold on. */
	// +optional
	EnforceOnKeyType *string `json:"enforceOnKeyType,omitempty"`
}

type SecuritypolicyExceedRedirectOptions struct {
	/* Target for the redirect action. This is required if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA. */
	// +optional
	Target *string `json:"target,omitempty"`

	/* Type of the redirect action. */
	Type string `json:"type"`
}

type SecuritypolicyExclusion struct {
	/* Request cookie whose value will be excluded from inspection during preconfigured WAF evaluation. */
	// +optional
	RequestCookie []SecuritypolicyRequestCookie `json:"requestCookie,omitempty"`

	/* Request header whose value will be excluded from inspection during preconfigured WAF evaluation. */
	// +optional
	RequestHeader []SecuritypolicyRequestHeader `json:"requestHeader,omitempty"`

	/* Request query parameter whose value will be excluded from inspection during preconfigured WAF evaluation.  Note that the parameter can be in the query string or in the POST body. */
	// +optional
	RequestQueryParam []SecuritypolicyRequestQueryParam `json:"requestQueryParam,omitempty"`

	/* Request URI from the request line to be excluded from inspection during preconfigured WAF evaluation. When specifying this field, the query or fragment part should be excluded. */
	// +optional
	RequestUri []SecuritypolicyRequestUri `json:"requestUri,omitempty"`

	/* A list of target rule IDs under the WAF rule set to apply the preconfigured WAF exclusion. If omitted, it refers to all the rule IDs under the WAF rule set. */
	// +optional
	TargetRuleIds []string `json:"targetRuleIds,omitempty"`

	/* Target WAF rule set to apply the preconfigured WAF exclusion. */
	TargetRuleSet string `json:"targetRuleSet"`
}

type SecuritypolicyExpr struct {
	/* Textual representation of an expression in Common Expression Language syntax. The application context of the containing message determines which well-known feature set of CEL is supported. */
	Expression string `json:"expression"`
}

type SecuritypolicyHeaderAction struct {
	/* The list of request headers to add or overwrite if they're already present. */
	RequestHeadersToAdds []SecuritypolicyRequestHeadersToAdds `json:"requestHeadersToAdds"`
}

type SecuritypolicyJsonCustomConfig struct {
	/* A list of custom Content-Type header values to apply the JSON parsing. */
	ContentTypes []string `json:"contentTypes"`
}

type SecuritypolicyLayer7DdosDefenseConfig struct {
	/* If set to true, enables CAAP for L7 DDoS detection. */
	// +optional
	Enable *bool `json:"enable,omitempty"`

	/* Rule visibility. Supported values include: "STANDARD", "PREMIUM". */
	// +optional
	RuleVisibility *string `json:"ruleVisibility,omitempty"`
}

type SecuritypolicyMatch struct {
	/* The configuration options available when specifying versioned_expr. This field must be specified if versioned_expr is specified and cannot be specified if versioned_expr is not specified. */
	// +optional
	Config *SecuritypolicyConfig `json:"config,omitempty"`

	/* User defined CEVAL expression. A CEVAL expression is used to specify match criteria such as origin.ip, source.region_code and contents in the request header. */
	// +optional
	Expr *SecuritypolicyExpr `json:"expr,omitempty"`

	/* Predefined rule expression. If this field is specified, config must also be specified. Available options:   SRC_IPS_V1: Must specify the corresponding src_ip_ranges field in config. */
	// +optional
	VersionedExpr *string `json:"versionedExpr,omitempty"`
}

type SecuritypolicyPreconfiguredWafConfig struct {
	/* An exclusion to apply during preconfigured WAF evaluation. */
	// +optional
	Exclusion []SecuritypolicyExclusion `json:"exclusion,omitempty"`
}

type SecuritypolicyRateLimitOptions struct {
	/* Can only be specified if the action for the rule is "rate_based_ban". If specified, determines the time (in seconds) the traffic will continue to be banned by the rate limit after the rate falls below the threshold. */
	// +optional
	BanDurationSec *int `json:"banDurationSec,omitempty"`

	/* Can only be specified if the action for the rule is "rate_based_ban". If specified, the key will be banned for the configured 'banDurationSec' when the number of requests that exceed the 'rateLimitThreshold' also exceed this 'banThreshold'. */
	// +optional
	BanThreshold *SecuritypolicyBanThreshold `json:"banThreshold,omitempty"`

	/* Action to take for requests that are under the configured rate limit threshold. Valid option is "allow" only. */
	ConformAction string `json:"conformAction"`

	/* Determines the key to enforce the rateLimitThreshold on. */
	// +optional
	EnforceOnKey *string `json:"enforceOnKey,omitempty"`

	/* Immutable. Enforce On Key Config of this security policy. */
	// +optional
	EnforceOnKeyConfigs []SecuritypolicyEnforceOnKeyConfigs `json:"enforceOnKeyConfigs,omitempty"`

	/* Rate limit key name applicable only for the following key types: HTTP_HEADER -- Name of the HTTP header whose value is taken as the key value. HTTP_COOKIE -- Name of the HTTP cookie whose value is taken as the key value. */
	// +optional
	EnforceOnKeyName *string `json:"enforceOnKeyName,omitempty"`

	/* Action to take for requests that are above the configured rate limit threshold, to either deny with a specified HTTP response code, or redirect to a different endpoint. Valid options are "deny()" where valid values for status are 403, 404, 429, and 502, and "redirect" where the redirect parameters come from exceedRedirectOptions below. */
	ExceedAction string `json:"exceedAction"`

	/* Parameters defining the redirect action that is used as the exceed action. Cannot be specified if the exceed action is not redirect. */
	// +optional
	ExceedRedirectOptions *SecuritypolicyExceedRedirectOptions `json:"exceedRedirectOptions,omitempty"`

	/* Threshold at which to begin ratelimiting. */
	RateLimitThreshold SecuritypolicyRateLimitThreshold `json:"rateLimitThreshold"`
}

type SecuritypolicyRateLimitThreshold struct {
	/* Number of HTTP(S) requests for calculating the threshold. */
	Count int `json:"count"`

	/* Interval over which the threshold is computed. */
	IntervalSec int `json:"intervalSec"`
}

type SecuritypolicyRecaptchaOptionsConfig struct {
	/* A field to supply a reCAPTCHA site key to be used for all the rules
	using the redirect action with the type of GOOGLE_RECAPTCHA under
	the security policy. The specified site key needs to be created from
	the reCAPTCHA API. The user is responsible for the validity of the
	specified site key. If not specified, a Google-managed site key is
	used. */
	RedirectSiteKeyRef v1alpha1.ResourceRef `json:"redirectSiteKeyRef"`
}

type SecuritypolicyRedirectOptions struct {
	/* Target for the redirect action. This is required if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA. */
	// +optional
	Target *string `json:"target,omitempty"`

	/* Type of the redirect action. Available options: EXTERNAL_302: Must specify the corresponding target field in config. GOOGLE_RECAPTCHA: Cannot specify target field in config. */
	Type string `json:"type"`
}

type SecuritypolicyRequestCookie struct {
	/* You can specify an exact match or a partial match by using a field operator and a field value. Available options: EQUALS: The operator matches if the field value equals the specified value. STARTS_WITH: The operator matches if the field value starts with the specified value. ENDS_WITH: The operator matches if the field value ends with the specified value. CONTAINS: The operator matches if the field value contains the specified value. EQUALS_ANY: The operator matches if the field value is any value. */
	Operator string `json:"operator"`

	/* A request field matching the specified value will be excluded from inspection during preconfigured WAF evaluation. The field value must be given if the field operator is not EQUALS_ANY, and cannot be given if the field operator is EQUALS_ANY. */
	// +optional
	Value *string `json:"value,omitempty"`
}

type SecuritypolicyRequestHeader struct {
	/* You can specify an exact match or a partial match by using a field operator and a field value. Available options: EQUALS: The operator matches if the field value equals the specified value. STARTS_WITH: The operator matches if the field value starts with the specified value. ENDS_WITH: The operator matches if the field value ends with the specified value. CONTAINS: The operator matches if the field value contains the specified value. EQUALS_ANY: The operator matches if the field value is any value. */
	Operator string `json:"operator"`

	/* A request field matching the specified value will be excluded from inspection during preconfigured WAF evaluation. The field value must be given if the field operator is not EQUALS_ANY, and cannot be given if the field operator is EQUALS_ANY. */
	// +optional
	Value *string `json:"value,omitempty"`
}

type SecuritypolicyRequestHeadersToAdds struct {
	/* The name of the header to set. */
	HeaderName string `json:"headerName"`

	/* The value to set the named header to. */
	// +optional
	HeaderValue *string `json:"headerValue,omitempty"`
}

type SecuritypolicyRequestQueryParam struct {
	/* You can specify an exact match or a partial match by using a field operator and a field value. Available options: EQUALS: The operator matches if the field value equals the specified value. STARTS_WITH: The operator matches if the field value starts with the specified value. ENDS_WITH: The operator matches if the field value ends with the specified value. CONTAINS: The operator matches if the field value contains the specified value. EQUALS_ANY: The operator matches if the field value is any value. */
	Operator string `json:"operator"`

	/* A request field matching the specified value will be excluded from inspection during preconfigured WAF evaluation. The field value must be given if the field operator is not EQUALS_ANY, and cannot be given if the field operator is EQUALS_ANY. */
	// +optional
	Value *string `json:"value,omitempty"`
}

type SecuritypolicyRequestUri struct {
	/* You can specify an exact match or a partial match by using a field operator and a field value. Available options: EQUALS: The operator matches if the field value equals the specified value. STARTS_WITH: The operator matches if the field value starts with the specified value. ENDS_WITH: The operator matches if the field value ends with the specified value. CONTAINS: The operator matches if the field value contains the specified value. EQUALS_ANY: The operator matches if the field value is any value. */
	Operator string `json:"operator"`

	/* A request field matching the specified value will be excluded from inspection during preconfigured WAF evaluation. The field value must be given if the field operator is not EQUALS_ANY, and cannot be given if the field operator is EQUALS_ANY. */
	// +optional
	Value *string `json:"value,omitempty"`
}

type SecuritypolicyRule struct {
	/* Action to take when match matches the request. */
	Action string `json:"action"`

	/* An optional description of this rule. Max size is 64. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Additional actions that are performed on headers. */
	// +optional
	HeaderAction *SecuritypolicyHeaderAction `json:"headerAction,omitempty"`

	/* A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding action is enforced. */
	Match SecuritypolicyMatch `json:"match"`

	/* Preconfigured WAF configuration to be applied for the rule. If the rule does not evaluate preconfigured WAF rules, i.e., if evaluatePreconfiguredWaf() is not used, this field will have no effect. */
	// +optional
	PreconfiguredWafConfig *SecuritypolicyPreconfiguredWafConfig `json:"preconfiguredWafConfig,omitempty"`

	/* When set to true, the action specified above is not enforced. Stackdriver logs for requests that trigger a preview action are annotated as such. */
	// +optional
	Preview *bool `json:"preview,omitempty"`

	/* An unique positive integer indicating the priority of evaluation for a rule. Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order. */
	Priority int `json:"priority"`

	/* Rate limit threshold for this security policy. Must be specified if the action is "rate_based_ban" or "throttle". Cannot be specified for any other actions. */
	// +optional
	RateLimitOptions *SecuritypolicyRateLimitOptions `json:"rateLimitOptions,omitempty"`

	/* Parameters defining the redirect action. Cannot be specified for any other actions. */
	// +optional
	RedirectOptions *SecuritypolicyRedirectOptions `json:"redirectOptions,omitempty"`
}

type ComputeSecurityPolicySpec struct {
	/* Adaptive Protection Config of this security policy. */
	// +optional
	AdaptiveProtectionConfig *SecuritypolicyAdaptiveProtectionConfig `json:"adaptiveProtectionConfig,omitempty"`

	/* Advanced Options Config of this security policy. */
	// +optional
	AdvancedOptionsConfig *SecuritypolicyAdvancedOptionsConfig `json:"advancedOptionsConfig,omitempty"`

	/* An optional description of this security policy. Max size is 2048. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* reCAPTCHA configuration options to be applied for the security policy. */
	// +optional
	RecaptchaOptionsConfig *SecuritypolicyRecaptchaOptionsConfig `json:"recaptchaOptionsConfig,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The set of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added. */
	// +optional
	Rule []SecuritypolicyRule `json:"rule,omitempty"`

	/* The type indicates the intended use of the security policy. CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. */
	// +optional
	Type *string `json:"type,omitempty"`
}

type ComputeSecurityPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeSecurityPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Fingerprint of this resource. */
	// +optional
	Fingerprint *string `json:"fingerprint,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The URI of the created resource. */
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeSecurityPolicy is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeSecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeSecurityPolicySpec   `json:"spec,omitempty"`
	Status ComputeSecurityPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeSecurityPolicyList contains a list of ComputeSecurityPolicy
type ComputeSecurityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeSecurityPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeSecurityPolicy{}, &ComputeSecurityPolicyList{})
}
