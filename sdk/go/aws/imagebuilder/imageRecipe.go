// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::ImageBuilder::ImageRecipe
type ImageRecipe struct {
	pulumi.CustomResourceState

	// Specify additional settings and launch scripts for your build instances.
	AdditionalInstanceConfiguration ImageRecipeAdditionalInstanceConfigurationPtrOutput `pulumi:"additionalInstanceConfiguration"`
	// The Amazon Resource Name (ARN) of the image recipe.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The block device mappings to apply when creating images from this recipe.
	BlockDeviceMappings ImageRecipeInstanceBlockDeviceMappingArrayOutput `pulumi:"blockDeviceMappings"`
	// The components of the image recipe.
	Components ImageRecipeComponentConfigurationArrayOutput `pulumi:"components"`
	// The description of the image recipe.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the image recipe.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent image of the image recipe.
	ParentImage pulumi.StringOutput `pulumi:"parentImage"`
	// The tags of the image recipe.
	Tags pulumi.AnyOutput `pulumi:"tags"`
	// The version of the image recipe.
	Version pulumi.StringOutput `pulumi:"version"`
	// The working directory to be used during build and test workflows.
	WorkingDirectory pulumi.StringPtrOutput `pulumi:"workingDirectory"`
}

// NewImageRecipe registers a new resource with the given unique name, arguments, and options.
func NewImageRecipe(ctx *pulumi.Context,
	name string, args *ImageRecipeArgs, opts ...pulumi.ResourceOption) (*ImageRecipe, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Components == nil {
		return nil, errors.New("invalid value for required argument 'Components'")
	}
	if args.ParentImage == nil {
		return nil, errors.New("invalid value for required argument 'ParentImage'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource ImageRecipe
	err := ctx.RegisterResource("aws-native:imagebuilder:ImageRecipe", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageRecipe gets an existing ImageRecipe resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageRecipe(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageRecipeState, opts ...pulumi.ResourceOption) (*ImageRecipe, error) {
	var resource ImageRecipe
	err := ctx.ReadResource("aws-native:imagebuilder:ImageRecipe", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageRecipe resources.
type imageRecipeState struct {
}

type ImageRecipeState struct {
}

func (ImageRecipeState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageRecipeState)(nil)).Elem()
}

type imageRecipeArgs struct {
	// Specify additional settings and launch scripts for your build instances.
	AdditionalInstanceConfiguration *ImageRecipeAdditionalInstanceConfiguration `pulumi:"additionalInstanceConfiguration"`
	// The block device mappings to apply when creating images from this recipe.
	BlockDeviceMappings []ImageRecipeInstanceBlockDeviceMapping `pulumi:"blockDeviceMappings"`
	// The components of the image recipe.
	Components []ImageRecipeComponentConfiguration `pulumi:"components"`
	// The description of the image recipe.
	Description *string `pulumi:"description"`
	// The name of the image recipe.
	Name *string `pulumi:"name"`
	// The parent image of the image recipe.
	ParentImage string `pulumi:"parentImage"`
	// The tags of the image recipe.
	Tags interface{} `pulumi:"tags"`
	// The version of the image recipe.
	Version string `pulumi:"version"`
	// The working directory to be used during build and test workflows.
	WorkingDirectory *string `pulumi:"workingDirectory"`
}

// The set of arguments for constructing a ImageRecipe resource.
type ImageRecipeArgs struct {
	// Specify additional settings and launch scripts for your build instances.
	AdditionalInstanceConfiguration ImageRecipeAdditionalInstanceConfigurationPtrInput
	// The block device mappings to apply when creating images from this recipe.
	BlockDeviceMappings ImageRecipeInstanceBlockDeviceMappingArrayInput
	// The components of the image recipe.
	Components ImageRecipeComponentConfigurationArrayInput
	// The description of the image recipe.
	Description pulumi.StringPtrInput
	// The name of the image recipe.
	Name pulumi.StringPtrInput
	// The parent image of the image recipe.
	ParentImage pulumi.StringInput
	// The tags of the image recipe.
	Tags pulumi.Input
	// The version of the image recipe.
	Version pulumi.StringInput
	// The working directory to be used during build and test workflows.
	WorkingDirectory pulumi.StringPtrInput
}

func (ImageRecipeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageRecipeArgs)(nil)).Elem()
}

type ImageRecipeInput interface {
	pulumi.Input

	ToImageRecipeOutput() ImageRecipeOutput
	ToImageRecipeOutputWithContext(ctx context.Context) ImageRecipeOutput
}

func (*ImageRecipe) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRecipe)(nil)).Elem()
}

func (i *ImageRecipe) ToImageRecipeOutput() ImageRecipeOutput {
	return i.ToImageRecipeOutputWithContext(context.Background())
}

func (i *ImageRecipe) ToImageRecipeOutputWithContext(ctx context.Context) ImageRecipeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRecipeOutput)
}

type ImageRecipeOutput struct{ *pulumi.OutputState }

func (ImageRecipeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRecipe)(nil)).Elem()
}

func (o ImageRecipeOutput) ToImageRecipeOutput() ImageRecipeOutput {
	return o
}

func (o ImageRecipeOutput) ToImageRecipeOutputWithContext(ctx context.Context) ImageRecipeOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageRecipeInput)(nil)).Elem(), &ImageRecipe{})
	pulumi.RegisterOutputType(ImageRecipeOutput{})
}
