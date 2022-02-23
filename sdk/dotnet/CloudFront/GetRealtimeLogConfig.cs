// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront
{
    public static class GetRealtimeLogConfig
    {
        /// <summary>
        /// Resource Type definition for AWS::CloudFront::RealtimeLogConfig
        /// </summary>
        public static Task<GetRealtimeLogConfigResult> InvokeAsync(GetRealtimeLogConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRealtimeLogConfigResult>("aws-native:cloudfront:getRealtimeLogConfig", args ?? new GetRealtimeLogConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::CloudFront::RealtimeLogConfig
        /// </summary>
        public static Output<GetRealtimeLogConfigResult> Invoke(GetRealtimeLogConfigInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRealtimeLogConfigResult>("aws-native:cloudfront:getRealtimeLogConfig", args ?? new GetRealtimeLogConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRealtimeLogConfigArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetRealtimeLogConfigArgs()
        {
        }
    }

    public sealed class GetRealtimeLogConfigInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetRealtimeLogConfigInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRealtimeLogConfigResult
    {
        public readonly string? Arn;
        public readonly ImmutableArray<Outputs.RealtimeLogConfigEndPoint> EndPoints;
        public readonly ImmutableArray<string> Fields;
        public readonly double? SamplingRate;

        [OutputConstructor]
        private GetRealtimeLogConfigResult(
            string? arn,

            ImmutableArray<Outputs.RealtimeLogConfigEndPoint> endPoints,

            ImmutableArray<string> fields,

            double? samplingRate)
        {
            Arn = arn;
            EndPoints = endPoints;
            Fields = fields;
            SamplingRate = samplingRate;
        }
    }
}