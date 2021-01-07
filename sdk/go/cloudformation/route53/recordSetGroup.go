// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html
type RecordSetGroup struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes RecordSetGroupAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties RecordSetGroupPropertiesOutput `pulumi:"properties"`
}

// NewRecordSetGroup registers a new resource with the given unique name, arguments, and options.
func NewRecordSetGroup(ctx *pulumi.Context,
	name string, args *RecordSetGroupArgs, opts ...pulumi.ResourceOption) (*RecordSetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource RecordSetGroup
	err := ctx.RegisterResource("cloudformation:Route53:RecordSetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRecordSetGroup gets an existing RecordSetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecordSetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordSetGroupState, opts ...pulumi.ResourceOption) (*RecordSetGroup, error) {
	var resource RecordSetGroup
	err := ctx.ReadResource("cloudformation:Route53:RecordSetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RecordSetGroup resources.
type recordSetGroupState struct {
	// The attributes associated with the resource
	Attributes *RecordSetGroupAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *RecordSetGroupProperties `pulumi:"properties"`
}

type RecordSetGroupState struct {
	// The attributes associated with the resource
	Attributes RecordSetGroupAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties RecordSetGroupPropertiesPtrInput
}

func (RecordSetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetGroupState)(nil)).Elem()
}

type recordSetGroupArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties RecordSetGroupProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a RecordSetGroup resource.
type RecordSetGroupArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties RecordSetGroupPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (RecordSetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetGroupArgs)(nil)).Elem()
}

type RecordSetGroupInput interface {
	pulumi.Input

	ToRecordSetGroupOutput() RecordSetGroupOutput
	ToRecordSetGroupOutputWithContext(ctx context.Context) RecordSetGroupOutput
}

func (*RecordSetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordSetGroup)(nil))
}

func (i *RecordSetGroup) ToRecordSetGroupOutput() RecordSetGroupOutput {
	return i.ToRecordSetGroupOutputWithContext(context.Background())
}

func (i *RecordSetGroup) ToRecordSetGroupOutputWithContext(ctx context.Context) RecordSetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetGroupOutput)
}

type RecordSetGroupOutput struct {
	*pulumi.OutputState
}

func (RecordSetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordSetGroup)(nil))
}

func (o RecordSetGroupOutput) ToRecordSetGroupOutput() RecordSetGroupOutput {
	return o
}

func (o RecordSetGroupOutput) ToRecordSetGroupOutputWithContext(ctx context.Context) RecordSetGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RecordSetGroupOutput{})
}
