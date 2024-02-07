// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package stepfunctions

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for Activity
type Activity struct {
	pulumi.CustomResourceState

	Arn  pulumi.StringOutput          `pulumi:"arn"`
	Name pulumi.StringOutput          `pulumi:"name"`
	Tags ActivityTagsEntryArrayOutput `pulumi:"tags"`
}

// NewActivity registers a new resource with the given unique name, arguments, and options.
func NewActivity(ctx *pulumi.Context,
	name string, args *ActivityArgs, opts ...pulumi.ResourceOption) (*Activity, error) {
	if args == nil {
		args = &ActivityArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Activity
	err := ctx.RegisterResource("aws-native:stepfunctions:Activity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActivity gets an existing Activity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActivity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActivityState, opts ...pulumi.ResourceOption) (*Activity, error) {
	var resource Activity
	err := ctx.ReadResource("aws-native:stepfunctions:Activity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Activity resources.
type activityState struct {
}

type ActivityState struct {
}

func (ActivityState) ElementType() reflect.Type {
	return reflect.TypeOf((*activityState)(nil)).Elem()
}

type activityArgs struct {
	Name *string             `pulumi:"name"`
	Tags []ActivityTagsEntry `pulumi:"tags"`
}

// The set of arguments for constructing a Activity resource.
type ActivityArgs struct {
	Name pulumi.StringPtrInput
	Tags ActivityTagsEntryArrayInput
}

func (ActivityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activityArgs)(nil)).Elem()
}

type ActivityInput interface {
	pulumi.Input

	ToActivityOutput() ActivityOutput
	ToActivityOutputWithContext(ctx context.Context) ActivityOutput
}

func (*Activity) ElementType() reflect.Type {
	return reflect.TypeOf((**Activity)(nil)).Elem()
}

func (i *Activity) ToActivityOutput() ActivityOutput {
	return i.ToActivityOutputWithContext(context.Background())
}

func (i *Activity) ToActivityOutputWithContext(ctx context.Context) ActivityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityOutput)
}

type ActivityOutput struct{ *pulumi.OutputState }

func (ActivityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Activity)(nil)).Elem()
}

func (o ActivityOutput) ToActivityOutput() ActivityOutput {
	return o
}

func (o ActivityOutput) ToActivityOutputWithContext(ctx context.Context) ActivityOutput {
	return o
}

func (o ActivityOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Activity) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ActivityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Activity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ActivityOutput) Tags() ActivityTagsEntryArrayOutput {
	return o.ApplyT(func(v *Activity) ActivityTagsEntryArrayOutput { return v.Tags }).(ActivityTagsEntryArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActivityInput)(nil)).Elem(), &Activity{})
	pulumi.RegisterOutputType(ActivityOutput{})
}
