// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::Logs::LogAnomalyDetector resource specifies a CloudWatch Logs LogAnomalyDetector.
func LookupLogAnomalyDetector(ctx *pulumi.Context, args *LookupLogAnomalyDetectorArgs, opts ...pulumi.InvokeOption) (*LookupLogAnomalyDetectorResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLogAnomalyDetectorResult
	err := ctx.Invoke("aws-native:logs:getLogAnomalyDetector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLogAnomalyDetectorArgs struct {
	// ARN of LogAnomalyDetector
	AnomalyDetectorArn string `pulumi:"anomalyDetectorArn"`
}

type LookupLogAnomalyDetectorResult struct {
	// ARN of LogAnomalyDetector
	AnomalyDetectorArn *string `pulumi:"anomalyDetectorArn"`
	// Current status of detector.
	AnomalyDetectorStatus *string  `pulumi:"anomalyDetectorStatus"`
	AnomalyVisibilityTime *float64 `pulumi:"anomalyVisibilityTime"`
	// When detector was created.
	CreationTimeStamp *float64 `pulumi:"creationTimeStamp"`
	// Name of detector
	DetectorName *string `pulumi:"detectorName"`
	// How often log group is evaluated
	EvaluationFrequency *LogAnomalyDetectorEvaluationFrequency `pulumi:"evaluationFrequency"`
	FilterPattern       *string                                `pulumi:"filterPattern"`
	// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// When detector was lsat modified.
	LastModifiedTimeStamp *float64 `pulumi:"lastModifiedTimeStamp"`
	// List of Arns for the given log group
	LogGroupArnList []string `pulumi:"logGroupArnList"`
}

func LookupLogAnomalyDetectorOutput(ctx *pulumi.Context, args LookupLogAnomalyDetectorOutputArgs, opts ...pulumi.InvokeOption) LookupLogAnomalyDetectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLogAnomalyDetectorResult, error) {
			args := v.(LookupLogAnomalyDetectorArgs)
			r, err := LookupLogAnomalyDetector(ctx, &args, opts...)
			var s LookupLogAnomalyDetectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLogAnomalyDetectorResultOutput)
}

type LookupLogAnomalyDetectorOutputArgs struct {
	// ARN of LogAnomalyDetector
	AnomalyDetectorArn pulumi.StringInput `pulumi:"anomalyDetectorArn"`
}

func (LookupLogAnomalyDetectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogAnomalyDetectorArgs)(nil)).Elem()
}

type LookupLogAnomalyDetectorResultOutput struct{ *pulumi.OutputState }

func (LookupLogAnomalyDetectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogAnomalyDetectorResult)(nil)).Elem()
}

func (o LookupLogAnomalyDetectorResultOutput) ToLookupLogAnomalyDetectorResultOutput() LookupLogAnomalyDetectorResultOutput {
	return o
}

func (o LookupLogAnomalyDetectorResultOutput) ToLookupLogAnomalyDetectorResultOutputWithContext(ctx context.Context) LookupLogAnomalyDetectorResultOutput {
	return o
}

// ARN of LogAnomalyDetector
func (o LookupLogAnomalyDetectorResultOutput) AnomalyDetectorArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLogAnomalyDetectorResult) *string { return v.AnomalyDetectorArn }).(pulumi.StringPtrOutput)
}

// Current status of detector.
func (o LookupLogAnomalyDetectorResultOutput) AnomalyDetectorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLogAnomalyDetectorResult) *string { return v.AnomalyDetectorStatus }).(pulumi.StringPtrOutput)
}

func (o LookupLogAnomalyDetectorResultOutput) AnomalyVisibilityTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupLogAnomalyDetectorResult) *float64 { return v.AnomalyVisibilityTime }).(pulumi.Float64PtrOutput)
}

// When detector was created.
func (o LookupLogAnomalyDetectorResultOutput) CreationTimeStamp() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupLogAnomalyDetectorResult) *float64 { return v.CreationTimeStamp }).(pulumi.Float64PtrOutput)
}

// Name of detector
func (o LookupLogAnomalyDetectorResultOutput) DetectorName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLogAnomalyDetectorResult) *string { return v.DetectorName }).(pulumi.StringPtrOutput)
}

// How often log group is evaluated
func (o LookupLogAnomalyDetectorResultOutput) EvaluationFrequency() LogAnomalyDetectorEvaluationFrequencyPtrOutput {
	return o.ApplyT(func(v LookupLogAnomalyDetectorResult) *LogAnomalyDetectorEvaluationFrequency {
		return v.EvaluationFrequency
	}).(LogAnomalyDetectorEvaluationFrequencyPtrOutput)
}

func (o LookupLogAnomalyDetectorResultOutput) FilterPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLogAnomalyDetectorResult) *string { return v.FilterPattern }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
func (o LookupLogAnomalyDetectorResultOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLogAnomalyDetectorResult) *string { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// When detector was lsat modified.
func (o LookupLogAnomalyDetectorResultOutput) LastModifiedTimeStamp() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupLogAnomalyDetectorResult) *float64 { return v.LastModifiedTimeStamp }).(pulumi.Float64PtrOutput)
}

// List of Arns for the given log group
func (o LookupLogAnomalyDetectorResultOutput) LogGroupArnList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLogAnomalyDetectorResult) []string { return v.LogGroupArnList }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLogAnomalyDetectorResultOutput{})
}
