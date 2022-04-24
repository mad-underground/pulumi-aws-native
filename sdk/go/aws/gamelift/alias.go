// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::GameLift::Alias resource creates an alias for an Amazon GameLift (GameLift) fleet destination.
type Alias struct {
	pulumi.CustomResourceState

	// Unique alias ID
	AliasId pulumi.StringOutput `pulumi:"aliasId"`
	// A human-readable description of the alias.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A descriptive label that is associated with an alias. Alias names do not need to be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// A routing configuration that specifies where traffic is directed for this alias, such as to a fleet or to a message.
	RoutingStrategy AliasRoutingStrategyOutput `pulumi:"routingStrategy"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoutingStrategy == nil {
		return nil, errors.New("invalid value for required argument 'RoutingStrategy'")
	}
	var resource Alias
	err := ctx.RegisterResource("aws-native:gamelift:Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	var resource Alias
	err := ctx.ReadResource("aws-native:gamelift:Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alias resources.
type aliasState struct {
}

type AliasState struct {
}

func (AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasState)(nil)).Elem()
}

type aliasArgs struct {
	// A human-readable description of the alias.
	Description *string `pulumi:"description"`
	// A descriptive label that is associated with an alias. Alias names do not need to be unique.
	Name *string `pulumi:"name"`
	// A routing configuration that specifies where traffic is directed for this alias, such as to a fleet or to a message.
	RoutingStrategy AliasRoutingStrategy `pulumi:"routingStrategy"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// A human-readable description of the alias.
	Description pulumi.StringPtrInput
	// A descriptive label that is associated with an alias. Alias names do not need to be unique.
	Name pulumi.StringPtrInput
	// A routing configuration that specifies where traffic is directed for this alias, such as to a fleet or to a message.
	RoutingStrategy AliasRoutingStrategyInput
}

func (AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasArgs)(nil)).Elem()
}

type AliasInput interface {
	pulumi.Input

	ToAliasOutput() AliasOutput
	ToAliasOutputWithContext(ctx context.Context) AliasOutput
}

func (*Alias) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (i *Alias) ToAliasOutput() AliasOutput {
	return i.ToAliasOutputWithContext(context.Background())
}

func (i *Alias) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasOutput)
}

type AliasOutput struct{ *pulumi.OutputState }

func (AliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (o AliasOutput) ToAliasOutput() AliasOutput {
	return o
}

func (o AliasOutput) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AliasInput)(nil)).Elem(), &Alias{})
	pulumi.RegisterOutputType(AliasOutput{})
}
