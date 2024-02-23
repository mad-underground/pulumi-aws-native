// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotfleetwise

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::IoTFleetWise::Vehicle Resource Type
func LookupVehicle(ctx *pulumi.Context, args *LookupVehicleArgs, opts ...pulumi.InvokeOption) (*LookupVehicleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVehicleResult
	err := ctx.Invoke("aws-native:iotfleetwise:getVehicle", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVehicleArgs struct {
	Name string `pulumi:"name"`
}

type LookupVehicleResult struct {
	Arn                  *string           `pulumi:"arn"`
	Attributes           map[string]string `pulumi:"attributes"`
	CreationTime         *string           `pulumi:"creationTime"`
	DecoderManifestArn   *string           `pulumi:"decoderManifestArn"`
	LastModificationTime *string           `pulumi:"lastModificationTime"`
	ModelManifestArn     *string           `pulumi:"modelManifestArn"`
	Tags                 []aws.Tag         `pulumi:"tags"`
}

func LookupVehicleOutput(ctx *pulumi.Context, args LookupVehicleOutputArgs, opts ...pulumi.InvokeOption) LookupVehicleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVehicleResult, error) {
			args := v.(LookupVehicleArgs)
			r, err := LookupVehicle(ctx, &args, opts...)
			var s LookupVehicleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVehicleResultOutput)
}

type LookupVehicleOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupVehicleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVehicleArgs)(nil)).Elem()
}

type LookupVehicleResultOutput struct{ *pulumi.OutputState }

func (LookupVehicleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVehicleResult)(nil)).Elem()
}

func (o LookupVehicleResultOutput) ToLookupVehicleResultOutput() LookupVehicleResultOutput {
	return o
}

func (o LookupVehicleResultOutput) ToLookupVehicleResultOutputWithContext(ctx context.Context) LookupVehicleResultOutput {
	return o
}

func (o LookupVehicleResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVehicleResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupVehicleResultOutput) Attributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVehicleResult) map[string]string { return v.Attributes }).(pulumi.StringMapOutput)
}

func (o LookupVehicleResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVehicleResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupVehicleResultOutput) DecoderManifestArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVehicleResult) *string { return v.DecoderManifestArn }).(pulumi.StringPtrOutput)
}

func (o LookupVehicleResultOutput) LastModificationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVehicleResult) *string { return v.LastModificationTime }).(pulumi.StringPtrOutput)
}

func (o LookupVehicleResultOutput) ModelManifestArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVehicleResult) *string { return v.ModelManifestArn }).(pulumi.StringPtrOutput)
}

func (o LookupVehicleResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupVehicleResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVehicleResultOutput{})
}
