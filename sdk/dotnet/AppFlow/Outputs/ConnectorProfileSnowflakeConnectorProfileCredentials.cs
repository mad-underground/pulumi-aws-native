// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppFlow.Outputs
{

    [OutputType]
    public sealed class ConnectorProfileSnowflakeConnectorProfileCredentials
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-snowflakeconnectorprofilecredentials.html#cfn-appflow-connectorprofile-snowflakeconnectorprofilecredentials-password
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-snowflakeconnectorprofilecredentials.html#cfn-appflow-connectorprofile-snowflakeconnectorprofilecredentials-username
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private ConnectorProfileSnowflakeConnectorProfileCredentials(
            string Password,

            string Username)
        {
            this.Password = Password;
            this.Username = Username;
        }
    }
}
