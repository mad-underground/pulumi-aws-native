// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Job templates enable you to preconfigure jobs so that you can deploy them to multiple sets of target devices.
func LookupJobTemplate(ctx *pulumi.Context, args *LookupJobTemplateArgs, opts ...pulumi.InvokeOption) (*LookupJobTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupJobTemplateResult
	err := ctx.Invoke("aws-native:iot:getJobTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobTemplateArgs struct {
	JobTemplateId string `pulumi:"jobTemplateId"`
}

type LookupJobTemplateResult struct {
	Arn                      *string                             `pulumi:"arn"`
	JobExecutionsRetryConfig *JobExecutionsRetryConfigProperties `pulumi:"jobExecutionsRetryConfig"`
	MaintenanceWindows       []JobTemplateMaintenanceWindow      `pulumi:"maintenanceWindows"`
}

func LookupJobTemplateOutput(ctx *pulumi.Context, args LookupJobTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupJobTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobTemplateResult, error) {
			args := v.(LookupJobTemplateArgs)
			r, err := LookupJobTemplate(ctx, &args, opts...)
			var s LookupJobTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobTemplateResultOutput)
}

type LookupJobTemplateOutputArgs struct {
	JobTemplateId pulumi.StringInput `pulumi:"jobTemplateId"`
}

func (LookupJobTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobTemplateArgs)(nil)).Elem()
}

type LookupJobTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupJobTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobTemplateResult)(nil)).Elem()
}

func (o LookupJobTemplateResultOutput) ToLookupJobTemplateResultOutput() LookupJobTemplateResultOutput {
	return o
}

func (o LookupJobTemplateResultOutput) ToLookupJobTemplateResultOutputWithContext(ctx context.Context) LookupJobTemplateResultOutput {
	return o
}

func (o LookupJobTemplateResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobTemplateResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupJobTemplateResultOutput) JobExecutionsRetryConfig() JobExecutionsRetryConfigPropertiesPtrOutput {
	return o.ApplyT(func(v LookupJobTemplateResult) *JobExecutionsRetryConfigProperties { return v.JobExecutionsRetryConfig }).(JobExecutionsRetryConfigPropertiesPtrOutput)
}

func (o LookupJobTemplateResultOutput) MaintenanceWindows() JobTemplateMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v LookupJobTemplateResult) []JobTemplateMaintenanceWindow { return v.MaintenanceWindows }).(JobTemplateMaintenanceWindowArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobTemplateResultOutput{})
}
