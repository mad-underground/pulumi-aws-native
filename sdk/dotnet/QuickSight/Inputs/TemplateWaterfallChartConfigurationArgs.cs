// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateWaterfallChartConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("categoryAxisDisplayOptions")]
        public Input<Inputs.TemplateAxisDisplayOptionsArgs>? CategoryAxisDisplayOptions { get; set; }

        [Input("categoryAxisLabelOptions")]
        public Input<Inputs.TemplateChartAxisLabelOptionsArgs>? CategoryAxisLabelOptions { get; set; }

        [Input("dataLabels")]
        public Input<Inputs.TemplateDataLabelOptionsArgs>? DataLabels { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.TemplateWaterfallChartFieldWellsArgs>? FieldWells { get; set; }

        [Input("legend")]
        public Input<Inputs.TemplateLegendOptionsArgs>? Legend { get; set; }

        [Input("primaryYAxisDisplayOptions")]
        public Input<Inputs.TemplateAxisDisplayOptionsArgs>? PrimaryYAxisDisplayOptions { get; set; }

        [Input("primaryYAxisLabelOptions")]
        public Input<Inputs.TemplateChartAxisLabelOptionsArgs>? PrimaryYAxisLabelOptions { get; set; }

        [Input("sortConfiguration")]
        public Input<Inputs.TemplateWaterfallChartSortConfigurationArgs>? SortConfiguration { get; set; }

        [Input("visualPalette")]
        public Input<Inputs.TemplateVisualPaletteArgs>? VisualPalette { get; set; }

        [Input("waterfallChartOptions")]
        public Input<Inputs.TemplateWaterfallChartOptionsArgs>? WaterfallChartOptions { get; set; }

        public TemplateWaterfallChartConfigurationArgs()
        {
        }
        public static new TemplateWaterfallChartConfigurationArgs Empty => new TemplateWaterfallChartConfigurationArgs();
    }
}
