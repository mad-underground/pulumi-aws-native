// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardSectionBasedLayoutCanvasSizeOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("paperCanvasSizeOptions")]
        public Input<Inputs.DashboardSectionBasedLayoutPaperCanvasSizeOptionsArgs>? PaperCanvasSizeOptions { get; set; }

        public DashboardSectionBasedLayoutCanvasSizeOptionsArgs()
        {
        }
        public static new DashboardSectionBasedLayoutCanvasSizeOptionsArgs Empty => new DashboardSectionBasedLayoutCanvasSizeOptionsArgs();
    }
}