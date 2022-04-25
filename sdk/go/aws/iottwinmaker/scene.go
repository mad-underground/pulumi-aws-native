// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iottwinmaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::IoTTwinMaker::Scene
type Scene struct {
	pulumi.CustomResourceState

	// The ARN of the scene.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A list of capabilities that the scene uses to render.
	Capabilities pulumi.StringArrayOutput `pulumi:"capabilities"`
	// The relative path that specifies the location of the content definition file.
	ContentLocation pulumi.StringOutput `pulumi:"contentLocation"`
	// The date and time when the scene was created.
	CreationDateTime pulumi.StringOutput `pulumi:"creationDateTime"`
	// The description of the scene.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the scene.
	SceneId pulumi.StringOutput `pulumi:"sceneId"`
	// A key-value pair to associate with a resource.
	Tags pulumi.AnyOutput `pulumi:"tags"`
	// The date and time of the current update.
	UpdateDateTime pulumi.StringOutput `pulumi:"updateDateTime"`
	// The ID of the scene.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewScene registers a new resource with the given unique name, arguments, and options.
func NewScene(ctx *pulumi.Context,
	name string, args *SceneArgs, opts ...pulumi.ResourceOption) (*Scene, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContentLocation == nil {
		return nil, errors.New("invalid value for required argument 'ContentLocation'")
	}
	if args.SceneId == nil {
		return nil, errors.New("invalid value for required argument 'SceneId'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	var resource Scene
	err := ctx.RegisterResource("aws-native:iottwinmaker:Scene", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScene gets an existing Scene resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScene(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SceneState, opts ...pulumi.ResourceOption) (*Scene, error) {
	var resource Scene
	err := ctx.ReadResource("aws-native:iottwinmaker:Scene", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Scene resources.
type sceneState struct {
}

type SceneState struct {
}

func (SceneState) ElementType() reflect.Type {
	return reflect.TypeOf((*sceneState)(nil)).Elem()
}

type sceneArgs struct {
	// A list of capabilities that the scene uses to render.
	Capabilities []string `pulumi:"capabilities"`
	// The relative path that specifies the location of the content definition file.
	ContentLocation string `pulumi:"contentLocation"`
	// The description of the scene.
	Description *string `pulumi:"description"`
	// The ID of the scene.
	SceneId string `pulumi:"sceneId"`
	// A key-value pair to associate with a resource.
	Tags interface{} `pulumi:"tags"`
	// The ID of the scene.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a Scene resource.
type SceneArgs struct {
	// A list of capabilities that the scene uses to render.
	Capabilities pulumi.StringArrayInput
	// The relative path that specifies the location of the content definition file.
	ContentLocation pulumi.StringInput
	// The description of the scene.
	Description pulumi.StringPtrInput
	// The ID of the scene.
	SceneId pulumi.StringInput
	// A key-value pair to associate with a resource.
	Tags pulumi.Input
	// The ID of the scene.
	WorkspaceId pulumi.StringInput
}

func (SceneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sceneArgs)(nil)).Elem()
}

type SceneInput interface {
	pulumi.Input

	ToSceneOutput() SceneOutput
	ToSceneOutputWithContext(ctx context.Context) SceneOutput
}

func (*Scene) ElementType() reflect.Type {
	return reflect.TypeOf((**Scene)(nil)).Elem()
}

func (i *Scene) ToSceneOutput() SceneOutput {
	return i.ToSceneOutputWithContext(context.Background())
}

func (i *Scene) ToSceneOutputWithContext(ctx context.Context) SceneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SceneOutput)
}

type SceneOutput struct{ *pulumi.OutputState }

func (SceneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Scene)(nil)).Elem()
}

func (o SceneOutput) ToSceneOutput() SceneOutput {
	return o
}

func (o SceneOutput) ToSceneOutputWithContext(ctx context.Context) SceneOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SceneInput)(nil)).Elem(), &Scene{})
	pulumi.RegisterOutputType(SceneOutput{})
}