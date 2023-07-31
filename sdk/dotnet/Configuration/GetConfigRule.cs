// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration
{
    public static class GetConfigRule
    {
        /// <summary>
        /// Schema for AWS Config ConfigRule
        /// </summary>
        public static Task<GetConfigRuleResult> InvokeAsync(GetConfigRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConfigRuleResult>("aws-native:configuration:getConfigRule", args ?? new GetConfigRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Schema for AWS Config ConfigRule
        /// </summary>
        public static Output<GetConfigRuleResult> Invoke(GetConfigRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConfigRuleResult>("aws-native:configuration:getConfigRule", args ?? new GetConfigRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name for the AWS Config rule
        /// </summary>
        [Input("configRuleName", required: true)]
        public string ConfigRuleName { get; set; } = null!;

        public GetConfigRuleArgs()
        {
        }
        public static new GetConfigRuleArgs Empty => new GetConfigRuleArgs();
    }

    public sealed class GetConfigRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name for the AWS Config rule
        /// </summary>
        [Input("configRuleName", required: true)]
        public Input<string> ConfigRuleName { get; set; } = null!;

        public GetConfigRuleInvokeArgs()
        {
        }
        public static new GetConfigRuleInvokeArgs Empty => new GetConfigRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetConfigRuleResult
    {
        /// <summary>
        /// ARN generated for the AWS Config rule 
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Compliance details of the Config rule
        /// </summary>
        public readonly Outputs.ComplianceProperties? Compliance;
        /// <summary>
        /// ID of the config rule
        /// </summary>
        public readonly string? ConfigRuleId;
        /// <summary>
        /// Description provided for the AWS Config rule
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// List of EvaluationModeConfiguration objects
        /// </summary>
        public readonly ImmutableArray<Outputs.ConfigRuleEvaluationModeConfiguration> EvaluationModes;
        /// <summary>
        /// JSON string passed the Lambda function
        /// </summary>
        public readonly string? InputParameters;
        /// <summary>
        /// Maximum frequency at which the rule has to be evaluated
        /// </summary>
        public readonly string? MaximumExecutionFrequency;
        /// <summary>
        /// Scope to constrain which resources can trigger the AWS Config rule
        /// </summary>
        public readonly Outputs.ConfigRuleScope? Scope;
        /// <summary>
        /// Source of events for the AWS Config rule
        /// </summary>
        public readonly Outputs.ConfigRuleSource? Source;

        [OutputConstructor]
        private GetConfigRuleResult(
            string? arn,

            Outputs.ComplianceProperties? compliance,

            string? configRuleId,

            string? description,

            ImmutableArray<Outputs.ConfigRuleEvaluationModeConfiguration> evaluationModes,

            string? inputParameters,

            string? maximumExecutionFrequency,

            Outputs.ConfigRuleScope? scope,

            Outputs.ConfigRuleSource? source)
        {
            Arn = arn;
            Compliance = compliance;
            ConfigRuleId = configRuleId;
            Description = description;
            EvaluationModes = evaluationModes;
            InputParameters = inputParameters;
            MaximumExecutionFrequency = maximumExecutionFrequency;
            Scope = scope;
            Source = source;
        }
    }
}
