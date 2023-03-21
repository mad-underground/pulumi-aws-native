// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPaginationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("pageNumber", required: true)]
        public Input<double> PageNumber { get; set; } = null!;

        [Input("pageSize", required: true)]
        public Input<double> PageSize { get; set; } = null!;

        public DashboardPaginationConfigurationArgs()
        {
        }
        public static new DashboardPaginationConfigurationArgs Empty => new DashboardPaginationConfigurationArgs();
    }
}
