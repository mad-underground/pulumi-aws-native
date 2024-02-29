// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GuardDuty Master resource schema
type Master struct {
	pulumi.CustomResourceState

	// Unique ID of the detector of the GuardDuty member account.
	DetectorId pulumi.StringOutput `pulumi:"detectorId"`
	// Value used to validate the master account to the member account.
	InvitationId pulumi.StringPtrOutput `pulumi:"invitationId"`
	// ID of the account used as the master account.
	MasterId pulumi.StringOutput `pulumi:"masterId"`
}

// NewMaster registers a new resource with the given unique name, arguments, and options.
func NewMaster(ctx *pulumi.Context,
	name string, args *MasterArgs, opts ...pulumi.ResourceOption) (*Master, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DetectorId == nil {
		return nil, errors.New("invalid value for required argument 'DetectorId'")
	}
	if args.MasterId == nil {
		return nil, errors.New("invalid value for required argument 'MasterId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"detectorId",
		"invitationId",
		"masterId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Master
	err := ctx.RegisterResource("aws-native:guardduty:Master", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaster gets an existing Master resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MasterState, opts ...pulumi.ResourceOption) (*Master, error) {
	var resource Master
	err := ctx.ReadResource("aws-native:guardduty:Master", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Master resources.
type masterState struct {
}

type MasterState struct {
}

func (MasterState) ElementType() reflect.Type {
	return reflect.TypeOf((*masterState)(nil)).Elem()
}

type masterArgs struct {
	// Unique ID of the detector of the GuardDuty member account.
	DetectorId string `pulumi:"detectorId"`
	// Value used to validate the master account to the member account.
	InvitationId *string `pulumi:"invitationId"`
	// ID of the account used as the master account.
	MasterId string `pulumi:"masterId"`
}

// The set of arguments for constructing a Master resource.
type MasterArgs struct {
	// Unique ID of the detector of the GuardDuty member account.
	DetectorId pulumi.StringInput
	// Value used to validate the master account to the member account.
	InvitationId pulumi.StringPtrInput
	// ID of the account used as the master account.
	MasterId pulumi.StringInput
}

func (MasterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*masterArgs)(nil)).Elem()
}

type MasterInput interface {
	pulumi.Input

	ToMasterOutput() MasterOutput
	ToMasterOutputWithContext(ctx context.Context) MasterOutput
}

func (*Master) ElementType() reflect.Type {
	return reflect.TypeOf((**Master)(nil)).Elem()
}

func (i *Master) ToMasterOutput() MasterOutput {
	return i.ToMasterOutputWithContext(context.Background())
}

func (i *Master) ToMasterOutputWithContext(ctx context.Context) MasterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MasterOutput)
}

type MasterOutput struct{ *pulumi.OutputState }

func (MasterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Master)(nil)).Elem()
}

func (o MasterOutput) ToMasterOutput() MasterOutput {
	return o
}

func (o MasterOutput) ToMasterOutputWithContext(ctx context.Context) MasterOutput {
	return o
}

// Unique ID of the detector of the GuardDuty member account.
func (o MasterOutput) DetectorId() pulumi.StringOutput {
	return o.ApplyT(func(v *Master) pulumi.StringOutput { return v.DetectorId }).(pulumi.StringOutput)
}

// Value used to validate the master account to the member account.
func (o MasterOutput) InvitationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Master) pulumi.StringPtrOutput { return v.InvitationId }).(pulumi.StringPtrOutput)
}

// ID of the account used as the master account.
func (o MasterOutput) MasterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Master) pulumi.StringOutput { return v.MasterId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MasterInput)(nil)).Elem(), &Master{})
	pulumi.RegisterOutputType(MasterOutput{})
}
