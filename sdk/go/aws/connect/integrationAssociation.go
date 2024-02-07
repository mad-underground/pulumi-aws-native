// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Connect::IntegrationAssociation
type IntegrationAssociation struct {
	pulumi.CustomResourceState

	InstanceId               pulumi.StringOutput                         `pulumi:"instanceId"`
	IntegrationArn           pulumi.StringOutput                         `pulumi:"integrationArn"`
	IntegrationAssociationId pulumi.StringOutput                         `pulumi:"integrationAssociationId"`
	IntegrationType          IntegrationAssociationIntegrationTypeOutput `pulumi:"integrationType"`
}

// NewIntegrationAssociation registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAssociation(ctx *pulumi.Context,
	name string, args *IntegrationAssociationArgs, opts ...pulumi.ResourceOption) (*IntegrationAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.IntegrationArn == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationArn'")
	}
	if args.IntegrationType == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationType'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"instanceId",
		"integrationArn",
		"integrationType",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IntegrationAssociation
	err := ctx.RegisterResource("aws-native:connect:IntegrationAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAssociation gets an existing IntegrationAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAssociationState, opts ...pulumi.ResourceOption) (*IntegrationAssociation, error) {
	var resource IntegrationAssociation
	err := ctx.ReadResource("aws-native:connect:IntegrationAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAssociation resources.
type integrationAssociationState struct {
}

type IntegrationAssociationState struct {
}

func (IntegrationAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAssociationState)(nil)).Elem()
}

type integrationAssociationArgs struct {
	InstanceId      string                                `pulumi:"instanceId"`
	IntegrationArn  string                                `pulumi:"integrationArn"`
	IntegrationType IntegrationAssociationIntegrationType `pulumi:"integrationType"`
}

// The set of arguments for constructing a IntegrationAssociation resource.
type IntegrationAssociationArgs struct {
	InstanceId      pulumi.StringInput
	IntegrationArn  pulumi.StringInput
	IntegrationType IntegrationAssociationIntegrationTypeInput
}

func (IntegrationAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAssociationArgs)(nil)).Elem()
}

type IntegrationAssociationInput interface {
	pulumi.Input

	ToIntegrationAssociationOutput() IntegrationAssociationOutput
	ToIntegrationAssociationOutputWithContext(ctx context.Context) IntegrationAssociationOutput
}

func (*IntegrationAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAssociation)(nil)).Elem()
}

func (i *IntegrationAssociation) ToIntegrationAssociationOutput() IntegrationAssociationOutput {
	return i.ToIntegrationAssociationOutputWithContext(context.Background())
}

func (i *IntegrationAssociation) ToIntegrationAssociationOutputWithContext(ctx context.Context) IntegrationAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAssociationOutput)
}

type IntegrationAssociationOutput struct{ *pulumi.OutputState }

func (IntegrationAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAssociation)(nil)).Elem()
}

func (o IntegrationAssociationOutput) ToIntegrationAssociationOutput() IntegrationAssociationOutput {
	return o
}

func (o IntegrationAssociationOutput) ToIntegrationAssociationOutputWithContext(ctx context.Context) IntegrationAssociationOutput {
	return o
}

func (o IntegrationAssociationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAssociation) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

func (o IntegrationAssociationOutput) IntegrationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAssociation) pulumi.StringOutput { return v.IntegrationArn }).(pulumi.StringOutput)
}

func (o IntegrationAssociationOutput) IntegrationAssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAssociation) pulumi.StringOutput { return v.IntegrationAssociationId }).(pulumi.StringOutput)
}

func (o IntegrationAssociationOutput) IntegrationType() IntegrationAssociationIntegrationTypeOutput {
	return o.ApplyT(func(v *IntegrationAssociation) IntegrationAssociationIntegrationTypeOutput { return v.IntegrationType }).(IntegrationAssociationIntegrationTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationAssociationInput)(nil)).Elem(), &IntegrationAssociation{})
	pulumi.RegisterOutputType(IntegrationAssociationOutput{})
}
