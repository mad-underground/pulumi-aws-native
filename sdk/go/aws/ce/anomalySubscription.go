// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ce

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. Create subscription to be notified
//
// Deprecated: AnomalySubscription is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type AnomalySubscription struct {
	pulumi.CustomResourceState

	// The accountId
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The frequency at which anomaly reports are sent over email.
	Frequency AnomalySubscriptionFrequencyOutput `pulumi:"frequency"`
	// A list of cost anomaly monitors.
	MonitorArnList pulumi.StringArrayOutput `pulumi:"monitorArnList"`
	// Tags to assign to subscription.
	ResourceTags AnomalySubscriptionResourceTagArrayOutput `pulumi:"resourceTags"`
	// A list of subscriber
	Subscribers     AnomalySubscriptionSubscriberArrayOutput `pulumi:"subscribers"`
	SubscriptionArn pulumi.StringOutput                      `pulumi:"subscriptionArn"`
	// The name of the subscription.
	SubscriptionName pulumi.StringOutput `pulumi:"subscriptionName"`
	// The dollar value that triggers a notification if the threshold is exceeded.
	Threshold pulumi.Float64PtrOutput `pulumi:"threshold"`
	// An Expression object in JSON String format used to specify the anomalies that you want to generate alerts for.
	ThresholdExpression pulumi.StringPtrOutput `pulumi:"thresholdExpression"`
}

// NewAnomalySubscription registers a new resource with the given unique name, arguments, and options.
func NewAnomalySubscription(ctx *pulumi.Context,
	name string, args *AnomalySubscriptionArgs, opts ...pulumi.ResourceOption) (*AnomalySubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.MonitorArnList == nil {
		return nil, errors.New("invalid value for required argument 'MonitorArnList'")
	}
	if args.Subscribers == nil {
		return nil, errors.New("invalid value for required argument 'Subscribers'")
	}
	if args.SubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AnomalySubscription
	err := ctx.RegisterResource("aws-native:ce:AnomalySubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnomalySubscription gets an existing AnomalySubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnomalySubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnomalySubscriptionState, opts ...pulumi.ResourceOption) (*AnomalySubscription, error) {
	var resource AnomalySubscription
	err := ctx.ReadResource("aws-native:ce:AnomalySubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnomalySubscription resources.
type anomalySubscriptionState struct {
}

type AnomalySubscriptionState struct {
}

func (AnomalySubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*anomalySubscriptionState)(nil)).Elem()
}

type anomalySubscriptionArgs struct {
	// The frequency at which anomaly reports are sent over email.
	Frequency AnomalySubscriptionFrequency `pulumi:"frequency"`
	// A list of cost anomaly monitors.
	MonitorArnList []string `pulumi:"monitorArnList"`
	// Tags to assign to subscription.
	ResourceTags []AnomalySubscriptionResourceTag `pulumi:"resourceTags"`
	// A list of subscriber
	Subscribers []AnomalySubscriptionSubscriber `pulumi:"subscribers"`
	// The name of the subscription.
	SubscriptionName string `pulumi:"subscriptionName"`
	// The dollar value that triggers a notification if the threshold is exceeded.
	Threshold *float64 `pulumi:"threshold"`
	// An Expression object in JSON String format used to specify the anomalies that you want to generate alerts for.
	ThresholdExpression *string `pulumi:"thresholdExpression"`
}

// The set of arguments for constructing a AnomalySubscription resource.
type AnomalySubscriptionArgs struct {
	// The frequency at which anomaly reports are sent over email.
	Frequency AnomalySubscriptionFrequencyInput
	// A list of cost anomaly monitors.
	MonitorArnList pulumi.StringArrayInput
	// Tags to assign to subscription.
	ResourceTags AnomalySubscriptionResourceTagArrayInput
	// A list of subscriber
	Subscribers AnomalySubscriptionSubscriberArrayInput
	// The name of the subscription.
	SubscriptionName pulumi.StringInput
	// The dollar value that triggers a notification if the threshold is exceeded.
	Threshold pulumi.Float64PtrInput
	// An Expression object in JSON String format used to specify the anomalies that you want to generate alerts for.
	ThresholdExpression pulumi.StringPtrInput
}

func (AnomalySubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*anomalySubscriptionArgs)(nil)).Elem()
}

