// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rekognition

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::Rekognition::Collection type creates an Amazon Rekognition Collection. A collection is a logical grouping of information about detected faces which can later be referenced for searches on the group
func LookupCollection(ctx *pulumi.Context, args *LookupCollectionArgs, opts ...pulumi.InvokeOption) (*LookupCollectionResult, error) {
	var rv LookupCollectionResult
	err := ctx.Invoke("aws-native:rekognition:getCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCollectionArgs struct {
	CollectionId string `pulumi:"collectionId"`
}

type LookupCollectionResult struct {
	Arn *string `pulumi:"arn"`
	// An array of key-value pairs to apply to this resource.
	Tags []CollectionTag `pulumi:"tags"`
}

func LookupCollectionOutput(ctx *pulumi.Context, args LookupCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupCollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCollectionResult, error) {
			args := v.(LookupCollectionArgs)
			r, err := LookupCollection(ctx, &args, opts...)
			return *r, err
		}).(LookupCollectionResultOutput)
}

type LookupCollectionOutputArgs struct {
	CollectionId pulumi.StringInput `pulumi:"collectionId"`
}

func (LookupCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCollectionArgs)(nil)).Elem()
}

type LookupCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCollectionResult)(nil)).Elem()
}

func (o LookupCollectionResultOutput) ToLookupCollectionResultOutput() LookupCollectionResultOutput {
	return o
}

func (o LookupCollectionResultOutput) ToLookupCollectionResultOutputWithContext(ctx context.Context) LookupCollectionResultOutput {
	return o
}

func (o LookupCollectionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCollectionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupCollectionResultOutput) Tags() CollectionTagArrayOutput {
	return o.ApplyT(func(v LookupCollectionResult) []CollectionTag { return v.Tags }).(CollectionTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCollectionResultOutput{})
}