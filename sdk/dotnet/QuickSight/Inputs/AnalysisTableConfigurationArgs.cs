// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisTableConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldOptions")]
        public Input<Inputs.AnalysisTableFieldOptionsArgs>? FieldOptions { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.AnalysisTableFieldWellsArgs>? FieldWells { get; set; }

        [Input("paginatedReportOptions")]
        public Input<Inputs.AnalysisTablePaginatedReportOptionsArgs>? PaginatedReportOptions { get; set; }

        [Input("sortConfiguration")]
        public Input<Inputs.AnalysisTableSortConfigurationArgs>? SortConfiguration { get; set; }

        [Input("tableInlineVisualizations")]
        private InputList<Inputs.AnalysisTableInlineVisualizationArgs>? _tableInlineVisualizations;
        public InputList<Inputs.AnalysisTableInlineVisualizationArgs> TableInlineVisualizations
        {
            get => _tableInlineVisualizations ?? (_tableInlineVisualizations = new InputList<Inputs.AnalysisTableInlineVisualizationArgs>());
            set => _tableInlineVisualizations = value;
        }

        [Input("tableOptions")]
        public Input<Inputs.AnalysisTableOptionsArgs>? TableOptions { get; set; }

        [Input("totalOptions")]
        public Input<Inputs.AnalysisTotalOptionsArgs>? TotalOptions { get; set; }

        public AnalysisTableConfigurationArgs()
        {
        }
        public static new AnalysisTableConfigurationArgs Empty => new AnalysisTableConfigurationArgs();
    }
}