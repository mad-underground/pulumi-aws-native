// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apprunner

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::AppRunner::Service resource specifies an AppRunner Service.
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceResult
	err := ctx.Invoke("aws-native:apprunner:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	// The Amazon Resource Name (ARN) of the AppRunner Service.
	ServiceArn string `pulumi:"serviceArn"`
}

type LookupServiceResult struct {
	HealthCheckConfiguration   *ServiceHealthCheckConfiguration   `pulumi:"healthCheckConfiguration"`
	InstanceConfiguration      *ServiceInstanceConfiguration      `pulumi:"instanceConfiguration"`
	NetworkConfiguration       *ServiceNetworkConfiguration       `pulumi:"networkConfiguration"`
	ObservabilityConfiguration *ServiceObservabilityConfiguration `pulumi:"observabilityConfiguration"`
	// The Amazon Resource Name (ARN) of the AppRunner Service.
	ServiceArn *string `pulumi:"serviceArn"`
	// The AppRunner Service Id
	ServiceId *string `pulumi:"serviceId"`
	// The Service Url of the AppRunner Service.
	ServiceUrl          *string                     `pulumi:"serviceUrl"`
	SourceConfiguration *ServiceSourceConfiguration `pulumi:"sourceConfiguration"`
	// AppRunner Service status.
	Status *string `pulumi:"status"`
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceResult, error) {
			args := v.(LookupServiceArgs)
			r, err := LookupService(ctx, &args, opts...)
			var s LookupServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceResultOutput)
}

type LookupServiceOutputArgs struct {
	// The Amazon Resource Name (ARN) of the AppRunner Service.
	ServiceArn pulumi.StringInput `pulumi:"serviceArn"`
}

func (LookupServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}

type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) HealthCheckConfiguration() ServiceHealthCheckConfigurationPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *ServiceHealthCheckConfiguration { return v.HealthCheckConfiguration }).(ServiceHealthCheckConfigurationPtrOutput)
}

func (o LookupServiceResultOutput) InstanceConfiguration() ServiceInstanceConfigurationPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *ServiceInstanceConfiguration { return v.InstanceConfiguration }).(ServiceInstanceConfigurationPtrOutput)
}

func (o LookupServiceResultOutput) NetworkConfiguration() ServiceNetworkConfigurationPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *ServiceNetworkConfiguration { return v.NetworkConfiguration }).(ServiceNetworkConfigurationPtrOutput)
}

func (o LookupServiceResultOutput) ObservabilityConfiguration() ServiceObservabilityConfigurationPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *ServiceObservabilityConfiguration { return v.ObservabilityConfiguration }).(ServiceObservabilityConfigurationPtrOutput)
}

// The Amazon Resource Name (ARN) of the AppRunner Service.
func (o LookupServiceResultOutput) ServiceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.ServiceArn }).(pulumi.StringPtrOutput)
}

// The AppRunner Service Id
func (o LookupServiceResultOutput) ServiceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.ServiceId }).(pulumi.StringPtrOutput)
}

// The Service Url of the AppRunner Service.
func (o LookupServiceResultOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) SourceConfiguration() ServiceSourceConfigurationPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *ServiceSourceConfiguration { return v.SourceConfiguration }).(ServiceSourceConfigurationPtrOutput)
}

// AppRunner Service status.
func (o LookupServiceResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}
