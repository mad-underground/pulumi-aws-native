// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::CoreDefinitionVersion
//
// Deprecated: CoreDefinitionVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type CoreDefinitionVersion struct {
	pulumi.CustomResourceState

	CoreDefinitionId pulumi.StringOutput                  `pulumi:"coreDefinitionId"`
	Cores            CoreDefinitionVersionCoreArrayOutput `pulumi:"cores"`
}

// NewCoreDefinitionVersion registers a new resource with the given unique name, arguments, and options.
func NewCoreDefinitionVersion(ctx *pulumi.Context,
	name string, args *CoreDefinitionVersionArgs, opts ...pulumi.ResourceOption) (*CoreDefinitionVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CoreDefinitionId == nil {
		return nil, errors.New("invalid value for required argument 'CoreDefinitionId'")
	}
	if args.Cores == nil {
		return nil, errors.New("invalid value for required argument 'Cores'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CoreDefinitionVersion
	err := ctx.RegisterResource("aws-native:greengrass:CoreDefinitionVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCoreDefinitionVersion gets an existing CoreDefinitionVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCoreDefinitionVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CoreDefinitionVersionState, opts ...pulumi.ResourceOption) (*CoreDefinitionVersion, error) {
	var resource CoreDefinitionVersion
	err := ctx.ReadResource("aws-native:greengrass:CoreDefinitionVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CoreDefinitionVersion resources.
type coreDefinitionVersionState struct {
}

type CoreDefinitionVersionState struct {
}

func (CoreDefinitionVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*coreDefinitionVersionState)(nil)).Elem()
}

type coreDefinitionVersionArgs struct {
	CoreDefinitionId string                      `pulumi:"coreDefinitionId"`
	Cores            []CoreDefinitionVersionCore `pulumi:"cores"`
}

// The set of arguments for constructing a CoreDefinitionVersion resource.
type CoreDefinitionVersionArgs struct {
	CoreDefinitionId pulumi.StringInput
	Cores            CoreDefinitionVersionCoreArrayInput
}

func (CoreDefinitionVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*coreDefinitionVersionArgs)(nil)).Elem()
}

type CoreDefinitionVersionInput interface {
	pulumi.Input

	ToCoreDefinitionVersionOutput() CoreDefinitionVersionOutput
	ToCoreDefinitionVersionOutputWithContext(ctx context.Context) CoreDefinitionVersionOutput
}

func (*CoreDefinitionVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreDefinitionVersion)(nil)).Elem()
}

func (i *CoreDefinitionVersion) ToCoreDefinitionVersionOutput() CoreDefinitionVersionOutput {
	return i.ToCoreDefinitionVersionOutputWithContext(context.Background())
}

func (i *CoreDefinitionVersion) ToCoreDefinitionVersionOutputWithContext(ctx context.Context) CoreDefinitionVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreDefinitionVersionOutput)
}

type CoreDefinitionVersionOutput struct{ *pulumi.OutputState }

func (CoreDefinitionVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreDefinitionVersion)(nil)).Elem()
}

func (o CoreDefinitionVersionOutput) ToCoreDefinitionVersionOutput() CoreDefinitionVersionOutput {
	return o
}

func (o CoreDefinitionVersionOutput) ToCoreDefinitionVersionOutputWithContext(ctx context.Context) CoreDefinitionVersionOutput {
	return o
}

func (o CoreDefinitionVersionOutput) CoreDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v *CoreDefinitionVersion) pulumi.StringOutput { return v.CoreDefinitionId }).(pulumi.StringOutput)
}

func (o CoreDefinitionVersionOutput) Cores() CoreDefinitionVersionCoreArrayOutput {
	return o.ApplyT(func(v *CoreDefinitionVersion) CoreDefinitionVersionCoreArrayOutput { return v.Cores }).(CoreDefinitionVersionCoreArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CoreDefinitionVersionInput)(nil)).Elem(), &CoreDefinitionVersion{})
	pulumi.RegisterOutputType(CoreDefinitionVersionOutput{})
}
