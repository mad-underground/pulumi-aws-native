// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::ImageBuilder::Image
type Image struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the image.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.
	ContainerRecipeArn pulumi.StringPtrOutput `pulumi:"containerRecipeArn"`
	// The Amazon Resource Name (ARN) of the distribution configuration.
	DistributionConfigurationArn pulumi.StringPtrOutput `pulumi:"distributionConfigurationArn"`
	// Collects additional information about the image being created, including the operating system (OS) version and package list.
	EnhancedImageMetadataEnabled pulumi.BoolPtrOutput `pulumi:"enhancedImageMetadataEnabled"`
	// The execution role name/ARN for the image build, if provided
	ExecutionRole pulumi.StringPtrOutput `pulumi:"executionRole"`
	// The AMI ID of the EC2 AMI in current region.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.
	ImageRecipeArn pulumi.StringPtrOutput `pulumi:"imageRecipeArn"`
	// Contains settings for vulnerability scans.
	ImageScanningConfiguration ImageScanningConfigurationPtrOutput `pulumi:"imageScanningConfiguration"`
	// The image tests configuration used when creating this image.
	ImageTestsConfiguration ImageTestsConfigurationPtrOutput `pulumi:"imageTestsConfiguration"`
	// URI for containers created in current Region with default ECR image tag
	ImageUri pulumi.StringOutput `pulumi:"imageUri"`
	// The Amazon Resource Name (ARN) of the infrastructure configuration.
	InfrastructureConfigurationArn pulumi.StringPtrOutput `pulumi:"infrastructureConfigurationArn"`
	// The name of the image.
	Name pulumi.StringOutput `pulumi:"name"`
	// The tags associated with the image.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Workflows to define the image build process
	Workflows ImageWorkflowConfigurationArrayOutput `pulumi:"workflows"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		args = &ImageArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"containerRecipeArn",
		"distributionConfigurationArn",
		"enhancedImageMetadataEnabled",
		"imageRecipeArn",
		"imageScanningConfiguration",
		"imageTestsConfiguration",
		"infrastructureConfigurationArn",
		"tags.*",
		"workflows[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Image
	err := ctx.RegisterResource("aws-native:imagebuilder:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("aws-native:imagebuilder:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
}

type ImageState struct {
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.
	ContainerRecipeArn *string `pulumi:"containerRecipeArn"`
	// The Amazon Resource Name (ARN) of the distribution configuration.
	DistributionConfigurationArn *string `pulumi:"distributionConfigurationArn"`
	// Collects additional information about the image being created, including the operating system (OS) version and package list.
	EnhancedImageMetadataEnabled *bool `pulumi:"enhancedImageMetadataEnabled"`
	// The execution role name/ARN for the image build, if provided
	ExecutionRole *string `pulumi:"executionRole"`
	// The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.
	ImageRecipeArn *string `pulumi:"imageRecipeArn"`
	// Contains settings for vulnerability scans.
	ImageScanningConfiguration *ImageScanningConfiguration `pulumi:"imageScanningConfiguration"`
	// The image tests configuration used when creating this image.
	ImageTestsConfiguration *ImageTestsConfiguration `pulumi:"imageTestsConfiguration"`
	// The Amazon Resource Name (ARN) of the infrastructure configuration.
	InfrastructureConfigurationArn *string `pulumi:"infrastructureConfigurationArn"`
	// The tags associated with the image.
	Tags map[string]string `pulumi:"tags"`
	// Workflows to define the image build process
	Workflows []ImageWorkflowConfiguration `pulumi:"workflows"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.
	ContainerRecipeArn pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the distribution configuration.
	DistributionConfigurationArn pulumi.StringPtrInput
	// Collects additional information about the image being created, including the operating system (OS) version and package list.
	EnhancedImageMetadataEnabled pulumi.BoolPtrInput
	// The execution role name/ARN for the image build, if provided
	ExecutionRole pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.
	ImageRecipeArn pulumi.StringPtrInput
	// Contains settings for vulnerability scans.
	ImageScanningConfiguration ImageScanningConfigurationPtrInput
	// The image tests configuration used when creating this image.
	ImageTestsConfiguration ImageTestsConfigurationPtrInput
	// The Amazon Resource Name (ARN) of the infrastructure configuration.
	InfrastructureConfigurationArn pulumi.StringPtrInput
	// The tags associated with the image.
	Tags pulumi.StringMapInput
	// Workflows to define the image build process
	Workflows ImageWorkflowConfigurationArrayInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

type ImageOutput struct{ *pulumi.OutputState }

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

// The Amazon Resource Name (ARN) of the image.
func (o ImageOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.
func (o ImageOutput) ContainerRecipeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.ContainerRecipeArn }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the distribution configuration.
func (o ImageOutput) DistributionConfigurationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.DistributionConfigurationArn }).(pulumi.StringPtrOutput)
}

// Collects additional information about the image being created, including the operating system (OS) version and package list.
func (o ImageOutput) EnhancedImageMetadataEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.BoolPtrOutput { return v.EnhancedImageMetadataEnabled }).(pulumi.BoolPtrOutput)
}

// The execution role name/ARN for the image build, if provided
func (o ImageOutput) ExecutionRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.ExecutionRole }).(pulumi.StringPtrOutput)
}

// The AMI ID of the EC2 AMI in current region.
func (o ImageOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.
func (o ImageOutput) ImageRecipeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.ImageRecipeArn }).(pulumi.StringPtrOutput)
}

// Contains settings for vulnerability scans.
func (o ImageOutput) ImageScanningConfiguration() ImageScanningConfigurationPtrOutput {
	return o.ApplyT(func(v *Image) ImageScanningConfigurationPtrOutput { return v.ImageScanningConfiguration }).(ImageScanningConfigurationPtrOutput)
}

// The image tests configuration used when creating this image.
func (o ImageOutput) ImageTestsConfiguration() ImageTestsConfigurationPtrOutput {
	return o.ApplyT(func(v *Image) ImageTestsConfigurationPtrOutput { return v.ImageTestsConfiguration }).(ImageTestsConfigurationPtrOutput)
}

// URI for containers created in current Region with default ECR image tag
func (o ImageOutput) ImageUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ImageUri }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the infrastructure configuration.
func (o ImageOutput) InfrastructureConfigurationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.InfrastructureConfigurationArn }).(pulumi.StringPtrOutput)
}

// The name of the image.
func (o ImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The tags associated with the image.
func (o ImageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Image) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Workflows to define the image build process
func (o ImageOutput) Workflows() ImageWorkflowConfigurationArrayOutput {
	return o.ApplyT(func(v *Image) ImageWorkflowConfigurationArrayOutput { return v.Workflows }).(ImageWorkflowConfigurationArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInput)(nil)).Elem(), &Image{})
	pulumi.RegisterOutputType(ImageOutput{})
}
