// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ServiceCatalog::TagOption
func LookupTagOption(ctx *pulumi.Context, args *LookupTagOptionArgs, opts ...pulumi.InvokeOption) (*LookupTagOptionResult, error) {
	var rv LookupTagOptionResult
	err := ctx.Invoke("aws-native:servicecatalog:getTagOption", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagOptionArgs struct {
	Id string `pulumi:"id"`
}

type LookupTagOptionResult struct {
	Active *bool   `pulumi:"active"`
	Id     *string `pulumi:"id"`
}

func LookupTagOptionOutput(ctx *pulumi.Context, args LookupTagOptionOutputArgs, opts ...pulumi.InvokeOption) LookupTagOptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagOptionResult, error) {
			args := v.(LookupTagOptionArgs)
			r, err := LookupTagOption(ctx, &args, opts...)
			return *r, err
		}).(LookupTagOptionResultOutput)
}

type LookupTagOptionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupTagOptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagOptionArgs)(nil)).Elem()
}

type LookupTagOptionResultOutput struct{ *pulumi.OutputState }

func (LookupTagOptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagOptionResult)(nil)).Elem()
}

func (o LookupTagOptionResultOutput) ToLookupTagOptionResultOutput() LookupTagOptionResultOutput {
	return o
}

func (o LookupTagOptionResultOutput) ToLookupTagOptionResultOutputWithContext(ctx context.Context) LookupTagOptionResultOutput {
	return o
}

func (o LookupTagOptionResultOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTagOptionResult) *bool { return v.Active }).(pulumi.BoolPtrOutput)
}

func (o LookupTagOptionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTagOptionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagOptionResultOutput{})
}
