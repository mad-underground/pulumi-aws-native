// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for a Redshift-managed VPC endpoint.
type EndpointAccess struct {
	pulumi.CustomResourceState

	// The DNS address of the endpoint.
	Address pulumi.StringOutput `pulumi:"address"`
	// A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account
	ClusterIdentifier pulumi.StringPtrOutput `pulumi:"clusterIdentifier"`
	// The time (UTC) that the endpoint was created.
	EndpointCreateTime pulumi.StringOutput `pulumi:"endpointCreateTime"`
	// The name of the endpoint.
	EndpointName pulumi.StringOutput `pulumi:"endpointName"`
	// The status of the endpoint.
	EndpointStatus pulumi.StringOutput `pulumi:"endpointStatus"`
	// The port number on which the cluster accepts incoming connections.
	Port pulumi.IntOutput `pulumi:"port"`
	// The AWS account ID of the owner of the cluster.
	ResourceOwner pulumi.StringPtrOutput `pulumi:"resourceOwner"`
	// The subnet group name where Amazon Redshift chooses to deploy the endpoint.
	SubnetGroupName pulumi.StringPtrOutput `pulumi:"subnetGroupName"`
	// The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.
	VpcEndpoint VpcEndpointPropertiesOutput `pulumi:"vpcEndpoint"`
	// A list of vpc security group ids to apply to the created endpoint access.
	VpcSecurityGroupIds pulumi.StringArrayOutput `pulumi:"vpcSecurityGroupIds"`
	// A list of Virtual Private Cloud (VPC) security groups to be associated with the endpoint.
	VpcSecurityGroups EndpointAccessVpcSecurityGroupArrayOutput `pulumi:"vpcSecurityGroups"`
}

// NewEndpointAccess registers a new resource with the given unique name, arguments, and options.
func NewEndpointAccess(ctx *pulumi.Context,
	name string, args *EndpointAccessArgs, opts ...pulumi.ResourceOption) (*EndpointAccess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.VpcSecurityGroupIds == nil {
		return nil, errors.New("invalid value for required argument 'VpcSecurityGroupIds'")
	}
	var resource EndpointAccess
	err := ctx.RegisterResource("aws-native:redshift:EndpointAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointAccess gets an existing EndpointAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointAccessState, opts ...pulumi.ResourceOption) (*EndpointAccess, error) {
	var resource EndpointAccess
	err := ctx.ReadResource("aws-native:redshift:EndpointAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointAccess resources.
type endpointAccessState struct {
}

type EndpointAccessState struct {
}

func (EndpointAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointAccessState)(nil)).Elem()
}

type endpointAccessArgs struct {
	// A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account
	ClusterIdentifier *string `pulumi:"clusterIdentifier"`
	// The name of the endpoint.
	EndpointName string `pulumi:"endpointName"`
	// The AWS account ID of the owner of the cluster.
	ResourceOwner *string `pulumi:"resourceOwner"`
	// The subnet group name where Amazon Redshift chooses to deploy the endpoint.
	SubnetGroupName *string `pulumi:"subnetGroupName"`
	// A list of vpc security group ids to apply to the created endpoint access.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
}

// The set of arguments for constructing a EndpointAccess resource.
type EndpointAccessArgs struct {
	// A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account
	ClusterIdentifier pulumi.StringPtrInput
	// The name of the endpoint.
	EndpointName pulumi.StringInput
	// The AWS account ID of the owner of the cluster.
	ResourceOwner pulumi.StringPtrInput
	// The subnet group name where Amazon Redshift chooses to deploy the endpoint.
	SubnetGroupName pulumi.StringPtrInput
	// A list of vpc security group ids to apply to the created endpoint access.
	VpcSecurityGroupIds pulumi.StringArrayInput
}

func (EndpointAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointAccessArgs)(nil)).Elem()
}

type EndpointAccessInput interface {
	pulumi.Input

	ToEndpointAccessOutput() EndpointAccessOutput
	ToEndpointAccessOutputWithContext(ctx context.Context) EndpointAccessOutput
}

func (*EndpointAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointAccess)(nil))
}

func (i *EndpointAccess) ToEndpointAccessOutput() EndpointAccessOutput {
	return i.ToEndpointAccessOutputWithContext(context.Background())
}

func (i *EndpointAccess) ToEndpointAccessOutputWithContext(ctx context.Context) EndpointAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAccessOutput)
}

type EndpointAccessOutput struct{ *pulumi.OutputState }

func (EndpointAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointAccess)(nil))
}

func (o EndpointAccessOutput) ToEndpointAccessOutput() EndpointAccessOutput {
	return o
}

func (o EndpointAccessOutput) ToEndpointAccessOutputWithContext(ctx context.Context) EndpointAccessOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointAccessInput)(nil)).Elem(), &EndpointAccess{})
	pulumi.RegisterOutputType(EndpointAccessOutput{})
}