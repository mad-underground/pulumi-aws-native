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
    public sealed class AnalysisKPISortConfiguration
    {
        public readonly ImmutableArray<Outputs.AnalysisFieldSortOptions> TrendGroupSort;

        [OutputConstructor]
        private AnalysisKPISortConfiguration(ImmutableArray<Outputs.AnalysisFieldSortOptions> trendGroupSort)
        {
            TrendGroupSort = trendGroupSort;
        }
    }
}