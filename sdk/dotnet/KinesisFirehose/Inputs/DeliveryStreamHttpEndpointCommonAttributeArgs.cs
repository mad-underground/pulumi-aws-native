// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.KinesisFirehose.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointcommonattribute.html
    /// </summary>
    public sealed class DeliveryStreamHttpEndpointCommonAttributeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointcommonattribute.html#cfn-kinesisfirehose-deliverystream-httpendpointcommonattribute-attributename
        /// </summary>
        [Input("AttributeName", required: true)]
        public Input<string> AttributeName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointcommonattribute.html#cfn-kinesisfirehose-deliverystream-httpendpointcommonattribute-attributevalue
        /// </summary>
        [Input("AttributeValue", required: true)]
        public Input<string> AttributeValue { get; set; } = null!;

        public DeliveryStreamHttpEndpointCommonAttributeArgs()
        {
        }
    }
}