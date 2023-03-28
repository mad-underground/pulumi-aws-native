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
    public sealed class AnalysisWaterfallChartAggregatedFieldWells
    {
        public readonly ImmutableArray<Outputs.AnalysisDimensionField> Breakdowns;
        public readonly ImmutableArray<Outputs.AnalysisDimensionField> Categories;
        public readonly ImmutableArray<Outputs.AnalysisMeasureField> Values;

        [OutputConstructor]
        private AnalysisWaterfallChartAggregatedFieldWells(
            ImmutableArray<Outputs.AnalysisDimensionField> breakdowns,

            ImmutableArray<Outputs.AnalysisDimensionField> categories,

            ImmutableArray<Outputs.AnalysisMeasureField> values)
        {
            Breakdowns = breakdowns;
            Categories = categories;
            Values = values;
        }
    }
}