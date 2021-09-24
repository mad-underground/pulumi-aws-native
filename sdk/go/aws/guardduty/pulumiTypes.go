// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DetectorCFNDataSourceConfigurations struct {
	S3Logs *DetectorCFNS3LogsConfiguration `pulumi:"s3Logs"`
}

// DetectorCFNDataSourceConfigurationsInput is an input type that accepts DetectorCFNDataSourceConfigurationsArgs and DetectorCFNDataSourceConfigurationsOutput values.
// You can construct a concrete instance of `DetectorCFNDataSourceConfigurationsInput` via:
//
//          DetectorCFNDataSourceConfigurationsArgs{...}
type DetectorCFNDataSourceConfigurationsInput interface {
	pulumi.Input

	ToDetectorCFNDataSourceConfigurationsOutput() DetectorCFNDataSourceConfigurationsOutput
	ToDetectorCFNDataSourceConfigurationsOutputWithContext(context.Context) DetectorCFNDataSourceConfigurationsOutput
}

type DetectorCFNDataSourceConfigurationsArgs struct {
	S3Logs DetectorCFNS3LogsConfigurationPtrInput `pulumi:"s3Logs"`
}

func (DetectorCFNDataSourceConfigurationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DetectorCFNDataSourceConfigurations)(nil)).Elem()
}

func (i DetectorCFNDataSourceConfigurationsArgs) ToDetectorCFNDataSourceConfigurationsOutput() DetectorCFNDataSourceConfigurationsOutput {
	return i.ToDetectorCFNDataSourceConfigurationsOutputWithContext(context.Background())
}

func (i DetectorCFNDataSourceConfigurationsArgs) ToDetectorCFNDataSourceConfigurationsOutputWithContext(ctx context.Context) DetectorCFNDataSourceConfigurationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorCFNDataSourceConfigurationsOutput)
}

func (i DetectorCFNDataSourceConfigurationsArgs) ToDetectorCFNDataSourceConfigurationsPtrOutput() DetectorCFNDataSourceConfigurationsPtrOutput {
	return i.ToDetectorCFNDataSourceConfigurationsPtrOutputWithContext(context.Background())
}

func (i DetectorCFNDataSourceConfigurationsArgs) ToDetectorCFNDataSourceConfigurationsPtrOutputWithContext(ctx context.Context) DetectorCFNDataSourceConfigurationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorCFNDataSourceConfigurationsOutput).ToDetectorCFNDataSourceConfigurationsPtrOutputWithContext(ctx)
}

// DetectorCFNDataSourceConfigurationsPtrInput is an input type that accepts DetectorCFNDataSourceConfigurationsArgs, DetectorCFNDataSourceConfigurationsPtr and DetectorCFNDataSourceConfigurationsPtrOutput values.
// You can construct a concrete instance of `DetectorCFNDataSourceConfigurationsPtrInput` via:
//
//          DetectorCFNDataSourceConfigurationsArgs{...}
//
//  or:
//
//          nil
type DetectorCFNDataSourceConfigurationsPtrInput interface {
	pulumi.Input

	ToDetectorCFNDataSourceConfigurationsPtrOutput() DetectorCFNDataSourceConfigurationsPtrOutput
	ToDetectorCFNDataSourceConfigurationsPtrOutputWithContext(context.Context) DetectorCFNDataSourceConfigurationsPtrOutput
}

type detectorCFNDataSourceConfigurationsPtrType DetectorCFNDataSourceConfigurationsArgs

func DetectorCFNDataSourceConfigurationsPtr(v *DetectorCFNDataSourceConfigurationsArgs) DetectorCFNDataSourceConfigurationsPtrInput {
	return (*detectorCFNDataSourceConfigurationsPtrType)(v)
}

func (*detectorCFNDataSourceConfigurationsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DetectorCFNDataSourceConfigurations)(nil)).Elem()
}

