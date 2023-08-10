// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor.Outputs
{

    /// <summary>
    /// &lt;p&gt;Access configuration parameters.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class SourceLocationAccessConfiguration
    {
        public readonly Pulumi.AwsNative.MediaTailor.SourceLocationAccessType? AccessType;
        public readonly Outputs.SourceLocationSecretsManagerAccessTokenConfiguration? SecretsManagerAccessTokenConfiguration;

        [OutputConstructor]
        private SourceLocationAccessConfiguration(
            Pulumi.AwsNative.MediaTailor.SourceLocationAccessType? accessType,

            Outputs.SourceLocationSecretsManagerAccessTokenConfiguration? secretsManagerAccessTokenConfiguration)
        {
            AccessType = accessType;
            SecretsManagerAccessTokenConfiguration = secretsManagerAccessTokenConfiguration;
        }
    }
}
