// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ServiceDiscovery::Service
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("aws-native:servicediscovery:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	Id string `pulumi:"id"`
}

type LookupServiceResult struct {
	Arn               *string                   `pulumi:"arn"`
	Description       *string                   `pulumi:"description"`
	DnsConfig         *ServiceDnsConfig         `pulumi:"dnsConfig"`
	HealthCheckConfig *ServiceHealthCheckConfig `pulumi:"healthCheckConfig"`
	Id                *string                   `pulumi:"id"`
	Tags              []ServiceTag              `pulumi:"tags"`
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceResult, error) {
			args := v.(LookupServiceArgs)
			r, err := LookupService(ctx, &args, opts...)
			return *r, err
		}).(LookupServiceResultOutput)
}

type LookupServiceOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
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

func (o LookupServiceResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) DnsConfig() ServiceDnsConfigPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *ServiceDnsConfig { return v.DnsConfig }).(ServiceDnsConfigPtrOutput)
}

func (o LookupServiceResultOutput) HealthCheckConfig() ServiceHealthCheckConfigPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *ServiceHealthCheckConfig { return v.HealthCheckConfig }).(ServiceHealthCheckConfigPtrOutput)
}

func (o LookupServiceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) Tags() ServiceTagArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []ServiceTag { return v.Tags }).(ServiceTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}