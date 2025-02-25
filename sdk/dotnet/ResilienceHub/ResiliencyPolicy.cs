// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ResilienceHub
{
    /// <summary>
    /// Resource Type Definition for Resiliency Policy.
    /// </summary>
    [AwsNativeResourceType("aws-native:resiliencehub:ResiliencyPolicy")]
    public partial class ResiliencyPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Data Location Constraint of the Policy.
        /// </summary>
        [Output("dataLocationConstraint")]
        public Output<Pulumi.AwsNative.ResilienceHub.ResiliencyPolicyDataLocationConstraint?> DataLocationConstraint { get; private set; } = null!;

        [Output("policy")]
        public Output<ImmutableDictionary<string, Outputs.ResiliencyPolicyFailurePolicy>> Policy { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the Resiliency Policy.
        /// </summary>
        [Output("policyArn")]
        public Output<string> PolicyArn { get; private set; } = null!;

        /// <summary>
        /// Description of Resiliency Policy.
        /// </summary>
        [Output("policyDescription")]
        public Output<string?> PolicyDescription { get; private set; } = null!;

        /// <summary>
        /// Name of Resiliency Policy.
        /// </summary>
        [Output("policyName")]
        public Output<string> PolicyName { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resiliency Policy Tier.
        /// </summary>
        [Output("tier")]
        public Output<Pulumi.AwsNative.ResilienceHub.ResiliencyPolicyTier> Tier { get; private set; } = null!;


        /// <summary>
        /// Create a ResiliencyPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResiliencyPolicy(string name, ResiliencyPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:resiliencehub:ResiliencyPolicy", name, args ?? new ResiliencyPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResiliencyPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:resiliencehub:ResiliencyPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ResiliencyPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResiliencyPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ResiliencyPolicy(name, id, options);
        }
    }

    public sealed class ResiliencyPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Data Location Constraint of the Policy.
        /// </summary>
        [Input("dataLocationConstraint")]
        public Input<Pulumi.AwsNative.ResilienceHub.ResiliencyPolicyDataLocationConstraint>? DataLocationConstraint { get; set; }

        [Input("policy", required: true)]
        private InputMap<Inputs.ResiliencyPolicyFailurePolicyArgs>? _policy;
        public InputMap<Inputs.ResiliencyPolicyFailurePolicyArgs> Policy
        {
            get => _policy ?? (_policy = new InputMap<Inputs.ResiliencyPolicyFailurePolicyArgs>());
            set => _policy = value;
        }

        /// <summary>
        /// Description of Resiliency Policy.
        /// </summary>
        [Input("policyDescription")]
        public Input<string>? PolicyDescription { get; set; }

        /// <summary>
        /// Name of Resiliency Policy.
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Resiliency Policy Tier.
        /// </summary>
        [Input("tier", required: true)]
        public Input<Pulumi.AwsNative.ResilienceHub.ResiliencyPolicyTier> Tier { get; set; } = null!;

        public ResiliencyPolicyArgs()
        {
        }
        public static new ResiliencyPolicyArgs Empty => new ResiliencyPolicyArgs();
    }
}
