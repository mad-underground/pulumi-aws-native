// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datazone

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::DataZone::EnvironmentBlueprintConfiguration Resource Type
func LookupEnvironmentBlueprintConfiguration(ctx *pulumi.Context, args *LookupEnvironmentBlueprintConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentBlueprintConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentBlueprintConfigurationResult
	err := ctx.Invoke("aws-native:datazone:getEnvironmentBlueprintConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentBlueprintConfigurationArgs struct {
	DomainId               string `pulumi:"domainId"`
	EnvironmentBlueprintId string `pulumi:"environmentBlueprintId"`
}

type LookupEnvironmentBlueprintConfigurationResult struct {
	CreatedAt              *string                                              `pulumi:"createdAt"`
	DomainId               *string                                              `pulumi:"domainId"`
	EnabledRegions         []string                                             `pulumi:"enabledRegions"`
	EnvironmentBlueprintId *string                                              `pulumi:"environmentBlueprintId"`
	ManageAccessRoleArn    *string                                              `pulumi:"manageAccessRoleArn"`
	ProvisioningRoleArn    *string                                              `pulumi:"provisioningRoleArn"`
	RegionalParameters     []EnvironmentBlueprintConfigurationRegionalParameter `pulumi:"regionalParameters"`
	UpdatedAt              *string                                              `pulumi:"updatedAt"`
}

func LookupEnvironmentBlueprintConfigurationOutput(ctx *pulumi.Context, args LookupEnvironmentBlueprintConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentBlueprintConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentBlueprintConfigurationResult, error) {
			args := v.(LookupEnvironmentBlueprintConfigurationArgs)
			r, err := LookupEnvironmentBlueprintConfiguration(ctx, &args, opts...)
			var s LookupEnvironmentBlueprintConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentBlueprintConfigurationResultOutput)
}

type LookupEnvironmentBlueprintConfigurationOutputArgs struct {
	DomainId               pulumi.StringInput `pulumi:"domainId"`
	EnvironmentBlueprintId pulumi.StringInput `pulumi:"environmentBlueprintId"`
}

func (LookupEnvironmentBlueprintConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentBlueprintConfigurationArgs)(nil)).Elem()
}

type LookupEnvironmentBlueprintConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentBlueprintConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentBlueprintConfigurationResult)(nil)).Elem()
}

func (o LookupEnvironmentBlueprintConfigurationResultOutput) ToLookupEnvironmentBlueprintConfigurationResultOutput() LookupEnvironmentBlueprintConfigurationResultOutput {
	return o
}

func (o LookupEnvironmentBlueprintConfigurationResultOutput) ToLookupEnvironmentBlueprintConfigurationResultOutputWithContext(ctx context.Context) LookupEnvironmentBlueprintConfigurationResultOutput {
	return o
}

func (o LookupEnvironmentBlueprintConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupEnvironmentBlueprintConfigurationResult] {
	return pulumix.Output[LookupEnvironmentBlueprintConfigurationResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupEnvironmentBlueprintConfigurationResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentBlueprintConfigurationResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentBlueprintConfigurationResultOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentBlueprintConfigurationResult) *string { return v.DomainId }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentBlueprintConfigurationResultOutput) EnabledRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEnvironmentBlueprintConfigurationResult) []string { return v.EnabledRegions }).(pulumi.StringArrayOutput)
}

func (o LookupEnvironmentBlueprintConfigurationResultOutput) EnvironmentBlueprintId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentBlueprintConfigurationResult) *string { return v.EnvironmentBlueprintId }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentBlueprintConfigurationResultOutput) ManageAccessRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentBlueprintConfigurationResult) *string { return v.ManageAccessRoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentBlueprintConfigurationResultOutput) ProvisioningRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentBlueprintConfigurationResult) *string { return v.ProvisioningRoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentBlueprintConfigurationResultOutput) RegionalParameters() EnvironmentBlueprintConfigurationRegionalParameterArrayOutput {
	return o.ApplyT(func(v LookupEnvironmentBlueprintConfigurationResult) []EnvironmentBlueprintConfigurationRegionalParameter {
		return v.RegionalParameters
	}).(EnvironmentBlueprintConfigurationRegionalParameterArrayOutput)
}

func (o LookupEnvironmentBlueprintConfigurationResultOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentBlueprintConfigurationResult) *string { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentBlueprintConfigurationResultOutput{})
}
