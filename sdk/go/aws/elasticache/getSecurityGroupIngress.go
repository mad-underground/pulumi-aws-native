// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElastiCache::SecurityGroupIngress
func LookupSecurityGroupIngress(ctx *pulumi.Context, args *LookupSecurityGroupIngressArgs, opts ...pulumi.InvokeOption) (*LookupSecurityGroupIngressResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityGroupIngressResult
	err := ctx.Invoke("aws-native:elasticache:getSecurityGroupIngress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityGroupIngressArgs struct {
	Id string `pulumi:"id"`
}

type LookupSecurityGroupIngressResult struct {
	CacheSecurityGroupName  *string `pulumi:"cacheSecurityGroupName"`
	EC2SecurityGroupName    *string `pulumi:"eC2SecurityGroupName"`
	EC2SecurityGroupOwnerId *string `pulumi:"eC2SecurityGroupOwnerId"`
	Id                      *string `pulumi:"id"`
}

func LookupSecurityGroupIngressOutput(ctx *pulumi.Context, args LookupSecurityGroupIngressOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityGroupIngressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityGroupIngressResult, error) {
			args := v.(LookupSecurityGroupIngressArgs)
			r, err := LookupSecurityGroupIngress(ctx, &args, opts...)
			var s LookupSecurityGroupIngressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityGroupIngressResultOutput)
}

type LookupSecurityGroupIngressOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupSecurityGroupIngressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupIngressArgs)(nil)).Elem()
}

type LookupSecurityGroupIngressResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityGroupIngressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupIngressResult)(nil)).Elem()
}

func (o LookupSecurityGroupIngressResultOutput) ToLookupSecurityGroupIngressResultOutput() LookupSecurityGroupIngressResultOutput {
	return o
}

func (o LookupSecurityGroupIngressResultOutput) ToLookupSecurityGroupIngressResultOutputWithContext(ctx context.Context) LookupSecurityGroupIngressResultOutput {
	return o
}

func (o LookupSecurityGroupIngressResultOutput) CacheSecurityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityGroupIngressResult) *string { return v.CacheSecurityGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityGroupIngressResultOutput) EC2SecurityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityGroupIngressResult) *string { return v.EC2SecurityGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityGroupIngressResultOutput) EC2SecurityGroupOwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityGroupIngressResult) *string { return v.EC2SecurityGroupOwnerId }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityGroupIngressResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityGroupIngressResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityGroupIngressResultOutput{})
}
