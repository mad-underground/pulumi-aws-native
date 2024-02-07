// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53resolver

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Route53Resolver::ResolverEndpoint
//
// Deprecated: ResolverEndpoint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ResolverEndpoint struct {
	pulumi.CustomResourceState

	Arn                   pulumi.StringOutput                         `pulumi:"arn"`
	Direction             pulumi.StringOutput                         `pulumi:"direction"`
	HostVpcId             pulumi.StringOutput                         `pulumi:"hostVpcId"`
	IpAddressCount        pulumi.StringOutput                         `pulumi:"ipAddressCount"`
	IpAddresses           ResolverEndpointIpAddressRequestArrayOutput `pulumi:"ipAddresses"`
	Name                  pulumi.StringPtrOutput                      `pulumi:"name"`
	OutpostArn            pulumi.StringPtrOutput                      `pulumi:"outpostArn"`
	PreferredInstanceType pulumi.StringPtrOutput                      `pulumi:"preferredInstanceType"`
	Protocols             pulumi.StringArrayOutput                    `pulumi:"protocols"`
	ResolverEndpointId    pulumi.StringOutput                         `pulumi:"resolverEndpointId"`
	ResolverEndpointType  pulumi.StringPtrOutput                      `pulumi:"resolverEndpointType"`
	SecurityGroupIds      pulumi.StringArrayOutput                    `pulumi:"securityGroupIds"`
	Tags                  ResolverEndpointTagArrayOutput              `pulumi:"tags"`
}

// NewResolverEndpoint registers a new resource with the given unique name, arguments, and options.
func NewResolverEndpoint(ctx *pulumi.Context,
	name string, args *ResolverEndpointArgs, opts ...pulumi.ResourceOption) (*ResolverEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.IpAddresses == nil {
		return nil, errors.New("invalid value for required argument 'IpAddresses'")
	}
	if args.SecurityGroupIds == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupIds'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"direction",
		"outpostArn",
		"preferredInstanceType",
		"securityGroupIds[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResolverEndpoint
	err := ctx.RegisterResource("aws-native:route53resolver:ResolverEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverEndpoint gets an existing ResolverEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverEndpointState, opts ...pulumi.ResourceOption) (*ResolverEndpoint, error) {
	var resource ResolverEndpoint
	err := ctx.ReadResource("aws-native:route53resolver:ResolverEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverEndpoint resources.
type resolverEndpointState struct {
}

type ResolverEndpointState struct {
}

func (ResolverEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverEndpointState)(nil)).Elem()
}

type resolverEndpointArgs struct {
	Direction             string                             `pulumi:"direction"`
	IpAddresses           []ResolverEndpointIpAddressRequest `pulumi:"ipAddresses"`
	Name                  *string                            `pulumi:"name"`
	OutpostArn            *string                            `pulumi:"outpostArn"`
	PreferredInstanceType *string                            `pulumi:"preferredInstanceType"`
	Protocols             []string                           `pulumi:"protocols"`
	ResolverEndpointType  *string                            `pulumi:"resolverEndpointType"`
	SecurityGroupIds      []string                           `pulumi:"securityGroupIds"`
	Tags                  []ResolverEndpointTag              `pulumi:"tags"`
}

// The set of arguments for constructing a ResolverEndpoint resource.
type ResolverEndpointArgs struct {
	Direction             pulumi.StringInput
	IpAddresses           ResolverEndpointIpAddressRequestArrayInput
	Name                  pulumi.StringPtrInput
	OutpostArn            pulumi.StringPtrInput
	PreferredInstanceType pulumi.StringPtrInput
	Protocols             pulumi.StringArrayInput
	ResolverEndpointType  pulumi.StringPtrInput
	SecurityGroupIds      pulumi.StringArrayInput
	Tags                  ResolverEndpointTagArrayInput
}

func (ResolverEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverEndpointArgs)(nil)).Elem()
}

type ResolverEndpointInput interface {
	pulumi.Input

	ToResolverEndpointOutput() ResolverEndpointOutput
	ToResolverEndpointOutputWithContext(ctx context.Context) ResolverEndpointOutput
}

func (*ResolverEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverEndpoint)(nil)).Elem()
}

func (i *ResolverEndpoint) ToResolverEndpointOutput() ResolverEndpointOutput {
	return i.ToResolverEndpointOutputWithContext(context.Background())
}

func (i *ResolverEndpoint) ToResolverEndpointOutputWithContext(ctx context.Context) ResolverEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverEndpointOutput)
}

type ResolverEndpointOutput struct{ *pulumi.OutputState }

func (ResolverEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverEndpoint)(nil)).Elem()
}

func (o ResolverEndpointOutput) ToResolverEndpointOutput() ResolverEndpointOutput {
	return o
}

func (o ResolverEndpointOutput) ToResolverEndpointOutputWithContext(ctx context.Context) ResolverEndpointOutput {
	return o
}

func (o ResolverEndpointOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverEndpoint) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ResolverEndpointOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverEndpoint) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

func (o ResolverEndpointOutput) HostVpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverEndpoint) pulumi.StringOutput { return v.HostVpcId }).(pulumi.StringOutput)
}

func (o ResolverEndpointOutput) IpAddressCount() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverEndpoint) pulumi.StringOutput { return v.IpAddressCount }).(pulumi.StringOutput)
}

func (o ResolverEndpointOutput) IpAddresses() ResolverEndpointIpAddressRequestArrayOutput {
	return o.ApplyT(func(v *ResolverEndpoint) ResolverEndpointIpAddressRequestArrayOutput { return v.IpAddresses }).(ResolverEndpointIpAddressRequestArrayOutput)
}

func (o ResolverEndpointOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResolverEndpoint) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResolverEndpointOutput) OutpostArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResolverEndpoint) pulumi.StringPtrOutput { return v.OutpostArn }).(pulumi.StringPtrOutput)
}

func (o ResolverEndpointOutput) PreferredInstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResolverEndpoint) pulumi.StringPtrOutput { return v.PreferredInstanceType }).(pulumi.StringPtrOutput)
}

func (o ResolverEndpointOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResolverEndpoint) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

func (o ResolverEndpointOutput) ResolverEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverEndpoint) pulumi.StringOutput { return v.ResolverEndpointId }).(pulumi.StringOutput)
}

func (o ResolverEndpointOutput) ResolverEndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResolverEndpoint) pulumi.StringPtrOutput { return v.ResolverEndpointType }).(pulumi.StringPtrOutput)
}

func (o ResolverEndpointOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResolverEndpoint) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o ResolverEndpointOutput) Tags() ResolverEndpointTagArrayOutput {
	return o.ApplyT(func(v *ResolverEndpoint) ResolverEndpointTagArrayOutput { return v.Tags }).(ResolverEndpointTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverEndpointInput)(nil)).Elem(), &ResolverEndpoint{})
	pulumi.RegisterOutputType(ResolverEndpointOutput{})
}
