// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An example resource schema demonstrating some basic constructs and validation rules.
func LookupFunctionConfiguration(ctx *pulumi.Context, args *LookupFunctionConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupFunctionConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFunctionConfigurationResult
	err := ctx.Invoke("aws-native:appsync:getFunctionConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFunctionConfigurationArgs struct {
	// The ARN for the function generated by the service
	FunctionArn string `pulumi:"functionArn"`
}

type LookupFunctionConfigurationResult struct {
	// The resolver code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
	Code *string `pulumi:"code"`
	// The name of data source this function will attach.
	DataSourceName *string `pulumi:"dataSourceName"`
	// The function description.
	Description *string `pulumi:"description"`
	// The ARN for the function generated by the service
	FunctionArn *string `pulumi:"functionArn"`
	// The unique identifier for the function generated by the service
	FunctionId *string `pulumi:"functionId"`
	// The version of the request mapping template. Currently, only the 2018-05-29 version of the template is supported.
	FunctionVersion *string `pulumi:"functionVersion"`
	// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a BatchInvoke operation.
	MaxBatchSize *int `pulumi:"maxBatchSize"`
	// The name of the function.
	Name *string `pulumi:"name"`
	// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
	RequestMappingTemplate *string `pulumi:"requestMappingTemplate"`
	// The Function response mapping template.
	ResponseMappingTemplate *string `pulumi:"responseMappingTemplate"`
	// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.
	Runtime *FunctionConfigurationAppSyncRuntime `pulumi:"runtime"`
	// Describes a Sync configuration for a resolver. Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is invoked.
	SyncConfig *FunctionConfigurationSyncConfig `pulumi:"syncConfig"`
}

func LookupFunctionConfigurationOutput(ctx *pulumi.Context, args LookupFunctionConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupFunctionConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFunctionConfigurationResult, error) {
			args := v.(LookupFunctionConfigurationArgs)
			r, err := LookupFunctionConfiguration(ctx, &args, opts...)
			var s LookupFunctionConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFunctionConfigurationResultOutput)
}

type LookupFunctionConfigurationOutputArgs struct {
	// The ARN for the function generated by the service
	FunctionArn pulumi.StringInput `pulumi:"functionArn"`
}

func (LookupFunctionConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionConfigurationArgs)(nil)).Elem()
}

type LookupFunctionConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupFunctionConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionConfigurationResult)(nil)).Elem()
}

func (o LookupFunctionConfigurationResultOutput) ToLookupFunctionConfigurationResultOutput() LookupFunctionConfigurationResultOutput {
	return o
}

func (o LookupFunctionConfigurationResultOutput) ToLookupFunctionConfigurationResultOutputWithContext(ctx context.Context) LookupFunctionConfigurationResultOutput {
	return o
}

func (o LookupFunctionConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFunctionConfigurationResult] {
	return pulumix.Output[LookupFunctionConfigurationResult]{
		OutputState: o.OutputState,
	}
}

// The resolver code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
func (o LookupFunctionConfigurationResultOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.Code }).(pulumi.StringPtrOutput)
}

// The name of data source this function will attach.
func (o LookupFunctionConfigurationResultOutput) DataSourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.DataSourceName }).(pulumi.StringPtrOutput)
}

// The function description.
func (o LookupFunctionConfigurationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The ARN for the function generated by the service
func (o LookupFunctionConfigurationResultOutput) FunctionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.FunctionArn }).(pulumi.StringPtrOutput)
}

// The unique identifier for the function generated by the service
func (o LookupFunctionConfigurationResultOutput) FunctionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.FunctionId }).(pulumi.StringPtrOutput)
}

// The version of the request mapping template. Currently, only the 2018-05-29 version of the template is supported.
func (o LookupFunctionConfigurationResultOutput) FunctionVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.FunctionVersion }).(pulumi.StringPtrOutput)
}

// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a BatchInvoke operation.
func (o LookupFunctionConfigurationResultOutput) MaxBatchSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *int { return v.MaxBatchSize }).(pulumi.IntPtrOutput)
}

// The name of the function.
func (o LookupFunctionConfigurationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
func (o LookupFunctionConfigurationResultOutput) RequestMappingTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.RequestMappingTemplate }).(pulumi.StringPtrOutput)
}

// The Function response mapping template.
func (o LookupFunctionConfigurationResultOutput) ResponseMappingTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.ResponseMappingTemplate }).(pulumi.StringPtrOutput)
}

// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.
func (o LookupFunctionConfigurationResultOutput) Runtime() FunctionConfigurationAppSyncRuntimePtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *FunctionConfigurationAppSyncRuntime { return v.Runtime }).(FunctionConfigurationAppSyncRuntimePtrOutput)
}

// Describes a Sync configuration for a resolver. Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is invoked.
func (o LookupFunctionConfigurationResultOutput) SyncConfig() FunctionConfigurationSyncConfigPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *FunctionConfigurationSyncConfig { return v.SyncConfig }).(FunctionConfigurationSyncConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFunctionConfigurationResultOutput{})
}
