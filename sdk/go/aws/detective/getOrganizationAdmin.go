// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package detective

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Detective::OrganizationAdmin
func LookupOrganizationAdmin(ctx *pulumi.Context, args *LookupOrganizationAdminArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationAdminResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrganizationAdminResult
	err := ctx.Invoke("aws-native:detective:getOrganizationAdmin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrganizationAdminArgs struct {
	// The account ID of the account that should be registered as your Organization's delegated administrator for Detective
	AccountId string `pulumi:"accountId"`
}

type LookupOrganizationAdminResult struct {
	// The Detective graph ARN
	GraphArn *string `pulumi:"graphArn"`
}

func LookupOrganizationAdminOutput(ctx *pulumi.Context, args LookupOrganizationAdminOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationAdminResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationAdminResult, error) {
			args := v.(LookupOrganizationAdminArgs)
			r, err := LookupOrganizationAdmin(ctx, &args, opts...)
			var s LookupOrganizationAdminResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrganizationAdminResultOutput)
}

type LookupOrganizationAdminOutputArgs struct {
	// The account ID of the account that should be registered as your Organization's delegated administrator for Detective
	AccountId pulumi.StringInput `pulumi:"accountId"`
}

func (LookupOrganizationAdminOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationAdminArgs)(nil)).Elem()
}

type LookupOrganizationAdminResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationAdminResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationAdminResult)(nil)).Elem()
}

func (o LookupOrganizationAdminResultOutput) ToLookupOrganizationAdminResultOutput() LookupOrganizationAdminResultOutput {
	return o
}

func (o LookupOrganizationAdminResultOutput) ToLookupOrganizationAdminResultOutputWithContext(ctx context.Context) LookupOrganizationAdminResultOutput {
	return o
}

// The Detective graph ARN
func (o LookupOrganizationAdminResultOutput) GraphArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrganizationAdminResult) *string { return v.GraphArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationAdminResultOutput{})
}
