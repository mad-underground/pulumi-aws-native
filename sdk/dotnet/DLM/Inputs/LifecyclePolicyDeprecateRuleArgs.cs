// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DLM.Inputs
{

    public sealed class LifecyclePolicyDeprecateRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("count")]
        public Input<int>? Count { get; set; }

        [Input("interval")]
        public Input<int>? Interval { get; set; }

        [Input("intervalUnit")]
        public Input<string>? IntervalUnit { get; set; }

        public LifecyclePolicyDeprecateRuleArgs()
        {
        }
        public static new LifecyclePolicyDeprecateRuleArgs Empty => new LifecyclePolicyDeprecateRuleArgs();
    }
}
