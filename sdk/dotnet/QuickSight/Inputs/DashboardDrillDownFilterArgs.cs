// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardDrillDownFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("categoryFilter")]
        public Input<Inputs.DashboardCategoryDrillDownFilterArgs>? CategoryFilter { get; set; }

        [Input("numericEqualityFilter")]
        public Input<Inputs.DashboardNumericEqualityDrillDownFilterArgs>? NumericEqualityFilter { get; set; }

        [Input("timeRangeFilter")]
        public Input<Inputs.DashboardTimeRangeDrillDownFilterArgs>? TimeRangeFilter { get; set; }

        public DashboardDrillDownFilterArgs()
        {
        }
        public static new DashboardDrillDownFilterArgs Empty => new DashboardDrillDownFilterArgs();
    }
}