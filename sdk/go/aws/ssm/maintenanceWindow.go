// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SSM::MaintenanceWindow
//
// Deprecated: MaintenanceWindow is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type MaintenanceWindow struct {
	pulumi.CustomResourceState

	AllowUnassociatedTargets pulumi.BoolOutput               `pulumi:"allowUnassociatedTargets"`
	Cutoff                   pulumi.IntOutput                `pulumi:"cutoff"`
	Description              pulumi.StringPtrOutput          `pulumi:"description"`
	Duration                 pulumi.IntOutput                `pulumi:"duration"`
	EndDate                  pulumi.StringPtrOutput          `pulumi:"endDate"`
	Name                     pulumi.StringOutput             `pulumi:"name"`
	Schedule                 pulumi.StringOutput             `pulumi:"schedule"`
	ScheduleOffset           pulumi.IntPtrOutput             `pulumi:"scheduleOffset"`
	ScheduleTimezone         pulumi.StringPtrOutput          `pulumi:"scheduleTimezone"`
	StartDate                pulumi.StringPtrOutput          `pulumi:"startDate"`
	Tags                     MaintenanceWindowTagArrayOutput `pulumi:"tags"`
}

// NewMaintenanceWindow registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceWindow(ctx *pulumi.Context,
	name string, args *MaintenanceWindowArgs, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowUnassociatedTargets == nil {
		return nil, errors.New("invalid value for required argument 'AllowUnassociatedTargets'")
	}
	if args.Cutoff == nil {
		return nil, errors.New("invalid value for required argument 'Cutoff'")
	}
	if args.Duration == nil {
		return nil, errors.New("invalid value for required argument 'Duration'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	var resource MaintenanceWindow
	err := ctx.RegisterResource("aws-native:ssm:MaintenanceWindow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceWindow gets an existing MaintenanceWindow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceWindow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceWindowState, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	var resource MaintenanceWindow
	err := ctx.ReadResource("aws-native:ssm:MaintenanceWindow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceWindow resources.
type maintenanceWindowState struct {
}

type MaintenanceWindowState struct {
}

func (MaintenanceWindowState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowState)(nil)).Elem()
}

type maintenanceWindowArgs struct {
	AllowUnassociatedTargets bool                   `pulumi:"allowUnassociatedTargets"`
	Cutoff                   int                    `pulumi:"cutoff"`
	Description              *string                `pulumi:"description"`
	Duration                 int                    `pulumi:"duration"`
	EndDate                  *string                `pulumi:"endDate"`
	Name                     *string                `pulumi:"name"`
	Schedule                 string                 `pulumi:"schedule"`
	ScheduleOffset           *int                   `pulumi:"scheduleOffset"`
	ScheduleTimezone         *string                `pulumi:"scheduleTimezone"`
	StartDate                *string                `pulumi:"startDate"`
	Tags                     []MaintenanceWindowTag `pulumi:"tags"`
}

// The set of arguments for constructing a MaintenanceWindow resource.
type MaintenanceWindowArgs struct {
	AllowUnassociatedTargets pulumi.BoolInput
	Cutoff                   pulumi.IntInput
	Description              pulumi.StringPtrInput
	Duration                 pulumi.IntInput
	EndDate                  pulumi.StringPtrInput
	Name                     pulumi.StringPtrInput
	Schedule                 pulumi.StringInput
	ScheduleOffset           pulumi.IntPtrInput
	ScheduleTimezone         pulumi.StringPtrInput
	StartDate                pulumi.StringPtrInput
	Tags                     MaintenanceWindowTagArrayInput
}

func (MaintenanceWindowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowArgs)(nil)).Elem()
}

type MaintenanceWindowInput interface {
	pulumi.Input

	ToMaintenanceWindowOutput() MaintenanceWindowOutput
	ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput
}

func (*MaintenanceWindow) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (i *MaintenanceWindow) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return i.ToMaintenanceWindowOutputWithContext(context.Background())
}

func (i *MaintenanceWindow) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowOutput)
}

type MaintenanceWindowOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) AllowUnassociatedTargets() pulumi.BoolOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.BoolOutput { return v.AllowUnassociatedTargets }).(pulumi.BoolOutput)
}

func (o MaintenanceWindowOutput) Cutoff() pulumi.IntOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.IntOutput { return v.Cutoff }).(pulumi.IntOutput)
}

func (o MaintenanceWindowOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowOutput) Duration() pulumi.IntOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.IntOutput { return v.Duration }).(pulumi.IntOutput)
}

func (o MaintenanceWindowOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringPtrOutput { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MaintenanceWindowOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringOutput { return v.Schedule }).(pulumi.StringOutput)
}

func (o MaintenanceWindowOutput) ScheduleOffset() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.IntPtrOutput { return v.ScheduleOffset }).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowOutput) ScheduleTimezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringPtrOutput { return v.ScheduleTimezone }).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringPtrOutput { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowOutput) Tags() MaintenanceWindowTagArrayOutput {
	return o.ApplyT(func(v *MaintenanceWindow) MaintenanceWindowTagArrayOutput { return v.Tags }).(MaintenanceWindowTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowInput)(nil)).Elem(), &MaintenanceWindow{})
	pulumi.RegisterOutputType(MaintenanceWindowOutput{})
}
