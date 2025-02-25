// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DocDB::DBSubnetGroup
//
// Deprecated: DbSubnetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type DbSubnetGroup struct {
	pulumi.CustomResourceState

	AwsId                    pulumi.StringOutput      `pulumi:"awsId"`
	DbSubnetGroupDescription pulumi.StringOutput      `pulumi:"dbSubnetGroupDescription"`
	DbSubnetGroupName        pulumi.StringPtrOutput   `pulumi:"dbSubnetGroupName"`
	SubnetIds                pulumi.StringArrayOutput `pulumi:"subnetIds"`
	Tags                     aws.TagArrayOutput       `pulumi:"tags"`
}

// NewDbSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewDbSubnetGroup(ctx *pulumi.Context,
	name string, args *DbSubnetGroupArgs, opts ...pulumi.ResourceOption) (*DbSubnetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbSubnetGroupDescription == nil {
		return nil, errors.New("invalid value for required argument 'DbSubnetGroupDescription'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"dbSubnetGroupName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DbSubnetGroup
	err := ctx.RegisterResource("aws-native:docdb:DbSubnetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDbSubnetGroup gets an existing DbSubnetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDbSubnetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DbSubnetGroupState, opts ...pulumi.ResourceOption) (*DbSubnetGroup, error) {
	var resource DbSubnetGroup
	err := ctx.ReadResource("aws-native:docdb:DbSubnetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DbSubnetGroup resources.
type dbSubnetGroupState struct {
}

type DbSubnetGroupState struct {
}

func (DbSubnetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbSubnetGroupState)(nil)).Elem()
}

type dbSubnetGroupArgs struct {
	DbSubnetGroupDescription string    `pulumi:"dbSubnetGroupDescription"`
	DbSubnetGroupName        *string   `pulumi:"dbSubnetGroupName"`
	SubnetIds                []string  `pulumi:"subnetIds"`
	Tags                     []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a DbSubnetGroup resource.
type DbSubnetGroupArgs struct {
	DbSubnetGroupDescription pulumi.StringInput
	DbSubnetGroupName        pulumi.StringPtrInput
	SubnetIds                pulumi.StringArrayInput
	Tags                     aws.TagArrayInput
}

func (DbSubnetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbSubnetGroupArgs)(nil)).Elem()
}

type DbSubnetGroupInput interface {
	pulumi.Input

	ToDbSubnetGroupOutput() DbSubnetGroupOutput
	ToDbSubnetGroupOutputWithContext(ctx context.Context) DbSubnetGroupOutput
}

func (*DbSubnetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DbSubnetGroup)(nil)).Elem()
}

func (i *DbSubnetGroup) ToDbSubnetGroupOutput() DbSubnetGroupOutput {
	return i.ToDbSubnetGroupOutputWithContext(context.Background())
}

func (i *DbSubnetGroup) ToDbSubnetGroupOutputWithContext(ctx context.Context) DbSubnetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbSubnetGroupOutput)
}

type DbSubnetGroupOutput struct{ *pulumi.OutputState }

func (DbSubnetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DbSubnetGroup)(nil)).Elem()
}

func (o DbSubnetGroupOutput) ToDbSubnetGroupOutput() DbSubnetGroupOutput {
	return o
}

func (o DbSubnetGroupOutput) ToDbSubnetGroupOutputWithContext(ctx context.Context) DbSubnetGroupOutput {
	return o
}

func (o DbSubnetGroupOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *DbSubnetGroup) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o DbSubnetGroupOutput) DbSubnetGroupDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *DbSubnetGroup) pulumi.StringOutput { return v.DbSubnetGroupDescription }).(pulumi.StringOutput)
}

func (o DbSubnetGroupOutput) DbSubnetGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DbSubnetGroup) pulumi.StringPtrOutput { return v.DbSubnetGroupName }).(pulumi.StringPtrOutput)
}

func (o DbSubnetGroupOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DbSubnetGroup) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func (o DbSubnetGroupOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *DbSubnetGroup) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DbSubnetGroupInput)(nil)).Elem(), &DbSubnetGroup{})
	pulumi.RegisterOutputType(DbSubnetGroupOutput{})
}
