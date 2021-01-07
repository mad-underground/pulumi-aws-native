// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.EMR.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html
    /// </summary>
    public sealed class InstanceFleetConfigPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-clusterid
        /// </summary>
        [Input("ClusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-instancefleettype
        /// </summary>
        [Input("InstanceFleetType", required: true)]
        public Input<string> InstanceFleetType { get; set; } = null!;

        [Input("InstanceTypeConfigs")]
        private InputList<Inputs.InstanceFleetConfigInstanceTypeConfigArgs>? _InstanceTypeConfigs;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-instancetypeconfigs
        /// </summary>
        public InputList<Inputs.InstanceFleetConfigInstanceTypeConfigArgs> InstanceTypeConfigs
        {
            get => _InstanceTypeConfigs ?? (_InstanceTypeConfigs = new InputList<Inputs.InstanceFleetConfigInstanceTypeConfigArgs>());
            set => _InstanceTypeConfigs = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-launchspecifications
        /// </summary>
        [Input("LaunchSpecifications")]
        public Input<Inputs.InstanceFleetConfigInstanceFleetProvisioningSpecificationsArgs>? LaunchSpecifications { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-name
        /// </summary>
        [Input("Name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-targetondemandcapacity
        /// </summary>
        [Input("TargetOnDemandCapacity")]
        public Input<int>? TargetOnDemandCapacity { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-targetspotcapacity
        /// </summary>
        [Input("TargetSpotCapacity")]
        public Input<int>? TargetSpotCapacity { get; set; }

        public InstanceFleetConfigPropertiesArgs()
        {
        }
    }
}