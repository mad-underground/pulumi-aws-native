// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package detective

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-graph.html
type Graph struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes GraphAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties GraphPropertiesOutput `pulumi:"properties"`
}

// NewGraph registers a new resource with the given unique name, arguments, and options.
func NewGraph(ctx *pulumi.Context,
	name string, args *GraphArgs, opts ...pulumi.ResourceOption) (*Graph, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource Graph
	err := ctx.RegisterResource("cloudformation:Detective:Graph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGraph gets an existing Graph resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GraphState, opts ...pulumi.ResourceOption) (*Graph, error) {
	var resource Graph
	err := ctx.ReadResource("cloudformation:Detective:Graph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Graph resources.
type graphState struct {
	// The attributes associated with the resource
	Attributes *GraphAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *GraphProperties `pulumi:"properties"`
}

type GraphState struct {
	// The attributes associated with the resource
	Attributes GraphAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties GraphPropertiesPtrInput
}

func (GraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*graphState)(nil)).Elem()
}

type graphArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties GraphProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a Graph resource.
type GraphArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties GraphPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (GraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*graphArgs)(nil)).Elem()
}

type GraphInput interface {
	pulumi.Input

	ToGraphOutput() GraphOutput
	ToGraphOutputWithContext(ctx context.Context) GraphOutput
}

func (*Graph) ElementType() reflect.Type {
	return reflect.TypeOf((*Graph)(nil))
}

func (i *Graph) ToGraphOutput() GraphOutput {
	return i.ToGraphOutputWithContext(context.Background())
}

func (i *Graph) ToGraphOutputWithContext(ctx context.Context) GraphOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphOutput)
}

type GraphOutput struct {
	*pulumi.OutputState
}

func (GraphOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Graph)(nil))
}

func (o GraphOutput) ToGraphOutput() GraphOutput {
	return o
}

func (o GraphOutput) ToGraphOutputWithContext(ctx context.Context) GraphOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GraphOutput{})
}
