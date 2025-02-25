// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IAM::ServerCertificate
type ServerCertificate struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the server certificate
	Arn                   pulumi.StringOutput    `pulumi:"arn"`
	CertificateBody       pulumi.StringPtrOutput `pulumi:"certificateBody"`
	CertificateChain      pulumi.StringPtrOutput `pulumi:"certificateChain"`
	Path                  pulumi.StringPtrOutput `pulumi:"path"`
	PrivateKey            pulumi.StringPtrOutput `pulumi:"privateKey"`
	ServerCertificateName pulumi.StringPtrOutput `pulumi:"serverCertificateName"`
	Tags                  aws.TagArrayOutput     `pulumi:"tags"`
}

// NewServerCertificate registers a new resource with the given unique name, arguments, and options.
func NewServerCertificate(ctx *pulumi.Context,
	name string, args *ServerCertificateArgs, opts ...pulumi.ResourceOption) (*ServerCertificate, error) {
	if args == nil {
		args = &ServerCertificateArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"certificateBody",
		"certificateChain",
		"privateKey",
		"serverCertificateName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServerCertificate
	err := ctx.RegisterResource("aws-native:iam:ServerCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerCertificate gets an existing ServerCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerCertificateState, opts ...pulumi.ResourceOption) (*ServerCertificate, error) {
	var resource ServerCertificate
	err := ctx.ReadResource("aws-native:iam:ServerCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerCertificate resources.
type serverCertificateState struct {
}

type ServerCertificateState struct {
}

func (ServerCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverCertificateState)(nil)).Elem()
}

type serverCertificateArgs struct {
	CertificateBody       *string   `pulumi:"certificateBody"`
	CertificateChain      *string   `pulumi:"certificateChain"`
	Path                  *string   `pulumi:"path"`
	PrivateKey            *string   `pulumi:"privateKey"`
	ServerCertificateName *string   `pulumi:"serverCertificateName"`
	Tags                  []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a ServerCertificate resource.
type ServerCertificateArgs struct {
	CertificateBody       pulumi.StringPtrInput
	CertificateChain      pulumi.StringPtrInput
	Path                  pulumi.StringPtrInput
	PrivateKey            pulumi.StringPtrInput
	ServerCertificateName pulumi.StringPtrInput
	Tags                  aws.TagArrayInput
}

func (ServerCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverCertificateArgs)(nil)).Elem()
}

type ServerCertificateInput interface {
	pulumi.Input

	ToServerCertificateOutput() ServerCertificateOutput
	ToServerCertificateOutputWithContext(ctx context.Context) ServerCertificateOutput
}

func (*ServerCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerCertificate)(nil)).Elem()
}

func (i *ServerCertificate) ToServerCertificateOutput() ServerCertificateOutput {
	return i.ToServerCertificateOutputWithContext(context.Background())
}

func (i *ServerCertificate) ToServerCertificateOutputWithContext(ctx context.Context) ServerCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerCertificateOutput)
}

type ServerCertificateOutput struct{ *pulumi.OutputState }

func (ServerCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerCertificate)(nil)).Elem()
}

func (o ServerCertificateOutput) ToServerCertificateOutput() ServerCertificateOutput {
	return o
}

func (o ServerCertificateOutput) ToServerCertificateOutputWithContext(ctx context.Context) ServerCertificateOutput {
	return o
}

// Amazon Resource Name (ARN) of the server certificate
func (o ServerCertificateOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ServerCertificateOutput) CertificateBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringPtrOutput { return v.CertificateBody }).(pulumi.StringPtrOutput)
}

func (o ServerCertificateOutput) CertificateChain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringPtrOutput { return v.CertificateChain }).(pulumi.StringPtrOutput)
}

func (o ServerCertificateOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ServerCertificateOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringPtrOutput { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o ServerCertificateOutput) ServerCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerCertificate) pulumi.StringPtrOutput { return v.ServerCertificateName }).(pulumi.StringPtrOutput)
}

func (o ServerCertificateOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *ServerCertificate) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerCertificateInput)(nil)).Elem(), &ServerCertificate{})
	pulumi.RegisterOutputType(ServerCertificateOutput{})
}
