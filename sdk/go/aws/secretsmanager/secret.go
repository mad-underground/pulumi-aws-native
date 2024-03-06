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
//	For information about creating a secret in the console, see [Create a secret](https://docs.aws.amazon.com/secretsmanager/latest/userguide/manage_create-basic-secret.html). For information about creating a secret using the CLI or SDK, see [CreateSecret](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_CreateSecret.html).
//	For information about retrieving a secret in code, see [Retrieve secrets from Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieving-secrets.html).
type Secret struct {
	pulumi.CustomResourceState

	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// The description of the secret.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A structure that specifies how to generate a password to encrypt and store in the secret. To include a specific string in the secret, use ``SecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.
	//  We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.
	GenerateSecretString SecretGenerateSecretStringPtrOutput `pulumi:"generateSecretString"`
	// The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret. An alias is always prefixed by ``alias/``, for example ``alias/aws/secretsmanager``. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html).
	//  To use a KMS key in a different account, use the key ARN or the alias ARN.
	//  If you don't specify this value, then Secrets Manager uses the key ``aws/secretsmanager``. If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
	//  If the secret is in a different AWS account from the credentials calling the API, then you can't use ``aws/secretsmanager`` to encrypt the secret, and you must create and use a customer managed KMS key.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// The name of the new secret.
	//  The secret name can contain ASCII letters, numbers, and the following characters: /_+=.@-
	//  Do not end your secret name with a hyphen followed by six characters. If you do so, you risk confusion and unexpected results when searching for a secret by partial ARN. Secrets Manager automatically adds a hyphen and six random characters after the secret name at the end of the ARN.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// A custom type that specifies a ``Region`` and the ``KmsKeyId`` for a replica secret.
	ReplicaRegions SecretReplicaRegionArrayOutput `pulumi:"replicaRegions"`
	// The text to encrypt and store in the secret. We recommend you use a JSON structure of key/value pairs for your secret value. To generate a random password, use ``GenerateSecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.
	SecretString pulumi.StringPtrOutput `pulumi:"secretString"`
	// A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string, for example:
	//   ``[{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]``
	//  Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a different tag from one with key "abc".
	//  Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret.
	//  If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an ``Access Denied`` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#auth-and-access_tags2).
	//  For information about how to format a JSON parameter for the various command line tool environments, see [Using JSON for Parameters](https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json). If your command-line tool or SDK requires quotation marks around the parameter, you should use single quotes to avoid confusion with the double quotes required in the JSON text.
	//  The following restrictions apply to tags:
	//   +  Maximum number of tags per secret: 50
	//   +  Maximum key length: 127 Unicode characters in UTF-8
	//   +  Maximum value length: 255 Unicode characters in UTF-8
	//   +  Tag keys and values are case sensitive.
	//   +  Do not use the ``aws:`` prefix in your tag names or values because AWS reserves it for AWS use. You can't edit or delete tag names or values with this prefix. Tags with this prefix do not count against your tags per secret limit.
	//   +  If you use your tagging schema across multiple services and resources, other services might have restrictions on allowed characters. Generally allowed characters: letters, spaces, and numbers representable in UTF-8, plus the following special characters: + - = . _ : / @.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewSecret registers a new resource with the given unique name, arguments, and options.
func NewSecret(ctx *pulumi.Context,
	name string, args *SecretArgs, opts ...pulumi.ResourceOption) (*Secret, error) {
	if args == nil {
		args = &SecretArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Secret
	err := ctx.RegisterResource("aws-native:secretsmanager:Secret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecret gets an existing Secret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretState, opts ...pulumi.ResourceOption) (*Secret, error) {
	var resource Secret
	err := ctx.ReadResource("aws-native:secretsmanager:Secret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Secret resources.
type secretState struct {
}

type SecretState struct {
}

func (SecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretState)(nil)).Elem()
}

type secretArgs struct {
	// The description of the secret.
	Description *string `pulumi:"description"`
	// A structure that specifies how to generate a password to encrypt and store in the secret. To include a specific string in the secret, use ``SecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.
	//  We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.
	GenerateSecretString *SecretGenerateSecretString `pulumi:"generateSecretString"`
	// The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret. An alias is always prefixed by ``alias/``, for example ``alias/aws/secretsmanager``. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html).
	//  To use a KMS key in a different account, use the key ARN or the alias ARN.
	//  If you don't specify this value, then Secrets Manager uses the key ``aws/secretsmanager``. If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
	//  If the secret is in a different AWS account from the credentials calling the API, then you can't use ``aws/secretsmanager`` to encrypt the secret, and you must create and use a customer managed KMS key.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The name of the new secret.
	//  The secret name can contain ASCII letters, numbers, and the following characters: /_+=.@-
	//  Do not end your secret name with a hyphen followed by six characters. If you do so, you risk confusion and unexpected results when searching for a secret by partial ARN. Secrets Manager automatically adds a hyphen and six random characters after the secret name at the end of the ARN.
	Name *string `pulumi:"name"`
	// A custom type that specifies a ``Region`` and the ``KmsKeyId`` for a replica secret.
	ReplicaRegions []SecretReplicaRegion `pulumi:"replicaRegions"`
	// The text to encrypt and store in the secret. We recommend you use a JSON structure of key/value pairs for your secret value. To generate a random password, use ``GenerateSecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.
	SecretString *string `pulumi:"secretString"`
	// A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string, for example:
	//   ``[{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]``
	//  Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a different tag from one with key "abc".
	//  Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret.
	//  If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an ``Access Denied`` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#auth-and-access_tags2).
	//  For information about how to format a JSON parameter for the various command line tool environments, see [Using JSON for Parameters](https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json). If your command-line tool or SDK requires quotation marks around the parameter, you should use single quotes to avoid confusion with the double quotes required in the JSON text.
	//  The following restrictions apply to tags:
	//   +  Maximum number of tags per secret: 50
	//   +  Maximum key length: 127 Unicode characters in UTF-8
	//   +  Maximum value length: 255 Unicode characters in UTF-8
	//   +  Tag keys and values are case sensitive.
	//   +  Do not use the ``aws:`` prefix in your tag names or values because AWS reserves it for AWS use. You can't edit or delete tag names or values with this prefix. Tags with this prefix do not count against your tags per secret limit.
	//   +  If you use your tagging schema across multiple services and resources, other services might have restrictions on allowed characters. Generally allowed characters: letters, spaces, and numbers representable in UTF-8, plus the following special characters: + - = . _ : / @.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a Secret resource.
type SecretArgs struct {
	// The description of the secret.
	Description pulumi.StringPtrInput
	// A structure that specifies how to generate a password to encrypt and store in the secret. To include a specific string in the secret, use ``SecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.
	//  We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.
	GenerateSecretString SecretGenerateSecretStringPtrInput
	// The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret. An alias is always prefixed by ``alias/``, for example ``alias/aws/secretsmanager``. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html).
	//  To use a KMS key in a different account, use the key ARN or the alias ARN.
	//  If you don't specify this value, then Secrets Manager uses the key ``aws/secretsmanager``. If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
	//  If the secret is in a different AWS account from the credentials calling the API, then you can't use ``aws/secretsmanager`` to encrypt the secret, and you must create and use a customer managed KMS key.
	KmsKeyId pulumi.StringPtrInput
	// The name of the new secret.
	//  The secret name can contain ASCII letters, numbers, and the following characters: /_+=.@-
	//  Do not end your secret name with a hyphen followed by six characters. If you do so, you risk confusion and unexpected results when searching for a secret by partial ARN. Secrets Manager automatically adds a hyphen and six random characters after the secret name at the end of the ARN.
	Name pulumi.StringPtrInput
	// A custom type that specifies a ``Region`` and the ``KmsKeyId`` for a replica secret.
	ReplicaRegions SecretReplicaRegionArrayInput
	// The text to encrypt and store in the secret. We recommend you use a JSON structure of key/value pairs for your secret value. To generate a random password, use ``GenerateSecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.
	SecretString pulumi.StringPtrInput
	// A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string, for example:
	//   ``[{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]``
	//  Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a different tag from one with key "abc".
	//  Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret.
	//  If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an ``Access Denied`` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#auth-and-access_tags2).
	//  For information about how to format a JSON parameter for the various command line tool environments, see [Using JSON for Parameters](https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json). If your command-line tool or SDK requires quotation marks around the parameter, you should use single quotes to avoid confusion with the double quotes required in the JSON text.
	//  The following restrictions apply to tags:
	//   +  Maximum number of tags per secret: 50
	//   +  Maximum key length: 127 Unicode characters in UTF-8
	//   +  Maximum value length: 255 Unicode characters in UTF-8
	//   +  Tag keys and values are case sensitive.
	//   +  Do not use the ``aws:`` prefix in your tag names or values because AWS reserves it for AWS use. You can't edit or delete tag names or values with this prefix. Tags with this prefix do not count against your tags per secret limit.
	//   +  If you use your tagging schema across multiple services and resources, other services might have restrictions on allowed characters. Generally allowed characters: letters, spaces, and numbers representable in UTF-8, plus the following special characters: + - = . _ : / @.
	Tags aws.TagArrayInput
}

func (SecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretArgs)(nil)).Elem()
}

type SecretInput interface {
	pulumi.Input

	ToSecretOutput() SecretOutput
	ToSecretOutputWithContext(ctx context.Context) SecretOutput
}

func (*Secret) ElementType() reflect.Type {
	return reflect.TypeOf((**Secret)(nil)).Elem()
}

func (i *Secret) ToSecretOutput() SecretOutput {
	return i.ToSecretOutputWithContext(context.Background())
}

func (i *Secret) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretOutput)
}

