// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPivotTableRowsLabelOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("customLabel")]
        public Input<string>? CustomLabel { get; set; }

        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? Visibility { get; set; }

        public DashboardPivotTableRowsLabelOptionsArgs()
        {
        }
        public static new DashboardPivotTableRowsLabelOptionsArgs Empty => new DashboardPivotTableRowsLabelOptionsArgs();
    }
}
