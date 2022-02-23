// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ServiceCatalog::LaunchNotificationConstraint
func LookupLaunchNotificationConstraint(ctx *pulumi.Context, args *LookupLaunchNotificationConstraintArgs, opts ...pulumi.InvokeOption) (*LookupLaunchNotificationConstraintResult, error) {
	var rv LookupLaunchNotificationConstraintResult
	err := ctx.Invoke("aws-native:servicecatalog:getLaunchNotificationConstraint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLaunchNotificationConstraintArgs struct {
	Id string `pulumi:"id"`
}

type LookupLaunchNotificationConstraintResult struct {
	AcceptLanguage   *string  `pulumi:"acceptLanguage"`
	Description      *string  `pulumi:"description"`
	Id               *string  `pulumi:"id"`
	NotificationArns []string `pulumi:"notificationArns"`
}

func LookupLaunchNotificationConstraintOutput(ctx *pulumi.Context, args LookupLaunchNotificationConstraintOutputArgs, opts ...pulumi.InvokeOption) LookupLaunchNotificationConstraintResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLaunchNotificationConstraintResult, error) {
			args := v.(LookupLaunchNotificationConstraintArgs)
			r, err := LookupLaunchNotificationConstraint(ctx, &args, opts...)
			return *r, err
		}).(LookupLaunchNotificationConstraintResultOutput)
}

type LookupLaunchNotificationConstraintOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupLaunchNotificationConstraintOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLaunchNotificationConstraintArgs)(nil)).Elem()
}

type LookupLaunchNotificationConstraintResultOutput struct{ *pulumi.OutputState }

func (LookupLaunchNotificationConstraintResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLaunchNotificationConstraintResult)(nil)).Elem()
}

func (o LookupLaunchNotificationConstraintResultOutput) ToLookupLaunchNotificationConstraintResultOutput() LookupLaunchNotificationConstraintResultOutput {
	return o
}

func (o LookupLaunchNotificationConstraintResultOutput) ToLookupLaunchNotificationConstraintResultOutputWithContext(ctx context.Context) LookupLaunchNotificationConstraintResultOutput {
	return o
}

func (o LookupLaunchNotificationConstraintResultOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLaunchNotificationConstraintResult) *string { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

func (o LookupLaunchNotificationConstraintResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLaunchNotificationConstraintResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLaunchNotificationConstraintResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLaunchNotificationConstraintResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupLaunchNotificationConstraintResultOutput) NotificationArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLaunchNotificationConstraintResult) []string { return v.NotificationArns }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLaunchNotificationConstraintResultOutput{})
}