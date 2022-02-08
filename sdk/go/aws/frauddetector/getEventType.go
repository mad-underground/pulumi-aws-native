// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package frauddetector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A resource schema for an EventType in Amazon Fraud Detector.
func LookupEventType(ctx *pulumi.Context, args *LookupEventTypeArgs, opts ...pulumi.InvokeOption) (*LookupEventTypeResult, error) {
	var rv LookupEventTypeResult
	err := ctx.Invoke("aws-native:frauddetector:getEventType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventTypeArgs struct {
	// The ARN of the event type.
	Arn string `pulumi:"arn"`
}

type LookupEventTypeResult struct {
	// The ARN of the event type.
	Arn *string `pulumi:"arn"`
	// The time when the event type was created.
	CreatedTime *string `pulumi:"createdTime"`
	// The description of the event type.
	Description    *string                  `pulumi:"description"`
	EntityTypes    []EventTypeEntityType    `pulumi:"entityTypes"`
	EventVariables []EventTypeEventVariable `pulumi:"eventVariables"`
	Labels         []EventTypeLabel         `pulumi:"labels"`
	// The time when the event type was last updated.
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// Tags associated with this event type.
	Tags []EventTypeTag `pulumi:"tags"`
}

func LookupEventTypeOutput(ctx *pulumi.Context, args LookupEventTypeOutputArgs, opts ...pulumi.InvokeOption) LookupEventTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventTypeResult, error) {
			args := v.(LookupEventTypeArgs)
			r, err := LookupEventType(ctx, &args, opts...)
			return *r, err
		}).(LookupEventTypeResultOutput)
}

type LookupEventTypeOutputArgs struct {
	// The ARN of the event type.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupEventTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventTypeArgs)(nil)).Elem()
}

type LookupEventTypeResultOutput struct{ *pulumi.OutputState }

func (LookupEventTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventTypeResult)(nil)).Elem()
}

func (o LookupEventTypeResultOutput) ToLookupEventTypeResultOutput() LookupEventTypeResultOutput {
	return o
}

func (o LookupEventTypeResultOutput) ToLookupEventTypeResultOutputWithContext(ctx context.Context) LookupEventTypeResultOutput {
	return o
}

// The ARN of the event type.
func (o LookupEventTypeResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventTypeResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The time when the event type was created.
func (o LookupEventTypeResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventTypeResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

// The description of the event type.
func (o LookupEventTypeResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventTypeResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupEventTypeResultOutput) EntityTypes() EventTypeEntityTypeArrayOutput {
	return o.ApplyT(func(v LookupEventTypeResult) []EventTypeEntityType { return v.EntityTypes }).(EventTypeEntityTypeArrayOutput)
}

func (o LookupEventTypeResultOutput) EventVariables() EventTypeEventVariableArrayOutput {
	return o.ApplyT(func(v LookupEventTypeResult) []EventTypeEventVariable { return v.EventVariables }).(EventTypeEventVariableArrayOutput)
}

func (o LookupEventTypeResultOutput) Labels() EventTypeLabelArrayOutput {
	return o.ApplyT(func(v LookupEventTypeResult) []EventTypeLabel { return v.Labels }).(EventTypeLabelArrayOutput)
}

// The time when the event type was last updated.
func (o LookupEventTypeResultOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventTypeResult) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

// Tags associated with this event type.
func (o LookupEventTypeResultOutput) Tags() EventTypeTagArrayOutput {
	return o.ApplyT(func(v LookupEventTypeResult) []EventTypeTag { return v.Tags }).(EventTypeTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventTypeResultOutput{})
}
