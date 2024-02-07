// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cleanrooms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a stored analysis within a collaboration
type AnalysisTemplate struct {
	pulumi.CustomResourceState

	// The member who can query can provide this placeholder for a literal data value in an analysis template
	AnalysisParameters         AnalysisTemplateAnalysisParameterArrayOutput `pulumi:"analysisParameters"`
	AnalysisTemplateIdentifier pulumi.StringOutput                          `pulumi:"analysisTemplateIdentifier"`
	Arn                        pulumi.StringOutput                          `pulumi:"arn"`
	CollaborationArn           pulumi.StringOutput                          `pulumi:"collaborationArn"`
	CollaborationIdentifier    pulumi.StringOutput                          `pulumi:"collaborationIdentifier"`
	Description                pulumi.StringPtrOutput                       `pulumi:"description"`
	Format                     AnalysisTemplateFormatOutput                 `pulumi:"format"`
	MembershipArn              pulumi.StringOutput                          `pulumi:"membershipArn"`
	MembershipIdentifier       pulumi.StringOutput                          `pulumi:"membershipIdentifier"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	Schema                     AnalysisTemplateAnalysisSchemaOutput         `pulumi:"schema"`
	Source                     AnalysisTemplateAnalysisSourceOutput         `pulumi:"source"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms analysis template.
	Tags AnalysisTemplateTagArrayOutput `pulumi:"tags"`
}

