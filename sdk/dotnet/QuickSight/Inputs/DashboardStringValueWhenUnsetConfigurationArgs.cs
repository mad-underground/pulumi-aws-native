// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardStringValueWhenUnsetConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("customValue")]
        public Input<string>? CustomValue { get; set; }

        [Input("valueWhenUnsetOption")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardValueWhenUnsetOption>? ValueWhenUnsetOption { get; set; }

        public DashboardStringValueWhenUnsetConfigurationArgs()
        {
        }
        public static new DashboardStringValueWhenUnsetConfigurationArgs Empty => new DashboardStringValueWhenUnsetConfigurationArgs();
    }
}
