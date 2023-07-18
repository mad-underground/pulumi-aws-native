// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema for AWS::ServiceCatalog::ServiceAction
type ServiceAction struct {
	pulumi.CustomResourceState

	AcceptLanguage ServiceActionAcceptLanguagePtrOutput        `pulumi:"acceptLanguage"`
	Definition     ServiceActionDefinitionParameterArrayOutput `pulumi:"definition"`
	DefinitionType ServiceActionDefinitionTypeOutput           `pulumi:"definitionType"`
	Description    pulumi.StringPtrOutput                      `pulumi:"description"`
	Name           pulumi.StringOutput                         `pulumi:"name"`
}

// NewServiceAction registers a new resource with the given unique name, arguments, and options.
func NewServiceAction(ctx *pulumi.Context,
	name string, args *ServiceActionArgs, opts ...pulumi.ResourceOption) (*ServiceAction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.DefinitionType == nil {
		return nil, errors.New("invalid value for required argument 'DefinitionType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceAction
	err := ctx.RegisterResource("aws-native:servicecatalog:ServiceAction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceAction gets an existing ServiceAction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceActionState, opts ...pulumi.ResourceOption) (*ServiceAction, error) {
	var resource ServiceAction
	err := ctx.ReadResource("aws-native:servicecatalog:ServiceAction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceAction resources.
type serviceActionState struct {
}

type ServiceActionState struct {
}

func (ServiceActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceActionState)(nil)).Elem()
}

type serviceActionArgs struct {
	AcceptLanguage *ServiceActionAcceptLanguage       `pulumi:"acceptLanguage"`
	Definition     []ServiceActionDefinitionParameter `pulumi:"definition"`
	DefinitionType ServiceActionDefinitionType        `pulumi:"definitionType"`
	Description    *string                            `pulumi:"description"`
	Name           *string                            `pulumi:"name"`
}

// The set of arguments for constructing a ServiceAction resource.
type ServiceActionArgs struct {
	AcceptLanguage ServiceActionAcceptLanguagePtrInput
	Definition     ServiceActionDefinitionParameterArrayInput
	DefinitionType ServiceActionDefinitionTypeInput
	Description    pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
}

func (ServiceActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceActionArgs)(nil)).Elem()
}

type ServiceActionInput interface {
	pulumi.Input

	ToServiceActionOutput() ServiceActionOutput
	ToServiceActionOutputWithContext(ctx context.Context) ServiceActionOutput
}

func (*ServiceAction) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAction)(nil)).Elem()
}

func (i *ServiceAction) ToServiceActionOutput() ServiceActionOutput {
	return i.ToServiceActionOutputWithContext(context.Background())
}

func (i *ServiceAction) ToServiceActionOutputWithContext(ctx context.Context) ServiceActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceActionOutput)
}

type ServiceActionOutput struct{ *pulumi.OutputState }

func (ServiceActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAction)(nil)).Elem()
}

func (o ServiceActionOutput) ToServiceActionOutput() ServiceActionOutput {
	return o
}

func (o ServiceActionOutput) ToServiceActionOutputWithContext(ctx context.Context) ServiceActionOutput {
	return o
}

func (o ServiceActionOutput) AcceptLanguage() ServiceActionAcceptLanguagePtrOutput {
	return o.ApplyT(func(v *ServiceAction) ServiceActionAcceptLanguagePtrOutput { return v.AcceptLanguage }).(ServiceActionAcceptLanguagePtrOutput)
}

func (o ServiceActionOutput) Definition() ServiceActionDefinitionParameterArrayOutput {
	return o.ApplyT(func(v *ServiceAction) ServiceActionDefinitionParameterArrayOutput { return v.Definition }).(ServiceActionDefinitionParameterArrayOutput)
}

func (o ServiceActionOutput) DefinitionType() ServiceActionDefinitionTypeOutput {
	return o.ApplyT(func(v *ServiceAction) ServiceActionDefinitionTypeOutput { return v.DefinitionType }).(ServiceActionDefinitionTypeOutput)
}

func (o ServiceActionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAction) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ServiceActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceAction) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceActionInput)(nil)).Elem(), &ServiceAction{})
	pulumi.RegisterOutputType(ServiceActionOutput{})
}
