// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisHistogramConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("binOptions")]
        public Input<Inputs.AnalysisHistogramBinOptionsArgs>? BinOptions { get; set; }

        [Input("dataLabels")]
        public Input<Inputs.AnalysisDataLabelOptionsArgs>? DataLabels { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.AnalysisHistogramFieldWellsArgs>? FieldWells { get; set; }

        [Input("tooltip")]
        public Input<Inputs.AnalysisTooltipOptionsArgs>? Tooltip { get; set; }

        [Input("visualPalette")]
        public Input<Inputs.AnalysisVisualPaletteArgs>? VisualPalette { get; set; }

        [Input("xAxisDisplayOptions")]
        public Input<Inputs.AnalysisAxisDisplayOptionsArgs>? XAxisDisplayOptions { get; set; }

        [Input("xAxisLabelOptions")]
        public Input<Inputs.AnalysisChartAxisLabelOptionsArgs>? XAxisLabelOptions { get; set; }

        [Input("yAxisDisplayOptions")]
        public Input<Inputs.AnalysisAxisDisplayOptionsArgs>? YAxisDisplayOptions { get; set; }

        public AnalysisHistogramConfigurationArgs()
        {
        }
        public static new AnalysisHistogramConfigurationArgs Empty => new AnalysisHistogramConfigurationArgs();
    }
}