// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateBoxPlotChartConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("boxPlotOptions")]
        public Input<Inputs.TemplateBoxPlotOptionsArgs>? BoxPlotOptions { get; set; }

        [Input("categoryAxis")]
        public Input<Inputs.TemplateAxisDisplayOptionsArgs>? CategoryAxis { get; set; }

        [Input("categoryLabelOptions")]
        public Input<Inputs.TemplateChartAxisLabelOptionsArgs>? CategoryLabelOptions { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.TemplateBoxPlotFieldWellsArgs>? FieldWells { get; set; }

        [Input("legend")]
        public Input<Inputs.TemplateLegendOptionsArgs>? Legend { get; set; }

        [Input("primaryYAxisDisplayOptions")]
        public Input<Inputs.TemplateAxisDisplayOptionsArgs>? PrimaryYAxisDisplayOptions { get; set; }

        [Input("primaryYAxisLabelOptions")]
        public Input<Inputs.TemplateChartAxisLabelOptionsArgs>? PrimaryYAxisLabelOptions { get; set; }

        [Input("referenceLines")]
        private InputList<Inputs.TemplateReferenceLineArgs>? _referenceLines;
        public InputList<Inputs.TemplateReferenceLineArgs> ReferenceLines
        {
            get => _referenceLines ?? (_referenceLines = new InputList<Inputs.TemplateReferenceLineArgs>());
            set => _referenceLines = value;
        }

        [Input("sortConfiguration")]
        public Input<Inputs.TemplateBoxPlotSortConfigurationArgs>? SortConfiguration { get; set; }

        [Input("tooltip")]
        public Input<Inputs.TemplateTooltipOptionsArgs>? Tooltip { get; set; }

        [Input("visualPalette")]
        public Input<Inputs.TemplateVisualPaletteArgs>? VisualPalette { get; set; }

        public TemplateBoxPlotChartConfigurationArgs()
        {
        }
        public static new TemplateBoxPlotChartConfigurationArgs Empty => new TemplateBoxPlotChartConfigurationArgs();
    }
}
