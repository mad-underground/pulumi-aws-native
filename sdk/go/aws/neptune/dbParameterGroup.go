// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Neptune::DBParameterGroup
//
// Deprecated: DbParameterGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type DbParameterGroup struct {
	pulumi.CustomResourceState

	AwsId       pulumi.StringOutput    `pulumi:"awsId"`
	Description pulumi.StringOutput    `pulumi:"description"`
	Family      pulumi.StringOutput    `pulumi:"family"`
	Name        pulumi.StringPtrOutput `pulumi:"name"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Neptune::DBParameterGroup` for more information about the expected schema for this property.
	Parameters pulumi.AnyOutput   `pulumi:"parameters"`
	Tags       aws.TagArrayOutput `pulumi:"tags"`
}

// NewDbParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewDbParameterGroup(ctx *pulumi.Context,
	name string, args *DbParameterGroupArgs, opts ...pulumi.ResourceOption) (*DbParameterGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Family == nil {
		return nil, errors.New("invalid value for required argument 'Family'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"description",
		"family",
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DbParameterGroup
	err := ctx.RegisterResource("aws-native:neptune:DbParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDbParameterGroup gets an existing DbParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDbParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DbParameterGroupState, opts ...pulumi.ResourceOption) (*DbParameterGroup, error) {
	var resource DbParameterGroup
	err := ctx.ReadResource("aws-native:neptune:DbParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DbParameterGroup resources.
type dbParameterGroupState struct {
}

type DbParameterGroupState struct {
}

func (DbParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbParameterGroupState)(nil)).Elem()
}

type dbParameterGroupArgs struct {
	Description string  `pulumi:"description"`
	Family      string  `pulumi:"family"`
	Name        *string `pulumi:"name"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Neptune::DBParameterGroup` for more information about the expected schema for this property.
	Parameters interface{} `pulumi:"parameters"`
	Tags       []aws.Tag   `pulumi:"tags"`
}

// The set of arguments for constructing a DbParameterGroup resource.
type DbParameterGroupArgs struct {
	Description pulumi.StringInput
	Family      pulumi.StringInput
	Name        pulumi.StringPtrInput
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Neptune::DBParameterGroup` for more information about the expected schema for this property.
	Parameters pulumi.Input
	Tags       aws.TagArrayInput
}

func (DbParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbParameterGroupArgs)(nil)).Elem()
}

type DbParameterGroupInput interface {
	pulumi.Input

	ToDbParameterGroupOutput() DbParameterGroupOutput
	ToDbParameterGroupOutputWithContext(ctx context.Context) DbParameterGroupOutput
}

func (*DbParameterGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DbParameterGroup)(nil)).Elem()
}

func (i *DbParameterGroup) ToDbParameterGroupOutput() DbParameterGroupOutput {
	return i.ToDbParameterGroupOutputWithContext(context.Background())
}

func (i *DbParameterGroup) ToDbParameterGroupOutputWithContext(ctx context.Context) DbParameterGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbParameterGroupOutput)
}

type DbParameterGroupOutput struct{ *pulumi.OutputState }

func (DbParameterGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DbParameterGroup)(nil)).Elem()
}

func (o DbParameterGroupOutput) ToDbParameterGroupOutput() DbParameterGroupOutput {
	return o
}

func (o DbParameterGroupOutput) ToDbParameterGroupOutputWithContext(ctx context.Context) DbParameterGroupOutput {
	return o
}

func (o DbParameterGroupOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *DbParameterGroup) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o DbParameterGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DbParameterGroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o DbParameterGroupOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v *DbParameterGroup) pulumi.StringOutput { return v.Family }).(pulumi.StringOutput)
}

func (o DbParameterGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DbParameterGroup) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Neptune::DBParameterGroup` for more information about the expected schema for this property.
func (o DbParameterGroupOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *DbParameterGroup) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

func (o DbParameterGroupOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *DbParameterGroup) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DbParameterGroupInput)(nil)).Elem(), &DbParameterGroup{})
	pulumi.RegisterOutputType(DbParameterGroupOutput{})
}
