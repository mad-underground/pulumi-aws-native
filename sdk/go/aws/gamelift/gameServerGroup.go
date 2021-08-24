// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html
type GameServerGroup struct {
	pulumi.CustomResourceState

	AutoScalingGroupArn pulumi.StringOutput `pulumi:"autoScalingGroupArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-autoscalingpolicy
	AutoScalingPolicy GameServerGroupAutoScalingPolicyPtrOutput `pulumi:"autoScalingPolicy"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-balancingstrategy
	BalancingStrategy pulumi.StringPtrOutput `pulumi:"balancingStrategy"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-deleteoption
	DeleteOption       pulumi.StringPtrOutput `pulumi:"deleteOption"`
	GameServerGroupArn pulumi.StringOutput    `pulumi:"gameServerGroupArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-gameservergroupname
	GameServerGroupName pulumi.StringOutput `pulumi:"gameServerGroupName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-gameserverprotectionpolicy
	GameServerProtectionPolicy pulumi.StringPtrOutput `pulumi:"gameServerProtectionPolicy"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-instancedefinitions
	InstanceDefinitions GameServerGroupInstanceDefinitionArrayOutput `pulumi:"instanceDefinitions"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-launchtemplate
	LaunchTemplate GameServerGroupLaunchTemplateOutput `pulumi:"launchTemplate"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-maxsize
	MaxSize pulumi.Float64PtrOutput `pulumi:"maxSize"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-minsize
	MinSize pulumi.Float64PtrOutput `pulumi:"minSize"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-rolearn
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-vpcsubnets
	VpcSubnets pulumi.StringArrayOutput `pulumi:"vpcSubnets"`
}

// NewGameServerGroup registers a new resource with the given unique name, arguments, and options.
func NewGameServerGroup(ctx *pulumi.Context,
	name string, args *GameServerGroupArgs, opts ...pulumi.ResourceOption) (*GameServerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GameServerGroupName == nil {
		return nil, errors.New("invalid value for required argument 'GameServerGroupName'")
	}
	if args.InstanceDefinitions == nil {
		return nil, errors.New("invalid value for required argument 'InstanceDefinitions'")
	}
	if args.LaunchTemplate == nil {
		return nil, errors.New("invalid value for required argument 'LaunchTemplate'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource GameServerGroup
	err := ctx.RegisterResource("aws-native:gamelift:GameServerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGameServerGroup gets an existing GameServerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGameServerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GameServerGroupState, opts ...pulumi.ResourceOption) (*GameServerGroup, error) {
	var resource GameServerGroup
	err := ctx.ReadResource("aws-native:gamelift:GameServerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GameServerGroup resources.
type gameServerGroupState struct {
}

type GameServerGroupState struct {
}

func (GameServerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*gameServerGroupState)(nil)).Elem()
}

type gameServerGroupArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-autoscalingpolicy
	AutoScalingPolicy *GameServerGroupAutoScalingPolicy `pulumi:"autoScalingPolicy"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-balancingstrategy
	BalancingStrategy *string `pulumi:"balancingStrategy"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-deleteoption
	DeleteOption *string `pulumi:"deleteOption"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-gameservergroupname
	GameServerGroupName string `pulumi:"gameServerGroupName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-gameserverprotectionpolicy
	GameServerProtectionPolicy *string `pulumi:"gameServerProtectionPolicy"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-instancedefinitions
	InstanceDefinitions []GameServerGroupInstanceDefinition `pulumi:"instanceDefinitions"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-launchtemplate
	LaunchTemplate GameServerGroupLaunchTemplate `pulumi:"launchTemplate"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-maxsize
	MaxSize *float64 `pulumi:"maxSize"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-minsize
	MinSize *float64 `pulumi:"minSize"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-rolearn
	RoleArn string `pulumi:"roleArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-tags
	Tags []aws.Tag `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-vpcsubnets
	VpcSubnets []string `pulumi:"vpcSubnets"`
}

// The set of arguments for constructing a GameServerGroup resource.
type GameServerGroupArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-autoscalingpolicy
	AutoScalingPolicy GameServerGroupAutoScalingPolicyPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-balancingstrategy
	BalancingStrategy pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-deleteoption
	DeleteOption pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-gameservergroupname
	GameServerGroupName pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-gameserverprotectionpolicy
	GameServerProtectionPolicy pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-instancedefinitions
	InstanceDefinitions GameServerGroupInstanceDefinitionArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-launchtemplate
	LaunchTemplate GameServerGroupLaunchTemplateInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-maxsize
	MaxSize pulumi.Float64PtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-minsize
	MinSize pulumi.Float64PtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-rolearn
	RoleArn pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-tags
	Tags aws.TagArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-vpcsubnets
	VpcSubnets pulumi.StringArrayInput
}

func (GameServerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gameServerGroupArgs)(nil)).Elem()
}

type GameServerGroupInput interface {
	pulumi.Input

	ToGameServerGroupOutput() GameServerGroupOutput
	ToGameServerGroupOutputWithContext(ctx context.Context) GameServerGroupOutput
}

func (*GameServerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*GameServerGroup)(nil))
}

func (i *GameServerGroup) ToGameServerGroupOutput() GameServerGroupOutput {
	return i.ToGameServerGroupOutputWithContext(context.Background())
}

func (i *GameServerGroup) ToGameServerGroupOutputWithContext(ctx context.Context) GameServerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameServerGroupOutput)
}

type GameServerGroupOutput struct{ *pulumi.OutputState }

func (GameServerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GameServerGroup)(nil))
}

func (o GameServerGroupOutput) ToGameServerGroupOutput() GameServerGroupOutput {
	return o
}

func (o GameServerGroupOutput) ToGameServerGroupOutputWithContext(ctx context.Context) GameServerGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GameServerGroupOutput{})
}