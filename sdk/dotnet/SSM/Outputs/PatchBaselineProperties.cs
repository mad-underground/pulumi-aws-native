// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.SSM.Outputs
{

    [OutputType]
    public sealed class PatchBaselineProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-approvalrules
        /// </summary>
        public readonly Outputs.PatchBaselineRuleGroup? ApprovalRules;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-approvedpatches
        /// </summary>
        public readonly ImmutableArray<string> ApprovedPatches;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-approvedpatchescompliancelevel
        /// </summary>
        public readonly string? ApprovedPatchesComplianceLevel;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-approvedpatchesenablenonsecurity
        /// </summary>
        public readonly bool? ApprovedPatchesEnableNonSecurity;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-globalfilters
        /// </summary>
        public readonly Outputs.PatchBaselinePatchFilterGroup? GlobalFilters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-operatingsystem
        /// </summary>
        public readonly string? OperatingSystem;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-patchgroups
        /// </summary>
        public readonly ImmutableArray<string> PatchGroups;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-rejectedpatches
        /// </summary>
        public readonly ImmutableArray<string> RejectedPatches;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-rejectedpatchesaction
        /// </summary>
        public readonly string? RejectedPatchesAction;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-sources
        /// </summary>
        public readonly ImmutableArray<Outputs.PatchBaselinePatchSource> Sources;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html#cfn-ssm-patchbaseline-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;

        [OutputConstructor]
        private PatchBaselineProperties(
            Outputs.PatchBaselineRuleGroup? ApprovalRules,

            ImmutableArray<string> ApprovedPatches,

            string? ApprovedPatchesComplianceLevel,

            bool? ApprovedPatchesEnableNonSecurity,

            string? Description,

            Outputs.PatchBaselinePatchFilterGroup? GlobalFilters,

            string Name,

            string? OperatingSystem,

            ImmutableArray<string> PatchGroups,

            ImmutableArray<string> RejectedPatches,

            string? RejectedPatchesAction,

            ImmutableArray<Outputs.PatchBaselinePatchSource> Sources,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags)
        {
            this.ApprovalRules = ApprovalRules;
            this.ApprovedPatches = ApprovedPatches;
            this.ApprovedPatchesComplianceLevel = ApprovedPatchesComplianceLevel;
            this.ApprovedPatchesEnableNonSecurity = ApprovedPatchesEnableNonSecurity;
            this.Description = Description;
            this.GlobalFilters = GlobalFilters;
            this.Name = Name;
            this.OperatingSystem = OperatingSystem;
            this.PatchGroups = PatchGroups;
            this.RejectedPatches = RejectedPatches;
            this.RejectedPatchesAction = RejectedPatchesAction;
            this.Sources = Sources;
            this.Tags = Tags;
        }
    }
}