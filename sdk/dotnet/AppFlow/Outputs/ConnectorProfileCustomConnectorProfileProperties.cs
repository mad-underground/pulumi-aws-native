// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class ConnectorProfileCustomConnectorProfileProperties
    {
        public readonly Outputs.ConnectorProfileOAuth2Properties? OAuth2Properties;
        public readonly ImmutableDictionary<string, string>? ProfileProperties;

        [OutputConstructor]
        private ConnectorProfileCustomConnectorProfileProperties(
            Outputs.ConnectorProfileOAuth2Properties? oAuth2Properties,

            ImmutableDictionary<string, string>? profileProperties)
        {
            OAuth2Properties = oAuth2Properties;
            ProfileProperties = profileProperties;
        }
    }
}
