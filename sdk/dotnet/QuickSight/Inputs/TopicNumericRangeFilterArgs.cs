// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TopicNumericRangeFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("aggregation")]
        public Input<Pulumi.AwsNative.QuickSight.TopicNamedFilterAggType>? Aggregation { get; set; }

        [Input("constant")]
        public Input<Inputs.TopicRangeFilterConstantArgs>? Constant { get; set; }

        [Input("inclusive")]
        public Input<bool>? Inclusive { get; set; }

        public TopicNumericRangeFilterArgs()
        {
        }
        public static new TopicNumericRangeFilterArgs Empty => new TopicNumericRangeFilterArgs();
    }
}
