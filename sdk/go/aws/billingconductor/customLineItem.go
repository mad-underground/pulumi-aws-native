// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package billingconductor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A custom line item is an one time charge that is applied to a specific billing group's bill.
//
// Deprecated: CustomLineItem is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type CustomLineItem struct {
	pulumi.CustomResourceState

	// ARN
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Number of source values associated to this custom line item
	AssociationSize pulumi.IntOutput `pulumi:"associationSize"`
	// Billing Group ARN
	BillingGroupArn    pulumi.StringOutput                       `pulumi:"billingGroupArn"`
	BillingPeriodRange CustomLineItemBillingPeriodRangePtrOutput `pulumi:"billingPeriodRange"`
	// Creation timestamp in UNIX epoch time format
	CreationTime                pulumi.IntOutput                     `pulumi:"creationTime"`
	CurrencyCode                CustomLineItemCurrencyCodeOutput     `pulumi:"currencyCode"`
	CustomLineItemChargeDetails CustomLineItemChargeDetailsPtrOutput `pulumi:"customLineItemChargeDetails"`
	Description                 pulumi.StringPtrOutput               `pulumi:"description"`
	// Latest modified timestamp in UNIX epoch time format
	LastModifiedTime pulumi.IntOutput    `pulumi:"lastModifiedTime"`
	Name             pulumi.StringOutput `pulumi:"name"`
	ProductCode      pulumi.StringOutput `pulumi:"productCode"`
}

// NewCustomLineItem registers a new resource with the given unique name, arguments, and options.
func NewCustomLineItem(ctx *pulumi.Context,
	name string, args *CustomLineItemArgs, opts ...pulumi.ResourceOption) (*CustomLineItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingGroupArn == nil {
		return nil, errors.New("invalid value for required argument 'BillingGroupArn'")
	}
	var resource CustomLineItem
	err := ctx.RegisterResource("aws-native:billingconductor:CustomLineItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomLineItem gets an existing CustomLineItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomLineItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomLineItemState, opts ...pulumi.ResourceOption) (*CustomLineItem, error) {
	var resource CustomLineItem
	err := ctx.ReadResource("aws-native:billingconductor:CustomLineItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomLineItem resources.
type customLineItemState struct {
}

type CustomLineItemState struct {
}

func (CustomLineItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*customLineItemState)(nil)).Elem()
}

type customLineItemArgs struct {
	// Billing Group ARN
	BillingGroupArn             string                            `pulumi:"billingGroupArn"`
	BillingPeriodRange          *CustomLineItemBillingPeriodRange `pulumi:"billingPeriodRange"`
	CustomLineItemChargeDetails *CustomLineItemChargeDetails      `pulumi:"customLineItemChargeDetails"`
	Description                 *string                           `pulumi:"description"`
	Name                        *string                           `pulumi:"name"`
}

// The set of arguments for constructing a CustomLineItem resource.
type CustomLineItemArgs struct {
	// Billing Group ARN
	BillingGroupArn             pulumi.StringInput
	BillingPeriodRange          CustomLineItemBillingPeriodRangePtrInput
	CustomLineItemChargeDetails CustomLineItemChargeDetailsPtrInput
	Description                 pulumi.StringPtrInput
	Name                        pulumi.StringPtrInput
}

func (CustomLineItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customLineItemArgs)(nil)).Elem()
}

type CustomLineItemInput interface {
	pulumi.Input

	ToCustomLineItemOutput() CustomLineItemOutput
	ToCustomLineItemOutputWithContext(ctx context.Context) CustomLineItemOutput
}

func (*CustomLineItem) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLineItem)(nil)).Elem()
}

func (i *CustomLineItem) ToCustomLineItemOutput() CustomLineItemOutput {
	return i.ToCustomLineItemOutputWithContext(context.Background())
}

func (i *CustomLineItem) ToCustomLineItemOutputWithContext(ctx context.Context) CustomLineItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLineItemOutput)
}

type CustomLineItemOutput struct{ *pulumi.OutputState }

func (CustomLineItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLineItem)(nil)).Elem()
}

func (o CustomLineItemOutput) ToCustomLineItemOutput() CustomLineItemOutput {
	return o
}

func (o CustomLineItemOutput) ToCustomLineItemOutputWithContext(ctx context.Context) CustomLineItemOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomLineItemInput)(nil)).Elem(), &CustomLineItem{})
	pulumi.RegisterOutputType(CustomLineItemOutput{})
}
