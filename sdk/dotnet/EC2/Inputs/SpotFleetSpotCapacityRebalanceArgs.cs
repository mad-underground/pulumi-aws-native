// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class SpotFleetSpotCapacityRebalanceArgs : Pulumi.ResourceArgs
    {
        [Input("replacementStrategy")]
        public Input<Pulumi.AwsNative.EC2.SpotFleetSpotCapacityRebalanceReplacementStrategy>? ReplacementStrategy { get; set; }

        [Input("terminationDelay")]
        public Input<int>? TerminationDelay { get; set; }

        public SpotFleetSpotCapacityRebalanceArgs()
        {
        }
    }
}
