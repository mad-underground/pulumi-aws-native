// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Redshift::ClusterSecurityGroupIngress
//
// Deprecated: ClusterSecurityGroupIngress is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ClusterSecurityGroupIngress struct {
	pulumi.CustomResourceState

	CIDRIP                   pulumi.StringPtrOutput `pulumi:"cIDRIP"`
	ClusterSecurityGroupName pulumi.StringOutput    `pulumi:"clusterSecurityGroupName"`
	EC2SecurityGroupName     pulumi.StringPtrOutput `pulumi:"eC2SecurityGroupName"`
	EC2SecurityGroupOwnerId  pulumi.StringPtrOutput `pulumi:"eC2SecurityGroupOwnerId"`
}

// NewClusterSecurityGroupIngress registers a new resource with the given unique name, arguments, and options.
func NewClusterSecurityGroupIngress(ctx *pulumi.Context,
	name string, args *ClusterSecurityGroupIngressArgs, opts ...pulumi.ResourceOption) (*ClusterSecurityGroupIngress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterSecurityGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterSecurityGroupName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterSecurityGroupIngress
	err := ctx.RegisterResource("aws-native:redshift:ClusterSecurityGroupIngress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterSecurityGroupIngress gets an existing ClusterSecurityGroupIngress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterSecurityGroupIngress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterSecurityGroupIngressState, opts ...pulumi.ResourceOption) (*ClusterSecurityGroupIngress, error) {
	var resource ClusterSecurityGroupIngress
	err := ctx.ReadResource("aws-native:redshift:ClusterSecurityGroupIngress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterSecurityGroupIngress resources.
type clusterSecurityGroupIngressState struct {
}

type ClusterSecurityGroupIngressState struct {
}

func (ClusterSecurityGroupIngressState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterSecurityGroupIngressState)(nil)).Elem()
}

type clusterSecurityGroupIngressArgs struct {
	CIDRIP                   *string `pulumi:"cIDRIP"`
	ClusterSecurityGroupName string  `pulumi:"clusterSecurityGroupName"`
	EC2SecurityGroupName     *string `pulumi:"eC2SecurityGroupName"`
	EC2SecurityGroupOwnerId  *string `pulumi:"eC2SecurityGroupOwnerId"`
}

// The set of arguments for constructing a ClusterSecurityGroupIngress resource.
type ClusterSecurityGroupIngressArgs struct {
	CIDRIP                   pulumi.StringPtrInput
	ClusterSecurityGroupName pulumi.StringInput
	EC2SecurityGroupName     pulumi.StringPtrInput
	EC2SecurityGroupOwnerId  pulumi.StringPtrInput
}

func (ClusterSecurityGroupIngressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterSecurityGroupIngressArgs)(nil)).Elem()
}

type ClusterSecurityGroupIngressInput interface {
	pulumi.Input

	ToClusterSecurityGroupIngressOutput() ClusterSecurityGroupIngressOutput
	ToClusterSecurityGroupIngressOutputWithContext(ctx context.Context) ClusterSecurityGroupIngressOutput
}

func (*ClusterSecurityGroupIngress) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSecurityGroupIngress)(nil)).Elem()
}

func (i *ClusterSecurityGroupIngress) ToClusterSecurityGroupIngressOutput() ClusterSecurityGroupIngressOutput {
	return i.ToClusterSecurityGroupIngressOutputWithContext(context.Background())
}

func (i *ClusterSecurityGroupIngress) ToClusterSecurityGroupIngressOutputWithContext(ctx context.Context) ClusterSecurityGroupIngressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSecurityGroupIngressOutput)
}

type ClusterSecurityGroupIngressOutput struct{ *pulumi.OutputState }

func (ClusterSecurityGroupIngressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSecurityGroupIngress)(nil)).Elem()
}

func (o ClusterSecurityGroupIngressOutput) ToClusterSecurityGroupIngressOutput() ClusterSecurityGroupIngressOutput {
	return o
}

func (o ClusterSecurityGroupIngressOutput) ToClusterSecurityGroupIngressOutputWithContext(ctx context.Context) ClusterSecurityGroupIngressOutput {
	return o
}

func (o ClusterSecurityGroupIngressOutput) CIDRIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSecurityGroupIngress) pulumi.StringPtrOutput { return v.CIDRIP }).(pulumi.StringPtrOutput)
}

func (o ClusterSecurityGroupIngressOutput) ClusterSecurityGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterSecurityGroupIngress) pulumi.StringOutput { return v.ClusterSecurityGroupName }).(pulumi.StringOutput)
}

func (o ClusterSecurityGroupIngressOutput) EC2SecurityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSecurityGroupIngress) pulumi.StringPtrOutput { return v.EC2SecurityGroupName }).(pulumi.StringPtrOutput)
}

func (o ClusterSecurityGroupIngressOutput) EC2SecurityGroupOwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSecurityGroupIngress) pulumi.StringPtrOutput { return v.EC2SecurityGroupOwnerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSecurityGroupIngressInput)(nil)).Elem(), &ClusterSecurityGroupIngress{})
	pulumi.RegisterOutputType(ClusterSecurityGroupIngressOutput{})
}
