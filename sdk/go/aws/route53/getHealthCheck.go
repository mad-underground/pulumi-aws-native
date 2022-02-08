// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Route53::HealthCheck.
func LookupHealthCheck(ctx *pulumi.Context, args *LookupHealthCheckArgs, opts ...pulumi.InvokeOption) (*LookupHealthCheckResult, error) {
	var rv LookupHealthCheckResult
	err := ctx.Invoke("aws-native:route53:getHealthCheck", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHealthCheckArgs struct {
	HealthCheckId string `pulumi:"healthCheckId"`
}

type LookupHealthCheckResult struct {
	// A complex type that contains information about the health check.
	HealthCheckConfig *HealthCheckConfigProperties `pulumi:"healthCheckConfig"`
	HealthCheckId     *string                      `pulumi:"healthCheckId"`
	// An array of key-value pairs to apply to this resource.
	HealthCheckTags []HealthCheckTag `pulumi:"healthCheckTags"`
}

func LookupHealthCheckOutput(ctx *pulumi.Context, args LookupHealthCheckOutputArgs, opts ...pulumi.InvokeOption) LookupHealthCheckResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHealthCheckResult, error) {
			args := v.(LookupHealthCheckArgs)
			r, err := LookupHealthCheck(ctx, &args, opts...)
			return *r, err
		}).(LookupHealthCheckResultOutput)
}

type LookupHealthCheckOutputArgs struct {
	HealthCheckId pulumi.StringInput `pulumi:"healthCheckId"`
}

func (LookupHealthCheckOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHealthCheckArgs)(nil)).Elem()
}

type LookupHealthCheckResultOutput struct{ *pulumi.OutputState }

func (LookupHealthCheckResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHealthCheckResult)(nil)).Elem()
}

func (o LookupHealthCheckResultOutput) ToLookupHealthCheckResultOutput() LookupHealthCheckResultOutput {
	return o
}

func (o LookupHealthCheckResultOutput) ToLookupHealthCheckResultOutputWithContext(ctx context.Context) LookupHealthCheckResultOutput {
	return o
}

// A complex type that contains information about the health check.
func (o LookupHealthCheckResultOutput) HealthCheckConfig() HealthCheckConfigPropertiesPtrOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) *HealthCheckConfigProperties { return v.HealthCheckConfig }).(HealthCheckConfigPropertiesPtrOutput)
}

func (o LookupHealthCheckResultOutput) HealthCheckId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) *string { return v.HealthCheckId }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupHealthCheckResultOutput) HealthCheckTags() HealthCheckTagArrayOutput {
	return o.ApplyT(func(v LookupHealthCheckResult) []HealthCheckTag { return v.HealthCheckTags }).(HealthCheckTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHealthCheckResultOutput{})
}
