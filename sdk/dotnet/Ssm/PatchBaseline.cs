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
    /// Resource Type definition for AWS::SSM::PatchBaseline
    /// </summary>
    [AwsNativeResourceType("aws-native:ssm:PatchBaseline")]
    public partial class PatchBaseline : global::Pulumi.CustomResource
    {
        [Output("approvalRules")]
        public Output<Outputs.PatchBaselineRuleGroup?> ApprovalRules { get; private set; } = null!;

        /// <summary>
        /// A list of explicitly approved patches for the baseline.
        /// </summary>
        [Output("approvedPatches")]
        public Output<ImmutableArray<string>> ApprovedPatches { get; private set; } = null!;

        /// <summary>
        /// Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. The default value is UNSPECIFIED.
        /// </summary>
        [Output("approvedPatchesComplianceLevel")]
        public Output<Pulumi.AwsNative.Ssm.PatchBaselineApprovedPatchesComplianceLevel?> ApprovedPatchesComplianceLevel { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances. The default value is 'false'. Applies to Linux instances only.
        /// </summary>
        [Output("approvedPatchesEnableNonSecurity")]
        public Output<bool?> ApprovedPatchesEnableNonSecurity { get; private set; } = null!;

        /// <summary>
        /// The ID of the patch baseline.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// Set the baseline as default baseline. Only registering to default patch baseline is allowed.
        /// </summary>
        [Output("defaultBaseline")]
        public Output<bool?> DefaultBaseline { get; private set; } = null!;

        /// <summary>
        /// The description of the patch baseline.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A set of global filters used to include patches in the baseline.
        /// </summary>
        [Output("globalFilters")]
        public Output<Outputs.PatchBaselinePatchFilterGroup?> GlobalFilters { get; private set; } = null!;

        /// <summary>
        /// The name of the patch baseline.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Defines the operating system the patch baseline applies to. The Default value is WINDOWS.
        /// </summary>
        [Output("operatingSystem")]
        public Output<Pulumi.AwsNative.Ssm.PatchBaselineOperatingSystem?> OperatingSystem { get; private set; } = null!;

        /// <summary>
        /// PatchGroups is used to associate instances with a specific patch baseline
        /// </summary>
        [Output("patchGroups")]
        public Output<ImmutableArray<string>> PatchGroups { get; private set; } = null!;

        /// <summary>
        /// A list of explicitly rejected patches for the baseline.
        /// </summary>
        [Output("rejectedPatches")]
        public Output<ImmutableArray<string>> RejectedPatches { get; private set; } = null!;

        /// <summary>
        /// The action for Patch Manager to take on patches included in the RejectedPackages list.
        /// </summary>
        [Output("rejectedPatchesAction")]
        public Output<Pulumi.AwsNative.Ssm.PatchBaselineRejectedPatchesAction?> RejectedPatchesAction { get; private set; } = null!;

        /// <summary>
        /// Information about the patches to use to update the instances, including target operating systems and source repository. Applies to Linux instances only.
        /// </summary>
        [Output("sources")]
        public Output<ImmutableArray<Outputs.PatchBaselinePatchSource>> Sources { get; private set; } = null!;

        /// <summary>
        /// Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


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
                ReplaceOnChanges =
                {
                    "operatingSystem",
                },
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

        /// <summary>
        /// A list of explicitly approved patches for the baseline.
        /// </summary>
        public InputList<string> ApprovedPatches
        {
            get => _approvedPatches ?? (_approvedPatches = new InputList<string>());
            set => _approvedPatches = value;
        }

        /// <summary>
        /// Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. The default value is UNSPECIFIED.
        /// </summary>
        [Input("approvedPatchesComplianceLevel")]
        public Input<Pulumi.AwsNative.Ssm.PatchBaselineApprovedPatchesComplianceLevel>? ApprovedPatchesComplianceLevel { get; set; }

        /// <summary>
        /// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances. The default value is 'false'. Applies to Linux instances only.
        /// </summary>
        [Input("approvedPatchesEnableNonSecurity")]
        public Input<bool>? ApprovedPatchesEnableNonSecurity { get; set; }

        /// <summary>
        /// Set the baseline as default baseline. Only registering to default patch baseline is allowed.
        /// </summary>
        [Input("defaultBaseline")]
        public Input<bool>? DefaultBaseline { get; set; }

        /// <summary>
        /// The description of the patch baseline.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A set of global filters used to include patches in the baseline.
        /// </summary>
        [Input("globalFilters")]
        public Input<Inputs.PatchBaselinePatchFilterGroupArgs>? GlobalFilters { get; set; }

        /// <summary>
        /// The name of the patch baseline.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Defines the operating system the patch baseline applies to. The Default value is WINDOWS.
        /// </summary>
        [Input("operatingSystem")]
        public Input<Pulumi.AwsNative.Ssm.PatchBaselineOperatingSystem>? OperatingSystem { get; set; }

        [Input("patchGroups")]
        private InputList<string>? _patchGroups;

        /// <summary>
        /// PatchGroups is used to associate instances with a specific patch baseline
        /// </summary>
        public InputList<string> PatchGroups
        {
            get => _patchGroups ?? (_patchGroups = new InputList<string>());
            set => _patchGroups = value;
        }

        [Input("rejectedPatches")]
        private InputList<string>? _rejectedPatches;

        /// <summary>
        /// A list of explicitly rejected patches for the baseline.
        /// </summary>
        public InputList<string> RejectedPatches
        {
            get => _rejectedPatches ?? (_rejectedPatches = new InputList<string>());
            set => _rejectedPatches = value;
        }

        /// <summary>
        /// The action for Patch Manager to take on patches included in the RejectedPackages list.
        /// </summary>
        [Input("rejectedPatchesAction")]
        public Input<Pulumi.AwsNative.Ssm.PatchBaselineRejectedPatchesAction>? RejectedPatchesAction { get; set; }

        [Input("sources")]
        private InputList<Inputs.PatchBaselinePatchSourceArgs>? _sources;

        /// <summary>
        /// Information about the patches to use to update the instances, including target operating systems and source repository. Applies to Linux instances only.
        /// </summary>
        public InputList<Inputs.PatchBaselinePatchSourceArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.PatchBaselinePatchSourceArgs>());
            set => _sources = value;
        }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public PatchBaselineArgs()
        {
        }
        public static new PatchBaselineArgs Empty => new PatchBaselineArgs();
    }
}
