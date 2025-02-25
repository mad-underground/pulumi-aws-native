// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::BaiduChannel
//
// Deprecated: BaiduChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type BaiduChannel struct {
	pulumi.CustomResourceState

	ApiKey        pulumi.StringOutput  `pulumi:"apiKey"`
	ApplicationId pulumi.StringOutput  `pulumi:"applicationId"`
	AwsId         pulumi.StringOutput  `pulumi:"awsId"`
	Enabled       pulumi.BoolPtrOutput `pulumi:"enabled"`
	SecretKey     pulumi.StringOutput  `pulumi:"secretKey"`
}

// NewBaiduChannel registers a new resource with the given unique name, arguments, and options.
func NewBaiduChannel(ctx *pulumi.Context,
	name string, args *BaiduChannelArgs, opts ...pulumi.ResourceOption) (*BaiduChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiKey == nil {
		return nil, errors.New("invalid value for required argument 'ApiKey'")
	}
	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.SecretKey == nil {
		return nil, errors.New("invalid value for required argument 'SecretKey'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"applicationId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BaiduChannel
	err := ctx.RegisterResource("aws-native:pinpoint:BaiduChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBaiduChannel gets an existing BaiduChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBaiduChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BaiduChannelState, opts ...pulumi.ResourceOption) (*BaiduChannel, error) {
	var resource BaiduChannel
	err := ctx.ReadResource("aws-native:pinpoint:BaiduChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BaiduChannel resources.
type baiduChannelState struct {
}

type BaiduChannelState struct {
}

func (BaiduChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*baiduChannelState)(nil)).Elem()
}

type baiduChannelArgs struct {
	ApiKey        string `pulumi:"apiKey"`
	ApplicationId string `pulumi:"applicationId"`
	Enabled       *bool  `pulumi:"enabled"`
	SecretKey     string `pulumi:"secretKey"`
}

// The set of arguments for constructing a BaiduChannel resource.
type BaiduChannelArgs struct {
	ApiKey        pulumi.StringInput
	ApplicationId pulumi.StringInput
	Enabled       pulumi.BoolPtrInput
	SecretKey     pulumi.StringInput
}

func (BaiduChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*baiduChannelArgs)(nil)).Elem()
}

type BaiduChannelInput interface {
	pulumi.Input

	ToBaiduChannelOutput() BaiduChannelOutput
	ToBaiduChannelOutputWithContext(ctx context.Context) BaiduChannelOutput
}

func (*BaiduChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**BaiduChannel)(nil)).Elem()
}

func (i *BaiduChannel) ToBaiduChannelOutput() BaiduChannelOutput {
	return i.ToBaiduChannelOutputWithContext(context.Background())
}

func (i *BaiduChannel) ToBaiduChannelOutputWithContext(ctx context.Context) BaiduChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaiduChannelOutput)
}

type BaiduChannelOutput struct{ *pulumi.OutputState }

func (BaiduChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BaiduChannel)(nil)).Elem()
}

func (o BaiduChannelOutput) ToBaiduChannelOutput() BaiduChannelOutput {
	return o
}

func (o BaiduChannelOutput) ToBaiduChannelOutputWithContext(ctx context.Context) BaiduChannelOutput {
	return o
}

func (o BaiduChannelOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v *BaiduChannel) pulumi.StringOutput { return v.ApiKey }).(pulumi.StringOutput)
}

func (o BaiduChannelOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *BaiduChannel) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o BaiduChannelOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *BaiduChannel) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o BaiduChannelOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BaiduChannel) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o BaiduChannelOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *BaiduChannel) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BaiduChannelInput)(nil)).Elem(), &BaiduChannel{})
	pulumi.RegisterOutputType(BaiduChannelOutput{})
}
