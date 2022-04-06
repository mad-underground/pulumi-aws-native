// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lambda::Url
func LookupUrl(ctx *pulumi.Context, args *LookupUrlArgs, opts ...pulumi.InvokeOption) (*LookupUrlResult, error) {
	var rv LookupUrlResult
	err := ctx.Invoke("aws-native:lambda:getUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUrlArgs struct {
	// The fully qualified Amazon Resource Name (ARN) of the function associated with the Function URL.
	FunctionArn string `pulumi:"functionArn"`
}

type LookupUrlResult struct {
	// Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
	AuthType *UrlAuthType `pulumi:"authType"`
	Cors     *UrlCors     `pulumi:"cors"`
	// The fully qualified Amazon Resource Name (ARN) of the function associated with the Function URL.
	FunctionArn *string `pulumi:"functionArn"`
	// The generated url for this resource.
	FunctionUrl *string `pulumi:"functionUrl"`
}

func LookupUrlOutput(ctx *pulumi.Context, args LookupUrlOutputArgs, opts ...pulumi.InvokeOption) LookupUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUrlResult, error) {
			args := v.(LookupUrlArgs)
			r, err := LookupUrl(ctx, &args, opts...)
			return *r, err
		}).(LookupUrlResultOutput)
}

type LookupUrlOutputArgs struct {
	// The fully qualified Amazon Resource Name (ARN) of the function associated with the Function URL.
	FunctionArn pulumi.StringInput `pulumi:"functionArn"`
}

func (LookupUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUrlArgs)(nil)).Elem()
}

type LookupUrlResultOutput struct{ *pulumi.OutputState }

func (LookupUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUrlResult)(nil)).Elem()
}

func (o LookupUrlResultOutput) ToLookupUrlResultOutput() LookupUrlResultOutput {
	return o
}

func (o LookupUrlResultOutput) ToLookupUrlResultOutputWithContext(ctx context.Context) LookupUrlResultOutput {
	return o
}

// Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
func (o LookupUrlResultOutput) AuthType() UrlAuthTypePtrOutput {
	return o.ApplyT(func(v LookupUrlResult) *UrlAuthType { return v.AuthType }).(UrlAuthTypePtrOutput)
}

func (o LookupUrlResultOutput) Cors() UrlCorsPtrOutput {
	return o.ApplyT(func(v LookupUrlResult) *UrlCors { return v.Cors }).(UrlCorsPtrOutput)
}

// The fully qualified Amazon Resource Name (ARN) of the function associated with the Function URL.
func (o LookupUrlResultOutput) FunctionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUrlResult) *string { return v.FunctionArn }).(pulumi.StringPtrOutput)
}

// The generated url for this resource.
func (o LookupUrlResultOutput) FunctionUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUrlResult) *string { return v.FunctionUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUrlResultOutput{})
}
