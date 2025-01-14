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
    public sealed class TemplateBoxPlotChartConfiguration
    {
        public readonly Outputs.TemplateBoxPlotOptions? BoxPlotOptions;
        public readonly Outputs.TemplateAxisDisplayOptions? CategoryAxis;
        public readonly Outputs.TemplateChartAxisLabelOptions? CategoryLabelOptions;
        public readonly Outputs.TemplateBoxPlotFieldWells? FieldWells;
        public readonly Outputs.TemplateLegendOptions? Legend;
        public readonly Outputs.TemplateAxisDisplayOptions? PrimaryYAxisDisplayOptions;
        public readonly Outputs.TemplateChartAxisLabelOptions? PrimaryYAxisLabelOptions;
        public readonly ImmutableArray<Outputs.TemplateReferenceLine> ReferenceLines;
        public readonly Outputs.TemplateBoxPlotSortConfiguration? SortConfiguration;
        public readonly Outputs.TemplateTooltipOptions? Tooltip;
        public readonly Outputs.TemplateVisualPalette? VisualPalette;

        [OutputConstructor]
        private TemplateBoxPlotChartConfiguration(
            Outputs.TemplateBoxPlotOptions? boxPlotOptions,

            Outputs.TemplateAxisDisplayOptions? categoryAxis,

            Outputs.TemplateChartAxisLabelOptions? categoryLabelOptions,

            Outputs.TemplateBoxPlotFieldWells? fieldWells,

            Outputs.TemplateLegendOptions? legend,

            Outputs.TemplateAxisDisplayOptions? primaryYAxisDisplayOptions,

            Outputs.TemplateChartAxisLabelOptions? primaryYAxisLabelOptions,

            ImmutableArray<Outputs.TemplateReferenceLine> referenceLines,

            Outputs.TemplateBoxPlotSortConfiguration? sortConfiguration,

            Outputs.TemplateTooltipOptions? tooltip,

            Outputs.TemplateVisualPalette? visualPalette)
        {
            BoxPlotOptions = boxPlotOptions;
            CategoryAxis = categoryAxis;
            CategoryLabelOptions = categoryLabelOptions;
            FieldWells = fieldWells;
            Legend = legend;
            PrimaryYAxisDisplayOptions = primaryYAxisDisplayOptions;
            PrimaryYAxisLabelOptions = primaryYAxisLabelOptions;
            ReferenceLines = referenceLines;
            SortConfiguration = sortConfiguration;
            Tooltip = tooltip;
            VisualPalette = visualPalette;
        }
    }
}
