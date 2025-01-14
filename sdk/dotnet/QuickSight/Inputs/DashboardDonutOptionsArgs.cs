// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardDonutOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("arcOptions")]
        public Input<Inputs.DashboardArcOptionsArgs>? ArcOptions { get; set; }

        [Input("donutCenterOptions")]
        public Input<Inputs.DashboardDonutCenterOptionsArgs>? DonutCenterOptions { get; set; }

        public DashboardDonutOptionsArgs()
        {
        }
        public static new DashboardDonutOptionsArgs Empty => new DashboardDonutOptionsArgs();
    }
}
