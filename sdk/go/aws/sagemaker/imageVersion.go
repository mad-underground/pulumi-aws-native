// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::ImageVersion
type ImageVersion struct {
	pulumi.CustomResourceState

	Alias           pulumi.StringPtrOutput              `pulumi:"alias"`
	Aliases         pulumi.StringArrayOutput            `pulumi:"aliases"`
	BaseImage       pulumi.StringOutput                 `pulumi:"baseImage"`
	ContainerImage  pulumi.StringOutput                 `pulumi:"containerImage"`
	Horovod         pulumi.BoolPtrOutput                `pulumi:"horovod"`
	ImageArn        pulumi.StringOutput                 `pulumi:"imageArn"`
	ImageName       pulumi.StringOutput                 `pulumi:"imageName"`
	ImageVersionArn pulumi.StringOutput                 `pulumi:"imageVersionArn"`
	JobType         ImageVersionJobTypePtrOutput        `pulumi:"jobType"`
	MlFramework     pulumi.StringPtrOutput              `pulumi:"mlFramework"`
	Processor       ImageVersionProcessorPtrOutput      `pulumi:"processor"`
	ProgrammingLang pulumi.StringPtrOutput              `pulumi:"programmingLang"`
	ReleaseNotes    pulumi.StringPtrOutput              `pulumi:"releaseNotes"`
	VendorGuidance  ImageVersionVendorGuidancePtrOutput `pulumi:"vendorGuidance"`
	Version         pulumi.IntOutput                    `pulumi:"version"`
}

// NewImageVersion registers a new resource with the given unique name, arguments, and options.
func NewImageVersion(ctx *pulumi.Context,
	name string, args *ImageVersionArgs, opts ...pulumi.ResourceOption) (*ImageVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaseImage == nil {
		return nil, errors.New("invalid value for required argument 'BaseImage'")
	}
	if args.ImageName == nil {
		return nil, errors.New("invalid value for required argument 'ImageName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"baseImage",
		"imageName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ImageVersion
	err := ctx.RegisterResource("aws-native:sagemaker:ImageVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageVersion gets an existing ImageVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageVersionState, opts ...pulumi.ResourceOption) (*ImageVersion, error) {
	var resource ImageVersion
	err := ctx.ReadResource("aws-native:sagemaker:ImageVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageVersion resources.
type imageVersionState struct {
}

type ImageVersionState struct {
}

func (ImageVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageVersionState)(nil)).Elem()
}

type imageVersionArgs struct {
	Alias           *string                     `pulumi:"alias"`
	Aliases         []string                    `pulumi:"aliases"`
	BaseImage       string                      `pulumi:"baseImage"`
	Horovod         *bool                       `pulumi:"horovod"`
	ImageName       string                      `pulumi:"imageName"`
	JobType         *ImageVersionJobType        `pulumi:"jobType"`
	MlFramework     *string                     `pulumi:"mlFramework"`
	Processor       *ImageVersionProcessor      `pulumi:"processor"`
	ProgrammingLang *string                     `pulumi:"programmingLang"`
	ReleaseNotes    *string                     `pulumi:"releaseNotes"`
	VendorGuidance  *ImageVersionVendorGuidance `pulumi:"vendorGuidance"`
}

// The set of arguments for constructing a ImageVersion resource.
type ImageVersionArgs struct {
	Alias           pulumi.StringPtrInput
	Aliases         pulumi.StringArrayInput
	BaseImage       pulumi.StringInput
	Horovod         pulumi.BoolPtrInput
	ImageName       pulumi.StringInput
	JobType         ImageVersionJobTypePtrInput
	MlFramework     pulumi.StringPtrInput
	Processor       ImageVersionProcessorPtrInput
	ProgrammingLang pulumi.StringPtrInput
	ReleaseNotes    pulumi.StringPtrInput
	VendorGuidance  ImageVersionVendorGuidancePtrInput
}

func (ImageVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageVersionArgs)(nil)).Elem()
}

type ImageVersionInput interface {
	pulumi.Input

	ToImageVersionOutput() ImageVersionOutput
	ToImageVersionOutputWithContext(ctx context.Context) ImageVersionOutput
}

func (*ImageVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageVersion)(nil)).Elem()
}

func (i *ImageVersion) ToImageVersionOutput() ImageVersionOutput {
	return i.ToImageVersionOutputWithContext(context.Background())
}

func (i *ImageVersion) ToImageVersionOutputWithContext(ctx context.Context) ImageVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageVersionOutput)
}

type ImageVersionOutput struct{ *pulumi.OutputState }

func (ImageVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageVersion)(nil)).Elem()
}

func (o ImageVersionOutput) ToImageVersionOutput() ImageVersionOutput {
	return o
}

func (o ImageVersionOutput) ToImageVersionOutputWithContext(ctx context.Context) ImageVersionOutput {
	return o
}

func (o ImageVersionOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageVersion) pulumi.StringPtrOutput { return v.Alias }).(pulumi.StringPtrOutput)
}

func (o ImageVersionOutput) Aliases() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ImageVersion) pulumi.StringArrayOutput { return v.Aliases }).(pulumi.StringArrayOutput)
}

func (o ImageVersionOutput) BaseImage() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageVersion) pulumi.StringOutput { return v.BaseImage }).(pulumi.StringOutput)
}

func (o ImageVersionOutput) ContainerImage() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageVersion) pulumi.StringOutput { return v.ContainerImage }).(pulumi.StringOutput)
}

func (o ImageVersionOutput) Horovod() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ImageVersion) pulumi.BoolPtrOutput { return v.Horovod }).(pulumi.BoolPtrOutput)
}

func (o ImageVersionOutput) ImageArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageVersion) pulumi.StringOutput { return v.ImageArn }).(pulumi.StringOutput)
}

func (o ImageVersionOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageVersion) pulumi.StringOutput { return v.ImageName }).(pulumi.StringOutput)
}

func (o ImageVersionOutput) ImageVersionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageVersion) pulumi.StringOutput { return v.ImageVersionArn }).(pulumi.StringOutput)
}

func (o ImageVersionOutput) JobType() ImageVersionJobTypePtrOutput {
	return o.ApplyT(func(v *ImageVersion) ImageVersionJobTypePtrOutput { return v.JobType }).(ImageVersionJobTypePtrOutput)
}

func (o ImageVersionOutput) MlFramework() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageVersion) pulumi.StringPtrOutput { return v.MlFramework }).(pulumi.StringPtrOutput)
}

func (o ImageVersionOutput) Processor() ImageVersionProcessorPtrOutput {
	return o.ApplyT(func(v *ImageVersion) ImageVersionProcessorPtrOutput { return v.Processor }).(ImageVersionProcessorPtrOutput)
}

func (o ImageVersionOutput) ProgrammingLang() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageVersion) pulumi.StringPtrOutput { return v.ProgrammingLang }).(pulumi.StringPtrOutput)
}

func (o ImageVersionOutput) ReleaseNotes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageVersion) pulumi.StringPtrOutput { return v.ReleaseNotes }).(pulumi.StringPtrOutput)
}

func (o ImageVersionOutput) VendorGuidance() ImageVersionVendorGuidancePtrOutput {
	return o.ApplyT(func(v *ImageVersion) ImageVersionVendorGuidancePtrOutput { return v.VendorGuidance }).(ImageVersionVendorGuidancePtrOutput)
}

func (o ImageVersionOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *ImageVersion) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageVersionInput)(nil)).Elem(), &ImageVersion{})
	pulumi.RegisterOutputType(ImageVersionOutput{})
}