func (i *detectorCFNDataSourceConfigurationsPtrType) ToDetectorCFNDataSourceConfigurationsPtrOutput() DetectorCFNDataSourceConfigurationsPtrOutput {
	return i.ToDetectorCFNDataSourceConfigurationsPtrOutputWithContext(context.Background())
}

func (i *detectorCFNDataSourceConfigurationsPtrType) ToDetectorCFNDataSourceConfigurationsPtrOutputWithContext(ctx context.Context) DetectorCFNDataSourceConfigurationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorCFNDataSourceConfigurationsPtrOutput)
}

type DetectorCFNDataSourceConfigurationsOutput struct{ *pulumi.OutputState }

func (DetectorCFNDataSourceConfigurationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DetectorCFNDataSourceConfigurations)(nil)).Elem()
}

func (o DetectorCFNDataSourceConfigurationsOutput) ToDetectorCFNDataSourceConfigurationsOutput() DetectorCFNDataSourceConfigurationsOutput {
	return o
}

func (o DetectorCFNDataSourceConfigurationsOutput) ToDetectorCFNDataSourceConfigurationsOutputWithContext(ctx context.Context) DetectorCFNDataSourceConfigurationsOutput {
	return o
}

func (o DetectorCFNDataSourceConfigurationsOutput) ToDetectorCFNDataSourceConfigurationsPtrOutput() DetectorCFNDataSourceConfigurationsPtrOutput {
	return o.ToDetectorCFNDataSourceConfigurationsPtrOutputWithContext(context.Background())
}

func (o DetectorCFNDataSourceConfigurationsOutput) ToDetectorCFNDataSourceConfigurationsPtrOutputWithContext(ctx context.Context) DetectorCFNDataSourceConfigurationsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DetectorCFNDataSourceConfigurations) *DetectorCFNDataSourceConfigurations {
		return &v
	}).(DetectorCFNDataSourceConfigurationsPtrOutput)
}

func (o DetectorCFNDataSourceConfigurationsOutput) S3Logs() DetectorCFNS3LogsConfigurationPtrOutput {
	return o.ApplyT(func(v DetectorCFNDataSourceConfigurations) *DetectorCFNS3LogsConfiguration { return v.S3Logs }).(DetectorCFNS3LogsConfigurationPtrOutput)
}

type DetectorCFNDataSourceConfigurationsPtrOutput struct{ *pulumi.OutputState }

func (DetectorCFNDataSourceConfigurationsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DetectorCFNDataSourceConfigurations)(nil)).Elem()
}

func (o DetectorCFNDataSourceConfigurationsPtrOutput) ToDetectorCFNDataSourceConfigurationsPtrOutput() DetectorCFNDataSourceConfigurationsPtrOutput {
	return o
}

func (o DetectorCFNDataSourceConfigurationsPtrOutput) ToDetectorCFNDataSourceConfigurationsPtrOutputWithContext(ctx context.Context) DetectorCFNDataSourceConfigurationsPtrOutput {
	return o
}

func (o DetectorCFNDataSourceConfigurationsPtrOutput) Elem() DetectorCFNDataSourceConfigurationsOutput {
	return o.ApplyT(func(v *DetectorCFNDataSourceConfigurations) DetectorCFNDataSourceConfigurations {
		if v != nil {
			return *v
		}
		var ret DetectorCFNDataSourceConfigurations
		return ret
	}).(DetectorCFNDataSourceConfigurationsOutput)
}

func (o DetectorCFNDataSourceConfigurationsPtrOutput) S3Logs() DetectorCFNS3LogsConfigurationPtrOutput {
	return o.ApplyT(func(v *DetectorCFNDataSourceConfigurations) *DetectorCFNS3LogsConfiguration {
		if v == nil {
			return nil
		}
		return v.S3Logs
	}).(DetectorCFNS3LogsConfigurationPtrOutput)
}

type DetectorCFNS3LogsConfiguration struct {
	Enable *bool `pulumi:"enable"`
}

