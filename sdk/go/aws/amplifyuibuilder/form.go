// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amplifyuibuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::AmplifyUIBuilder::Form Resource Type
type Form struct {
	pulumi.CustomResourceState

	AppId             pulumi.StringPtrOutput        `pulumi:"appId"`
	AwsId             pulumi.StringOutput           `pulumi:"awsId"`
	Cta               FormCtaPtrOutput              `pulumi:"cta"`
	DataType          FormDataTypeConfigPtrOutput   `pulumi:"dataType"`
	EnvironmentName   pulumi.StringPtrOutput        `pulumi:"environmentName"`
	Fields            FormFieldConfigMapOutput      `pulumi:"fields"`
	FormActionType    FormActionTypePtrOutput       `pulumi:"formActionType"`
	LabelDecorator    FormLabelDecoratorPtrOutput   `pulumi:"labelDecorator"`
	Name              pulumi.StringPtrOutput        `pulumi:"name"`
	SchemaVersion     pulumi.StringPtrOutput        `pulumi:"schemaVersion"`
	SectionalElements FormSectionalElementMapOutput `pulumi:"sectionalElements"`
	Style             FormStylePtrOutput            `pulumi:"style"`
	Tags              pulumi.StringMapOutput        `pulumi:"tags"`
}

// NewForm registers a new resource with the given unique name, arguments, and options.
func NewForm(ctx *pulumi.Context,
	name string, args *FormArgs, opts ...pulumi.ResourceOption) (*Form, error) {
	if args == nil {
		args = &FormArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"appId",
		"environmentName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Form
	err := ctx.RegisterResource("aws-native:amplifyuibuilder:Form", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetForm gets an existing Form resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetForm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FormState, opts ...pulumi.ResourceOption) (*Form, error) {
	var resource Form
	err := ctx.ReadResource("aws-native:amplifyuibuilder:Form", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Form resources.
type formState struct {
}

type FormState struct {
}

func (FormState) ElementType() reflect.Type {
	return reflect.TypeOf((*formState)(nil)).Elem()
}

type formArgs struct {
	AppId             *string                         `pulumi:"appId"`
	Cta               *FormCta                        `pulumi:"cta"`
	DataType          *FormDataTypeConfig             `pulumi:"dataType"`
	EnvironmentName   *string                         `pulumi:"environmentName"`
	Fields            map[string]FormFieldConfig      `pulumi:"fields"`
	FormActionType    *FormActionType                 `pulumi:"formActionType"`
	LabelDecorator    *FormLabelDecorator             `pulumi:"labelDecorator"`
	Name              *string                         `pulumi:"name"`
	SchemaVersion     *string                         `pulumi:"schemaVersion"`
	SectionalElements map[string]FormSectionalElement `pulumi:"sectionalElements"`
	Style             *FormStyle                      `pulumi:"style"`
	Tags              map[string]string               `pulumi:"tags"`
}

// The set of arguments for constructing a Form resource.
type FormArgs struct {
	AppId             pulumi.StringPtrInput
	Cta               FormCtaPtrInput
	DataType          FormDataTypeConfigPtrInput
	EnvironmentName   pulumi.StringPtrInput
	Fields            FormFieldConfigMapInput
	FormActionType    FormActionTypePtrInput
	LabelDecorator    FormLabelDecoratorPtrInput
	Name              pulumi.StringPtrInput
	SchemaVersion     pulumi.StringPtrInput
	SectionalElements FormSectionalElementMapInput
	Style             FormStylePtrInput
	Tags              pulumi.StringMapInput
}

func (FormArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*formArgs)(nil)).Elem()
}

type FormInput interface {
	pulumi.Input

	ToFormOutput() FormOutput
	ToFormOutputWithContext(ctx context.Context) FormOutput
}

func (*Form) ElementType() reflect.Type {
	return reflect.TypeOf((**Form)(nil)).Elem()
}

func (i *Form) ToFormOutput() FormOutput {
	return i.ToFormOutputWithContext(context.Background())
}

func (i *Form) ToFormOutputWithContext(ctx context.Context) FormOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormOutput)
}

type FormOutput struct{ *pulumi.OutputState }

func (FormOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Form)(nil)).Elem()
}

func (o FormOutput) ToFormOutput() FormOutput {
	return o
}

func (o FormOutput) ToFormOutputWithContext(ctx context.Context) FormOutput {
	return o
}

func (o FormOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Form) pulumi.StringPtrOutput { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o FormOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Form) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o FormOutput) Cta() FormCtaPtrOutput {
	return o.ApplyT(func(v *Form) FormCtaPtrOutput { return v.Cta }).(FormCtaPtrOutput)
}

func (o FormOutput) DataType() FormDataTypeConfigPtrOutput {
	return o.ApplyT(func(v *Form) FormDataTypeConfigPtrOutput { return v.DataType }).(FormDataTypeConfigPtrOutput)
}

func (o FormOutput) EnvironmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Form) pulumi.StringPtrOutput { return v.EnvironmentName }).(pulumi.StringPtrOutput)
}

func (o FormOutput) Fields() FormFieldConfigMapOutput {
	return o.ApplyT(func(v *Form) FormFieldConfigMapOutput { return v.Fields }).(FormFieldConfigMapOutput)
}

func (o FormOutput) FormActionType() FormActionTypePtrOutput {
	return o.ApplyT(func(v *Form) FormActionTypePtrOutput { return v.FormActionType }).(FormActionTypePtrOutput)
}

func (o FormOutput) LabelDecorator() FormLabelDecoratorPtrOutput {
	return o.ApplyT(func(v *Form) FormLabelDecoratorPtrOutput { return v.LabelDecorator }).(FormLabelDecoratorPtrOutput)
}

func (o FormOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Form) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FormOutput) SchemaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Form) pulumi.StringPtrOutput { return v.SchemaVersion }).(pulumi.StringPtrOutput)
}

func (o FormOutput) SectionalElements() FormSectionalElementMapOutput {
	return o.ApplyT(func(v *Form) FormSectionalElementMapOutput { return v.SectionalElements }).(FormSectionalElementMapOutput)
}

func (o FormOutput) Style() FormStylePtrOutput {
	return o.ApplyT(func(v *Form) FormStylePtrOutput { return v.Style }).(FormStylePtrOutput)
}

func (o FormOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Form) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FormInput)(nil)).Elem(), &Form{})
	pulumi.RegisterOutputType(FormOutput{})
}
