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
    public sealed class AnalysisTableConfiguration
    {
        public readonly Outputs.AnalysisTableFieldOptions? FieldOptions;
        public readonly Outputs.AnalysisTableFieldWells? FieldWells;
        public readonly Outputs.AnalysisTablePaginatedReportOptions? PaginatedReportOptions;
        public readonly Outputs.AnalysisTableSortConfiguration? SortConfiguration;
        public readonly ImmutableArray<Outputs.AnalysisTableInlineVisualization> TableInlineVisualizations;
        public readonly Outputs.AnalysisTableOptions? TableOptions;
        public readonly Outputs.AnalysisTotalOptions? TotalOptions;

        [OutputConstructor]
        private AnalysisTableConfiguration(
            Outputs.AnalysisTableFieldOptions? fieldOptions,

            Outputs.AnalysisTableFieldWells? fieldWells,

            Outputs.AnalysisTablePaginatedReportOptions? paginatedReportOptions,

            Outputs.AnalysisTableSortConfiguration? sortConfiguration,

            ImmutableArray<Outputs.AnalysisTableInlineVisualization> tableInlineVisualizations,

            Outputs.AnalysisTableOptions? tableOptions,

            Outputs.AnalysisTotalOptions? totalOptions)
        {
            FieldOptions = fieldOptions;
            FieldWells = fieldWells;
            PaginatedReportOptions = paginatedReportOptions;
            SortConfiguration = sortConfiguration;
            TableInlineVisualizations = tableInlineVisualizations;
            TableOptions = tableOptions;
            TotalOptions = totalOptions;
        }
    }
}
