// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPivotTablePaginatedReportOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("overflowColumnHeaderVisibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? OverflowColumnHeaderVisibility { get; set; }

        [Input("verticalOverflowVisibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? VerticalOverflowVisibility { get; set; }

        public DashboardPivotTablePaginatedReportOptionsArgs()
        {
        }
        public static new DashboardPivotTablePaginatedReportOptionsArgs Empty => new DashboardPivotTablePaginatedReportOptionsArgs();
    }
}
