// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::CodeRepository
//
// Deprecated: CodeRepository is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type CodeRepository struct {
	pulumi.CustomResourceState

	CodeRepositoryName pulumi.StringPtrOutput        `pulumi:"codeRepositoryName"`
	GitConfig          CodeRepositoryGitConfigOutput `pulumi:"gitConfig"`
	Tags               CodeRepositoryTagArrayOutput  `pulumi:"tags"`
}

// NewCodeRepository registers a new resource with the given unique name, arguments, and options.
func NewCodeRepository(ctx *pulumi.Context,
	name string, args *CodeRepositoryArgs, opts ...pulumi.ResourceOption) (*CodeRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GitConfig == nil {
		return nil, errors.New("invalid value for required argument 'GitConfig'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CodeRepository
	err := ctx.RegisterResource("aws-native:sagemaker:CodeRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCodeRepository gets an existing CodeRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCodeRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CodeRepositoryState, opts ...pulumi.ResourceOption) (*CodeRepository, error) {
	var resource CodeRepository
	err := ctx.ReadResource("aws-native:sagemaker:CodeRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CodeRepository resources.
type codeRepositoryState struct {
}

type CodeRepositoryState struct {
}

func (CodeRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*codeRepositoryState)(nil)).Elem()
}

type codeRepositoryArgs struct {
	CodeRepositoryName *string                 `pulumi:"codeRepositoryName"`
	GitConfig          CodeRepositoryGitConfig `pulumi:"gitConfig"`
	Tags               []CodeRepositoryTag     `pulumi:"tags"`
}

// The set of arguments for constructing a CodeRepository resource.
type CodeRepositoryArgs struct {
	CodeRepositoryName pulumi.StringPtrInput
	GitConfig          CodeRepositoryGitConfigInput
	Tags               CodeRepositoryTagArrayInput
}

func (CodeRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*codeRepositoryArgs)(nil)).Elem()
}

type CodeRepositoryInput interface {
	pulumi.Input

	ToCodeRepositoryOutput() CodeRepositoryOutput
	ToCodeRepositoryOutputWithContext(ctx context.Context) CodeRepositoryOutput
}

func (*CodeRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeRepository)(nil)).Elem()
}

func (i *CodeRepository) ToCodeRepositoryOutput() CodeRepositoryOutput {
	return i.ToCodeRepositoryOutputWithContext(context.Background())
}

func (i *CodeRepository) ToCodeRepositoryOutputWithContext(ctx context.Context) CodeRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodeRepositoryOutput)
}

type CodeRepositoryOutput struct{ *pulumi.OutputState }

func (CodeRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodeRepository)(nil)).Elem()
}

func (o CodeRepositoryOutput) ToCodeRepositoryOutput() CodeRepositoryOutput {
	return o
}

func (o CodeRepositoryOutput) ToCodeRepositoryOutputWithContext(ctx context.Context) CodeRepositoryOutput {
	return o
}

func (o CodeRepositoryOutput) CodeRepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodeRepository) pulumi.StringPtrOutput { return v.CodeRepositoryName }).(pulumi.StringPtrOutput)
}

func (o CodeRepositoryOutput) GitConfig() CodeRepositoryGitConfigOutput {
	return o.ApplyT(func(v *CodeRepository) CodeRepositoryGitConfigOutput { return v.GitConfig }).(CodeRepositoryGitConfigOutput)
}

func (o CodeRepositoryOutput) Tags() CodeRepositoryTagArrayOutput {
	return o.ApplyT(func(v *CodeRepository) CodeRepositoryTagArrayOutput { return v.Tags }).(CodeRepositoryTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CodeRepositoryInput)(nil)).Elem(), &CodeRepository{})
	pulumi.RegisterOutputType(CodeRepositoryOutput{})
}
