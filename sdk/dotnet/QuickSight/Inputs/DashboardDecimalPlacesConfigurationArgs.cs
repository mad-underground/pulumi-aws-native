// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardDecimalPlacesConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("decimalPlaces", required: true)]
        public Input<double> DecimalPlaces { get; set; } = null!;

        public DashboardDecimalPlacesConfigurationArgs()
        {
        }
        public static new DashboardDecimalPlacesConfigurationArgs Empty => new DashboardDecimalPlacesConfigurationArgs();
    }
}
