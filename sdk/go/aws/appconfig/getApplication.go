// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppConfig::Application
func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("aws-native:appconfig:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	Id string `pulumi:"id"`
}

type LookupApplicationResult struct {
	Description *string           `pulumi:"description"`
	Id          *string           `pulumi:"id"`
	Name        *string           `pulumi:"name"`
	Tags        []ApplicationTags `pulumi:"tags"`
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationResult, error) {
			args := v.(LookupApplicationArgs)
			r, err := LookupApplication(ctx, &args, opts...)
			return *r, err
		}).(LookupApplicationResultOutput)
}

type LookupApplicationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}

type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationResultOutput) Tags() ApplicationTagsArrayOutput {
	return o.ApplyT(func(v LookupApplicationResult) []ApplicationTags { return v.Tags }).(ApplicationTagsArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
