// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::RDS::DBClusterParameterGroup resource creates a new Amazon RDS DB cluster parameter group. For more information, see Managing an Amazon Aurora DB Cluster in the Amazon Aurora User Guide.
type DBClusterParameterGroup struct {
	pulumi.CustomResourceState

	DBClusterParameterGroupName pulumi.StringPtrOutput `pulumi:"dBClusterParameterGroupName"`
	// A friendly description for this DB cluster parameter group.
	Description pulumi.StringOutput `pulumi:"description"`
	// The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a DB engine and engine version compatible with that DB cluster parameter group family.
	Family pulumi.StringOutput `pulumi:"family"`
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	Parameters pulumi.AnyOutput `pulumi:"parameters"`
	// The list of tags for the cluster parameter group.
	Tags DBClusterParameterGroupTagArrayOutput `pulumi:"tags"`
}

// NewDBClusterParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewDBClusterParameterGroup(ctx *pulumi.Context,
	name string, args *DBClusterParameterGroupArgs, opts ...pulumi.ResourceOption) (*DBClusterParameterGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Family == nil {
		return nil, errors.New("invalid value for required argument 'Family'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DBClusterParameterGroup
	err := ctx.RegisterResource("aws-native:rds:DBClusterParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDBClusterParameterGroup gets an existing DBClusterParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDBClusterParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DBClusterParameterGroupState, opts ...pulumi.ResourceOption) (*DBClusterParameterGroup, error) {
	var resource DBClusterParameterGroup
	err := ctx.ReadResource("aws-native:rds:DBClusterParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DBClusterParameterGroup resources.
type dbclusterParameterGroupState struct {
}

type DBClusterParameterGroupState struct {
}

func (DBClusterParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbclusterParameterGroupState)(nil)).Elem()
}

type dbclusterParameterGroupArgs struct {
	DBClusterParameterGroupName *string `pulumi:"dBClusterParameterGroupName"`
	// A friendly description for this DB cluster parameter group.
	Description string `pulumi:"description"`
	// The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a DB engine and engine version compatible with that DB cluster parameter group family.
	Family string `pulumi:"family"`
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	Parameters interface{} `pulumi:"parameters"`
	// The list of tags for the cluster parameter group.
	Tags []DBClusterParameterGroupTag `pulumi:"tags"`
}

// The set of arguments for constructing a DBClusterParameterGroup resource.
type DBClusterParameterGroupArgs struct {
	DBClusterParameterGroupName pulumi.StringPtrInput
	// A friendly description for this DB cluster parameter group.
	Description pulumi.StringInput
	// The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a DB engine and engine version compatible with that DB cluster parameter group family.
	Family pulumi.StringInput
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	Parameters pulumi.Input
	// The list of tags for the cluster parameter group.
	Tags DBClusterParameterGroupTagArrayInput
}

func (DBClusterParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbclusterParameterGroupArgs)(nil)).Elem()
}

type DBClusterParameterGroupInput interface {
	pulumi.Input

	ToDBClusterParameterGroupOutput() DBClusterParameterGroupOutput
	ToDBClusterParameterGroupOutputWithContext(ctx context.Context) DBClusterParameterGroupOutput
}

func (*DBClusterParameterGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DBClusterParameterGroup)(nil)).Elem()
}

func (i *DBClusterParameterGroup) ToDBClusterParameterGroupOutput() DBClusterParameterGroupOutput {
	return i.ToDBClusterParameterGroupOutputWithContext(context.Background())
}

func (i *DBClusterParameterGroup) ToDBClusterParameterGroupOutputWithContext(ctx context.Context) DBClusterParameterGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterParameterGroupOutput)
}

type DBClusterParameterGroupOutput struct{ *pulumi.OutputState }

func (DBClusterParameterGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DBClusterParameterGroup)(nil)).Elem()
}

func (o DBClusterParameterGroupOutput) ToDBClusterParameterGroupOutput() DBClusterParameterGroupOutput {
	return o
}

func (o DBClusterParameterGroupOutput) ToDBClusterParameterGroupOutputWithContext(ctx context.Context) DBClusterParameterGroupOutput {
	return o
}

func (o DBClusterParameterGroupOutput) DBClusterParameterGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DBClusterParameterGroup) pulumi.StringPtrOutput { return v.DBClusterParameterGroupName }).(pulumi.StringPtrOutput)
}

// A friendly description for this DB cluster parameter group.
func (o DBClusterParameterGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterParameterGroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a DB engine and engine version compatible with that DB cluster parameter group family.
func (o DBClusterParameterGroupOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v *DBClusterParameterGroup) pulumi.StringOutput { return v.Family }).(pulumi.StringOutput)
}

// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
func (o DBClusterParameterGroupOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *DBClusterParameterGroup) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

// The list of tags for the cluster parameter group.
func (o DBClusterParameterGroupOutput) Tags() DBClusterParameterGroupTagArrayOutput {
	return o.ApplyT(func(v *DBClusterParameterGroup) DBClusterParameterGroupTagArrayOutput { return v.Tags }).(DBClusterParameterGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterParameterGroupInput)(nil)).Elem(), &DBClusterParameterGroup{})
	pulumi.RegisterOutputType(DBClusterParameterGroupOutput{})
}
