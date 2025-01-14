// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Adds the specified outbound (egress) rule to a security group.
//
//	An outbound rule permits instances to send traffic to the specified IPv4 or IPv6 address range, the IP addresses that are specified by a prefix list, or the instances that are associated with a destination security group. For more information, see [Security group rules](https://docs.aws.amazon.com/vpc/latest/userguide/security-group-rules.html).
//	You must specify exactly one of the following destinations: an IPv4 or IPv6 address range, a prefix list, or a security group. Otherwise, the stack launches successfully but the rule is not added to the security group.
//	You must specify a protocol for each rule (for example, TCP). If the protocol is TCP or UDP, you must also specify a port or port range. If the protocol is ICMP or ICMPv6, you must also specify the ICMP/ICMPv6 type and code. To specify all types or all codes, use -1.
//	Rule changes are propagated to instances associated with the security group as quickly as possible
func LookupSecurityGroupEgress(ctx *pulumi.Context, args *LookupSecurityGroupEgressArgs, opts ...pulumi.InvokeOption) (*LookupSecurityGroupEgressResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityGroupEgressResult
	err := ctx.Invoke("aws-native:ec2:getSecurityGroupEgress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityGroupEgressArgs struct {
	Id string `pulumi:"id"`
}

type LookupSecurityGroupEgressResult struct {
	// The description of an egress (outbound) security group rule.
	//  Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*
	Description *string `pulumi:"description"`
	Id          *string `pulumi:"id"`
}

func LookupSecurityGroupEgressOutput(ctx *pulumi.Context, args LookupSecurityGroupEgressOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityGroupEgressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityGroupEgressResult, error) {
			args := v.(LookupSecurityGroupEgressArgs)
			r, err := LookupSecurityGroupEgress(ctx, &args, opts...)
			var s LookupSecurityGroupEgressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityGroupEgressResultOutput)
}

type LookupSecurityGroupEgressOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupSecurityGroupEgressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupEgressArgs)(nil)).Elem()
}

type LookupSecurityGroupEgressResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityGroupEgressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupEgressResult)(nil)).Elem()
}

func (o LookupSecurityGroupEgressResultOutput) ToLookupSecurityGroupEgressResultOutput() LookupSecurityGroupEgressResultOutput {
	return o
}

func (o LookupSecurityGroupEgressResultOutput) ToLookupSecurityGroupEgressResultOutputWithContext(ctx context.Context) LookupSecurityGroupEgressResultOutput {
	return o
}

// The description of an egress (outbound) security group rule.
//
//	Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*
func (o LookupSecurityGroupEgressResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityGroupEgressResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityGroupEgressResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityGroupEgressResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityGroupEgressResultOutput{})
}