// DetectorCFNS3LogsConfigurationInput is an input type that accepts DetectorCFNS3LogsConfigurationArgs and DetectorCFNS3LogsConfigurationOutput values.
// You can construct a concrete instance of `DetectorCFNS3LogsConfigurationInput` via:
//
//          DetectorCFNS3LogsConfigurationArgs{...}
type DetectorCFNS3LogsConfigurationInput interface {
	pulumi.Input

	ToDetectorCFNS3LogsConfigurationOutput() DetectorCFNS3LogsConfigurationOutput
	ToDetectorCFNS3LogsConfigurationOutputWithContext(context.Context) DetectorCFNS3LogsConfigurationOutput
}

type DetectorCFNS3LogsConfigurationArgs struct {
	Enable pulumi.BoolPtrInput `pulumi:"enable"`
}

func (DetectorCFNS3LogsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DetectorCFNS3LogsConfiguration)(nil)).Elem()
}

func (i DetectorCFNS3LogsConfigurationArgs) ToDetectorCFNS3LogsConfigurationOutput() DetectorCFNS3LogsConfigurationOutput {
	return i.ToDetectorCFNS3LogsConfigurationOutputWithContext(context.Background())
}

func (i DetectorCFNS3LogsConfigurationArgs) ToDetectorCFNS3LogsConfigurationOutputWithContext(ctx context.Context) DetectorCFNS3LogsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorCFNS3LogsConfigurationOutput)
}

func (i DetectorCFNS3LogsConfigurationArgs) ToDetectorCFNS3LogsConfigurationPtrOutput() DetectorCFNS3LogsConfigurationPtrOutput {
	return i.ToDetectorCFNS3LogsConfigurationPtrOutputWithContext(context.Background())
}

func (i DetectorCFNS3LogsConfigurationArgs) ToDetectorCFNS3LogsConfigurationPtrOutputWithContext(ctx context.Context) DetectorCFNS3LogsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorCFNS3LogsConfigurationOutput).ToDetectorCFNS3LogsConfigurationPtrOutputWithContext(ctx)
}

// DetectorCFNS3LogsConfigurationPtrInput is an input type that accepts DetectorCFNS3LogsConfigurationArgs, DetectorCFNS3LogsConfigurationPtr and DetectorCFNS3LogsConfigurationPtrOutput values.
// You can construct a concrete instance of `DetectorCFNS3LogsConfigurationPtrInput` via:
//
//          DetectorCFNS3LogsConfigurationArgs{...}
//
//  or:
//
//          nil
type DetectorCFNS3LogsConfigurationPtrInput interface {
	pulumi.Input

	ToDetectorCFNS3LogsConfigurationPtrOutput() DetectorCFNS3LogsConfigurationPtrOutput
	ToDetectorCFNS3LogsConfigurationPtrOutputWithContext(context.Context) DetectorCFNS3LogsConfigurationPtrOutput
}

type detectorCFNS3LogsConfigurationPtrType DetectorCFNS3LogsConfigurationArgs

func DetectorCFNS3LogsConfigurationPtr(v *DetectorCFNS3LogsConfigurationArgs) DetectorCFNS3LogsConfigurationPtrInput {
	return (*detectorCFNS3LogsConfigurationPtrType)(v)
}

func (*detectorCFNS3LogsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DetectorCFNS3LogsConfiguration)(nil)).Elem()
}

func (i *detectorCFNS3LogsConfigurationPtrType) ToDetectorCFNS3LogsConfigurationPtrOutput() DetectorCFNS3LogsConfigurationPtrOutput {
	return i.ToDetectorCFNS3LogsConfigurationPtrOutputWithContext(context.Background())
}

func (i *detectorCFNS3LogsConfigurationPtrType) ToDetectorCFNS3LogsConfigurationPtrOutputWithContext(ctx context.Context) DetectorCFNS3LogsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorCFNS3LogsConfigurationPtrOutput)
}

