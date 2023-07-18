// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppStream::AppBlock
type AppBlock struct {
	pulumi.CustomResourceState

	Arn                    pulumi.StringOutput            `pulumi:"arn"`
	CreatedTime            pulumi.StringOutput            `pulumi:"createdTime"`
	Description            pulumi.StringPtrOutput         `pulumi:"description"`
	DisplayName            pulumi.StringPtrOutput         `pulumi:"displayName"`
	Name                   pulumi.StringOutput            `pulumi:"name"`
	PackagingType          pulumi.StringPtrOutput         `pulumi:"packagingType"`
	PostSetupScriptDetails AppBlockScriptDetailsPtrOutput `pulumi:"postSetupScriptDetails"`
	SetupScriptDetails     AppBlockScriptDetailsPtrOutput `pulumi:"setupScriptDetails"`
	SourceS3Location       AppBlockS3LocationOutput       `pulumi:"sourceS3Location"`
	Tags                   AppBlockTagArrayOutput         `pulumi:"tags"`
}

// NewAppBlock registers a new resource with the given unique name, arguments, and options.
func NewAppBlock(ctx *pulumi.Context,
	name string, args *AppBlockArgs, opts ...pulumi.ResourceOption) (*AppBlock, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceS3Location == nil {
		return nil, errors.New("invalid value for required argument 'SourceS3Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AppBlock
	err := ctx.RegisterResource("aws-native:appstream:AppBlock", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppBlock gets an existing AppBlock resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppBlock(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppBlockState, opts ...pulumi.ResourceOption) (*AppBlock, error) {
	var resource AppBlock
	err := ctx.ReadResource("aws-native:appstream:AppBlock", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppBlock resources.
type appBlockState struct {
}

type AppBlockState struct {
}

func (AppBlockState) ElementType() reflect.Type {
	return reflect.TypeOf((*appBlockState)(nil)).Elem()
}

type appBlockArgs struct {
	Description            *string                `pulumi:"description"`
	DisplayName            *string                `pulumi:"displayName"`
	Name                   *string                `pulumi:"name"`
	PackagingType          *string                `pulumi:"packagingType"`
	PostSetupScriptDetails *AppBlockScriptDetails `pulumi:"postSetupScriptDetails"`
	SetupScriptDetails     *AppBlockScriptDetails `pulumi:"setupScriptDetails"`
	SourceS3Location       AppBlockS3Location     `pulumi:"sourceS3Location"`
	Tags                   []AppBlockTag          `pulumi:"tags"`
}

// The set of arguments for constructing a AppBlock resource.
type AppBlockArgs struct {
	Description            pulumi.StringPtrInput
	DisplayName            pulumi.StringPtrInput
	Name                   pulumi.StringPtrInput
	PackagingType          pulumi.StringPtrInput
	PostSetupScriptDetails AppBlockScriptDetailsPtrInput
	SetupScriptDetails     AppBlockScriptDetailsPtrInput
	SourceS3Location       AppBlockS3LocationInput
	Tags                   AppBlockTagArrayInput
}

func (AppBlockArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appBlockArgs)(nil)).Elem()
}

type AppBlockInput interface {
	pulumi.Input

	ToAppBlockOutput() AppBlockOutput
	ToAppBlockOutputWithContext(ctx context.Context) AppBlockOutput
}

func (*AppBlock) ElementType() reflect.Type {
	return reflect.TypeOf((**AppBlock)(nil)).Elem()
}

func (i *AppBlock) ToAppBlockOutput() AppBlockOutput {
	return i.ToAppBlockOutputWithContext(context.Background())
}

func (i *AppBlock) ToAppBlockOutputWithContext(ctx context.Context) AppBlockOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppBlockOutput)
}

type AppBlockOutput struct{ *pulumi.OutputState }

func (AppBlockOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppBlock)(nil)).Elem()
}

func (o AppBlockOutput) ToAppBlockOutput() AppBlockOutput {
	return o
}

func (o AppBlockOutput) ToAppBlockOutputWithContext(ctx context.Context) AppBlockOutput {
	return o
}

func (o AppBlockOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *AppBlock) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o AppBlockOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AppBlock) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o AppBlockOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppBlock) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AppBlockOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppBlock) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o AppBlockOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppBlock) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppBlockOutput) PackagingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppBlock) pulumi.StringPtrOutput { return v.PackagingType }).(pulumi.StringPtrOutput)
}

func (o AppBlockOutput) PostSetupScriptDetails() AppBlockScriptDetailsPtrOutput {
	return o.ApplyT(func(v *AppBlock) AppBlockScriptDetailsPtrOutput { return v.PostSetupScriptDetails }).(AppBlockScriptDetailsPtrOutput)
}

func (o AppBlockOutput) SetupScriptDetails() AppBlockScriptDetailsPtrOutput {
	return o.ApplyT(func(v *AppBlock) AppBlockScriptDetailsPtrOutput { return v.SetupScriptDetails }).(AppBlockScriptDetailsPtrOutput)
}

func (o AppBlockOutput) SourceS3Location() AppBlockS3LocationOutput {
	return o.ApplyT(func(v *AppBlock) AppBlockS3LocationOutput { return v.SourceS3Location }).(AppBlockS3LocationOutput)
}

func (o AppBlockOutput) Tags() AppBlockTagArrayOutput {
	return o.ApplyT(func(v *AppBlock) AppBlockTagArrayOutput { return v.Tags }).(AppBlockTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppBlockInput)(nil)).Elem(), &AppBlock{})
	pulumi.RegisterOutputType(AppBlockOutput{})
}
