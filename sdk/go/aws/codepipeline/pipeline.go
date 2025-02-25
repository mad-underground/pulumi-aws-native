// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codepipeline

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CodePipeline::Pipeline
//
// Deprecated: Pipeline is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Pipeline struct {
	pulumi.CustomResourceState

	ArtifactStore                  PipelineArtifactStorePtrOutput         `pulumi:"artifactStore"`
	ArtifactStores                 PipelineArtifactStoreMapArrayOutput    `pulumi:"artifactStores"`
	AwsId                          pulumi.StringOutput                    `pulumi:"awsId"`
	DisableInboundStageTransitions PipelineStageTransitionArrayOutput     `pulumi:"disableInboundStageTransitions"`
	ExecutionMode                  pulumi.StringPtrOutput                 `pulumi:"executionMode"`
	Name                           pulumi.StringPtrOutput                 `pulumi:"name"`
	PipelineType                   pulumi.StringPtrOutput                 `pulumi:"pipelineType"`
	RestartExecutionOnUpdate       pulumi.BoolPtrOutput                   `pulumi:"restartExecutionOnUpdate"`
	RoleArn                        pulumi.StringOutput                    `pulumi:"roleArn"`
	Stages                         PipelineStageDeclarationArrayOutput    `pulumi:"stages"`
	Tags                           aws.TagArrayOutput                     `pulumi:"tags"`
	Triggers                       PipelineTriggerDeclarationArrayOutput  `pulumi:"triggers"`
	Variables                      PipelineVariableDeclarationArrayOutput `pulumi:"variables"`
	Version                        pulumi.StringOutput                    `pulumi:"version"`
}

// NewPipeline registers a new resource with the given unique name, arguments, and options.
func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.Stages == nil {
		return nil, errors.New("invalid value for required argument 'Stages'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Pipeline
	err := ctx.RegisterResource("aws-native:codepipeline:Pipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipeline gets an existing Pipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineState, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	var resource Pipeline
	err := ctx.ReadResource("aws-native:codepipeline:Pipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pipeline resources.
type pipelineState struct {
}

type PipelineState struct {
}

func (PipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineState)(nil)).Elem()
}

type pipelineArgs struct {
	ArtifactStore                  *PipelineArtifactStore        `pulumi:"artifactStore"`
	ArtifactStores                 []PipelineArtifactStoreMap    `pulumi:"artifactStores"`
	DisableInboundStageTransitions []PipelineStageTransition     `pulumi:"disableInboundStageTransitions"`
	ExecutionMode                  *string                       `pulumi:"executionMode"`
	Name                           *string                       `pulumi:"name"`
	PipelineType                   *string                       `pulumi:"pipelineType"`
	RestartExecutionOnUpdate       *bool                         `pulumi:"restartExecutionOnUpdate"`
	RoleArn                        string                        `pulumi:"roleArn"`
	Stages                         []PipelineStageDeclaration    `pulumi:"stages"`
	Tags                           []aws.Tag                     `pulumi:"tags"`
	Triggers                       []PipelineTriggerDeclaration  `pulumi:"triggers"`
	Variables                      []PipelineVariableDeclaration `pulumi:"variables"`
}

// The set of arguments for constructing a Pipeline resource.
type PipelineArgs struct {
	ArtifactStore                  PipelineArtifactStorePtrInput
	ArtifactStores                 PipelineArtifactStoreMapArrayInput
	DisableInboundStageTransitions PipelineStageTransitionArrayInput
	ExecutionMode                  pulumi.StringPtrInput
	Name                           pulumi.StringPtrInput
	PipelineType                   pulumi.StringPtrInput
	RestartExecutionOnUpdate       pulumi.BoolPtrInput
	RoleArn                        pulumi.StringInput
	Stages                         PipelineStageDeclarationArrayInput
	Tags                           aws.TagArrayInput
	Triggers                       PipelineTriggerDeclarationArrayInput
	Variables                      PipelineVariableDeclarationArrayInput
}

func (PipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineArgs)(nil)).Elem()
}

type PipelineInput interface {
	pulumi.Input

	ToPipelineOutput() PipelineOutput
	ToPipelineOutputWithContext(ctx context.Context) PipelineOutput
}

func (*Pipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil)).Elem()
}

func (i *Pipeline) ToPipelineOutput() PipelineOutput {
	return i.ToPipelineOutputWithContext(context.Background())
}

func (i *Pipeline) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOutput)
}

type PipelineOutput struct{ *pulumi.OutputState }

func (PipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil)).Elem()
}

func (o PipelineOutput) ToPipelineOutput() PipelineOutput {
	return o
}

func (o PipelineOutput) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return o
}

func (o PipelineOutput) ArtifactStore() PipelineArtifactStorePtrOutput {
	return o.ApplyT(func(v *Pipeline) PipelineArtifactStorePtrOutput { return v.ArtifactStore }).(PipelineArtifactStorePtrOutput)
}

func (o PipelineOutput) ArtifactStores() PipelineArtifactStoreMapArrayOutput {
	return o.ApplyT(func(v *Pipeline) PipelineArtifactStoreMapArrayOutput { return v.ArtifactStores }).(PipelineArtifactStoreMapArrayOutput)
}

func (o PipelineOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o PipelineOutput) DisableInboundStageTransitions() PipelineStageTransitionArrayOutput {
	return o.ApplyT(func(v *Pipeline) PipelineStageTransitionArrayOutput { return v.DisableInboundStageTransitions }).(PipelineStageTransitionArrayOutput)
}

func (o PipelineOutput) ExecutionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringPtrOutput { return v.ExecutionMode }).(pulumi.StringPtrOutput)
}

func (o PipelineOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PipelineOutput) PipelineType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringPtrOutput { return v.PipelineType }).(pulumi.StringPtrOutput)
}

func (o PipelineOutput) RestartExecutionOnUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.BoolPtrOutput { return v.RestartExecutionOnUpdate }).(pulumi.BoolPtrOutput)
}

func (o PipelineOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

func (o PipelineOutput) Stages() PipelineStageDeclarationArrayOutput {
	return o.ApplyT(func(v *Pipeline) PipelineStageDeclarationArrayOutput { return v.Stages }).(PipelineStageDeclarationArrayOutput)
}

func (o PipelineOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Pipeline) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func (o PipelineOutput) Triggers() PipelineTriggerDeclarationArrayOutput {
	return o.ApplyT(func(v *Pipeline) PipelineTriggerDeclarationArrayOutput { return v.Triggers }).(PipelineTriggerDeclarationArrayOutput)
}

func (o PipelineOutput) Variables() PipelineVariableDeclarationArrayOutput {
	return o.ApplyT(func(v *Pipeline) PipelineVariableDeclarationArrayOutput { return v.Variables }).(PipelineVariableDeclarationArrayOutput)
}

func (o PipelineOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineInput)(nil)).Elem(), &Pipeline{})
	pulumi.RegisterOutputType(PipelineOutput{})
}
