// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardComparisonConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("comparisonFormat")]
        public Input<Inputs.DashboardComparisonFormatConfigurationArgs>? ComparisonFormat { get; set; }

        [Input("comparisonMethod")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardComparisonMethod>? ComparisonMethod { get; set; }

        public DashboardComparisonConfigurationArgs()
        {
        }
        public static new DashboardComparisonConfigurationArgs Empty => new DashboardComparisonConfigurationArgs();
    }
}
