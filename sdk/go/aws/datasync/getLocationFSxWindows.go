// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DataSync::LocationFSxWindows.
func LookupLocationFSxWindows(ctx *pulumi.Context, args *LookupLocationFSxWindowsArgs, opts ...pulumi.InvokeOption) (*LookupLocationFSxWindowsResult, error) {
	var rv LookupLocationFSxWindowsResult
	err := ctx.Invoke("aws-native:datasync:getLocationFSxWindows", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLocationFSxWindowsArgs struct {
	// The Amazon Resource Name (ARN) of the Amazon FSx for Windows file system location that is created.
	LocationArn string `pulumi:"locationArn"`
}

type LookupLocationFSxWindowsResult struct {
	// The Amazon Resource Name (ARN) of the Amazon FSx for Windows file system location that is created.
	LocationArn *string `pulumi:"locationArn"`
	// The URL of the FSx for Windows location that was described.
	LocationUri *string `pulumi:"locationUri"`
	// An array of key-value pairs to apply to this resource.
	Tags []LocationFSxWindowsTag `pulumi:"tags"`
}

func LookupLocationFSxWindowsOutput(ctx *pulumi.Context, args LookupLocationFSxWindowsOutputArgs, opts ...pulumi.InvokeOption) LookupLocationFSxWindowsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocationFSxWindowsResult, error) {
			args := v.(LookupLocationFSxWindowsArgs)
			r, err := LookupLocationFSxWindows(ctx, &args, opts...)
			return *r, err
		}).(LookupLocationFSxWindowsResultOutput)
}

type LookupLocationFSxWindowsOutputArgs struct {
	// The Amazon Resource Name (ARN) of the Amazon FSx for Windows file system location that is created.
	LocationArn pulumi.StringInput `pulumi:"locationArn"`
}

func (LookupLocationFSxWindowsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationFSxWindowsArgs)(nil)).Elem()
}

type LookupLocationFSxWindowsResultOutput struct{ *pulumi.OutputState }

func (LookupLocationFSxWindowsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationFSxWindowsResult)(nil)).Elem()
}

func (o LookupLocationFSxWindowsResultOutput) ToLookupLocationFSxWindowsResultOutput() LookupLocationFSxWindowsResultOutput {
	return o
}

func (o LookupLocationFSxWindowsResultOutput) ToLookupLocationFSxWindowsResultOutputWithContext(ctx context.Context) LookupLocationFSxWindowsResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the Amazon FSx for Windows file system location that is created.
func (o LookupLocationFSxWindowsResultOutput) LocationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationFSxWindowsResult) *string { return v.LocationArn }).(pulumi.StringPtrOutput)
}

// The URL of the FSx for Windows location that was described.
func (o LookupLocationFSxWindowsResultOutput) LocationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationFSxWindowsResult) *string { return v.LocationUri }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupLocationFSxWindowsResultOutput) Tags() LocationFSxWindowsTagArrayOutput {
	return o.ApplyT(func(v LookupLocationFSxWindowsResult) []LocationFSxWindowsTag { return v.Tags }).(LocationFSxWindowsTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocationFSxWindowsResultOutput{})
}
