// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::InternetGateway
type InternetGateway struct {
	pulumi.CustomResourceState

	// ID of internet gateway.
	InternetGatewayId pulumi.StringOutput `pulumi:"internetGatewayId"`
	// Any tags to assign to the internet gateway.
	Tags InternetGatewayTagArrayOutput `pulumi:"tags"`
}

// NewInternetGateway registers a new resource with the given unique name, arguments, and options.
func NewInternetGateway(ctx *pulumi.Context,
	name string, args *InternetGatewayArgs, opts ...pulumi.ResourceOption) (*InternetGateway, error) {
	if args == nil {
		args = &InternetGatewayArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InternetGateway
	err := ctx.RegisterResource("aws-native:ec2:InternetGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInternetGateway gets an existing InternetGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInternetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InternetGatewayState, opts ...pulumi.ResourceOption) (*InternetGateway, error) {
	var resource InternetGateway
	err := ctx.ReadResource("aws-native:ec2:InternetGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InternetGateway resources.
type internetGatewayState struct {
}

type InternetGatewayState struct {
}

func (InternetGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*internetGatewayState)(nil)).Elem()
}

type internetGatewayArgs struct {
	// Any tags to assign to the internet gateway.
	Tags []InternetGatewayTag `pulumi:"tags"`
}

// The set of arguments for constructing a InternetGateway resource.
type InternetGatewayArgs struct {
	// Any tags to assign to the internet gateway.
	Tags InternetGatewayTagArrayInput
}

func (InternetGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*internetGatewayArgs)(nil)).Elem()
}

type InternetGatewayInput interface {
	pulumi.Input

	ToInternetGatewayOutput() InternetGatewayOutput
	ToInternetGatewayOutputWithContext(ctx context.Context) InternetGatewayOutput
}

func (*InternetGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**InternetGateway)(nil)).Elem()
}

func (i *InternetGateway) ToInternetGatewayOutput() InternetGatewayOutput {
	return i.ToInternetGatewayOutputWithContext(context.Background())
}

func (i *InternetGateway) ToInternetGatewayOutputWithContext(ctx context.Context) InternetGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetGatewayOutput)
}

type InternetGatewayOutput struct{ *pulumi.OutputState }

func (InternetGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InternetGateway)(nil)).Elem()
}

func (o InternetGatewayOutput) ToInternetGatewayOutput() InternetGatewayOutput {
	return o
}

func (o InternetGatewayOutput) ToInternetGatewayOutputWithContext(ctx context.Context) InternetGatewayOutput {
	return o
}

// ID of internet gateway.
func (o InternetGatewayOutput) InternetGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetGateway) pulumi.StringOutput { return v.InternetGatewayId }).(pulumi.StringOutput)
}

// Any tags to assign to the internet gateway.
func (o InternetGatewayOutput) Tags() InternetGatewayTagArrayOutput {
	return o.ApplyT(func(v *InternetGateway) InternetGatewayTagArrayOutput { return v.Tags }).(InternetGatewayTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InternetGatewayInput)(nil)).Elem(), &InternetGateway{})
	pulumi.RegisterOutputType(InternetGatewayOutput{})
}
