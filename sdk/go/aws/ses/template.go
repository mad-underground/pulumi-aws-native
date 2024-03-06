// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SES::Template
type Template struct {
	pulumi.CustomResourceState

	AwsId    pulumi.StringOutput   `pulumi:"awsId"`
	Template TemplateTypePtrOutput `pulumi:"template"`
}

// NewTemplate registers a new resource with the given unique name, arguments, and options.
func NewTemplate(ctx *pulumi.Context,
	name string, args *TemplateArgs, opts ...pulumi.ResourceOption) (*Template, error) {
	if args == nil {
		args = &TemplateArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"template.templateName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Template
	err := ctx.RegisterResource("aws-native:ses:Template", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:ses:Template", name, id, state, &resource, opts...)
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
	Template *TemplateType `pulumi:"template"`
}

// The set of arguments for constructing a Template resource.
type TemplateArgs struct {
	Template TemplateTypePtrInput
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

func (o TemplateOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o TemplateOutput) Template() TemplateTypePtrOutput {
	return o.ApplyT(func(v *Template) TemplateTypePtrOutput { return v.Template }).(TemplateTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateInput)(nil)).Elem(), &Template{})
	pulumi.RegisterOutputType(TemplateOutput{})
}
