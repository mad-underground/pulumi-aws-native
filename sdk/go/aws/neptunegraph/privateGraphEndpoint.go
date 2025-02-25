// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package neptunegraph

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::NeptuneGraph::PrivateGraphEndpoint resource creates an Amazon NeptuneGraph PrivateGraphEndpoint.
type PrivateGraphEndpoint struct {
	pulumi.CustomResourceState

	// The auto-generated Graph Id assigned by the service.
	GraphIdentifier pulumi.StringOutput `pulumi:"graphIdentifier"`
	// PrivateGraphEndpoint resource identifier generated by concatenating the associated GraphIdentifier and VpcId with an underscore separator.
	//
	//  For example, if GraphIdentifier is `g-12a3bcdef4` and VpcId is `vpc-0a12bc34567de8f90`, the generated PrivateGraphEndpointIdentifier will be `g-12a3bcdef4_vpc-0a12bc34567de8f90`
	PrivateGraphEndpointIdentifier pulumi.StringOutput `pulumi:"privateGraphEndpointIdentifier"`
	// The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// VPC endpoint that provides a private connection between the Graph and specified VPC.
	VpcEndpointId pulumi.StringOutput `pulumi:"vpcEndpointId"`
	// The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewPrivateGraphEndpoint registers a new resource with the given unique name, arguments, and options.
func NewPrivateGraphEndpoint(ctx *pulumi.Context,
	name string, args *PrivateGraphEndpointArgs, opts ...pulumi.ResourceOption) (*PrivateGraphEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GraphIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'GraphIdentifier'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"graphIdentifier",
		"securityGroupIds[*]",
		"subnetIds[*]",
		"vpcId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrivateGraphEndpoint
	err := ctx.RegisterResource("aws-native:neptunegraph:PrivateGraphEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateGraphEndpoint gets an existing PrivateGraphEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateGraphEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateGraphEndpointState, opts ...pulumi.ResourceOption) (*PrivateGraphEndpoint, error) {
	var resource PrivateGraphEndpoint
	err := ctx.ReadResource("aws-native:neptunegraph:PrivateGraphEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateGraphEndpoint resources.
type privateGraphEndpointState struct {
}

type PrivateGraphEndpointState struct {
}

func (PrivateGraphEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateGraphEndpointState)(nil)).Elem()
}

type privateGraphEndpointArgs struct {
	// The auto-generated Graph Id assigned by the service.
	GraphIdentifier string `pulumi:"graphIdentifier"`
	// The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
	SubnetIds []string `pulumi:"subnetIds"`
	// The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a PrivateGraphEndpoint resource.
type PrivateGraphEndpointArgs struct {
	// The auto-generated Graph Id assigned by the service.
	GraphIdentifier pulumi.StringInput
	// The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
	SecurityGroupIds pulumi.StringArrayInput
	// The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
	SubnetIds pulumi.StringArrayInput
	// The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
	VpcId pulumi.StringInput
}

func (PrivateGraphEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateGraphEndpointArgs)(nil)).Elem()
}

type PrivateGraphEndpointInput interface {
	pulumi.Input

	ToPrivateGraphEndpointOutput() PrivateGraphEndpointOutput
	ToPrivateGraphEndpointOutputWithContext(ctx context.Context) PrivateGraphEndpointOutput
}

func (*PrivateGraphEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateGraphEndpoint)(nil)).Elem()
}

func (i *PrivateGraphEndpoint) ToPrivateGraphEndpointOutput() PrivateGraphEndpointOutput {
	return i.ToPrivateGraphEndpointOutputWithContext(context.Background())
}

func (i *PrivateGraphEndpoint) ToPrivateGraphEndpointOutputWithContext(ctx context.Context) PrivateGraphEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateGraphEndpointOutput)
}

type PrivateGraphEndpointOutput struct{ *pulumi.OutputState }

func (PrivateGraphEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateGraphEndpoint)(nil)).Elem()
}

func (o PrivateGraphEndpointOutput) ToPrivateGraphEndpointOutput() PrivateGraphEndpointOutput {
	return o
}

func (o PrivateGraphEndpointOutput) ToPrivateGraphEndpointOutputWithContext(ctx context.Context) PrivateGraphEndpointOutput {
	return o
}

// The auto-generated Graph Id assigned by the service.
func (o PrivateGraphEndpointOutput) GraphIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateGraphEndpoint) pulumi.StringOutput { return v.GraphIdentifier }).(pulumi.StringOutput)
}

// PrivateGraphEndpoint resource identifier generated by concatenating the associated GraphIdentifier and VpcId with an underscore separator.
//
//	For example, if GraphIdentifier is `g-12a3bcdef4` and VpcId is `vpc-0a12bc34567de8f90`, the generated PrivateGraphEndpointIdentifier will be `g-12a3bcdef4_vpc-0a12bc34567de8f90`
func (o PrivateGraphEndpointOutput) PrivateGraphEndpointIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateGraphEndpoint) pulumi.StringOutput { return v.PrivateGraphEndpointIdentifier }).(pulumi.StringOutput)
}

// The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
func (o PrivateGraphEndpointOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateGraphEndpoint) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
func (o PrivateGraphEndpointOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateGraphEndpoint) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// VPC endpoint that provides a private connection between the Graph and specified VPC.
func (o PrivateGraphEndpointOutput) VpcEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateGraphEndpoint) pulumi.StringOutput { return v.VpcEndpointId }).(pulumi.StringOutput)
}

// The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
func (o PrivateGraphEndpointOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateGraphEndpoint) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateGraphEndpointInput)(nil)).Elem(), &PrivateGraphEndpoint{})
	pulumi.RegisterOutputType(PrivateGraphEndpointOutput{})
}
