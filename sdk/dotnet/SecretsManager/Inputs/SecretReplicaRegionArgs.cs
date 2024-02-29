// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecretsManager.Inputs
{

    /// <summary>
    /// Specifies a ``Region`` and the ``KmsKeyId`` for a replica secret.
    /// </summary>
    public sealed class SecretReplicaRegionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN, key ID, or alias of the KMS key to encrypt the secret. If you don't include this field, Secrets Manager uses ``aws/secretsmanager``.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// A string that represents a ``Region``, for example "us-east-1".
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public SecretReplicaRegionArgs()
        {
        }
        public static new SecretReplicaRegionArgs Empty => new SecretReplicaRegionArgs();
    }
}
