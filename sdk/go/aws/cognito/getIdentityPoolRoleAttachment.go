// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cognito::IdentityPoolRoleAttachment
func LookupIdentityPoolRoleAttachment(ctx *pulumi.Context, args *LookupIdentityPoolRoleAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupIdentityPoolRoleAttachmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIdentityPoolRoleAttachmentResult
	err := ctx.Invoke("aws-native:cognito:getIdentityPoolRoleAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIdentityPoolRoleAttachmentArgs struct {
	Id string `pulumi:"id"`
}

type LookupIdentityPoolRoleAttachmentResult struct {
	Id           *string     `pulumi:"id"`
	RoleMappings interface{} `pulumi:"roleMappings"`
	Roles        interface{} `pulumi:"roles"`
}

func LookupIdentityPoolRoleAttachmentOutput(ctx *pulumi.Context, args LookupIdentityPoolRoleAttachmentOutputArgs, opts ...pulumi.InvokeOption) LookupIdentityPoolRoleAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIdentityPoolRoleAttachmentResult, error) {
			args := v.(LookupIdentityPoolRoleAttachmentArgs)
			r, err := LookupIdentityPoolRoleAttachment(ctx, &args, opts...)
			var s LookupIdentityPoolRoleAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIdentityPoolRoleAttachmentResultOutput)
}

type LookupIdentityPoolRoleAttachmentOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupIdentityPoolRoleAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdentityPoolRoleAttachmentArgs)(nil)).Elem()
}

type LookupIdentityPoolRoleAttachmentResultOutput struct{ *pulumi.OutputState }

func (LookupIdentityPoolRoleAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdentityPoolRoleAttachmentResult)(nil)).Elem()
}

func (o LookupIdentityPoolRoleAttachmentResultOutput) ToLookupIdentityPoolRoleAttachmentResultOutput() LookupIdentityPoolRoleAttachmentResultOutput {
	return o
}

func (o LookupIdentityPoolRoleAttachmentResultOutput) ToLookupIdentityPoolRoleAttachmentResultOutputWithContext(ctx context.Context) LookupIdentityPoolRoleAttachmentResultOutput {
	return o
}

func (o LookupIdentityPoolRoleAttachmentResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentityPoolRoleAttachmentResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupIdentityPoolRoleAttachmentResultOutput) RoleMappings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIdentityPoolRoleAttachmentResult) interface{} { return v.RoleMappings }).(pulumi.AnyOutput)
}

func (o LookupIdentityPoolRoleAttachmentResultOutput) Roles() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIdentityPoolRoleAttachmentResult) interface{} { return v.Roles }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIdentityPoolRoleAttachmentResultOutput{})
}
