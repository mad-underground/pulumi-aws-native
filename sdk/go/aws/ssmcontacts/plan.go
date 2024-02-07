// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssmcontacts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Engagement Plan for a SSM Incident Manager Contact.
type Plan struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the contact.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Contact ID for the AWS SSM Incident Manager Contact to associate the plan.
	ContactId pulumi.StringPtrOutput `pulumi:"contactId"`
	// Rotation Ids to associate with Oncall Contact for engagement.
	RotationIds pulumi.StringArrayOutput `pulumi:"rotationIds"`
	// The stages that an escalation plan or engagement plan engages contacts and contact methods in.
	Stages PlanStageArrayOutput `pulumi:"stages"`
}

// NewPlan registers a new resource with the given unique name, arguments, and options.
func NewPlan(ctx *pulumi.Context,
	name string, args *PlanArgs, opts ...pulumi.ResourceOption) (*Plan, error) {
	if args == nil {
		args = &PlanArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"contactId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Plan
	err := ctx.RegisterResource("aws-native:ssmcontacts:Plan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlan gets an existing Plan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlanState, opts ...pulumi.ResourceOption) (*Plan, error) {
	var resource Plan
	err := ctx.ReadResource("aws-native:ssmcontacts:Plan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Plan resources.
type planState struct {
}

type PlanState struct {
}

func (PlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*planState)(nil)).Elem()
}

type planArgs struct {
	// Contact ID for the AWS SSM Incident Manager Contact to associate the plan.
	ContactId *string `pulumi:"contactId"`
	// Rotation Ids to associate with Oncall Contact for engagement.
	RotationIds []string `pulumi:"rotationIds"`
	// The stages that an escalation plan or engagement plan engages contacts and contact methods in.
	Stages []PlanStage `pulumi:"stages"`
}

// The set of arguments for constructing a Plan resource.
type PlanArgs struct {
	// Contact ID for the AWS SSM Incident Manager Contact to associate the plan.
	ContactId pulumi.StringPtrInput
	// Rotation Ids to associate with Oncall Contact for engagement.
	RotationIds pulumi.StringArrayInput
	// The stages that an escalation plan or engagement plan engages contacts and contact methods in.
	Stages PlanStageArrayInput
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*planArgs)(nil)).Elem()
}

type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(ctx context.Context) PlanOutput
}

func (*Plan) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (i *Plan) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i *Plan) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}

type PlanOutput struct{ *pulumi.OutputState }

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

// The Amazon Resource Name (ARN) of the contact.
func (o PlanOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Plan) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Contact ID for the AWS SSM Incident Manager Contact to associate the plan.
func (o PlanOutput) ContactId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) pulumi.StringPtrOutput { return v.ContactId }).(pulumi.StringPtrOutput)
}

// Rotation Ids to associate with Oncall Contact for engagement.
func (o PlanOutput) RotationIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Plan) pulumi.StringArrayOutput { return v.RotationIds }).(pulumi.StringArrayOutput)
}

// The stages that an escalation plan or engagement plan engages contacts and contact methods in.
func (o PlanOutput) Stages() PlanStageArrayOutput {
	return o.ApplyT(func(v *Plan) PlanStageArrayOutput { return v.Stages }).(PlanStageArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PlanInput)(nil)).Elem(), &Plan{})
	pulumi.RegisterOutputType(PlanOutput{})
}
