// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::RDS::OptionGroup
//
// Deprecated: OptionGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type OptionGroup struct {
	pulumi.CustomResourceState

	EngineName             pulumi.StringOutput                       `pulumi:"engineName"`
	MajorEngineVersion     pulumi.StringOutput                       `pulumi:"majorEngineVersion"`
	OptionConfigurations   OptionGroupOptionConfigurationArrayOutput `pulumi:"optionConfigurations"`
	OptionGroupDescription pulumi.StringOutput                       `pulumi:"optionGroupDescription"`
	Tags                   OptionGroupTagArrayOutput                 `pulumi:"tags"`
}

// NewOptionGroup registers a new resource with the given unique name, arguments, and options.
func NewOptionGroup(ctx *pulumi.Context,
	name string, args *OptionGroupArgs, opts ...pulumi.ResourceOption) (*OptionGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EngineName == nil {
		return nil, errors.New("invalid value for required argument 'EngineName'")
	}
	if args.MajorEngineVersion == nil {
		return nil, errors.New("invalid value for required argument 'MajorEngineVersion'")
	}
	if args.OptionConfigurations == nil {
		return nil, errors.New("invalid value for required argument 'OptionConfigurations'")
	}
	if args.OptionGroupDescription == nil {
		return nil, errors.New("invalid value for required argument 'OptionGroupDescription'")
	}
	var resource OptionGroup
	err := ctx.RegisterResource("aws-native:rds:OptionGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOptionGroup gets an existing OptionGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOptionGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OptionGroupState, opts ...pulumi.ResourceOption) (*OptionGroup, error) {
	var resource OptionGroup
	err := ctx.ReadResource("aws-native:rds:OptionGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OptionGroup resources.
type optionGroupState struct {
}

type OptionGroupState struct {
}

func (OptionGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*optionGroupState)(nil)).Elem()
}

type optionGroupArgs struct {
	EngineName             string                           `pulumi:"engineName"`
	MajorEngineVersion     string                           `pulumi:"majorEngineVersion"`
	OptionConfigurations   []OptionGroupOptionConfiguration `pulumi:"optionConfigurations"`
	OptionGroupDescription string                           `pulumi:"optionGroupDescription"`
	Tags                   []OptionGroupTag                 `pulumi:"tags"`
}

// The set of arguments for constructing a OptionGroup resource.
type OptionGroupArgs struct {
	EngineName             pulumi.StringInput
	MajorEngineVersion     pulumi.StringInput
	OptionConfigurations   OptionGroupOptionConfigurationArrayInput
	OptionGroupDescription pulumi.StringInput
	Tags                   OptionGroupTagArrayInput
}

func (OptionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*optionGroupArgs)(nil)).Elem()
}

type OptionGroupInput interface {
	pulumi.Input

	ToOptionGroupOutput() OptionGroupOutput
	ToOptionGroupOutputWithContext(ctx context.Context) OptionGroupOutput
}

func (*OptionGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*OptionGroup)(nil))
}

func (i *OptionGroup) ToOptionGroupOutput() OptionGroupOutput {
	return i.ToOptionGroupOutputWithContext(context.Background())
}

func (i *OptionGroup) ToOptionGroupOutputWithContext(ctx context.Context) OptionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OptionGroupOutput)
}

type OptionGroupOutput struct{ *pulumi.OutputState }

func (OptionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OptionGroup)(nil))
}

func (o OptionGroupOutput) ToOptionGroupOutput() OptionGroupOutput {
	return o
}

func (o OptionGroupOutput) ToOptionGroupOutputWithContext(ctx context.Context) OptionGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OptionGroupOutput{})
}