type DetectorCFNS3LogsConfigurationOutput struct{ *pulumi.OutputState }

func (DetectorCFNS3LogsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DetectorCFNS3LogsConfiguration)(nil)).Elem()
}

func (o DetectorCFNS3LogsConfigurationOutput) ToDetectorCFNS3LogsConfigurationOutput() DetectorCFNS3LogsConfigurationOutput {
	return o
}

func (o DetectorCFNS3LogsConfigurationOutput) ToDetectorCFNS3LogsConfigurationOutputWithContext(ctx context.Context) DetectorCFNS3LogsConfigurationOutput {
	return o
}

func (o DetectorCFNS3LogsConfigurationOutput) ToDetectorCFNS3LogsConfigurationPtrOutput() DetectorCFNS3LogsConfigurationPtrOutput {
	return o.ToDetectorCFNS3LogsConfigurationPtrOutputWithContext(context.Background())
}

func (o DetectorCFNS3LogsConfigurationOutput) ToDetectorCFNS3LogsConfigurationPtrOutputWithContext(ctx context.Context) DetectorCFNS3LogsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DetectorCFNS3LogsConfiguration) *DetectorCFNS3LogsConfiguration {
		return &v
	}).(DetectorCFNS3LogsConfigurationPtrOutput)
}

func (o DetectorCFNS3LogsConfigurationOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DetectorCFNS3LogsConfiguration) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

type DetectorCFNS3LogsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (DetectorCFNS3LogsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DetectorCFNS3LogsConfiguration)(nil)).Elem()
}

func (o DetectorCFNS3LogsConfigurationPtrOutput) ToDetectorCFNS3LogsConfigurationPtrOutput() DetectorCFNS3LogsConfigurationPtrOutput {
	return o
}

func (o DetectorCFNS3LogsConfigurationPtrOutput) ToDetectorCFNS3LogsConfigurationPtrOutputWithContext(ctx context.Context) DetectorCFNS3LogsConfigurationPtrOutput {
	return o
}

func (o DetectorCFNS3LogsConfigurationPtrOutput) Elem() DetectorCFNS3LogsConfigurationOutput {
	return o.ApplyT(func(v *DetectorCFNS3LogsConfiguration) DetectorCFNS3LogsConfiguration {
		if v != nil {
			return *v
		}
		var ret DetectorCFNS3LogsConfiguration
		return ret
	}).(DetectorCFNS3LogsConfigurationOutput)
}

func (o DetectorCFNS3LogsConfigurationPtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DetectorCFNS3LogsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.Enable
	}).(pulumi.BoolPtrOutput)
}

type FilterCondition struct {
	Eq  []string `pulumi:"eq"`
	Gte *int     `pulumi:"gte"`
	Lt  *int     `pulumi:"lt"`
	Lte *int     `pulumi:"lte"`
	Neq []string `pulumi:"neq"`
}

// FilterConditionInput is an input type that accepts FilterConditionArgs and FilterConditionOutput values.
// You can construct a concrete instance of `FilterConditionInput` via:
//
//          FilterConditionArgs{...}
type FilterConditionInput interface {
	pulumi.Input

	ToFilterConditionOutput() FilterConditionOutput
	ToFilterConditionOutputWithContext(context.Context) FilterConditionOutput
}

type FilterConditionArgs struct {
	Eq  pulumi.StringArrayInput `pulumi:"eq"`
	Gte pulumi.IntPtrInput      `pulumi:"gte"`
	Lt  pulumi.IntPtrInput      `pulumi:"lt"`
	Lte pulumi.IntPtrInput      `pulumi:"lte"`
	Neq pulumi.StringArrayInput `pulumi:"neq"`
}

func (FilterConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterCondition)(nil)).Elem()
}

func (i FilterConditionArgs) ToFilterConditionOutput() FilterConditionOutput {
	return i.ToFilterConditionOutputWithContext(context.Background())
}

