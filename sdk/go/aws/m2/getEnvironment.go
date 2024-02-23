// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package m2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a runtime environment that can run migrated mainframe applications.
func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentResult
	err := ctx.Invoke("aws-native:m2:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	// The Amazon Resource Name (ARN) of the runtime environment.
	EnvironmentArn string `pulumi:"environmentArn"`
}

type LookupEnvironmentResult struct {
	// The version of the runtime engine for the environment.
	EngineVersion *string `pulumi:"engineVersion"`
	// The Amazon Resource Name (ARN) of the runtime environment.
	EnvironmentArn *string `pulumi:"environmentArn"`
	// The unique identifier of the environment.
	EnvironmentId          *string                            `pulumi:"environmentId"`
	HighAvailabilityConfig *EnvironmentHighAvailabilityConfig `pulumi:"highAvailabilityConfig"`
	// The type of instance underlying the environment.
	InstanceType *string `pulumi:"instanceType"`
	// Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
	// Tags associated to this environment.
	Tags map[string]string `pulumi:"tags"`
}

func LookupEnvironmentOutput(ctx *pulumi.Context, args LookupEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentResult, error) {
			args := v.(LookupEnvironmentArgs)
			r, err := LookupEnvironment(ctx, &args, opts...)
			var s LookupEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentResultOutput)
}

type LookupEnvironmentOutputArgs struct {
	// The Amazon Resource Name (ARN) of the runtime environment.
	EnvironmentArn pulumi.StringInput `pulumi:"environmentArn"`
}

func (LookupEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentArgs)(nil)).Elem()
}

type LookupEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentResult)(nil)).Elem()
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutput() LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutputWithContext(ctx context.Context) LookupEnvironmentResultOutput {
	return o
}

// The version of the runtime engine for the environment.
func (o LookupEnvironmentResultOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the runtime environment.
func (o LookupEnvironmentResultOutput) EnvironmentArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.EnvironmentArn }).(pulumi.StringPtrOutput)
}

// The unique identifier of the environment.
func (o LookupEnvironmentResultOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) HighAvailabilityConfig() EnvironmentHighAvailabilityConfigPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *EnvironmentHighAvailabilityConfig { return v.HighAvailabilityConfig }).(EnvironmentHighAvailabilityConfigPtrOutput)
}

// The type of instance underlying the environment.
func (o LookupEnvironmentResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

// Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.
func (o LookupEnvironmentResultOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

// Tags associated to this environment.
func (o LookupEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentResultOutput{})
}
