// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::APNSVoipChannel
func LookupApnsVoipChannel(ctx *pulumi.Context, args *LookupApnsVoipChannelArgs, opts ...pulumi.InvokeOption) (*LookupApnsVoipChannelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApnsVoipChannelResult
	err := ctx.Invoke("aws-native:pinpoint:getApnsVoipChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApnsVoipChannelArgs struct {
	Id string `pulumi:"id"`
}

type LookupApnsVoipChannelResult struct {
	BundleId                    *string `pulumi:"bundleId"`
	Certificate                 *string `pulumi:"certificate"`
	DefaultAuthenticationMethod *string `pulumi:"defaultAuthenticationMethod"`
	Enabled                     *bool   `pulumi:"enabled"`
	Id                          *string `pulumi:"id"`
	PrivateKey                  *string `pulumi:"privateKey"`
	TeamId                      *string `pulumi:"teamId"`
	TokenKey                    *string `pulumi:"tokenKey"`
	TokenKeyId                  *string `pulumi:"tokenKeyId"`
}

func LookupApnsVoipChannelOutput(ctx *pulumi.Context, args LookupApnsVoipChannelOutputArgs, opts ...pulumi.InvokeOption) LookupApnsVoipChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApnsVoipChannelResult, error) {
			args := v.(LookupApnsVoipChannelArgs)
			r, err := LookupApnsVoipChannel(ctx, &args, opts...)
			var s LookupApnsVoipChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApnsVoipChannelResultOutput)
}

type LookupApnsVoipChannelOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupApnsVoipChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApnsVoipChannelArgs)(nil)).Elem()
}

type LookupApnsVoipChannelResultOutput struct{ *pulumi.OutputState }

func (LookupApnsVoipChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApnsVoipChannelResult)(nil)).Elem()
}

func (o LookupApnsVoipChannelResultOutput) ToLookupApnsVoipChannelResultOutput() LookupApnsVoipChannelResultOutput {
	return o
}

func (o LookupApnsVoipChannelResultOutput) ToLookupApnsVoipChannelResultOutputWithContext(ctx context.Context) LookupApnsVoipChannelResultOutput {
	return o
}

func (o LookupApnsVoipChannelResultOutput) BundleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApnsVoipChannelResult) *string { return v.BundleId }).(pulumi.StringPtrOutput)
}

func (o LookupApnsVoipChannelResultOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApnsVoipChannelResult) *string { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o LookupApnsVoipChannelResultOutput) DefaultAuthenticationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApnsVoipChannelResult) *string { return v.DefaultAuthenticationMethod }).(pulumi.StringPtrOutput)
}

func (o LookupApnsVoipChannelResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApnsVoipChannelResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupApnsVoipChannelResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApnsVoipChannelResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupApnsVoipChannelResultOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApnsVoipChannelResult) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o LookupApnsVoipChannelResultOutput) TeamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApnsVoipChannelResult) *string { return v.TeamId }).(pulumi.StringPtrOutput)
}

func (o LookupApnsVoipChannelResultOutput) TokenKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApnsVoipChannelResult) *string { return v.TokenKey }).(pulumi.StringPtrOutput)
}

func (o LookupApnsVoipChannelResultOutput) TokenKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApnsVoipChannelResult) *string { return v.TokenKeyId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApnsVoipChannelResultOutput{})
}
