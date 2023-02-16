// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSM
{
    /// <summary>
    /// Resource Type definition for AWS::SSM::PatchBaseline
    /// </summary>
    [Obsolete(@"PatchBaseline is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ssm:PatchBaseline")]
    public partial class PatchBaseline : global::Pulumi.CustomResource
    {
        [Output("approvalRules")]
        public Output<Outputs.PatchBaselineRuleGroup?> ApprovalRules { get; private set; } = null!;

        [Output("approvedPatches")]
        public Output<ImmutableArray<string>> ApprovedPatches { get; private set; } = null!;

        [Output("approvedPatchesComplianceLevel")]
        public Output<string?> ApprovedPatchesComplianceLevel { get; private set; } = null!;

        [Output("approvedPatchesEnableNonSecurity")]
        public Output<bool?> ApprovedPatchesEnableNonSecurity { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("globalFilters")]
        public Output<Outputs.PatchBaselinePatchFilterGroup?> GlobalFilters { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("operatingSystem")]
        public Output<string?> OperatingSystem { get; private set; } = null!;

        [Output("patchGroups")]
        public Output<ImmutableArray<string>> PatchGroups { get; private set; } = null!;

        [Output("rejectedPatches")]
        public Output<ImmutableArray<string>> RejectedPatches { get; private set; } = null!;

        [Output("rejectedPatchesAction")]
        public Output<string?> RejectedPatchesAction { get; private set; } = null!;

        [Output("sources")]
        public Output<ImmutableArray<Outputs.PatchBaselinePatchSource>> Sources { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.PatchBaselineTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a PatchBaseline resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PatchBaseline(string name, PatchBaselineArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ssm:PatchBaseline", name, args ?? new PatchBaselineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PatchBaseline(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ssm:PatchBaseline", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PatchBaseline resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PatchBaseline Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PatchBaseline(name, id, options);
        }
    }

    public sealed class PatchBaselineArgs : global::Pulumi.ResourceArgs
    {
        [Input("approvalRules")]
        public Input<Inputs.PatchBaselineRuleGroupArgs>? ApprovalRules { get; set; }

        [Input("approvedPatches")]
        private InputList<string>? _approvedPatches;
        public InputList<string> ApprovedPatches
        {
            get => _approvedPatches ?? (_approvedPatches = new InputList<string>());
            set => _approvedPatches = value;
        }

        [Input("approvedPatchesComplianceLevel")]
        public Input<string>? ApprovedPatchesComplianceLevel { get; set; }

        [Input("approvedPatchesEnableNonSecurity")]
        public Input<bool>? ApprovedPatchesEnableNonSecurity { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("globalFilters")]
        public Input<Inputs.PatchBaselinePatchFilterGroupArgs>? GlobalFilters { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("operatingSystem")]
        public Input<string>? OperatingSystem { get; set; }

        [Input("patchGroups")]
        private InputList<string>? _patchGroups;
        public InputList<string> PatchGroups
        {
            get => _patchGroups ?? (_patchGroups = new InputList<string>());
            set => _patchGroups = value;
        }

        [Input("rejectedPatches")]
        private InputList<string>? _rejectedPatches;
        public InputList<string> RejectedPatches
        {
            get => _rejectedPatches ?? (_rejectedPatches = new InputList<string>());
            set => _rejectedPatches = value;
        }

        [Input("rejectedPatchesAction")]
        public Input<string>? RejectedPatchesAction { get; set; }

        [Input("sources")]
        private InputList<Inputs.PatchBaselinePatchSourceArgs>? _sources;
        public InputList<Inputs.PatchBaselinePatchSourceArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.PatchBaselinePatchSourceArgs>());
            set => _sources = value;
        }

        [Input("tags")]
        private InputList<Inputs.PatchBaselineTagArgs>? _tags;
        public InputList<Inputs.PatchBaselineTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.PatchBaselineTagArgs>());
            set => _tags = value;
        }

        public PatchBaselineArgs()
        {
        }
        public static new PatchBaselineArgs Empty => new PatchBaselineArgs();
    }
}
