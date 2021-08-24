// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53RecoveryReadiness.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-resource.html
    /// </summary>
    [OutputType]
    public sealed class ResourceSetResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-resource.html#cfn-route53recoveryreadiness-resourceset-resource-componentid
        /// </summary>
        public readonly string? ComponentId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-resource.html#cfn-route53recoveryreadiness-resourceset-resource-dnstargetresource
        /// </summary>
        public readonly Outputs.ResourceSetDNSTargetResource? DnsTargetResource;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-resource.html#cfn-route53recoveryreadiness-resourceset-resource-readinessscopes
        /// </summary>
        public readonly ImmutableArray<string> ReadinessScopes;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53recoveryreadiness-resourceset-resource.html#cfn-route53recoveryreadiness-resourceset-resource-resourcearn
        /// </summary>
        public readonly string? ResourceArn;

        [OutputConstructor]
        private ResourceSetResource(
            string? componentId,

            Outputs.ResourceSetDNSTargetResource? dnsTargetResource,

            ImmutableArray<string> readinessScopes,

            string? resourceArn)
        {
            ComponentId = componentId;
            DnsTargetResource = dnsTargetResource;
            ReadinessScopes = readinessScopes;
            ResourceArn = resourceArn;
        }
    }
}