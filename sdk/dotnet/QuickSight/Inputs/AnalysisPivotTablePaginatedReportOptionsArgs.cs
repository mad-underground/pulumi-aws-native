// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisPivotTablePaginatedReportOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("overflowColumnHeaderVisibility")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisVisibility>? OverflowColumnHeaderVisibility { get; set; }

        [Input("verticalOverflowVisibility")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisVisibility>? VerticalOverflowVisibility { get; set; }

        public AnalysisPivotTablePaginatedReportOptionsArgs()
        {
        }
        public static new AnalysisPivotTablePaginatedReportOptionsArgs Empty => new AnalysisPivotTablePaginatedReportOptionsArgs();
    }
}