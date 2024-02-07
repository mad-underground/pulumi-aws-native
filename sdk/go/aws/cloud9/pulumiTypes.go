// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud9

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type EnvironmentEc2Repository struct {
	PathComponent string `pulumi:"pathComponent"`
	RepositoryUrl string `pulumi:"repositoryUrl"`
}

// EnvironmentEc2RepositoryInput is an input type that accepts EnvironmentEc2RepositoryArgs and EnvironmentEc2RepositoryOutput values.
// You can construct a concrete instance of `EnvironmentEc2RepositoryInput` via:
//
//	EnvironmentEc2RepositoryArgs{...}
type EnvironmentEc2RepositoryInput interface {
	pulumi.Input

	ToEnvironmentEc2RepositoryOutput() EnvironmentEc2RepositoryOutput
	ToEnvironmentEc2RepositoryOutputWithContext(context.Context) EnvironmentEc2RepositoryOutput
}

type EnvironmentEc2RepositoryArgs struct {
	PathComponent pulumi.StringInput `pulumi:"pathComponent"`
	RepositoryUrl pulumi.StringInput `pulumi:"repositoryUrl"`
}

func (EnvironmentEc2RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentEc2Repository)(nil)).Elem()
}

func (i EnvironmentEc2RepositoryArgs) ToEnvironmentEc2RepositoryOutput() EnvironmentEc2RepositoryOutput {
	return i.ToEnvironmentEc2RepositoryOutputWithContext(context.Background())
}

func (i EnvironmentEc2RepositoryArgs) ToEnvironmentEc2RepositoryOutputWithContext(ctx context.Context) EnvironmentEc2RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentEc2RepositoryOutput)
}

// EnvironmentEc2RepositoryArrayInput is an input type that accepts EnvironmentEc2RepositoryArray and EnvironmentEc2RepositoryArrayOutput values.
// You can construct a concrete instance of `EnvironmentEc2RepositoryArrayInput` via:
//
//	EnvironmentEc2RepositoryArray{ EnvironmentEc2RepositoryArgs{...} }
type EnvironmentEc2RepositoryArrayInput interface {
	pulumi.Input

	ToEnvironmentEc2RepositoryArrayOutput() EnvironmentEc2RepositoryArrayOutput
	ToEnvironmentEc2RepositoryArrayOutputWithContext(context.Context) EnvironmentEc2RepositoryArrayOutput
}

type EnvironmentEc2RepositoryArray []EnvironmentEc2RepositoryInput

func (EnvironmentEc2RepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentEc2Repository)(nil)).Elem()
}

func (i EnvironmentEc2RepositoryArray) ToEnvironmentEc2RepositoryArrayOutput() EnvironmentEc2RepositoryArrayOutput {
	return i.ToEnvironmentEc2RepositoryArrayOutputWithContext(context.Background())
}

func (i EnvironmentEc2RepositoryArray) ToEnvironmentEc2RepositoryArrayOutputWithContext(ctx context.Context) EnvironmentEc2RepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentEc2RepositoryArrayOutput)
}

type EnvironmentEc2RepositoryOutput struct{ *pulumi.OutputState }

func (EnvironmentEc2RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentEc2Repository)(nil)).Elem()
}

func (o EnvironmentEc2RepositoryOutput) ToEnvironmentEc2RepositoryOutput() EnvironmentEc2RepositoryOutput {
	return o
}

func (o EnvironmentEc2RepositoryOutput) ToEnvironmentEc2RepositoryOutputWithContext(ctx context.Context) EnvironmentEc2RepositoryOutput {
	return o
}

func (o EnvironmentEc2RepositoryOutput) PathComponent() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentEc2Repository) string { return v.PathComponent }).(pulumi.StringOutput)
}

func (o EnvironmentEc2RepositoryOutput) RepositoryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentEc2Repository) string { return v.RepositoryUrl }).(pulumi.StringOutput)
}

type EnvironmentEc2RepositoryArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentEc2RepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentEc2Repository)(nil)).Elem()
}

func (o EnvironmentEc2RepositoryArrayOutput) ToEnvironmentEc2RepositoryArrayOutput() EnvironmentEc2RepositoryArrayOutput {
	return o
}

func (o EnvironmentEc2RepositoryArrayOutput) ToEnvironmentEc2RepositoryArrayOutputWithContext(ctx context.Context) EnvironmentEc2RepositoryArrayOutput {
	return o
}

