// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew
{
    public static class GetRuleset
    {
        /// <summary>
        /// Resource schema for AWS::DataBrew::Ruleset.
        /// </summary>
        public static Task<GetRulesetResult> InvokeAsync(GetRulesetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRulesetResult>("aws-native:databrew:getRuleset", args ?? new GetRulesetArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::DataBrew::Ruleset.
        /// </summary>
        public static Output<GetRulesetResult> Invoke(GetRulesetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRulesetResult>("aws-native:databrew:getRuleset", args ?? new GetRulesetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRulesetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Ruleset
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetRulesetArgs()
        {
        }
        public static new GetRulesetArgs Empty => new GetRulesetArgs();
    }

    public sealed class GetRulesetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Ruleset
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetRulesetInvokeArgs()
        {
        }
        public static new GetRulesetInvokeArgs Empty => new GetRulesetInvokeArgs();
    }


    [OutputType]
    public sealed class GetRulesetResult
    {
        /// <summary>
        /// Description of the Ruleset
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// List of the data quality rules in the ruleset
        /// </summary>
        public readonly ImmutableArray<Outputs.RulesetRule> Rules;
        public readonly ImmutableArray<Outputs.RulesetTag> Tags;

        [OutputConstructor]
        private GetRulesetResult(
            string? description,

            ImmutableArray<Outputs.RulesetRule> rules,

            ImmutableArray<Outputs.RulesetTag> tags)
        {
            Description = description;
            Rules = rules;
            Tags = tags;
        }
    }
}
