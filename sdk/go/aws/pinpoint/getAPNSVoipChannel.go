// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::APNSVoipChannel
func LookupAPNSVoipChannel(ctx *pulumi.Context, args *LookupAPNSVoipChannelArgs, opts ...pulumi.InvokeOption) (*LookupAPNSVoipChannelResult, error) {
	var rv LookupAPNSVoipChannelResult
	err := ctx.Invoke("aws-native:pinpoint:getAPNSVoipChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAPNSVoipChannelArgs struct {
	Id string `pulumi:"id"`
}

type LookupAPNSVoipChannelResult struct {
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

func LookupAPNSVoipChannelOutput(ctx *pulumi.Context, args LookupAPNSVoipChannelOutputArgs, opts ...pulumi.InvokeOption) LookupAPNSVoipChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAPNSVoipChannelResult, error) {
			args := v.(LookupAPNSVoipChannelArgs)
			r, err := LookupAPNSVoipChannel(ctx, &args, opts...)
			return *r, err
		}).(LookupAPNSVoipChannelResultOutput)
}

type LookupAPNSVoipChannelOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupAPNSVoipChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAPNSVoipChannelArgs)(nil)).Elem()
}

type LookupAPNSVoipChannelResultOutput struct{ *pulumi.OutputState }

func (LookupAPNSVoipChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAPNSVoipChannelResult)(nil)).Elem()
}

func (o LookupAPNSVoipChannelResultOutput) ToLookupAPNSVoipChannelResultOutput() LookupAPNSVoipChannelResultOutput {
	return o
}

func (o LookupAPNSVoipChannelResultOutput) ToLookupAPNSVoipChannelResultOutputWithContext(ctx context.Context) LookupAPNSVoipChannelResultOutput {
	return o
}

func (o LookupAPNSVoipChannelResultOutput) BundleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSVoipChannelResult) *string { return v.BundleId }).(pulumi.StringPtrOutput)
}

func (o LookupAPNSVoipChannelResultOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSVoipChannelResult) *string { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o LookupAPNSVoipChannelResultOutput) DefaultAuthenticationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSVoipChannelResult) *string { return v.DefaultAuthenticationMethod }).(pulumi.StringPtrOutput)
}

func (o LookupAPNSVoipChannelResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAPNSVoipChannelResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupAPNSVoipChannelResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSVoipChannelResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupAPNSVoipChannelResultOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSVoipChannelResult) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o LookupAPNSVoipChannelResultOutput) TeamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSVoipChannelResult) *string { return v.TeamId }).(pulumi.StringPtrOutput)
}

func (o LookupAPNSVoipChannelResultOutput) TokenKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSVoipChannelResult) *string { return v.TokenKey }).(pulumi.StringPtrOutput)
}

func (o LookupAPNSVoipChannelResultOutput) TokenKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPNSVoipChannelResult) *string { return v.TokenKeyId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAPNSVoipChannelResultOutput{})
}
