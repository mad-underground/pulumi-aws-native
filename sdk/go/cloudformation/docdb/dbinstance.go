// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html
type DBInstance struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes DBInstanceAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties DBInstancePropertiesOutput `pulumi:"properties"`
}

// NewDBInstance registers a new resource with the given unique name, arguments, and options.
func NewDBInstance(ctx *pulumi.Context,
	name string, args *DBInstanceArgs, opts ...pulumi.ResourceOption) (*DBInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource DBInstance
	err := ctx.RegisterResource("cloudformation:DocDB:DBInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDBInstance gets an existing DBInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDBInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DBInstanceState, opts ...pulumi.ResourceOption) (*DBInstance, error) {
	var resource DBInstance
	err := ctx.ReadResource("cloudformation:DocDB:DBInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DBInstance resources.
type dbinstanceState struct {
	// The attributes associated with the resource
	Attributes *DBInstanceAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *DBInstanceProperties `pulumi:"properties"`
}

type DBInstanceState struct {
	// The attributes associated with the resource
	Attributes DBInstanceAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties DBInstancePropertiesPtrInput
}

func (DBInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbinstanceState)(nil)).Elem()
}

type dbinstanceArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties DBInstanceProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a DBInstance resource.
type DBInstanceArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties DBInstancePropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (DBInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbinstanceArgs)(nil)).Elem()
}

type DBInstanceInput interface {
	pulumi.Input

	ToDBInstanceOutput() DBInstanceOutput
	ToDBInstanceOutputWithContext(ctx context.Context) DBInstanceOutput
}

func (*DBInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*DBInstance)(nil))
}

func (i *DBInstance) ToDBInstanceOutput() DBInstanceOutput {
	return i.ToDBInstanceOutputWithContext(context.Background())
}

func (i *DBInstance) ToDBInstanceOutputWithContext(ctx context.Context) DBInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBInstanceOutput)
}

type DBInstanceOutput struct {
	*pulumi.OutputState
}

func (DBInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBInstance)(nil))
}

func (o DBInstanceOutput) ToDBInstanceOutput() DBInstanceOutput {
	return o
}

func (o DBInstanceOutput) ToDBInstanceOutputWithContext(ctx context.Context) DBInstanceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DBInstanceOutput{})
}
