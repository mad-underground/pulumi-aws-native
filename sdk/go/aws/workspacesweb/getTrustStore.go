// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workspacesweb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::WorkSpacesWeb::TrustStore Resource Type
func LookupTrustStore(ctx *pulumi.Context, args *LookupTrustStoreArgs, opts ...pulumi.InvokeOption) (*LookupTrustStoreResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTrustStoreResult
	err := ctx.Invoke("aws-native:workspacesweb:getTrustStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrustStoreArgs struct {
	TrustStoreArn string `pulumi:"trustStoreArn"`
}

type LookupTrustStoreResult struct {
	AssociatedPortalArns []string        `pulumi:"associatedPortalArns"`
	CertificateList      []string        `pulumi:"certificateList"`
	Tags                 []TrustStoreTag `pulumi:"tags"`
	TrustStoreArn        *string         `pulumi:"trustStoreArn"`
}

func LookupTrustStoreOutput(ctx *pulumi.Context, args LookupTrustStoreOutputArgs, opts ...pulumi.InvokeOption) LookupTrustStoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTrustStoreResult, error) {
			args := v.(LookupTrustStoreArgs)
			r, err := LookupTrustStore(ctx, &args, opts...)
			var s LookupTrustStoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTrustStoreResultOutput)
}

type LookupTrustStoreOutputArgs struct {
	TrustStoreArn pulumi.StringInput `pulumi:"trustStoreArn"`
}

func (LookupTrustStoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrustStoreArgs)(nil)).Elem()
}

type LookupTrustStoreResultOutput struct{ *pulumi.OutputState }

func (LookupTrustStoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrustStoreResult)(nil)).Elem()
}

func (o LookupTrustStoreResultOutput) ToLookupTrustStoreResultOutput() LookupTrustStoreResultOutput {
	return o
}

func (o LookupTrustStoreResultOutput) ToLookupTrustStoreResultOutputWithContext(ctx context.Context) LookupTrustStoreResultOutput {
	return o
}

func (o LookupTrustStoreResultOutput) AssociatedPortalArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTrustStoreResult) []string { return v.AssociatedPortalArns }).(pulumi.StringArrayOutput)
}

func (o LookupTrustStoreResultOutput) CertificateList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTrustStoreResult) []string { return v.CertificateList }).(pulumi.StringArrayOutput)
}

func (o LookupTrustStoreResultOutput) Tags() TrustStoreTagArrayOutput {
	return o.ApplyT(func(v LookupTrustStoreResult) []TrustStoreTag { return v.Tags }).(TrustStoreTagArrayOutput)
}

func (o LookupTrustStoreResultOutput) TrustStoreArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrustStoreResult) *string { return v.TrustStoreArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTrustStoreResultOutput{})
}
