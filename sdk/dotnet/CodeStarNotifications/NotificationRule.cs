// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeStarNotifications
{
    /// <summary>
    /// Resource Type definition for AWS::CodeStarNotifications::NotificationRule
    /// </summary>
    [AwsNativeResourceType("aws-native:codestarnotifications:NotificationRule")]
    public partial class NotificationRule : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("createdBy")]
        public Output<string?> CreatedBy { get; private set; } = null!;

        [Output("detailType")]
        public Output<Pulumi.AwsNative.CodeStarNotifications.NotificationRuleDetailType> DetailType { get; private set; } = null!;

        [Output("eventTypeId")]
        public Output<string?> EventTypeId { get; private set; } = null!;

        [Output("eventTypeIds")]
        public Output<ImmutableArray<string>> EventTypeIds { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("resource")]
        public Output<string> Resource { get; private set; } = null!;

        [Output("status")]
        public Output<Pulumi.AwsNative.CodeStarNotifications.NotificationRuleStatus?> Status { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("targetAddress")]
        public Output<string?> TargetAddress { get; private set; } = null!;

        [Output("targets")]
        public Output<ImmutableArray<Outputs.NotificationRuleTarget>> Targets { get; private set; } = null!;


        /// <summary>
        /// Create a NotificationRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NotificationRule(string name, NotificationRuleArgs args, CustomResourceOptions? options = null)
            : base("aws-native:codestarnotifications:NotificationRule", name, args ?? new NotificationRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NotificationRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:codestarnotifications:NotificationRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "resource",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NotificationRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NotificationRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NotificationRule(name, id, options);
        }
    }

    public sealed class NotificationRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("createdBy")]
        public Input<string>? CreatedBy { get; set; }

        [Input("detailType", required: true)]
        public Input<Pulumi.AwsNative.CodeStarNotifications.NotificationRuleDetailType> DetailType { get; set; } = null!;

        [Input("eventTypeId")]
        public Input<string>? EventTypeId { get; set; }

        [Input("eventTypeIds", required: true)]
        private InputList<string>? _eventTypeIds;
        public InputList<string> EventTypeIds
        {
            get => _eventTypeIds ?? (_eventTypeIds = new InputList<string>());
            set => _eventTypeIds = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        [Input("status")]
        public Input<Pulumi.AwsNative.CodeStarNotifications.NotificationRuleStatus>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("targetAddress")]
        public Input<string>? TargetAddress { get; set; }

        [Input("targets", required: true)]
        private InputList<Inputs.NotificationRuleTargetArgs>? _targets;
        public InputList<Inputs.NotificationRuleTargetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.NotificationRuleTargetArgs>());
            set => _targets = value;
        }

        public NotificationRuleArgs()
        {
        }
        public static new NotificationRuleArgs Empty => new NotificationRuleArgs();
    }
}
