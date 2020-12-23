// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.GameLift.Outputs
{

    [OutputType]
    public sealed class GameServerGroupProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-autoscalingpolicy
        /// </summary>
        public readonly Outputs.GameServerGroupAutoScalingPolicy? AutoScalingPolicy;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-balancingstrategy
        /// </summary>
        public readonly string? BalancingStrategy;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-deleteoption
        /// </summary>
        public readonly string? DeleteOption;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-gameservergroupname
        /// </summary>
        public readonly string GameServerGroupName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-gameserverprotectionpolicy
        /// </summary>
        public readonly string? GameServerProtectionPolicy;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-instancedefinitions
        /// </summary>
        public readonly Outputs.GameServerGroupInstanceDefinitions InstanceDefinitions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-launchtemplate
        /// </summary>
        public readonly Outputs.GameServerGroupLaunchTemplate LaunchTemplate;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-maxsize
        /// </summary>
        public readonly double? MaxSize;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-minsize
        /// </summary>
        public readonly double? MinSize;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-rolearn
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-tags
        /// </summary>
        public readonly Outputs.GameServerGroupTags? Tags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gameservergroup.html#cfn-gamelift-gameservergroup-vpcsubnets
        /// </summary>
        public readonly Outputs.GameServerGroupVpcSubnets? VpcSubnets;

        [OutputConstructor]
        private GameServerGroupProperties(
            Outputs.GameServerGroupAutoScalingPolicy? AutoScalingPolicy,

            string? BalancingStrategy,

            string? DeleteOption,

            string GameServerGroupName,

            string? GameServerProtectionPolicy,

            Outputs.GameServerGroupInstanceDefinitions InstanceDefinitions,

            Outputs.GameServerGroupLaunchTemplate LaunchTemplate,

            double? MaxSize,

            double? MinSize,

            string RoleArn,

            Outputs.GameServerGroupTags? Tags,

            Outputs.GameServerGroupVpcSubnets? VpcSubnets)
        {
            this.AutoScalingPolicy = AutoScalingPolicy;
            this.BalancingStrategy = BalancingStrategy;
            this.DeleteOption = DeleteOption;
            this.GameServerGroupName = GameServerGroupName;
            this.GameServerProtectionPolicy = GameServerProtectionPolicy;
            this.InstanceDefinitions = InstanceDefinitions;
            this.LaunchTemplate = LaunchTemplate;
            this.MaxSize = MaxSize;
            this.MinSize = MinSize;
            this.RoleArn = RoleArn;
            this.Tags = Tags;
            this.VpcSubnets = VpcSubnets;
        }
    }
}
