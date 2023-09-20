// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor
{
    public static class GetLiveSource
    {
        /// <summary>
        /// Definition of AWS::MediaTailor::LiveSource Resource Type
        /// </summary>
        public static Task<GetLiveSourceResult> InvokeAsync(GetLiveSourceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLiveSourceResult>("aws-native:mediatailor:getLiveSource", args ?? new GetLiveSourceArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::MediaTailor::LiveSource Resource Type
        /// </summary>
        public static Output<GetLiveSourceResult> Invoke(GetLiveSourceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLiveSourceResult>("aws-native:mediatailor:getLiveSource", args ?? new GetLiveSourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLiveSourceArgs : global::Pulumi.InvokeArgs
    {
        [Input("liveSourceName", required: true)]
        public string LiveSourceName { get; set; } = null!;

        [Input("sourceLocationName", required: true)]
        public string SourceLocationName { get; set; } = null!;

        public GetLiveSourceArgs()
        {
        }
        public static new GetLiveSourceArgs Empty => new GetLiveSourceArgs();
    }

    public sealed class GetLiveSourceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("liveSourceName", required: true)]
        public Input<string> LiveSourceName { get; set; } = null!;

        [Input("sourceLocationName", required: true)]
        public Input<string> SourceLocationName { get; set; } = null!;

        public GetLiveSourceInvokeArgs()
        {
        }
        public static new GetLiveSourceInvokeArgs Empty => new GetLiveSourceInvokeArgs();
    }


    [OutputType]
    public sealed class GetLiveSourceResult
    {
        /// <summary>
        /// &lt;p&gt;The ARN of the live source.&lt;/p&gt;
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// &lt;p&gt;A list of HTTP package configuration parameters for this live source.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.LiveSourceHttpPackageConfiguration> HttpPackageConfigurations;
        /// <summary>
        /// The tags to assign to the live source.
        /// </summary>
        public readonly ImmutableArray<Outputs.LiveSourceTag> Tags;

        [OutputConstructor]
        private GetLiveSourceResult(
            string? arn,

            ImmutableArray<Outputs.LiveSourceHttpPackageConfiguration> httpPackageConfigurations,

            ImmutableArray<Outputs.LiveSourceTag> tags)
        {
            Arn = arn;
            HttpPackageConfigurations = httpPackageConfigurations;
            Tags = tags;
        }
    }
}