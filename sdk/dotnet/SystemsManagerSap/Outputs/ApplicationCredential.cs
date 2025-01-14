// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SystemsManagerSap.Outputs
{

    [OutputType]
    public sealed class ApplicationCredential
    {
        public readonly Pulumi.AwsNative.SystemsManagerSap.ApplicationCredentialCredentialType? CredentialType;
        public readonly string? DatabaseName;
        public readonly string? SecretId;

        [OutputConstructor]
        private ApplicationCredential(
            Pulumi.AwsNative.SystemsManagerSap.ApplicationCredentialCredentialType? credentialType,

            string? databaseName,

            string? secretId)
        {
            CredentialType = credentialType;
            DatabaseName = databaseName;
            SecretId = secretId;
        }
    }
}
