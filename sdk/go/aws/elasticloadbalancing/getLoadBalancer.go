// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElasticLoadBalancing::LoadBalancer
func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("aws-native:elasticloadbalancing:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoadBalancerArgs struct {
	Id string `pulumi:"id"`
}

type LookupLoadBalancerResult struct {
	AccessLoggingPolicy           *LoadBalancerAccessLoggingPolicy        `pulumi:"accessLoggingPolicy"`
	AppCookieStickinessPolicy     []LoadBalancerAppCookieStickinessPolicy `pulumi:"appCookieStickinessPolicy"`
	AvailabilityZones             []string                                `pulumi:"availabilityZones"`
	CanonicalHostedZoneName       *string                                 `pulumi:"canonicalHostedZoneName"`
	CanonicalHostedZoneNameID     *string                                 `pulumi:"canonicalHostedZoneNameID"`
	ConnectionDrainingPolicy      *LoadBalancerConnectionDrainingPolicy   `pulumi:"connectionDrainingPolicy"`
	ConnectionSettings            *LoadBalancerConnectionSettings         `pulumi:"connectionSettings"`
	CrossZone                     *bool                                   `pulumi:"crossZone"`
	DNSName                       *string                                 `pulumi:"dNSName"`
	HealthCheck                   *LoadBalancerHealthCheck                `pulumi:"healthCheck"`
	Id                            *string                                 `pulumi:"id"`
	Instances                     []string                                `pulumi:"instances"`
	LBCookieStickinessPolicy      []LoadBalancerLBCookieStickinessPolicy  `pulumi:"lBCookieStickinessPolicy"`
	Listeners                     []LoadBalancerListeners                 `pulumi:"listeners"`
	Policies                      []LoadBalancerPolicies                  `pulumi:"policies"`
	SecurityGroups                []string                                `pulumi:"securityGroups"`
	SourceSecurityGroupGroupName  *string                                 `pulumi:"sourceSecurityGroupGroupName"`
	SourceSecurityGroupOwnerAlias *string                                 `pulumi:"sourceSecurityGroupOwnerAlias"`
	Subnets                       []string                                `pulumi:"subnets"`
	Tags                          []LoadBalancerTag                       `pulumi:"tags"`
}

func LookupLoadBalancerOutput(ctx *pulumi.Context, args LookupLoadBalancerOutputArgs, opts ...pulumi.InvokeOption) LookupLoadBalancerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoadBalancerResult, error) {
			args := v.(LookupLoadBalancerArgs)
			r, err := LookupLoadBalancer(ctx, &args, opts...)
			var s LookupLoadBalancerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLoadBalancerResultOutput)
}

type LookupLoadBalancerOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupLoadBalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerArgs)(nil)).Elem()
}

type LookupLoadBalancerResultOutput struct{ *pulumi.OutputState }

func (LookupLoadBalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerResult)(nil)).Elem()
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutput() LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutputWithContext(ctx context.Context) LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) AccessLoggingPolicy() LoadBalancerAccessLoggingPolicyPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *LoadBalancerAccessLoggingPolicy { return v.AccessLoggingPolicy }).(LoadBalancerAccessLoggingPolicyPtrOutput)
}

func (o LookupLoadBalancerResultOutput) AppCookieStickinessPolicy() LoadBalancerAppCookieStickinessPolicyArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []LoadBalancerAppCookieStickinessPolicy {
		return v.AppCookieStickinessPolicy
	}).(LoadBalancerAppCookieStickinessPolicyArrayOutput)
}

func (o LookupLoadBalancerResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o LookupLoadBalancerResultOutput) CanonicalHostedZoneName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.CanonicalHostedZoneName }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) CanonicalHostedZoneNameID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.CanonicalHostedZoneNameID }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) ConnectionDrainingPolicy() LoadBalancerConnectionDrainingPolicyPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *LoadBalancerConnectionDrainingPolicy {
		return v.ConnectionDrainingPolicy
	}).(LoadBalancerConnectionDrainingPolicyPtrOutput)
}

func (o LookupLoadBalancerResultOutput) ConnectionSettings() LoadBalancerConnectionSettingsPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *LoadBalancerConnectionSettings { return v.ConnectionSettings }).(LoadBalancerConnectionSettingsPtrOutput)
}

func (o LookupLoadBalancerResultOutput) CrossZone() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *bool { return v.CrossZone }).(pulumi.BoolPtrOutput)
}

func (o LookupLoadBalancerResultOutput) DNSName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.DNSName }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) HealthCheck() LoadBalancerHealthCheckPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *LoadBalancerHealthCheck { return v.HealthCheck }).(LoadBalancerHealthCheckPtrOutput)
}

func (o LookupLoadBalancerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) Instances() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []string { return v.Instances }).(pulumi.StringArrayOutput)
}

func (o LookupLoadBalancerResultOutput) LBCookieStickinessPolicy() LoadBalancerLBCookieStickinessPolicyArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []LoadBalancerLBCookieStickinessPolicy {
		return v.LBCookieStickinessPolicy
	}).(LoadBalancerLBCookieStickinessPolicyArrayOutput)
}

func (o LookupLoadBalancerResultOutput) Listeners() LoadBalancerListenersArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []LoadBalancerListeners { return v.Listeners }).(LoadBalancerListenersArrayOutput)
}

func (o LookupLoadBalancerResultOutput) Policies() LoadBalancerPoliciesArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []LoadBalancerPolicies { return v.Policies }).(LoadBalancerPoliciesArrayOutput)
}

func (o LookupLoadBalancerResultOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []string { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

func (o LookupLoadBalancerResultOutput) SourceSecurityGroupGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.SourceSecurityGroupGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) SourceSecurityGroupOwnerAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.SourceSecurityGroupOwnerAlias }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) Subnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []string { return v.Subnets }).(pulumi.StringArrayOutput)
}

func (o LookupLoadBalancerResultOutput) Tags() LoadBalancerTagArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []LoadBalancerTag { return v.Tags }).(LoadBalancerTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoadBalancerResultOutput{})
}
