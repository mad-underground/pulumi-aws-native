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
    public sealed class DashboardGaugeChartFieldWells
    {
        public readonly ImmutableArray<Outputs.DashboardMeasureField> TargetValues;
        public readonly ImmutableArray<Outputs.DashboardMeasureField> Values;

        [OutputConstructor]
        private DashboardGaugeChartFieldWells(
            ImmutableArray<Outputs.DashboardMeasureField> targetValues,

            ImmutableArray<Outputs.DashboardMeasureField> values)
        {
            TargetValues = targetValues;
            Values = values;
        }
    }
}
