// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudTrail
{
    public static class GetChannel
    {
        /// <summary>
        /// A channel receives events from a specific source (such as an on-premises storage solution or application, or a partner event data source), and delivers the events to one or more event data stores. You use channels to ingest events into CloudTrail from sources outside AWS.
        /// </summary>
        public static Task<GetChannelResult> InvokeAsync(GetChannelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetChannelResult>("aws-native:cloudtrail:getChannel", args ?? new GetChannelArgs(), options.WithDefaults());

        /// <summary>
        /// A channel receives events from a specific source (such as an on-premises storage solution or application, or a partner event data source), and delivers the events to one or more event data stores. You use channels to ingest events into CloudTrail from sources outside AWS.
        /// </summary>
        public static Output<GetChannelResult> Invoke(GetChannelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetChannelResult>("aws-native:cloudtrail:getChannel", args ?? new GetChannelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetChannelArgs : global::Pulumi.InvokeArgs
    {
        [Input("channelArn", required: true)]
        public string ChannelArn { get; set; } = null!;

        public GetChannelArgs()
        {
        }
        public static new GetChannelArgs Empty => new GetChannelArgs();
    }

    public sealed class GetChannelInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("channelArn", required: true)]
        public Input<string> ChannelArn { get; set; } = null!;

        public GetChannelInvokeArgs()
        {
        }
        public static new GetChannelInvokeArgs Empty => new GetChannelInvokeArgs();
    }


    [OutputType]
    public sealed class GetChannelResult
    {
        public readonly string? ChannelArn;
        /// <summary>
        /// One or more resources to which events arriving through a channel are logged and stored.
        /// </summary>
        public readonly ImmutableArray<Outputs.ChannelDestination> Destinations;
        public readonly string? Name;

        [OutputConstructor]
        private GetChannelResult(
            string? channelArn,

            ImmutableArray<Outputs.ChannelDestination> destinations,

            string? name)
        {
            ChannelArn = channelArn;
            Destinations = destinations;
            Name = name;
        }
    }
}
