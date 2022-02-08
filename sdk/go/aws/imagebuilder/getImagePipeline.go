// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::ImageBuilder::ImagePipeline
func LookupImagePipeline(ctx *pulumi.Context, args *LookupImagePipelineArgs, opts ...pulumi.InvokeOption) (*LookupImagePipelineResult, error) {
	var rv LookupImagePipelineResult
	err := ctx.Invoke("aws-native:imagebuilder:getImagePipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupImagePipelineArgs struct {
	// The Amazon Resource Name (ARN) of the image pipeline.
	Arn string `pulumi:"arn"`
}

type LookupImagePipelineResult struct {
	// The Amazon Resource Name (ARN) of the image pipeline.
	Arn *string `pulumi:"arn"`
	// The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.
	ContainerRecipeArn *string `pulumi:"containerRecipeArn"`
	// The description of the image pipeline.
	Description *string `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.
	DistributionConfigurationArn *string `pulumi:"distributionConfigurationArn"`
	// Collects additional information about the image being created, including the operating system (OS) version and package list.
	EnhancedImageMetadataEnabled *bool `pulumi:"enhancedImageMetadataEnabled"`
	// The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.
	ImageRecipeArn *string `pulumi:"imageRecipeArn"`
	// The image tests configuration of the image pipeline.
	ImageTestsConfiguration *ImagePipelineImageTestsConfiguration `pulumi:"imageTestsConfiguration"`
	// The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.
	InfrastructureConfigurationArn *string `pulumi:"infrastructureConfigurationArn"`
	// The schedule of the image pipeline.
	Schedule *ImagePipelineSchedule `pulumi:"schedule"`
	// The status of the image pipeline.
	Status *ImagePipelineStatus `pulumi:"status"`
	// The tags of this image pipeline.
	Tags interface{} `pulumi:"tags"`
}

func LookupImagePipelineOutput(ctx *pulumi.Context, args LookupImagePipelineOutputArgs, opts ...pulumi.InvokeOption) LookupImagePipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImagePipelineResult, error) {
			args := v.(LookupImagePipelineArgs)
			r, err := LookupImagePipeline(ctx, &args, opts...)
			return *r, err
		}).(LookupImagePipelineResultOutput)
}

type LookupImagePipelineOutputArgs struct {
	// The Amazon Resource Name (ARN) of the image pipeline.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupImagePipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImagePipelineArgs)(nil)).Elem()
}

type LookupImagePipelineResultOutput struct{ *pulumi.OutputState }

func (LookupImagePipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImagePipelineResult)(nil)).Elem()
}

func (o LookupImagePipelineResultOutput) ToLookupImagePipelineResultOutput() LookupImagePipelineResultOutput {
	return o
}

func (o LookupImagePipelineResultOutput) ToLookupImagePipelineResultOutputWithContext(ctx context.Context) LookupImagePipelineResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the image pipeline.
func (o LookupImagePipelineResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.
func (o LookupImagePipelineResultOutput) ContainerRecipeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) *string { return v.ContainerRecipeArn }).(pulumi.StringPtrOutput)
}

// The description of the image pipeline.
func (o LookupImagePipelineResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.
func (o LookupImagePipelineResultOutput) DistributionConfigurationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) *string { return v.DistributionConfigurationArn }).(pulumi.StringPtrOutput)
}

// Collects additional information about the image being created, including the operating system (OS) version and package list.
func (o LookupImagePipelineResultOutput) EnhancedImageMetadataEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) *bool { return v.EnhancedImageMetadataEnabled }).(pulumi.BoolPtrOutput)
}

// The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.
func (o LookupImagePipelineResultOutput) ImageRecipeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) *string { return v.ImageRecipeArn }).(pulumi.StringPtrOutput)
}

// The image tests configuration of the image pipeline.
func (o LookupImagePipelineResultOutput) ImageTestsConfiguration() ImagePipelineImageTestsConfigurationPtrOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) *ImagePipelineImageTestsConfiguration {
		return v.ImageTestsConfiguration
	}).(ImagePipelineImageTestsConfigurationPtrOutput)
}

// The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.
func (o LookupImagePipelineResultOutput) InfrastructureConfigurationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) *string { return v.InfrastructureConfigurationArn }).(pulumi.StringPtrOutput)
}

// The schedule of the image pipeline.
func (o LookupImagePipelineResultOutput) Schedule() ImagePipelineSchedulePtrOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) *ImagePipelineSchedule { return v.Schedule }).(ImagePipelineSchedulePtrOutput)
}

// The status of the image pipeline.
func (o LookupImagePipelineResultOutput) Status() ImagePipelineStatusPtrOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) *ImagePipelineStatus { return v.Status }).(ImagePipelineStatusPtrOutput)
}

// The tags of this image pipeline.
func (o LookupImagePipelineResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImagePipelineResultOutput{})
}
