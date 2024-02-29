// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    /// <summary>
    /// Specifies a block device mapping for a launch template. You must specify ``DeviceName`` plus exactly one of the following properties: ``Ebs``, ``NoDevice``, or ``VirtualName``.
    ///   ``BlockDeviceMapping`` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html).
    /// </summary>
    public sealed class LaunchTemplateBlockDeviceMappingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The device name (for example, /dev/sdh or xvdh).
        /// </summary>
        [Input("deviceName")]
        public Input<string>? DeviceName { get; set; }

        /// <summary>
        /// Parameters used to automatically set up EBS volumes when the instance is launched.
        /// </summary>
        [Input("ebs")]
        public Input<Inputs.LaunchTemplateEbsArgs>? Ebs { get; set; }

        /// <summary>
        /// To omit the device from the block device mapping, specify an empty string.
        /// </summary>
        [Input("noDevice")]
        public Input<string>? NoDevice { get; set; }

        /// <summary>
        /// The virtual device name (ephemeralN). Instance store volumes are numbered starting from 0. An instance type with 2 available instance store volumes can specify mappings for ephemeral0 and ephemeral1. The number of available instance store volumes depends on the instance type. After you connect to the instance, you must mount the volume.
        /// </summary>
        [Input("virtualName")]
        public Input<string>? VirtualName { get; set; }

        public LaunchTemplateBlockDeviceMappingArgs()
        {
        }
        public static new LaunchTemplateBlockDeviceMappingArgs Empty => new LaunchTemplateBlockDeviceMappingArgs();
    }
}
