// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint
{
    public static class GetSegment
    {
        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::Segment
        /// </summary>
        public static Task<GetSegmentResult> InvokeAsync(GetSegmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSegmentResult>("aws-native:pinpoint:getSegment", args ?? new GetSegmentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::Segment
        /// </summary>
        public static Output<GetSegmentResult> Invoke(GetSegmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSegmentResult>("aws-native:pinpoint:getSegment", args ?? new GetSegmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSegmentArgs : global::Pulumi.InvokeArgs
    {
        [Input("segmentId", required: true)]
        public string SegmentId { get; set; } = null!;

        public GetSegmentArgs()
        {
        }
        public static new GetSegmentArgs Empty => new GetSegmentArgs();
    }

    public sealed class GetSegmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("segmentId", required: true)]
        public Input<string> SegmentId { get; set; } = null!;

        public GetSegmentInvokeArgs()
        {
        }
        public static new GetSegmentInvokeArgs Empty => new GetSegmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetSegmentResult
    {
        public readonly string? Arn;
        public readonly Outputs.SegmentDimensions? Dimensions;
        public readonly string? Name;
        public readonly Outputs.SegmentGroups? SegmentGroups;
        public readonly string? SegmentId;
        /// <summary>
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Pinpoint::Segment` for more information about the expected schema for this property.
        /// </summary>
        public readonly object? Tags;

        [OutputConstructor]
        private GetSegmentResult(
            string? arn,

            Outputs.SegmentDimensions? dimensions,

            string? name,

            Outputs.SegmentGroups? segmentGroups,

            string? segmentId,

            object? tags)
        {
            Arn = arn;
            Dimensions = dimensions;
            Name = name;
            SegmentGroups = segmentGroups;
            SegmentId = segmentId;
            Tags = tags;
        }
    }
}
