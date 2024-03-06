// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package macie

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Macie CustomDataIdentifier resource schema
type CustomDataIdentifier struct {
	pulumi.CustomResourceState

	// Custom data identifier ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Custom data identifier ID.
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// Description of custom data identifier.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Words to be ignored.
	IgnoreWords pulumi.StringArrayOutput `pulumi:"ignoreWords"`
	// Keywords to be matched against.
	Keywords pulumi.StringArrayOutput `pulumi:"keywords"`
	// Maximum match distance.
	MaximumMatchDistance pulumi.IntPtrOutput `pulumi:"maximumMatchDistance"`
	// Name of custom data identifier.
	Name pulumi.StringOutput `pulumi:"name"`
	// Regular expression for custom data identifier.
	Regex pulumi.StringOutput `pulumi:"regex"`
	// A collection of tags associated with a resource
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewCustomDataIdentifier registers a new resource with the given unique name, arguments, and options.
func NewCustomDataIdentifier(ctx *pulumi.Context,
	name string, args *CustomDataIdentifierArgs, opts ...pulumi.ResourceOption) (*CustomDataIdentifier, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Regex == nil {
		return nil, errors.New("invalid value for required argument 'Regex'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"description",
		"ignoreWords[*]",
		"keywords[*]",
		"maximumMatchDistance",
		"name",
		"regex",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomDataIdentifier
	err := ctx.RegisterResource("aws-native:macie:CustomDataIdentifier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomDataIdentifier gets an existing CustomDataIdentifier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomDataIdentifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomDataIdentifierState, opts ...pulumi.ResourceOption) (*CustomDataIdentifier, error) {
	var resource CustomDataIdentifier
	err := ctx.ReadResource("aws-native:macie:CustomDataIdentifier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomDataIdentifier resources.
type customDataIdentifierState struct {
}

type CustomDataIdentifierState struct {
}

func (CustomDataIdentifierState) ElementType() reflect.Type {
	return reflect.TypeOf((*customDataIdentifierState)(nil)).Elem()
}

type customDataIdentifierArgs struct {
	// Description of custom data identifier.
	Description *string `pulumi:"description"`
	// Words to be ignored.
	IgnoreWords []string `pulumi:"ignoreWords"`
	// Keywords to be matched against.
	Keywords []string `pulumi:"keywords"`
	// Maximum match distance.
	MaximumMatchDistance *int `pulumi:"maximumMatchDistance"`
	// Name of custom data identifier.
	Name *string `pulumi:"name"`
	// Regular expression for custom data identifier.
	Regex string `pulumi:"regex"`
	// A collection of tags associated with a resource
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a CustomDataIdentifier resource.
type CustomDataIdentifierArgs struct {
	// Description of custom data identifier.
	Description pulumi.StringPtrInput
	// Words to be ignored.
	IgnoreWords pulumi.StringArrayInput
	// Keywords to be matched against.
	Keywords pulumi.StringArrayInput
	// Maximum match distance.
	MaximumMatchDistance pulumi.IntPtrInput
	// Name of custom data identifier.
	Name pulumi.StringPtrInput
	// Regular expression for custom data identifier.
	Regex pulumi.StringInput
	// A collection of tags associated with a resource
	Tags aws.TagArrayInput
}

func (CustomDataIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customDataIdentifierArgs)(nil)).Elem()
}

type CustomDataIdentifierInput interface {
	pulumi.Input

	ToCustomDataIdentifierOutput() CustomDataIdentifierOutput
	ToCustomDataIdentifierOutputWithContext(ctx context.Context) CustomDataIdentifierOutput
}

func (*CustomDataIdentifier) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDataIdentifier)(nil)).Elem()
}

func (i *CustomDataIdentifier) ToCustomDataIdentifierOutput() CustomDataIdentifierOutput {
	return i.ToCustomDataIdentifierOutputWithContext(context.Background())
}

func (i *CustomDataIdentifier) ToCustomDataIdentifierOutputWithContext(ctx context.Context) CustomDataIdentifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDataIdentifierOutput)
}

type CustomDataIdentifierOutput struct{ *pulumi.OutputState }

func (CustomDataIdentifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDataIdentifier)(nil)).Elem()
}

func (o CustomDataIdentifierOutput) ToCustomDataIdentifierOutput() CustomDataIdentifierOutput {
	return o
}

func (o CustomDataIdentifierOutput) ToCustomDataIdentifierOutputWithContext(ctx context.Context) CustomDataIdentifierOutput {
	return o
}

// Custom data identifier ARN.
func (o CustomDataIdentifierOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Custom data identifier ID.
func (o CustomDataIdentifierOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// Description of custom data identifier.
func (o CustomDataIdentifierOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Words to be ignored.
func (o CustomDataIdentifierOutput) IgnoreWords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringArrayOutput { return v.IgnoreWords }).(pulumi.StringArrayOutput)
}

// Keywords to be matched against.
func (o CustomDataIdentifierOutput) Keywords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringArrayOutput { return v.Keywords }).(pulumi.StringArrayOutput)
}

// Maximum match distance.
func (o CustomDataIdentifierOutput) MaximumMatchDistance() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.IntPtrOutput { return v.MaximumMatchDistance }).(pulumi.IntPtrOutput)
}

// Name of custom data identifier.
func (o CustomDataIdentifierOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Regular expression for custom data identifier.
func (o CustomDataIdentifierOutput) Regex() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringOutput { return v.Regex }).(pulumi.StringOutput)
}

// A collection of tags associated with a resource
func (o CustomDataIdentifierOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDataIdentifierInput)(nil)).Elem(), &CustomDataIdentifier{})
	pulumi.RegisterOutputType(CustomDataIdentifierOutput{})
}
