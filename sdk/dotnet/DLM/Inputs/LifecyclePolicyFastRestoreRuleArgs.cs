// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DLM.Inputs
{

    public sealed class LifecyclePolicyFastRestoreRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        [Input("count")]
        public Input<int>? Count { get; set; }

        [Input("interval")]
        public Input<int>? Interval { get; set; }

        [Input("intervalUnit")]
        public Input<string>? IntervalUnit { get; set; }

        public LifecyclePolicyFastRestoreRuleArgs()
        {
        }
        public static new LifecyclePolicyFastRestoreRuleArgs Empty => new LifecyclePolicyFastRestoreRuleArgs();
    }
}
