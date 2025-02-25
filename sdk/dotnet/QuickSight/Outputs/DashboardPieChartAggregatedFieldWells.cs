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
    public sealed class DashboardPieChartAggregatedFieldWells
    {
        public readonly ImmutableArray<Outputs.DashboardDimensionField> Category;
        public readonly ImmutableArray<Outputs.DashboardDimensionField> SmallMultiples;
        public readonly ImmutableArray<Outputs.DashboardMeasureField> Values;

        [OutputConstructor]
        private DashboardPieChartAggregatedFieldWells(
            ImmutableArray<Outputs.DashboardDimensionField> category,

            ImmutableArray<Outputs.DashboardDimensionField> smallMultiples,

            ImmutableArray<Outputs.DashboardMeasureField> values)
        {
            Category = category;
            SmallMultiples = smallMultiples;
            Values = values;
        }
    }
}
