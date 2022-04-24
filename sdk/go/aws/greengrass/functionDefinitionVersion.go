// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::FunctionDefinitionVersion
//
// Deprecated: FunctionDefinitionVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type FunctionDefinitionVersion struct {
	pulumi.CustomResourceState

	DefaultConfig        FunctionDefinitionVersionDefaultConfigPtrOutput `pulumi:"defaultConfig"`
	FunctionDefinitionId pulumi.StringOutput                             `pulumi:"functionDefinitionId"`
	Functions            FunctionDefinitionVersionFunctionArrayOutput    `pulumi:"functions"`
}

// NewFunctionDefinitionVersion registers a new resource with the given unique name, arguments, and options.
func NewFunctionDefinitionVersion(ctx *pulumi.Context,
	name string, args *FunctionDefinitionVersionArgs, opts ...pulumi.ResourceOption) (*FunctionDefinitionVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FunctionDefinitionId == nil {
		return nil, errors.New("invalid value for required argument 'FunctionDefinitionId'")
	}
	if args.Functions == nil {
		return nil, errors.New("invalid value for required argument 'Functions'")
	}
	var resource FunctionDefinitionVersion
	err := ctx.RegisterResource("aws-native:greengrass:FunctionDefinitionVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionDefinitionVersion gets an existing FunctionDefinitionVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionDefinitionVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionDefinitionVersionState, opts ...pulumi.ResourceOption) (*FunctionDefinitionVersion, error) {
	var resource FunctionDefinitionVersion
	err := ctx.ReadResource("aws-native:greengrass:FunctionDefinitionVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionDefinitionVersion resources.
type functionDefinitionVersionState struct {
}

type FunctionDefinitionVersionState struct {
}

func (FunctionDefinitionVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionDefinitionVersionState)(nil)).Elem()
}

type functionDefinitionVersionArgs struct {
	DefaultConfig        *FunctionDefinitionVersionDefaultConfig `pulumi:"defaultConfig"`
	FunctionDefinitionId string                                  `pulumi:"functionDefinitionId"`
	Functions            []FunctionDefinitionVersionFunction     `pulumi:"functions"`
}

// The set of arguments for constructing a FunctionDefinitionVersion resource.
type FunctionDefinitionVersionArgs struct {
	DefaultConfig        FunctionDefinitionVersionDefaultConfigPtrInput
	FunctionDefinitionId pulumi.StringInput
	Functions            FunctionDefinitionVersionFunctionArrayInput
}

func (FunctionDefinitionVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionDefinitionVersionArgs)(nil)).Elem()
}

type FunctionDefinitionVersionInput interface {
	pulumi.Input

	ToFunctionDefinitionVersionOutput() FunctionDefinitionVersionOutput
	ToFunctionDefinitionVersionOutputWithContext(ctx context.Context) FunctionDefinitionVersionOutput
}

func (*FunctionDefinitionVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionDefinitionVersion)(nil)).Elem()
}

func (i *FunctionDefinitionVersion) ToFunctionDefinitionVersionOutput() FunctionDefinitionVersionOutput {
	return i.ToFunctionDefinitionVersionOutputWithContext(context.Background())
}

func (i *FunctionDefinitionVersion) ToFunctionDefinitionVersionOutputWithContext(ctx context.Context) FunctionDefinitionVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionDefinitionVersionOutput)
}

type FunctionDefinitionVersionOutput struct{ *pulumi.OutputState }

func (FunctionDefinitionVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionDefinitionVersion)(nil)).Elem()
}

func (o FunctionDefinitionVersionOutput) ToFunctionDefinitionVersionOutput() FunctionDefinitionVersionOutput {
	return o
}

func (o FunctionDefinitionVersionOutput) ToFunctionDefinitionVersionOutputWithContext(ctx context.Context) FunctionDefinitionVersionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionDefinitionVersionInput)(nil)).Elem(), &FunctionDefinitionVersion{})
	pulumi.RegisterOutputType(FunctionDefinitionVersionOutput{})
}
