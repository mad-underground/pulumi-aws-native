// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GuardDuty::Master
//
// Deprecated: Master is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Master struct {
	pulumi.CustomResourceState

	DetectorId   pulumi.StringOutput    `pulumi:"detectorId"`
	InvitationId pulumi.StringPtrOutput `pulumi:"invitationId"`
	MasterId     pulumi.StringOutput    `pulumi:"masterId"`
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
	DetectorId   string  `pulumi:"detectorId"`
	InvitationId *string `pulumi:"invitationId"`
}

// The set of arguments for constructing a Master resource.
type MasterArgs struct {
	DetectorId   pulumi.StringInput
	InvitationId pulumi.StringPtrInput
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
	return reflect.TypeOf((*Master)(nil))
}

func (i *Master) ToMasterOutput() MasterOutput {
	return i.ToMasterOutputWithContext(context.Background())
}

func (i *Master) ToMasterOutputWithContext(ctx context.Context) MasterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MasterOutput)
}

type MasterOutput struct{ *pulumi.OutputState }

func (MasterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Master)(nil))
}

func (o MasterOutput) ToMasterOutput() MasterOutput {
	return o
}

func (o MasterOutput) ToMasterOutputWithContext(ctx context.Context) MasterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MasterOutput{})
}