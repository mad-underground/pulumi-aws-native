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
    public sealed class DashboardTopBottomFilter
    {
        public readonly ImmutableArray<Outputs.DashboardAggregationSortConfiguration> AggregationSortConfigurations;
        public readonly Outputs.DashboardColumnIdentifier Column;
        public readonly string FilterId;
        public readonly double? Limit;
        public readonly string? ParameterName;
        public readonly Pulumi.AwsNative.QuickSight.DashboardTimeGranularity? TimeGranularity;

        [OutputConstructor]
        private DashboardTopBottomFilter(
            ImmutableArray<Outputs.DashboardAggregationSortConfiguration> aggregationSortConfigurations,

            Outputs.DashboardColumnIdentifier column,

            string filterId,

            double? limit,

            string? parameterName,

            Pulumi.AwsNative.QuickSight.DashboardTimeGranularity? timeGranularity)
        {
            AggregationSortConfigurations = aggregationSortConfigurations;
            Column = column;
            FilterId = filterId;
            Limit = limit;
            ParameterName = parameterName;
            TimeGranularity = timeGranularity;
        }
    }
}
