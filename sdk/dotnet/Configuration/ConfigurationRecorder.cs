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
    /// Resource Type definition for AWS::Config::ConfigurationRecorder
    /// </summary>
    [Obsolete(@"ConfigurationRecorder is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:configuration:ConfigurationRecorder")]
    public partial class ConfigurationRecorder : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("recordingGroup")]
        public Output<Outputs.ConfigurationRecorderRecordingGroup?> RecordingGroup { get; private set; } = null!;

        [Output("recordingMode")]
        public Output<Outputs.ConfigurationRecorderRecordingMode?> RecordingMode { get; private set; } = null!;

        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigurationRecorder resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigurationRecorder(string name, ConfigurationRecorderArgs args, CustomResourceOptions? options = null)
            : base("aws-native:configuration:ConfigurationRecorder", name, args ?? new ConfigurationRecorderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigurationRecorder(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:configuration:ConfigurationRecorder", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigurationRecorder resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigurationRecorder Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConfigurationRecorder(name, id, options);
        }
    }

    public sealed class ConfigurationRecorderArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("recordingGroup")]
        public Input<Inputs.ConfigurationRecorderRecordingGroupArgs>? RecordingGroup { get; set; }

        [Input("recordingMode")]
        public Input<Inputs.ConfigurationRecorderRecordingModeArgs>? RecordingMode { get; set; }

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        public ConfigurationRecorderArgs()
        {
        }
        public static new ConfigurationRecorderArgs Empty => new ConfigurationRecorderArgs();
    }
}