func (i FilterConditionArgs) ToFilterConditionOutputWithContext(ctx context.Context) FilterConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterConditionOutput)
}

func (i FilterConditionArgs) ToFilterConditionPtrOutput() FilterConditionPtrOutput {
	return i.ToFilterConditionPtrOutputWithContext(context.Background())
}

func (i FilterConditionArgs) ToFilterConditionPtrOutputWithContext(ctx context.Context) FilterConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterConditionOutput).ToFilterConditionPtrOutputWithContext(ctx)
}

// FilterConditionPtrInput is an input type that accepts FilterConditionArgs, FilterConditionPtr and FilterConditionPtrOutput values.
// You can construct a concrete instance of `FilterConditionPtrInput` via:
//
//          FilterConditionArgs{...}
//
//  or:
//
//          nil
type FilterConditionPtrInput interface {
	pulumi.Input

	ToFilterConditionPtrOutput() FilterConditionPtrOutput
	ToFilterConditionPtrOutputWithContext(context.Context) FilterConditionPtrOutput
}

type filterConditionPtrType FilterConditionArgs

func FilterConditionPtr(v *FilterConditionArgs) FilterConditionPtrInput {
	return (*filterConditionPtrType)(v)
}

func (*filterConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FilterCondition)(nil)).Elem()
}

func (i *filterConditionPtrType) ToFilterConditionPtrOutput() FilterConditionPtrOutput {
	return i.ToFilterConditionPtrOutputWithContext(context.Background())
}

func (i *filterConditionPtrType) ToFilterConditionPtrOutputWithContext(ctx context.Context) FilterConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterConditionPtrOutput)
}

type FilterConditionOutput struct{ *pulumi.OutputState }

func (FilterConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterCondition)(nil)).Elem()
}

func (o FilterConditionOutput) ToFilterConditionOutput() FilterConditionOutput {
	return o
}

func (o FilterConditionOutput) ToFilterConditionOutputWithContext(ctx context.Context) FilterConditionOutput {
	return o
}

func (o FilterConditionOutput) ToFilterConditionPtrOutput() FilterConditionPtrOutput {
	return o.ToFilterConditionPtrOutputWithContext(context.Background())
}

func (o FilterConditionOutput) ToFilterConditionPtrOutputWithContext(ctx context.Context) FilterConditionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FilterCondition) *FilterCondition {
		return &v
	}).(FilterConditionPtrOutput)
}

func (o FilterConditionOutput) Eq() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FilterCondition) []string { return v.Eq }).(pulumi.StringArrayOutput)
}

func (o FilterConditionOutput) Gte() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FilterCondition) *int { return v.Gte }).(pulumi.IntPtrOutput)
}

func (o FilterConditionOutput) Lt() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FilterCondition) *int { return v.Lt }).(pulumi.IntPtrOutput)
}

func (o FilterConditionOutput) Lte() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FilterCondition) *int { return v.Lte }).(pulumi.IntPtrOutput)
}

func (o FilterConditionOutput) Neq() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FilterCondition) []string { return v.Neq }).(pulumi.StringArrayOutput)
}

type FilterConditionPtrOutput struct{ *pulumi.OutputState }

func (FilterConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FilterCondition)(nil)).Elem()
}

func (o FilterConditionPtrOutput) ToFilterConditionPtrOutput() FilterConditionPtrOutput {
	return o
}

func (o FilterConditionPtrOutput) ToFilterConditionPtrOutputWithContext(ctx context.Context) FilterConditionPtrOutput {
	return o
}

func (o FilterConditionPtrOutput) Elem() FilterConditionOutput {
	return o.ApplyT(func(v *FilterCondition) FilterCondition {
		if v != nil {
			return *v
		}
		var ret FilterCondition
		return ret
	}).(FilterConditionOutput)
}

