// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3objectlambda

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::S3ObjectLambda::AccessPoint resource is an Amazon S3ObjectLambda resource type that you can use to add computation to S3 actions
type AccessPoint struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// The date and time when the Object lambda Access Point was created.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The name you want to assign to this Object lambda Access Point.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions
	ObjectLambdaConfiguration AccessPointObjectLambdaConfigurationPtrOutput `pulumi:"objectLambdaConfiguration"`
	PolicyStatus              PolicyStatusPropertiesOutput                  `pulumi:"policyStatus"`
	// The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
	PublicAccessBlockConfiguration AccessPointPublicAccessBlockConfigurationOutput `pulumi:"publicAccessBlockConfiguration"`
}

// NewAccessPoint registers a new resource with the given unique name, arguments, and options.
func NewAccessPoint(ctx *pulumi.Context,
	name string, args *AccessPointArgs, opts ...pulumi.ResourceOption) (*AccessPoint, error) {
	if args == nil {
		args = &AccessPointArgs{}
	}

	var resource AccessPoint
	err := ctx.RegisterResource("aws-native:s3objectlambda:AccessPoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPoint gets an existing AccessPoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPointState, opts ...pulumi.ResourceOption) (*AccessPoint, error) {
	var resource AccessPoint
	err := ctx.ReadResource("aws-native:s3objectlambda:AccessPoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPoint resources.
type accessPointState struct {
}

type AccessPointState struct {
}

func (AccessPointState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPointState)(nil)).Elem()
}

type accessPointArgs struct {
	// The name you want to assign to this Object lambda Access Point.
	Name *string `pulumi:"name"`
	// The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions
	ObjectLambdaConfiguration *AccessPointObjectLambdaConfiguration `pulumi:"objectLambdaConfiguration"`
}

// The set of arguments for constructing a AccessPoint resource.
type AccessPointArgs struct {
	// The name you want to assign to this Object lambda Access Point.
	Name pulumi.StringPtrInput
	// The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions
	ObjectLambdaConfiguration AccessPointObjectLambdaConfigurationPtrInput
}

func (AccessPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPointArgs)(nil)).Elem()
}

type AccessPointInput interface {
	pulumi.Input

	ToAccessPointOutput() AccessPointOutput
	ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput
}

func (*AccessPoint) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPoint)(nil))
}

func (i *AccessPoint) ToAccessPointOutput() AccessPointOutput {
	return i.ToAccessPointOutputWithContext(context.Background())
}

func (i *AccessPoint) ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointOutput)
}

type AccessPointOutput struct{ *pulumi.OutputState }

func (AccessPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPoint)(nil))
}

func (o AccessPointOutput) ToAccessPointOutput() AccessPointOutput {
	return o
}

func (o AccessPointOutput) ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccessPointOutput{})
}
