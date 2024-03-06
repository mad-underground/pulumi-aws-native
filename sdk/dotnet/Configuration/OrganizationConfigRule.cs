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
    /// Resource Type definition for AWS::Config::OrganizationConfigRule
    /// </summary>
    [Obsolete(@"OrganizationConfigRule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:configuration:OrganizationConfigRule")]
    public partial class OrganizationConfigRule : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("excludedAccounts")]
        public Output<ImmutableArray<string>> ExcludedAccounts { get; private set; } = null!;

        [Output("organizationConfigRuleName")]
        public Output<string> OrganizationConfigRuleName { get; private set; } = null!;

        [Output("organizationCustomPolicyRuleMetadata")]
        public Output<Outputs.OrganizationConfigRuleOrganizationCustomPolicyRuleMetadata?> OrganizationCustomPolicyRuleMetadata { get; private set; } = null!;

        [Output("organizationCustomRuleMetadata")]
        public Output<Outputs.OrganizationConfigRuleOrganizationCustomRuleMetadata?> OrganizationCustomRuleMetadata { get; private set; } = null!;

        [Output("organizationManagedRuleMetadata")]
        public Output<Outputs.OrganizationConfigRuleOrganizationManagedRuleMetadata?> OrganizationManagedRuleMetadata { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationConfigRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationConfigRule(string name, OrganizationConfigRuleArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:configuration:OrganizationConfigRule", name, args ?? new OrganizationConfigRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationConfigRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:configuration:OrganizationConfigRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "organizationConfigRuleName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OrganizationConfigRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationConfigRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OrganizationConfigRule(name, id, options);
        }
    }

    public sealed class OrganizationConfigRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludedAccounts")]
        private InputList<string>? _excludedAccounts;
        public InputList<string> ExcludedAccounts
        {
            get => _excludedAccounts ?? (_excludedAccounts = new InputList<string>());
            set => _excludedAccounts = value;
        }

        [Input("organizationConfigRuleName")]
        public Input<string>? OrganizationConfigRuleName { get; set; }

        [Input("organizationCustomPolicyRuleMetadata")]
        public Input<Inputs.OrganizationConfigRuleOrganizationCustomPolicyRuleMetadataArgs>? OrganizationCustomPolicyRuleMetadata { get; set; }

        [Input("organizationCustomRuleMetadata")]
        public Input<Inputs.OrganizationConfigRuleOrganizationCustomRuleMetadataArgs>? OrganizationCustomRuleMetadata { get; set; }

        [Input("organizationManagedRuleMetadata")]
        public Input<Inputs.OrganizationConfigRuleOrganizationManagedRuleMetadataArgs>? OrganizationManagedRuleMetadata { get; set; }

        public OrganizationConfigRuleArgs()
        {
        }
        public static new OrganizationConfigRuleArgs Empty => new OrganizationConfigRuleArgs();
    }
}
