// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTTwinMaker
{
    /// <summary>
    /// Resource schema for AWS::IoTTwinMaker::SyncJob
    /// </summary>
    [AwsNativeResourceType("aws-native:iottwinmaker:SyncJob")]
    public partial class SyncJob : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the SyncJob.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The date and time when the sync job was created.
        /// </summary>
        [Output("creationDateTime")]
        public Output<string> CreationDateTime { get; private set; } = null!;

        /// <summary>
        /// The state of SyncJob.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The IAM Role that execute SyncJob.
        /// </summary>
        [Output("syncRole")]
        public Output<string> SyncRole { get; private set; } = null!;

        /// <summary>
        /// The source of the SyncJob.
        /// </summary>
        [Output("syncSource")]
        public Output<string> SyncSource { get; private set; } = null!;

        /// <summary>
        /// A key-value pair to associate with a resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The date and time when the sync job was updated.
        /// </summary>
        [Output("updateDateTime")]
        public Output<string> UpdateDateTime { get; private set; } = null!;

        /// <summary>
        /// The ID of the workspace.
        /// </summary>
        [Output("workspaceId")]
        public Output<string> WorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a SyncJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SyncJob(string name, SyncJobArgs args, CustomResourceOptions? options = null)
            : base("aws-native:iottwinmaker:SyncJob", name, args ?? new SyncJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SyncJob(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:iottwinmaker:SyncJob", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "syncRole",
                    "syncSource",
                    "tags.*",
                    "workspaceId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SyncJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SyncJob Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SyncJob(name, id, options);
        }
    }

    public sealed class SyncJobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IAM Role that execute SyncJob.
        /// </summary>
        [Input("syncRole", required: true)]
        public Input<string> SyncRole { get; set; } = null!;

        /// <summary>
        /// The source of the SyncJob.
        /// </summary>
        [Input("syncSource", required: true)]
        public Input<string> SyncSource { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A key-value pair to associate with a resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the workspace.
        /// </summary>
        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public SyncJobArgs()
        {
        }
        public static new SyncJobArgs Empty => new SyncJobArgs();
    }
}
