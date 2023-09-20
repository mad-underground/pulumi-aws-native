// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    /// <summary>
    /// Visibility Metric of the RuleGroup.
    /// </summary>
    public sealed class RuleGroupVisibilityConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudWatchMetricsEnabled", required: true)]
        public Input<bool> CloudWatchMetricsEnabled { get; set; } = null!;

        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        [Input("sampledRequestsEnabled", required: true)]
        public Input<bool> SampledRequestsEnabled { get; set; } = null!;

        public RuleGroupVisibilityConfigArgs()
        {
        }
        public static new RuleGroupVisibilityConfigArgs Empty => new RuleGroupVisibilityConfigArgs();
    }
}