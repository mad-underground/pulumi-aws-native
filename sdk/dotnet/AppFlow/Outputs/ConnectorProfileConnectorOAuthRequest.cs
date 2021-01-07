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
    public sealed class ConnectorProfileConnectorOAuthRequest
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectoroauthrequest.html#cfn-appflow-connectorprofile-connectoroauthrequest-authcode
        /// </summary>
        public readonly string? AuthCode;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-connectoroauthrequest.html#cfn-appflow-connectorprofile-connectoroauthrequest-redirecturi
        /// </summary>
        public readonly string? RedirectUri;

        [OutputConstructor]
        private ConnectorProfileConnectorOAuthRequest(
            string? AuthCode,

            string? RedirectUri)
        {
            this.AuthCode = AuthCode;
            this.RedirectUri = RedirectUri;
        }
    }
}
