// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Redshift::ClusterParameterGroup
type ClusterParameterGroup struct {
	pulumi.CustomResourceState

	// A description of the parameter group.
	Description pulumi.StringOutput `pulumi:"description"`
	// The Amazon Redshift engine version to which the cluster parameter group applies. The cluster engine version determines the set of parameters.
	ParameterGroupFamily pulumi.StringOutput `pulumi:"parameterGroupFamily"`
	// The name of the cluster parameter group.
	ParameterGroupName pulumi.StringOutput `pulumi:"parameterGroupName"`
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	Parameters ClusterParameterGroupParameterArrayOutput `pulumi:"parameters"`
	// An array of key-value pairs to apply to this resource.
	Tags ClusterParameterGroupTagArrayOutput `pulumi:"tags"`
}

// NewClusterParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewClusterParameterGroup(ctx *pulumi.Context,
	name string, args *ClusterParameterGroupArgs, opts ...pulumi.ResourceOption) (*ClusterParameterGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.ParameterGroupFamily == nil {
		return nil, errors.New("invalid value for required argument 'ParameterGroupFamily'")
	}
	var resource ClusterParameterGroup
	err := ctx.RegisterResource("aws-native:redshift:ClusterParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterParameterGroup gets an existing ClusterParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterParameterGroupState, opts ...pulumi.ResourceOption) (*ClusterParameterGroup, error) {
	var resource ClusterParameterGroup
	err := ctx.ReadResource("aws-native:redshift:ClusterParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterParameterGroup resources.
type clusterParameterGroupState struct {
}

type ClusterParameterGroupState struct {
}

func (ClusterParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterParameterGroupState)(nil)).Elem()
}

type clusterParameterGroupArgs struct {
	// A description of the parameter group.
	Description string `pulumi:"description"`
	// The Amazon Redshift engine version to which the cluster parameter group applies. The cluster engine version determines the set of parameters.
	ParameterGroupFamily string `pulumi:"parameterGroupFamily"`
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	Parameters []ClusterParameterGroupParameter `pulumi:"parameters"`
	// An array of key-value pairs to apply to this resource.
	Tags []ClusterParameterGroupTag `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterParameterGroup resource.
type ClusterParameterGroupArgs struct {
	// A description of the parameter group.
	Description pulumi.StringInput
	// The Amazon Redshift engine version to which the cluster parameter group applies. The cluster engine version determines the set of parameters.
	ParameterGroupFamily pulumi.StringInput
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	Parameters ClusterParameterGroupParameterArrayInput
	// An array of key-value pairs to apply to this resource.
	Tags ClusterParameterGroupTagArrayInput
}

func (ClusterParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterParameterGroupArgs)(nil)).Elem()
}

type ClusterParameterGroupInput interface {
	pulumi.Input

	ToClusterParameterGroupOutput() ClusterParameterGroupOutput
	ToClusterParameterGroupOutputWithContext(ctx context.Context) ClusterParameterGroupOutput
}

func (*ClusterParameterGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterParameterGroup)(nil)).Elem()
}

func (i *ClusterParameterGroup) ToClusterParameterGroupOutput() ClusterParameterGroupOutput {
	return i.ToClusterParameterGroupOutputWithContext(context.Background())
}

func (i *ClusterParameterGroup) ToClusterParameterGroupOutputWithContext(ctx context.Context) ClusterParameterGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupOutput)
}

type ClusterParameterGroupOutput struct{ *pulumi.OutputState }

func (ClusterParameterGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterParameterGroup)(nil)).Elem()
}

func (o ClusterParameterGroupOutput) ToClusterParameterGroupOutput() ClusterParameterGroupOutput {
	return o
}

func (o ClusterParameterGroupOutput) ToClusterParameterGroupOutputWithContext(ctx context.Context) ClusterParameterGroupOutput {
	return o
}

// A description of the parameter group.
func (o ClusterParameterGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterParameterGroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The Amazon Redshift engine version to which the cluster parameter group applies. The cluster engine version determines the set of parameters.
func (o ClusterParameterGroupOutput) ParameterGroupFamily() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterParameterGroup) pulumi.StringOutput { return v.ParameterGroupFamily }).(pulumi.StringOutput)
}

// The name of the cluster parameter group.
func (o ClusterParameterGroupOutput) ParameterGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterParameterGroup) pulumi.StringOutput { return v.ParameterGroupName }).(pulumi.StringOutput)
}

// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
func (o ClusterParameterGroupOutput) Parameters() ClusterParameterGroupParameterArrayOutput {
	return o.ApplyT(func(v *ClusterParameterGroup) ClusterParameterGroupParameterArrayOutput { return v.Parameters }).(ClusterParameterGroupParameterArrayOutput)
}

// An array of key-value pairs to apply to this resource.
func (o ClusterParameterGroupOutput) Tags() ClusterParameterGroupTagArrayOutput {
	return o.ApplyT(func(v *ClusterParameterGroup) ClusterParameterGroupTagArrayOutput { return v.Tags }).(ClusterParameterGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterParameterGroupInput)(nil)).Elem(), &ClusterParameterGroup{})
	pulumi.RegisterOutputType(ClusterParameterGroupOutput{})
}
