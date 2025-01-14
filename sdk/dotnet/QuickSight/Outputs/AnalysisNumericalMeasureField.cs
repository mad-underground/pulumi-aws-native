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
    public sealed class AnalysisNumericalMeasureField
    {
        public readonly Outputs.AnalysisNumericalAggregationFunction? AggregationFunction;
        public readonly Outputs.AnalysisColumnIdentifier Column;
        public readonly string FieldId;
        public readonly Outputs.AnalysisNumberFormatConfiguration? FormatConfiguration;

        [OutputConstructor]
        private AnalysisNumericalMeasureField(
            Outputs.AnalysisNumericalAggregationFunction? aggregationFunction,

            Outputs.AnalysisColumnIdentifier column,

            string fieldId,

            Outputs.AnalysisNumberFormatConfiguration? formatConfiguration)
        {
            AggregationFunction = aggregationFunction;
            Column = column;
            FieldId = fieldId;
            FormatConfiguration = formatConfiguration;
        }
    }
}
