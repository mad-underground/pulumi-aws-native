// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema for AWS::ServiceCatalog::ServiceAction
func LookupServiceAction(ctx *pulumi.Context, args *LookupServiceActionArgs, opts ...pulumi.InvokeOption) (*LookupServiceActionResult, error) {
	var rv LookupServiceActionResult
	err := ctx.Invoke("aws-native:servicecatalog:getServiceAction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceActionArgs struct {
	Id string `pulumi:"id"`
}

type LookupServiceActionResult struct {
	AcceptLanguage *ServiceActionAcceptLanguage       `pulumi:"acceptLanguage"`
	Definition     []ServiceActionDefinitionParameter `pulumi:"definition"`
	DefinitionType *ServiceActionDefinitionType       `pulumi:"definitionType"`
	Description    *string                            `pulumi:"description"`
	Id             *string                            `pulumi:"id"`
	Name           *string                            `pulumi:"name"`
}

func LookupServiceActionOutput(ctx *pulumi.Context, args LookupServiceActionOutputArgs, opts ...pulumi.InvokeOption) LookupServiceActionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceActionResult, error) {
			args := v.(LookupServiceActionArgs)
			r, err := LookupServiceAction(ctx, &args, opts...)
			return *r, err
		}).(LookupServiceActionResultOutput)
}

type LookupServiceActionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupServiceActionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceActionArgs)(nil)).Elem()
}

type LookupServiceActionResultOutput struct{ *pulumi.OutputState }

func (LookupServiceActionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceActionResult)(nil)).Elem()
}

func (o LookupServiceActionResultOutput) ToLookupServiceActionResultOutput() LookupServiceActionResultOutput {
	return o
}

func (o LookupServiceActionResultOutput) ToLookupServiceActionResultOutputWithContext(ctx context.Context) LookupServiceActionResultOutput {
	return o
}

func (o LookupServiceActionResultOutput) AcceptLanguage() ServiceActionAcceptLanguagePtrOutput {
	return o.ApplyT(func(v LookupServiceActionResult) *ServiceActionAcceptLanguage { return v.AcceptLanguage }).(ServiceActionAcceptLanguagePtrOutput)
}

func (o LookupServiceActionResultOutput) Definition() ServiceActionDefinitionParameterArrayOutput {
	return o.ApplyT(func(v LookupServiceActionResult) []ServiceActionDefinitionParameter { return v.Definition }).(ServiceActionDefinitionParameterArrayOutput)
}

func (o LookupServiceActionResultOutput) DefinitionType() ServiceActionDefinitionTypePtrOutput {
	return o.ApplyT(func(v LookupServiceActionResult) *ServiceActionDefinitionType { return v.DefinitionType }).(ServiceActionDefinitionTypePtrOutput)
}

func (o LookupServiceActionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceActionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupServiceActionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceActionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupServiceActionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceActionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceActionResultOutput{})
}
