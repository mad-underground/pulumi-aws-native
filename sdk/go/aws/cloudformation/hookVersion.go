// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Publishes new or first hook version to AWS CloudFormation Registry.
type HookVersion struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the type, here the HookVersion. This is used to uniquely identify a HookVersion resource
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
	ExecutionRoleArn pulumi.StringPtrOutput `pulumi:"executionRoleArn"`
	// Indicates if this type version is the current default version
	IsDefaultVersion pulumi.BoolOutput `pulumi:"isDefaultVersion"`
	// Specifies logging configuration information for a type.
	LoggingConfig HookVersionLoggingConfigPtrOutput `pulumi:"loggingConfig"`
	// A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.
	//
	// For information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.
	SchemaHandlerPackage pulumi.StringOutput `pulumi:"schemaHandlerPackage"`
	// The Amazon Resource Name (ARN) of the type without the versionID.
	TypeArn pulumi.StringOutput `pulumi:"typeArn"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName pulumi.StringOutput `pulumi:"typeName"`
	// The ID of the version of the type represented by this hook instance.
	VersionId pulumi.StringOutput `pulumi:"versionId"`
	// The scope at which the type is visible and usable in CloudFormation operations.
	//
	// Valid values include:
	//
	// PRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.
	//
	// PUBLIC: The type is publically visible and usable within any Amazon account.
	Visibility HookVersionVisibilityOutput `pulumi:"visibility"`
}

// NewHookVersion registers a new resource with the given unique name, arguments, and options.
func NewHookVersion(ctx *pulumi.Context,
	name string, args *HookVersionArgs, opts ...pulumi.ResourceOption) (*HookVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SchemaHandlerPackage == nil {
		return nil, errors.New("invalid value for required argument 'SchemaHandlerPackage'")
	}
	if args.TypeName == nil {
		return nil, errors.New("invalid value for required argument 'TypeName'")
	}
	var resource HookVersion
	err := ctx.RegisterResource("aws-native:cloudformation:HookVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHookVersion gets an existing HookVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHookVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HookVersionState, opts ...pulumi.ResourceOption) (*HookVersion, error) {
	var resource HookVersion
	err := ctx.ReadResource("aws-native:cloudformation:HookVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HookVersion resources.
type hookVersionState struct {
}

type HookVersionState struct {
}

func (HookVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*hookVersionState)(nil)).Elem()
}

type hookVersionArgs struct {
	// The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
	ExecutionRoleArn *string `pulumi:"executionRoleArn"`
	// Specifies logging configuration information for a type.
	LoggingConfig *HookVersionLoggingConfig `pulumi:"loggingConfig"`
	// A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.
	//
	// For information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.
	SchemaHandlerPackage string `pulumi:"schemaHandlerPackage"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName string `pulumi:"typeName"`
}

// The set of arguments for constructing a HookVersion resource.
type HookVersionArgs struct {
	// The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
	ExecutionRoleArn pulumi.StringPtrInput
	// Specifies logging configuration information for a type.
	LoggingConfig HookVersionLoggingConfigPtrInput
	// A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.
	//
	// For information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.
	SchemaHandlerPackage pulumi.StringInput
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName pulumi.StringInput
}

func (HookVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hookVersionArgs)(nil)).Elem()
}

type HookVersionInput interface {
	pulumi.Input

	ToHookVersionOutput() HookVersionOutput
	ToHookVersionOutputWithContext(ctx context.Context) HookVersionOutput
}

func (*HookVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**HookVersion)(nil)).Elem()
}

func (i *HookVersion) ToHookVersionOutput() HookVersionOutput {
	return i.ToHookVersionOutputWithContext(context.Background())
}

func (i *HookVersion) ToHookVersionOutputWithContext(ctx context.Context) HookVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HookVersionOutput)
}

type HookVersionOutput struct{ *pulumi.OutputState }

func (HookVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HookVersion)(nil)).Elem()
}

func (o HookVersionOutput) ToHookVersionOutput() HookVersionOutput {
	return o
}

func (o HookVersionOutput) ToHookVersionOutputWithContext(ctx context.Context) HookVersionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HookVersionInput)(nil)).Elem(), &HookVersion{})
	pulumi.RegisterOutputType(HookVersionOutput{})
}
