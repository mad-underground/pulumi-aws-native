// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Outputs
{

    /// <summary>
    /// Specifies information required when integrating with Amazon Cognito to authenticate users.
    /// </summary>
    [OutputType]
    public sealed class ListenerAuthenticateCognitoConfig
    {
        /// <summary>
        /// The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? AuthenticationRequestExtraParams;
        /// <summary>
        /// The behavior if the user is not authenticated. The following are possible values:
        ///   +  deny```` - Return an HTTP 401 Unauthorized error.
        ///   +  allow```` - Allow the request to be forwarded to the target.
        ///   +  authenticate```` - Redirect the request to the IdP authorization endpoint. This is the default value.
        /// </summary>
        public readonly string? OnUnauthenticatedRequest;
        /// <summary>
        /// The set of user claims to be requested from the IdP. The default is ``openid``.
        ///  To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
        /// </summary>
        public readonly string? Scope;
        /// <summary>
        /// The name of the cookie used to maintain session information. The default is AWSELBAuthSessionCookie.
        /// </summary>
        public readonly string? SessionCookieName;
        /// <summary>
        /// The maximum duration of the authentication session, in seconds. The default is 604800 seconds (7 days).
        /// </summary>
        public readonly string? SessionTimeout;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon Cognito user pool.
        /// </summary>
        public readonly string UserPoolArn;
        /// <summary>
        /// The ID of the Amazon Cognito user pool client.
        /// </summary>
        public readonly string UserPoolClientId;
        /// <summary>
        /// The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.
        /// </summary>
        public readonly string UserPoolDomain;

        [OutputConstructor]
        private ListenerAuthenticateCognitoConfig(
            ImmutableDictionary<string, string>? authenticationRequestExtraParams,

            string? onUnauthenticatedRequest,

            string? scope,

            string? sessionCookieName,

            string? sessionTimeout,

            string userPoolArn,

            string userPoolClientId,

            string userPoolDomain)
        {
            AuthenticationRequestExtraParams = authenticationRequestExtraParams;
            OnUnauthenticatedRequest = onUnauthenticatedRequest;
            Scope = scope;
            SessionCookieName = sessionCookieName;
            SessionTimeout = sessionTimeout;
            UserPoolArn = userPoolArn;
            UserPoolClientId = userPoolClientId;
            UserPoolDomain = userPoolDomain;
        }
    }
}
