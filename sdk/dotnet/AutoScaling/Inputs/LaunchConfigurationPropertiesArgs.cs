// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AutoScaling.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html
    /// </summary>
    public sealed class LaunchConfigurationPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cf-as-launchconfig-associatepubip
        /// </summary>
        [Input("AssociatePublicIpAddress")]
        public Input<bool>? AssociatePublicIpAddress { get; set; }

        [Input("BlockDeviceMappings")]
        private InputList<Inputs.LaunchConfigurationBlockDeviceMappingArgs>? _BlockDeviceMappings;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-blockdevicemappings
        /// </summary>
        public InputList<Inputs.LaunchConfigurationBlockDeviceMappingArgs> BlockDeviceMappings
        {
            get => _BlockDeviceMappings ?? (_BlockDeviceMappings = new InputList<Inputs.LaunchConfigurationBlockDeviceMappingArgs>());
            set => _BlockDeviceMappings = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-classiclinkvpcid
        /// </summary>
        [Input("ClassicLinkVPCId")]
        public Input<string>? ClassicLinkVPCId { get; set; }

        [Input("ClassicLinkVPCSecurityGroups")]
        private InputList<string>? _ClassicLinkVPCSecurityGroups;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-classiclinkvpcsecuritygroups
        /// </summary>
        public InputList<string> ClassicLinkVPCSecurityGroups
        {
            get => _ClassicLinkVPCSecurityGroups ?? (_ClassicLinkVPCSecurityGroups = new InputList<string>());
            set => _ClassicLinkVPCSecurityGroups = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-ebsoptimized
        /// </summary>
        [Input("EbsOptimized")]
        public Input<bool>? EbsOptimized { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-iaminstanceprofile
        /// </summary>
        [Input("IamInstanceProfile")]
        public Input<string>? IamInstanceProfile { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-imageid
        /// </summary>
        [Input("ImageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instanceid
        /// </summary>
        [Input("InstanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instancemonitoring
        /// </summary>
        [Input("InstanceMonitoring")]
        public Input<bool>? InstanceMonitoring { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instancetype
        /// </summary>
        [Input("InstanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-kernelid
        /// </summary>
        [Input("KernelId")]
        public Input<string>? KernelId { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-keyname
        /// </summary>
        [Input("KeyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-autoscaling-launchconfig-launchconfigurationname
        /// </summary>
        [Input("LaunchConfigurationName")]
        public Input<string>? LaunchConfigurationName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-autoscaling-launchconfig-metadataoptions
        /// </summary>
        [Input("MetadataOptions")]
        public Input<Inputs.LaunchConfigurationMetadataOptionsArgs>? MetadataOptions { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-placementtenancy
        /// </summary>
        [Input("PlacementTenancy")]
        public Input<string>? PlacementTenancy { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-ramdiskid
        /// </summary>
        [Input("RamDiskId")]
        public Input<string>? RamDiskId { get; set; }

        [Input("SecurityGroups")]
        private InputList<string>? _SecurityGroups;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-securitygroups
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _SecurityGroups ?? (_SecurityGroups = new InputList<string>());
            set => _SecurityGroups = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-spotprice
        /// </summary>
        [Input("SpotPrice")]
        public Input<string>? SpotPrice { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-userdata
        /// </summary>
        [Input("UserData")]
        public Input<string>? UserData { get; set; }

        public LaunchConfigurationPropertiesArgs()
        {
        }
    }
}
