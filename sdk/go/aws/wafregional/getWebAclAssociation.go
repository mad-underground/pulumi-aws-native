// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::WAFRegional::WebACLAssociation
func LookupWebAclAssociation(ctx *pulumi.Context, args *LookupWebAclAssociationArgs, opts ...pulumi.InvokeOption) (*LookupWebAclAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAclAssociationResult
	err := ctx.Invoke("aws-native:wafregional:getWebAclAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAclAssociationArgs struct {
	Id string `pulumi:"id"`
}

type LookupWebAclAssociationResult struct {
	Id *string `pulumi:"id"`
}

func LookupWebAclAssociationOutput(ctx *pulumi.Context, args LookupWebAclAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupWebAclAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAclAssociationResult, error) {
			args := v.(LookupWebAclAssociationArgs)
			r, err := LookupWebAclAssociation(ctx, &args, opts...)
			var s LookupWebAclAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAclAssociationResultOutput)
}

type LookupWebAclAssociationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupWebAclAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAclAssociationArgs)(nil)).Elem()
}

type LookupWebAclAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupWebAclAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAclAssociationResult)(nil)).Elem()
}

func (o LookupWebAclAssociationResultOutput) ToLookupWebAclAssociationResultOutput() LookupWebAclAssociationResultOutput {
	return o
}

func (o LookupWebAclAssociationResultOutput) ToLookupWebAclAssociationResultOutputWithContext(ctx context.Context) LookupWebAclAssociationResultOutput {
	return o
}

func (o LookupWebAclAssociationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAclAssociationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAclAssociationResultOutput{})
}
