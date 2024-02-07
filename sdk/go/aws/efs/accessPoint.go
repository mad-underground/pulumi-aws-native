// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EFS::AccessPoint
type AccessPoint struct {
	pulumi.CustomResourceState

	AccessPointId   pulumi.StringOutput       `pulumi:"accessPointId"`
	AccessPointTags AccessPointTagArrayOutput `pulumi:"accessPointTags"`
	Arn             pulumi.StringOutput       `pulumi:"arn"`
	// (optional) A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent creation.
	ClientToken pulumi.StringPtrOutput `pulumi:"clientToken"`
	// The ID of the EFS file system that the access point provides access to.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
	// The operating system user and group applied to all file system requests made using the access point.
	PosixUser AccessPointPosixUserPtrOutput `pulumi:"posixUser"`
	// Specifies the directory on the Amazon EFS file system that the access point exposes as the root directory of your file system to NFS clients using the access point. The clients using the access point can only access the root directory and below. If the RootDirectory>Path specified does not exist, EFS creates it and applies the CreationInfo settings when a client connects to an access point. When specifying a RootDirectory, you need to provide the Path, and the CreationInfo is optional.
	RootDirectory AccessPointRootDirectoryPtrOutput `pulumi:"rootDirectory"`
}

// NewAccessPoint registers a new resource with the given unique name, arguments, and options.
func NewAccessPoint(ctx *pulumi.Context,
	name string, args *AccessPointArgs, opts ...pulumi.ResourceOption) (*AccessPoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileSystemId == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"clientToken",
		"fileSystemId",
		"posixUser",
		"rootDirectory",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccessPoint
	err := ctx.RegisterResource("aws-native:efs:AccessPoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPoint gets an existing AccessPoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPointState, opts ...pulumi.ResourceOption) (*AccessPoint, error) {
	var resource AccessPoint
	err := ctx.ReadResource("aws-native:efs:AccessPoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPoint resources.
type accessPointState struct {
}

type AccessPointState struct {
}

func (AccessPointState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPointState)(nil)).Elem()
}

type accessPointArgs struct {
	AccessPointTags []AccessPointTag `pulumi:"accessPointTags"`
	// (optional) A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent creation.
	ClientToken *string `pulumi:"clientToken"`
	// The ID of the EFS file system that the access point provides access to.
	FileSystemId string `pulumi:"fileSystemId"`
	// The operating system user and group applied to all file system requests made using the access point.
	PosixUser *AccessPointPosixUser `pulumi:"posixUser"`
	// Specifies the directory on the Amazon EFS file system that the access point exposes as the root directory of your file system to NFS clients using the access point. The clients using the access point can only access the root directory and below. If the RootDirectory>Path specified does not exist, EFS creates it and applies the CreationInfo settings when a client connects to an access point. When specifying a RootDirectory, you need to provide the Path, and the CreationInfo is optional.
	RootDirectory *AccessPointRootDirectory `pulumi:"rootDirectory"`
}

// The set of arguments for constructing a AccessPoint resource.
type AccessPointArgs struct {
	AccessPointTags AccessPointTagArrayInput
	// (optional) A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent creation.
	ClientToken pulumi.StringPtrInput
	// The ID of the EFS file system that the access point provides access to.
	FileSystemId pulumi.StringInput
	// The operating system user and group applied to all file system requests made using the access point.
	PosixUser AccessPointPosixUserPtrInput
	// Specifies the directory on the Amazon EFS file system that the access point exposes as the root directory of your file system to NFS clients using the access point. The clients using the access point can only access the root directory and below. If the RootDirectory>Path specified does not exist, EFS creates it and applies the CreationInfo settings when a client connects to an access point. When specifying a RootDirectory, you need to provide the Path, and the CreationInfo is optional.
	RootDirectory AccessPointRootDirectoryPtrInput
}

func (AccessPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPointArgs)(nil)).Elem()
}

type AccessPointInput interface {
	pulumi.Input

	ToAccessPointOutput() AccessPointOutput
	ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput
}

func (*AccessPoint) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPoint)(nil)).Elem()
}

func (i *AccessPoint) ToAccessPointOutput() AccessPointOutput {
	return i.ToAccessPointOutputWithContext(context.Background())
}

func (i *AccessPoint) ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointOutput)
}

type AccessPointOutput struct{ *pulumi.OutputState }

func (AccessPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPoint)(nil)).Elem()
}

func (o AccessPointOutput) ToAccessPointOutput() AccessPointOutput {
	return o
}

func (o AccessPointOutput) ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput {
	return o
}

func (o AccessPointOutput) AccessPointId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.AccessPointId }).(pulumi.StringOutput)
}

func (o AccessPointOutput) AccessPointTags() AccessPointTagArrayOutput {
	return o.ApplyT(func(v *AccessPoint) AccessPointTagArrayOutput { return v.AccessPointTags }).(AccessPointTagArrayOutput)
}

func (o AccessPointOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// (optional) A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent creation.
func (o AccessPointOutput) ClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringPtrOutput { return v.ClientToken }).(pulumi.StringPtrOutput)
}

// The ID of the EFS file system that the access point provides access to.
func (o AccessPointOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

// The operating system user and group applied to all file system requests made using the access point.
func (o AccessPointOutput) PosixUser() AccessPointPosixUserPtrOutput {
	return o.ApplyT(func(v *AccessPoint) AccessPointPosixUserPtrOutput { return v.PosixUser }).(AccessPointPosixUserPtrOutput)
}

// Specifies the directory on the Amazon EFS file system that the access point exposes as the root directory of your file system to NFS clients using the access point. The clients using the access point can only access the root directory and below. If the RootDirectory>Path specified does not exist, EFS creates it and applies the CreationInfo settings when a client connects to an access point. When specifying a RootDirectory, you need to provide the Path, and the CreationInfo is optional.
func (o AccessPointOutput) RootDirectory() AccessPointRootDirectoryPtrOutput {
	return o.ApplyT(func(v *AccessPoint) AccessPointRootDirectoryPtrOutput { return v.RootDirectory }).(AccessPointRootDirectoryPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPointInput)(nil)).Elem(), &AccessPoint{})
	pulumi.RegisterOutputType(AccessPointOutput{})
}
