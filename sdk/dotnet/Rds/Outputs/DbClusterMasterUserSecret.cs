// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rds.Outputs
{

    [OutputType]
    public sealed class DbClusterMasterUserSecret
    {
        /// <summary>
        /// The AWS KMS key identifier that is used to encrypt the secret.
        /// </summary>
        public readonly string? KmsKeyId;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the secret.
        /// </summary>
        public readonly string? SecretArn;

        [OutputConstructor]
        private DbClusterMasterUserSecret(
            string? kmsKeyId,

            string? secretArn)
        {
            KmsKeyId = kmsKeyId;
            SecretArn = secretArn;
        }
    }
}