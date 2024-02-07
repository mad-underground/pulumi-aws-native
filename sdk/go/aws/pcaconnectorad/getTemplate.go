// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pcaconnectorad

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a template that defines certificate configurations, both for issuance and client handling
func LookupTemplate(ctx *pulumi.Context, args *LookupTemplateArgs, opts ...pulumi.InvokeOption) (*LookupTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTemplateResult
	err := ctx.Invoke("aws-native:pcaconnectorad:getTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplateArgs struct {
	TemplateArn string `pulumi:"templateArn"`
}

type LookupTemplateResult struct {
	TemplateArn *string `pulumi:"templateArn"`
}

func LookupTemplateOutput(ctx *pulumi.Context, args LookupTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTemplateResult, error) {
			args := v.(LookupTemplateArgs)
			r, err := LookupTemplate(ctx, &args, opts...)
			var s LookupTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTemplateResultOutput)
}

type LookupTemplateOutputArgs struct {
	TemplateArn pulumi.StringInput `pulumi:"templateArn"`
}

func (LookupTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateArgs)(nil)).Elem()
}

type LookupTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateResult)(nil)).Elem()
}

func (o LookupTemplateResultOutput) ToLookupTemplateResultOutput() LookupTemplateResultOutput {
	return o
}

func (o LookupTemplateResultOutput) ToLookupTemplateResultOutputWithContext(ctx context.Context) LookupTemplateResultOutput {
	return o
}

func (o LookupTemplateResultOutput) TemplateArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateResult) *string { return v.TemplateArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTemplateResultOutput{})
}
