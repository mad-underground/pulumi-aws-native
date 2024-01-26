// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::GuardDuty::Filter
func LookupFilter(ctx *pulumi.Context, args *LookupFilterArgs, opts ...pulumi.InvokeOption) (*LookupFilterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFilterResult
	err := ctx.Invoke("aws-native:guardduty:getFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFilterArgs struct {
	DetectorId string `pulumi:"detectorId"`
	Name       string `pulumi:"name"`
}

type LookupFilterResult struct {
	Action          *string                `pulumi:"action"`
	Description     *string                `pulumi:"description"`
	FindingCriteria *FilterFindingCriteria `pulumi:"findingCriteria"`
	Rank            *int                   `pulumi:"rank"`
	Tags            []FilterTagItem        `pulumi:"tags"`
}

func LookupFilterOutput(ctx *pulumi.Context, args LookupFilterOutputArgs, opts ...pulumi.InvokeOption) LookupFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFilterResult, error) {
			args := v.(LookupFilterArgs)
			r, err := LookupFilter(ctx, &args, opts...)
			var s LookupFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFilterResultOutput)
}

type LookupFilterOutputArgs struct {
	DetectorId pulumi.StringInput `pulumi:"detectorId"`
	Name       pulumi.StringInput `pulumi:"name"`
}

func (LookupFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFilterArgs)(nil)).Elem()
}

type LookupFilterResultOutput struct{ *pulumi.OutputState }

func (LookupFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFilterResult)(nil)).Elem()
}

func (o LookupFilterResultOutput) ToLookupFilterResultOutput() LookupFilterResultOutput {
	return o
}

func (o LookupFilterResultOutput) ToLookupFilterResultOutputWithContext(ctx context.Context) LookupFilterResultOutput {
	return o
}

func (o LookupFilterResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFilterResult] {
	return pulumix.Output[LookupFilterResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupFilterResultOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFilterResult) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o LookupFilterResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFilterResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFilterResultOutput) FindingCriteria() FilterFindingCriteriaPtrOutput {
	return o.ApplyT(func(v LookupFilterResult) *FilterFindingCriteria { return v.FindingCriteria }).(FilterFindingCriteriaPtrOutput)
}

func (o LookupFilterResultOutput) Rank() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFilterResult) *int { return v.Rank }).(pulumi.IntPtrOutput)
}

func (o LookupFilterResultOutput) Tags() FilterTagItemArrayOutput {
	return o.ApplyT(func(v LookupFilterResult) []FilterTagItem { return v.Tags }).(FilterTagItemArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFilterResultOutput{})
}
