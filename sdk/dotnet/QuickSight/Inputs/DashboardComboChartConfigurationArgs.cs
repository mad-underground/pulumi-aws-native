// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardComboChartConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("barDataLabels")]
        public Input<Inputs.DashboardDataLabelOptionsArgs>? BarDataLabels { get; set; }

        [Input("barsArrangement")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardBarsArrangement>? BarsArrangement { get; set; }

        [Input("categoryAxis")]
        public Input<Inputs.DashboardAxisDisplayOptionsArgs>? CategoryAxis { get; set; }

        [Input("categoryLabelOptions")]
        public Input<Inputs.DashboardChartAxisLabelOptionsArgs>? CategoryLabelOptions { get; set; }

        [Input("colorLabelOptions")]
        public Input<Inputs.DashboardChartAxisLabelOptionsArgs>? ColorLabelOptions { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.DashboardComboChartFieldWellsArgs>? FieldWells { get; set; }

        [Input("legend")]
        public Input<Inputs.DashboardLegendOptionsArgs>? Legend { get; set; }

        [Input("lineDataLabels")]
        public Input<Inputs.DashboardDataLabelOptionsArgs>? LineDataLabels { get; set; }

        [Input("primaryYAxisDisplayOptions")]
        public Input<Inputs.DashboardAxisDisplayOptionsArgs>? PrimaryYAxisDisplayOptions { get; set; }

        [Input("primaryYAxisLabelOptions")]
        public Input<Inputs.DashboardChartAxisLabelOptionsArgs>? PrimaryYAxisLabelOptions { get; set; }

        [Input("referenceLines")]
        private InputList<Inputs.DashboardReferenceLineArgs>? _referenceLines;
        public InputList<Inputs.DashboardReferenceLineArgs> ReferenceLines
        {
            get => _referenceLines ?? (_referenceLines = new InputList<Inputs.DashboardReferenceLineArgs>());
            set => _referenceLines = value;
        }

        [Input("secondaryYAxisDisplayOptions")]
        public Input<Inputs.DashboardAxisDisplayOptionsArgs>? SecondaryYAxisDisplayOptions { get; set; }

        [Input("secondaryYAxisLabelOptions")]
        public Input<Inputs.DashboardChartAxisLabelOptionsArgs>? SecondaryYAxisLabelOptions { get; set; }

        [Input("sortConfiguration")]
        public Input<Inputs.DashboardComboChartSortConfigurationArgs>? SortConfiguration { get; set; }

        [Input("tooltip")]
        public Input<Inputs.DashboardTooltipOptionsArgs>? Tooltip { get; set; }

        [Input("visualPalette")]
        public Input<Inputs.DashboardVisualPaletteArgs>? VisualPalette { get; set; }

        public DashboardComboChartConfigurationArgs()
        {
        }
        public static new DashboardComboChartConfigurationArgs Empty => new DashboardComboChartConfigurationArgs();
    }
}
