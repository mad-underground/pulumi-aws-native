// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeBuild
{
    /// <summary>
    /// Resource Type definition for AWS::CodeBuild::Project
    /// </summary>
    [Obsolete(@"Project is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:codebuild:Project")]
    public partial class Project : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("artifacts")]
        public Output<Outputs.ProjectArtifacts> Artifacts { get; private set; } = null!;

        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("badgeEnabled")]
        public Output<bool?> BadgeEnabled { get; private set; } = null!;

        [Output("buildBatchConfig")]
        public Output<Outputs.ProjectBuildBatchConfig?> BuildBatchConfig { get; private set; } = null!;

        [Output("cache")]
        public Output<Outputs.ProjectCache?> Cache { get; private set; } = null!;

        [Output("concurrentBuildLimit")]
        public Output<int?> ConcurrentBuildLimit { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("encryptionKey")]
        public Output<string?> EncryptionKey { get; private set; } = null!;

        [Output("environment")]
        public Output<Outputs.ProjectEnvironment> Environment { get; private set; } = null!;

        [Output("fileSystemLocations")]
        public Output<ImmutableArray<Outputs.ProjectFileSystemLocation>> FileSystemLocations { get; private set; } = null!;

        [Output("logsConfig")]
        public Output<Outputs.ProjectLogsConfig?> LogsConfig { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("queuedTimeoutInMinutes")]
        public Output<int?> QueuedTimeoutInMinutes { get; private set; } = null!;

        [Output("resourceAccessRole")]
        public Output<string?> ResourceAccessRole { get; private set; } = null!;

        [Output("secondaryArtifacts")]
        public Output<ImmutableArray<Outputs.ProjectArtifacts>> SecondaryArtifacts { get; private set; } = null!;

        [Output("secondarySourceVersions")]
        public Output<ImmutableArray<Outputs.ProjectSourceVersion>> SecondarySourceVersions { get; private set; } = null!;

        [Output("secondarySources")]
        public Output<ImmutableArray<Outputs.ProjectSource>> SecondarySources { get; private set; } = null!;

        [Output("serviceRole")]
        public Output<string> ServiceRole { get; private set; } = null!;

        [Output("source")]
        public Output<Outputs.ProjectSource> Source { get; private set; } = null!;

        [Output("sourceVersion")]
        public Output<string?> SourceVersion { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        [Output("timeoutInMinutes")]
        public Output<int?> TimeoutInMinutes { get; private set; } = null!;

        [Output("triggers")]
        public Output<Outputs.ProjectTriggers?> Triggers { get; private set; } = null!;

        [Output("visibility")]
        public Output<string?> Visibility { get; private set; } = null!;

        [Output("vpcConfig")]
        public Output<Outputs.ProjectVpcConfig?> VpcConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs args, CustomResourceOptions? options = null)
            : base("aws-native:codebuild:Project", name, args ?? new ProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:codebuild:Project", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Project(name, id, options);
        }
    }

    public sealed class ProjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("artifacts", required: true)]
        public Input<Inputs.ProjectArtifactsArgs> Artifacts { get; set; } = null!;

        [Input("badgeEnabled")]
        public Input<bool>? BadgeEnabled { get; set; }

        [Input("buildBatchConfig")]
        public Input<Inputs.ProjectBuildBatchConfigArgs>? BuildBatchConfig { get; set; }

        [Input("cache")]
        public Input<Inputs.ProjectCacheArgs>? Cache { get; set; }

        [Input("concurrentBuildLimit")]
        public Input<int>? ConcurrentBuildLimit { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("encryptionKey")]
        public Input<string>? EncryptionKey { get; set; }

        [Input("environment", required: true)]
        public Input<Inputs.ProjectEnvironmentArgs> Environment { get; set; } = null!;

        [Input("fileSystemLocations")]
        private InputList<Inputs.ProjectFileSystemLocationArgs>? _fileSystemLocations;
        public InputList<Inputs.ProjectFileSystemLocationArgs> FileSystemLocations
        {
            get => _fileSystemLocations ?? (_fileSystemLocations = new InputList<Inputs.ProjectFileSystemLocationArgs>());
            set => _fileSystemLocations = value;
        }

        [Input("logsConfig")]
        public Input<Inputs.ProjectLogsConfigArgs>? LogsConfig { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("queuedTimeoutInMinutes")]
        public Input<int>? QueuedTimeoutInMinutes { get; set; }

        [Input("resourceAccessRole")]
        public Input<string>? ResourceAccessRole { get; set; }

        [Input("secondaryArtifacts")]
        private InputList<Inputs.ProjectArtifactsArgs>? _secondaryArtifacts;
        public InputList<Inputs.ProjectArtifactsArgs> SecondaryArtifacts
        {
            get => _secondaryArtifacts ?? (_secondaryArtifacts = new InputList<Inputs.ProjectArtifactsArgs>());
            set => _secondaryArtifacts = value;
        }

        [Input("secondarySourceVersions")]
        private InputList<Inputs.ProjectSourceVersionArgs>? _secondarySourceVersions;
        public InputList<Inputs.ProjectSourceVersionArgs> SecondarySourceVersions
        {
            get => _secondarySourceVersions ?? (_secondarySourceVersions = new InputList<Inputs.ProjectSourceVersionArgs>());
            set => _secondarySourceVersions = value;
        }

        [Input("secondarySources")]
        private InputList<Inputs.ProjectSourceArgs>? _secondarySources;
        public InputList<Inputs.ProjectSourceArgs> SecondarySources
        {
            get => _secondarySources ?? (_secondarySources = new InputList<Inputs.ProjectSourceArgs>());
            set => _secondarySources = value;
        }

        [Input("serviceRole", required: true)]
        public Input<string> ServiceRole { get; set; } = null!;

        [Input("source", required: true)]
        public Input<Inputs.ProjectSourceArgs> Source { get; set; } = null!;

        [Input("sourceVersion")]
        public Input<string>? SourceVersion { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("timeoutInMinutes")]
        public Input<int>? TimeoutInMinutes { get; set; }

        [Input("triggers")]
        public Input<Inputs.ProjectTriggersArgs>? Triggers { get; set; }

        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        [Input("vpcConfig")]
        public Input<Inputs.ProjectVpcConfigArgs>? VpcConfig { get; set; }

        public ProjectArgs()
        {
        }
        public static new ProjectArgs Empty => new ProjectArgs();
    }
}
