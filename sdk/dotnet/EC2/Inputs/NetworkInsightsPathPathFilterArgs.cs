// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class NetworkInsightsPathPathFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationAddress")]
        public Input<string>? DestinationAddress { get; set; }

        [Input("destinationPortRange")]
        public Input<Inputs.NetworkInsightsPathFilterPortRangeArgs>? DestinationPortRange { get; set; }

        [Input("sourceAddress")]
        public Input<string>? SourceAddress { get; set; }

        [Input("sourcePortRange")]
        public Input<Inputs.NetworkInsightsPathFilterPortRangeArgs>? SourcePortRange { get; set; }

        public NetworkInsightsPathPathFilterArgs()
        {
        }
        public static new NetworkInsightsPathPathFilterArgs Empty => new NetworkInsightsPathPathFilterArgs();
    }
}