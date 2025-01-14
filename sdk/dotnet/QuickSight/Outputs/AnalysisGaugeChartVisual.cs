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
    public sealed class AnalysisGaugeChartVisual
    {
        public readonly ImmutableArray<Outputs.AnalysisVisualCustomAction> Actions;
        public readonly Outputs.AnalysisGaugeChartConfiguration? ChartConfiguration;
        public readonly Outputs.AnalysisGaugeChartConditionalFormatting? ConditionalFormatting;
        public readonly Outputs.AnalysisVisualSubtitleLabelOptions? Subtitle;
        public readonly Outputs.AnalysisVisualTitleLabelOptions? Title;
        public readonly string VisualId;

        [OutputConstructor]
        private AnalysisGaugeChartVisual(
            ImmutableArray<Outputs.AnalysisVisualCustomAction> actions,

            Outputs.AnalysisGaugeChartConfiguration? chartConfiguration,

            Outputs.AnalysisGaugeChartConditionalFormatting? conditionalFormatting,

            Outputs.AnalysisVisualSubtitleLabelOptions? subtitle,

            Outputs.AnalysisVisualTitleLabelOptions? title,

            string visualId)
        {
            Actions = actions;
            ChartConfiguration = chartConfiguration;
            ConditionalFormatting = conditionalFormatting;
            Subtitle = subtitle;
            Title = title;
            VisualId = visualId;
        }
    }
}
