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

// Resource Type definition for AWS::ServiceCatalog::TagOptionAssociation
//
// Deprecated: TagOptionAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type TagOptionAssociation struct {
	pulumi.CustomResourceState

	AwsId       pulumi.StringOutput `pulumi:"awsId"`
	ResourceId  pulumi.StringOutput `pulumi:"resourceId"`
	TagOptionId pulumi.StringOutput `pulumi:"tagOptionId"`
}

// NewTagOptionAssociation registers a new resource with the given unique name, arguments, and options.
func NewTagOptionAssociation(ctx *pulumi.Context,
	name string, args *TagOptionAssociationArgs, opts ...pulumi.ResourceOption) (*TagOptionAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.TagOptionId == nil {
		return nil, errors.New("invalid value for required argument 'TagOptionId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"resourceId",
		"tagOptionId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TagOptionAssociation
	err := ctx.RegisterResource("aws-native:servicecatalog:TagOptionAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagOptionAssociation gets an existing TagOptionAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagOptionAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagOptionAssociationState, opts ...pulumi.ResourceOption) (*TagOptionAssociation, error) {
	var resource TagOptionAssociation
	err := ctx.ReadResource("aws-native:servicecatalog:TagOptionAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagOptionAssociation resources.
type tagOptionAssociationState struct {
}

type TagOptionAssociationState struct {
}

func (TagOptionAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagOptionAssociationState)(nil)).Elem()
}

type tagOptionAssociationArgs struct {
	ResourceId  string `pulumi:"resourceId"`
	TagOptionId string `pulumi:"tagOptionId"`
}

// The set of arguments for constructing a TagOptionAssociation resource.
type TagOptionAssociationArgs struct {
	ResourceId  pulumi.StringInput
	TagOptionId pulumi.StringInput
}

func (TagOptionAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagOptionAssociationArgs)(nil)).Elem()
}

type TagOptionAssociationInput interface {
	pulumi.Input

	ToTagOptionAssociationOutput() TagOptionAssociationOutput
	ToTagOptionAssociationOutputWithContext(ctx context.Context) TagOptionAssociationOutput
}

func (*TagOptionAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**TagOptionAssociation)(nil)).Elem()
}

func (i *TagOptionAssociation) ToTagOptionAssociationOutput() TagOptionAssociationOutput {
	return i.ToTagOptionAssociationOutputWithContext(context.Background())
}

func (i *TagOptionAssociation) ToTagOptionAssociationOutputWithContext(ctx context.Context) TagOptionAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOptionAssociationOutput)
}

type TagOptionAssociationOutput struct{ *pulumi.OutputState }

func (TagOptionAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagOptionAssociation)(nil)).Elem()
}

func (o TagOptionAssociationOutput) ToTagOptionAssociationOutput() TagOptionAssociationOutput {
	return o
}

func (o TagOptionAssociationOutput) ToTagOptionAssociationOutputWithContext(ctx context.Context) TagOptionAssociationOutput {
	return o
}

func (o TagOptionAssociationOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *TagOptionAssociation) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o TagOptionAssociationOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TagOptionAssociation) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

func (o TagOptionAssociationOutput) TagOptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *TagOptionAssociation) pulumi.StringOutput { return v.TagOptionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagOptionAssociationInput)(nil)).Elem(), &TagOptionAssociation{})
	pulumi.RegisterOutputType(TagOptionAssociationOutput{})
}
