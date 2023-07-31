// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration
{
    /// <summary>
    /// Schema for AWS Config ConfigRule
    /// </summary>
    [AwsNativeResourceType("aws-native:configuration:ConfigRule")]
    public partial class ConfigRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN generated for the AWS Config rule 
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Compliance details of the Config rule
        /// </summary>
        [Output("compliance")]
        public Output<Outputs.ComplianceProperties?> Compliance { get; private set; } = null!;

        /// <summary>
        /// ID of the config rule
        /// </summary>
        [Output("configRuleId")]
        public Output<string> ConfigRuleId { get; private set; } = null!;

        /// <summary>
        /// Name for the AWS Config rule
        /// </summary>
        [Output("configRuleName")]
        public Output<string?> ConfigRuleName { get; private set; } = null!;

        /// <summary>
        /// Description provided for the AWS Config rule
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of EvaluationModeConfiguration objects
        /// </summary>
        [Output("evaluationModes")]
        public Output<ImmutableArray<Outputs.ConfigRuleEvaluationModeConfiguration>> EvaluationModes { get; private set; } = null!;

        /// <summary>
        /// JSON string passed the Lambda function
        /// </summary>
        [Output("inputParameters")]
        public Output<string?> InputParameters { get; private set; } = null!;

        /// <summary>
        /// Maximum frequency at which the rule has to be evaluated
        /// </summary>
        [Output("maximumExecutionFrequency")]
        public Output<string?> MaximumExecutionFrequency { get; private set; } = null!;

        /// <summary>
        /// Scope to constrain which resources can trigger the AWS Config rule
        /// </summary>
        [Output("scope")]
        public Output<Outputs.ConfigRuleScope?> Scope { get; private set; } = null!;

        /// <summary>
        /// Source of events for the AWS Config rule
        /// </summary>
        [Output("source")]
        public Output<Outputs.ConfigRuleSource> Source { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigRule(string name, ConfigRuleArgs args, CustomResourceOptions? options = null)
            : base("aws-native:configuration:ConfigRule", name, args ?? new ConfigRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:configuration:ConfigRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ConfigRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConfigRule(name, id, options);
        }
    }

    public sealed class ConfigRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Compliance details of the Config rule
        /// </summary>
        [Input("compliance")]
        public Input<Inputs.CompliancePropertiesArgs>? Compliance { get; set; }

        /// <summary>
        /// Name for the AWS Config rule
        /// </summary>
        [Input("configRuleName")]
        public Input<string>? ConfigRuleName { get; set; }

        /// <summary>
        /// Description provided for the AWS Config rule
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("evaluationModes")]
        private InputList<Inputs.ConfigRuleEvaluationModeConfigurationArgs>? _evaluationModes;

        /// <summary>
        /// List of EvaluationModeConfiguration objects
        /// </summary>
        public InputList<Inputs.ConfigRuleEvaluationModeConfigurationArgs> EvaluationModes
        {
            get => _evaluationModes ?? (_evaluationModes = new InputList<Inputs.ConfigRuleEvaluationModeConfigurationArgs>());
            set => _evaluationModes = value;
        }

        /// <summary>
        /// JSON string passed the Lambda function
        /// </summary>
        [Input("inputParameters")]
        public Input<string>? InputParameters { get; set; }

        /// <summary>
        /// Maximum frequency at which the rule has to be evaluated
        /// </summary>
        [Input("maximumExecutionFrequency")]
        public Input<string>? MaximumExecutionFrequency { get; set; }

        /// <summary>
        /// Scope to constrain which resources can trigger the AWS Config rule
        /// </summary>
        [Input("scope")]
        public Input<Inputs.ConfigRuleScopeArgs>? Scope { get; set; }

        /// <summary>
        /// Source of events for the AWS Config rule
        /// </summary>
        [Input("source", required: true)]
        public Input<Inputs.ConfigRuleSourceArgs> Source { get; set; } = null!;

        public ConfigRuleArgs()
        {
        }
        public static new ConfigRuleArgs Empty => new ConfigRuleArgs();
    }
}
