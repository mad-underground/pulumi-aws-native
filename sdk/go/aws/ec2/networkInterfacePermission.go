// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::NetworkInterfacePermission
//
// Deprecated: NetworkInterfacePermission is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type NetworkInterfacePermission struct {
	pulumi.CustomResourceState

	AwsAccountId       pulumi.StringOutput `pulumi:"awsAccountId"`
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	Permission         pulumi.StringOutput `pulumi:"permission"`
}

// NewNetworkInterfacePermission registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfacePermission(ctx *pulumi.Context,
	name string, args *NetworkInterfacePermissionArgs, opts ...pulumi.ResourceOption) (*NetworkInterfacePermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AwsAccountId == nil {
		return nil, errors.New("invalid value for required argument 'AwsAccountId'")
	}
	if args.NetworkInterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaceId'")
	}
	if args.Permission == nil {
		return nil, errors.New("invalid value for required argument 'Permission'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"awsAccountId",
		"networkInterfaceId",
		"permission",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkInterfacePermission
	err := ctx.RegisterResource("aws-native:ec2:NetworkInterfacePermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfacePermission gets an existing NetworkInterfacePermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfacePermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfacePermissionState, opts ...pulumi.ResourceOption) (*NetworkInterfacePermission, error) {
	var resource NetworkInterfacePermission
	err := ctx.ReadResource("aws-native:ec2:NetworkInterfacePermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfacePermission resources.
type networkInterfacePermissionState struct {
}

type NetworkInterfacePermissionState struct {
}

func (NetworkInterfacePermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfacePermissionState)(nil)).Elem()
}

type networkInterfacePermissionArgs struct {
	AwsAccountId       string `pulumi:"awsAccountId"`
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	Permission         string `pulumi:"permission"`
}

// The set of arguments for constructing a NetworkInterfacePermission resource.
type NetworkInterfacePermissionArgs struct {
	AwsAccountId       pulumi.StringInput
	NetworkInterfaceId pulumi.StringInput
	Permission         pulumi.StringInput
}

func (NetworkInterfacePermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfacePermissionArgs)(nil)).Elem()
}

type NetworkInterfacePermissionInput interface {
	pulumi.Input

	ToNetworkInterfacePermissionOutput() NetworkInterfacePermissionOutput
	ToNetworkInterfacePermissionOutputWithContext(ctx context.Context) NetworkInterfacePermissionOutput
}

func (*NetworkInterfacePermission) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfacePermission)(nil)).Elem()
}

func (i *NetworkInterfacePermission) ToNetworkInterfacePermissionOutput() NetworkInterfacePermissionOutput {
	return i.ToNetworkInterfacePermissionOutputWithContext(context.Background())
}

func (i *NetworkInterfacePermission) ToNetworkInterfacePermissionOutputWithContext(ctx context.Context) NetworkInterfacePermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfacePermissionOutput)
}

type NetworkInterfacePermissionOutput struct{ *pulumi.OutputState }

func (NetworkInterfacePermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfacePermission)(nil)).Elem()
}

func (o NetworkInterfacePermissionOutput) ToNetworkInterfacePermissionOutput() NetworkInterfacePermissionOutput {
	return o
}

func (o NetworkInterfacePermissionOutput) ToNetworkInterfacePermissionOutputWithContext(ctx context.Context) NetworkInterfacePermissionOutput {
	return o
}

func (o NetworkInterfacePermissionOutput) AwsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterfacePermission) pulumi.StringOutput { return v.AwsAccountId }).(pulumi.StringOutput)
}

func (o NetworkInterfacePermissionOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterfacePermission) pulumi.StringOutput { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

func (o NetworkInterfacePermissionOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterfacePermission) pulumi.StringOutput { return v.Permission }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfacePermissionInput)(nil)).Elem(), &NetworkInterfacePermission{})
	pulumi.RegisterOutputType(NetworkInterfacePermissionOutput{})
}
