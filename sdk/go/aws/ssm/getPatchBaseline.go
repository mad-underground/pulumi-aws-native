// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SSM::PatchBaseline
func LookupPatchBaseline(ctx *pulumi.Context, args *LookupPatchBaselineArgs, opts ...pulumi.InvokeOption) (*LookupPatchBaselineResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPatchBaselineResult
	err := ctx.Invoke("aws-native:ssm:getPatchBaseline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPatchBaselineArgs struct {
	// The ID of the patch baseline.
	Id string `pulumi:"id"`
}

type LookupPatchBaselineResult struct {
	ApprovalRules *PatchBaselineRuleGroup `pulumi:"approvalRules"`
	// A list of explicitly approved patches for the baseline.
	ApprovedPatches []string `pulumi:"approvedPatches"`
	// Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. The default value is UNSPECIFIED.
	ApprovedPatchesComplianceLevel *PatchBaselineApprovedPatchesComplianceLevel `pulumi:"approvedPatchesComplianceLevel"`
	// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances. The default value is 'false'. Applies to Linux instances only.
	ApprovedPatchesEnableNonSecurity *bool `pulumi:"approvedPatchesEnableNonSecurity"`
	// Set the baseline as default baseline. Only registering to default patch baseline is allowed.
	DefaultBaseline *bool `pulumi:"defaultBaseline"`
	// The description of the patch baseline.
	Description *string `pulumi:"description"`
	// A set of global filters used to include patches in the baseline.
	GlobalFilters *PatchBaselinePatchFilterGroup `pulumi:"globalFilters"`
	// The ID of the patch baseline.
	Id *string `pulumi:"id"`
	// The name of the patch baseline.
	Name *string `pulumi:"name"`
	// PatchGroups is used to associate instances with a specific patch baseline
	PatchGroups []string `pulumi:"patchGroups"`
	// A list of explicitly rejected patches for the baseline.
	RejectedPatches []string `pulumi:"rejectedPatches"`
	// The action for Patch Manager to take on patches included in the RejectedPackages list.
	RejectedPatchesAction *PatchBaselineRejectedPatchesAction `pulumi:"rejectedPatchesAction"`
	// Information about the patches to use to update the instances, including target operating systems and source repository. Applies to Linux instances only.
	Sources []PatchBaselinePatchSource `pulumi:"sources"`
	// Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways.
	Tags []PatchBaselineTag `pulumi:"tags"`
}

func LookupPatchBaselineOutput(ctx *pulumi.Context, args LookupPatchBaselineOutputArgs, opts ...pulumi.InvokeOption) LookupPatchBaselineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPatchBaselineResult, error) {
			args := v.(LookupPatchBaselineArgs)
			r, err := LookupPatchBaseline(ctx, &args, opts...)
			var s LookupPatchBaselineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPatchBaselineResultOutput)
}

type LookupPatchBaselineOutputArgs struct {
	// The ID of the patch baseline.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupPatchBaselineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPatchBaselineArgs)(nil)).Elem()
}

type LookupPatchBaselineResultOutput struct{ *pulumi.OutputState }

func (LookupPatchBaselineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPatchBaselineResult)(nil)).Elem()
}

func (o LookupPatchBaselineResultOutput) ToLookupPatchBaselineResultOutput() LookupPatchBaselineResultOutput {
	return o
}

func (o LookupPatchBaselineResultOutput) ToLookupPatchBaselineResultOutputWithContext(ctx context.Context) LookupPatchBaselineResultOutput {
	return o
}

func (o LookupPatchBaselineResultOutput) ApprovalRules() PatchBaselineRuleGroupPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *PatchBaselineRuleGroup { return v.ApprovalRules }).(PatchBaselineRuleGroupPtrOutput)
}

// A list of explicitly approved patches for the baseline.
func (o LookupPatchBaselineResultOutput) ApprovedPatches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []string { return v.ApprovedPatches }).(pulumi.StringArrayOutput)
}

// Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. The default value is UNSPECIFIED.
func (o LookupPatchBaselineResultOutput) ApprovedPatchesComplianceLevel() PatchBaselineApprovedPatchesComplianceLevelPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *PatchBaselineApprovedPatchesComplianceLevel {
		return v.ApprovedPatchesComplianceLevel
	}).(PatchBaselineApprovedPatchesComplianceLevelPtrOutput)
}

// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances. The default value is 'false'. Applies to Linux instances only.
func (o LookupPatchBaselineResultOutput) ApprovedPatchesEnableNonSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *bool { return v.ApprovedPatchesEnableNonSecurity }).(pulumi.BoolPtrOutput)
}

// Set the baseline as default baseline. Only registering to default patch baseline is allowed.
func (o LookupPatchBaselineResultOutput) DefaultBaseline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *bool { return v.DefaultBaseline }).(pulumi.BoolPtrOutput)
}

// The description of the patch baseline.
func (o LookupPatchBaselineResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A set of global filters used to include patches in the baseline.
func (o LookupPatchBaselineResultOutput) GlobalFilters() PatchBaselinePatchFilterGroupPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *PatchBaselinePatchFilterGroup { return v.GlobalFilters }).(PatchBaselinePatchFilterGroupPtrOutput)
}

// The ID of the patch baseline.
func (o LookupPatchBaselineResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the patch baseline.
func (o LookupPatchBaselineResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// PatchGroups is used to associate instances with a specific patch baseline
func (o LookupPatchBaselineResultOutput) PatchGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []string { return v.PatchGroups }).(pulumi.StringArrayOutput)
}

// A list of explicitly rejected patches for the baseline.
func (o LookupPatchBaselineResultOutput) RejectedPatches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []string { return v.RejectedPatches }).(pulumi.StringArrayOutput)
}

// The action for Patch Manager to take on patches included in the RejectedPackages list.
func (o LookupPatchBaselineResultOutput) RejectedPatchesAction() PatchBaselineRejectedPatchesActionPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *PatchBaselineRejectedPatchesAction { return v.RejectedPatchesAction }).(PatchBaselineRejectedPatchesActionPtrOutput)
}

// Information about the patches to use to update the instances, including target operating systems and source repository. Applies to Linux instances only.
func (o LookupPatchBaselineResultOutput) Sources() PatchBaselinePatchSourceArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []PatchBaselinePatchSource { return v.Sources }).(PatchBaselinePatchSourceArrayOutput)
}

// Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways.
func (o LookupPatchBaselineResultOutput) Tags() PatchBaselineTagArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []PatchBaselineTag { return v.Tags }).(PatchBaselineTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPatchBaselineResultOutput{})
}
