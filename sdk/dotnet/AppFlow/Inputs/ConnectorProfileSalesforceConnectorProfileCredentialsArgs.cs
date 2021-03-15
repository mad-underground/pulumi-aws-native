// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-salesforceconnectorprofilecredentials.html
    /// </summary>
    public sealed class ConnectorProfileSalesforceConnectorProfileCredentialsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-salesforceconnectorprofilecredentials.html#cfn-appflow-connectorprofile-salesforceconnectorprofilecredentials-accesstoken
        /// </summary>
        [Input("accessToken")]
        public Input<string>? AccessToken { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-salesforceconnectorprofilecredentials.html#cfn-appflow-connectorprofile-salesforceconnectorprofilecredentials-clientcredentialsarn
        /// </summary>
        [Input("clientCredentialsArn")]
        public Input<string>? ClientCredentialsArn { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-salesforceconnectorprofilecredentials.html#cfn-appflow-connectorprofile-salesforceconnectorprofilecredentials-connectoroauthrequest
        /// </summary>
        [Input("connectorOAuthRequest")]
        public Input<Inputs.ConnectorProfileConnectorOAuthRequestArgs>? ConnectorOAuthRequest { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-salesforceconnectorprofilecredentials.html#cfn-appflow-connectorprofile-salesforceconnectorprofilecredentials-refreshtoken
        /// </summary>
        [Input("refreshToken")]
        public Input<string>? RefreshToken { get; set; }

        public ConnectorProfileSalesforceConnectorProfileCredentialsArgs()
        {
        }
    }
}
