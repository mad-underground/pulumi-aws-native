// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor.Inputs
{

    /// <summary>
    /// &lt;p&gt;AWS Secrets Manager access token configuration parameters. For information about Secrets Manager access token authentication, see &lt;a href="https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-access-configuration-access-token.html"&gt;Working with AWS Secrets Manager access token authentication&lt;/a&gt;.&lt;/p&gt;
    /// </summary>
    public sealed class SourceLocationSecretsManagerAccessTokenConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The name of the HTTP header used to supply the access token in requests to the source location.&lt;/p&gt;
        /// </summary>
        [Input("headerName")]
        public Input<string>? HeaderName { get; set; }

        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the access token.&lt;/p&gt;
        /// </summary>
        [Input("secretArn")]
        public Input<string>? SecretArn { get; set; }

        /// <summary>
        /// &lt;p&gt;The AWS Secrets Manager &lt;a href="https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_CreateSecret.html#SecretsManager-CreateSecret-request-SecretString.html"&gt;SecretString&lt;/a&gt; key associated with the access token. MediaTailor uses the key to look up SecretString key and value pair containing the access token.&lt;/p&gt;
        /// </summary>
        [Input("secretStringKey")]
        public Input<string>? SecretStringKey { get; set; }

        public SourceLocationSecretsManagerAccessTokenConfigurationArgs()
        {
        }
        public static new SourceLocationSecretsManagerAccessTokenConfigurationArgs Empty => new SourceLocationSecretsManagerAccessTokenConfigurationArgs();
    }
}
