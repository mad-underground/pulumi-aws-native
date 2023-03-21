// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplatePivotTableConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldOptions")]
        public Input<Inputs.TemplatePivotTableFieldOptionsArgs>? FieldOptions { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.TemplatePivotTableFieldWellsArgs>? FieldWells { get; set; }

        [Input("paginatedReportOptions")]
        public Input<Inputs.TemplatePivotTablePaginatedReportOptionsArgs>? PaginatedReportOptions { get; set; }

        [Input("sortConfiguration")]
        public Input<Inputs.TemplatePivotTableSortConfigurationArgs>? SortConfiguration { get; set; }

        [Input("tableOptions")]
        public Input<Inputs.TemplatePivotTableOptionsArgs>? TableOptions { get; set; }

        [Input("totalOptions")]
        public Input<Inputs.TemplatePivotTableTotalOptionsArgs>? TotalOptions { get; set; }

        public TemplatePivotTableConfigurationArgs()
        {
        }
        public static new TemplatePivotTableConfigurationArgs Empty => new TemplatePivotTableConfigurationArgs();
    }
}
