// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Aps
{
    /// <summary>
    /// Resource Type definition for AWS::APS::Workspace
    /// </summary>
    [AwsNativeResourceType("aws-native:aps:Workspace")]
    public partial class Workspace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The AMP Workspace alert manager definition data
        /// </summary>
        [Output("alertManagerDefinition")]
        public Output<string?> AlertManagerDefinition { get; private set; } = null!;

        /// <summary>
        /// AMP Workspace alias.
        /// </summary>
        [Output("alias")]
        public Output<string?> Alias { get; private set; } = null!;

        /// <summary>
        /// Workspace arn.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// KMS Key ARN used to encrypt and decrypt AMP workspace data.
        /// </summary>
        [Output("kmsKeyArn")]
        public Output<string?> KmsKeyArn { get; private set; } = null!;

        [Output("loggingConfiguration")]
        public Output<Outputs.WorkspaceLoggingConfiguration?> LoggingConfiguration { get; private set; } = null!;

        /// <summary>
        /// AMP Workspace prometheus endpoint
        /// </summary>
        [Output("prometheusEndpoint")]
        public Output<string> PrometheusEndpoint { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Required to identify a specific APS Workspace.
        /// </summary>
        [Output("workspaceId")]
        public Output<string> WorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a Workspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workspace(string name, WorkspaceArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:aps:Workspace", name, args ?? new WorkspaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workspace(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:aps:Workspace", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "kmsKeyArn",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Workspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workspace Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Workspace(name, id, options);
        }
    }

    public sealed class WorkspaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AMP Workspace alert manager definition data
        /// </summary>
        [Input("alertManagerDefinition")]
        public Input<string>? AlertManagerDefinition { get; set; }

        /// <summary>
        /// AMP Workspace alias.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// KMS Key ARN used to encrypt and decrypt AMP workspace data.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        [Input("loggingConfiguration")]
        public Input<Inputs.WorkspaceLoggingConfigurationArgs>? LoggingConfiguration { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public WorkspaceArgs()
        {
        }
        public static new WorkspaceArgs Empty => new WorkspaceArgs();
    }
}
