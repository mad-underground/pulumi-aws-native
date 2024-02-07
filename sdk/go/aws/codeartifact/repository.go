// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codeartifact

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The resource schema to create a CodeArtifact repository.
type Repository struct {
	pulumi.CustomResourceState

	// The ARN of the repository.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A text description of the repository.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the domain that contains the repository.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The 12-digit account ID of the AWS account that owns the domain.
	DomainOwner pulumi.StringOutput `pulumi:"domainOwner"`
	// A list of external connections associated with the repository.
	ExternalConnections pulumi.StringArrayOutput `pulumi:"externalConnections"`
	// The name of the repository. This is used for GetAtt
	Name pulumi.StringOutput `pulumi:"name"`
	// The access control resource policy on the provided repository.
	PermissionsPolicyDocument pulumi.AnyOutput `pulumi:"permissionsPolicyDocument"`
	// The name of the repository.
	RepositoryName pulumi.StringOutput `pulumi:"repositoryName"`
	// An array of key-value pairs to apply to this resource.
	Tags RepositoryTagArrayOutput `pulumi:"tags"`
	// A list of upstream repositories associated with the repository.
	Upstreams pulumi.StringArrayOutput `pulumi:"upstreams"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"domainName",
		"domainOwner",
		"repositoryName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Repository
	err := ctx.RegisterResource("aws-native:codeartifact:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("aws-native:codeartifact:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
}

type RepositoryState struct {
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// A text description of the repository.
	Description *string `pulumi:"description"`
	// The name of the domain that contains the repository.
	DomainName string `pulumi:"domainName"`
	// The 12-digit account ID of the AWS account that owns the domain.
	DomainOwner *string `pulumi:"domainOwner"`
	// A list of external connections associated with the repository.
	ExternalConnections []string `pulumi:"externalConnections"`
	// The access control resource policy on the provided repository.
	PermissionsPolicyDocument interface{} `pulumi:"permissionsPolicyDocument"`
	// The name of the repository.
	RepositoryName *string `pulumi:"repositoryName"`
	// An array of key-value pairs to apply to this resource.
	Tags []RepositoryTag `pulumi:"tags"`
	// A list of upstream repositories associated with the repository.
	Upstreams []string `pulumi:"upstreams"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// A text description of the repository.
	Description pulumi.StringPtrInput
	// The name of the domain that contains the repository.
	DomainName pulumi.StringInput
	// The 12-digit account ID of the AWS account that owns the domain.
	DomainOwner pulumi.StringPtrInput
	// A list of external connections associated with the repository.
	ExternalConnections pulumi.StringArrayInput
	// The access control resource policy on the provided repository.
	PermissionsPolicyDocument pulumi.Input
	// The name of the repository.
	RepositoryName pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags RepositoryTagArrayInput
	// A list of upstream repositories associated with the repository.
	Upstreams pulumi.StringArrayInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (*Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (i *Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

type RepositoryOutput struct{ *pulumi.OutputState }

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

// The ARN of the repository.
func (o RepositoryOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A text description of the repository.
func (o RepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the domain that contains the repository.
func (o RepositoryOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The 12-digit account ID of the AWS account that owns the domain.
func (o RepositoryOutput) DomainOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.DomainOwner }).(pulumi.StringOutput)
}

// A list of external connections associated with the repository.
func (o RepositoryOutput) ExternalConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringArrayOutput { return v.ExternalConnections }).(pulumi.StringArrayOutput)
}

// The name of the repository. This is used for GetAtt
func (o RepositoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The access control resource policy on the provided repository.
func (o RepositoryOutput) PermissionsPolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v *Repository) pulumi.AnyOutput { return v.PermissionsPolicyDocument }).(pulumi.AnyOutput)
}

// The name of the repository.
func (o RepositoryOutput) RepositoryName() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.RepositoryName }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this resource.
func (o RepositoryOutput) Tags() RepositoryTagArrayOutput {
	return o.ApplyT(func(v *Repository) RepositoryTagArrayOutput { return v.Tags }).(RepositoryTagArrayOutput)
}

// A list of upstream repositories associated with the repository.
func (o RepositoryOutput) Upstreams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringArrayOutput { return v.Upstreams }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryInput)(nil)).Elem(), &Repository{})
	pulumi.RegisterOutputType(RepositoryOutput{})
}
