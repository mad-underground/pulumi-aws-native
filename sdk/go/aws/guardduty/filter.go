// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GuardDuty::Filter
type Filter struct {
	pulumi.CustomResourceState

	Action          pulumi.StringPtrOutput      `pulumi:"action"`
	Description     pulumi.StringPtrOutput      `pulumi:"description"`
	DetectorId      pulumi.StringPtrOutput      `pulumi:"detectorId"`
	FindingCriteria FilterFindingCriteriaOutput `pulumi:"findingCriteria"`
	Name            pulumi.StringPtrOutput      `pulumi:"name"`
	Rank            pulumi.IntPtrOutput         `pulumi:"rank"`
	Tags            FilterTagItemArrayOutput    `pulumi:"tags"`
}

// NewFilter registers a new resource with the given unique name, arguments, and options.
func NewFilter(ctx *pulumi.Context,
	name string, args *FilterArgs, opts ...pulumi.ResourceOption) (*Filter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FindingCriteria == nil {
		return nil, errors.New("invalid value for required argument 'FindingCriteria'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"detectorId",
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Filter
	err := ctx.RegisterResource("aws-native:guardduty:Filter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFilter gets an existing Filter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FilterState, opts ...pulumi.ResourceOption) (*Filter, error) {
	var resource Filter
	err := ctx.ReadResource("aws-native:guardduty:Filter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Filter resources.
type filterState struct {
}

type FilterState struct {
}

func (FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*filterState)(nil)).Elem()
}

type filterArgs struct {
	Action          *string               `pulumi:"action"`
	Description     *string               `pulumi:"description"`
	DetectorId      *string               `pulumi:"detectorId"`
	FindingCriteria FilterFindingCriteria `pulumi:"findingCriteria"`
	Name            *string               `pulumi:"name"`
	Rank            *int                  `pulumi:"rank"`
	Tags            []FilterTagItem       `pulumi:"tags"`
}

// The set of arguments for constructing a Filter resource.
type FilterArgs struct {
	Action          pulumi.StringPtrInput
	Description     pulumi.StringPtrInput
	DetectorId      pulumi.StringPtrInput
	FindingCriteria FilterFindingCriteriaInput
	Name            pulumi.StringPtrInput
	Rank            pulumi.IntPtrInput
	Tags            FilterTagItemArrayInput
}

func (FilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*filterArgs)(nil)).Elem()
}

type FilterInput interface {
	pulumi.Input

	ToFilterOutput() FilterOutput
	ToFilterOutputWithContext(ctx context.Context) FilterOutput
}

func (*Filter) ElementType() reflect.Type {
	return reflect.TypeOf((**Filter)(nil)).Elem()
}

func (i *Filter) ToFilterOutput() FilterOutput {
	return i.ToFilterOutputWithContext(context.Background())
}

func (i *Filter) ToFilterOutputWithContext(ctx context.Context) FilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterOutput)
}

type FilterOutput struct{ *pulumi.OutputState }

func (FilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Filter)(nil)).Elem()
}

func (o FilterOutput) ToFilterOutput() FilterOutput {
	return o
}

func (o FilterOutput) ToFilterOutputWithContext(ctx context.Context) FilterOutput {
	return o
}

func (o FilterOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringPtrOutput { return v.Action }).(pulumi.StringPtrOutput)
}

func (o FilterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o FilterOutput) DetectorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringPtrOutput { return v.DetectorId }).(pulumi.StringPtrOutput)
}

func (o FilterOutput) FindingCriteria() FilterFindingCriteriaOutput {
	return o.ApplyT(func(v *Filter) FilterFindingCriteriaOutput { return v.FindingCriteria }).(FilterFindingCriteriaOutput)
}

func (o FilterOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Filter) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FilterOutput) Rank() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Filter) pulumi.IntPtrOutput { return v.Rank }).(pulumi.IntPtrOutput)
}

func (o FilterOutput) Tags() FilterTagItemArrayOutput {
	return o.ApplyT(func(v *Filter) FilterTagItemArrayOutput { return v.Tags }).(FilterTagItemArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FilterInput)(nil)).Elem(), &Filter{})
	pulumi.RegisterOutputType(FilterOutput{})
}
