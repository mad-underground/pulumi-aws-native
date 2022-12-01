// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElasticLoadBalancingV2::TargetGroup
func LookupTargetGroup(ctx *pulumi.Context, args *LookupTargetGroupArgs, opts ...pulumi.InvokeOption) (*LookupTargetGroupResult, error) {
	var rv LookupTargetGroupResult
	err := ctx.Invoke("aws-native:elasticloadbalancingv2:getTargetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTargetGroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupTargetGroupResult struct {
	HealthCheckEnabled         *bool                          `pulumi:"healthCheckEnabled"`
	HealthCheckIntervalSeconds *int                           `pulumi:"healthCheckIntervalSeconds"`
	HealthCheckPath            *string                        `pulumi:"healthCheckPath"`
	HealthCheckPort            *string                        `pulumi:"healthCheckPort"`
	HealthCheckProtocol        *string                        `pulumi:"healthCheckProtocol"`
	HealthCheckTimeoutSeconds  *int                           `pulumi:"healthCheckTimeoutSeconds"`
	HealthyThresholdCount      *int                           `pulumi:"healthyThresholdCount"`
	Id                         *string                        `pulumi:"id"`
	LoadBalancerArns           []string                       `pulumi:"loadBalancerArns"`
	Matcher                    *TargetGroupMatcher            `pulumi:"matcher"`
	Tags                       []TargetGroupTag               `pulumi:"tags"`
	TargetGroupArn             *string                        `pulumi:"targetGroupArn"`
	TargetGroupAttributes      []TargetGroupAttribute         `pulumi:"targetGroupAttributes"`
	TargetGroupFullName        *string                        `pulumi:"targetGroupFullName"`
	TargetGroupName            *string                        `pulumi:"targetGroupName"`
	Targets                    []TargetGroupTargetDescription `pulumi:"targets"`
	UnhealthyThresholdCount    *int                           `pulumi:"unhealthyThresholdCount"`
}

func LookupTargetGroupOutput(ctx *pulumi.Context, args LookupTargetGroupOutputArgs, opts ...pulumi.InvokeOption) LookupTargetGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTargetGroupResult, error) {
			args := v.(LookupTargetGroupArgs)
			r, err := LookupTargetGroup(ctx, &args, opts...)
			var s LookupTargetGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTargetGroupResultOutput)
}

type LookupTargetGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupTargetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetGroupArgs)(nil)).Elem()
}

type LookupTargetGroupResultOutput struct{ *pulumi.OutputState }

func (LookupTargetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetGroupResult)(nil)).Elem()
}

func (o LookupTargetGroupResultOutput) ToLookupTargetGroupResultOutput() LookupTargetGroupResultOutput {
	return o
}

func (o LookupTargetGroupResultOutput) ToLookupTargetGroupResultOutputWithContext(ctx context.Context) LookupTargetGroupResultOutput {
	return o
}

func (o LookupTargetGroupResultOutput) HealthCheckEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *bool { return v.HealthCheckEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupTargetGroupResultOutput) HealthCheckIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *int { return v.HealthCheckIntervalSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupTargetGroupResultOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.HealthCheckPath }).(pulumi.StringPtrOutput)
}

func (o LookupTargetGroupResultOutput) HealthCheckPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.HealthCheckPort }).(pulumi.StringPtrOutput)
}

func (o LookupTargetGroupResultOutput) HealthCheckProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.HealthCheckProtocol }).(pulumi.StringPtrOutput)
}

func (o LookupTargetGroupResultOutput) HealthCheckTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *int { return v.HealthCheckTimeoutSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupTargetGroupResultOutput) HealthyThresholdCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *int { return v.HealthyThresholdCount }).(pulumi.IntPtrOutput)
}

func (o LookupTargetGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupTargetGroupResultOutput) LoadBalancerArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) []string { return v.LoadBalancerArns }).(pulumi.StringArrayOutput)
}

func (o LookupTargetGroupResultOutput) Matcher() TargetGroupMatcherPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *TargetGroupMatcher { return v.Matcher }).(TargetGroupMatcherPtrOutput)
}

func (o LookupTargetGroupResultOutput) Tags() TargetGroupTagArrayOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) []TargetGroupTag { return v.Tags }).(TargetGroupTagArrayOutput)
}

func (o LookupTargetGroupResultOutput) TargetGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.TargetGroupArn }).(pulumi.StringPtrOutput)
}

func (o LookupTargetGroupResultOutput) TargetGroupAttributes() TargetGroupAttributeArrayOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) []TargetGroupAttribute { return v.TargetGroupAttributes }).(TargetGroupAttributeArrayOutput)
}

func (o LookupTargetGroupResultOutput) TargetGroupFullName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.TargetGroupFullName }).(pulumi.StringPtrOutput)
}

func (o LookupTargetGroupResultOutput) TargetGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.TargetGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupTargetGroupResultOutput) Targets() TargetGroupTargetDescriptionArrayOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) []TargetGroupTargetDescription { return v.Targets }).(TargetGroupTargetDescriptionArrayOutput)
}

func (o LookupTargetGroupResultOutput) UnhealthyThresholdCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *int { return v.UnhealthyThresholdCount }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTargetGroupResultOutput{})
}
