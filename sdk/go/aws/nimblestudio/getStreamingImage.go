// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nimblestudio

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a streaming session machine image that can be used to launch a streaming session
func LookupStreamingImage(ctx *pulumi.Context, args *LookupStreamingImageArgs, opts ...pulumi.InvokeOption) (*LookupStreamingImageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStreamingImageResult
	err := ctx.Invoke("aws-native:nimblestudio:getStreamingImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamingImageArgs struct {
	StreamingImageId string `pulumi:"streamingImageId"`
	// <p>The studioId. </p>
	StudioId string `pulumi:"studioId"`
}

type LookupStreamingImageResult struct {
	// <p>A human-readable description of the streaming image.</p>
	Description             *string                                `pulumi:"description"`
	EncryptionConfiguration *StreamingImageEncryptionConfiguration `pulumi:"encryptionConfiguration"`
	// <p>The list of EULAs that must be accepted before a Streaming Session can be started using this streaming image.</p>
	EulaIds []string `pulumi:"eulaIds"`
	// <p>A friendly name for a streaming image resource.</p>
	Name *string `pulumi:"name"`
	// <p>The owner of the streaming image, either the studioId that contains the streaming image, or 'amazon' for images that are provided by Amazon Nimble Studio.</p>
	Owner *string `pulumi:"owner"`
	// <p>The platform of the streaming image, either WINDOWS or LINUX.</p>
	Platform         *string `pulumi:"platform"`
	StreamingImageId *string `pulumi:"streamingImageId"`
}

func LookupStreamingImageOutput(ctx *pulumi.Context, args LookupStreamingImageOutputArgs, opts ...pulumi.InvokeOption) LookupStreamingImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStreamingImageResult, error) {
			args := v.(LookupStreamingImageArgs)
			r, err := LookupStreamingImage(ctx, &args, opts...)
			var s LookupStreamingImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStreamingImageResultOutput)
}

type LookupStreamingImageOutputArgs struct {
	StreamingImageId pulumi.StringInput `pulumi:"streamingImageId"`
	// <p>The studioId. </p>
	StudioId pulumi.StringInput `pulumi:"studioId"`
}

func (LookupStreamingImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingImageArgs)(nil)).Elem()
}

type LookupStreamingImageResultOutput struct{ *pulumi.OutputState }

func (LookupStreamingImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingImageResult)(nil)).Elem()
}

func (o LookupStreamingImageResultOutput) ToLookupStreamingImageResultOutput() LookupStreamingImageResultOutput {
	return o
}

func (o LookupStreamingImageResultOutput) ToLookupStreamingImageResultOutputWithContext(ctx context.Context) LookupStreamingImageResultOutput {
	return o
}

// <p>A human-readable description of the streaming image.</p>
func (o LookupStreamingImageResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingImageResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingImageResultOutput) EncryptionConfiguration() StreamingImageEncryptionConfigurationPtrOutput {
	return o.ApplyT(func(v LookupStreamingImageResult) *StreamingImageEncryptionConfiguration {
		return v.EncryptionConfiguration
	}).(StreamingImageEncryptionConfigurationPtrOutput)
}

// <p>The list of EULAs that must be accepted before a Streaming Session can be started using this streaming image.</p>
func (o LookupStreamingImageResultOutput) EulaIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStreamingImageResult) []string { return v.EulaIds }).(pulumi.StringArrayOutput)
}

// <p>A friendly name for a streaming image resource.</p>
func (o LookupStreamingImageResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingImageResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// <p>The owner of the streaming image, either the studioId that contains the streaming image, or 'amazon' for images that are provided by Amazon Nimble Studio.</p>
func (o LookupStreamingImageResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingImageResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

// <p>The platform of the streaming image, either WINDOWS or LINUX.</p>
func (o LookupStreamingImageResultOutput) Platform() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingImageResult) *string { return v.Platform }).(pulumi.StringPtrOutput)
}

func (o LookupStreamingImageResultOutput) StreamingImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingImageResult) *string { return v.StreamingImageId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStreamingImageResultOutput{})
}
