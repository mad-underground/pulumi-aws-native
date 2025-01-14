// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPanelTitleOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("fontConfiguration")]
        public Input<Inputs.DashboardFontConfigurationArgs>? FontConfiguration { get; set; }

        [Input("horizontalTextAlignment")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardHorizontalTextAlignment>? HorizontalTextAlignment { get; set; }

        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? Visibility { get; set; }

        public DashboardPanelTitleOptionsArgs()
        {
        }
        public static new DashboardPanelTitleOptionsArgs Empty => new DashboardPanelTitleOptionsArgs();
    }
}
