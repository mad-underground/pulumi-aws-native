// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Athena.Outputs
{

    [OutputType]
    public sealed class WorkGroupEncryptionConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-encryptionconfiguration.html#cfn-athena-workgroup-encryptionconfiguration-encryptionoption
        /// </summary>
        public readonly string EncryptionOption;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-encryptionconfiguration.html#cfn-athena-workgroup-encryptionconfiguration-kmskey
        /// </summary>
        public readonly string? KmsKey;

        [OutputConstructor]
        private WorkGroupEncryptionConfiguration(
            string encryptionOption,

            string? kmsKey)
        {
            EncryptionOption = encryptionOption;
            KmsKey = kmsKey;
        }
    }
}
