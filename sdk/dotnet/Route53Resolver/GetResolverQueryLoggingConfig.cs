// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53Resolver
{
    public static class GetResolverQueryLoggingConfig
    {
        /// <summary>
        /// Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfig.
        /// </summary>
        public static Task<GetResolverQueryLoggingConfigResult> InvokeAsync(GetResolverQueryLoggingConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResolverQueryLoggingConfigResult>("aws-native:route53resolver:getResolverQueryLoggingConfig", args ?? new GetResolverQueryLoggingConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfig.
        /// </summary>
        public static Output<GetResolverQueryLoggingConfigResult> Invoke(GetResolverQueryLoggingConfigInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResolverQueryLoggingConfigResult>("aws-native:route53resolver:getResolverQueryLoggingConfig", args ?? new GetResolverQueryLoggingConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResolverQueryLoggingConfigArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ResourceId
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetResolverQueryLoggingConfigArgs()
        {
        }
    }

    public sealed class GetResolverQueryLoggingConfigInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ResourceId
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetResolverQueryLoggingConfigInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResolverQueryLoggingConfigResult
    {
        /// <summary>
        /// Arn
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Count
        /// </summary>
        public readonly int? AssociationCount;
        /// <summary>
        /// Rfc3339TimeString
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// The id of the creator request.
        /// </summary>
        public readonly string? CreatorRequestId;
        /// <summary>
        /// ResourceId
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// AccountId
        /// </summary>
        public readonly string? OwnerId;
        /// <summary>
        /// ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.
        /// </summary>
        public readonly Pulumi.AwsNative.Route53Resolver.ResolverQueryLoggingConfigShareStatus? ShareStatus;
        /// <summary>
        /// ResolverQueryLogConfigStatus, possible values are CREATING, CREATED, DELETED AND FAILED.
        /// </summary>
        public readonly Pulumi.AwsNative.Route53Resolver.ResolverQueryLoggingConfigStatus? Status;

        [OutputConstructor]
        private GetResolverQueryLoggingConfigResult(
            string? arn,

            int? associationCount,

            string? creationTime,

            string? creatorRequestId,

            string? id,

            string? ownerId,

            Pulumi.AwsNative.Route53Resolver.ResolverQueryLoggingConfigShareStatus? shareStatus,

            Pulumi.AwsNative.Route53Resolver.ResolverQueryLoggingConfigStatus? status)
        {
            Arn = arn;
            AssociationCount = associationCount;
            CreationTime = creationTime;
            CreatorRequestId = creatorRequestId;
            Id = id;
            OwnerId = ownerId;
            ShareStatus = shareStatus;
            Status = status;
        }
    }
}
