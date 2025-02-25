// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront
{
    public static class GetDistribution
    {
        /// <summary>
        /// A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.
        /// </summary>
        public static Task<GetDistributionResult> InvokeAsync(GetDistributionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDistributionResult>("aws-native:cloudfront:getDistribution", args ?? new GetDistributionArgs(), options.WithDefaults());

        /// <summary>
        /// A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.
        /// </summary>
        public static Output<GetDistributionResult> Invoke(GetDistributionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDistributionResult>("aws-native:cloudfront:getDistribution", args ?? new GetDistributionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDistributionArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDistributionArgs()
        {
        }
        public static new GetDistributionArgs Empty => new GetDistributionArgs();
    }

    public sealed class GetDistributionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDistributionInvokeArgs()
        {
        }
        public static new GetDistributionInvokeArgs Empty => new GetDistributionInvokeArgs();
    }


    [OutputType]
    public sealed class GetDistributionResult
    {
        /// <summary>
        /// The distribution's configuration.
        /// </summary>
        public readonly Outputs.DistributionConfig? DistributionConfig;
        public readonly string? DomainName;
        public readonly string? Id;
        /// <summary>
        /// A complex type that contains zero or more ``Tag`` elements.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetDistributionResult(
            Outputs.DistributionConfig? distributionConfig,

            string? domainName,

            string? id,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            DistributionConfig = distributionConfig;
            DomainName = domainName;
            Id = id;
            Tags = tags;
        }
    }
}
