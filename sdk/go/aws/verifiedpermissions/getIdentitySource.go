// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package verifiedpermissions

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::VerifiedPermissions::IdentitySource Resource Type
func LookupIdentitySource(ctx *pulumi.Context, args *LookupIdentitySourceArgs, opts ...pulumi.InvokeOption) (*LookupIdentitySourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIdentitySourceResult
	err := ctx.Invoke("aws-native:verifiedpermissions:getIdentitySource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIdentitySourceArgs struct {
	IdentitySourceId string `pulumi:"identitySourceId"`
	PolicyStoreId    string `pulumi:"policyStoreId"`
}

type LookupIdentitySourceResult struct {
	Configuration       *IdentitySourceConfiguration `pulumi:"configuration"`
	Details             *IdentitySourceDetails       `pulumi:"details"`
	IdentitySourceId    *string                      `pulumi:"identitySourceId"`
	PrincipalEntityType *string                      `pulumi:"principalEntityType"`
}

func LookupIdentitySourceOutput(ctx *pulumi.Context, args LookupIdentitySourceOutputArgs, opts ...pulumi.InvokeOption) LookupIdentitySourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIdentitySourceResult, error) {
			args := v.(LookupIdentitySourceArgs)
			r, err := LookupIdentitySource(ctx, &args, opts...)
			var s LookupIdentitySourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIdentitySourceResultOutput)
}

type LookupIdentitySourceOutputArgs struct {
	IdentitySourceId pulumi.StringInput `pulumi:"identitySourceId"`
	PolicyStoreId    pulumi.StringInput `pulumi:"policyStoreId"`
}

func (LookupIdentitySourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdentitySourceArgs)(nil)).Elem()
}

type LookupIdentitySourceResultOutput struct{ *pulumi.OutputState }

func (LookupIdentitySourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdentitySourceResult)(nil)).Elem()
}

func (o LookupIdentitySourceResultOutput) ToLookupIdentitySourceResultOutput() LookupIdentitySourceResultOutput {
	return o
}

func (o LookupIdentitySourceResultOutput) ToLookupIdentitySourceResultOutputWithContext(ctx context.Context) LookupIdentitySourceResultOutput {
	return o
}

func (o LookupIdentitySourceResultOutput) Configuration() IdentitySourceConfigurationPtrOutput {
	return o.ApplyT(func(v LookupIdentitySourceResult) *IdentitySourceConfiguration { return v.Configuration }).(IdentitySourceConfigurationPtrOutput)
}

func (o LookupIdentitySourceResultOutput) Details() IdentitySourceDetailsPtrOutput {
	return o.ApplyT(func(v LookupIdentitySourceResult) *IdentitySourceDetails { return v.Details }).(IdentitySourceDetailsPtrOutput)
}

func (o LookupIdentitySourceResultOutput) IdentitySourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentitySourceResult) *string { return v.IdentitySourceId }).(pulumi.StringPtrOutput)
}

func (o LookupIdentitySourceResultOutput) PrincipalEntityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIdentitySourceResult) *string { return v.PrincipalEntityType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIdentitySourceResultOutput{})
}
