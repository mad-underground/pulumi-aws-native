// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VerifiedPermissions.Outputs
{

    [OutputType]
    public sealed class IdentitySourceCognitoUserPoolConfiguration
    {
        public readonly ImmutableArray<string> ClientIds;
        public readonly string UserPoolArn;

        [OutputConstructor]
        private IdentitySourceCognitoUserPoolConfiguration(
            ImmutableArray<string> clientIds,

            string userPoolArn)
        {
            ClientIds = clientIds;
            UserPoolArn = userPoolArn;
        }
    }
}
