// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.OpsWorks.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html
    /// </summary>
    public sealed class InstancePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-agentversion
        /// </summary>
        [Input("AgentVersion")]
        public Input<string>? AgentVersion { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-amiid
        /// </summary>
        [Input("AmiId")]
        public Input<string>? AmiId { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-architecture
        /// </summary>
        [Input("Architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-autoscalingtype
        /// </summary>
        [Input("AutoScalingType")]
        public Input<string>? AutoScalingType { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-availabilityzone
        /// </summary>
        [Input("AvailabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("BlockDeviceMappings")]
        private InputList<Inputs.InstanceBlockDeviceMappingArgs>? _BlockDeviceMappings;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-blockdevicemappings
        /// </summary>
        public InputList<Inputs.InstanceBlockDeviceMappingArgs> BlockDeviceMappings
        {
            get => _BlockDeviceMappings ?? (_BlockDeviceMappings = new InputList<Inputs.InstanceBlockDeviceMappingArgs>());
            set => _BlockDeviceMappings = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-ebsoptimized
        /// </summary>
        [Input("EbsOptimized")]
        public Input<bool>? EbsOptimized { get; set; }

        [Input("ElasticIps")]
        private InputList<string>? _ElasticIps;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-elasticips
        /// </summary>
        public InputList<string> ElasticIps
        {
            get => _ElasticIps ?? (_ElasticIps = new InputList<string>());
            set => _ElasticIps = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-hostname
        /// </summary>
        [Input("Hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-installupdatesonboot
        /// </summary>
        [Input("InstallUpdatesOnBoot")]
        public Input<bool>? InstallUpdatesOnBoot { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-instancetype
        /// </summary>
        [Input("InstanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        [Input("LayerIds", required: true)]
        private InputList<string>? _LayerIds;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-layerids
        /// </summary>
        public InputList<string> LayerIds
        {
            get => _LayerIds ?? (_LayerIds = new InputList<string>());
            set => _LayerIds = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-os
        /// </summary>
        [Input("Os")]
        public Input<string>? Os { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-rootdevicetype
        /// </summary>
        [Input("RootDeviceType")]
        public Input<string>? RootDeviceType { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-sshkeyname
        /// </summary>
        [Input("SshKeyName")]
        public Input<string>? SshKeyName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-stackid
        /// </summary>
        [Input("StackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-subnetid
        /// </summary>
        [Input("SubnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-tenancy
        /// </summary>
        [Input("Tenancy")]
        public Input<string>? Tenancy { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-timebasedautoscaling
        /// </summary>
        [Input("TimeBasedAutoScaling")]
        public Input<Inputs.InstanceTimeBasedAutoScalingArgs>? TimeBasedAutoScaling { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-virtualizationtype
        /// </summary>
        [Input("VirtualizationType")]
        public Input<string>? VirtualizationType { get; set; }

        [Input("Volumes")]
        private InputList<string>? _Volumes;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-volumes
        /// </summary>
        public InputList<string> Volumes
        {
            get => _Volumes ?? (_Volumes = new InputList<string>());
            set => _Volumes = value;
        }

        public InstancePropertiesArgs()
        {
        }
    }
}