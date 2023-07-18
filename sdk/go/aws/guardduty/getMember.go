// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GuardDuty::Member
func LookupMember(ctx *pulumi.Context, args *LookupMemberArgs, opts ...pulumi.InvokeOption) (*LookupMemberResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMemberResult
	err := ctx.Invoke("aws-native:guardduty:getMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMemberArgs struct {
	MemberId string `pulumi:"memberId"`
}

type LookupMemberResult struct {
	DisableEmailNotification *bool   `pulumi:"disableEmailNotification"`
	Message                  *string `pulumi:"message"`
	Status                   *string `pulumi:"status"`
}

func LookupMemberOutput(ctx *pulumi.Context, args LookupMemberOutputArgs, opts ...pulumi.InvokeOption) LookupMemberResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMemberResult, error) {
			args := v.(LookupMemberArgs)
			r, err := LookupMember(ctx, &args, opts...)
			var s LookupMemberResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMemberResultOutput)
}

type LookupMemberOutputArgs struct {
	MemberId pulumi.StringInput `pulumi:"memberId"`
}

func (LookupMemberOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMemberArgs)(nil)).Elem()
}

type LookupMemberResultOutput struct{ *pulumi.OutputState }

func (LookupMemberResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMemberResult)(nil)).Elem()
}

func (o LookupMemberResultOutput) ToLookupMemberResultOutput() LookupMemberResultOutput {
	return o
}

func (o LookupMemberResultOutput) ToLookupMemberResultOutputWithContext(ctx context.Context) LookupMemberResultOutput {
	return o
}

func (o LookupMemberResultOutput) DisableEmailNotification() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMemberResult) *bool { return v.DisableEmailNotification }).(pulumi.BoolPtrOutput)
}

func (o LookupMemberResultOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMemberResult) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o LookupMemberResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMemberResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMemberResultOutput{})
}