type AnomalySubscriptionInput interface {
	pulumi.Input

	ToAnomalySubscriptionOutput() AnomalySubscriptionOutput
	ToAnomalySubscriptionOutputWithContext(ctx context.Context) AnomalySubscriptionOutput
}

func (*AnomalySubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**AnomalySubscription)(nil)).Elem()
}

func (i *AnomalySubscription) ToAnomalySubscriptionOutput() AnomalySubscriptionOutput {
	return i.ToAnomalySubscriptionOutputWithContext(context.Background())
}

func (i *AnomalySubscription) ToAnomalySubscriptionOutputWithContext(ctx context.Context) AnomalySubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnomalySubscriptionOutput)
}

type AnomalySubscriptionOutput struct{ *pulumi.OutputState }

func (AnomalySubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnomalySubscription)(nil)).Elem()
}

func (o AnomalySubscriptionOutput) ToAnomalySubscriptionOutput() AnomalySubscriptionOutput {
	return o
}

func (o AnomalySubscriptionOutput) ToAnomalySubscriptionOutputWithContext(ctx context.Context) AnomalySubscriptionOutput {
	return o
}

// The accountId
func (o AnomalySubscriptionOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySubscription) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The frequency at which anomaly reports are sent over email.
func (o AnomalySubscriptionOutput) Frequency() AnomalySubscriptionFrequencyOutput {
	return o.ApplyT(func(v *AnomalySubscription) AnomalySubscriptionFrequencyOutput { return v.Frequency }).(AnomalySubscriptionFrequencyOutput)
}

// A list of cost anomaly monitors.
func (o AnomalySubscriptionOutput) MonitorArnList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AnomalySubscription) pulumi.StringArrayOutput { return v.MonitorArnList }).(pulumi.StringArrayOutput)
}

// Tags to assign to subscription.
func (o AnomalySubscriptionOutput) ResourceTags() AnomalySubscriptionResourceTagArrayOutput {
	return o.ApplyT(func(v *AnomalySubscription) AnomalySubscriptionResourceTagArrayOutput { return v.ResourceTags }).(AnomalySubscriptionResourceTagArrayOutput)
}

// A list of subscriber
func (o AnomalySubscriptionOutput) Subscribers() AnomalySubscriptionSubscriberArrayOutput {
	return o.ApplyT(func(v *AnomalySubscription) AnomalySubscriptionSubscriberArrayOutput { return v.Subscribers }).(AnomalySubscriptionSubscriberArrayOutput)
}

func (o AnomalySubscriptionOutput) SubscriptionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySubscription) pulumi.StringOutput { return v.SubscriptionArn }).(pulumi.StringOutput)
}

// The name of the subscription.
func (o AnomalySubscriptionOutput) SubscriptionName() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySubscription) pulumi.StringOutput { return v.SubscriptionName }).(pulumi.StringOutput)
}

// The dollar value that triggers a notification if the threshold is exceeded.
func (o AnomalySubscriptionOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AnomalySubscription) pulumi.Float64PtrOutput { return v.Threshold }).(pulumi.Float64PtrOutput)
}

// An Expression object in JSON String format used to specify the anomalies that you want to generate alerts for.
func (o AnomalySubscriptionOutput) ThresholdExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnomalySubscription) pulumi.StringPtrOutput { return v.ThresholdExpression }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnomalySubscriptionInput)(nil)).Elem(), &AnomalySubscription{})
	pulumi.RegisterOutputType(AnomalySubscriptionOutput{})
}
