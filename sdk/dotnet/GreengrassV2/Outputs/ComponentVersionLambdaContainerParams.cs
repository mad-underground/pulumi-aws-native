// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2.Outputs
{

    [OutputType]
    public sealed class ComponentVersionLambdaContainerParams
    {
        public readonly ImmutableArray<Outputs.ComponentVersionLambdaDeviceMount> Devices;
        public readonly int? MemorySizeInKb;
        public readonly bool? MountRoSysfs;
        public readonly ImmutableArray<Outputs.ComponentVersionLambdaVolumeMount> Volumes;

        [OutputConstructor]
        private ComponentVersionLambdaContainerParams(
            ImmutableArray<Outputs.ComponentVersionLambdaDeviceMount> devices,

            int? memorySizeInKb,

            bool? mountRoSysfs,

            ImmutableArray<Outputs.ComponentVersionLambdaVolumeMount> volumes)
        {
            Devices = devices;
            MemorySizeInKb = memorySizeInKb;
            MountRoSysfs = mountRoSysfs;
            Volumes = volumes;
        }
    }
}