type SecretOutput struct{ *pulumi.OutputState }

func (SecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Secret)(nil)).Elem()
}

func (o SecretOutput) ToSecretOutput() SecretOutput {
	return o
}

func (o SecretOutput) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return o
}

func (o SecretOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// The description of the secret.
func (o SecretOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A structure that specifies how to generate a password to encrypt and store in the secret. To include a specific string in the secret, use “SecretString“ instead. If you omit both “GenerateSecretString“ and “SecretString“, you create an empty secret. When you make a change to this property, a new secret version is created.
//
//	We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.
func (o SecretOutput) GenerateSecretString() SecretGenerateSecretStringPtrOutput {
	return o.ApplyT(func(v *Secret) SecretGenerateSecretStringPtrOutput { return v.GenerateSecretString }).(SecretGenerateSecretStringPtrOutput)
}

// The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret. An alias is always prefixed by “alias/“, for example “alias/aws/secretsmanager“. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html).
//
//	To use a KMS key in a different account, use the key ARN or the alias ARN.
//	If you don't specify this value, then Secrets Manager uses the key ``aws/secretsmanager``. If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
//	If the secret is in a different AWS account from the credentials calling the API, then you can't use ``aws/secretsmanager`` to encrypt the secret, and you must create and use a customer managed KMS key.
func (o SecretOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// The name of the new secret.
//
//	The secret name can contain ASCII letters, numbers, and the following characters: /_+=.@-
//	Do not end your secret name with a hyphen followed by six characters. If you do so, you risk confusion and unexpected results when searching for a secret by partial ARN. Secrets Manager automatically adds a hyphen and six random characters after the secret name at the end of the ARN.
func (o SecretOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// A custom type that specifies a “Region“ and the “KmsKeyId“ for a replica secret.
func (o SecretOutput) ReplicaRegions() SecretReplicaRegionArrayOutput {
	return o.ApplyT(func(v *Secret) SecretReplicaRegionArrayOutput { return v.ReplicaRegions }).(SecretReplicaRegionArrayOutput)
}

// The text to encrypt and store in the secret. We recommend you use a JSON structure of key/value pairs for your secret value. To generate a random password, use “GenerateSecretString“ instead. If you omit both “GenerateSecretString“ and “SecretString“, you create an empty secret. When you make a change to this property, a new secret version is created.
func (o SecretOutput) SecretString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Secret) pulumi.StringPtrOutput { return v.SecretString }).(pulumi.StringPtrOutput)
}

// A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string, for example:
//
//	 ``[{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]``
//	Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a different tag from one with key "abc".
//	Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret.
//	If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an ``Access Denied`` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#auth-and-access_tags2).
//	For information about how to format a JSON parameter for the various command line tool environments, see [Using JSON for Parameters](https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json). If your command-line tool or SDK requires quotation marks around the parameter, you should use single quotes to avoid confusion with the double quotes required in the JSON text.
//	The following restrictions apply to tags:
//	 +  Maximum number of tags per secret: 50
//	 +  Maximum key length: 127 Unicode characters in UTF-8
//	 +  Maximum value length: 255 Unicode characters in UTF-8
//	 +  Tag keys and values are case sensitive.
//	 +  Do not use the ``aws:`` prefix in your tag names or values because AWS reserves it for AWS use. You can't edit or delete tag names or values with this prefix. Tags with this prefix do not count against your tags per secret limit.
//	 +  If you use your tagging schema across multiple services and resources, other services might have restrictions on allowed characters. Generally allowed characters: letters, spaces, and numbers representable in UTF-8, plus the following special characters: + - = . _ : / @.
func (o SecretOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Secret) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretInput)(nil)).Elem(), &Secret{})
	pulumi.RegisterOutputType(SecretOutput{})
}
