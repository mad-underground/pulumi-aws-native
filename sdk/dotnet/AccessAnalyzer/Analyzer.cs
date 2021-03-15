// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AccessAnalyzer
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html
    /// </summary>
    [AwsNativeResourceType("aws-native:AccessAnalyzer:Analyzer")]
    public partial class Analyzer : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-analyzername
        /// </summary>
        [Output("analyzerName")]
        public Output<string?> AnalyzerName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-archiverules
        /// </summary>
        [Output("archiveRules")]
        public Output<ImmutableArray<Outputs.AnalyzerArchiveRule>> ArchiveRules { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Analyzer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Analyzer(string name, AnalyzerArgs args, CustomResourceOptions? options = null)
            : base("aws-native:AccessAnalyzer:Analyzer", name, args ?? new AnalyzerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Analyzer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:AccessAnalyzer:Analyzer", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Analyzer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Analyzer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Analyzer(name, id, options);
        }
    }

    public sealed class AnalyzerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-analyzername
        /// </summary>
        [Input("analyzerName")]
        public Input<string>? AnalyzerName { get; set; }

        [Input("archiveRules")]
        private InputList<Inputs.AnalyzerArchiveRuleArgs>? _archiveRules;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-archiverules
        /// </summary>
        public InputList<Inputs.AnalyzerArchiveRuleArgs> ArchiveRules
        {
            get => _archiveRules ?? (_archiveRules = new InputList<Inputs.AnalyzerArchiveRuleArgs>());
            set => _archiveRules = value;
        }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accessanalyzer-analyzer.html#cfn-accessanalyzer-analyzer-type
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AnalyzerArgs()
        {
        }
    }
}
