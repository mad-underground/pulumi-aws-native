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
    public sealed class AnalysisGaugeChartConditionalFormattingOption
    {
        public readonly Outputs.AnalysisGaugeChartArcConditionalFormatting? Arc;
        public readonly Outputs.AnalysisGaugeChartPrimaryValueConditionalFormatting? PrimaryValue;

        [OutputConstructor]
        private AnalysisGaugeChartConditionalFormattingOption(
            Outputs.AnalysisGaugeChartArcConditionalFormatting? arc,

            Outputs.AnalysisGaugeChartPrimaryValueConditionalFormatting? primaryValue)
        {
            Arc = arc;
            PrimaryValue = primaryValue;
        }
    }
}
