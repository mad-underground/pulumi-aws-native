// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Logs::LogGroup
func LookupLogGroup(ctx *pulumi.Context, args *LookupLogGroupArgs, opts ...pulumi.InvokeOption) (*LookupLogGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLogGroupResult
	err := ctx.Invoke("aws-native:logs:getLogGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLogGroupArgs struct {
	// The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.
	LogGroupName string `pulumi:"logGroupName"`
}

type LookupLogGroupResult struct {
	// The CloudWatch log group ARN.
	Arn *string `pulumi:"arn"`
	// The body of the policy document you want to use for this topic.
	//
	// You can only add one policy per topic.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 30720
	DataProtectionPolicy interface{} `pulumi:"dataProtectionPolicy"`
	// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The class of the log group. Possible values are: STANDARD and INFREQUENT_ACCESS, with STANDARD being the default class
	LogGroupClass *LogGroupClass `pulumi:"logGroupClass"`
	// The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, and 3653.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// An array of key-value pairs to apply to this resource.
	Tags []LogGroupTag `pulumi:"tags"`
}

func LookupLogGroupOutput(ctx *pulumi.Context, args LookupLogGroupOutputArgs, opts ...pulumi.InvokeOption) LookupLogGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLogGroupResult, error) {
			args := v.(LookupLogGroupArgs)
			r, err := LookupLogGroup(ctx, &args, opts...)
			var s LookupLogGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLogGroupResultOutput)
}

type LookupLogGroupOutputArgs struct {
	// The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.
	LogGroupName pulumi.StringInput `pulumi:"logGroupName"`
}

func (LookupLogGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogGroupArgs)(nil)).Elem()
}

type LookupLogGroupResultOutput struct{ *pulumi.OutputState }

func (LookupLogGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogGroupResult)(nil)).Elem()
}

func (o LookupLogGroupResultOutput) ToLookupLogGroupResultOutput() LookupLogGroupResultOutput {
	return o
}

func (o LookupLogGroupResultOutput) ToLookupLogGroupResultOutputWithContext(ctx context.Context) LookupLogGroupResultOutput {
	return o
}

// The CloudWatch log group ARN.
func (o LookupLogGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLogGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The body of the policy document you want to use for this topic.
//
// You can only add one policy per topic.
//
// The policy must be in JSON string format.
//
// Length Constraints: Maximum length of 30720
func (o LookupLogGroupResultOutput) DataProtectionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupLogGroupResult) interface{} { return v.DataProtectionPolicy }).(pulumi.AnyOutput)
}

// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
func (o LookupLogGroupResultOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLogGroupResult) *string { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// The class of the log group. Possible values are: STANDARD and INFREQUENT_ACCESS, with STANDARD being the default class
func (o LookupLogGroupResultOutput) LogGroupClass() LogGroupClassPtrOutput {
	return o.ApplyT(func(v LookupLogGroupResult) *LogGroupClass { return v.LogGroupClass }).(LogGroupClassPtrOutput)
}

// The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, and 3653.
func (o LookupLogGroupResultOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLogGroupResult) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupLogGroupResultOutput) Tags() LogGroupTagArrayOutput {
	return o.ApplyT(func(v LookupLogGroupResult) []LogGroupTag { return v.Tags }).(LogGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLogGroupResultOutput{})
}
