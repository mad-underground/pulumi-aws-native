// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ResilienceHub
{
    public static class GetResiliencyPolicy
    {
        /// <summary>
        /// Resource Type Definition for Resiliency Policy.
        /// </summary>
        public static Task<GetResiliencyPolicyResult> InvokeAsync(GetResiliencyPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResiliencyPolicyResult>("aws-native:resiliencehub:getResiliencyPolicy", args ?? new GetResiliencyPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type Definition for Resiliency Policy.
        /// </summary>
        public static Output<GetResiliencyPolicyResult> Invoke(GetResiliencyPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResiliencyPolicyResult>("aws-native:resiliencehub:getResiliencyPolicy", args ?? new GetResiliencyPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResiliencyPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the Resiliency Policy.
        /// </summary>
        [Input("policyArn", required: true)]
        public string PolicyArn { get; set; } = null!;

        public GetResiliencyPolicyArgs()
        {
        }
        public static new GetResiliencyPolicyArgs Empty => new GetResiliencyPolicyArgs();
    }

    public sealed class GetResiliencyPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the Resiliency Policy.
        /// </summary>
        [Input("policyArn", required: true)]
        public Input<string> PolicyArn { get; set; } = null!;

        public GetResiliencyPolicyInvokeArgs()
        {
        }
        public static new GetResiliencyPolicyInvokeArgs Empty => new GetResiliencyPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetResiliencyPolicyResult
    {
        /// <summary>
        /// Data Location Constraint of the Policy.
        /// </summary>
        public readonly Pulumi.AwsNative.ResilienceHub.ResiliencyPolicyDataLocationConstraint? DataLocationConstraint;
        public readonly Outputs.ResiliencyPolicyPolicyMap? Policy;
        /// <summary>
        /// Amazon Resource Name (ARN) of the Resiliency Policy.
        /// </summary>
        public readonly string? PolicyArn;
        /// <summary>
        /// Description of Resiliency Policy.
        /// </summary>
        public readonly string? PolicyDescription;
        /// <summary>
        /// Name of Resiliency Policy.
        /// </summary>
        public readonly string? PolicyName;
        public readonly Outputs.ResiliencyPolicyTagMap? Tags;
        /// <summary>
        /// Resiliency Policy Tier.
        /// </summary>
        public readonly Pulumi.AwsNative.ResilienceHub.ResiliencyPolicyTier? Tier;

        [OutputConstructor]
        private GetResiliencyPolicyResult(
            Pulumi.AwsNative.ResilienceHub.ResiliencyPolicyDataLocationConstraint? dataLocationConstraint,

            Outputs.ResiliencyPolicyPolicyMap? policy,

            string? policyArn,

            string? policyDescription,

            string? policyName,

            Outputs.ResiliencyPolicyTagMap? tags,

            Pulumi.AwsNative.ResilienceHub.ResiliencyPolicyTier? tier)
        {
            DataLocationConstraint = dataLocationConstraint;
            Policy = policy;
            PolicyArn = policyArn;
            PolicyDescription = policyDescription;
            PolicyName = policyName;
            Tags = tags;
            Tier = tier;
        }
    }
}
