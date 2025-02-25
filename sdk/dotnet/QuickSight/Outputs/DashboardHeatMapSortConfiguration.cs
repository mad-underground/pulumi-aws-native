// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardHeatMapSortConfiguration
    {
        public readonly Outputs.DashboardItemsLimitConfiguration? HeatMapColumnItemsLimitConfiguration;
        public readonly ImmutableArray<Outputs.DashboardFieldSortOptions> HeatMapColumnSort;
        public readonly Outputs.DashboardItemsLimitConfiguration? HeatMapRowItemsLimitConfiguration;
        public readonly ImmutableArray<Outputs.DashboardFieldSortOptions> HeatMapRowSort;

        [OutputConstructor]
        private DashboardHeatMapSortConfiguration(
            Outputs.DashboardItemsLimitConfiguration? heatMapColumnItemsLimitConfiguration,

            ImmutableArray<Outputs.DashboardFieldSortOptions> heatMapColumnSort,

            Outputs.DashboardItemsLimitConfiguration? heatMapRowItemsLimitConfiguration,

            ImmutableArray<Outputs.DashboardFieldSortOptions> heatMapRowSort)
        {
            HeatMapColumnItemsLimitConfiguration = heatMapColumnItemsLimitConfiguration;
            HeatMapColumnSort = heatMapColumnSort;
            HeatMapRowItemsLimitConfiguration = heatMapRowItemsLimitConfiguration;
            HeatMapRowSort = heatMapRowSort;
        }
    }
}
