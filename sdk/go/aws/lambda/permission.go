// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “AWS::Lambda::Permission“ resource grants an AWS service or another account permission to use a function. You can apply the policy at the function level, or specify a qualifier to restrict access to a single version or alias. If you use a qualifier, the invoker must use the full Amazon Resource Name (ARN) of that version or alias to invoke the function.
//
//	To grant permission to another account, specify the account ID as the ``Principal``. To grant permission to an organization defined in AOlong, specify the organization ID as the ``PrincipalOrgID``. For AWS services, the principal is a domain-style identifier defined by the service, like ``s3.amazonaws.com`` or ``sns.amazonaws.com``. For AWS services, you can also specify the ARN of the associated resource as the ``SourceArn``. If you grant permission to a service principal without specifying the source, other accounts could potentially configure resources in their account to invoke your Lambda function.
//	If your function has a fu
type Permission struct {
	pulumi.CustomResourceState

	// The action that the principal can use on the function. For example, ``lambda:InvokeFunction`` or ``lambda:GetFunction``.
	Action pulumi.StringOutput `pulumi:"action"`
	AwsId  pulumi.StringOutput `pulumi:"awsId"`
	// For Alexa Smart Home functions, a token that the invoker must supply.
	EventSourceToken pulumi.StringPtrOutput `pulumi:"eventSourceToken"`
	// The name of the Lambda function, version, or alias.
	//   **Name formats**
	//  +   *Function name* – ``my-function`` (name-only), ``my-function:v1`` (with alias).
	//   +   *Function ARN* – ``arn:aws:lambda:us-west-2:123456789012:function:my-function``.
	//   +   *Partial ARN* – ``123456789012:function:my-function``.
	//
	//  You can append a version number or alias to any of the formats. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	FunctionName pulumi.StringOutput `pulumi:"functionName"`
	// The type of authentication that your function URL uses. Set to ``AWS_IAM`` if you want to restrict access to authenticated users only. Set to ``NONE`` if you want to bypass IAM authentication to create a public endpoint. For more information, see [Security and auth model for Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html).
	FunctionUrlAuthType PermissionFunctionUrlAuthTypePtrOutput `pulumi:"functionUrlAuthType"`
	// The AWS-service or AWS-account that invokes the function. If you specify a service, use ``SourceArn`` or ``SourceAccount`` to limit who can invoke the function through that service.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The identifier for your organization in AOlong. Use this to grant permissions to all the AWS-accounts under this organization.
	PrincipalOrgId pulumi.StringPtrOutput `pulumi:"principalOrgId"`
	// For AWS-service, the ID of the AWS-account that owns the resource. Use this together with ``SourceArn`` to ensure that the specified account owns the resource. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.
	SourceAccount pulumi.StringPtrOutput `pulumi:"sourceAccount"`
	// For AWS-services, the ARN of the AWS resource that invokes the function. For example, an Amazon S3 bucket or Amazon SNS topic.
	//  Note that Lambda configures the comparison using the ``StringLike`` operator.
	SourceArn pulumi.StringPtrOutput `pulumi:"sourceArn"`
}

