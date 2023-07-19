// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Inputs
{

    /// <summary>
    /// The outbound caller ID name, number, and outbound whisper flow.
    /// </summary>
    public sealed class QueueOutboundCallerConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("outboundCallerIdName")]
        public Input<string>? OutboundCallerIdName { get; set; }

        [Input("outboundCallerIdNumberArn")]
        public Input<string>? OutboundCallerIdNumberArn { get; set; }

        [Input("outboundFlowArn")]
        public Input<string>? OutboundFlowArn { get; set; }

        public QueueOutboundCallerConfigArgs()
        {
        }
        public static new QueueOutboundCallerConfigArgs Empty => new QueueOutboundCallerConfigArgs();
    }
}
