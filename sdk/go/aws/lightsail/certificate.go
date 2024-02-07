// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An example resource schema demonstrating some basic constructs and validation rules.
type Certificate struct {
	pulumi.CustomResourceState

	CertificateArn pulumi.StringOutput `pulumi:"certificateArn"`
	// The name for the certificate.
	CertificateName pulumi.StringOutput `pulumi:"certificateName"`
	// The domain name (e.g., example.com ) for the certificate.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The validation status of the certificate.
	Status pulumi.StringOutput `pulumi:"status"`
	// An array of strings that specify the alternate domains (e.g., example2.com) and subdomains (e.g., blog.example.com) for the certificate.
	SubjectAlternativeNames pulumi.StringArrayOutput `pulumi:"subjectAlternativeNames"`
	// An array of key-value pairs to apply to this resource.
	Tags CertificateTagArrayOutput `pulumi:"tags"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"certificateName",
		"domainName",
		"subjectAlternativeNames[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterResource("aws-native:lightsail:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("aws-native:lightsail:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
}

type CertificateState struct {
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// The name for the certificate.
	CertificateName *string `pulumi:"certificateName"`
	// The domain name (e.g., example.com ) for the certificate.
	DomainName string `pulumi:"domainName"`
	// An array of strings that specify the alternate domains (e.g., example2.com) and subdomains (e.g., blog.example.com) for the certificate.
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
	// An array of key-value pairs to apply to this resource.
	Tags []CertificateTag `pulumi:"tags"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// The name for the certificate.
	CertificateName pulumi.StringPtrInput
	// The domain name (e.g., example.com ) for the certificate.
	DomainName pulumi.StringInput
	// An array of strings that specify the alternate domains (e.g., example2.com) and subdomains (e.g., blog.example.com) for the certificate.
	SubjectAlternativeNames pulumi.StringArrayInput
	// An array of key-value pairs to apply to this resource.
	Tags CertificateTagArrayInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func (o CertificateOutput) CertificateArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CertificateArn }).(pulumi.StringOutput)
}

// The name for the certificate.
func (o CertificateOutput) CertificateName() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CertificateName }).(pulumi.StringOutput)
}

// The domain name (e.g., example.com ) for the certificate.
func (o CertificateOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The validation status of the certificate.
func (o CertificateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// An array of strings that specify the alternate domains (e.g., example2.com) and subdomains (e.g., blog.example.com) for the certificate.
func (o CertificateOutput) SubjectAlternativeNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringArrayOutput { return v.SubjectAlternativeNames }).(pulumi.StringArrayOutput)
}

// An array of key-value pairs to apply to this resource.
func (o CertificateOutput) Tags() CertificateTagArrayOutput {
	return o.ApplyT(func(v *Certificate) CertificateTagArrayOutput { return v.Tags }).(CertificateTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterOutputType(CertificateOutput{})
}
