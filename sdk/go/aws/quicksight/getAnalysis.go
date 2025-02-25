// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the AWS::QuickSight::Analysis Resource Type.
func LookupAnalysis(ctx *pulumi.Context, args *LookupAnalysisArgs, opts ...pulumi.InvokeOption) (*LookupAnalysisResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAnalysisResult
	err := ctx.Invoke("aws-native:quicksight:getAnalysis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAnalysisArgs struct {
	AnalysisId   string `pulumi:"analysisId"`
	AwsAccountId string `pulumi:"awsAccountId"`
}

type LookupAnalysisResult struct {
	Arn             *string                      `pulumi:"arn"`
	CreatedTime     *string                      `pulumi:"createdTime"`
	DataSetArns     []string                     `pulumi:"dataSetArns"`
	Errors          []AnalysisError              `pulumi:"errors"`
	LastUpdatedTime *string                      `pulumi:"lastUpdatedTime"`
	Name            *string                      `pulumi:"name"`
	Permissions     []AnalysisResourcePermission `pulumi:"permissions"`
	Sheets          []AnalysisSheet              `pulumi:"sheets"`
	Tags            []aws.Tag                    `pulumi:"tags"`
	ThemeArn        *string                      `pulumi:"themeArn"`
}

func LookupAnalysisOutput(ctx *pulumi.Context, args LookupAnalysisOutputArgs, opts ...pulumi.InvokeOption) LookupAnalysisResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAnalysisResult, error) {
			args := v.(LookupAnalysisArgs)
			r, err := LookupAnalysis(ctx, &args, opts...)
			var s LookupAnalysisResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAnalysisResultOutput)
}

type LookupAnalysisOutputArgs struct {
	AnalysisId   pulumi.StringInput `pulumi:"analysisId"`
	AwsAccountId pulumi.StringInput `pulumi:"awsAccountId"`
}

func (LookupAnalysisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnalysisArgs)(nil)).Elem()
}

type LookupAnalysisResultOutput struct{ *pulumi.OutputState }

func (LookupAnalysisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnalysisResult)(nil)).Elem()
}

func (o LookupAnalysisResultOutput) ToLookupAnalysisResultOutput() LookupAnalysisResultOutput {
	return o
}

func (o LookupAnalysisResultOutput) ToLookupAnalysisResultOutputWithContext(ctx context.Context) LookupAnalysisResultOutput {
	return o
}

func (o LookupAnalysisResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalysisResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupAnalysisResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalysisResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o LookupAnalysisResultOutput) DataSetArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAnalysisResult) []string { return v.DataSetArns }).(pulumi.StringArrayOutput)
}

func (o LookupAnalysisResultOutput) Errors() AnalysisErrorArrayOutput {
	return o.ApplyT(func(v LookupAnalysisResult) []AnalysisError { return v.Errors }).(AnalysisErrorArrayOutput)
}

func (o LookupAnalysisResultOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalysisResult) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

func (o LookupAnalysisResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalysisResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupAnalysisResultOutput) Permissions() AnalysisResourcePermissionArrayOutput {
	return o.ApplyT(func(v LookupAnalysisResult) []AnalysisResourcePermission { return v.Permissions }).(AnalysisResourcePermissionArrayOutput)
}

func (o LookupAnalysisResultOutput) Sheets() AnalysisSheetArrayOutput {
	return o.ApplyT(func(v LookupAnalysisResult) []AnalysisSheet { return v.Sheets }).(AnalysisSheetArrayOutput)
}

func (o LookupAnalysisResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupAnalysisResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func (o LookupAnalysisResultOutput) ThemeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalysisResult) *string { return v.ThemeArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAnalysisResultOutput{})
}
