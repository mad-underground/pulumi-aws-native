// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardFreeFormLayoutElementBorderStyleArgs : global::Pulumi.ResourceArgs
    {
        [Input("color")]
        public Input<string>? Color { get; set; }

        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? Visibility { get; set; }

        public DashboardFreeFormLayoutElementBorderStyleArgs()
        {
        }
        public static new DashboardFreeFormLayoutElementBorderStyleArgs Empty => new DashboardFreeFormLayoutElementBorderStyleArgs();
    }
}
