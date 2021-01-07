// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package amazonmq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html
type Broker struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes BrokerAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties BrokerPropertiesOutput `pulumi:"properties"`
}

// NewBroker registers a new resource with the given unique name, arguments, and options.
func NewBroker(ctx *pulumi.Context,
	name string, args *BrokerArgs, opts ...pulumi.ResourceOption) (*Broker, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource Broker
	err := ctx.RegisterResource("cloudformation:AmazonMQ:Broker", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBroker gets an existing Broker resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBroker(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BrokerState, opts ...pulumi.ResourceOption) (*Broker, error) {
	var resource Broker
	err := ctx.ReadResource("cloudformation:AmazonMQ:Broker", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Broker resources.
type brokerState struct {
	// The attributes associated with the resource
	Attributes *BrokerAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *BrokerProperties `pulumi:"properties"`
}

type BrokerState struct {
	// The attributes associated with the resource
	Attributes BrokerAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties BrokerPropertiesPtrInput
}

func (BrokerState) ElementType() reflect.Type {
	return reflect.TypeOf((*brokerState)(nil)).Elem()
}

type brokerArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties BrokerProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a Broker resource.
type BrokerArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties BrokerPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (BrokerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*brokerArgs)(nil)).Elem()
}

type BrokerInput interface {
	pulumi.Input

	ToBrokerOutput() BrokerOutput
	ToBrokerOutputWithContext(ctx context.Context) BrokerOutput
}

func (*Broker) ElementType() reflect.Type {
	return reflect.TypeOf((*Broker)(nil))
}

func (i *Broker) ToBrokerOutput() BrokerOutput {
	return i.ToBrokerOutputWithContext(context.Background())
}

func (i *Broker) ToBrokerOutputWithContext(ctx context.Context) BrokerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BrokerOutput)
}

type BrokerOutput struct {
	*pulumi.OutputState
}

func (BrokerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Broker)(nil))
}

func (o BrokerOutput) ToBrokerOutput() BrokerOutput {
	return o
}

func (o BrokerOutput) ToBrokerOutputWithContext(ctx context.Context) BrokerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BrokerOutput{})
}