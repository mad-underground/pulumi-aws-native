// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::RDS::DBSecurityGroupIngress
func LookupDBSecurityGroupIngress(ctx *pulumi.Context, args *LookupDBSecurityGroupIngressArgs, opts ...pulumi.InvokeOption) (*LookupDBSecurityGroupIngressResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDBSecurityGroupIngressResult
	err := ctx.Invoke("aws-native:rds:getDBSecurityGroupIngress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDBSecurityGroupIngressArgs struct {
	Id string `pulumi:"id"`
}

type LookupDBSecurityGroupIngressResult struct {
	CIDRIP                  *string `pulumi:"cIDRIP"`
	DBSecurityGroupName     *string `pulumi:"dBSecurityGroupName"`
	EC2SecurityGroupId      *string `pulumi:"eC2SecurityGroupId"`
	EC2SecurityGroupName    *string `pulumi:"eC2SecurityGroupName"`
	EC2SecurityGroupOwnerId *string `pulumi:"eC2SecurityGroupOwnerId"`
	Id                      *string `pulumi:"id"`
}

func LookupDBSecurityGroupIngressOutput(ctx *pulumi.Context, args LookupDBSecurityGroupIngressOutputArgs, opts ...pulumi.InvokeOption) LookupDBSecurityGroupIngressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDBSecurityGroupIngressResult, error) {
			args := v.(LookupDBSecurityGroupIngressArgs)
			r, err := LookupDBSecurityGroupIngress(ctx, &args, opts...)
			var s LookupDBSecurityGroupIngressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDBSecurityGroupIngressResultOutput)
}

type LookupDBSecurityGroupIngressOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDBSecurityGroupIngressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBSecurityGroupIngressArgs)(nil)).Elem()
}

type LookupDBSecurityGroupIngressResultOutput struct{ *pulumi.OutputState }

func (LookupDBSecurityGroupIngressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBSecurityGroupIngressResult)(nil)).Elem()
}

func (o LookupDBSecurityGroupIngressResultOutput) ToLookupDBSecurityGroupIngressResultOutput() LookupDBSecurityGroupIngressResultOutput {
	return o
}

func (o LookupDBSecurityGroupIngressResultOutput) ToLookupDBSecurityGroupIngressResultOutputWithContext(ctx context.Context) LookupDBSecurityGroupIngressResultOutput {
	return o
}

func (o LookupDBSecurityGroupIngressResultOutput) CIDRIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBSecurityGroupIngressResult) *string { return v.CIDRIP }).(pulumi.StringPtrOutput)
}

func (o LookupDBSecurityGroupIngressResultOutput) DBSecurityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBSecurityGroupIngressResult) *string { return v.DBSecurityGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupDBSecurityGroupIngressResultOutput) EC2SecurityGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBSecurityGroupIngressResult) *string { return v.EC2SecurityGroupId }).(pulumi.StringPtrOutput)
}

func (o LookupDBSecurityGroupIngressResultOutput) EC2SecurityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBSecurityGroupIngressResult) *string { return v.EC2SecurityGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupDBSecurityGroupIngressResultOutput) EC2SecurityGroupOwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBSecurityGroupIngressResult) *string { return v.EC2SecurityGroupOwnerId }).(pulumi.StringPtrOutput)
}

func (o LookupDBSecurityGroupIngressResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBSecurityGroupIngressResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDBSecurityGroupIngressResultOutput{})
}
