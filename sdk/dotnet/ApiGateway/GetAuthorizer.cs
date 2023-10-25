// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    public static class GetAuthorizer
    {
        /// <summary>
        /// The ``AWS::ApiGateway::Authorizer`` resource creates an authorization layer that API Gateway activates for methods that have authorization enabled. API Gateway activates the authorizer when a client calls those methods.
        /// </summary>
        public static Task<GetAuthorizerResult> InvokeAsync(GetAuthorizerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAuthorizerResult>("aws-native:apigateway:getAuthorizer", args ?? new GetAuthorizerArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::ApiGateway::Authorizer`` resource creates an authorization layer that API Gateway activates for methods that have authorization enabled. API Gateway activates the authorizer when a client calls those methods.
        /// </summary>
        public static Output<GetAuthorizerResult> Invoke(GetAuthorizerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthorizerResult>("aws-native:apigateway:getAuthorizer", args ?? new GetAuthorizerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthorizerArgs : global::Pulumi.InvokeArgs
    {
        [Input("authorizerId", required: true)]
        public string AuthorizerId { get; set; } = null!;

        /// <summary>
        /// The string identifier of the associated RestApi.
        /// </summary>
        [Input("restApiId", required: true)]
        public string RestApiId { get; set; } = null!;

        public GetAuthorizerArgs()
        {
        }
        public static new GetAuthorizerArgs Empty => new GetAuthorizerArgs();
    }

    public sealed class GetAuthorizerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("authorizerId", required: true)]
        public Input<string> AuthorizerId { get; set; } = null!;

        /// <summary>
        /// The string identifier of the associated RestApi.
        /// </summary>
        [Input("restApiId", required: true)]
        public Input<string> RestApiId { get; set; } = null!;

        public GetAuthorizerInvokeArgs()
        {
        }
        public static new GetAuthorizerInvokeArgs Empty => new GetAuthorizerInvokeArgs();
    }


    [OutputType]
    public sealed class GetAuthorizerResult
    {
        /// <summary>
        /// Optional customer-defined field, used in OpenAPI imports and exports without functional impact.
        /// </summary>
        public readonly string? AuthType;
        /// <summary>
        /// Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null.
        /// </summary>
        public readonly string? AuthorizerCredentials;
        public readonly string? AuthorizerId;
        /// <summary>
        /// The TTL in seconds of cached authorizer results. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway will cache authorizer responses. If this field is not set, the default value is 300. The maximum value is 3600, or 1 hour.
        /// </summary>
        public readonly int? AuthorizerResultTtlInSeconds;
        /// <summary>
        /// Specifies the authorizer's Uniform Resource Identifier (URI). For ``TOKEN`` or ``REQUEST`` authorizers, this must be a well-formed Lambda function URI, for example, ``arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations``. In general, the URI has this form ``arn:aws:apigateway:{region}:lambda:path/{service_api}``, where ``{region}`` is the same as the region hosting the Lambda function, ``path`` indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial ``/``. For Lambda functions, this is usually of the form ``/2015-03-31/functions/[FunctionARN]/invocations``.
        /// </summary>
        public readonly string? AuthorizerUri;
        /// <summary>
        /// The identity source for which authorization is requested. For a ``TOKEN`` or ``COGNITO_USER_POOLS`` authorizer, this is required and specifies the request header mapping expression for the custom header holding the authorization token submitted by the client. For example, if the token header name is ``Auth``, the header mapping expression is ``method.request.header.Auth``. For the ``REQUEST`` authorizer, this is required when authorization caching is enabled. The value is a comma-separated string of one or more mapping expressions of the specified request parameters. For example, if an ``Auth`` header, a ``Name`` query string parameter are defined as identity sources, this value is ``method.request.header.Auth, method.request.querystring.Name``. These parameters will be used to derive the authorization caching key and to perform runtime validation of the ``REQUEST`` authorizer by verifying all of the identity-related request parameters are present, not null and non-empty. Only when this is true does the authorizer invoke the authorizer Lambda function, otherwise, it returns a 401 Unauthorized response without calling the Lambda function. The valid value is a string of comma-separated mapping expressions of the specified request parameters. When the authorization caching is not enabled, this property is optional.
        /// </summary>
        public readonly string? IdentitySource;
        /// <summary>
        /// A validation expression for the incoming identity token. For ``TOKEN`` authorizers, this value is a regular expression. For ``COGNITO_USER_POOLS`` authorizers, API Gateway will match the ``aud`` field of the incoming token from the client against the specified regular expression. It will invoke the authorizer's Lambda function when there is a match. Otherwise, it will return a 401 Unauthorized response without calling the Lambda function. The validation expression does not apply to the ``REQUEST`` authorizer.
        /// </summary>
        public readonly string? IdentityValidationExpression;
        /// <summary>
        /// The name of the authorizer.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A list of the Amazon Cognito user pool ARNs for the ``COGNITO_USER_POOLS`` authorizer. Each element is of this format: ``arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}``. For a ``TOKEN`` or ``REQUEST`` authorizer, this is not defined.
        /// </summary>
        public readonly ImmutableArray<string> ProviderArns;
        /// <summary>
        /// The authorizer type. Valid values are ``TOKEN`` for a Lambda function using a single authorization token submitted in a custom header, ``REQUEST`` for a Lambda function using incoming request parameters, and ``COGNITO_USER_POOLS`` for using an Amazon Cognito user pool.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetAuthorizerResult(
            string? authType,

            string? authorizerCredentials,

            string? authorizerId,

            int? authorizerResultTtlInSeconds,

            string? authorizerUri,

            string? identitySource,

            string? identityValidationExpression,

            string? name,

            ImmutableArray<string> providerArns,

            string? type)
        {
            AuthType = authType;
            AuthorizerCredentials = authorizerCredentials;
            AuthorizerId = authorizerId;
            AuthorizerResultTtlInSeconds = authorizerResultTtlInSeconds;
            AuthorizerUri = authorizerUri;
            IdentitySource = identitySource;
            IdentityValidationExpression = identityValidationExpression;
            Name = name;
            ProviderArns = providerArns;
            Type = type;
        }
    }
}
