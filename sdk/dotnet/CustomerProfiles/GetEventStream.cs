// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles
{
    public static class GetEventStream
    {
        /// <summary>
        /// An Event Stream resource of Amazon Connect Customer Profiles
        /// </summary>
        public static Task<GetEventStreamResult> InvokeAsync(GetEventStreamArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEventStreamResult>("aws-native:customerprofiles:getEventStream", args ?? new GetEventStreamArgs(), options.WithDefaults());

        /// <summary>
        /// An Event Stream resource of Amazon Connect Customer Profiles
        /// </summary>
        public static Output<GetEventStreamResult> Invoke(GetEventStreamInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEventStreamResult>("aws-native:customerprofiles:getEventStream", args ?? new GetEventStreamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEventStreamArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name of the domain.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        /// <summary>
        /// The name of the event stream.
        /// </summary>
        [Input("eventStreamName", required: true)]
        public string EventStreamName { get; set; } = null!;

        public GetEventStreamArgs()
        {
        }
        public static new GetEventStreamArgs Empty => new GetEventStreamArgs();
    }

    public sealed class GetEventStreamInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name of the domain.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The name of the event stream.
        /// </summary>
        [Input("eventStreamName", required: true)]
        public Input<string> EventStreamName { get; set; } = null!;

        public GetEventStreamInvokeArgs()
        {
        }
        public static new GetEventStreamInvokeArgs Empty => new GetEventStreamInvokeArgs();
    }


    [OutputType]
    public sealed class GetEventStreamResult
    {
        /// <summary>
        /// The timestamp of when the export was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// Details regarding the Kinesis stream.
        /// </summary>
        public readonly Outputs.DestinationDetailsProperties? DestinationDetails;
        /// <summary>
        /// A unique identifier for the event stream.
        /// </summary>
        public readonly string? EventStreamArn;
        /// <summary>
        /// The operational state of destination stream for export.
        /// </summary>
        public readonly Pulumi.AwsNative.CustomerProfiles.EventStreamState? State;
        /// <summary>
        /// The tags used to organize, track, or control access for this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.EventStreamTag> Tags;

        [OutputConstructor]
        private GetEventStreamResult(
            string? createdAt,

            Outputs.DestinationDetailsProperties? destinationDetails,

            string? eventStreamArn,

            Pulumi.AwsNative.CustomerProfiles.EventStreamState? state,

            ImmutableArray<Outputs.EventStreamTag> tags)
        {
            CreatedAt = createdAt;
            DestinationDetails = destinationDetails;
            EventStreamArn = eventStreamArn;
            State = state;
            Tags = tags;
        }
    }
}