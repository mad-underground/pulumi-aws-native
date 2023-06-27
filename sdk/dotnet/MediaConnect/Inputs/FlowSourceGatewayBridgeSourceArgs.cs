// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Inputs
{

    /// <summary>
    /// The source configuration for cloud flows receiving a stream from a bridge.
    /// </summary>
    public sealed class FlowSourceGatewayBridgeSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the bridge feeding this flow.
        /// </summary>
        [Input("bridgeArn", required: true)]
        public Input<string> BridgeArn { get; set; } = null!;

        /// <summary>
        /// The name of the VPC interface attachment to use for this bridge source.
        /// </summary>
        [Input("vpcInterfaceAttachment")]
        public Input<Inputs.FlowSourceVpcInterfaceAttachmentArgs>? VpcInterfaceAttachment { get; set; }

        public FlowSourceGatewayBridgeSourceArgs()
        {
        }
        public static new FlowSourceGatewayBridgeSourceArgs Empty => new FlowSourceGatewayBridgeSourceArgs();
    }
}