// NewAnalysisTemplate registers a new resource with the given unique name, arguments, and options.
func NewAnalysisTemplate(ctx *pulumi.Context,
	name string, args *AnalysisTemplateArgs, opts ...pulumi.ResourceOption) (*AnalysisTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Format == nil {
		return nil, errors.New("invalid value for required argument 'Format'")
	}
	if args.MembershipIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'MembershipIdentifier'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"analysisParameters[*]",
		"format",
		"membershipIdentifier",
		"name",
		"source",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AnalysisTemplate
	err := ctx.RegisterResource("aws-native:cleanrooms:AnalysisTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalysisTemplate gets an existing AnalysisTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalysisTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalysisTemplateState, opts ...pulumi.ResourceOption) (*AnalysisTemplate, error) {
	var resource AnalysisTemplate
	err := ctx.ReadResource("aws-native:cleanrooms:AnalysisTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnalysisTemplate resources.
type analysisTemplateState struct {
}

type AnalysisTemplateState struct {
}

func (AnalysisTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*analysisTemplateState)(nil)).Elem()
}

type analysisTemplateArgs struct {
	// The member who can query can provide this placeholder for a literal data value in an analysis template
	AnalysisParameters   []AnalysisTemplateAnalysisParameter `pulumi:"analysisParameters"`
	Description          *string                             `pulumi:"description"`
	Format               AnalysisTemplateFormat              `pulumi:"format"`
	MembershipIdentifier string                              `pulumi:"membershipIdentifier"`
	Name                 *string                             `pulumi:"name"`
	Source               AnalysisTemplateAnalysisSource      `pulumi:"source"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms analysis template.
	Tags []AnalysisTemplateTag `pulumi:"tags"`
}

// The set of arguments for constructing a AnalysisTemplate resource.
type AnalysisTemplateArgs struct {
	// The member who can query can provide this placeholder for a literal data value in an analysis template
	AnalysisParameters   AnalysisTemplateAnalysisParameterArrayInput
	Description          pulumi.StringPtrInput
	Format               AnalysisTemplateFormatInput
	MembershipIdentifier pulumi.StringInput
	Name                 pulumi.StringPtrInput
	Source               AnalysisTemplateAnalysisSourceInput
	// An arbitrary set of tags (key-value pairs) for this cleanrooms analysis template.
	Tags AnalysisTemplateTagArrayInput
}

func (AnalysisTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analysisTemplateArgs)(nil)).Elem()
}

type AnalysisTemplateInput interface {
	pulumi.Input

	ToAnalysisTemplateOutput() AnalysisTemplateOutput
	ToAnalysisTemplateOutputWithContext(ctx context.Context) AnalysisTemplateOutput
}

func (*AnalysisTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalysisTemplate)(nil)).Elem()
}

func (i *AnalysisTemplate) ToAnalysisTemplateOutput() AnalysisTemplateOutput {
	return i.ToAnalysisTemplateOutputWithContext(context.Background())
}

func (i *AnalysisTemplate) ToAnalysisTemplateOutputWithContext(ctx context.Context) AnalysisTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalysisTemplateOutput)
}

type AnalysisTemplateOutput struct{ *pulumi.OutputState }

func (AnalysisTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalysisTemplate)(nil)).Elem()
}

func (o AnalysisTemplateOutput) ToAnalysisTemplateOutput() AnalysisTemplateOutput {
	return o
}

func (o AnalysisTemplateOutput) ToAnalysisTemplateOutputWithContext(ctx context.Context) AnalysisTemplateOutput {
	return o
}

// The member who can query can provide this placeholder for a literal data value in an analysis template
func (o AnalysisTemplateOutput) AnalysisParameters() AnalysisTemplateAnalysisParameterArrayOutput {
	return o.ApplyT(func(v *AnalysisTemplate) AnalysisTemplateAnalysisParameterArrayOutput { return v.AnalysisParameters }).(AnalysisTemplateAnalysisParameterArrayOutput)
}

func (o AnalysisTemplateOutput) AnalysisTemplateIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalysisTemplate) pulumi.StringOutput { return v.AnalysisTemplateIdentifier }).(pulumi.StringOutput)
}

func (o AnalysisTemplateOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalysisTemplate) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o AnalysisTemplateOutput) CollaborationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalysisTemplate) pulumi.StringOutput { return v.CollaborationArn }).(pulumi.StringOutput)
}

func (o AnalysisTemplateOutput) CollaborationIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalysisTemplate) pulumi.StringOutput { return v.CollaborationIdentifier }).(pulumi.StringOutput)
}

func (o AnalysisTemplateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnalysisTemplate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AnalysisTemplateOutput) Format() AnalysisTemplateFormatOutput {
	return o.ApplyT(func(v *AnalysisTemplate) AnalysisTemplateFormatOutput { return v.Format }).(AnalysisTemplateFormatOutput)
}

func (o AnalysisTemplateOutput) MembershipArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalysisTemplate) pulumi.StringOutput { return v.MembershipArn }).(pulumi.StringOutput)
}

func (o AnalysisTemplateOutput) MembershipIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalysisTemplate) pulumi.StringOutput { return v.MembershipIdentifier }).(pulumi.StringOutput)
}

func (o AnalysisTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AnalysisTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AnalysisTemplateOutput) Schema() AnalysisTemplateAnalysisSchemaOutput {
	return o.ApplyT(func(v *AnalysisTemplate) AnalysisTemplateAnalysisSchemaOutput { return v.Schema }).(AnalysisTemplateAnalysisSchemaOutput)
}

func (o AnalysisTemplateOutput) Source() AnalysisTemplateAnalysisSourceOutput {
	return o.ApplyT(func(v *AnalysisTemplate) AnalysisTemplateAnalysisSourceOutput { return v.Source }).(AnalysisTemplateAnalysisSourceOutput)
}

// An arbitrary set of tags (key-value pairs) for this cleanrooms analysis template.
func (o AnalysisTemplateOutput) Tags() AnalysisTemplateTagArrayOutput {
	return o.ApplyT(func(v *AnalysisTemplate) AnalysisTemplateTagArrayOutput { return v.Tags }).(AnalysisTemplateTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnalysisTemplateInput)(nil)).Elem(), &AnalysisTemplate{})
	pulumi.RegisterOutputType(AnalysisTemplateOutput{})
}
