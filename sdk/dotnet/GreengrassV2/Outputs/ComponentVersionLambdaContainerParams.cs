// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdacontainerparams.html
    /// </summary>
    [OutputType]
    public sealed class ComponentVersionLambdaContainerParams
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdacontainerparams.html#cfn-greengrassv2-componentversion-lambdacontainerparams-devices
        /// </summary>
        public readonly ImmutableArray<Outputs.ComponentVersionLambdaDeviceMount> Devices;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdacontainerparams.html#cfn-greengrassv2-componentversion-lambdacontainerparams-memorysizeinkb
        /// </summary>
        public readonly int? MemorySizeInKB;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdacontainerparams.html#cfn-greengrassv2-componentversion-lambdacontainerparams-mountrosysfs
        /// </summary>
        public readonly bool? MountROSysfs;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdacontainerparams.html#cfn-greengrassv2-componentversion-lambdacontainerparams-volumes
        /// </summary>
        public readonly ImmutableArray<Outputs.ComponentVersionLambdaVolumeMount> Volumes;

        [OutputConstructor]
        private ComponentVersionLambdaContainerParams(
            ImmutableArray<Outputs.ComponentVersionLambdaDeviceMount> devices,

            int? memorySizeInKB,

            bool? mountROSysfs,

            ImmutableArray<Outputs.ComponentVersionLambdaVolumeMount> volumes)
        {
            Devices = devices;
            MemorySizeInKB = memorySizeInKB;
            MountROSysfs = mountROSysfs;
            Volumes = volumes;
        }
    }
}