// NewPermission registers a new resource with the given unique name, arguments, and options.
func NewPermission(ctx *pulumi.Context,
	name string, args *PermissionArgs, opts ...pulumi.ResourceOption) (*Permission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.FunctionName == nil {
		return nil, errors.New("invalid value for required argument 'FunctionName'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"action",
		"eventSourceToken",
		"functionName",
		"functionUrlAuthType",
		"principal",
		"principalOrgId",
		"sourceAccount",
		"sourceArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Permission
	err := ctx.RegisterResource("aws-native:lambda:Permission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermission gets an existing Permission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionState, opts ...pulumi.ResourceOption) (*Permission, error) {
	var resource Permission
	err := ctx.ReadResource("aws-native:lambda:Permission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Permission resources.
type permissionState struct {
}

type PermissionState struct {
}

func (PermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionState)(nil)).Elem()
}

type permissionArgs struct {
	// The action that the principal can use on the function. For example, ``lambda:InvokeFunction`` or ``lambda:GetFunction``.
	Action string `pulumi:"action"`
	// For Alexa Smart Home functions, a token that the invoker must supply.
	EventSourceToken *string `pulumi:"eventSourceToken"`
	// The name of the Lambda function, version, or alias.
	//   **Name formats**
	//  +   *Function name* – ``my-function`` (name-only), ``my-function:v1`` (with alias).
	//   +   *Function ARN* – ``arn:aws:lambda:us-west-2:123456789012:function:my-function``.
	//   +   *Partial ARN* – ``123456789012:function:my-function``.
	//
	//  You can append a version number or alias to any of the formats. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	FunctionName string `pulumi:"functionName"`
	// The type of authentication that your function URL uses. Set to ``AWS_IAM`` if you want to restrict access to authenticated users only. Set to ``NONE`` if you want to bypass IAM authentication to create a public endpoint. For more information, see [Security and auth model for Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html).
	FunctionUrlAuthType *PermissionFunctionUrlAuthType `pulumi:"functionUrlAuthType"`
	// The AWS-service or AWS-account that invokes the function. If you specify a service, use ``SourceArn`` or ``SourceAccount`` to limit who can invoke the function through that service.
	Principal string `pulumi:"principal"`
	// The identifier for your organization in AOlong. Use this to grant permissions to all the AWS-accounts under this organization.
	PrincipalOrgId *string `pulumi:"principalOrgId"`
	// For AWS-service, the ID of the AWS-account that owns the resource. Use this together with ``SourceArn`` to ensure that the specified account owns the resource. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.
	SourceAccount *string `pulumi:"sourceAccount"`
	// For AWS-services, the ARN of the AWS resource that invokes the function. For example, an Amazon S3 bucket or Amazon SNS topic.
	//  Note that Lambda configures the comparison using the ``StringLike`` operator.
	SourceArn *string `pulumi:"sourceArn"`
}

// The set of arguments for constructing a Permission resource.
type PermissionArgs struct {
	// The action that the principal can use on the function. For example, ``lambda:InvokeFunction`` or ``lambda:GetFunction``.
	Action pulumi.StringInput
	// For Alexa Smart Home functions, a token that the invoker must supply.
	EventSourceToken pulumi.StringPtrInput
	// The name of the Lambda function, version, or alias.
	//   **Name formats**
	//  +   *Function name* – ``my-function`` (name-only), ``my-function:v1`` (with alias).
	//   +   *Function ARN* – ``arn:aws:lambda:us-west-2:123456789012:function:my-function``.
	//   +   *Partial ARN* – ``123456789012:function:my-function``.
	//
	//  You can append a version number or alias to any of the formats. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	FunctionName pulumi.StringInput
	// The type of authentication that your function URL uses. Set to ``AWS_IAM`` if you want to restrict access to authenticated users only. Set to ``NONE`` if you want to bypass IAM authentication to create a public endpoint. For more information, see [Security and auth model for Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html).
	FunctionUrlAuthType PermissionFunctionUrlAuthTypePtrInput
	// The AWS-service or AWS-account that invokes the function. If you specify a service, use ``SourceArn`` or ``SourceAccount`` to limit who can invoke the function through that service.
	Principal pulumi.StringInput
	// The identifier for your organization in AOlong. Use this to grant permissions to all the AWS-accounts under this organization.
	PrincipalOrgId pulumi.StringPtrInput
	// For AWS-service, the ID of the AWS-account that owns the resource. Use this together with ``SourceArn`` to ensure that the specified account owns the resource. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.
	SourceAccount pulumi.StringPtrInput
	// For AWS-services, the ARN of the AWS resource that invokes the function. For example, an Amazon S3 bucket or Amazon SNS topic.
	//  Note that Lambda configures the comparison using the ``StringLike`` operator.
	SourceArn pulumi.StringPtrInput
}

func (PermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionArgs)(nil)).Elem()
}

type PermissionInput interface {
	pulumi.Input

	ToPermissionOutput() PermissionOutput
	ToPermissionOutputWithContext(ctx context.Context) PermissionOutput
}

func (*Permission) ElementType() reflect.Type {
	return reflect.TypeOf((**Permission)(nil)).Elem()
}

func (i *Permission) ToPermissionOutput() PermissionOutput {
	return i.ToPermissionOutputWithContext(context.Background())
}

func (i *Permission) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionOutput)
}

type PermissionOutput struct{ *pulumi.OutputState }

func (PermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Permission)(nil)).Elem()
}

func (o PermissionOutput) ToPermissionOutput() PermissionOutput {
	return o
}

func (o PermissionOutput) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return o
}

// The action that the principal can use on the function. For example, “lambda:InvokeFunction“ or “lambda:GetFunction“.
func (o PermissionOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

func (o PermissionOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// For Alexa Smart Home functions, a token that the invoker must supply.
func (o PermissionOutput) EventSourceToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringPtrOutput { return v.EventSourceToken }).(pulumi.StringPtrOutput)
}

// The name of the Lambda function, version, or alias.
//
//	 **Name formats**
//	+   *Function name* – ``my-function`` (name-only), ``my-function:v1`` (with alias).
//	 +   *Function ARN* – ``arn:aws:lambda:us-west-2:123456789012:function:my-function``.
//	 +   *Partial ARN* – ``123456789012:function:my-function``.
//
//	You can append a version number or alias to any of the formats. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
func (o PermissionOutput) FunctionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringOutput { return v.FunctionName }).(pulumi.StringOutput)
}

// The type of authentication that your function URL uses. Set to “AWS_IAM“ if you want to restrict access to authenticated users only. Set to “NONE“ if you want to bypass IAM authentication to create a public endpoint. For more information, see [Security and auth model for Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html).
func (o PermissionOutput) FunctionUrlAuthType() PermissionFunctionUrlAuthTypePtrOutput {
	return o.ApplyT(func(v *Permission) PermissionFunctionUrlAuthTypePtrOutput { return v.FunctionUrlAuthType }).(PermissionFunctionUrlAuthTypePtrOutput)
}

// The AWS-service or AWS-account that invokes the function. If you specify a service, use “SourceArn“ or “SourceAccount“ to limit who can invoke the function through that service.
func (o PermissionOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// The identifier for your organization in AOlong. Use this to grant permissions to all the AWS-accounts under this organization.
func (o PermissionOutput) PrincipalOrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringPtrOutput { return v.PrincipalOrgId }).(pulumi.StringPtrOutput)
}

// For AWS-service, the ID of the AWS-account that owns the resource. Use this together with “SourceArn“ to ensure that the specified account owns the resource. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.
func (o PermissionOutput) SourceAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringPtrOutput { return v.SourceAccount }).(pulumi.StringPtrOutput)
}

// For AWS-services, the ARN of the AWS resource that invokes the function. For example, an Amazon S3 bucket or Amazon SNS topic.
//
//	Note that Lambda configures the comparison using the ``StringLike`` operator.
func (o PermissionOutput) SourceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Permission) pulumi.StringPtrOutput { return v.SourceArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionInput)(nil)).Elem(), &Permission{})
	pulumi.RegisterOutputType(PermissionOutput{})
}
