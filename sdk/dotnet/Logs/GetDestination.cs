// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Logs
{
    public static class GetDestination
    {
        /// <summary>
        /// The AWS::Logs::Destination resource specifies a CloudWatch Logs destination. A destination encapsulates a physical resource (such as an Amazon Kinesis data stream) and enables you to subscribe that resource to a stream of log events.
        /// </summary>
        public static Task<GetDestinationResult> InvokeAsync(GetDestinationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDestinationResult>("aws-native:logs:getDestination", args ?? new GetDestinationArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::Logs::Destination resource specifies a CloudWatch Logs destination. A destination encapsulates a physical resource (such as an Amazon Kinesis data stream) and enables you to subscribe that resource to a stream of log events.
        /// </summary>
        public static Output<GetDestinationResult> Invoke(GetDestinationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDestinationResult>("aws-native:logs:getDestination", args ?? new GetDestinationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDestinationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the destination resource
        /// </summary>
        [Input("destinationName", required: true)]
        public string DestinationName { get; set; } = null!;

        public GetDestinationArgs()
        {
        }
        public static new GetDestinationArgs Empty => new GetDestinationArgs();
    }

    public sealed class GetDestinationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the destination resource
        /// </summary>
        [Input("destinationName", required: true)]
        public Input<string> DestinationName { get; set; } = null!;

        public GetDestinationInvokeArgs()
        {
        }
        public static new GetDestinationInvokeArgs Empty => new GetDestinationInvokeArgs();
    }


    [OutputType]
    public sealed class GetDestinationResult
    {
        public readonly string? Arn;
        /// <summary>
        /// An IAM policy document that governs which AWS accounts can create subscription filters against this destination.
        /// </summary>
        public readonly string? DestinationPolicy;
        /// <summary>
        /// The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource
        /// </summary>
        public readonly string? RoleArn;
        /// <summary>
        /// The ARN of the physical target where the log events are delivered (for example, a Kinesis stream)
        /// </summary>
        public readonly string? TargetArn;

        [OutputConstructor]
        private GetDestinationResult(
            string? arn,

            string? destinationPolicy,

            string? roleArn,

            string? targetArn)
        {
            Arn = arn;
            DestinationPolicy = destinationPolicy;
            RoleArn = roleArn;
            TargetArn = targetArn;
        }
    }
}
