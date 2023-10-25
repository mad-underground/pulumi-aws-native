// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of the AWS::QuickSight::Template Resource Type.
type Template struct {
	pulumi.CustomResourceState

	Arn                pulumi.StringOutput                   `pulumi:"arn"`
	AwsAccountId       pulumi.StringOutput                   `pulumi:"awsAccountId"`
	CreatedTime        pulumi.StringOutput                   `pulumi:"createdTime"`
	Definition         TemplateVersionDefinitionPtrOutput    `pulumi:"definition"`
	LastUpdatedTime    pulumi.StringOutput                   `pulumi:"lastUpdatedTime"`
	Name               pulumi.StringPtrOutput                `pulumi:"name"`
	Permissions        TemplateResourcePermissionArrayOutput `pulumi:"permissions"`
	SourceEntity       TemplateSourceEntityPtrOutput         `pulumi:"sourceEntity"`
	Tags               TemplateTagArrayOutput                `pulumi:"tags"`
	TemplateId         pulumi.StringOutput                   `pulumi:"templateId"`
	ValidationStrategy TemplateValidationStrategyPtrOutput   `pulumi:"validationStrategy"`
	Version            TemplateVersionOutput                 `pulumi:"version"`
	VersionDescription pulumi.StringPtrOutput                `pulumi:"versionDescription"`
}

// NewTemplate registers a new resource with the given unique name, arguments, and options.
func NewTemplate(ctx *pulumi.Context,
	name string, args *TemplateArgs, opts ...pulumi.ResourceOption) (*Template, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AwsAccountId == nil {
		return nil, errors.New("invalid value for required argument 'AwsAccountId'")
	}
	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"awsAccountId",
		"templateId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Template
	err := ctx.RegisterResource("aws-native:quicksight:Template", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplate gets an existing Template resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateState, opts ...pulumi.ResourceOption) (*Template, error) {
	var resource Template
	err := ctx.ReadResource("aws-native:quicksight:Template", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Template resources.
type templateState struct {
}

type TemplateState struct {
}

func (TemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateState)(nil)).Elem()
}

type templateArgs struct {
	AwsAccountId       string                       `pulumi:"awsAccountId"`
	Definition         *TemplateVersionDefinition   `pulumi:"definition"`
	Name               *string                      `pulumi:"name"`
	Permissions        []TemplateResourcePermission `pulumi:"permissions"`
	SourceEntity       *TemplateSourceEntity        `pulumi:"sourceEntity"`
	Tags               []TemplateTag                `pulumi:"tags"`
	TemplateId         string                       `pulumi:"templateId"`
	ValidationStrategy *TemplateValidationStrategy  `pulumi:"validationStrategy"`
	VersionDescription *string                      `pulumi:"versionDescription"`
}

// The set of arguments for constructing a Template resource.
type TemplateArgs struct {
	AwsAccountId       pulumi.StringInput
	Definition         TemplateVersionDefinitionPtrInput
	Name               pulumi.StringPtrInput
	Permissions        TemplateResourcePermissionArrayInput
	SourceEntity       TemplateSourceEntityPtrInput
	Tags               TemplateTagArrayInput
	TemplateId         pulumi.StringInput
	ValidationStrategy TemplateValidationStrategyPtrInput
	VersionDescription pulumi.StringPtrInput
}

func (TemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArgs)(nil)).Elem()
}

type TemplateInput interface {
	pulumi.Input

	ToTemplateOutput() TemplateOutput
	ToTemplateOutputWithContext(ctx context.Context) TemplateOutput
}

func (*Template) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (i *Template) ToTemplateOutput() TemplateOutput {
	return i.ToTemplateOutputWithContext(context.Background())
}

func (i *Template) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateOutput)
}

func (i *Template) ToOutput(ctx context.Context) pulumix.Output[*Template] {
	return pulumix.Output[*Template]{
		OutputState: i.ToTemplateOutputWithContext(ctx).OutputState,
	}
}

type TemplateOutput struct{ *pulumi.OutputState }

func (TemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (o TemplateOutput) ToTemplateOutput() TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return o
}

func (o TemplateOutput) ToOutput(ctx context.Context) pulumix.Output[*Template] {
	return pulumix.Output[*Template]{
		OutputState: o.OutputState,
	}
}

func (o TemplateOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o TemplateOutput) AwsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.AwsAccountId }).(pulumi.StringOutput)
}

func (o TemplateOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o TemplateOutput) Definition() TemplateVersionDefinitionPtrOutput {
	return o.ApplyT(func(v *Template) TemplateVersionDefinitionPtrOutput { return v.Definition }).(TemplateVersionDefinitionPtrOutput)
}

func (o TemplateOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o TemplateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Template) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TemplateOutput) Permissions() TemplateResourcePermissionArrayOutput {
	return o.ApplyT(func(v *Template) TemplateResourcePermissionArrayOutput { return v.Permissions }).(TemplateResourcePermissionArrayOutput)
}

func (o TemplateOutput) SourceEntity() TemplateSourceEntityPtrOutput {
	return o.ApplyT(func(v *Template) TemplateSourceEntityPtrOutput { return v.SourceEntity }).(TemplateSourceEntityPtrOutput)
}

func (o TemplateOutput) Tags() TemplateTagArrayOutput {
	return o.ApplyT(func(v *Template) TemplateTagArrayOutput { return v.Tags }).(TemplateTagArrayOutput)
}

func (o TemplateOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

func (o TemplateOutput) ValidationStrategy() TemplateValidationStrategyPtrOutput {
	return o.ApplyT(func(v *Template) TemplateValidationStrategyPtrOutput { return v.ValidationStrategy }).(TemplateValidationStrategyPtrOutput)
}

func (o TemplateOutput) Version() TemplateVersionOutput {
	return o.ApplyT(func(v *Template) TemplateVersionOutput { return v.Version }).(TemplateVersionOutput)
}

func (o TemplateOutput) VersionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Template) pulumi.StringPtrOutput { return v.VersionDescription }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateInput)(nil)).Elem(), &Template{})
	pulumi.RegisterOutputType(TemplateOutput{})
}
