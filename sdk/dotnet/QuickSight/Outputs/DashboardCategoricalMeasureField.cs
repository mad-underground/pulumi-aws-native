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
    public sealed class DashboardCategoricalMeasureField
    {
        public readonly Pulumi.AwsNative.QuickSight.DashboardCategoricalAggregationFunction? AggregationFunction;
        public readonly Outputs.DashboardColumnIdentifier Column;
        public readonly string FieldId;
        public readonly Outputs.DashboardStringFormatConfiguration? FormatConfiguration;

        [OutputConstructor]
        private DashboardCategoricalMeasureField(
            Pulumi.AwsNative.QuickSight.DashboardCategoricalAggregationFunction? aggregationFunction,

            Outputs.DashboardColumnIdentifier column,

            string fieldId,

            Outputs.DashboardStringFormatConfiguration? formatConfiguration)
        {
            AggregationFunction = aggregationFunction;
            Column = column;
            FieldId = fieldId;
            FormatConfiguration = formatConfiguration;
        }
    }
}
