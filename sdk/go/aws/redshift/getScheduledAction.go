// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `AWS::Redshift::ScheduledAction` resource creates an Amazon Redshift Scheduled Action.
func LookupScheduledAction(ctx *pulumi.Context, args *LookupScheduledActionArgs, opts ...pulumi.InvokeOption) (*LookupScheduledActionResult, error) {
	var rv LookupScheduledActionResult
	err := ctx.Invoke("aws-native:redshift:getScheduledAction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledActionArgs struct {
	// The name of the scheduled action. The name must be unique within an account.
	ScheduledActionName string `pulumi:"scheduledActionName"`
}

type LookupScheduledActionResult struct {
	// If true, the schedule is enabled. If false, the scheduled action does not trigger.
	Enable *bool `pulumi:"enable"`
	// The end time in UTC of the scheduled action. After this time, the scheduled action does not trigger.
	EndTime *string `pulumi:"endTime"`
	// The IAM role to assume to run the target action.
	IamRole *string `pulumi:"iamRole"`
	// List of times when the scheduled action will run.
	NextInvocations []string `pulumi:"nextInvocations"`
	// The schedule in `at( )` or `cron( )` format.
	Schedule *string `pulumi:"schedule"`
	// The description of the scheduled action.
	ScheduledActionDescription *string `pulumi:"scheduledActionDescription"`
	// The start time in UTC of the scheduled action. Before this time, the scheduled action does not trigger.
	StartTime *string `pulumi:"startTime"`
	// The state of the scheduled action.
	State *ScheduledActionStateEnum `pulumi:"state"`
	// A JSON format string of the Amazon Redshift API operation with input parameters.
	TargetAction *ScheduledActionType `pulumi:"targetAction"`
}

func LookupScheduledActionOutput(ctx *pulumi.Context, args LookupScheduledActionOutputArgs, opts ...pulumi.InvokeOption) LookupScheduledActionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduledActionResult, error) {
			args := v.(LookupScheduledActionArgs)
			r, err := LookupScheduledAction(ctx, &args, opts...)
			return *r, err
		}).(LookupScheduledActionResultOutput)
}

type LookupScheduledActionOutputArgs struct {
	// The name of the scheduled action. The name must be unique within an account.
	ScheduledActionName pulumi.StringInput `pulumi:"scheduledActionName"`
}

func (LookupScheduledActionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledActionArgs)(nil)).Elem()
}

type LookupScheduledActionResultOutput struct{ *pulumi.OutputState }

func (LookupScheduledActionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledActionResult)(nil)).Elem()
}

func (o LookupScheduledActionResultOutput) ToLookupScheduledActionResultOutput() LookupScheduledActionResultOutput {
	return o
}

func (o LookupScheduledActionResultOutput) ToLookupScheduledActionResultOutputWithContext(ctx context.Context) LookupScheduledActionResultOutput {
	return o
}

// If true, the schedule is enabled. If false, the scheduled action does not trigger.
func (o LookupScheduledActionResultOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

// The end time in UTC of the scheduled action. After this time, the scheduled action does not trigger.
func (o LookupScheduledActionResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// The IAM role to assume to run the target action.
func (o LookupScheduledActionResultOutput) IamRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.IamRole }).(pulumi.StringPtrOutput)
}

// List of times when the scheduled action will run.
func (o LookupScheduledActionResultOutput) NextInvocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) []string { return v.NextInvocations }).(pulumi.StringArrayOutput)
}

// The schedule in `at( )` or `cron( )` format.
func (o LookupScheduledActionResultOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.Schedule }).(pulumi.StringPtrOutput)
}

// The description of the scheduled action.
func (o LookupScheduledActionResultOutput) ScheduledActionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.ScheduledActionDescription }).(pulumi.StringPtrOutput)
}

// The start time in UTC of the scheduled action. Before this time, the scheduled action does not trigger.
func (o LookupScheduledActionResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

// The state of the scheduled action.
func (o LookupScheduledActionResultOutput) State() ScheduledActionStateEnumPtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *ScheduledActionStateEnum { return v.State }).(ScheduledActionStateEnumPtrOutput)
}

// A JSON format string of the Amazon Redshift API operation with input parameters.
func (o LookupScheduledActionResultOutput) TargetAction() ScheduledActionTypePtrOutput {
	return o.ApplyT(func(v LookupScheduledActionResult) *ScheduledActionType { return v.TargetAction }).(ScheduledActionTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduledActionResultOutput{})
}
