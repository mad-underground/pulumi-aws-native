// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “AWS::ApiGateway::Stage“ resource creates a stage for a deployment.
func LookupStage(ctx *pulumi.Context, args *LookupStageArgs, opts ...pulumi.InvokeOption) (*LookupStageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStageResult
	err := ctx.Invoke("aws-native:apigateway:getStage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStageArgs struct {
	// The string identifier of the associated RestApi.
	RestApiId string `pulumi:"restApiId"`
	// The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway. Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.
	StageName string `pulumi:"stageName"`
}

type LookupStageResult struct {
	// Access log settings, including the access log format and access log destination ARN.
	AccessLogSetting *StageAccessLogSetting `pulumi:"accessLogSetting"`
	// Specifies whether a cache cluster is enabled for the stage.
	CacheClusterEnabled *bool `pulumi:"cacheClusterEnabled"`
	// The stage's cache capacity in GB. For more information about choosing a cache size, see [Enabling API caching to enhance responsiveness](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html).
	CacheClusterSize *string `pulumi:"cacheClusterSize"`
	// Settings for the canary deployment in this stage.
	CanarySetting *StageCanarySetting `pulumi:"canarySetting"`
	// The identifier of a client certificate for an API stage.
	ClientCertificateId *string `pulumi:"clientCertificateId"`
	// The identifier of the Deployment that the stage points to.
	DeploymentId *string `pulumi:"deploymentId"`
	// The stage's description.
	Description *string `pulumi:"description"`
	// The version of the associated API documentation.
	DocumentationVersion *string `pulumi:"documentationVersion"`
	// A map that defines the method settings for a Stage resource. Keys (designated as ``/{method_setting_key`` below) are method paths defined as ``{resource_path}/{http_method}`` for an individual method override, or ``/\*/\*`` for overriding all methods in the stage.
	MethodSettings []StageMethodSetting `pulumi:"methodSettings"`
	// The collection of tags. Each tag element is associated with a given resource.
	Tags []StageTag `pulumi:"tags"`
	// Specifies whether active tracing with X-ray is enabled for the Stage.
	TracingEnabled *bool `pulumi:"tracingEnabled"`
	// A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value. Variable names are limited to alphanumeric characters. Values must match the following regular expression: ``[A-Za-z0-9-._~:/?#&=,]+``.
	Variables interface{} `pulumi:"variables"`
}

func LookupStageOutput(ctx *pulumi.Context, args LookupStageOutputArgs, opts ...pulumi.InvokeOption) LookupStageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStageResult, error) {
			args := v.(LookupStageArgs)
			r, err := LookupStage(ctx, &args, opts...)
			var s LookupStageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStageResultOutput)
}

type LookupStageOutputArgs struct {
	// The string identifier of the associated RestApi.
	RestApiId pulumi.StringInput `pulumi:"restApiId"`
	// The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway. Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.
	StageName pulumi.StringInput `pulumi:"stageName"`
}

func (LookupStageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStageArgs)(nil)).Elem()
}

type LookupStageResultOutput struct{ *pulumi.OutputState }

func (LookupStageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStageResult)(nil)).Elem()
}

func (o LookupStageResultOutput) ToLookupStageResultOutput() LookupStageResultOutput {
	return o
}

func (o LookupStageResultOutput) ToLookupStageResultOutputWithContext(ctx context.Context) LookupStageResultOutput {
	return o
}

// Access log settings, including the access log format and access log destination ARN.
func (o LookupStageResultOutput) AccessLogSetting() StageAccessLogSettingPtrOutput {
	return o.ApplyT(func(v LookupStageResult) *StageAccessLogSetting { return v.AccessLogSetting }).(StageAccessLogSettingPtrOutput)
}

// Specifies whether a cache cluster is enabled for the stage.
func (o LookupStageResultOutput) CacheClusterEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStageResult) *bool { return v.CacheClusterEnabled }).(pulumi.BoolPtrOutput)
}

// The stage's cache capacity in GB. For more information about choosing a cache size, see [Enabling API caching to enhance responsiveness](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html).
func (o LookupStageResultOutput) CacheClusterSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStageResult) *string { return v.CacheClusterSize }).(pulumi.StringPtrOutput)
}

// Settings for the canary deployment in this stage.
func (o LookupStageResultOutput) CanarySetting() StageCanarySettingPtrOutput {
	return o.ApplyT(func(v LookupStageResult) *StageCanarySetting { return v.CanarySetting }).(StageCanarySettingPtrOutput)
}

// The identifier of a client certificate for an API stage.
func (o LookupStageResultOutput) ClientCertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStageResult) *string { return v.ClientCertificateId }).(pulumi.StringPtrOutput)
}

// The identifier of the Deployment that the stage points to.
func (o LookupStageResultOutput) DeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStageResult) *string { return v.DeploymentId }).(pulumi.StringPtrOutput)
}

// The stage's description.
func (o LookupStageResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStageResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The version of the associated API documentation.
func (o LookupStageResultOutput) DocumentationVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStageResult) *string { return v.DocumentationVersion }).(pulumi.StringPtrOutput)
}

// A map that defines the method settings for a Stage resource. Keys (designated as “/{method_setting_key“ below) are method paths defined as “{resource_path}/{http_method}“ for an individual method override, or “/\*/\*“ for overriding all methods in the stage.
func (o LookupStageResultOutput) MethodSettings() StageMethodSettingArrayOutput {
	return o.ApplyT(func(v LookupStageResult) []StageMethodSetting { return v.MethodSettings }).(StageMethodSettingArrayOutput)
}

// The collection of tags. Each tag element is associated with a given resource.
func (o LookupStageResultOutput) Tags() StageTagArrayOutput {
	return o.ApplyT(func(v LookupStageResult) []StageTag { return v.Tags }).(StageTagArrayOutput)
}

// Specifies whether active tracing with X-ray is enabled for the Stage.
func (o LookupStageResultOutput) TracingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStageResult) *bool { return v.TracingEnabled }).(pulumi.BoolPtrOutput)
}

// A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value. Variable names are limited to alphanumeric characters. Values must match the following regular expression: “[A-Za-z0-9-._~:/?#&=,]+“.
func (o LookupStageResultOutput) Variables() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupStageResult) interface{} { return v.Variables }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStageResultOutput{})
}
