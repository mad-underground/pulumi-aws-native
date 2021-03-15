// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertyvalue.html
    /// </summary>
    public sealed class TopicRuleAssetPropertyValueArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertyvalue.html#cfn-iot-topicrule-assetpropertyvalue-quality
        /// </summary>
        [Input("quality")]
        public Input<string>? Quality { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertyvalue.html#cfn-iot-topicrule-assetpropertyvalue-timestamp
        /// </summary>
        [Input("timestamp", required: true)]
        public Input<Inputs.TopicRuleAssetPropertyTimestampArgs> Timestamp { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertyvalue.html#cfn-iot-topicrule-assetpropertyvalue-value
        /// </summary>
        [Input("value", required: true)]
        public Input<Inputs.TopicRuleAssetPropertyVariantArgs> Value { get; set; } = null!;

        public TopicRuleAssetPropertyValueArgs()
        {
        }
    }
}
