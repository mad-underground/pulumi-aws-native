// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package acmpca

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Permission set on private certificate authority
type Permission struct {
	pulumi.CustomResourceState

	// The actions that the specified AWS service principal can use. Actions IssueCertificate, GetCertificate and ListPermissions must be provided.
	Actions pulumi.StringArrayOutput `pulumi:"actions"`
	// The Amazon Resource Name (ARN) of the Private Certificate Authority that grants the permission.
	CertificateAuthorityArn pulumi.StringOutput `pulumi:"certificateAuthorityArn"`
	// The AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The ID of the calling account.
	SourceAccount pulumi.StringPtrOutput `pulumi:"sourceAccount"`
}

// NewPermission registers a new resource with the given unique name, arguments, and options.
func NewPermission(ctx *pulumi.Context,
	name string, args *PermissionArgs, opts ...pulumi.ResourceOption) (*Permission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.CertificateAuthorityArn == nil {
		return nil, errors.New("invalid value for required argument 'CertificateAuthorityArn'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Permission
	err := ctx.RegisterResource("aws-native:acmpca:Permission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermission gets an existing Permission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionState, opts ...pulumi.ResourceOption) (*Permission, error) {
	var resource Permission
	err := ctx.ReadResource("aws-native:acmpca:Permission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Permission resources.
type permissionState struct {
}

type PermissionState struct {
}

func (PermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionState)(nil)).Elem()
}

type permissionArgs struct {
	// The actions that the specified AWS service principal can use. Actions IssueCertificate, GetCertificate and ListPermissions must be provided.
	Actions []string `pulumi:"actions"`
	// The Amazon Resource Name (ARN) of the Private Certificate Authority that grants the permission.
	CertificateAuthorityArn string `pulumi:"certificateAuthorityArn"`
	// The AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com.
	Principal string `pulumi:"principal"`
	// The ID of the calling account.
	SourceAccount *string `pulumi:"sourceAccount"`
}

// The set of arguments for constructing a Permission resource.
type PermissionArgs struct {
	// The actions that the specified AWS service principal can use. Actions IssueCertificate, GetCertificate and ListPermissions must be provided.
	Actions pulumi.StringArrayInput
	// The Amazon Resource Name (ARN) of the Private Certificate Authority that grants the permission.
	CertificateAuthorityArn pulumi.StringInput
	// The AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com.
	Principal pulumi.StringInput
	// The ID of the calling account.
	SourceAccount pulumi.StringPtrInput
}

func (PermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionArgs)(nil)).Elem()
}

type PermissionInput interface {
	pulumi.Input

	ToPermissionOutput() PermissionOutput
	ToPermissionOutputWithContext(ctx context.Context) PermissionOutput
}

func (*Permission) ElementType() reflect.Type {
	return reflect.TypeOf((**Permission)(nil)).Elem()
}

func (i *Permission) ToPermissionOutput() PermissionOutput {
	return i.ToPermissionOutputWithContext(context.Background())
}

func (i *Permission) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionOutput)
}

type PermissionOutput struct{ *pulumi.OutputState }

func (PermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Permission)(nil)).Elem()
}

func (o PermissionOutput) ToPermissionOutput() PermissionOutput {
	return o
}

func (o PermissionOutput) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return o
}

// The actions that the specified AWS service principal can use. Actions IssueCertificate, GetCertificate and ListPermissions must be provided.
func (o PermissionOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringArrayOutput { return v.Actions }).(pulumi.StringArrayOutput)
}

// The Amazon Resource Name (ARN) of the Private Certificate Authority that grants the permission.
func (o PermissionOutput) CertificateAuthorityArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringOutput { return v.CertificateAuthorityArn }).(pulumi.StringOutput)
}

// The AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com.
func (o PermissionOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The ID of the calling account.
func (o PermissionOutput) SourceAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringPtrOutput { return v.SourceAccount }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionInput)(nil)).Elem(), &Permission{})
	pulumi.RegisterOutputType(PermissionOutput{})
}
