// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::EmailTemplate
//
// Deprecated: EmailTemplate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type EmailTemplate struct {
	pulumi.CustomResourceState

	Arn                  pulumi.StringOutput    `pulumi:"arn"`
	AwsId                pulumi.StringOutput    `pulumi:"awsId"`
	DefaultSubstitutions pulumi.StringPtrOutput `pulumi:"defaultSubstitutions"`
	HtmlPart             pulumi.StringPtrOutput `pulumi:"htmlPart"`
	Subject              pulumi.StringOutput    `pulumi:"subject"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Pinpoint::EmailTemplate` for more information about the expected schema for this property.
	Tags                pulumi.AnyOutput       `pulumi:"tags"`
	TemplateDescription pulumi.StringPtrOutput `pulumi:"templateDescription"`
	TemplateName        pulumi.StringOutput    `pulumi:"templateName"`
	TextPart            pulumi.StringPtrOutput `pulumi:"textPart"`
}

// NewEmailTemplate registers a new resource with the given unique name, arguments, and options.
func NewEmailTemplate(ctx *pulumi.Context,
	name string, args *EmailTemplateArgs, opts ...pulumi.ResourceOption) (*EmailTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Subject == nil {
		return nil, errors.New("invalid value for required argument 'Subject'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"templateName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EmailTemplate
	err := ctx.RegisterResource("aws-native:pinpoint:EmailTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailTemplate gets an existing EmailTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailTemplateState, opts ...pulumi.ResourceOption) (*EmailTemplate, error) {
	var resource EmailTemplate
	err := ctx.ReadResource("aws-native:pinpoint:EmailTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailTemplate resources.
type emailTemplateState struct {
}

type EmailTemplateState struct {
}

func (EmailTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailTemplateState)(nil)).Elem()
}

type emailTemplateArgs struct {
	DefaultSubstitutions *string `pulumi:"defaultSubstitutions"`
	HtmlPart             *string `pulumi:"htmlPart"`
	Subject              string  `pulumi:"subject"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Pinpoint::EmailTemplate` for more information about the expected schema for this property.
	Tags                interface{} `pulumi:"tags"`
	TemplateDescription *string     `pulumi:"templateDescription"`
	TemplateName        string      `pulumi:"templateName"`
	TextPart            *string     `pulumi:"textPart"`
}

// The set of arguments for constructing a EmailTemplate resource.
type EmailTemplateArgs struct {
	DefaultSubstitutions pulumi.StringPtrInput
	HtmlPart             pulumi.StringPtrInput
	Subject              pulumi.StringInput
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Pinpoint::EmailTemplate` for more information about the expected schema for this property.
	Tags                pulumi.Input
	TemplateDescription pulumi.StringPtrInput
	TemplateName        pulumi.StringInput
	TextPart            pulumi.StringPtrInput
}

func (EmailTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailTemplateArgs)(nil)).Elem()
}

type EmailTemplateInput interface {
	pulumi.Input

	ToEmailTemplateOutput() EmailTemplateOutput
	ToEmailTemplateOutputWithContext(ctx context.Context) EmailTemplateOutput
}

func (*EmailTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailTemplate)(nil)).Elem()
}

func (i *EmailTemplate) ToEmailTemplateOutput() EmailTemplateOutput {
	return i.ToEmailTemplateOutputWithContext(context.Background())
}

func (i *EmailTemplate) ToEmailTemplateOutputWithContext(ctx context.Context) EmailTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailTemplateOutput)
}

type EmailTemplateOutput struct{ *pulumi.OutputState }

func (EmailTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailTemplate)(nil)).Elem()
}

func (o EmailTemplateOutput) ToEmailTemplateOutput() EmailTemplateOutput {
	return o
}

func (o EmailTemplateOutput) ToEmailTemplateOutputWithContext(ctx context.Context) EmailTemplateOutput {
	return o
}

func (o EmailTemplateOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o EmailTemplateOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o EmailTemplateOutput) DefaultSubstitutions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringPtrOutput { return v.DefaultSubstitutions }).(pulumi.StringPtrOutput)
}

func (o EmailTemplateOutput) HtmlPart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringPtrOutput { return v.HtmlPart }).(pulumi.StringPtrOutput)
}

func (o EmailTemplateOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringOutput { return v.Subject }).(pulumi.StringOutput)
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Pinpoint::EmailTemplate` for more information about the expected schema for this property.
func (o EmailTemplateOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func (o EmailTemplateOutput) TemplateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringPtrOutput { return v.TemplateDescription }).(pulumi.StringPtrOutput)
}

func (o EmailTemplateOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

func (o EmailTemplateOutput) TextPart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailTemplate) pulumi.StringPtrOutput { return v.TextPart }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmailTemplateInput)(nil)).Elem(), &EmailTemplate{})
	pulumi.RegisterOutputType(EmailTemplateOutput{})
}
