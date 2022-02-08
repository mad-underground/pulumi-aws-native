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
	// The Amazon Resource Name (ARN) of the Function URL.
	Arn string `pulumi:"arn"`
}

type LookupUrlResult struct {
	// The Amazon Resource Name (ARN) of the Function URL.
	Arn *string `pulumi:"arn"`
	// Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
	AuthorizationType *UrlAuthorizationType `pulumi:"authorizationType"`
	Cors              *UrlCors              `pulumi:"cors"`
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
	// The Amazon Resource Name (ARN) of the Function URL.
	Arn pulumi.StringInput `pulumi:"arn"`
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

// The Amazon Resource Name (ARN) of the Function URL.
func (o LookupUrlResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUrlResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
func (o LookupUrlResultOutput) AuthorizationType() UrlAuthorizationTypePtrOutput {
	return o.ApplyT(func(v LookupUrlResult) *UrlAuthorizationType { return v.AuthorizationType }).(UrlAuthorizationTypePtrOutput)
}

func (o LookupUrlResultOutput) Cors() UrlCorsPtrOutput {
	return o.ApplyT(func(v LookupUrlResult) *UrlCors { return v.Cors }).(UrlCorsPtrOutput)
}

// The generated url for this resource.
func (o LookupUrlResultOutput) FunctionUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUrlResult) *string { return v.FunctionUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUrlResultOutput{})
}
