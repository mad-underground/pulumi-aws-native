// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class LaunchTemplateSpotOptionsArgs : Pulumi.ResourceArgs
    {
        [Input("blockDurationMinutes")]
        public Input<int>? BlockDurationMinutes { get; set; }

        [Input("instanceInterruptionBehavior")]
        public Input<string>? InstanceInterruptionBehavior { get; set; }

        [Input("maxPrice")]
        public Input<string>? MaxPrice { get; set; }

        [Input("spotInstanceType")]
        public Input<string>? SpotInstanceType { get; set; }

        [Input("validUntil")]
        public Input<string>? ValidUntil { get; set; }

        public LaunchTemplateSpotOptionsArgs()
        {
        }
    }
}