// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kinesis
{
    public static class GetStreamConsumer
    {
        /// <summary>
        /// Resource Type definition for AWS::Kinesis::StreamConsumer
        /// </summary>
        public static Task<GetStreamConsumerResult> InvokeAsync(GetStreamConsumerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStreamConsumerResult>("aws-native:kinesis:getStreamConsumer", args ?? new GetStreamConsumerArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Kinesis::StreamConsumer
        /// </summary>
        public static Output<GetStreamConsumerResult> Invoke(GetStreamConsumerInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetStreamConsumerResult>("aws-native:kinesis:getStreamConsumer", args ?? new GetStreamConsumerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStreamConsumerArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetStreamConsumerArgs()
        {
        }
    }

    public sealed class GetStreamConsumerInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetStreamConsumerInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStreamConsumerResult
    {
        public readonly string? ConsumerARN;
        public readonly string? ConsumerCreationTimestamp;
        public readonly string? ConsumerStatus;
        public readonly string? Id;

        [OutputConstructor]
        private GetStreamConsumerResult(
            string? consumerARN,

            string? consumerCreationTimestamp,

            string? consumerStatus,

            string? id)
        {
            ConsumerARN = consumerARN;
            ConsumerCreationTimestamp = consumerCreationTimestamp;
            ConsumerStatus = consumerStatus;
            Id = id;
        }
    }
}