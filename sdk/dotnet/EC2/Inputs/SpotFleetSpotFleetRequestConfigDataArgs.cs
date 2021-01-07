// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.EC2.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html
    /// </summary>
    public sealed class SpotFleetSpotFleetRequestConfigDataArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-allocationstrategy
        /// </summary>
        [Input("AllocationStrategy")]
        public Input<string>? AllocationStrategy { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-excesscapacityterminationpolicy
        /// </summary>
        [Input("ExcessCapacityTerminationPolicy")]
        public Input<string>? ExcessCapacityTerminationPolicy { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-iamfleetrole
        /// </summary>
        [Input("IamFleetRole", required: true)]
        public Input<string> IamFleetRole { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-instanceinterruptionbehavior
        /// </summary>
        [Input("InstanceInterruptionBehavior")]
        public Input<string>? InstanceInterruptionBehavior { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-instancepoolstousecount
        /// </summary>
        [Input("InstancePoolsToUseCount")]
        public Input<int>? InstancePoolsToUseCount { get; set; }

        [Input("LaunchSpecifications")]
        private InputList<Inputs.SpotFleetSpotFleetLaunchSpecificationArgs>? _LaunchSpecifications;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications
        /// </summary>
        public InputList<Inputs.SpotFleetSpotFleetLaunchSpecificationArgs> LaunchSpecifications
        {
            get => _LaunchSpecifications ?? (_LaunchSpecifications = new InputList<Inputs.SpotFleetSpotFleetLaunchSpecificationArgs>());
            set => _LaunchSpecifications = value;
        }

        [Input("LaunchTemplateConfigs")]
        private InputList<Inputs.SpotFleetLaunchTemplateConfigArgs>? _LaunchTemplateConfigs;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-launchtemplateconfigs
        /// </summary>
        public InputList<Inputs.SpotFleetLaunchTemplateConfigArgs> LaunchTemplateConfigs
        {
            get => _LaunchTemplateConfigs ?? (_LaunchTemplateConfigs = new InputList<Inputs.SpotFleetLaunchTemplateConfigArgs>());
            set => _LaunchTemplateConfigs = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-loadbalancersconfig
        /// </summary>
        [Input("LoadBalancersConfig")]
        public Input<Inputs.SpotFleetLoadBalancersConfigArgs>? LoadBalancersConfig { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-ondemandallocationstrategy
        /// </summary>
        [Input("OnDemandAllocationStrategy")]
        public Input<string>? OnDemandAllocationStrategy { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-ondemandmaxtotalprice
        /// </summary>
        [Input("OnDemandMaxTotalPrice")]
        public Input<string>? OnDemandMaxTotalPrice { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-ondemandtargetcapacity
        /// </summary>
        [Input("OnDemandTargetCapacity")]
        public Input<int>? OnDemandTargetCapacity { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-replaceunhealthyinstances
        /// </summary>
        [Input("ReplaceUnhealthyInstances")]
        public Input<bool>? ReplaceUnhealthyInstances { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-spotmaintenancestrategies
        /// </summary>
        [Input("SpotMaintenanceStrategies")]
        public Input<Inputs.SpotFleetSpotMaintenanceStrategiesArgs>? SpotMaintenanceStrategies { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-spotmaxtotalprice
        /// </summary>
        [Input("SpotMaxTotalPrice")]
        public Input<string>? SpotMaxTotalPrice { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-spotprice
        /// </summary>
        [Input("SpotPrice")]
        public Input<string>? SpotPrice { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-targetcapacity
        /// </summary>
        [Input("TargetCapacity", required: true)]
        public Input<int> TargetCapacity { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-terminateinstanceswithexpiration
        /// </summary>
        [Input("TerminateInstancesWithExpiration")]
        public Input<bool>? TerminateInstancesWithExpiration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-type
        /// </summary>
        [Input("Type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-validfrom
        /// </summary>
        [Input("ValidFrom")]
        public Input<string>? ValidFrom { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-validuntil
        /// </summary>
        [Input("ValidUntil")]
        public Input<string>? ValidUntil { get; set; }

        public SpotFleetSpotFleetRequestConfigDataArgs()
        {
        }
    }
}