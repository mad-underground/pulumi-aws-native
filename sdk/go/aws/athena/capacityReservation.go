// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package athena

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Athena::CapacityReservation
type CapacityReservation struct {
	pulumi.CustomResourceState

	// The number of DPUs Athena has provisioned and allocated for the reservation
	AllocatedDpus                   pulumi.IntOutput                                            `pulumi:"allocatedDpus"`
	Arn                             pulumi.StringOutput                                         `pulumi:"arn"`
	CapacityAssignmentConfiguration CapacityReservationCapacityAssignmentConfigurationPtrOutput `pulumi:"capacityAssignmentConfiguration"`
	// The date and time the reservation was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The timestamp when the last successful allocated was made
	LastSuccessfulAllocationTime pulumi.StringOutput `pulumi:"lastSuccessfulAllocationTime"`
	// The reservation name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the reservation.
	Status CapacityReservationStatusOutput `pulumi:"status"`
	// An array of key-value pairs to apply to this resource.
	Tags CapacityReservationTagArrayOutput `pulumi:"tags"`
	// The number of DPUs to request to be allocated to the reservation.
	TargetDpus pulumi.IntOutput `pulumi:"targetDpus"`
}

// NewCapacityReservation registers a new resource with the given unique name, arguments, and options.
func NewCapacityReservation(ctx *pulumi.Context,
	name string, args *CapacityReservationArgs, opts ...pulumi.ResourceOption) (*CapacityReservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TargetDpus == nil {
		return nil, errors.New("invalid value for required argument 'TargetDpus'")
	}
	var resource CapacityReservation
	err := ctx.RegisterResource("aws-native:athena:CapacityReservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCapacityReservation gets an existing CapacityReservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCapacityReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityReservationState, opts ...pulumi.ResourceOption) (*CapacityReservation, error) {
	var resource CapacityReservation
	err := ctx.ReadResource("aws-native:athena:CapacityReservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CapacityReservation resources.
type capacityReservationState struct {
}

type CapacityReservationState struct {
}

func (CapacityReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationState)(nil)).Elem()
}

type capacityReservationArgs struct {
	CapacityAssignmentConfiguration *CapacityReservationCapacityAssignmentConfiguration `pulumi:"capacityAssignmentConfiguration"`
	// The reservation name.
	Name *string `pulumi:"name"`
	// An array of key-value pairs to apply to this resource.
	Tags []CapacityReservationTag `pulumi:"tags"`
	// The number of DPUs to request to be allocated to the reservation.
	TargetDpus int `pulumi:"targetDpus"`
}

// The set of arguments for constructing a CapacityReservation resource.
type CapacityReservationArgs struct {
	CapacityAssignmentConfiguration CapacityReservationCapacityAssignmentConfigurationPtrInput
	// The reservation name.
	Name pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags CapacityReservationTagArrayInput
	// The number of DPUs to request to be allocated to the reservation.
	TargetDpus pulumi.IntInput
}

func (CapacityReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationArgs)(nil)).Elem()
}

type CapacityReservationInput interface {
	pulumi.Input

	ToCapacityReservationOutput() CapacityReservationOutput
	ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput
}

func (*CapacityReservation) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservation)(nil)).Elem()
}

func (i *CapacityReservation) ToCapacityReservationOutput() CapacityReservationOutput {
	return i.ToCapacityReservationOutputWithContext(context.Background())
}

func (i *CapacityReservation) ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationOutput)
}

type CapacityReservationOutput struct{ *pulumi.OutputState }

func (CapacityReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservation)(nil)).Elem()
}

func (o CapacityReservationOutput) ToCapacityReservationOutput() CapacityReservationOutput {
	return o
}

func (o CapacityReservationOutput) ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput {
	return o
}

// The number of DPUs Athena has provisioned and allocated for the reservation
func (o CapacityReservationOutput) AllocatedDpus() pulumi.IntOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.IntOutput { return v.AllocatedDpus }).(pulumi.IntOutput)
}

func (o CapacityReservationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o CapacityReservationOutput) CapacityAssignmentConfiguration() CapacityReservationCapacityAssignmentConfigurationPtrOutput {
	return o.ApplyT(func(v *CapacityReservation) CapacityReservationCapacityAssignmentConfigurationPtrOutput {
		return v.CapacityAssignmentConfiguration
	}).(CapacityReservationCapacityAssignmentConfigurationPtrOutput)
}

// The date and time the reservation was created.
func (o CapacityReservationOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The timestamp when the last successful allocated was made
func (o CapacityReservationOutput) LastSuccessfulAllocationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.LastSuccessfulAllocationTime }).(pulumi.StringOutput)
}

// The reservation name.
func (o CapacityReservationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the reservation.
func (o CapacityReservationOutput) Status() CapacityReservationStatusOutput {
	return o.ApplyT(func(v *CapacityReservation) CapacityReservationStatusOutput { return v.Status }).(CapacityReservationStatusOutput)
}

// An array of key-value pairs to apply to this resource.
func (o CapacityReservationOutput) Tags() CapacityReservationTagArrayOutput {
	return o.ApplyT(func(v *CapacityReservation) CapacityReservationTagArrayOutput { return v.Tags }).(CapacityReservationTagArrayOutput)
}

// The number of DPUs to request to be allocated to the reservation.
func (o CapacityReservationOutput) TargetDpus() pulumi.IntOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.IntOutput { return v.TargetDpus }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CapacityReservationInput)(nil)).Elem(), &CapacityReservation{})
	pulumi.RegisterOutputType(CapacityReservationOutput{})
}
