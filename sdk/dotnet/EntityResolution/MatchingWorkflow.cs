// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EntityResolution
{
    /// <summary>
    /// MatchingWorkflow defined in AWS Entity Resolution service
    /// </summary>
    [AwsNativeResourceType("aws-native:entityresolution:MatchingWorkflow")]
    public partial class MatchingWorkflow : global::Pulumi.CustomResource
    {
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The description of the MatchingWorkflow
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("inputSourceConfig")]
        public Output<ImmutableArray<Outputs.MatchingWorkflowInputSource>> InputSourceConfig { get; private set; } = null!;

        [Output("outputSourceConfig")]
        public Output<ImmutableArray<Outputs.MatchingWorkflowOutputSource>> OutputSourceConfig { get; private set; } = null!;

        [Output("resolutionTechniques")]
        public Output<Outputs.MatchingWorkflowResolutionTechniques> ResolutionTechniques { get; private set; } = null!;

        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.MatchingWorkflowTag>> Tags { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        [Output("workflowArn")]
        public Output<string> WorkflowArn { get; private set; } = null!;

        /// <summary>
        /// The name of the MatchingWorkflow
        /// </summary>
        [Output("workflowName")]
        public Output<string> WorkflowName { get; private set; } = null!;


        /// <summary>
        /// Create a MatchingWorkflow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MatchingWorkflow(string name, MatchingWorkflowArgs args, CustomResourceOptions? options = null)
            : base("aws-native:entityresolution:MatchingWorkflow", name, args ?? new MatchingWorkflowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MatchingWorkflow(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:entityresolution:MatchingWorkflow", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "workflowName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MatchingWorkflow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MatchingWorkflow Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MatchingWorkflow(name, id, options);
        }
    }

    public sealed class MatchingWorkflowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the MatchingWorkflow
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("inputSourceConfig", required: true)]
        private InputList<Inputs.MatchingWorkflowInputSourceArgs>? _inputSourceConfig;
        public InputList<Inputs.MatchingWorkflowInputSourceArgs> InputSourceConfig
        {
            get => _inputSourceConfig ?? (_inputSourceConfig = new InputList<Inputs.MatchingWorkflowInputSourceArgs>());
            set => _inputSourceConfig = value;
        }

        [Input("outputSourceConfig", required: true)]
        private InputList<Inputs.MatchingWorkflowOutputSourceArgs>? _outputSourceConfig;
        public InputList<Inputs.MatchingWorkflowOutputSourceArgs> OutputSourceConfig
        {
            get => _outputSourceConfig ?? (_outputSourceConfig = new InputList<Inputs.MatchingWorkflowOutputSourceArgs>());
            set => _outputSourceConfig = value;
        }

        [Input("resolutionTechniques", required: true)]
        public Input<Inputs.MatchingWorkflowResolutionTechniquesArgs> ResolutionTechniques { get; set; } = null!;

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.MatchingWorkflowTagArgs>? _tags;
        public InputList<Inputs.MatchingWorkflowTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.MatchingWorkflowTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the MatchingWorkflow
        /// </summary>
        [Input("workflowName", required: true)]
        public Input<string> WorkflowName { get; set; } = null!;

        public MatchingWorkflowArgs()
        {
        }
        public static new MatchingWorkflowArgs Empty => new MatchingWorkflowArgs();
    }
}