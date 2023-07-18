// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ServiceCatalog::PortfolioPrincipalAssociation
//
// Deprecated: PortfolioPrincipalAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type PortfolioPrincipalAssociation struct {
	pulumi.CustomResourceState

	AcceptLanguage pulumi.StringPtrOutput `pulumi:"acceptLanguage"`
	PortfolioId    pulumi.StringOutput    `pulumi:"portfolioId"`
	PrincipalARN   pulumi.StringOutput    `pulumi:"principalARN"`
	PrincipalType  pulumi.StringOutput    `pulumi:"principalType"`
}

// NewPortfolioPrincipalAssociation registers a new resource with the given unique name, arguments, and options.
func NewPortfolioPrincipalAssociation(ctx *pulumi.Context,
	name string, args *PortfolioPrincipalAssociationArgs, opts ...pulumi.ResourceOption) (*PortfolioPrincipalAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PortfolioId == nil {
		return nil, errors.New("invalid value for required argument 'PortfolioId'")
	}
	if args.PrincipalARN == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalARN'")
	}
	if args.PrincipalType == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PortfolioPrincipalAssociation
	err := ctx.RegisterResource("aws-native:servicecatalog:PortfolioPrincipalAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPortfolioPrincipalAssociation gets an existing PortfolioPrincipalAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPortfolioPrincipalAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortfolioPrincipalAssociationState, opts ...pulumi.ResourceOption) (*PortfolioPrincipalAssociation, error) {
	var resource PortfolioPrincipalAssociation
	err := ctx.ReadResource("aws-native:servicecatalog:PortfolioPrincipalAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PortfolioPrincipalAssociation resources.
type portfolioPrincipalAssociationState struct {
}

type PortfolioPrincipalAssociationState struct {
}

func (PortfolioPrincipalAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*portfolioPrincipalAssociationState)(nil)).Elem()
}

type portfolioPrincipalAssociationArgs struct {
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	PortfolioId    string  `pulumi:"portfolioId"`
	PrincipalARN   string  `pulumi:"principalARN"`
	PrincipalType  string  `pulumi:"principalType"`
}

// The set of arguments for constructing a PortfolioPrincipalAssociation resource.
type PortfolioPrincipalAssociationArgs struct {
	AcceptLanguage pulumi.StringPtrInput
	PortfolioId    pulumi.StringInput
	PrincipalARN   pulumi.StringInput
	PrincipalType  pulumi.StringInput
}

func (PortfolioPrincipalAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portfolioPrincipalAssociationArgs)(nil)).Elem()
}

type PortfolioPrincipalAssociationInput interface {
	pulumi.Input

	ToPortfolioPrincipalAssociationOutput() PortfolioPrincipalAssociationOutput
	ToPortfolioPrincipalAssociationOutputWithContext(ctx context.Context) PortfolioPrincipalAssociationOutput
}

func (*PortfolioPrincipalAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**PortfolioPrincipalAssociation)(nil)).Elem()
}

func (i *PortfolioPrincipalAssociation) ToPortfolioPrincipalAssociationOutput() PortfolioPrincipalAssociationOutput {
	return i.ToPortfolioPrincipalAssociationOutputWithContext(context.Background())
}

func (i *PortfolioPrincipalAssociation) ToPortfolioPrincipalAssociationOutputWithContext(ctx context.Context) PortfolioPrincipalAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortfolioPrincipalAssociationOutput)
}

type PortfolioPrincipalAssociationOutput struct{ *pulumi.OutputState }

func (PortfolioPrincipalAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PortfolioPrincipalAssociation)(nil)).Elem()
}

func (o PortfolioPrincipalAssociationOutput) ToPortfolioPrincipalAssociationOutput() PortfolioPrincipalAssociationOutput {
	return o
}

func (o PortfolioPrincipalAssociationOutput) ToPortfolioPrincipalAssociationOutputWithContext(ctx context.Context) PortfolioPrincipalAssociationOutput {
	return o
}

func (o PortfolioPrincipalAssociationOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PortfolioPrincipalAssociation) pulumi.StringPtrOutput { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

func (o PortfolioPrincipalAssociationOutput) PortfolioId() pulumi.StringOutput {
	return o.ApplyT(func(v *PortfolioPrincipalAssociation) pulumi.StringOutput { return v.PortfolioId }).(pulumi.StringOutput)
}

func (o PortfolioPrincipalAssociationOutput) PrincipalARN() pulumi.StringOutput {
	return o.ApplyT(func(v *PortfolioPrincipalAssociation) pulumi.StringOutput { return v.PrincipalARN }).(pulumi.StringOutput)
}

func (o PortfolioPrincipalAssociationOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *PortfolioPrincipalAssociation) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PortfolioPrincipalAssociationInput)(nil)).Elem(), &PortfolioPrincipalAssociation{})
	pulumi.RegisterOutputType(PortfolioPrincipalAssociationOutput{})
}
