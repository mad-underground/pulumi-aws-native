// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Config::ConfigurationRecorder
func LookupConfigurationRecorder(ctx *pulumi.Context, args *LookupConfigurationRecorderArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationRecorderResult, error) {
	var rv LookupConfigurationRecorderResult
	err := ctx.Invoke("aws-native:configuration:getConfigurationRecorder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationRecorderArgs struct {
	Id string `pulumi:"id"`
}

type LookupConfigurationRecorderResult struct {
	Id             *string                              `pulumi:"id"`
	RecordingGroup *ConfigurationRecorderRecordingGroup `pulumi:"recordingGroup"`
	RoleARN        *string                              `pulumi:"roleARN"`
}

func LookupConfigurationRecorderOutput(ctx *pulumi.Context, args LookupConfigurationRecorderOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationRecorderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationRecorderResult, error) {
			args := v.(LookupConfigurationRecorderArgs)
			r, err := LookupConfigurationRecorder(ctx, &args, opts...)
			return *r, err
		}).(LookupConfigurationRecorderResultOutput)
}

type LookupConfigurationRecorderOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupConfigurationRecorderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationRecorderArgs)(nil)).Elem()
}

type LookupConfigurationRecorderResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationRecorderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationRecorderResult)(nil)).Elem()
}

func (o LookupConfigurationRecorderResultOutput) ToLookupConfigurationRecorderResultOutput() LookupConfigurationRecorderResultOutput {
	return o
}

func (o LookupConfigurationRecorderResultOutput) ToLookupConfigurationRecorderResultOutputWithContext(ctx context.Context) LookupConfigurationRecorderResultOutput {
	return o
}

func (o LookupConfigurationRecorderResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationRecorderResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationRecorderResultOutput) RecordingGroup() ConfigurationRecorderRecordingGroupPtrOutput {
	return o.ApplyT(func(v LookupConfigurationRecorderResult) *ConfigurationRecorderRecordingGroup {
		return v.RecordingGroup
	}).(ConfigurationRecorderRecordingGroupPtrOutput)
}

func (o LookupConfigurationRecorderResultOutput) RoleARN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationRecorderResult) *string { return v.RoleARN }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationRecorderResultOutput{})
}
