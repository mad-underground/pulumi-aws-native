// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::EC2::TransitGatewayMulticastDomainAssociation type
func LookupTransitGatewayMulticastDomainAssociation(ctx *pulumi.Context, args *LookupTransitGatewayMulticastDomainAssociationArgs, opts ...pulumi.InvokeOption) (*LookupTransitGatewayMulticastDomainAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTransitGatewayMulticastDomainAssociationResult
	err := ctx.Invoke("aws-native:ec2:getTransitGatewayMulticastDomainAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransitGatewayMulticastDomainAssociationArgs struct {
	// The IDs of the subnets to associate with the transit gateway multicast domain.
	SubnetId string `pulumi:"subnetId"`
	// The ID of the transit gateway attachment.
	TransitGatewayAttachmentId string `pulumi:"transitGatewayAttachmentId"`
	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId string `pulumi:"transitGatewayMulticastDomainId"`
}

type LookupTransitGatewayMulticastDomainAssociationResult struct {
	// The ID of the resource.
	ResourceId *string `pulumi:"resourceId"`
	// The type of resource, for example a VPC attachment.
	ResourceType *string `pulumi:"resourceType"`
	// The state of the subnet association.
	State *string `pulumi:"state"`
}

func LookupTransitGatewayMulticastDomainAssociationOutput(ctx *pulumi.Context, args LookupTransitGatewayMulticastDomainAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupTransitGatewayMulticastDomainAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTransitGatewayMulticastDomainAssociationResult, error) {
			args := v.(LookupTransitGatewayMulticastDomainAssociationArgs)
			r, err := LookupTransitGatewayMulticastDomainAssociation(ctx, &args, opts...)
			var s LookupTransitGatewayMulticastDomainAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTransitGatewayMulticastDomainAssociationResultOutput)
}

type LookupTransitGatewayMulticastDomainAssociationOutputArgs struct {
	// The IDs of the subnets to associate with the transit gateway multicast domain.
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
	// The ID of the transit gateway attachment.
	TransitGatewayAttachmentId pulumi.StringInput `pulumi:"transitGatewayAttachmentId"`
	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId pulumi.StringInput `pulumi:"transitGatewayMulticastDomainId"`
}

func (LookupTransitGatewayMulticastDomainAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransitGatewayMulticastDomainAssociationArgs)(nil)).Elem()
}

type LookupTransitGatewayMulticastDomainAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupTransitGatewayMulticastDomainAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransitGatewayMulticastDomainAssociationResult)(nil)).Elem()
}

func (o LookupTransitGatewayMulticastDomainAssociationResultOutput) ToLookupTransitGatewayMulticastDomainAssociationResultOutput() LookupTransitGatewayMulticastDomainAssociationResultOutput {
	return o
}

func (o LookupTransitGatewayMulticastDomainAssociationResultOutput) ToLookupTransitGatewayMulticastDomainAssociationResultOutputWithContext(ctx context.Context) LookupTransitGatewayMulticastDomainAssociationResultOutput {
	return o
}

// The ID of the resource.
func (o LookupTransitGatewayMulticastDomainAssociationResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransitGatewayMulticastDomainAssociationResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The type of resource, for example a VPC attachment.
func (o LookupTransitGatewayMulticastDomainAssociationResultOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransitGatewayMulticastDomainAssociationResult) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

// The state of the subnet association.
func (o LookupTransitGatewayMulticastDomainAssociationResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTransitGatewayMulticastDomainAssociationResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTransitGatewayMulticastDomainAssociationResultOutput{})
}
