// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rekognition

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::Rekognition::Project type creates an Amazon Rekognition CustomLabels Project. A project is a grouping of the resources needed to create and manage Dataset and ProjectVersions.
type Project struct {
	pulumi.CustomResourceState

	Arn         pulumi.StringOutput `pulumi:"arn"`
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterResource("aws-native:rekognition:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("aws-native:rekognition:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
}

type ProjectState struct {
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	ProjectName *string `pulumi:"projectName"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	ProjectName pulumi.StringPtrInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

func (o ProjectOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ProjectOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterOutputType(ProjectOutput{})
}
