// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.EC2.Outputs
{

    [OutputType]
    public sealed class SpotFleetSpotFleetRequestConfigData
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-allocationstrategy
        /// </summary>
        public readonly string? AllocationStrategy;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-excesscapacityterminationpolicy
        /// </summary>
        public readonly string? ExcessCapacityTerminationPolicy;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-iamfleetrole
        /// </summary>
        public readonly string IamFleetRole;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-instanceinterruptionbehavior
        /// </summary>
        public readonly string? InstanceInterruptionBehavior;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-instancepoolstousecount
        /// </summary>
        public readonly int? InstancePoolsToUseCount;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications
        /// </summary>
        public readonly ImmutableArray<Outputs.SpotFleetSpotFleetLaunchSpecification> LaunchSpecifications;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-launchtemplateconfigs
        /// </summary>
        public readonly ImmutableArray<Outputs.SpotFleetLaunchTemplateConfig> LaunchTemplateConfigs;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-loadbalancersconfig
        /// </summary>
        public readonly Outputs.SpotFleetLoadBalancersConfig? LoadBalancersConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-ondemandallocationstrategy
        /// </summary>
        public readonly string? OnDemandAllocationStrategy;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-ondemandmaxtotalprice
        /// </summary>
        public readonly string? OnDemandMaxTotalPrice;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-ondemandtargetcapacity
        /// </summary>
        public readonly int? OnDemandTargetCapacity;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-replaceunhealthyinstances
        /// </summary>
        public readonly bool? ReplaceUnhealthyInstances;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-spotmaintenancestrategies
        /// </summary>
        public readonly Outputs.SpotFleetSpotMaintenanceStrategies? SpotMaintenanceStrategies;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-spotmaxtotalprice
        /// </summary>
        public readonly string? SpotMaxTotalPrice;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-spotprice
        /// </summary>
        public readonly string? SpotPrice;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-targetcapacity
        /// </summary>
        public readonly int TargetCapacity;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-terminateinstanceswithexpiration
        /// </summary>
        public readonly bool? TerminateInstancesWithExpiration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-type
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-validfrom
        /// </summary>
        public readonly string? ValidFrom;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-validuntil
        /// </summary>
        public readonly string? ValidUntil;

        [OutputConstructor]
        private SpotFleetSpotFleetRequestConfigData(
            string? AllocationStrategy,

            string? ExcessCapacityTerminationPolicy,

            string IamFleetRole,

            string? InstanceInterruptionBehavior,

            int? InstancePoolsToUseCount,

            ImmutableArray<Outputs.SpotFleetSpotFleetLaunchSpecification> LaunchSpecifications,

            ImmutableArray<Outputs.SpotFleetLaunchTemplateConfig> LaunchTemplateConfigs,

            Outputs.SpotFleetLoadBalancersConfig? LoadBalancersConfig,

            string? OnDemandAllocationStrategy,

            string? OnDemandMaxTotalPrice,

            int? OnDemandTargetCapacity,

            bool? ReplaceUnhealthyInstances,

            Outputs.SpotFleetSpotMaintenanceStrategies? SpotMaintenanceStrategies,

            string? SpotMaxTotalPrice,

            string? SpotPrice,

            int TargetCapacity,

            bool? TerminateInstancesWithExpiration,

            string? Type,

            string? ValidFrom,

            string? ValidUntil)
        {
            this.AllocationStrategy = AllocationStrategy;
            this.ExcessCapacityTerminationPolicy = ExcessCapacityTerminationPolicy;
            this.IamFleetRole = IamFleetRole;
            this.InstanceInterruptionBehavior = InstanceInterruptionBehavior;
            this.InstancePoolsToUseCount = InstancePoolsToUseCount;
            this.LaunchSpecifications = LaunchSpecifications;
            this.LaunchTemplateConfigs = LaunchTemplateConfigs;
            this.LoadBalancersConfig = LoadBalancersConfig;
            this.OnDemandAllocationStrategy = OnDemandAllocationStrategy;
            this.OnDemandMaxTotalPrice = OnDemandMaxTotalPrice;
            this.OnDemandTargetCapacity = OnDemandTargetCapacity;
            this.ReplaceUnhealthyInstances = ReplaceUnhealthyInstances;
            this.SpotMaintenanceStrategies = SpotMaintenanceStrategies;
            this.SpotMaxTotalPrice = SpotMaxTotalPrice;
            this.SpotPrice = SpotPrice;
            this.TargetCapacity = TargetCapacity;
            this.TerminateInstancesWithExpiration = TerminateInstancesWithExpiration;
            this.Type = Type;
            this.ValidFrom = ValidFrom;
            this.ValidUntil = ValidUntil;
        }
    }
}
