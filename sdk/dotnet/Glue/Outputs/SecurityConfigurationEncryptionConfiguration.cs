// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    [OutputType]
    public sealed class SecurityConfigurationEncryptionConfiguration
    {
        public readonly Outputs.SecurityConfigurationCloudWatchEncryption? CloudWatchEncryption;
        public readonly Outputs.SecurityConfigurationJobBookmarksEncryption? JobBookmarksEncryption;
        public readonly Outputs.SecurityConfigurationS3Encryptions? S3Encryptions;

        [OutputConstructor]
        private SecurityConfigurationEncryptionConfiguration(
            Outputs.SecurityConfigurationCloudWatchEncryption? cloudWatchEncryption,

            Outputs.SecurityConfigurationJobBookmarksEncryption? jobBookmarksEncryption,

            Outputs.SecurityConfigurationS3Encryptions? s3Encryptions)
        {
            CloudWatchEncryption = cloudWatchEncryption;
            JobBookmarksEncryption = jobBookmarksEncryption;
            S3Encryptions = s3Encryptions;
        }
    }
}
