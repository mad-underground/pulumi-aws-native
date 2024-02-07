// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::S3::AccessGrantsLocation resource is an Amazon S3 resource type hosted in an access grants instance which can be the target of S3 access grants.
func LookupAccessGrantsLocation(ctx *pulumi.Context, args *LookupAccessGrantsLocationArgs, opts ...pulumi.InvokeOption) (*LookupAccessGrantsLocationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessGrantsLocationResult
	err := ctx.Invoke("aws-native:s3:getAccessGrantsLocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessGrantsLocationArgs struct {
	// The unique identifier for the specified Access Grants location.
	AccessGrantsLocationId string `pulumi:"accessGrantsLocationId"`
}

type LookupAccessGrantsLocationResult struct {
	// The Amazon Resource Name (ARN) of the specified Access Grants location.
	AccessGrantsLocationArn *string `pulumi:"accessGrantsLocationArn"`
	// The unique identifier for the specified Access Grants location.
	AccessGrantsLocationId *string `pulumi:"accessGrantsLocationId"`
	// The Amazon Resource Name (ARN) of the access grant location's associated IAM role.
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// Descriptor for where the location actually points
	LocationScope *string `pulumi:"locationScope"`
}

func LookupAccessGrantsLocationOutput(ctx *pulumi.Context, args LookupAccessGrantsLocationOutputArgs, opts ...pulumi.InvokeOption) LookupAccessGrantsLocationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessGrantsLocationResult, error) {
			args := v.(LookupAccessGrantsLocationArgs)
			r, err := LookupAccessGrantsLocation(ctx, &args, opts...)
			var s LookupAccessGrantsLocationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessGrantsLocationResultOutput)
}

type LookupAccessGrantsLocationOutputArgs struct {
	// The unique identifier for the specified Access Grants location.
	AccessGrantsLocationId pulumi.StringInput `pulumi:"accessGrantsLocationId"`
}

func (LookupAccessGrantsLocationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessGrantsLocationArgs)(nil)).Elem()
}

type LookupAccessGrantsLocationResultOutput struct{ *pulumi.OutputState }

func (LookupAccessGrantsLocationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessGrantsLocationResult)(nil)).Elem()
}

func (o LookupAccessGrantsLocationResultOutput) ToLookupAccessGrantsLocationResultOutput() LookupAccessGrantsLocationResultOutput {
	return o
}

func (o LookupAccessGrantsLocationResultOutput) ToLookupAccessGrantsLocationResultOutputWithContext(ctx context.Context) LookupAccessGrantsLocationResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the specified Access Grants location.
func (o LookupAccessGrantsLocationResultOutput) AccessGrantsLocationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessGrantsLocationResult) *string { return v.AccessGrantsLocationArn }).(pulumi.StringPtrOutput)
}

// The unique identifier for the specified Access Grants location.
func (o LookupAccessGrantsLocationResultOutput) AccessGrantsLocationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessGrantsLocationResult) *string { return v.AccessGrantsLocationId }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the access grant location's associated IAM role.
func (o LookupAccessGrantsLocationResultOutput) IamRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessGrantsLocationResult) *string { return v.IamRoleArn }).(pulumi.StringPtrOutput)
}

// Descriptor for where the location actually points
func (o LookupAccessGrantsLocationResultOutput) LocationScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessGrantsLocationResult) *string { return v.LocationScope }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessGrantsLocationResultOutput{})
}
