// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::LoggerDefinition
func LookupLoggerDefinition(ctx *pulumi.Context, args *LookupLoggerDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupLoggerDefinitionResult, error) {
	var rv LookupLoggerDefinitionResult
	err := ctx.Invoke("aws-native:greengrass:getLoggerDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoggerDefinitionArgs struct {
	Id string `pulumi:"id"`
}

type LookupLoggerDefinitionResult struct {
	Arn              *string     `pulumi:"arn"`
	Id               *string     `pulumi:"id"`
	LatestVersionArn *string     `pulumi:"latestVersionArn"`
	Name             *string     `pulumi:"name"`
	Tags             interface{} `pulumi:"tags"`
}

func LookupLoggerDefinitionOutput(ctx *pulumi.Context, args LookupLoggerDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupLoggerDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoggerDefinitionResult, error) {
			args := v.(LookupLoggerDefinitionArgs)
			r, err := LookupLoggerDefinition(ctx, &args, opts...)
			return *r, err
		}).(LookupLoggerDefinitionResultOutput)
}

type LookupLoggerDefinitionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupLoggerDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoggerDefinitionArgs)(nil)).Elem()
}

type LookupLoggerDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupLoggerDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoggerDefinitionResult)(nil)).Elem()
}

func (o LookupLoggerDefinitionResultOutput) ToLookupLoggerDefinitionResultOutput() LookupLoggerDefinitionResultOutput {
	return o
}

func (o LookupLoggerDefinitionResultOutput) ToLookupLoggerDefinitionResultOutputWithContext(ctx context.Context) LookupLoggerDefinitionResultOutput {
	return o
}

func (o LookupLoggerDefinitionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoggerDefinitionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupLoggerDefinitionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoggerDefinitionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupLoggerDefinitionResultOutput) LatestVersionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoggerDefinitionResult) *string { return v.LatestVersionArn }).(pulumi.StringPtrOutput)
}

func (o LookupLoggerDefinitionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoggerDefinitionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupLoggerDefinitionResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupLoggerDefinitionResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoggerDefinitionResultOutput{})
}