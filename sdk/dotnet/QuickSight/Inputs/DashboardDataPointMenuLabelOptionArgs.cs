// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardDataPointMenuLabelOptionArgs : global::Pulumi.ResourceArgs
    {
        [Input("availabilityStatus")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardBehavior>? AvailabilityStatus { get; set; }

        public DashboardDataPointMenuLabelOptionArgs()
        {
        }
        public static new DashboardDataPointMenuLabelOptionArgs Empty => new DashboardDataPointMenuLabelOptionArgs();
    }
}
