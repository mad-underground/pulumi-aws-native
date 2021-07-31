// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html
    /// </summary>
    [OutputType]
    public sealed class ImageRecipeInstanceBlockDeviceMapping
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-devicename
        /// </summary>
        public readonly string? DeviceName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-ebs
        /// </summary>
        public readonly Outputs.ImageRecipeEbsInstanceBlockDeviceSpecification? Ebs;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-nodevice
        /// </summary>
        public readonly string? NoDevice;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html#cfn-imagebuilder-imagerecipe-instanceblockdevicemapping-virtualname
        /// </summary>
        public readonly string? VirtualName;

        [OutputConstructor]
        private ImageRecipeInstanceBlockDeviceMapping(
            string? deviceName,

            Outputs.ImageRecipeEbsInstanceBlockDeviceSpecification? ebs,

            string? noDevice,

            string? virtualName)
        {
            DeviceName = deviceName;
            Ebs = ebs;
            NoDevice = noDevice;
            VirtualName = virtualName;
        }
    }
}