func (o EnvironmentEc2RepositoryArrayOutput) Index(i pulumi.IntInput) EnvironmentEc2RepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentEc2Repository {
		return vs[0].([]EnvironmentEc2Repository)[vs[1].(int)]
	}).(EnvironmentEc2RepositoryOutput)
}

type EnvironmentEc2Tag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// EnvironmentEc2TagInput is an input type that accepts EnvironmentEc2TagArgs and EnvironmentEc2TagOutput values.
// You can construct a concrete instance of `EnvironmentEc2TagInput` via:
//
//	EnvironmentEc2TagArgs{...}
type EnvironmentEc2TagInput interface {
	pulumi.Input

	ToEnvironmentEc2TagOutput() EnvironmentEc2TagOutput
	ToEnvironmentEc2TagOutputWithContext(context.Context) EnvironmentEc2TagOutput
}

type EnvironmentEc2TagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (EnvironmentEc2TagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentEc2Tag)(nil)).Elem()
}

func (i EnvironmentEc2TagArgs) ToEnvironmentEc2TagOutput() EnvironmentEc2TagOutput {
	return i.ToEnvironmentEc2TagOutputWithContext(context.Background())
}

func (i EnvironmentEc2TagArgs) ToEnvironmentEc2TagOutputWithContext(ctx context.Context) EnvironmentEc2TagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentEc2TagOutput)
}

// EnvironmentEc2TagArrayInput is an input type that accepts EnvironmentEc2TagArray and EnvironmentEc2TagArrayOutput values.
// You can construct a concrete instance of `EnvironmentEc2TagArrayInput` via:
//
//	EnvironmentEc2TagArray{ EnvironmentEc2TagArgs{...} }
type EnvironmentEc2TagArrayInput interface {
	pulumi.Input

	ToEnvironmentEc2TagArrayOutput() EnvironmentEc2TagArrayOutput
	ToEnvironmentEc2TagArrayOutputWithContext(context.Context) EnvironmentEc2TagArrayOutput
}

type EnvironmentEc2TagArray []EnvironmentEc2TagInput

func (EnvironmentEc2TagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentEc2Tag)(nil)).Elem()
}

func (i EnvironmentEc2TagArray) ToEnvironmentEc2TagArrayOutput() EnvironmentEc2TagArrayOutput {
	return i.ToEnvironmentEc2TagArrayOutputWithContext(context.Background())
}

func (i EnvironmentEc2TagArray) ToEnvironmentEc2TagArrayOutputWithContext(ctx context.Context) EnvironmentEc2TagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentEc2TagArrayOutput)
}

type EnvironmentEc2TagOutput struct{ *pulumi.OutputState }

func (EnvironmentEc2TagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentEc2Tag)(nil)).Elem()
}

func (o EnvironmentEc2TagOutput) ToEnvironmentEc2TagOutput() EnvironmentEc2TagOutput {
	return o
}

func (o EnvironmentEc2TagOutput) ToEnvironmentEc2TagOutputWithContext(ctx context.Context) EnvironmentEc2TagOutput {
	return o
}

func (o EnvironmentEc2TagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentEc2Tag) string { return v.Key }).(pulumi.StringOutput)
}

func (o EnvironmentEc2TagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentEc2Tag) string { return v.Value }).(pulumi.StringOutput)
}

type EnvironmentEc2TagArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentEc2TagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentEc2Tag)(nil)).Elem()
}

func (o EnvironmentEc2TagArrayOutput) ToEnvironmentEc2TagArrayOutput() EnvironmentEc2TagArrayOutput {
	return o
}

func (o EnvironmentEc2TagArrayOutput) ToEnvironmentEc2TagArrayOutputWithContext(ctx context.Context) EnvironmentEc2TagArrayOutput {
	return o
}

func (o EnvironmentEc2TagArrayOutput) Index(i pulumi.IntInput) EnvironmentEc2TagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentEc2Tag {
		return vs[0].([]EnvironmentEc2Tag)[vs[1].(int)]
	}).(EnvironmentEc2TagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentEc2RepositoryInput)(nil)).Elem(), EnvironmentEc2RepositoryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentEc2RepositoryArrayInput)(nil)).Elem(), EnvironmentEc2RepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentEc2TagInput)(nil)).Elem(), EnvironmentEc2TagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentEc2TagArrayInput)(nil)).Elem(), EnvironmentEc2TagArray{})
	pulumi.RegisterOutputType(EnvironmentEc2RepositoryOutput{})
	pulumi.RegisterOutputType(EnvironmentEc2RepositoryArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentEc2TagOutput{})
	pulumi.RegisterOutputType(EnvironmentEc2TagArrayOutput{})
}
