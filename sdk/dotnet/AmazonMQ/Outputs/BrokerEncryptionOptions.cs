// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AmazonMQ.Outputs
{

    [OutputType]
    public sealed class BrokerEncryptionOptions
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-encryptionoptions.html#cfn-amazonmq-broker-encryptionoptions-kmskeyid
        /// </summary>
        public readonly string? KmsKeyId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-broker-encryptionoptions.html#cfn-amazonmq-broker-encryptionoptions-useawsownedkey
        /// </summary>
        public readonly bool UseAwsOwnedKey;

        [OutputConstructor]
        private BrokerEncryptionOptions(
            string? KmsKeyId,

            bool UseAwsOwnedKey)
        {
            this.KmsKeyId = KmsKeyId;
            this.UseAwsOwnedKey = UseAwsOwnedKey;
        }
    }
}