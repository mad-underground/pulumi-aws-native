// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package nimblestudio

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html
type StreamingImage struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-ec2imageid
	Ec2ImageId pulumi.StringOutput      `pulumi:"ec2ImageId"`
	EulaIds    pulumi.StringArrayOutput `pulumi:"eulaIds"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-name
	Name             pulumi.StringOutput `pulumi:"name"`
	Owner            pulumi.StringOutput `pulumi:"owner"`
	Platform         pulumi.StringOutput `pulumi:"platform"`
	StreamingImageId pulumi.StringOutput `pulumi:"streamingImageId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-studioid
	StudioId pulumi.StringOutput `pulumi:"studioId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewStreamingImage registers a new resource with the given unique name, arguments, and options.
func NewStreamingImage(ctx *pulumi.Context,
	name string, args *StreamingImageArgs, opts ...pulumi.ResourceOption) (*StreamingImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ec2ImageId == nil {
		return nil, errors.New("invalid value for required argument 'Ec2ImageId'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.StudioId == nil {
		return nil, errors.New("invalid value for required argument 'StudioId'")
	}
	var resource StreamingImage
	err := ctx.RegisterResource("aws-native:nimblestudio:StreamingImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingImage gets an existing StreamingImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingImageState, opts ...pulumi.ResourceOption) (*StreamingImage, error) {
	var resource StreamingImage
	err := ctx.ReadResource("aws-native:nimblestudio:StreamingImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingImage resources.
type streamingImageState struct {
}

type StreamingImageState struct {
}

func (StreamingImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingImageState)(nil)).Elem()
}

type streamingImageArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-ec2imageid
	Ec2ImageId string `pulumi:"ec2ImageId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-name
	Name string `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-studioid
	StudioId string `pulumi:"studioId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StreamingImage resource.
type StreamingImageArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-ec2imageid
	Ec2ImageId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-name
	Name pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-studioid
	StudioId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-streamingimage.html#cfn-nimblestudio-streamingimage-tags
	Tags pulumi.StringMapInput
}

func (StreamingImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingImageArgs)(nil)).Elem()
}

type StreamingImageInput interface {
	pulumi.Input

	ToStreamingImageOutput() StreamingImageOutput
	ToStreamingImageOutputWithContext(ctx context.Context) StreamingImageOutput
}

func (*StreamingImage) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingImage)(nil))
}

func (i *StreamingImage) ToStreamingImageOutput() StreamingImageOutput {
	return i.ToStreamingImageOutputWithContext(context.Background())
}

func (i *StreamingImage) ToStreamingImageOutputWithContext(ctx context.Context) StreamingImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingImageOutput)
}

type StreamingImageOutput struct{ *pulumi.OutputState }

func (StreamingImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingImage)(nil))
}

func (o StreamingImageOutput) ToStreamingImageOutput() StreamingImageOutput {
	return o
}

func (o StreamingImageOutput) ToStreamingImageOutputWithContext(ctx context.Context) StreamingImageOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StreamingImageOutput{})
}