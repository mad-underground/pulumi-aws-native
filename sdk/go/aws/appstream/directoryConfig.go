// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppStream::DirectoryConfig
type DirectoryConfig struct {
	pulumi.CustomResourceState

	CertificateBasedAuthProperties       DirectoryConfigCertificateBasedAuthPropertiesPtrOutput `pulumi:"certificateBasedAuthProperties"`
	DirectoryName                        pulumi.StringOutput                                    `pulumi:"directoryName"`
	OrganizationalUnitDistinguishedNames pulumi.StringArrayOutput                               `pulumi:"organizationalUnitDistinguishedNames"`
	ServiceAccountCredentials            DirectoryConfigServiceAccountCredentialsOutput         `pulumi:"serviceAccountCredentials"`
}

// NewDirectoryConfig registers a new resource with the given unique name, arguments, and options.
func NewDirectoryConfig(ctx *pulumi.Context,
	name string, args *DirectoryConfigArgs, opts ...pulumi.ResourceOption) (*DirectoryConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DirectoryName == nil {
		return nil, errors.New("invalid value for required argument 'DirectoryName'")
	}
	if args.OrganizationalUnitDistinguishedNames == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationalUnitDistinguishedNames'")
	}
	if args.ServiceAccountCredentials == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountCredentials'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DirectoryConfig
	err := ctx.RegisterResource("aws-native:appstream:DirectoryConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDirectoryConfig gets an existing DirectoryConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDirectoryConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DirectoryConfigState, opts ...pulumi.ResourceOption) (*DirectoryConfig, error) {
	var resource DirectoryConfig
	err := ctx.ReadResource("aws-native:appstream:DirectoryConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DirectoryConfig resources.
type directoryConfigState struct {
}

type DirectoryConfigState struct {
}

func (DirectoryConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryConfigState)(nil)).Elem()
}

type directoryConfigArgs struct {
	CertificateBasedAuthProperties       *DirectoryConfigCertificateBasedAuthProperties `pulumi:"certificateBasedAuthProperties"`
	DirectoryName                        string                                         `pulumi:"directoryName"`
	OrganizationalUnitDistinguishedNames []string                                       `pulumi:"organizationalUnitDistinguishedNames"`
	ServiceAccountCredentials            DirectoryConfigServiceAccountCredentials       `pulumi:"serviceAccountCredentials"`
}

// The set of arguments for constructing a DirectoryConfig resource.
type DirectoryConfigArgs struct {
	CertificateBasedAuthProperties       DirectoryConfigCertificateBasedAuthPropertiesPtrInput
	DirectoryName                        pulumi.StringInput
	OrganizationalUnitDistinguishedNames pulumi.StringArrayInput
	ServiceAccountCredentials            DirectoryConfigServiceAccountCredentialsInput
}

func (DirectoryConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryConfigArgs)(nil)).Elem()
}

type DirectoryConfigInput interface {
	pulumi.Input

	ToDirectoryConfigOutput() DirectoryConfigOutput
	ToDirectoryConfigOutputWithContext(ctx context.Context) DirectoryConfigOutput
}

func (*DirectoryConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryConfig)(nil)).Elem()
}

func (i *DirectoryConfig) ToDirectoryConfigOutput() DirectoryConfigOutput {
	return i.ToDirectoryConfigOutputWithContext(context.Background())
}

func (i *DirectoryConfig) ToDirectoryConfigOutputWithContext(ctx context.Context) DirectoryConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryConfigOutput)
}

type DirectoryConfigOutput struct{ *pulumi.OutputState }

func (DirectoryConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryConfig)(nil)).Elem()
}

func (o DirectoryConfigOutput) ToDirectoryConfigOutput() DirectoryConfigOutput {
	return o
}

func (o DirectoryConfigOutput) ToDirectoryConfigOutputWithContext(ctx context.Context) DirectoryConfigOutput {
	return o
}

func (o DirectoryConfigOutput) CertificateBasedAuthProperties() DirectoryConfigCertificateBasedAuthPropertiesPtrOutput {
	return o.ApplyT(func(v *DirectoryConfig) DirectoryConfigCertificateBasedAuthPropertiesPtrOutput {
		return v.CertificateBasedAuthProperties
	}).(DirectoryConfigCertificateBasedAuthPropertiesPtrOutput)
}

func (o DirectoryConfigOutput) DirectoryName() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryConfig) pulumi.StringOutput { return v.DirectoryName }).(pulumi.StringOutput)
}

func (o DirectoryConfigOutput) OrganizationalUnitDistinguishedNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DirectoryConfig) pulumi.StringArrayOutput { return v.OrganizationalUnitDistinguishedNames }).(pulumi.StringArrayOutput)
}

func (o DirectoryConfigOutput) ServiceAccountCredentials() DirectoryConfigServiceAccountCredentialsOutput {
	return o.ApplyT(func(v *DirectoryConfig) DirectoryConfigServiceAccountCredentialsOutput {
		return v.ServiceAccountCredentials
	}).(DirectoryConfigServiceAccountCredentialsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryConfigInput)(nil)).Elem(), &DirectoryConfig{})
	pulumi.RegisterOutputType(DirectoryConfigOutput{})
}
