// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gamecast

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::GameCast::StreamGroup Resource Type
type StreamGroup struct {
	pulumi.CustomResourceState

	// ARN of the resource.
	Arn                pulumi.StringOutput                    `pulumi:"arn"`
	DefaultApplication StreamGroupDefaultApplicationPtrOutput `pulumi:"defaultApplication"`
	// Descriptive label for the resource, not a unique ID.
	Description pulumi.StringOutput `pulumi:"description"`
	// DesiredCapacity is the target number of stream sessions customer sets.
	DesiredCapacity pulumi.IntPtrOutput             `pulumi:"desiredCapacity"`
	StreamClass     StreamGroupStreamClassPtrOutput `pulumi:"streamClass"`
	Tags            StreamGroupTagsPtrOutput        `pulumi:"tags"`
}

// NewStreamGroup registers a new resource with the given unique name, arguments, and options.
func NewStreamGroup(ctx *pulumi.Context,
	name string, args *StreamGroupArgs, opts ...pulumi.ResourceOption) (*StreamGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StreamGroup
	err := ctx.RegisterResource("aws-native:gamecast:StreamGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamGroup gets an existing StreamGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamGroupState, opts ...pulumi.ResourceOption) (*StreamGroup, error) {
	var resource StreamGroup
	err := ctx.ReadResource("aws-native:gamecast:StreamGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamGroup resources.
type streamGroupState struct {
}

type StreamGroupState struct {
}

func (StreamGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamGroupState)(nil)).Elem()
}

type streamGroupArgs struct {
	DefaultApplication *StreamGroupDefaultApplication `pulumi:"defaultApplication"`
	// Descriptive label for the resource, not a unique ID.
	Description string `pulumi:"description"`
	// DesiredCapacity is the target number of stream sessions customer sets.
	DesiredCapacity *int                    `pulumi:"desiredCapacity"`
	StreamClass     *StreamGroupStreamClass `pulumi:"streamClass"`
	Tags            *StreamGroupTags        `pulumi:"tags"`
}

// The set of arguments for constructing a StreamGroup resource.
type StreamGroupArgs struct {
	DefaultApplication StreamGroupDefaultApplicationPtrInput
	// Descriptive label for the resource, not a unique ID.
	Description pulumi.StringInput
	// DesiredCapacity is the target number of stream sessions customer sets.
	DesiredCapacity pulumi.IntPtrInput
	StreamClass     StreamGroupStreamClassPtrInput
	Tags            StreamGroupTagsPtrInput
}

func (StreamGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamGroupArgs)(nil)).Elem()
}

type StreamGroupInput interface {
	pulumi.Input

	ToStreamGroupOutput() StreamGroupOutput
	ToStreamGroupOutputWithContext(ctx context.Context) StreamGroupOutput
}

func (*StreamGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamGroup)(nil)).Elem()
}

func (i *StreamGroup) ToStreamGroupOutput() StreamGroupOutput {
	return i.ToStreamGroupOutputWithContext(context.Background())
}

func (i *StreamGroup) ToStreamGroupOutputWithContext(ctx context.Context) StreamGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamGroupOutput)
}

type StreamGroupOutput struct{ *pulumi.OutputState }

func (StreamGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamGroup)(nil)).Elem()
}

func (o StreamGroupOutput) ToStreamGroupOutput() StreamGroupOutput {
	return o
}

func (o StreamGroupOutput) ToStreamGroupOutputWithContext(ctx context.Context) StreamGroupOutput {
	return o
}

// ARN of the resource.
func (o StreamGroupOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamGroup) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o StreamGroupOutput) DefaultApplication() StreamGroupDefaultApplicationPtrOutput {
	return o.ApplyT(func(v *StreamGroup) StreamGroupDefaultApplicationPtrOutput { return v.DefaultApplication }).(StreamGroupDefaultApplicationPtrOutput)
}

// Descriptive label for the resource, not a unique ID.
func (o StreamGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamGroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// DesiredCapacity is the target number of stream sessions customer sets.
func (o StreamGroupOutput) DesiredCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StreamGroup) pulumi.IntPtrOutput { return v.DesiredCapacity }).(pulumi.IntPtrOutput)
}

func (o StreamGroupOutput) StreamClass() StreamGroupStreamClassPtrOutput {
	return o.ApplyT(func(v *StreamGroup) StreamGroupStreamClassPtrOutput { return v.StreamClass }).(StreamGroupStreamClassPtrOutput)
}

func (o StreamGroupOutput) Tags() StreamGroupTagsPtrOutput {
	return o.ApplyT(func(v *StreamGroup) StreamGroupTagsPtrOutput { return v.Tags }).(StreamGroupTagsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StreamGroupInput)(nil)).Elem(), &StreamGroup{})
	pulumi.RegisterOutputType(StreamGroupOutput{})
}
