// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html
type DBProxyTargetGroup struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-connectionpoolconfigurationinfo
	ConnectionPoolConfigurationInfo DBProxyTargetGroupConnectionPoolConfigurationInfoFormatPtrOutput `pulumi:"connectionPoolConfigurationInfo"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbclusteridentifiers
	DBClusterIdentifiers pulumi.StringArrayOutput `pulumi:"dBClusterIdentifiers"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbinstanceidentifiers
	DBInstanceIdentifiers pulumi.StringArrayOutput `pulumi:"dBInstanceIdentifiers"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbproxyname
	DBProxyName    pulumi.StringOutput `pulumi:"dBProxyName"`
	TargetGroupArn pulumi.StringOutput `pulumi:"targetGroupArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-targetgroupname
	TargetGroupName pulumi.StringOutput `pulumi:"targetGroupName"`
}

// NewDBProxyTargetGroup registers a new resource with the given unique name, arguments, and options.
func NewDBProxyTargetGroup(ctx *pulumi.Context,
	name string, args *DBProxyTargetGroupArgs, opts ...pulumi.ResourceOption) (*DBProxyTargetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DBProxyName == nil {
		return nil, errors.New("invalid value for required argument 'DBProxyName'")
	}
	if args.TargetGroupName == nil {
		return nil, errors.New("invalid value for required argument 'TargetGroupName'")
	}
	var resource DBProxyTargetGroup
	err := ctx.RegisterResource("aws-native:rds:DBProxyTargetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDBProxyTargetGroup gets an existing DBProxyTargetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDBProxyTargetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DBProxyTargetGroupState, opts ...pulumi.ResourceOption) (*DBProxyTargetGroup, error) {
	var resource DBProxyTargetGroup
	err := ctx.ReadResource("aws-native:rds:DBProxyTargetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DBProxyTargetGroup resources.
type dbproxyTargetGroupState struct {
}

type DBProxyTargetGroupState struct {
}

func (DBProxyTargetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbproxyTargetGroupState)(nil)).Elem()
}

type dbproxyTargetGroupArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-connectionpoolconfigurationinfo
	ConnectionPoolConfigurationInfo *DBProxyTargetGroupConnectionPoolConfigurationInfoFormat `pulumi:"connectionPoolConfigurationInfo"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbclusteridentifiers
	DBClusterIdentifiers []string `pulumi:"dBClusterIdentifiers"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbinstanceidentifiers
	DBInstanceIdentifiers []string `pulumi:"dBInstanceIdentifiers"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbproxyname
	DBProxyName string `pulumi:"dBProxyName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-targetgroupname
	TargetGroupName string `pulumi:"targetGroupName"`
}

// The set of arguments for constructing a DBProxyTargetGroup resource.
type DBProxyTargetGroupArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-connectionpoolconfigurationinfo
	ConnectionPoolConfigurationInfo DBProxyTargetGroupConnectionPoolConfigurationInfoFormatPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbclusteridentifiers
	DBClusterIdentifiers pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbinstanceidentifiers
	DBInstanceIdentifiers pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbproxyname
	DBProxyName pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-targetgroupname
	TargetGroupName pulumi.StringInput
}

func (DBProxyTargetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbproxyTargetGroupArgs)(nil)).Elem()
}

type DBProxyTargetGroupInput interface {
	pulumi.Input

	ToDBProxyTargetGroupOutput() DBProxyTargetGroupOutput
	ToDBProxyTargetGroupOutputWithContext(ctx context.Context) DBProxyTargetGroupOutput
}

func (*DBProxyTargetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*DBProxyTargetGroup)(nil))
}

func (i *DBProxyTargetGroup) ToDBProxyTargetGroupOutput() DBProxyTargetGroupOutput {
	return i.ToDBProxyTargetGroupOutputWithContext(context.Background())
}

func (i *DBProxyTargetGroup) ToDBProxyTargetGroupOutputWithContext(ctx context.Context) DBProxyTargetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBProxyTargetGroupOutput)
}

type DBProxyTargetGroupOutput struct{ *pulumi.OutputState }

func (DBProxyTargetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBProxyTargetGroup)(nil))
}

func (o DBProxyTargetGroupOutput) ToDBProxyTargetGroupOutput() DBProxyTargetGroupOutput {
	return o
}

func (o DBProxyTargetGroupOutput) ToDBProxyTargetGroupOutputWithContext(ctx context.Context) DBProxyTargetGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DBProxyTargetGroupOutput{})
}