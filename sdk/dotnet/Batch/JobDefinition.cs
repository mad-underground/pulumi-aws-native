// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch
{
    /// <summary>
    /// Resource Type definition for AWS::Batch::JobDefinition
    /// </summary>
    [Obsolete(@"JobDefinition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:batch:JobDefinition")]
    public partial class JobDefinition : Pulumi.CustomResource
    {
        [Output("containerProperties")]
        public Output<Outputs.JobDefinitionContainerProperties?> ContainerProperties { get; private set; } = null!;

        [Output("jobDefinitionName")]
        public Output<string?> JobDefinitionName { get; private set; } = null!;

        [Output("nodeProperties")]
        public Output<Outputs.JobDefinitionNodeProperties?> NodeProperties { get; private set; } = null!;

        [Output("parameters")]
        public Output<object?> Parameters { get; private set; } = null!;

        [Output("platformCapabilities")]
        public Output<ImmutableArray<string>> PlatformCapabilities { get; private set; } = null!;

        [Output("propagateTags")]
        public Output<bool?> PropagateTags { get; private set; } = null!;

        [Output("retryStrategy")]
        public Output<Outputs.JobDefinitionRetryStrategy?> RetryStrategy { get; private set; } = null!;

        [Output("tags")]
        public Output<object?> Tags { get; private set; } = null!;

        [Output("timeout")]
        public Output<Outputs.JobDefinitionTimeout?> Timeout { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a JobDefinition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public JobDefinition(string name, JobDefinitionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:batch:JobDefinition", name, args ?? new JobDefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private JobDefinition(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:batch:JobDefinition", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing JobDefinition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static JobDefinition Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new JobDefinition(name, id, options);
        }
    }

    public sealed class JobDefinitionArgs : Pulumi.ResourceArgs
    {
        [Input("containerProperties")]
        public Input<Inputs.JobDefinitionContainerPropertiesArgs>? ContainerProperties { get; set; }

        [Input("jobDefinitionName")]
        public Input<string>? JobDefinitionName { get; set; }

        [Input("nodeProperties")]
        public Input<Inputs.JobDefinitionNodePropertiesArgs>? NodeProperties { get; set; }

        [Input("parameters")]
        public Input<object>? Parameters { get; set; }

        [Input("platformCapabilities")]
        private InputList<string>? _platformCapabilities;
        public InputList<string> PlatformCapabilities
        {
            get => _platformCapabilities ?? (_platformCapabilities = new InputList<string>());
            set => _platformCapabilities = value;
        }

        [Input("propagateTags")]
        public Input<bool>? PropagateTags { get; set; }

        [Input("retryStrategy")]
        public Input<Inputs.JobDefinitionRetryStrategyArgs>? RetryStrategy { get; set; }

        [Input("tags")]
        public Input<object>? Tags { get; set; }

        [Input("timeout")]
        public Input<Inputs.JobDefinitionTimeoutArgs>? Timeout { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public JobDefinitionArgs()
        {
        }
    }
}