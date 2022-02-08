// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElasticBeanstalk::ConfigurationTemplate
func LookupConfigurationTemplate(ctx *pulumi.Context, args *LookupConfigurationTemplateArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationTemplateResult, error) {
	var rv LookupConfigurationTemplateResult
	err := ctx.Invoke("aws-native:elasticbeanstalk:getConfigurationTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationTemplateArgs struct {
	Id string `pulumi:"id"`
}

type LookupConfigurationTemplateResult struct {
	Description    *string                                           `pulumi:"description"`
	Id             *string                                           `pulumi:"id"`
	OptionSettings []ConfigurationTemplateConfigurationOptionSetting `pulumi:"optionSettings"`
}

func LookupConfigurationTemplateOutput(ctx *pulumi.Context, args LookupConfigurationTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationTemplateResult, error) {
			args := v.(LookupConfigurationTemplateArgs)
			r, err := LookupConfigurationTemplate(ctx, &args, opts...)
			return *r, err
		}).(LookupConfigurationTemplateResultOutput)
}

type LookupConfigurationTemplateOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupConfigurationTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationTemplateArgs)(nil)).Elem()
}

type LookupConfigurationTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationTemplateResult)(nil)).Elem()
}

func (o LookupConfigurationTemplateResultOutput) ToLookupConfigurationTemplateResultOutput() LookupConfigurationTemplateResultOutput {
	return o
}

func (o LookupConfigurationTemplateResultOutput) ToLookupConfigurationTemplateResultOutputWithContext(ctx context.Context) LookupConfigurationTemplateResultOutput {
	return o
}

func (o LookupConfigurationTemplateResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationTemplateResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationTemplateResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationTemplateResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationTemplateResultOutput) OptionSettings() ConfigurationTemplateConfigurationOptionSettingArrayOutput {
	return o.ApplyT(func(v LookupConfigurationTemplateResult) []ConfigurationTemplateConfigurationOptionSetting {
		return v.OptionSettings
	}).(ConfigurationTemplateConfigurationOptionSettingArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationTemplateResultOutput{})
}
