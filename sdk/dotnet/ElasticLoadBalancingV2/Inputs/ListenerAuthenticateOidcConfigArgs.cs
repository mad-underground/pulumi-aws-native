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
    /// Specifies information required using an identity provide (IdP) that is compliant with OpenID Connect (OIDC) to authenticate users.
    /// </summary>
    public sealed class ListenerAuthenticateOidcConfigArgs : global::Pulumi.ResourceArgs
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
        /// The authorization endpoint of the IdP. This must be a full URL, including the HTTPS protocol, the domain, and the path.
        /// </summary>
        [Input("authorizationEndpoint", required: true)]
        public Input<string> AuthorizationEndpoint { get; set; } = null!;

        /// <summary>
        /// The OAuth 2.0 client identifier.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The OAuth 2.0 client secret. This parameter is required if you are creating a rule. If you are modifying a rule, you can omit this parameter if you set ``UseExistingClientSecret`` to true.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// The OIDC issuer identifier of the IdP. This must be a full URL, including the HTTPS protocol, the domain, and the path.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

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
        /// The token endpoint of the IdP. This must be a full URL, including the HTTPS protocol, the domain, and the path.
        /// </summary>
        [Input("tokenEndpoint", required: true)]
        public Input<string> TokenEndpoint { get; set; } = null!;

        /// <summary>
        /// Indicates whether to use the existing client secret when modifying a rule. If you are creating a rule, you can omit this parameter or set it to false.
        /// </summary>
        [Input("useExistingClientSecret")]
        public Input<bool>? UseExistingClientSecret { get; set; }

        /// <summary>
        /// The user info endpoint of the IdP. This must be a full URL, including the HTTPS protocol, the domain, and the path.
        /// </summary>
        [Input("userInfoEndpoint", required: true)]
        public Input<string> UserInfoEndpoint { get; set; } = null!;

        public ListenerAuthenticateOidcConfigArgs()
        {
        }
        public static new ListenerAuthenticateOidcConfigArgs Empty => new ListenerAuthenticateOidcConfigArgs();
    }
}
