// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Connect::SecurityKey
type SecurityKey struct {
	pulumi.CustomResourceState

	AssociationId pulumi.StringOutput `pulumi:"associationId"`
	InstanceId    pulumi.StringOutput `pulumi:"instanceId"`
	Key           pulumi.StringOutput `pulumi:"key"`
}

// NewSecurityKey registers a new resource with the given unique name, arguments, and options.
func NewSecurityKey(ctx *pulumi.Context,
	name string, args *SecurityKeyArgs, opts ...pulumi.ResourceOption) (*SecurityKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecurityKey
	err := ctx.RegisterResource("aws-native:connect:SecurityKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityKey gets an existing SecurityKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityKeyState, opts ...pulumi.ResourceOption) (*SecurityKey, error) {
	var resource SecurityKey
	err := ctx.ReadResource("aws-native:connect:SecurityKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityKey resources.
type securityKeyState struct {
}

type SecurityKeyState struct {
}

func (SecurityKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityKeyState)(nil)).Elem()
}

type securityKeyArgs struct {
	InstanceId string `pulumi:"instanceId"`
	Key        string `pulumi:"key"`
}

// The set of arguments for constructing a SecurityKey resource.
type SecurityKeyArgs struct {
	InstanceId pulumi.StringInput
	Key        pulumi.StringInput
}

func (SecurityKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityKeyArgs)(nil)).Elem()
}

type SecurityKeyInput interface {
	pulumi.Input

	ToSecurityKeyOutput() SecurityKeyOutput
	ToSecurityKeyOutputWithContext(ctx context.Context) SecurityKeyOutput
}

func (*SecurityKey) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityKey)(nil)).Elem()
}

func (i *SecurityKey) ToSecurityKeyOutput() SecurityKeyOutput {
	return i.ToSecurityKeyOutputWithContext(context.Background())
}

func (i *SecurityKey) ToSecurityKeyOutputWithContext(ctx context.Context) SecurityKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityKeyOutput)
}

type SecurityKeyOutput struct{ *pulumi.OutputState }

func (SecurityKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityKey)(nil)).Elem()
}

func (o SecurityKeyOutput) ToSecurityKeyOutput() SecurityKeyOutput {
	return o
}

func (o SecurityKeyOutput) ToSecurityKeyOutputWithContext(ctx context.Context) SecurityKeyOutput {
	return o
}

func (o SecurityKeyOutput) AssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityKey) pulumi.StringOutput { return v.AssociationId }).(pulumi.StringOutput)
}

func (o SecurityKeyOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityKey) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

func (o SecurityKeyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityKey) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityKeyInput)(nil)).Elem(), &SecurityKey{})
	pulumi.RegisterOutputType(SecurityKeyOutput{})
}
