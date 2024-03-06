// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotsitewise

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::IoTSiteWise::Asset
func LookupAsset(ctx *pulumi.Context, args *LookupAssetArgs, opts ...pulumi.InvokeOption) (*LookupAssetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAssetResult
	err := ctx.Invoke("aws-native:iotsitewise:getAsset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssetArgs struct {
	// The ID of the asset
	AssetId string `pulumi:"assetId"`
}

type LookupAssetResult struct {
	// The ARN of the asset
	AssetArn *string `pulumi:"assetArn"`
	// A description for the asset
	AssetDescription *string `pulumi:"assetDescription"`
	// The External ID of the asset
	AssetExternalId  *string          `pulumi:"assetExternalId"`
	AssetHierarchies []AssetHierarchy `pulumi:"assetHierarchies"`
	// The ID of the asset
	AssetId *string `pulumi:"assetId"`
	// The ID of the asset model from which to create the asset.
	AssetModelId *string `pulumi:"assetModelId"`
	// A unique, friendly name for the asset.
	AssetName       *string         `pulumi:"assetName"`
	AssetProperties []AssetProperty `pulumi:"assetProperties"`
	// A list of key-value pairs that contain metadata for the asset.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupAssetOutput(ctx *pulumi.Context, args LookupAssetOutputArgs, opts ...pulumi.InvokeOption) LookupAssetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssetResult, error) {
			args := v.(LookupAssetArgs)
			r, err := LookupAsset(ctx, &args, opts...)
			var s LookupAssetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssetResultOutput)
}

type LookupAssetOutputArgs struct {
	// The ID of the asset
	AssetId pulumi.StringInput `pulumi:"assetId"`
}

func (LookupAssetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssetArgs)(nil)).Elem()
}

type LookupAssetResultOutput struct{ *pulumi.OutputState }

func (LookupAssetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssetResult)(nil)).Elem()
}

func (o LookupAssetResultOutput) ToLookupAssetResultOutput() LookupAssetResultOutput {
	return o
}

func (o LookupAssetResultOutput) ToLookupAssetResultOutputWithContext(ctx context.Context) LookupAssetResultOutput {
	return o
}

// The ARN of the asset
func (o LookupAssetResultOutput) AssetArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.AssetArn }).(pulumi.StringPtrOutput)
}

// A description for the asset
func (o LookupAssetResultOutput) AssetDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.AssetDescription }).(pulumi.StringPtrOutput)
}

// The External ID of the asset
func (o LookupAssetResultOutput) AssetExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.AssetExternalId }).(pulumi.StringPtrOutput)
}

func (o LookupAssetResultOutput) AssetHierarchies() AssetHierarchyArrayOutput {
	return o.ApplyT(func(v LookupAssetResult) []AssetHierarchy { return v.AssetHierarchies }).(AssetHierarchyArrayOutput)
}

// The ID of the asset
func (o LookupAssetResultOutput) AssetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.AssetId }).(pulumi.StringPtrOutput)
}

// The ID of the asset model from which to create the asset.
func (o LookupAssetResultOutput) AssetModelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.AssetModelId }).(pulumi.StringPtrOutput)
}

// A unique, friendly name for the asset.
func (o LookupAssetResultOutput) AssetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.AssetName }).(pulumi.StringPtrOutput)
}

func (o LookupAssetResultOutput) AssetProperties() AssetPropertyArrayOutput {
	return o.ApplyT(func(v LookupAssetResult) []AssetProperty { return v.AssetProperties }).(AssetPropertyArrayOutput)
}

// A list of key-value pairs that contain metadata for the asset.
func (o LookupAssetResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupAssetResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssetResultOutput{})
}
