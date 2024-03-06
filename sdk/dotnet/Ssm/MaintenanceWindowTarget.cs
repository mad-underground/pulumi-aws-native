// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ssm
{
    /// <summary>
    /// Resource Type definition for AWS::SSM::MaintenanceWindowTarget
    /// </summary>
    [Obsolete(@"MaintenanceWindowTarget is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ssm:MaintenanceWindowTarget")]
    public partial class MaintenanceWindowTarget : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("ownerInformation")]
        public Output<string?> OwnerInformation { get; private set; } = null!;

        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;

        [Output("targets")]
        public Output<ImmutableArray<Outputs.MaintenanceWindowTargetTargets>> Targets { get; private set; } = null!;

        [Output("windowId")]
        public Output<string> WindowId { get; private set; } = null!;


        /// <summary>
        /// Create a MaintenanceWindowTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MaintenanceWindowTarget(string name, MaintenanceWindowTargetArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ssm:MaintenanceWindowTarget", name, args ?? new MaintenanceWindowTargetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MaintenanceWindowTarget(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ssm:MaintenanceWindowTarget", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "windowId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MaintenanceWindowTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MaintenanceWindowTarget Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MaintenanceWindowTarget(name, id, options);
        }
    }

    public sealed class MaintenanceWindowTargetArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ownerInformation")]
        public Input<string>? OwnerInformation { get; set; }

        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        [Input("targets", required: true)]
        private InputList<Inputs.MaintenanceWindowTargetTargetsArgs>? _targets;
        public InputList<Inputs.MaintenanceWindowTargetTargetsArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.MaintenanceWindowTargetTargetsArgs>());
            set => _targets = value;
        }

        [Input("windowId", required: true)]
        public Input<string> WindowId { get; set; } = null!;

        public MaintenanceWindowTargetArgs()
        {
        }
        public static new MaintenanceWindowTargetArgs Empty => new MaintenanceWindowTargetArgs();
    }
}
