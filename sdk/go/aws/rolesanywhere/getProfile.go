// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rolesanywhere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::RolesAnywhere::Profile Resource Type
func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProfileResult
	err := ctx.Invoke("aws-native:rolesanywhere:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfileArgs struct {
	ProfileId string `pulumi:"profileId"`
}

type LookupProfileResult struct {
	DurationSeconds           *float64     `pulumi:"durationSeconds"`
	Enabled                   *bool        `pulumi:"enabled"`
	ManagedPolicyArns         []string     `pulumi:"managedPolicyArns"`
	Name                      *string      `pulumi:"name"`
	ProfileArn                *string      `pulumi:"profileArn"`
	ProfileId                 *string      `pulumi:"profileId"`
	RequireInstanceProperties *bool        `pulumi:"requireInstanceProperties"`
	RoleArns                  []string     `pulumi:"roleArns"`
	SessionPolicy             *string      `pulumi:"sessionPolicy"`
	Tags                      []ProfileTag `pulumi:"tags"`
}

func LookupProfileOutput(ctx *pulumi.Context, args LookupProfileOutputArgs, opts ...pulumi.InvokeOption) LookupProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProfileResult, error) {
			args := v.(LookupProfileArgs)
			r, err := LookupProfile(ctx, &args, opts...)
			var s LookupProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProfileResultOutput)
}

type LookupProfileOutputArgs struct {
	ProfileId pulumi.StringInput `pulumi:"profileId"`
}

func (LookupProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileArgs)(nil)).Elem()
}

type LookupProfileResultOutput struct{ *pulumi.OutputState }

func (LookupProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileResult)(nil)).Elem()
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutput() LookupProfileResultOutput {
	return o
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutputWithContext(ctx context.Context) LookupProfileResultOutput {
	return o
}

func (o LookupProfileResultOutput) DurationSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *float64 { return v.DurationSeconds }).(pulumi.Float64PtrOutput)
}

func (o LookupProfileResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupProfileResultOutput) ManagedPolicyArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProfileResult) []string { return v.ManagedPolicyArns }).(pulumi.StringArrayOutput)
}

func (o LookupProfileResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) ProfileArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.ProfileArn }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) ProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.ProfileId }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) RequireInstanceProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *bool { return v.RequireInstanceProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupProfileResultOutput) RoleArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProfileResult) []string { return v.RoleArns }).(pulumi.StringArrayOutput)
}

func (o LookupProfileResultOutput) SessionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.SessionPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) Tags() ProfileTagArrayOutput {
	return o.ApplyT(func(v LookupProfileResult) []ProfileTag { return v.Tags }).(ProfileTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProfileResultOutput{})
}
