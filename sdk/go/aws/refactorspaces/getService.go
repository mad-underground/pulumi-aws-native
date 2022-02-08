// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package refactorspaces

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::RefactorSpaces::Service Resource Type
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("aws-native:refactorspaces:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	ApplicationIdentifier string `pulumi:"applicationIdentifier"`
	EnvironmentIdentifier string `pulumi:"environmentIdentifier"`
	ServiceIdentifier     string `pulumi:"serviceIdentifier"`
}

type LookupServiceResult struct {
	Arn               *string `pulumi:"arn"`
	ServiceIdentifier *string `pulumi:"serviceIdentifier"`
	// Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
	Tags []ServiceTag `pulumi:"tags"`
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceResult, error) {
			args := v.(LookupServiceArgs)
			r, err := LookupService(ctx, &args, opts...)
			return *r, err
		}).(LookupServiceResultOutput)
}

type LookupServiceOutputArgs struct {
	ApplicationIdentifier pulumi.StringInput `pulumi:"applicationIdentifier"`
	EnvironmentIdentifier pulumi.StringInput `pulumi:"environmentIdentifier"`
	ServiceIdentifier     pulumi.StringInput `pulumi:"serviceIdentifier"`
}

func (LookupServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}

type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) ServiceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.ServiceIdentifier }).(pulumi.StringPtrOutput)
}

// Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
func (o LookupServiceResultOutput) Tags() ServiceTagArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []ServiceTag { return v.Tags }).(ServiceTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}
