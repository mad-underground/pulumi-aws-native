// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Inputs
{

    /// <summary>
    /// Specifies information required when integrating with Amazon Cognito to authenticate users.
    /// </summary>
    public sealed class ListenerAuthenticateCognitoConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("authenticationRequestExtraParams")]
        private InputMap<string>? _authenticationRequestExtraParams;

        /// <summary>
        /// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
        /// </summary>
        public InputMap<string> AuthenticationRequestExtraParams
        {
            get => _authenticationRequestExtraParams ?? (_authenticationRequestExtraParams = new InputMap<string>());
            set => _authenticationRequestExtraParams = value;
        }

        /// <summary>
        /// The behavior if the user is not authenticated. The following are possible values:
        ///   +  deny```` - Return an HTTP 401 Unauthorized error.
        ///   +  allow```` - Allow the request to be forwarded to the target.
        ///   +  authenticate```` - Redirect the request to the IdP authorization endpoint. This is the default value.
        /// </summary>
        [Input("onUnauthenticatedRequest")]
        public Input<string>? OnUnauthenticatedRequest { get; set; }

        /// <summary>
        /// The set of user claims to be requested from the IdP. The default is ``openid``.
        ///  To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The name of the cookie used to maintain session information. The default is AWSELBAuthSessionCookie.
        /// </summary>
        [Input("sessionCookieName")]
        public Input<string>? SessionCookieName { get; set; }

        /// <summary>
        /// The maximum duration of the authentication session, in seconds. The default is 604800 seconds (7 days).
        /// </summary>
        [Input("sessionTimeout")]
        public Input<string>? SessionTimeout { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon Cognito user pool.
        /// </summary>
        [Input("userPoolArn", required: true)]
        public Input<string> UserPoolArn { get; set; } = null!;

        /// <summary>
        /// The ID of the Amazon Cognito user pool client.
        /// </summary>
        [Input("userPoolClientId", required: true)]
        public Input<string> UserPoolClientId { get; set; } = null!;

        /// <summary>
        /// The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.
        /// </summary>
        [Input("userPoolDomain", required: true)]
        public Input<string> UserPoolDomain { get; set; } = null!;

        public ListenerAuthenticateCognitoConfigArgs()
        {
        }
        public static new ListenerAuthenticateCognitoConfigArgs Empty => new ListenerAuthenticateCognitoConfigArgs();
    }
}
