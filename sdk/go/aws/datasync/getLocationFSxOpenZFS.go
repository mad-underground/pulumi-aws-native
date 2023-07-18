// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DataSync::LocationFSxOpenZFS.
func LookupLocationFSxOpenZFS(ctx *pulumi.Context, args *LookupLocationFSxOpenZFSArgs, opts ...pulumi.InvokeOption) (*LookupLocationFSxOpenZFSResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocationFSxOpenZFSResult
	err := ctx.Invoke("aws-native:datasync:getLocationFSxOpenZFS", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLocationFSxOpenZFSArgs struct {
	// The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created.
	LocationArn string `pulumi:"locationArn"`
}

type LookupLocationFSxOpenZFSResult struct {
	// The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created.
	LocationArn *string `pulumi:"locationArn"`
	// The URL of the FSx OpenZFS that was described.
	LocationUri *string `pulumi:"locationUri"`
	// An array of key-value pairs to apply to this resource.
	Tags []LocationFSxOpenZFSTag `pulumi:"tags"`
}

func LookupLocationFSxOpenZFSOutput(ctx *pulumi.Context, args LookupLocationFSxOpenZFSOutputArgs, opts ...pulumi.InvokeOption) LookupLocationFSxOpenZFSResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocationFSxOpenZFSResult, error) {
			args := v.(LookupLocationFSxOpenZFSArgs)
			r, err := LookupLocationFSxOpenZFS(ctx, &args, opts...)
			var s LookupLocationFSxOpenZFSResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocationFSxOpenZFSResultOutput)
}

type LookupLocationFSxOpenZFSOutputArgs struct {
	// The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created.
	LocationArn pulumi.StringInput `pulumi:"locationArn"`
}

func (LookupLocationFSxOpenZFSOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationFSxOpenZFSArgs)(nil)).Elem()
}

type LookupLocationFSxOpenZFSResultOutput struct{ *pulumi.OutputState }

func (LookupLocationFSxOpenZFSResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocationFSxOpenZFSResult)(nil)).Elem()
}

func (o LookupLocationFSxOpenZFSResultOutput) ToLookupLocationFSxOpenZFSResultOutput() LookupLocationFSxOpenZFSResultOutput {
	return o
}

func (o LookupLocationFSxOpenZFSResultOutput) ToLookupLocationFSxOpenZFSResultOutputWithContext(ctx context.Context) LookupLocationFSxOpenZFSResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created.
func (o LookupLocationFSxOpenZFSResultOutput) LocationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationFSxOpenZFSResult) *string { return v.LocationArn }).(pulumi.StringPtrOutput)
}

// The URL of the FSx OpenZFS that was described.
func (o LookupLocationFSxOpenZFSResultOutput) LocationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocationFSxOpenZFSResult) *string { return v.LocationUri }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupLocationFSxOpenZFSResultOutput) Tags() LocationFSxOpenZFSTagArrayOutput {
	return o.ApplyT(func(v LookupLocationFSxOpenZFSResult) []LocationFSxOpenZFSTag { return v.Tags }).(LocationFSxOpenZFSTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocationFSxOpenZFSResultOutput{})
}
