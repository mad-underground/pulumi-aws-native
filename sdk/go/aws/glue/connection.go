// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Glue::Connection
//
// Deprecated: Connection is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Connection struct {
	pulumi.CustomResourceState

	AwsId           pulumi.StringOutput       `pulumi:"awsId"`
	CatalogId       pulumi.StringOutput       `pulumi:"catalogId"`
	ConnectionInput ConnectionInputTypeOutput `pulumi:"connectionInput"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CatalogId == nil {
		return nil, errors.New("invalid value for required argument 'CatalogId'")
	}
	if args.ConnectionInput == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionInput'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"catalogId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Connection
	err := ctx.RegisterResource("aws-native:glue:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("aws-native:glue:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
}

type ConnectionState struct {
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	CatalogId       string              `pulumi:"catalogId"`
	ConnectionInput ConnectionInputType `pulumi:"connectionInput"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	CatalogId       pulumi.StringInput
	ConnectionInput ConnectionInputTypeInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func (o ConnectionOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o ConnectionOutput) CatalogId() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.CatalogId }).(pulumi.StringOutput)
}

func (o ConnectionOutput) ConnectionInput() ConnectionInputTypeOutput {
	return o.ApplyT(func(v *Connection) ConnectionInputTypeOutput { return v.ConnectionInput }).(ConnectionInputTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterOutputType(ConnectionOutput{})
}
