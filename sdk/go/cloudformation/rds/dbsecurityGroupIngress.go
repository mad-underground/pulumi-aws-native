// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html
type DBSecurityGroupIngress struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes DBSecurityGroupIngressAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties DBSecurityGroupIngressPropertiesOutput `pulumi:"properties"`
}

// NewDBSecurityGroupIngress registers a new resource with the given unique name, arguments, and options.
func NewDBSecurityGroupIngress(ctx *pulumi.Context,
	name string, args *DBSecurityGroupIngressArgs, opts ...pulumi.ResourceOption) (*DBSecurityGroupIngress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource DBSecurityGroupIngress
	err := ctx.RegisterResource("cloudformation:RDS:DBSecurityGroupIngress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDBSecurityGroupIngress gets an existing DBSecurityGroupIngress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDBSecurityGroupIngress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DBSecurityGroupIngressState, opts ...pulumi.ResourceOption) (*DBSecurityGroupIngress, error) {
	var resource DBSecurityGroupIngress
	err := ctx.ReadResource("cloudformation:RDS:DBSecurityGroupIngress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DBSecurityGroupIngress resources.
type dbsecurityGroupIngressState struct {
	// The attributes associated with the resource
	Attributes *DBSecurityGroupIngressAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *DBSecurityGroupIngressProperties `pulumi:"properties"`
}

type DBSecurityGroupIngressState struct {
	// The attributes associated with the resource
	Attributes DBSecurityGroupIngressAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties DBSecurityGroupIngressPropertiesPtrInput
}

func (DBSecurityGroupIngressState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbsecurityGroupIngressState)(nil)).Elem()
}

type dbsecurityGroupIngressArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties DBSecurityGroupIngressProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a DBSecurityGroupIngress resource.
type DBSecurityGroupIngressArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties DBSecurityGroupIngressPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (DBSecurityGroupIngressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbsecurityGroupIngressArgs)(nil)).Elem()
}

type DBSecurityGroupIngressInput interface {
	pulumi.Input

	ToDBSecurityGroupIngressOutput() DBSecurityGroupIngressOutput
	ToDBSecurityGroupIngressOutputWithContext(ctx context.Context) DBSecurityGroupIngressOutput
}

func (*DBSecurityGroupIngress) ElementType() reflect.Type {
	return reflect.TypeOf((*DBSecurityGroupIngress)(nil))
}

func (i *DBSecurityGroupIngress) ToDBSecurityGroupIngressOutput() DBSecurityGroupIngressOutput {
	return i.ToDBSecurityGroupIngressOutputWithContext(context.Background())
}

func (i *DBSecurityGroupIngress) ToDBSecurityGroupIngressOutputWithContext(ctx context.Context) DBSecurityGroupIngressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBSecurityGroupIngressOutput)
}

type DBSecurityGroupIngressOutput struct {
	*pulumi.OutputState
}

func (DBSecurityGroupIngressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBSecurityGroupIngress)(nil))
}

func (o DBSecurityGroupIngressOutput) ToDBSecurityGroupIngressOutput() DBSecurityGroupIngressOutput {
	return o
}

func (o DBSecurityGroupIngressOutput) ToDBSecurityGroupIngressOutputWithContext(ctx context.Context) DBSecurityGroupIngressOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DBSecurityGroupIngressOutput{})
}