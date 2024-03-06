// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppConfig::Extension
type Extension struct {
	pulumi.CustomResourceState

	Actions ExtensionActionArrayMapOutput `pulumi:"actions"`
	Arn     pulumi.StringOutput           `pulumi:"arn"`
	AwsId   pulumi.StringOutput           `pulumi:"awsId"`
	// Description of the extension.
	Description         pulumi.StringPtrOutput `pulumi:"description"`
	LatestVersionNumber pulumi.IntPtrOutput    `pulumi:"latestVersionNumber"`
	// Name of the extension.
	Name       pulumi.StringOutput         `pulumi:"name"`
	Parameters ExtensionParameterMapOutput `pulumi:"parameters"`
	// An array of key-value tags to apply to this resource.
	Tags          aws.CreateOnlyTagArrayOutput `pulumi:"tags"`
	VersionNumber pulumi.IntOutput             `pulumi:"versionNumber"`
}

// NewExtension registers a new resource with the given unique name, arguments, and options.
func NewExtension(ctx *pulumi.Context,
	name string, args *ExtensionArgs, opts ...pulumi.ResourceOption) (*Extension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
		"tags[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Extension
	err := ctx.RegisterResource("aws-native:appconfig:Extension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtension gets an existing Extension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionState, opts ...pulumi.ResourceOption) (*Extension, error) {
	var resource Extension
	err := ctx.ReadResource("aws-native:appconfig:Extension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Extension resources.
type extensionState struct {
}

type ExtensionState struct {
}

func (ExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionState)(nil)).Elem()
}

type extensionArgs struct {
	Actions map[string][]ExtensionAction `pulumi:"actions"`
	// Description of the extension.
	Description         *string `pulumi:"description"`
	LatestVersionNumber *int    `pulumi:"latestVersionNumber"`
	// Name of the extension.
	Name       *string                       `pulumi:"name"`
	Parameters map[string]ExtensionParameter `pulumi:"parameters"`
	// An array of key-value tags to apply to this resource.
	Tags []aws.CreateOnlyTag `pulumi:"tags"`
}

// The set of arguments for constructing a Extension resource.
type ExtensionArgs struct {
	Actions ExtensionActionArrayMapInput
	// Description of the extension.
	Description         pulumi.StringPtrInput
	LatestVersionNumber pulumi.IntPtrInput
	// Name of the extension.
	Name       pulumi.StringPtrInput
	Parameters ExtensionParameterMapInput
	// An array of key-value tags to apply to this resource.
	Tags aws.CreateOnlyTagArrayInput
}

func (ExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionArgs)(nil)).Elem()
}

type ExtensionInput interface {
	pulumi.Input

	ToExtensionOutput() ExtensionOutput
	ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput
}

func (*Extension) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (i *Extension) ToExtensionOutput() ExtensionOutput {
	return i.ToExtensionOutputWithContext(context.Background())
}

func (i *Extension) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionOutput)
}

type ExtensionOutput struct{ *pulumi.OutputState }

func (ExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (o ExtensionOutput) ToExtensionOutput() ExtensionOutput {
	return o
}

func (o ExtensionOutput) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return o
}

func (o ExtensionOutput) Actions() ExtensionActionArrayMapOutput {
	return o.ApplyT(func(v *Extension) ExtensionActionArrayMapOutput { return v.Actions }).(ExtensionActionArrayMapOutput)
}

func (o ExtensionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ExtensionOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// Description of the extension.
func (o ExtensionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ExtensionOutput) LatestVersionNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.IntPtrOutput { return v.LatestVersionNumber }).(pulumi.IntPtrOutput)
}

// Name of the extension.
func (o ExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ExtensionOutput) Parameters() ExtensionParameterMapOutput {
	return o.ApplyT(func(v *Extension) ExtensionParameterMapOutput { return v.Parameters }).(ExtensionParameterMapOutput)
}

// An array of key-value tags to apply to this resource.
func (o ExtensionOutput) Tags() aws.CreateOnlyTagArrayOutput {
	return o.ApplyT(func(v *Extension) aws.CreateOnlyTagArrayOutput { return v.Tags }).(aws.CreateOnlyTagArrayOutput)
}

func (o ExtensionOutput) VersionNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Extension) pulumi.IntOutput { return v.VersionNumber }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensionInput)(nil)).Elem(), &Extension{})
	pulumi.RegisterOutputType(ExtensionOutput{})
}
