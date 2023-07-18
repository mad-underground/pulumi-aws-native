// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package budgets

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Budgets::Budget
//
// Deprecated: Budget is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Budget struct {
	pulumi.CustomResourceState

	Budget                       BudgetDataOutput                             `pulumi:"budget"`
	NotificationsWithSubscribers BudgetNotificationWithSubscribersArrayOutput `pulumi:"notificationsWithSubscribers"`
}

// NewBudget registers a new resource with the given unique name, arguments, and options.
func NewBudget(ctx *pulumi.Context,
	name string, args *BudgetArgs, opts ...pulumi.ResourceOption) (*Budget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Budget == nil {
		return nil, errors.New("invalid value for required argument 'Budget'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Budget
	err := ctx.RegisterResource("aws-native:budgets:Budget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBudget gets an existing Budget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBudget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BudgetState, opts ...pulumi.ResourceOption) (*Budget, error) {
	var resource Budget
	err := ctx.ReadResource("aws-native:budgets:Budget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Budget resources.
type budgetState struct {
}

type BudgetState struct {
}

func (BudgetState) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetState)(nil)).Elem()
}

type budgetArgs struct {
	Budget                       BudgetData                          `pulumi:"budget"`
	NotificationsWithSubscribers []BudgetNotificationWithSubscribers `pulumi:"notificationsWithSubscribers"`
}

// The set of arguments for constructing a Budget resource.
type BudgetArgs struct {
	Budget                       BudgetDataInput
	NotificationsWithSubscribers BudgetNotificationWithSubscribersArrayInput
}

func (BudgetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetArgs)(nil)).Elem()
}

type BudgetInput interface {
	pulumi.Input

	ToBudgetOutput() BudgetOutput
	ToBudgetOutputWithContext(ctx context.Context) BudgetOutput
}

func (*Budget) ElementType() reflect.Type {
	return reflect.TypeOf((**Budget)(nil)).Elem()
}

func (i *Budget) ToBudgetOutput() BudgetOutput {
	return i.ToBudgetOutputWithContext(context.Background())
}

func (i *Budget) ToBudgetOutputWithContext(ctx context.Context) BudgetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetOutput)
}

type BudgetOutput struct{ *pulumi.OutputState }

func (BudgetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Budget)(nil)).Elem()
}

func (o BudgetOutput) ToBudgetOutput() BudgetOutput {
	return o
}

func (o BudgetOutput) ToBudgetOutputWithContext(ctx context.Context) BudgetOutput {
	return o
}

func (o BudgetOutput) Budget() BudgetDataOutput {
	return o.ApplyT(func(v *Budget) BudgetDataOutput { return v.Budget }).(BudgetDataOutput)
}

func (o BudgetOutput) NotificationsWithSubscribers() BudgetNotificationWithSubscribersArrayOutput {
	return o.ApplyT(func(v *Budget) BudgetNotificationWithSubscribersArrayOutput { return v.NotificationsWithSubscribers }).(BudgetNotificationWithSubscribersArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetInput)(nil)).Elem(), &Budget{})
	pulumi.RegisterOutputType(BudgetOutput{})
}