func (o FilterConditionPtrOutput) Eq() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FilterCondition) []string {
		if v == nil {
			return nil
		}
		return v.Eq
	}).(pulumi.StringArrayOutput)
}

func (o FilterConditionPtrOutput) Gte() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FilterCondition) *int {
		if v == nil {
			return nil
		}
		return v.Gte
	}).(pulumi.IntPtrOutput)
}

func (o FilterConditionPtrOutput) Lt() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FilterCondition) *int {
		if v == nil {
			return nil
		}
		return v.Lt
	}).(pulumi.IntPtrOutput)
}

func (o FilterConditionPtrOutput) Lte() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FilterCondition) *int {
		if v == nil {
			return nil
		}
		return v.Lte
	}).(pulumi.IntPtrOutput)
}

func (o FilterConditionPtrOutput) Neq() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FilterCondition) []string {
		if v == nil {
			return nil
		}
		return v.Neq
	}).(pulumi.StringArrayOutput)
}

type FilterFindingCriteria struct {
	Criterion interface{}      `pulumi:"criterion"`
	ItemType  *FilterCondition `pulumi:"itemType"`
}

// FilterFindingCriteriaInput is an input type that accepts FilterFindingCriteriaArgs and FilterFindingCriteriaOutput values.
// You can construct a concrete instance of `FilterFindingCriteriaInput` via:
//
//          FilterFindingCriteriaArgs{...}
type FilterFindingCriteriaInput interface {
	pulumi.Input

	ToFilterFindingCriteriaOutput() FilterFindingCriteriaOutput
	ToFilterFindingCriteriaOutputWithContext(context.Context) FilterFindingCriteriaOutput
}

type FilterFindingCriteriaArgs struct {
	Criterion pulumi.Input            `pulumi:"criterion"`
	ItemType  FilterConditionPtrInput `pulumi:"itemType"`
}

func (FilterFindingCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterFindingCriteria)(nil)).Elem()
}

func (i FilterFindingCriteriaArgs) ToFilterFindingCriteriaOutput() FilterFindingCriteriaOutput {
	return i.ToFilterFindingCriteriaOutputWithContext(context.Background())
}

func (i FilterFindingCriteriaArgs) ToFilterFindingCriteriaOutputWithContext(ctx context.Context) FilterFindingCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFindingCriteriaOutput)
}

func (i FilterFindingCriteriaArgs) ToFilterFindingCriteriaPtrOutput() FilterFindingCriteriaPtrOutput {
	return i.ToFilterFindingCriteriaPtrOutputWithContext(context.Background())
}

func (i FilterFindingCriteriaArgs) ToFilterFindingCriteriaPtrOutputWithContext(ctx context.Context) FilterFindingCriteriaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFindingCriteriaOutput).ToFilterFindingCriteriaPtrOutputWithContext(ctx)
}

// FilterFindingCriteriaPtrInput is an input type that accepts FilterFindingCriteriaArgs, FilterFindingCriteriaPtr and FilterFindingCriteriaPtrOutput values.
// You can construct a concrete instance of `FilterFindingCriteriaPtrInput` via:
//
//          FilterFindingCriteriaArgs{...}
//
//  or:
//
//          nil
type FilterFindingCriteriaPtrInput interface {
	pulumi.Input

	ToFilterFindingCriteriaPtrOutput() FilterFindingCriteriaPtrOutput
	ToFilterFindingCriteriaPtrOutputWithContext(context.Context) FilterFindingCriteriaPtrOutput
}

type filterFindingCriteriaPtrType FilterFindingCriteriaArgs

func FilterFindingCriteriaPtr(v *FilterFindingCriteriaArgs) FilterFindingCriteriaPtrInput {
	return (*filterFindingCriteriaPtrType)(v)
}

func (*filterFindingCriteriaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FilterFindingCriteria)(nil)).Elem()
}

