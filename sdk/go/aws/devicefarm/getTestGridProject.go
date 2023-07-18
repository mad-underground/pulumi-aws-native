// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devicefarm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::DeviceFarm::TestGridProject creates a new TestGrid Project
func LookupTestGridProject(ctx *pulumi.Context, args *LookupTestGridProjectArgs, opts ...pulumi.InvokeOption) (*LookupTestGridProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTestGridProjectResult
	err := ctx.Invoke("aws-native:devicefarm:getTestGridProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTestGridProjectArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupTestGridProjectResult struct {
	Arn         *string              `pulumi:"arn"`
	Description *string              `pulumi:"description"`
	Name        *string              `pulumi:"name"`
	Tags        []TestGridProjectTag `pulumi:"tags"`
}

func LookupTestGridProjectOutput(ctx *pulumi.Context, args LookupTestGridProjectOutputArgs, opts ...pulumi.InvokeOption) LookupTestGridProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTestGridProjectResult, error) {
			args := v.(LookupTestGridProjectArgs)
			r, err := LookupTestGridProject(ctx, &args, opts...)
			var s LookupTestGridProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTestGridProjectResultOutput)
}

type LookupTestGridProjectOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupTestGridProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTestGridProjectArgs)(nil)).Elem()
}

type LookupTestGridProjectResultOutput struct{ *pulumi.OutputState }

func (LookupTestGridProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTestGridProjectResult)(nil)).Elem()
}

func (o LookupTestGridProjectResultOutput) ToLookupTestGridProjectResultOutput() LookupTestGridProjectResultOutput {
	return o
}

func (o LookupTestGridProjectResultOutput) ToLookupTestGridProjectResultOutputWithContext(ctx context.Context) LookupTestGridProjectResultOutput {
	return o
}

func (o LookupTestGridProjectResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTestGridProjectResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupTestGridProjectResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTestGridProjectResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupTestGridProjectResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTestGridProjectResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupTestGridProjectResultOutput) Tags() TestGridProjectTagArrayOutput {
	return o.ApplyT(func(v LookupTestGridProjectResult) []TestGridProjectTag { return v.Tags }).(TestGridProjectTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTestGridProjectResultOutput{})
}
