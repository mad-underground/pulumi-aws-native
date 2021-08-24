// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package athena

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html
type NamedQuery struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-database
	Database pulumi.StringOutput `pulumi:"database"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-name
	Name         pulumi.StringPtrOutput `pulumi:"name"`
	NamedQueryId pulumi.StringOutput    `pulumi:"namedQueryId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-querystring
	QueryString pulumi.StringOutput `pulumi:"queryString"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-workgroup
	WorkGroup pulumi.StringPtrOutput `pulumi:"workGroup"`
}

// NewNamedQuery registers a new resource with the given unique name, arguments, and options.
func NewNamedQuery(ctx *pulumi.Context,
	name string, args *NamedQueryArgs, opts ...pulumi.ResourceOption) (*NamedQuery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.QueryString == nil {
		return nil, errors.New("invalid value for required argument 'QueryString'")
	}
	var resource NamedQuery
	err := ctx.RegisterResource("aws-native:athena:NamedQuery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamedQuery gets an existing NamedQuery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamedQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamedQueryState, opts ...pulumi.ResourceOption) (*NamedQuery, error) {
	var resource NamedQuery
	err := ctx.ReadResource("aws-native:athena:NamedQuery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamedQuery resources.
type namedQueryState struct {
}

type NamedQueryState struct {
}

func (NamedQueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*namedQueryState)(nil)).Elem()
}

type namedQueryArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-database
	Database string `pulumi:"database"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-name
	Name *string `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-querystring
	QueryString string `pulumi:"queryString"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-workgroup
	WorkGroup *string `pulumi:"workGroup"`
}

// The set of arguments for constructing a NamedQuery resource.
type NamedQueryArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-database
	Database pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-name
	Name pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-querystring
	QueryString pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html#cfn-athena-namedquery-workgroup
	WorkGroup pulumi.StringPtrInput
}

func (NamedQueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namedQueryArgs)(nil)).Elem()
}

type NamedQueryInput interface {
	pulumi.Input

	ToNamedQueryOutput() NamedQueryOutput
	ToNamedQueryOutputWithContext(ctx context.Context) NamedQueryOutput
}

func (*NamedQuery) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedQuery)(nil))
}

func (i *NamedQuery) ToNamedQueryOutput() NamedQueryOutput {
	return i.ToNamedQueryOutputWithContext(context.Background())
}

func (i *NamedQuery) ToNamedQueryOutputWithContext(ctx context.Context) NamedQueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedQueryOutput)
}

type NamedQueryOutput struct{ *pulumi.OutputState }

func (NamedQueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedQuery)(nil))
}

func (o NamedQueryOutput) ToNamedQueryOutput() NamedQueryOutput {
	return o
}

func (o NamedQueryOutput) ToNamedQueryOutputWithContext(ctx context.Context) NamedQueryOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NamedQueryOutput{})
}