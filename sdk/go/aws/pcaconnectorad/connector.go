// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pcaconnectorad

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::PCAConnectorAD::Connector Resource Type
type Connector struct {
	pulumi.CustomResourceState

	CertificateAuthorityArn pulumi.StringOutput           `pulumi:"certificateAuthorityArn"`
	ConnectorArn            pulumi.StringOutput           `pulumi:"connectorArn"`
	DirectoryId             pulumi.StringOutput           `pulumi:"directoryId"`
	Tags                    pulumi.StringMapOutput        `pulumi:"tags"`
	VpcInformation          ConnectorVpcInformationOutput `pulumi:"vpcInformation"`
}

// NewConnector registers a new resource with the given unique name, arguments, and options.
func NewConnector(ctx *pulumi.Context,
	name string, args *ConnectorArgs, opts ...pulumi.ResourceOption) (*Connector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateAuthorityArn == nil {
		return nil, errors.New("invalid value for required argument 'CertificateAuthorityArn'")
	}
	if args.DirectoryId == nil {
		return nil, errors.New("invalid value for required argument 'DirectoryId'")
	}
	if args.VpcInformation == nil {
		return nil, errors.New("invalid value for required argument 'VpcInformation'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"certificateAuthorityArn",
		"directoryId",
		"vpcInformation",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Connector
	err := ctx.RegisterResource("aws-native:pcaconnectorad:Connector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnector gets an existing Connector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorState, opts ...pulumi.ResourceOption) (*Connector, error) {
	var resource Connector
	err := ctx.ReadResource("aws-native:pcaconnectorad:Connector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connector resources.
type connectorState struct {
}

type ConnectorState struct {
}

func (ConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorState)(nil)).Elem()
}

type connectorArgs struct {
	CertificateAuthorityArn string                  `pulumi:"certificateAuthorityArn"`
	DirectoryId             string                  `pulumi:"directoryId"`
	Tags                    map[string]string       `pulumi:"tags"`
	VpcInformation          ConnectorVpcInformation `pulumi:"vpcInformation"`
}

// The set of arguments for constructing a Connector resource.
type ConnectorArgs struct {
	CertificateAuthorityArn pulumi.StringInput
	DirectoryId             pulumi.StringInput
	Tags                    pulumi.StringMapInput
	VpcInformation          ConnectorVpcInformationInput
}

func (ConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorArgs)(nil)).Elem()
}

type ConnectorInput interface {
	pulumi.Input

	ToConnectorOutput() ConnectorOutput
	ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput
}

func (*Connector) ElementType() reflect.Type {
	return reflect.TypeOf((**Connector)(nil)).Elem()
}

func (i *Connector) ToConnectorOutput() ConnectorOutput {
	return i.ToConnectorOutputWithContext(context.Background())
}

func (i *Connector) ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorOutput)
}

type ConnectorOutput struct{ *pulumi.OutputState }

func (ConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connector)(nil)).Elem()
}

func (o ConnectorOutput) ToConnectorOutput() ConnectorOutput {
	return o
}

func (o ConnectorOutput) ToConnectorOutputWithContext(ctx context.Context) ConnectorOutput {
	return o
}

func (o ConnectorOutput) CertificateAuthorityArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.CertificateAuthorityArn }).(pulumi.StringOutput)
}

func (o ConnectorOutput) ConnectorArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.ConnectorArn }).(pulumi.StringOutput)
}

func (o ConnectorOutput) DirectoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringOutput { return v.DirectoryId }).(pulumi.StringOutput)
}

func (o ConnectorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Connector) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ConnectorOutput) VpcInformation() ConnectorVpcInformationOutput {
	return o.ApplyT(func(v *Connector) ConnectorVpcInformationOutput { return v.VpcInformation }).(ConnectorVpcInformationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectorInput)(nil)).Elem(), &Connector{})
	pulumi.RegisterOutputType(ConnectorOutput{})
}
