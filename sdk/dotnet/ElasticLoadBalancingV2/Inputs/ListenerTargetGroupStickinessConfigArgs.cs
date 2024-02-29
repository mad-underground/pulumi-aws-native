// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Inputs
{

    /// <summary>
    /// Information about the target group stickiness for a rule.
    /// </summary>
    public sealed class ListenerTargetGroupStickinessConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time period, in seconds, during which requests from a client should be routed to the same target group. The range is 1-604800 seconds (7 days).
        /// </summary>
        [Input("durationSeconds")]
        public Input<int>? DurationSeconds { get; set; }

        /// <summary>
        /// Indicates whether target group stickiness is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public ListenerTargetGroupStickinessConfigArgs()
        {
        }
        public static new ListenerTargetGroupStickinessConfigArgs Empty => new ListenerTargetGroupStickinessConfigArgs();
    }
}
