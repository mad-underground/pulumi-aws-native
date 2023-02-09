// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameCast
{
    /// <summary>
    /// Definition of AWS::GameCast::Application Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:gamecast:Application")]
    public partial class Application : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ApplicationSourceUri points to a S3 Uri to replicate game files from customer S3 bucket into Motif internal S3 bucket.
        /// The uri points to a S3 prefix that could contain many objects. Motif will attempt to copy all the S3 objects under that prefix.
        /// </summary>
        [Output("applicationSourceUri")]
        public Output<string?> ApplicationSourceUri { get; private set; } = null!;

        /// <summary>
        /// ARN of the resource.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Descriptive label for the resource, not a unique ID.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Executable path is a relative path to the game launcher executable.
        /// </summary>
        [Output("executablePath")]
        public Output<string?> ExecutablePath { get; private set; } = null!;

        /// <summary>
        /// A list of save file, registry key or log paths that are absolute paths that store game save files when the games
        /// are running on a Windows environment.
        /// </summary>
        [Output("logLocations")]
        public Output<ImmutableArray<string>> LogLocations { get; private set; } = null!;

        [Output("runtimeEnvironment")]
        public Output<Outputs.ApplicationRuntimeEnvironment?> RuntimeEnvironment { get; private set; } = null!;

        [Output("saveConfiguration")]
        public Output<Outputs.ApplicationSaveConfiguration?> SaveConfiguration { get; private set; } = null!;

        /// <summary>
        /// SaveKey is used as key to manage save files, meaning that different applications with the
        /// same SaveKey can share game save files generated by the game itself.
        /// </summary>
        [Output("saveKey")]
        public Output<string?> SaveKey { get; private set; } = null!;

        [Output("tags")]
        public Output<Outputs.ApplicationTags?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:gamecast:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:gamecast:Application", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Application(name, id, options);
        }
    }

    public sealed class ApplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ApplicationSourceUri points to a S3 Uri to replicate game files from customer S3 bucket into Motif internal S3 bucket.
        /// The uri points to a S3 prefix that could contain many objects. Motif will attempt to copy all the S3 objects under that prefix.
        /// </summary>
        [Input("applicationSourceUri")]
        public Input<string>? ApplicationSourceUri { get; set; }

        /// <summary>
        /// Descriptive label for the resource, not a unique ID.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Executable path is a relative path to the game launcher executable.
        /// </summary>
        [Input("executablePath")]
        public Input<string>? ExecutablePath { get; set; }

        [Input("logLocations")]
        private InputList<string>? _logLocations;

        /// <summary>
        /// A list of save file, registry key or log paths that are absolute paths that store game save files when the games
        /// are running on a Windows environment.
        /// </summary>
        public InputList<string> LogLocations
        {
            get => _logLocations ?? (_logLocations = new InputList<string>());
            set => _logLocations = value;
        }

        [Input("runtimeEnvironment")]
        public Input<Inputs.ApplicationRuntimeEnvironmentArgs>? RuntimeEnvironment { get; set; }

        [Input("saveConfiguration")]
        public Input<Inputs.ApplicationSaveConfigurationArgs>? SaveConfiguration { get; set; }

        /// <summary>
        /// SaveKey is used as key to manage save files, meaning that different applications with the
        /// same SaveKey can share game save files generated by the game itself.
        /// </summary>
        [Input("saveKey")]
        public Input<string>? SaveKey { get; set; }

        [Input("tags")]
        public Input<Inputs.ApplicationTagsArgs>? Tags { get; set; }

        public ApplicationArgs()
        {
        }
        public static new ApplicationArgs Empty => new ApplicationArgs();
    }
}