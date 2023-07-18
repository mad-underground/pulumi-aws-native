// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::CoreDefinition
//
// Deprecated: CoreDefinition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type CoreDefinition struct {
	pulumi.CustomResourceState

	Arn              pulumi.StringOutput                `pulumi:"arn"`
	InitialVersion   CoreDefinitionVersionTypePtrOutput `pulumi:"initialVersion"`
	LatestVersionArn pulumi.StringOutput                `pulumi:"latestVersionArn"`
	Name             pulumi.StringOutput                `pulumi:"name"`
	Tags             pulumi.AnyOutput                   `pulumi:"tags"`
}

// NewCoreDefinition registers a new resource with the given unique name, arguments, and options.
func NewCoreDefinition(ctx *pulumi.Context,
	name string, args *CoreDefinitionArgs, opts ...pulumi.ResourceOption) (*CoreDefinition, error) {
	if args == nil {
		args = &CoreDefinitionArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CoreDefinition
	err := ctx.RegisterResource("aws-native:greengrass:CoreDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCoreDefinition gets an existing CoreDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCoreDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CoreDefinitionState, opts ...pulumi.ResourceOption) (*CoreDefinition, error) {
	var resource CoreDefinition
	err := ctx.ReadResource("aws-native:greengrass:CoreDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CoreDefinition resources.
type coreDefinitionState struct {
}

type CoreDefinitionState struct {
}

func (CoreDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*coreDefinitionState)(nil)).Elem()
}

type coreDefinitionArgs struct {
	InitialVersion *CoreDefinitionVersionType `pulumi:"initialVersion"`
	Name           *string                    `pulumi:"name"`
	Tags           interface{}                `pulumi:"tags"`
}

// The set of arguments for constructing a CoreDefinition resource.
type CoreDefinitionArgs struct {
	InitialVersion CoreDefinitionVersionTypePtrInput
	Name           pulumi.StringPtrInput
	Tags           pulumi.Input
}

func (CoreDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*coreDefinitionArgs)(nil)).Elem()
}

type CoreDefinitionInput interface {
	pulumi.Input

	ToCoreDefinitionOutput() CoreDefinitionOutput
	ToCoreDefinitionOutputWithContext(ctx context.Context) CoreDefinitionOutput
}

func (*CoreDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreDefinition)(nil)).Elem()
}

func (i *CoreDefinition) ToCoreDefinitionOutput() CoreDefinitionOutput {
	return i.ToCoreDefinitionOutputWithContext(context.Background())
}

func (i *CoreDefinition) ToCoreDefinitionOutputWithContext(ctx context.Context) CoreDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreDefinitionOutput)
}

type CoreDefinitionOutput struct{ *pulumi.OutputState }

func (CoreDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreDefinition)(nil)).Elem()
}

func (o CoreDefinitionOutput) ToCoreDefinitionOutput() CoreDefinitionOutput {
	return o
}

func (o CoreDefinitionOutput) ToCoreDefinitionOutputWithContext(ctx context.Context) CoreDefinitionOutput {
	return o
}

func (o CoreDefinitionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CoreDefinition) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o CoreDefinitionOutput) InitialVersion() CoreDefinitionVersionTypePtrOutput {
	return o.ApplyT(func(v *CoreDefinition) CoreDefinitionVersionTypePtrOutput { return v.InitialVersion }).(CoreDefinitionVersionTypePtrOutput)
}

func (o CoreDefinitionOutput) LatestVersionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *CoreDefinition) pulumi.StringOutput { return v.LatestVersionArn }).(pulumi.StringOutput)
}

func (o CoreDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CoreDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CoreDefinitionOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *CoreDefinition) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CoreDefinitionInput)(nil)).Elem(), &CoreDefinition{})
	pulumi.RegisterOutputType(CoreDefinitionOutput{})
}
