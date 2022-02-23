// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package location

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Location::Map Resource Type
func LookupMap(ctx *pulumi.Context, args *LookupMapArgs, opts ...pulumi.InvokeOption) (*LookupMapResult, error) {
	var rv LookupMapResult
	err := ctx.Invoke("aws-native:location:getMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMapArgs struct {
	MapName string `pulumi:"mapName"`
}

type LookupMapResult struct {
	Arn        *string `pulumi:"arn"`
	CreateTime *string `pulumi:"createTime"`
	DataSource *string `pulumi:"dataSource"`
	MapArn     *string `pulumi:"mapArn"`
	UpdateTime *string `pulumi:"updateTime"`
}

func LookupMapOutput(ctx *pulumi.Context, args LookupMapOutputArgs, opts ...pulumi.InvokeOption) LookupMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMapResult, error) {
			args := v.(LookupMapArgs)
			r, err := LookupMap(ctx, &args, opts...)
			return *r, err
		}).(LookupMapResultOutput)
}

type LookupMapOutputArgs struct {
	MapName pulumi.StringInput `pulumi:"mapName"`
}

func (LookupMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMapArgs)(nil)).Elem()
}

type LookupMapResultOutput struct{ *pulumi.OutputState }

func (LookupMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMapResult)(nil)).Elem()
}

func (o LookupMapResultOutput) ToLookupMapResultOutput() LookupMapResultOutput {
	return o
}

func (o LookupMapResultOutput) ToLookupMapResultOutputWithContext(ctx context.Context) LookupMapResultOutput {
	return o
}

func (o LookupMapResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMapResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupMapResultOutput) CreateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMapResult) *string { return v.CreateTime }).(pulumi.StringPtrOutput)
}

func (o LookupMapResultOutput) DataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMapResult) *string { return v.DataSource }).(pulumi.StringPtrOutput)
}

func (o LookupMapResultOutput) MapArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMapResult) *string { return v.MapArn }).(pulumi.StringPtrOutput)
}

func (o LookupMapResultOutput) UpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMapResult) *string { return v.UpdateTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMapResultOutput{})
}