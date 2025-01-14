// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisHeatMapConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("colorScale")]
        public Input<Inputs.AnalysisColorScaleArgs>? ColorScale { get; set; }

        [Input("columnLabelOptions")]
        public Input<Inputs.AnalysisChartAxisLabelOptionsArgs>? ColumnLabelOptions { get; set; }

        [Input("dataLabels")]
        public Input<Inputs.AnalysisDataLabelOptionsArgs>? DataLabels { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.AnalysisHeatMapFieldWellsArgs>? FieldWells { get; set; }

        [Input("legend")]
        public Input<Inputs.AnalysisLegendOptionsArgs>? Legend { get; set; }

        [Input("rowLabelOptions")]
        public Input<Inputs.AnalysisChartAxisLabelOptionsArgs>? RowLabelOptions { get; set; }

        [Input("sortConfiguration")]
        public Input<Inputs.AnalysisHeatMapSortConfigurationArgs>? SortConfiguration { get; set; }

        [Input("tooltip")]
        public Input<Inputs.AnalysisTooltipOptionsArgs>? Tooltip { get; set; }

        public AnalysisHeatMapConfigurationArgs()
        {
        }
        public static new AnalysisHeatMapConfigurationArgs Empty => new AnalysisHeatMapConfigurationArgs();
    }
}
