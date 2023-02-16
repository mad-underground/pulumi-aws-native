// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DataSync::Task.
func LookupTask(ctx *pulumi.Context, args *LookupTaskArgs, opts ...pulumi.InvokeOption) (*LookupTaskResult, error) {
	var rv LookupTaskResult
	err := ctx.Invoke("aws-native:datasync:getTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTaskArgs struct {
	// The ARN of the task.
	TaskArn string `pulumi:"taskArn"`
}

type LookupTaskResult struct {
	// The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task.
	CloudWatchLogGroupArn           *string          `pulumi:"cloudWatchLogGroupArn"`
	DestinationNetworkInterfaceArns []string         `pulumi:"destinationNetworkInterfaceArns"`
	Excludes                        []TaskFilterRule `pulumi:"excludes"`
	Includes                        []TaskFilterRule `pulumi:"includes"`
	// The name of a task. This value is a text reference that is used to identify the task in the console.
	Name                       *string       `pulumi:"name"`
	Options                    *TaskOptions  `pulumi:"options"`
	Schedule                   *TaskSchedule `pulumi:"schedule"`
	SourceNetworkInterfaceArns []string      `pulumi:"sourceNetworkInterfaceArns"`
	// The status of the task that was described.
	Status *TaskStatus `pulumi:"status"`
	// An array of key-value pairs to apply to this resource.
	Tags []TaskTag `pulumi:"tags"`
	// The ARN of the task.
	TaskArn *string `pulumi:"taskArn"`
}

func LookupTaskOutput(ctx *pulumi.Context, args LookupTaskOutputArgs, opts ...pulumi.InvokeOption) LookupTaskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTaskResult, error) {
			args := v.(LookupTaskArgs)
			r, err := LookupTask(ctx, &args, opts...)
			var s LookupTaskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTaskResultOutput)
}

type LookupTaskOutputArgs struct {
	// The ARN of the task.
	TaskArn pulumi.StringInput `pulumi:"taskArn"`
}

func (LookupTaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskArgs)(nil)).Elem()
}

type LookupTaskResultOutput struct{ *pulumi.OutputState }

func (LookupTaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskResult)(nil)).Elem()
}

func (o LookupTaskResultOutput) ToLookupTaskResultOutput() LookupTaskResultOutput {
	return o
}

func (o LookupTaskResultOutput) ToLookupTaskResultOutputWithContext(ctx context.Context) LookupTaskResultOutput {
	return o
}

// The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task.
func (o LookupTaskResultOutput) CloudWatchLogGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *string { return v.CloudWatchLogGroupArn }).(pulumi.StringPtrOutput)
}

func (o LookupTaskResultOutput) DestinationNetworkInterfaceArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTaskResult) []string { return v.DestinationNetworkInterfaceArns }).(pulumi.StringArrayOutput)
}

func (o LookupTaskResultOutput) Excludes() TaskFilterRuleArrayOutput {
	return o.ApplyT(func(v LookupTaskResult) []TaskFilterRule { return v.Excludes }).(TaskFilterRuleArrayOutput)
}

func (o LookupTaskResultOutput) Includes() TaskFilterRuleArrayOutput {
	return o.ApplyT(func(v LookupTaskResult) []TaskFilterRule { return v.Includes }).(TaskFilterRuleArrayOutput)
}

// The name of a task. This value is a text reference that is used to identify the task in the console.
func (o LookupTaskResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupTaskResultOutput) Options() TaskOptionsPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *TaskOptions { return v.Options }).(TaskOptionsPtrOutput)
}

func (o LookupTaskResultOutput) Schedule() TaskSchedulePtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *TaskSchedule { return v.Schedule }).(TaskSchedulePtrOutput)
}

func (o LookupTaskResultOutput) SourceNetworkInterfaceArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTaskResult) []string { return v.SourceNetworkInterfaceArns }).(pulumi.StringArrayOutput)
}

// The status of the task that was described.
func (o LookupTaskResultOutput) Status() TaskStatusPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *TaskStatus { return v.Status }).(TaskStatusPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupTaskResultOutput) Tags() TaskTagArrayOutput {
	return o.ApplyT(func(v LookupTaskResult) []TaskTag { return v.Tags }).(TaskTagArrayOutput)
}

// The ARN of the task.
func (o LookupTaskResultOutput) TaskArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTaskResult) *string { return v.TaskArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTaskResultOutput{})
}
