// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub.Inputs
{

    public sealed class AutomationRuleDateRangeArgs : global::Pulumi.ResourceArgs
    {
        [Input("unit", required: true)]
        public Input<Pulumi.AwsNative.SecurityHub.AutomationRuleDateRangeUnit> Unit { get; set; } = null!;

        [Input("value", required: true)]
        public Input<double> Value { get; set; } = null!;

        public AutomationRuleDateRangeArgs()
        {
        }
        public static new AutomationRuleDateRangeArgs Empty => new AutomationRuleDateRangeArgs();
    }
}
