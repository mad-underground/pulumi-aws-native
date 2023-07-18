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
func LookupWebACLAssociation(ctx *pulumi.Context, args *LookupWebACLAssociationArgs, opts ...pulumi.InvokeOption) (*LookupWebACLAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWebACLAssociationResult
	err := ctx.Invoke("aws-native:wafregional:getWebACLAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebACLAssociationArgs struct {
	Id string `pulumi:"id"`
}

type LookupWebACLAssociationResult struct {
	Id *string `pulumi:"id"`
}

func LookupWebACLAssociationOutput(ctx *pulumi.Context, args LookupWebACLAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupWebACLAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebACLAssociationResult, error) {
			args := v.(LookupWebACLAssociationArgs)
			r, err := LookupWebACLAssociation(ctx, &args, opts...)
			var s LookupWebACLAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebACLAssociationResultOutput)
}

type LookupWebACLAssociationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupWebACLAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebACLAssociationArgs)(nil)).Elem()
}

type LookupWebACLAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupWebACLAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebACLAssociationResult)(nil)).Elem()
}

func (o LookupWebACLAssociationResultOutput) ToLookupWebACLAssociationResultOutput() LookupWebACLAssociationResultOutput {
	return o
}

func (o LookupWebACLAssociationResultOutput) ToLookupWebACLAssociationResultOutputWithContext(ctx context.Context) LookupWebACLAssociationResultOutput {
	return o
}

func (o LookupWebACLAssociationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebACLAssociationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebACLAssociationResultOutput{})
}
