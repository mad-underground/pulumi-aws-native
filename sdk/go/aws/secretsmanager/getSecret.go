// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package secretsmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new secret. A *secret* can be a password, a set of credentials such as a user name and password, an OAuth token, or other secret information that you store in an encrypted form in Secrets Manager.
//
//	For RDS master user credentials, see [AWS::RDS::DBCluster MasterUserSecret](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-masterusersecret.html).
//	To retrieve a secret in a CFNshort template, use a *dynamic reference*. For more information, see [Retrieve a secret in an resource](https://docs.aws.amazon.com/secretsmanager/latest/userguide/cfn-example_reference-secret.html).
//	A common scenario is to first create a secret with ``GenerateSecretString``, which generates a password, and then use a dynamic reference to retrieve the username and password from the secret to use as credentials for a new database. See the example *Creating a Redshift cluster and a secret for the admin credentials*.
//	For information about creating a secret in the c
func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSecretResult
	err := ctx.Invoke("aws-native:secretsmanager:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecretArgs struct {
	Id string `pulumi:"id"`
}

type LookupSecretResult struct {
	// The description of the secret.
	Description *string `pulumi:"description"`
	Id          *string `pulumi:"id"`
	// The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret. An alias is always prefixed by ``alias/``, for example ``alias/aws/secretsmanager``. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html).
	//  To use a KMS key in a different account, use the key ARN or the alias ARN.
	//  If you don't specify this value, then Secrets Manager uses the key ``aws/secretsmanager``. If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
	//  If the secret is in a different AWS account from the credentials calling the API, then you can't use ``aws/secretsmanager`` to encrypt the secret, and you must create and use a customer managed KMS key.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A custom type that specifies a ``Region`` and the ``KmsKeyId`` for a replica secret.
	ReplicaRegions []SecretReplicaRegion `pulumi:"replicaRegions"`
	// A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string, for example:
	//   ``[{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]``
	//  Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a different tag from one with key "abc".
	//  Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret.
	//  If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an ``Access Denied`` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazo
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupSecretOutput(ctx *pulumi.Context, args LookupSecretOutputArgs, opts ...pulumi.InvokeOption) LookupSecretResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecretResult, error) {
			args := v.(LookupSecretArgs)
			r, err := LookupSecret(ctx, &args, opts...)
			var s LookupSecretResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecretResultOutput)
}

type LookupSecretOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupSecretOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretArgs)(nil)).Elem()
}

type LookupSecretResultOutput struct{ *pulumi.OutputState }

func (LookupSecretResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretResult)(nil)).Elem()
}

func (o LookupSecretResultOutput) ToLookupSecretResultOutput() LookupSecretResultOutput {
	return o
}

func (o LookupSecretResultOutput) ToLookupSecretResultOutputWithContext(ctx context.Context) LookupSecretResultOutput {
	return o
}

// The description of the secret.
func (o LookupSecretResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSecretResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret. An alias is always prefixed by “alias/“, for example “alias/aws/secretsmanager“. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html).
//
//	To use a KMS key in a different account, use the key ARN or the alias ARN.
//	If you don't specify this value, then Secrets Manager uses the key ``aws/secretsmanager``. If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
//	If the secret is in a different AWS account from the credentials calling the API, then you can't use ``aws/secretsmanager`` to encrypt the secret, and you must create and use a customer managed KMS key.
func (o LookupSecretResultOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretResult) *string { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// A custom type that specifies a “Region“ and the “KmsKeyId“ for a replica secret.
func (o LookupSecretResultOutput) ReplicaRegions() SecretReplicaRegionArrayOutput {
	return o.ApplyT(func(v LookupSecretResult) []SecretReplicaRegion { return v.ReplicaRegions }).(SecretReplicaRegionArrayOutput)
}

// A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string, for example:
//
//	 ``[{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]``
//	Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a different tag from one with key "abc".
//	Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret.
//	If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an ``Access Denied`` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazo
func (o LookupSecretResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupSecretResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecretResultOutput{})
}
