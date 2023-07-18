// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devicefarm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::DeviceFarm::DevicePool creates a new Device Pool for a given DF Project
type DevicePool struct {
	pulumi.CustomResourceState

	Arn         pulumi.StringOutput       `pulumi:"arn"`
	Description pulumi.StringPtrOutput    `pulumi:"description"`
	MaxDevices  pulumi.IntPtrOutput       `pulumi:"maxDevices"`
	Name        pulumi.StringOutput       `pulumi:"name"`
	ProjectArn  pulumi.StringOutput       `pulumi:"projectArn"`
	Rules       DevicePoolRuleArrayOutput `pulumi:"rules"`
	Tags        DevicePoolTagArrayOutput  `pulumi:"tags"`
}

// NewDevicePool registers a new resource with the given unique name, arguments, and options.
func NewDevicePool(ctx *pulumi.Context,
	name string, args *DevicePoolArgs, opts ...pulumi.ResourceOption) (*DevicePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectArn == nil {
		return nil, errors.New("invalid value for required argument 'ProjectArn'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DevicePool
	err := ctx.RegisterResource("aws-native:devicefarm:DevicePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevicePool gets an existing DevicePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevicePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DevicePoolState, opts ...pulumi.ResourceOption) (*DevicePool, error) {
	var resource DevicePool
	err := ctx.ReadResource("aws-native:devicefarm:DevicePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DevicePool resources.
type devicePoolState struct {
}

type DevicePoolState struct {
}

func (DevicePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*devicePoolState)(nil)).Elem()
}

type devicePoolArgs struct {
	Description *string          `pulumi:"description"`
	MaxDevices  *int             `pulumi:"maxDevices"`
	Name        *string          `pulumi:"name"`
	ProjectArn  string           `pulumi:"projectArn"`
	Rules       []DevicePoolRule `pulumi:"rules"`
	Tags        []DevicePoolTag  `pulumi:"tags"`
}

// The set of arguments for constructing a DevicePool resource.
type DevicePoolArgs struct {
	Description pulumi.StringPtrInput
	MaxDevices  pulumi.IntPtrInput
	Name        pulumi.StringPtrInput
	ProjectArn  pulumi.StringInput
	Rules       DevicePoolRuleArrayInput
	Tags        DevicePoolTagArrayInput
}

func (DevicePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*devicePoolArgs)(nil)).Elem()
}

type DevicePoolInput interface {
	pulumi.Input

	ToDevicePoolOutput() DevicePoolOutput
	ToDevicePoolOutputWithContext(ctx context.Context) DevicePoolOutput
}

func (*DevicePool) ElementType() reflect.Type {
	return reflect.TypeOf((**DevicePool)(nil)).Elem()
}

func (i *DevicePool) ToDevicePoolOutput() DevicePoolOutput {
	return i.ToDevicePoolOutputWithContext(context.Background())
}

func (i *DevicePool) ToDevicePoolOutputWithContext(ctx context.Context) DevicePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicePoolOutput)
}

type DevicePoolOutput struct{ *pulumi.OutputState }

func (DevicePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DevicePool)(nil)).Elem()
}

func (o DevicePoolOutput) ToDevicePoolOutput() DevicePoolOutput {
	return o
}

func (o DevicePoolOutput) ToDevicePoolOutputWithContext(ctx context.Context) DevicePoolOutput {
	return o
}

func (o DevicePoolOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DevicePool) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o DevicePoolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevicePool) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DevicePoolOutput) MaxDevices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DevicePool) pulumi.IntPtrOutput { return v.MaxDevices }).(pulumi.IntPtrOutput)
}

func (o DevicePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DevicePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DevicePoolOutput) ProjectArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DevicePool) pulumi.StringOutput { return v.ProjectArn }).(pulumi.StringOutput)
}

func (o DevicePoolOutput) Rules() DevicePoolRuleArrayOutput {
	return o.ApplyT(func(v *DevicePool) DevicePoolRuleArrayOutput { return v.Rules }).(DevicePoolRuleArrayOutput)
}

func (o DevicePoolOutput) Tags() DevicePoolTagArrayOutput {
	return o.ApplyT(func(v *DevicePool) DevicePoolTagArrayOutput { return v.Tags }).(DevicePoolTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DevicePoolInput)(nil)).Elem(), &DevicePool{})
	pulumi.RegisterOutputType(DevicePoolOutput{})
}
