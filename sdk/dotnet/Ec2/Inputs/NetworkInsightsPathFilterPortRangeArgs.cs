// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class NetworkInsightsPathFilterPortRangeArgs : global::Pulumi.ResourceArgs
    {
        [Input("fromPort")]
        public Input<int>? FromPort { get; set; }

        [Input("toPort")]
        public Input<int>? ToPort { get; set; }

        public NetworkInsightsPathFilterPortRangeArgs()
        {
        }
        public static new NetworkInsightsPathFilterPortRangeArgs Empty => new NetworkInsightsPathFilterPortRangeArgs();
    }
}