// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3outposts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type Definition for AWS::S3Outposts::Bucket
func LookupBucket(ctx *pulumi.Context, args *LookupBucketArgs, opts ...pulumi.InvokeOption) (*LookupBucketResult, error) {
	var rv LookupBucketResult
	err := ctx.Invoke("aws-native:s3outposts:getBucket", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBucketArgs struct {
	// The Amazon Resource Name (ARN) of the specified bucket.
	Arn string `pulumi:"arn"`
}

type LookupBucketResult struct {
	// The Amazon Resource Name (ARN) of the specified bucket.
	Arn *string `pulumi:"arn"`
	// Rules that define how Amazon S3Outposts manages objects during their lifetime.
	LifecycleConfiguration *BucketLifecycleConfiguration `pulumi:"lifecycleConfiguration"`
	// An arbitrary set of tags (key-value pairs) for this S3Outposts bucket.
	Tags []BucketTag `pulumi:"tags"`
}

func LookupBucketOutput(ctx *pulumi.Context, args LookupBucketOutputArgs, opts ...pulumi.InvokeOption) LookupBucketResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBucketResult, error) {
			args := v.(LookupBucketArgs)
			r, err := LookupBucket(ctx, &args, opts...)
			return *r, err
		}).(LookupBucketResultOutput)
}

type LookupBucketOutputArgs struct {
	// The Amazon Resource Name (ARN) of the specified bucket.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupBucketOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketArgs)(nil)).Elem()
}

type LookupBucketResultOutput struct{ *pulumi.OutputState }

func (LookupBucketResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketResult)(nil)).Elem()
}

func (o LookupBucketResultOutput) ToLookupBucketResultOutput() LookupBucketResultOutput {
	return o
}

func (o LookupBucketResultOutput) ToLookupBucketResultOutputWithContext(ctx context.Context) LookupBucketResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the specified bucket.
func (o LookupBucketResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Rules that define how Amazon S3Outposts manages objects during their lifetime.
func (o LookupBucketResultOutput) LifecycleConfiguration() BucketLifecycleConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketLifecycleConfiguration { return v.LifecycleConfiguration }).(BucketLifecycleConfigurationPtrOutput)
}

// An arbitrary set of tags (key-value pairs) for this S3Outposts bucket.
func (o LookupBucketResultOutput) Tags() BucketTagArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []BucketTag { return v.Tags }).(BucketTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBucketResultOutput{})
}