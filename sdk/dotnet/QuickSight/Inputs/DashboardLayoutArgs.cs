// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardLayoutArgs : global::Pulumi.ResourceArgs
    {
        [Input("configuration", required: true)]
        public Input<Inputs.DashboardLayoutConfigurationArgs> Configuration { get; set; } = null!;

        public DashboardLayoutArgs()
        {
        }
        public static new DashboardLayoutArgs Empty => new DashboardLayoutArgs();
    }
}
