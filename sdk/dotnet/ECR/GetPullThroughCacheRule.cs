// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECR
{
    public static class GetPullThroughCacheRule
    {
        /// <summary>
        /// The AWS::ECR::PullThroughCacheRule resource configures the upstream registry configuration details for an Amazon Elastic Container Registry (Amazon Private ECR) pull-through cache.
        /// </summary>
        public static Task<GetPullThroughCacheRuleResult> InvokeAsync(GetPullThroughCacheRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPullThroughCacheRuleResult>("aws-native:ecr:getPullThroughCacheRule", args ?? new GetPullThroughCacheRuleArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::ECR::PullThroughCacheRule resource configures the upstream registry configuration details for an Amazon Elastic Container Registry (Amazon Private ECR) pull-through cache.
        /// </summary>
        public static Output<GetPullThroughCacheRuleResult> Invoke(GetPullThroughCacheRuleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPullThroughCacheRuleResult>("aws-native:ecr:getPullThroughCacheRule", args ?? new GetPullThroughCacheRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPullThroughCacheRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ECRRepositoryPrefix is a custom alias for upstream registry url.
        /// </summary>
        [Input("ecrRepositoryPrefix", required: true)]
        public string EcrRepositoryPrefix { get; set; } = null!;

        public GetPullThroughCacheRuleArgs()
        {
        }
    }

    public sealed class GetPullThroughCacheRuleInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ECRRepositoryPrefix is a custom alias for upstream registry url.
        /// </summary>
        [Input("ecrRepositoryPrefix", required: true)]
        public Input<string> EcrRepositoryPrefix { get; set; } = null!;

        public GetPullThroughCacheRuleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPullThroughCacheRuleResult
    {
        [OutputConstructor]
        private GetPullThroughCacheRuleResult()
        {
        }
    }
}