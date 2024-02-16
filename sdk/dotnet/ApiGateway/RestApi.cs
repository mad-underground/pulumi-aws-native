// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    /// <summary>
    /// The ``AWS::ApiGateway::RestApi`` resource creates a REST API. For more information, see [restapi:create](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateRestApi.html) in the *Amazon API Gateway REST API Reference*.
    ///  On January 1, 2016, the Swagger Specification was donated to the [OpenAPI initiative](https://docs.aws.amazon.com/https://www.openapis.org/), becoming the foundation of the OpenAPI Specification.
    /// </summary>
    [AwsNativeResourceType("aws-native:apigateway:RestApi")]
    public partial class RestApi : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The source of the API key for metering requests according to a usage plan. Valid values are: ``HEADER`` to read the API key from the ``X-API-Key`` header of a request. ``AUTHORIZER`` to read the API key from the ``UsageIdentifierKey`` from a custom authorizer.
        /// </summary>
        [Output("apiKeySourceType")]
        public Output<string?> ApiKeySourceType { get; private set; } = null!;

        /// <summary>
        /// The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
        /// </summary>
        [Output("binaryMediaTypes")]
        public Output<ImmutableArray<string>> BinaryMediaTypes { get; private set; } = null!;

        /// <summary>
        /// An OpenAPI specification that defines a set of RESTful APIs in JSON format. For YAML templates, you can also provide the specification in YAML format.
        /// </summary>
        [Output("body")]
        public Output<object?> Body { get; private set; } = null!;

        /// <summary>
        /// The Amazon Simple Storage Service (Amazon S3) location that points to an OpenAPI file, which defines a set of RESTful APIs in JSON or YAML format.
        /// </summary>
        [Output("bodyS3Location")]
        public Output<Outputs.RestApiS3Location?> BodyS3Location { get; private set; } = null!;

        /// <summary>
        /// The ID of the RestApi that you want to clone from.
        /// </summary>
        [Output("cloneFrom")]
        public Output<string?> CloneFrom { get; private set; } = null!;

        /// <summary>
        /// The description of the RestApi.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies whether clients can invoke your API by using the default ``execute-api`` endpoint. By default, clients can invoke your API with the default ``https://{api_id}.execute-api.{region}.amazonaws.com`` endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint
        /// </summary>
        [Output("disableExecuteApiEndpoint")]
        public Output<bool?> DisableExecuteApiEndpoint { get; private set; } = null!;

        /// <summary>
        /// A list of the endpoint types of the API. Use this property when creating an API. When importing an existing API, specify the endpoint configuration types using the ``Parameters`` property.
        /// </summary>
        [Output("endpointConfiguration")]
        public Output<Outputs.RestApiEndpointConfiguration?> EndpointConfiguration { get; private set; } = null!;

        /// <summary>
        /// A query parameter to indicate whether to rollback the API update (``true``) or not (``false``) when a warning is encountered. The default value is ``false``.
        /// </summary>
        [Output("failOnWarnings")]
        public Output<bool?> FailOnWarnings { get; private set; } = null!;

        /// <summary>
        /// A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API. When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value. Setting it to zero allows compression for any payload size.
        /// </summary>
        [Output("minimumCompressionSize")]
        public Output<int?> MinimumCompressionSize { get; private set; } = null!;

        /// <summary>
        /// This property applies only when you use OpenAPI to define your REST API. The ``Mode`` determines how API Gateway handles resource updates.
        ///  Valid values are ``overwrite`` or ``merge``. 
        ///  For ``overwrite``, the new API definition replaces the existing one. The existing API identifier remains unchanged.
        ///   For ``merge``, the new API definition is merged with the existing API.
        ///  If you don't specify this property, a default value is chosen. For REST APIs created before March 29, 2021, the default is ``overwrite``. For REST APIs created after March 29, 2021, the new API definition takes precedence, but any container types such as endpoint configurations and binary media types are merged with the existing API. 
        ///  Use the default mode to define top-level ``RestApi`` properties in addition to using OpenAPI. Generally, it's preferred to use API Gateway's OpenAPI extensions to model these properties.
        /// </summary>
        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        /// <summary>
        /// The name of the RestApi. A name is required if the REST API is not based on an OpenAPI specification.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Custom header parameters as part of the request. For example, to exclude DocumentationParts from an imported API, set ``ignore=documentation`` as a ``parameters`` value, as in the AWS CLI command of ``aws apigateway import-rest-api --parameters ignore=documentation --body 'file:///path/to/imported-api-body.json'``.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// A policy document that contains the permissions for the ``RestApi`` resource. To set the ARN for the policy, use the ``!Join`` intrinsic function with ``""`` as delimiter and values of ``"execute-api:/"`` and ``"*"``.
        /// </summary>
        [Output("policy")]
        public Output<object?> Policy { get; private set; } = null!;

        [Output("restApiId")]
        public Output<string> RestApiId { get; private set; } = null!;

        [Output("rootResourceId")]
        public Output<string> RootResourceId { get; private set; } = null!;

        /// <summary>
        /// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The tag key can be up to 128 characters and must not start with ``aws:``. The tag value can be up to 256 characters.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.RestApiTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a RestApi resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RestApi(string name, RestApiArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:apigateway:RestApi", name, args ?? new RestApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RestApi(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:apigateway:RestApi", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RestApi resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RestApi Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RestApi(name, id, options);
        }
    }

    public sealed class RestApiArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The source of the API key for metering requests according to a usage plan. Valid values are: ``HEADER`` to read the API key from the ``X-API-Key`` header of a request. ``AUTHORIZER`` to read the API key from the ``UsageIdentifierKey`` from a custom authorizer.
        /// </summary>
        [Input("apiKeySourceType")]
        public Input<string>? ApiKeySourceType { get; set; }

        [Input("binaryMediaTypes")]
        private InputList<string>? _binaryMediaTypes;

        /// <summary>
        /// The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
        /// </summary>
        public InputList<string> BinaryMediaTypes
        {
            get => _binaryMediaTypes ?? (_binaryMediaTypes = new InputList<string>());
            set => _binaryMediaTypes = value;
        }

        /// <summary>
        /// An OpenAPI specification that defines a set of RESTful APIs in JSON format. For YAML templates, you can also provide the specification in YAML format.
        /// </summary>
        [Input("body")]
        public Input<object>? Body { get; set; }

        /// <summary>
        /// The Amazon Simple Storage Service (Amazon S3) location that points to an OpenAPI file, which defines a set of RESTful APIs in JSON or YAML format.
        /// </summary>
        [Input("bodyS3Location")]
        public Input<Inputs.RestApiS3LocationArgs>? BodyS3Location { get; set; }

        /// <summary>
        /// The ID of the RestApi that you want to clone from.
        /// </summary>
        [Input("cloneFrom")]
        public Input<string>? CloneFrom { get; set; }

        /// <summary>
        /// The description of the RestApi.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether clients can invoke your API by using the default ``execute-api`` endpoint. By default, clients can invoke your API with the default ``https://{api_id}.execute-api.{region}.amazonaws.com`` endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint
        /// </summary>
        [Input("disableExecuteApiEndpoint")]
        public Input<bool>? DisableExecuteApiEndpoint { get; set; }

        /// <summary>
        /// A list of the endpoint types of the API. Use this property when creating an API. When importing an existing API, specify the endpoint configuration types using the ``Parameters`` property.
        /// </summary>
        [Input("endpointConfiguration")]
        public Input<Inputs.RestApiEndpointConfigurationArgs>? EndpointConfiguration { get; set; }

        /// <summary>
        /// A query parameter to indicate whether to rollback the API update (``true``) or not (``false``) when a warning is encountered. The default value is ``false``.
        /// </summary>
        [Input("failOnWarnings")]
        public Input<bool>? FailOnWarnings { get; set; }

        /// <summary>
        /// A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API. When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value. Setting it to zero allows compression for any payload size.
        /// </summary>
        [Input("minimumCompressionSize")]
        public Input<int>? MinimumCompressionSize { get; set; }

        /// <summary>
        /// This property applies only when you use OpenAPI to define your REST API. The ``Mode`` determines how API Gateway handles resource updates.
        ///  Valid values are ``overwrite`` or ``merge``. 
        ///  For ``overwrite``, the new API definition replaces the existing one. The existing API identifier remains unchanged.
        ///   For ``merge``, the new API definition is merged with the existing API.
        ///  If you don't specify this property, a default value is chosen. For REST APIs created before March 29, 2021, the default is ``overwrite``. For REST APIs created after March 29, 2021, the new API definition takes precedence, but any container types such as endpoint configurations and binary media types are merged with the existing API. 
        ///  Use the default mode to define top-level ``RestApi`` properties in addition to using OpenAPI. Generally, it's preferred to use API Gateway's OpenAPI extensions to model these properties.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The name of the RestApi. A name is required if the REST API is not based on an OpenAPI specification.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Custom header parameters as part of the request. For example, to exclude DocumentationParts from an imported API, set ``ignore=documentation`` as a ``parameters`` value, as in the AWS CLI command of ``aws apigateway import-rest-api --parameters ignore=documentation --body 'file:///path/to/imported-api-body.json'``.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// A policy document that contains the permissions for the ``RestApi`` resource. To set the ARN for the policy, use the ``!Join`` intrinsic function with ``""`` as delimiter and values of ``"execute-api:/"`` and ``"*"``.
        /// </summary>
        [Input("policy")]
        public Input<object>? Policy { get; set; }

        [Input("tags")]
        private InputList<Inputs.RestApiTagArgs>? _tags;

        /// <summary>
        /// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The tag key can be up to 128 characters and must not start with ``aws:``. The tag value can be up to 256 characters.
        /// </summary>
        public InputList<Inputs.RestApiTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.RestApiTagArgs>());
            set => _tags = value;
        }

        public RestApiArgs()
        {
        }
        public static new RestApiArgs Empty => new RestApiArgs();
    }
}
