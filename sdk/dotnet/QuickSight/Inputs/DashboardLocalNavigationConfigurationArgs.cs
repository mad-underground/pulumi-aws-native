// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardLocalNavigationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("targetSheetId", required: true)]
        public Input<string> TargetSheetId { get; set; } = null!;

        public DashboardLocalNavigationConfigurationArgs()
        {
        }
        public static new DashboardLocalNavigationConfigurationArgs Empty => new DashboardLocalNavigationConfigurationArgs();
    }
}
