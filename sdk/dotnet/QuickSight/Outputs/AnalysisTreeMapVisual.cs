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
    public sealed class AnalysisTreeMapVisual
    {
        public readonly ImmutableArray<Outputs.AnalysisVisualCustomAction> Actions;
        public readonly Outputs.AnalysisTreeMapConfiguration? ChartConfiguration;
        public readonly ImmutableArray<Outputs.AnalysisColumnHierarchy> ColumnHierarchies;
        public readonly Outputs.AnalysisVisualSubtitleLabelOptions? Subtitle;
        public readonly Outputs.AnalysisVisualTitleLabelOptions? Title;
        public readonly string VisualId;

        [OutputConstructor]
        private AnalysisTreeMapVisual(
            ImmutableArray<Outputs.AnalysisVisualCustomAction> actions,

            Outputs.AnalysisTreeMapConfiguration? chartConfiguration,

            ImmutableArray<Outputs.AnalysisColumnHierarchy> columnHierarchies,

            Outputs.AnalysisVisualSubtitleLabelOptions? subtitle,

            Outputs.AnalysisVisualTitleLabelOptions? title,

            string visualId)
        {
            Actions = actions;
            ChartConfiguration = chartConfiguration;
            ColumnHierarchies = columnHierarchies;
            Subtitle = subtitle;
            Title = title;
            VisualId = visualId;
        }
    }
}
