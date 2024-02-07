// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SSM::MaintenanceWindow
func LookupMaintenanceWindow(ctx *pulumi.Context, args *LookupMaintenanceWindowArgs, opts ...pulumi.InvokeOption) (*LookupMaintenanceWindowResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMaintenanceWindowResult
	err := ctx.Invoke("aws-native:ssm:getMaintenanceWindow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMaintenanceWindowArgs struct {
	Id string `pulumi:"id"`
}

type LookupMaintenanceWindowResult struct {
	AllowUnassociatedTargets *bool                  `pulumi:"allowUnassociatedTargets"`
	Cutoff                   *int                   `pulumi:"cutoff"`
	Description              *string                `pulumi:"description"`
	Duration                 *int                   `pulumi:"duration"`
	EndDate                  *string                `pulumi:"endDate"`
	Id                       *string                `pulumi:"id"`
	Name                     *string                `pulumi:"name"`
	Schedule                 *string                `pulumi:"schedule"`
	ScheduleOffset           *int                   `pulumi:"scheduleOffset"`
	ScheduleTimezone         *string                `pulumi:"scheduleTimezone"`
	StartDate                *string                `pulumi:"startDate"`
	Tags                     []MaintenanceWindowTag `pulumi:"tags"`
}

func LookupMaintenanceWindowOutput(ctx *pulumi.Context, args LookupMaintenanceWindowOutputArgs, opts ...pulumi.InvokeOption) LookupMaintenanceWindowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMaintenanceWindowResult, error) {
			args := v.(LookupMaintenanceWindowArgs)
			r, err := LookupMaintenanceWindow(ctx, &args, opts...)
			var s LookupMaintenanceWindowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMaintenanceWindowResultOutput)
}

type LookupMaintenanceWindowOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupMaintenanceWindowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMaintenanceWindowArgs)(nil)).Elem()
}

type LookupMaintenanceWindowResultOutput struct{ *pulumi.OutputState }

func (LookupMaintenanceWindowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMaintenanceWindowResult)(nil)).Elem()
}

func (o LookupMaintenanceWindowResultOutput) ToLookupMaintenanceWindowResultOutput() LookupMaintenanceWindowResultOutput {
	return o
}

func (o LookupMaintenanceWindowResultOutput) ToLookupMaintenanceWindowResultOutputWithContext(ctx context.Context) LookupMaintenanceWindowResultOutput {
	return o
}

func (o LookupMaintenanceWindowResultOutput) AllowUnassociatedTargets() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowResult) *bool { return v.AllowUnassociatedTargets }).(pulumi.BoolPtrOutput)
}

func (o LookupMaintenanceWindowResultOutput) Cutoff() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowResult) *int { return v.Cutoff }).(pulumi.IntPtrOutput)
}

func (o LookupMaintenanceWindowResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowResultOutput) Duration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowResult) *int { return v.Duration }).(pulumi.IntPtrOutput)
}

func (o LookupMaintenanceWindowResultOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowResult) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowResultOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowResult) *string { return v.Schedule }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowResultOutput) ScheduleOffset() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowResult) *int { return v.ScheduleOffset }).(pulumi.IntPtrOutput)
}

func (o LookupMaintenanceWindowResultOutput) ScheduleTimezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowResult) *string { return v.ScheduleTimezone }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowResultOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowResult) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceWindowResultOutput) Tags() MaintenanceWindowTagArrayOutput {
	return o.ApplyT(func(v LookupMaintenanceWindowResult) []MaintenanceWindowTag { return v.Tags }).(MaintenanceWindowTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMaintenanceWindowResultOutput{})
}
