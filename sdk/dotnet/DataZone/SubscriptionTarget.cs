// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone
{
    /// <summary>
    /// Subscription targets enables one to access the data to which you have subscribed in your projects.
    /// </summary>
    [AwsNativeResourceType("aws-native:datazone:SubscriptionTarget")]
    public partial class SubscriptionTarget : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The asset types that can be included in the subscription target.
        /// </summary>
        [Output("applicableAssetTypes")]
        public Output<ImmutableArray<string>> ApplicableAssetTypes { get; private set; } = null!;

        /// <summary>
        /// The authorized principals of the subscription target.
        /// </summary>
        [Output("authorizedPrincipals")]
        public Output<ImmutableArray<string>> AuthorizedPrincipals { get; private set; } = null!;

        /// <summary>
        /// The ID of the subscription target.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// The timestamp of when the subscription target was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The Amazon DataZone user who created the subscription target.
        /// </summary>
        [Output("createdBy")]
        public Output<string> CreatedBy { get; private set; } = null!;

        /// <summary>
        /// The ID of the Amazon DataZone domain in which subscription target is created.
        /// </summary>
        [Output("domainId")]
        public Output<string> DomainId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Amazon DataZone domain in which subscription target would be created.
        /// </summary>
        [Output("domainIdentifier")]
        public Output<string> DomainIdentifier { get; private set; } = null!;

        /// <summary>
        /// The ID of the environment in which subscription target is created.
        /// </summary>
        [Output("environmentId")]
        public Output<string> EnvironmentId { get; private set; } = null!;

        /// <summary>
        /// The ID of the environment in which subscription target would be created.
        /// </summary>
        [Output("environmentIdentifier")]
        public Output<string> EnvironmentIdentifier { get; private set; } = null!;

        /// <summary>
        /// The manage access role that is used to create the subscription target.
        /// </summary>
        [Output("manageAccessRole")]
        public Output<string> ManageAccessRole { get; private set; } = null!;

        /// <summary>
        /// The name of the subscription target.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The identifier of the project specified in the subscription target.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The provider of the subscription target.
        /// </summary>
        [Output("provider")]
        public Output<string?> Provider { get; private set; } = null!;

        /// <summary>
        /// The configuration of the subscription target.
        /// </summary>
        [Output("subscriptionTargetConfig")]
        public Output<ImmutableArray<Outputs.SubscriptionTargetForm>> SubscriptionTargetConfig { get; private set; } = null!;

        /// <summary>
        /// The type of the subscription target.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The timestamp of when the subscription target was updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The Amazon DataZone user who updated the subscription target.
        /// </summary>
        [Output("updatedBy")]
        public Output<string> UpdatedBy { get; private set; } = null!;


        /// <summary>
        /// Create a SubscriptionTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubscriptionTarget(string name, SubscriptionTargetArgs args, CustomResourceOptions? options = null)
            : base("aws-native:datazone:SubscriptionTarget", name, args ?? new SubscriptionTargetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SubscriptionTarget(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:datazone:SubscriptionTarget", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "domainIdentifier",
                    "environmentIdentifier",
                    "type",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SubscriptionTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubscriptionTarget Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SubscriptionTarget(name, id, options);
        }
    }

    public sealed class SubscriptionTargetArgs : global::Pulumi.ResourceArgs
    {
        [Input("applicableAssetTypes", required: true)]
        private InputList<string>? _applicableAssetTypes;

        /// <summary>
        /// The asset types that can be included in the subscription target.
        /// </summary>
        public InputList<string> ApplicableAssetTypes
        {
            get => _applicableAssetTypes ?? (_applicableAssetTypes = new InputList<string>());
            set => _applicableAssetTypes = value;
        }

        [Input("authorizedPrincipals", required: true)]
        private InputList<string>? _authorizedPrincipals;

        /// <summary>
        /// The authorized principals of the subscription target.
        /// </summary>
        public InputList<string> AuthorizedPrincipals
        {
            get => _authorizedPrincipals ?? (_authorizedPrincipals = new InputList<string>());
            set => _authorizedPrincipals = value;
        }

        /// <summary>
        /// The ID of the Amazon DataZone domain in which subscription target would be created.
        /// </summary>
        [Input("domainIdentifier", required: true)]
        public Input<string> DomainIdentifier { get; set; } = null!;

        /// <summary>
        /// The ID of the environment in which subscription target would be created.
        /// </summary>
        [Input("environmentIdentifier", required: true)]
        public Input<string> EnvironmentIdentifier { get; set; } = null!;

        /// <summary>
        /// The manage access role that is used to create the subscription target.
        /// </summary>
        [Input("manageAccessRole", required: true)]
        public Input<string> ManageAccessRole { get; set; } = null!;

        /// <summary>
        /// The name of the subscription target.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The provider of the subscription target.
        /// </summary>
        [Input("provider")]
        public Input<string>? Provider { get; set; }

        [Input("subscriptionTargetConfig", required: true)]
        private InputList<Inputs.SubscriptionTargetFormArgs>? _subscriptionTargetConfig;

        /// <summary>
        /// The configuration of the subscription target.
        /// </summary>
        public InputList<Inputs.SubscriptionTargetFormArgs> SubscriptionTargetConfig
        {
            get => _subscriptionTargetConfig ?? (_subscriptionTargetConfig = new InputList<Inputs.SubscriptionTargetFormArgs>());
            set => _subscriptionTargetConfig = value;
        }

        /// <summary>
        /// The type of the subscription target.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SubscriptionTargetArgs()
        {
        }
        public static new SubscriptionTargetArgs Empty => new SubscriptionTargetArgs();
    }
}
