// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appflow

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppFlow::ConnectorProfile
type ConnectorProfile struct {
	pulumi.CustomResourceState

	// Mode in which data transfer should be enabled. Private connection mode is currently enabled for Salesforce, Snowflake, Trendmicro and Singular
	ConnectionMode ConnectorProfileConnectionModeOutput `pulumi:"connectionMode"`
	// The label of the connector. The label is unique for each ConnectorRegistration in your AWS account. Only needed if calling for CUSTOMCONNECTOR connector type/.
	ConnectorLabel pulumi.StringPtrOutput `pulumi:"connectorLabel"`
	// Unique identifier for connector profile resources
	ConnectorProfileArn pulumi.StringOutput `pulumi:"connectorProfileArn"`
	// Connector specific configurations needed to create connector profile
	ConnectorProfileConfig ConnectorProfileConfigPtrOutput `pulumi:"connectorProfileConfig"`
	// The maximum number of items to retrieve in a single batch.
	ConnectorProfileName pulumi.StringOutput `pulumi:"connectorProfileName"`
	// List of Saas providers that need connector profile to be created
	ConnectorType ConnectorProfileConnectorTypeOutput `pulumi:"connectorType"`
	// A unique Arn for Connector-Profile resource
	CredentialsArn pulumi.StringOutput `pulumi:"credentialsArn"`
	// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
	KMSArn pulumi.StringPtrOutput `pulumi:"kMSArn"`
}

// NewConnectorProfile registers a new resource with the given unique name, arguments, and options.
func NewConnectorProfile(ctx *pulumi.Context,
	name string, args *ConnectorProfileArgs, opts ...pulumi.ResourceOption) (*ConnectorProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionMode == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionMode'")
	}
	if args.ConnectorType == nil {
		return nil, errors.New("invalid value for required argument 'ConnectorType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConnectorProfile
	err := ctx.RegisterResource("aws-native:appflow:ConnectorProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectorProfile gets an existing ConnectorProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectorProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorProfileState, opts ...pulumi.ResourceOption) (*ConnectorProfile, error) {
	var resource ConnectorProfile
	err := ctx.ReadResource("aws-native:appflow:ConnectorProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectorProfile resources.
type connectorProfileState struct {
}

type ConnectorProfileState struct {
}

func (ConnectorProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorProfileState)(nil)).Elem()
}

type connectorProfileArgs struct {
	// Mode in which data transfer should be enabled. Private connection mode is currently enabled for Salesforce, Snowflake, Trendmicro and Singular
	ConnectionMode ConnectorProfileConnectionMode `pulumi:"connectionMode"`
	// The label of the connector. The label is unique for each ConnectorRegistration in your AWS account. Only needed if calling for CUSTOMCONNECTOR connector type/.
	ConnectorLabel *string `pulumi:"connectorLabel"`
	// Connector specific configurations needed to create connector profile
	ConnectorProfileConfig *ConnectorProfileConfig `pulumi:"connectorProfileConfig"`
	// The maximum number of items to retrieve in a single batch.
	ConnectorProfileName *string `pulumi:"connectorProfileName"`
	// List of Saas providers that need connector profile to be created
	ConnectorType ConnectorProfileConnectorType `pulumi:"connectorType"`
	// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
	KMSArn *string `pulumi:"kMSArn"`
}

// The set of arguments for constructing a ConnectorProfile resource.
type ConnectorProfileArgs struct {
	// Mode in which data transfer should be enabled. Private connection mode is currently enabled for Salesforce, Snowflake, Trendmicro and Singular
	ConnectionMode ConnectorProfileConnectionModeInput
	// The label of the connector. The label is unique for each ConnectorRegistration in your AWS account. Only needed if calling for CUSTOMCONNECTOR connector type/.
	ConnectorLabel pulumi.StringPtrInput
	// Connector specific configurations needed to create connector profile
	ConnectorProfileConfig ConnectorProfileConfigPtrInput
	// The maximum number of items to retrieve in a single batch.
	ConnectorProfileName pulumi.StringPtrInput
	// List of Saas providers that need connector profile to be created
	ConnectorType ConnectorProfileConnectorTypeInput
	// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
	KMSArn pulumi.StringPtrInput
}

func (ConnectorProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorProfileArgs)(nil)).Elem()
}

type ConnectorProfileInput interface {
	pulumi.Input

	ToConnectorProfileOutput() ConnectorProfileOutput
	ToConnectorProfileOutputWithContext(ctx context.Context) ConnectorProfileOutput
}

func (*ConnectorProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorProfile)(nil)).Elem()
}

func (i *ConnectorProfile) ToConnectorProfileOutput() ConnectorProfileOutput {
	return i.ToConnectorProfileOutputWithContext(context.Background())
}

func (i *ConnectorProfile) ToConnectorProfileOutputWithContext(ctx context.Context) ConnectorProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorProfileOutput)
}

type ConnectorProfileOutput struct{ *pulumi.OutputState }

func (ConnectorProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorProfile)(nil)).Elem()
}

func (o ConnectorProfileOutput) ToConnectorProfileOutput() ConnectorProfileOutput {
	return o
}

func (o ConnectorProfileOutput) ToConnectorProfileOutputWithContext(ctx context.Context) ConnectorProfileOutput {
	return o
}

// Mode in which data transfer should be enabled. Private connection mode is currently enabled for Salesforce, Snowflake, Trendmicro and Singular
func (o ConnectorProfileOutput) ConnectionMode() ConnectorProfileConnectionModeOutput {
	return o.ApplyT(func(v *ConnectorProfile) ConnectorProfileConnectionModeOutput { return v.ConnectionMode }).(ConnectorProfileConnectionModeOutput)
}

// The label of the connector. The label is unique for each ConnectorRegistration in your AWS account. Only needed if calling for CUSTOMCONNECTOR connector type/.
func (o ConnectorProfileOutput) ConnectorLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorProfile) pulumi.StringPtrOutput { return v.ConnectorLabel }).(pulumi.StringPtrOutput)
}

// Unique identifier for connector profile resources
func (o ConnectorProfileOutput) ConnectorProfileArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorProfile) pulumi.StringOutput { return v.ConnectorProfileArn }).(pulumi.StringOutput)
}

// Connector specific configurations needed to create connector profile
func (o ConnectorProfileOutput) ConnectorProfileConfig() ConnectorProfileConfigPtrOutput {
	return o.ApplyT(func(v *ConnectorProfile) ConnectorProfileConfigPtrOutput { return v.ConnectorProfileConfig }).(ConnectorProfileConfigPtrOutput)
}

// The maximum number of items to retrieve in a single batch.
func (o ConnectorProfileOutput) ConnectorProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorProfile) pulumi.StringOutput { return v.ConnectorProfileName }).(pulumi.StringOutput)
}

// List of Saas providers that need connector profile to be created
func (o ConnectorProfileOutput) ConnectorType() ConnectorProfileConnectorTypeOutput {
	return o.ApplyT(func(v *ConnectorProfile) ConnectorProfileConnectorTypeOutput { return v.ConnectorType }).(ConnectorProfileConnectorTypeOutput)
}

// A unique Arn for Connector-Profile resource
func (o ConnectorProfileOutput) CredentialsArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorProfile) pulumi.StringOutput { return v.CredentialsArn }).(pulumi.StringOutput)
}

// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
func (o ConnectorProfileOutput) KMSArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorProfile) pulumi.StringPtrOutput { return v.KMSArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectorProfileInput)(nil)).Elem(), &ConnectorProfile{})
	pulumi.RegisterOutputType(ConnectorProfileOutput{})
}
