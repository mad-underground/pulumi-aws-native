// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardArcAxisDisplayRangeArgs : global::Pulumi.ResourceArgs
    {
        [Input("max")]
        public Input<double>? Max { get; set; }

        [Input("min")]
        public Input<double>? Min { get; set; }

        public DashboardArcAxisDisplayRangeArgs()
        {
        }
        public static new DashboardArcAxisDisplayRangeArgs Empty => new DashboardArcAxisDisplayRangeArgs();
    }
}
