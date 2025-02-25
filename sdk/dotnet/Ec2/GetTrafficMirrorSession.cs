// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetTrafficMirrorSession
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::TrafficMirrorSession
        /// </summary>
        public static Task<GetTrafficMirrorSessionResult> InvokeAsync(GetTrafficMirrorSessionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTrafficMirrorSessionResult>("aws-native:ec2:getTrafficMirrorSession", args ?? new GetTrafficMirrorSessionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::TrafficMirrorSession
        /// </summary>
        public static Output<GetTrafficMirrorSessionResult> Invoke(GetTrafficMirrorSessionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTrafficMirrorSessionResult>("aws-native:ec2:getTrafficMirrorSession", args ?? new GetTrafficMirrorSessionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTrafficMirrorSessionArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetTrafficMirrorSessionArgs()
        {
        }
        public static new GetTrafficMirrorSessionArgs Empty => new GetTrafficMirrorSessionArgs();
    }

    public sealed class GetTrafficMirrorSessionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetTrafficMirrorSessionInvokeArgs()
        {
        }
        public static new GetTrafficMirrorSessionInvokeArgs Empty => new GetTrafficMirrorSessionInvokeArgs();
    }


    [OutputType]
    public sealed class GetTrafficMirrorSessionResult
    {
        public readonly string? Description;
        public readonly string? Id;
        public readonly int? PacketLength;
        public readonly int? SessionNumber;
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        public readonly string? TrafficMirrorFilterId;
        public readonly string? TrafficMirrorTargetId;
        public readonly int? VirtualNetworkId;

        [OutputConstructor]
        private GetTrafficMirrorSessionResult(
            string? description,

            string? id,

            int? packetLength,

            int? sessionNumber,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? trafficMirrorFilterId,

            string? trafficMirrorTargetId,

            int? virtualNetworkId)
        {
            Description = description;
            Id = id;
            PacketLength = packetLength;
            SessionNumber = sessionNumber;
            Tags = tags;
            TrafficMirrorFilterId = trafficMirrorFilterId;
            TrafficMirrorTargetId = trafficMirrorTargetId;
            VirtualNetworkId = virtualNetworkId;
        }
    }
}
