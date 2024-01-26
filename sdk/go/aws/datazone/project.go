// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datazone

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Amazon DataZone projects are business use case–based groupings of people, assets (data), and tools used to simplify access to the AWS analytics.
type Project struct {
	pulumi.CustomResourceState

	// The timestamp of when the project was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The Amazon DataZone user who created the project.
	CreatedBy pulumi.StringOutput `pulumi:"createdBy"`
	// The description of the Amazon DataZone project.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The identifier of the Amazon DataZone domain in which the project was created.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// The ID of the Amazon DataZone domain in which this project is created.
	DomainIdentifier pulumi.StringOutput `pulumi:"domainIdentifier"`
	// The glossary terms that can be used in this Amazon DataZone project.
	GlossaryTerms pulumi.StringArrayOutput `pulumi:"glossaryTerms"`
	// The timestamp of when the project was last updated.
	LastUpdatedAt pulumi.StringOutput `pulumi:"lastUpdatedAt"`
	// The name of the Amazon DataZone project.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'DomainIdentifier'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"domainIdentifier",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterResource("aws-native:datazone:Project", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:datazone:Project", name, id, state, &resource, opts...)
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
	// The description of the Amazon DataZone project.
	Description *string `pulumi:"description"`
	// The ID of the Amazon DataZone domain in which this project is created.
	DomainIdentifier string `pulumi:"domainIdentifier"`
	// The glossary terms that can be used in this Amazon DataZone project.
	GlossaryTerms []string `pulumi:"glossaryTerms"`
	// The name of the Amazon DataZone project.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// The description of the Amazon DataZone project.
	Description pulumi.StringPtrInput
	// The ID of the Amazon DataZone domain in which this project is created.
	DomainIdentifier pulumi.StringInput
	// The glossary terms that can be used in this Amazon DataZone project.
	GlossaryTerms pulumi.StringArrayInput
	// The name of the Amazon DataZone project.
	Name pulumi.StringPtrInput
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

func (i *Project) ToOutput(ctx context.Context) pulumix.Output[*Project] {
	return pulumix.Output[*Project]{
		OutputState: i.ToProjectOutputWithContext(ctx).OutputState,
	}
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

func (o ProjectOutput) ToOutput(ctx context.Context) pulumix.Output[*Project] {
	return pulumix.Output[*Project]{
		OutputState: o.OutputState,
	}
}

// The timestamp of when the project was created.
func (o ProjectOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The Amazon DataZone user who created the project.
func (o ProjectOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

// The description of the Amazon DataZone project.
func (o ProjectOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The identifier of the Amazon DataZone domain in which the project was created.
func (o ProjectOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.DomainId }).(pulumi.StringOutput)
}

// The ID of the Amazon DataZone domain in which this project is created.
func (o ProjectOutput) DomainIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.DomainIdentifier }).(pulumi.StringOutput)
}

// The glossary terms that can be used in this Amazon DataZone project.
func (o ProjectOutput) GlossaryTerms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Project) pulumi.StringArrayOutput { return v.GlossaryTerms }).(pulumi.StringArrayOutput)
}

// The timestamp of when the project was last updated.
func (o ProjectOutput) LastUpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.LastUpdatedAt }).(pulumi.StringOutput)
}

// The name of the Amazon DataZone project.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterOutputType(ProjectOutput{})
}
