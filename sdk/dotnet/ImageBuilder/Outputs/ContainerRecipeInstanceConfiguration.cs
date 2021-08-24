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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-instanceconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class ContainerRecipeInstanceConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-instanceconfiguration.html#cfn-imagebuilder-containerrecipe-instanceconfiguration-blockdevicemappings
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerRecipeInstanceBlockDeviceMapping> BlockDeviceMappings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-instanceconfiguration.html#cfn-imagebuilder-containerrecipe-instanceconfiguration-image
        /// </summary>
        public readonly string? Image;

        [OutputConstructor]
        private ContainerRecipeInstanceConfiguration(
            ImmutableArray<Outputs.ContainerRecipeInstanceBlockDeviceMapping> blockDeviceMappings,

            string? image)
        {
            BlockDeviceMappings = blockDeviceMappings;
            Image = image;
        }
    }
}