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
    public sealed class AnalysisFunnelChartSortConfiguration
    {
        public readonly Outputs.AnalysisItemsLimitConfiguration? CategoryItemsLimit;
        public readonly ImmutableArray<Outputs.AnalysisFieldSortOptions> CategorySort;

        [OutputConstructor]
        private AnalysisFunnelChartSortConfiguration(
            Outputs.AnalysisItemsLimitConfiguration? categoryItemsLimit,

            ImmutableArray<Outputs.AnalysisFieldSortOptions> categorySort)
        {
            CategoryItemsLimit = categoryItemsLimit;
            CategorySort = categorySort;
        }
    }
}
