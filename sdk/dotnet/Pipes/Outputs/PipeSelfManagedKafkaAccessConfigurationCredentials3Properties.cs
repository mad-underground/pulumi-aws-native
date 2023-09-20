// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Outputs
{

    [OutputType]
    public sealed class PipeSelfManagedKafkaAccessConfigurationCredentials3Properties
    {
        /// <summary>
        /// Optional SecretManager ARN which stores the database credentials
        /// </summary>
        public readonly string ClientCertificateTlsAuth;

        [OutputConstructor]
        private PipeSelfManagedKafkaAccessConfigurationCredentials3Properties(string clientCertificateTlsAuth)
        {
            ClientCertificateTlsAuth = clientCertificateTlsAuth;
        }
    }
}