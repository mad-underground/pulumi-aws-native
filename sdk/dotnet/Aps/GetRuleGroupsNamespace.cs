// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Aps
{
    public static class GetRuleGroupsNamespace
    {
        /// <summary>
        /// RuleGroupsNamespace schema for cloudformation.
        /// </summary>
        public static Task<GetRuleGroupsNamespaceResult> InvokeAsync(GetRuleGroupsNamespaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRuleGroupsNamespaceResult>("aws-native:aps:getRuleGroupsNamespace", args ?? new GetRuleGroupsNamespaceArgs(), options.WithDefaults());

        /// <summary>
        /// RuleGroupsNamespace schema for cloudformation.
        /// </summary>
        public static Output<GetRuleGroupsNamespaceResult> Invoke(GetRuleGroupsNamespaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRuleGroupsNamespaceResult>("aws-native:aps:getRuleGroupsNamespace", args ?? new GetRuleGroupsNamespaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRuleGroupsNamespaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The RuleGroupsNamespace ARN.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetRuleGroupsNamespaceArgs()
        {
        }
        public static new GetRuleGroupsNamespaceArgs Empty => new GetRuleGroupsNamespaceArgs();
    }

    public sealed class GetRuleGroupsNamespaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The RuleGroupsNamespace ARN.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetRuleGroupsNamespaceInvokeArgs()
        {
        }
        public static new GetRuleGroupsNamespaceInvokeArgs Empty => new GetRuleGroupsNamespaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetRuleGroupsNamespaceResult
    {
        /// <summary>
        /// The RuleGroupsNamespace ARN.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The RuleGroupsNamespace data.
        /// </summary>
        public readonly string? Data;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleGroupsNamespaceTag> Tags;

        [OutputConstructor]
        private GetRuleGroupsNamespaceResult(
            string? arn,

            string? data,

            ImmutableArray<Outputs.RuleGroupsNamespaceTag> tags)
        {
            Arn = arn;
            Data = data;
            Tags = tags;
        }
    }
}
