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
    public sealed class TemplateTreeMapConfiguration
    {
        public readonly Outputs.TemplateChartAxisLabelOptions? ColorLabelOptions;
        public readonly Outputs.TemplateColorScale? ColorScale;
        public readonly Outputs.TemplateDataLabelOptions? DataLabels;
        public readonly Outputs.TemplateTreeMapFieldWells? FieldWells;
        public readonly Outputs.TemplateChartAxisLabelOptions? GroupLabelOptions;
        public readonly Outputs.TemplateLegendOptions? Legend;
        public readonly Outputs.TemplateChartAxisLabelOptions? SizeLabelOptions;
        public readonly Outputs.TemplateTreeMapSortConfiguration? SortConfiguration;
        public readonly Outputs.TemplateTooltipOptions? Tooltip;

        [OutputConstructor]
        private TemplateTreeMapConfiguration(
            Outputs.TemplateChartAxisLabelOptions? colorLabelOptions,

            Outputs.TemplateColorScale? colorScale,

            Outputs.TemplateDataLabelOptions? dataLabels,

            Outputs.TemplateTreeMapFieldWells? fieldWells,

            Outputs.TemplateChartAxisLabelOptions? groupLabelOptions,

            Outputs.TemplateLegendOptions? legend,

            Outputs.TemplateChartAxisLabelOptions? sizeLabelOptions,

            Outputs.TemplateTreeMapSortConfiguration? sortConfiguration,

            Outputs.TemplateTooltipOptions? tooltip)
        {
            ColorLabelOptions = colorLabelOptions;
            ColorScale = colorScale;
            DataLabels = dataLabels;
            FieldWells = fieldWells;
            GroupLabelOptions = groupLabelOptions;
            Legend = legend;
            SizeLabelOptions = sizeLabelOptions;
            SortConfiguration = sortConfiguration;
            Tooltip = tooltip;
        }
    }
}