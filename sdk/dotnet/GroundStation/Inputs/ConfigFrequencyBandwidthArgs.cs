// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Inputs
{

    public sealed class ConfigFrequencyBandwidthArgs : global::Pulumi.ResourceArgs
    {
        [Input("units")]
        public Input<Pulumi.AwsNative.GroundStation.ConfigBandwidthUnits>? Units { get; set; }

        [Input("value")]
        public Input<double>? Value { get; set; }

        public ConfigFrequencyBandwidthArgs()
        {
        }
        public static new ConfigFrequencyBandwidthArgs Empty => new ConfigFrequencyBandwidthArgs();
    }
}
