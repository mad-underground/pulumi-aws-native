// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Inputs
{

    public sealed class DeliveryStreamHttpEndpointCommonAttributeArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributeName", required: true)]
        public Input<string> AttributeName { get; set; } = null!;

        [Input("attributeValue", required: true)]
        public Input<string> AttributeValue { get; set; } = null!;

        public DeliveryStreamHttpEndpointCommonAttributeArgs()
        {
        }
        public static new DeliveryStreamHttpEndpointCommonAttributeArgs Empty => new DeliveryStreamHttpEndpointCommonAttributeArgs();
    }
}
