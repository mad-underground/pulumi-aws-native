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
    public sealed class TemplatePieChartConfiguration
    {
        public readonly Outputs.TemplateChartAxisLabelOptions? CategoryLabelOptions;
        public readonly ImmutableArray<Outputs.TemplateContributionAnalysisDefault> ContributionAnalysisDefaults;
        public readonly Outputs.TemplateDataLabelOptions? DataLabels;
        public readonly Outputs.TemplateDonutOptions? DonutOptions;
        public readonly Outputs.TemplatePieChartFieldWells? FieldWells;
        public readonly Outputs.TemplateLegendOptions? Legend;
        public readonly Outputs.TemplateSmallMultiplesOptions? SmallMultiplesOptions;
        public readonly Outputs.TemplatePieChartSortConfiguration? SortConfiguration;
        public readonly Outputs.TemplateTooltipOptions? Tooltip;
        public readonly Outputs.TemplateChartAxisLabelOptions? ValueLabelOptions;
        public readonly Outputs.TemplateVisualPalette? VisualPalette;

        [OutputConstructor]
        private TemplatePieChartConfiguration(
            Outputs.TemplateChartAxisLabelOptions? categoryLabelOptions,

            ImmutableArray<Outputs.TemplateContributionAnalysisDefault> contributionAnalysisDefaults,

            Outputs.TemplateDataLabelOptions? dataLabels,

            Outputs.TemplateDonutOptions? donutOptions,

            Outputs.TemplatePieChartFieldWells? fieldWells,

            Outputs.TemplateLegendOptions? legend,

            Outputs.TemplateSmallMultiplesOptions? smallMultiplesOptions,

            Outputs.TemplatePieChartSortConfiguration? sortConfiguration,

            Outputs.TemplateTooltipOptions? tooltip,

            Outputs.TemplateChartAxisLabelOptions? valueLabelOptions,

            Outputs.TemplateVisualPalette? visualPalette)
        {
            CategoryLabelOptions = categoryLabelOptions;
            ContributionAnalysisDefaults = contributionAnalysisDefaults;
            DataLabels = dataLabels;
            DonutOptions = donutOptions;
            FieldWells = fieldWells;
            Legend = legend;
            SmallMultiplesOptions = smallMultiplesOptions;
            SortConfiguration = sortConfiguration;
            Tooltip = tooltip;
            ValueLabelOptions = valueLabelOptions;
            VisualPalette = visualPalette;
        }
    }
}