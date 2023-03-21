// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardMetricComparisonComputationArgs : global::Pulumi.ResourceArgs
    {
        [Input("computationId", required: true)]
        public Input<string> ComputationId { get; set; } = null!;

        [Input("fromValue", required: true)]
        public Input<Inputs.DashboardMeasureFieldArgs> FromValue { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("targetValue", required: true)]
        public Input<Inputs.DashboardMeasureFieldArgs> TargetValue { get; set; } = null!;

        [Input("time", required: true)]
        public Input<Inputs.DashboardDimensionFieldArgs> Time { get; set; } = null!;

        public DashboardMetricComparisonComputationArgs()
        {
        }
        public static new DashboardMetricComparisonComputationArgs Empty => new DashboardMetricComparisonComputationArgs();
    }
}
