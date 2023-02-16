// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class ConnectorProfileInforNexusConnectorProfileCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Access Key portion of the credentials.
        /// </summary>
        [Input("accessKeyId", required: true)]
        public Input<string> AccessKeyId { get; set; } = null!;

        /// <summary>
        /// The encryption keys used to encrypt data.
        /// </summary>
        [Input("datakey", required: true)]
        public Input<string> Datakey { get; set; } = null!;

        /// <summary>
        /// The secret key used to sign requests.
        /// </summary>
        [Input("secretAccessKey", required: true)]
        public Input<string> SecretAccessKey { get; set; } = null!;

        /// <summary>
        /// The identiﬁer for the user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public ConnectorProfileInforNexusConnectorProfileCredentialsArgs()
        {
        }
        public static new ConnectorProfileInforNexusConnectorProfileCredentialsArgs Empty => new ConnectorProfileInforNexusConnectorProfileCredentialsArgs();
    }
}
