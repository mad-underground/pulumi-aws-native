// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Outputs
{

    [OutputType]
    public sealed class ConfigSpectrumConfig
    {
        public readonly Outputs.ConfigFrequencyBandwidth? Bandwidth;
        public readonly Outputs.ConfigFrequency? CenterFrequency;
        public readonly Pulumi.AwsNative.GroundStation.ConfigPolarization? Polarization;

        [OutputConstructor]
        private ConfigSpectrumConfig(
            Outputs.ConfigFrequencyBandwidth? bandwidth,

            Outputs.ConfigFrequency? centerFrequency,

            Pulumi.AwsNative.GroundStation.ConfigPolarization? polarization)
        {
            Bandwidth = bandwidth;
            CenterFrequency = centerFrequency;
            Polarization = polarization;
        }
    }
}
