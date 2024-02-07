// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElasticLoadBalancingV2::ListenerCertificate
func LookupListenerCertificate(ctx *pulumi.Context, args *LookupListenerCertificateArgs, opts ...pulumi.InvokeOption) (*LookupListenerCertificateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupListenerCertificateResult
	err := ctx.Invoke("aws-native:elasticloadbalancingv2:getListenerCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupListenerCertificateArgs struct {
	Id string `pulumi:"id"`
}

type LookupListenerCertificateResult struct {
	Certificates []ListenerCertificateCertificate `pulumi:"certificates"`
	Id           *string                          `pulumi:"id"`
}

func LookupListenerCertificateOutput(ctx *pulumi.Context, args LookupListenerCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupListenerCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupListenerCertificateResult, error) {
			args := v.(LookupListenerCertificateArgs)
			r, err := LookupListenerCertificate(ctx, &args, opts...)
			var s LookupListenerCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupListenerCertificateResultOutput)
}

type LookupListenerCertificateOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupListenerCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupListenerCertificateArgs)(nil)).Elem()
}

type LookupListenerCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupListenerCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupListenerCertificateResult)(nil)).Elem()
}

func (o LookupListenerCertificateResultOutput) ToLookupListenerCertificateResultOutput() LookupListenerCertificateResultOutput {
	return o
}

func (o LookupListenerCertificateResultOutput) ToLookupListenerCertificateResultOutputWithContext(ctx context.Context) LookupListenerCertificateResultOutput {
	return o
}

func (o LookupListenerCertificateResultOutput) Certificates() ListenerCertificateCertificateArrayOutput {
	return o.ApplyT(func(v LookupListenerCertificateResult) []ListenerCertificateCertificate { return v.Certificates }).(ListenerCertificateCertificateArrayOutput)
}

func (o LookupListenerCertificateResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupListenerCertificateResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupListenerCertificateResultOutput{})
}
