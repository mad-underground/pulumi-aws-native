// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DMS.Outputs
{

    [OutputType]
    public sealed class EndpointMongoDbSettings
    {
        public readonly string? AuthMechanism;
        public readonly string? AuthSource;
        public readonly string? AuthType;
        public readonly string? DatabaseName;
        public readonly string? DocsToInvestigate;
        public readonly string? ExtractDocId;
        public readonly string? NestingLevel;
        public readonly string? Password;
        public readonly int? Port;
        public readonly string? SecretsManagerAccessRoleArn;
        public readonly string? SecretsManagerSecretId;
        public readonly string? ServerName;
        public readonly string? Username;

        [OutputConstructor]
        private EndpointMongoDbSettings(
            string? authMechanism,

            string? authSource,

            string? authType,

            string? databaseName,

            string? docsToInvestigate,

            string? extractDocId,

            string? nestingLevel,

            string? password,

            int? port,

            string? secretsManagerAccessRoleArn,

            string? secretsManagerSecretId,

            string? serverName,

            string? username)
        {
            AuthMechanism = authMechanism;
            AuthSource = authSource;
            AuthType = authType;
            DatabaseName = databaseName;
            DocsToInvestigate = docsToInvestigate;
            ExtractDocId = extractDocId;
            NestingLevel = nestingLevel;
            Password = password;
            Port = port;
            SecretsManagerAccessRoleArn = secretsManagerAccessRoleArn;
            SecretsManagerSecretId = secretsManagerSecretId;
            ServerName = serverName;
            Username = username;
        }
    }
}