func (i *filterFindingCriteriaPtrType) ToFilterFindingCriteriaPtrOutput() FilterFindingCriteriaPtrOutput {
	return i.ToFilterFindingCriteriaPtrOutputWithContext(context.Background())
}

func (i *filterFindingCriteriaPtrType) ToFilterFindingCriteriaPtrOutputWithContext(ctx context.Context) FilterFindingCriteriaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFindingCriteriaPtrOutput)
}

type FilterFindingCriteriaOutput struct{ *pulumi.OutputState }

func (FilterFindingCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterFindingCriteria)(nil)).Elem()
}

func (o FilterFindingCriteriaOutput) ToFilterFindingCriteriaOutput() FilterFindingCriteriaOutput {
	return o
}

func (o FilterFindingCriteriaOutput) ToFilterFindingCriteriaOutputWithContext(ctx context.Context) FilterFindingCriteriaOutput {
	return o
}

func (o FilterFindingCriteriaOutput) ToFilterFindingCriteriaPtrOutput() FilterFindingCriteriaPtrOutput {
	return o.ToFilterFindingCriteriaPtrOutputWithContext(context.Background())
}

func (o FilterFindingCriteriaOutput) ToFilterFindingCriteriaPtrOutputWithContext(ctx context.Context) FilterFindingCriteriaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FilterFindingCriteria) *FilterFindingCriteria {
		return &v
	}).(FilterFindingCriteriaPtrOutput)
}

func (o FilterFindingCriteriaOutput) Criterion() pulumi.AnyOutput {
	return o.ApplyT(func(v FilterFindingCriteria) interface{} { return v.Criterion }).(pulumi.AnyOutput)
}

func (o FilterFindingCriteriaOutput) ItemType() FilterConditionPtrOutput {
	return o.ApplyT(func(v FilterFindingCriteria) *FilterCondition { return v.ItemType }).(FilterConditionPtrOutput)
}

type FilterFindingCriteriaPtrOutput struct{ *pulumi.OutputState }

func (FilterFindingCriteriaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FilterFindingCriteria)(nil)).Elem()
}

func (o FilterFindingCriteriaPtrOutput) ToFilterFindingCriteriaPtrOutput() FilterFindingCriteriaPtrOutput {
	return o
}

func (o FilterFindingCriteriaPtrOutput) ToFilterFindingCriteriaPtrOutputWithContext(ctx context.Context) FilterFindingCriteriaPtrOutput {
	return o
}

func (o FilterFindingCriteriaPtrOutput) Elem() FilterFindingCriteriaOutput {
	return o.ApplyT(func(v *FilterFindingCriteria) FilterFindingCriteria {
		if v != nil {
			return *v
		}
		var ret FilterFindingCriteria
		return ret
	}).(FilterFindingCriteriaOutput)
}

func (o FilterFindingCriteriaPtrOutput) Criterion() pulumi.AnyOutput {
	return o.ApplyT(func(v *FilterFindingCriteria) interface{} {
		if v == nil {
			return nil
		}
		return v.Criterion
	}).(pulumi.AnyOutput)
}

func (o FilterFindingCriteriaPtrOutput) ItemType() FilterConditionPtrOutput {
	return o.ApplyT(func(v *FilterFindingCriteria) *FilterCondition {
		if v == nil {
			return nil
		}
		return v.ItemType
	}).(FilterConditionPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DetectorCFNDataSourceConfigurationsOutput{})
	pulumi.RegisterOutputType(DetectorCFNDataSourceConfigurationsPtrOutput{})
	pulumi.RegisterOutputType(DetectorCFNS3LogsConfigurationOutput{})
	pulumi.RegisterOutputType(DetectorCFNS3LogsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(FilterConditionOutput{})
	pulumi.RegisterOutputType(FilterConditionPtrOutput{})
	pulumi.RegisterOutputType(FilterFindingCriteriaOutput{})
	pulumi.RegisterOutputType(FilterFindingCriteriaPtrOutput{})
}