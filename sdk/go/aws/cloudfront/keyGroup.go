// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keygroup.html
type KeyGroup struct {
	pulumi.CustomResourceState

	Id pulumi.StringOutput `pulumi:"id"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keygroup.html#cfn-cloudfront-keygroup-keygroupconfig
	KeyGroupConfig   KeyGroupKeyGroupConfigOutput `pulumi:"keyGroupConfig"`
	LastModifiedTime pulumi.StringOutput          `pulumi:"lastModifiedTime"`
}

// NewKeyGroup registers a new resource with the given unique name, arguments, and options.
func NewKeyGroup(ctx *pulumi.Context,
	name string, args *KeyGroupArgs, opts ...pulumi.ResourceOption) (*KeyGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyGroupConfig == nil {
		return nil, errors.New("invalid value for required argument 'KeyGroupConfig'")
	}
	var resource KeyGroup
	err := ctx.RegisterResource("aws-native:cloudfront:KeyGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyGroup gets an existing KeyGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyGroupState, opts ...pulumi.ResourceOption) (*KeyGroup, error) {
	var resource KeyGroup
	err := ctx.ReadResource("aws-native:cloudfront:KeyGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyGroup resources.
type keyGroupState struct {
}

type KeyGroupState struct {
}

func (KeyGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyGroupState)(nil)).Elem()
}

type keyGroupArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keygroup.html#cfn-cloudfront-keygroup-keygroupconfig
	KeyGroupConfig KeyGroupKeyGroupConfig `pulumi:"keyGroupConfig"`
}

// The set of arguments for constructing a KeyGroup resource.
type KeyGroupArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keygroup.html#cfn-cloudfront-keygroup-keygroupconfig
	KeyGroupConfig KeyGroupKeyGroupConfigInput
}

func (KeyGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyGroupArgs)(nil)).Elem()
}

type KeyGroupInput interface {
	pulumi.Input

	ToKeyGroupOutput() KeyGroupOutput
	ToKeyGroupOutputWithContext(ctx context.Context) KeyGroupOutput
}

func (*KeyGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyGroup)(nil))
}

func (i *KeyGroup) ToKeyGroupOutput() KeyGroupOutput {
	return i.ToKeyGroupOutputWithContext(context.Background())
}

func (i *KeyGroup) ToKeyGroupOutputWithContext(ctx context.Context) KeyGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyGroupOutput)
}

type KeyGroupOutput struct{ *pulumi.OutputState }

func (KeyGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyGroup)(nil))
}

func (o KeyGroupOutput) ToKeyGroupOutput() KeyGroupOutput {
	return o
}

func (o KeyGroupOutput) ToKeyGroupOutputWithContext(ctx context.Context) KeyGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KeyGroupOutput{})
}