// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codestar

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestar-githubrepository.html
type GitHubRepository struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes GitHubRepositoryAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties GitHubRepositoryPropertiesOutput `pulumi:"properties"`
}

// NewGitHubRepository registers a new resource with the given unique name, arguments, and options.
func NewGitHubRepository(ctx *pulumi.Context,
	name string, args *GitHubRepositoryArgs, opts ...pulumi.ResourceOption) (*GitHubRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource GitHubRepository
	err := ctx.RegisterResource("cloudformation:CodeStar:GitHubRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGitHubRepository gets an existing GitHubRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGitHubRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GitHubRepositoryState, opts ...pulumi.ResourceOption) (*GitHubRepository, error) {
	var resource GitHubRepository
	err := ctx.ReadResource("cloudformation:CodeStar:GitHubRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GitHubRepository resources.
type gitHubRepositoryState struct {
	// The attributes associated with the resource
	Attributes *GitHubRepositoryAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *GitHubRepositoryProperties `pulumi:"properties"`
}

type GitHubRepositoryState struct {
	// The attributes associated with the resource
	Attributes GitHubRepositoryAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties GitHubRepositoryPropertiesPtrInput
}

func (GitHubRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*gitHubRepositoryState)(nil)).Elem()
}

type gitHubRepositoryArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties GitHubRepositoryProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a GitHubRepository resource.
type GitHubRepositoryArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties GitHubRepositoryPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (GitHubRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gitHubRepositoryArgs)(nil)).Elem()
}

type GitHubRepositoryInput interface {
	pulumi.Input

	ToGitHubRepositoryOutput() GitHubRepositoryOutput
	ToGitHubRepositoryOutputWithContext(ctx context.Context) GitHubRepositoryOutput
}

func (*GitHubRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubRepository)(nil))
}

func (i *GitHubRepository) ToGitHubRepositoryOutput() GitHubRepositoryOutput {
	return i.ToGitHubRepositoryOutputWithContext(context.Background())
}

func (i *GitHubRepository) ToGitHubRepositoryOutputWithContext(ctx context.Context) GitHubRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubRepositoryOutput)
}

type GitHubRepositoryOutput struct {
	*pulumi.OutputState
}

func (GitHubRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubRepository)(nil))
}

func (o GitHubRepositoryOutput) ToGitHubRepositoryOutput() GitHubRepositoryOutput {
	return o
}

func (o GitHubRepositoryOutput) ToGitHubRepositoryOutputWithContext(ctx context.Context) GitHubRepositoryOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GitHubRepositoryOutput{})
}
