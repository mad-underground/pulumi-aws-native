// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecretsManager
{
    /// <summary>
    /// Creates a new secret. A *secret* can be a password, a set of credentials such as a user name and password, an OAuth token, or other secret information that you store in an encrypted form in Secrets Manager.
    ///  For RDS master user credentials, see [AWS::RDS::DBCluster MasterUserSecret](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-masterusersecret.html).
    ///  To retrieve a secret in a CFNshort template, use a *dynamic reference*. For more information, see [Retrieve a secret in an resource](https://docs.aws.amazon.com/secretsmanager/latest/userguide/cfn-example_reference-secret.html).
    ///  A common scenario is to first create a secret with ``GenerateSecretString``, which generates a password, and then use a dynamic reference to retrieve the username and password from the secret to use as credentials for a new database. See the example *Creating a Redshift cluster and a secret for the admin credentials*.
    ///  For information about creating a secret in the c
    /// </summary>
    [AwsNativeResourceType("aws-native:secretsmanager:Secret")]
    public partial class Secret : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// The description of the secret.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A structure that specifies how to generate a password to encrypt and store in the secret. To include a specific string in the secret, use ``SecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.
        ///  We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.
        /// </summary>
        [Output("generateSecretString")]
        public Output<Outputs.SecretGenerateSecretString?> GenerateSecretString { get; private set; } = null!;

        /// <summary>
        /// The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret. An alias is always prefixed by ``alias/``, for example ``alias/aws/secretsmanager``. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html).
        ///  To use a KMS key in a different account, use the key ARN or the alias ARN.
        ///  If you don't specify this value, then Secrets Manager uses the key ``aws/secretsmanager``. If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
        ///  If the secret is in a different AWS account from the credentials calling the API, then you can't use ``aws/secretsmanager`` to encrypt the secret, and you must create and use a customer managed KMS key.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string?> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// The name of the new secret.
        ///  The secret name can contain ASCII letters, numbers, and the following characters: /_+=.@-
        ///  Do not end your secret name with a hyphen followed by six characters. If you do so, you risk confusion and unexpected results when searching for a secret by partial ARN. Secrets Manager automatically adds a hyphen and six random characters after the secret name at the end of the ARN.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// A custom type that specifies a ``Region`` and the ``KmsKeyId`` for a replica secret.
        /// </summary>
        [Output("replicaRegions")]
        public Output<ImmutableArray<Outputs.SecretReplicaRegion>> ReplicaRegions { get; private set; } = null!;

        /// <summary>
        /// The text to encrypt and store in the secret. We recommend you use a JSON structure of key/value pairs for your secret value. To generate a random password, use ``GenerateSecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.
        /// </summary>
        [Output("secretString")]
        public Output<string?> SecretString { get; private set; } = null!;

        /// <summary>
        /// A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string, for example:
        ///   ``[{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]`` 
        ///  Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a different tag from one with key "abc".
        ///  Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret. 
        ///  If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an ``Access Denied`` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazo
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Secret resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Secret(string name, SecretArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:secretsmanager:Secret", name, args ?? new SecretArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Secret(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:secretsmanager:Secret", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Secret resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Secret Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Secret(name, id, options);
        }
    }

    public sealed class SecretArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the secret.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A structure that specifies how to generate a password to encrypt and store in the secret. To include a specific string in the secret, use ``SecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.
        ///  We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.
        /// </summary>
        [Input("generateSecretString")]
        public Input<Inputs.SecretGenerateSecretStringArgs>? GenerateSecretString { get; set; }

        /// <summary>
        /// The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret. An alias is always prefixed by ``alias/``, for example ``alias/aws/secretsmanager``. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html).
        ///  To use a KMS key in a different account, use the key ARN or the alias ARN.
        ///  If you don't specify this value, then Secrets Manager uses the key ``aws/secretsmanager``. If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
        ///  If the secret is in a different AWS account from the credentials calling the API, then you can't use ``aws/secretsmanager`` to encrypt the secret, and you must create and use a customer managed KMS key.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// The name of the new secret.
        ///  The secret name can contain ASCII letters, numbers, and the following characters: /_+=.@-
        ///  Do not end your secret name with a hyphen followed by six characters. If you do so, you risk confusion and unexpected results when searching for a secret by partial ARN. Secrets Manager automatically adds a hyphen and six random characters after the secret name at the end of the ARN.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("replicaRegions")]
        private InputList<Inputs.SecretReplicaRegionArgs>? _replicaRegions;

        /// <summary>
        /// A custom type that specifies a ``Region`` and the ``KmsKeyId`` for a replica secret.
        /// </summary>
        public InputList<Inputs.SecretReplicaRegionArgs> ReplicaRegions
        {
            get => _replicaRegions ?? (_replicaRegions = new InputList<Inputs.SecretReplicaRegionArgs>());
            set => _replicaRegions = value;
        }

        /// <summary>
        /// The text to encrypt and store in the secret. We recommend you use a JSON structure of key/value pairs for your secret value. To generate a random password, use ``GenerateSecretString`` instead. If you omit both ``GenerateSecretString`` and ``SecretString``, you create an empty secret. When you make a change to this property, a new secret version is created.
        /// </summary>
        [Input("secretString")]
        public Input<string>? SecretString { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string, for example:
        ///   ``[{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]`` 
        ///  Secrets Manager tag key names are case sensitive. A tag with the key "ABC" is a different tag from one with key "abc".
        ///  Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret. 
        ///  If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an ``Access Denied`` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazo
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public SecretArgs()
        {
        }
        public static new SecretArgs Empty => new SecretArgs();
    }
}
