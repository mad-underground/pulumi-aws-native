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
    /// The ``AWS::ApiGateway::Stage`` resource creates a stage for a deployment.
    /// </summary>
    [AwsNativeResourceType("aws-native:apigateway:Stage")]
    public partial class Stage : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access log settings, including the access log format and access log destination ARN.
        /// </summary>
        [Output("accessLogSetting")]
        public Output<Outputs.StageAccessLogSetting?> AccessLogSetting { get; private set; } = null!;

        /// <summary>
        /// Specifies whether a cache cluster is enabled for the stage.
        /// </summary>
        [Output("cacheClusterEnabled")]
        public Output<bool?> CacheClusterEnabled { get; private set; } = null!;

        /// <summary>
        /// The stage's cache capacity in GB. For more information about choosing a cache size, see [Enabling API caching to enhance responsiveness](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html).
        /// </summary>
        [Output("cacheClusterSize")]
        public Output<string?> CacheClusterSize { get; private set; } = null!;

        /// <summary>
        /// Settings for the canary deployment in this stage.
        /// </summary>
        [Output("canarySetting")]
        public Output<Outputs.StageCanarySetting?> CanarySetting { get; private set; } = null!;

        /// <summary>
        /// The identifier of a client certificate for an API stage.
        /// </summary>
        [Output("clientCertificateId")]
        public Output<string?> ClientCertificateId { get; private set; } = null!;

        /// <summary>
        /// The identifier of the Deployment that the stage points to.
        /// </summary>
        [Output("deploymentId")]
        public Output<string?> DeploymentId { get; private set; } = null!;

        /// <summary>
        /// The stage's description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The version of the associated API documentation.
        /// </summary>
        [Output("documentationVersion")]
        public Output<string?> DocumentationVersion { get; private set; } = null!;

        /// <summary>
        /// A map that defines the method settings for a Stage resource. Keys (designated as ``/{method_setting_key`` below) are method paths defined as ``{resource_path}/{http_method}`` for an individual method override, or ``/\*/\*`` for overriding all methods in the stage.
        /// </summary>
        [Output("methodSettings")]
        public Output<ImmutableArray<Outputs.StageMethodSetting>> MethodSettings { get; private set; } = null!;

        /// <summary>
        /// The string identifier of the associated RestApi.
        /// </summary>
        [Output("restApiId")]
        public Output<string> RestApiId { get; private set; } = null!;

        /// <summary>
        /// The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway. Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.
        /// </summary>
        [Output("stageName")]
        public Output<string?> StageName { get; private set; } = null!;

        /// <summary>
        /// The collection of tags. Each tag element is associated with a given resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.StageTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies whether active tracing with X-ray is enabled for the Stage.
        /// </summary>
        [Output("tracingEnabled")]
        public Output<bool?> TracingEnabled { get; private set; } = null!;

        /// <summary>
        /// A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value. Variable names are limited to alphanumeric characters. Values must match the following regular expression: ``[A-Za-z0-9-._~:/?#&amp;=,]+``.
        /// </summary>
        [Output("variables")]
        public Output<ImmutableDictionary<string, string>?> Variables { get; private set; } = null!;


        /// <summary>
        /// Create a Stage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Stage(string name, StageArgs args, CustomResourceOptions? options = null)
            : base("aws-native:apigateway:Stage", name, args ?? new StageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Stage(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:apigateway:Stage", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "restApiId",
                    "stageName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Stage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Stage Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Stage(name, id, options);
        }
    }

    public sealed class StageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access log settings, including the access log format and access log destination ARN.
        /// </summary>
        [Input("accessLogSetting")]
        public Input<Inputs.StageAccessLogSettingArgs>? AccessLogSetting { get; set; }

        /// <summary>
        /// Specifies whether a cache cluster is enabled for the stage.
        /// </summary>
        [Input("cacheClusterEnabled")]
        public Input<bool>? CacheClusterEnabled { get; set; }

        /// <summary>
        /// The stage's cache capacity in GB. For more information about choosing a cache size, see [Enabling API caching to enhance responsiveness](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html).
        /// </summary>
        [Input("cacheClusterSize")]
        public Input<string>? CacheClusterSize { get; set; }

        /// <summary>
        /// Settings for the canary deployment in this stage.
        /// </summary>
        [Input("canarySetting")]
        public Input<Inputs.StageCanarySettingArgs>? CanarySetting { get; set; }

        /// <summary>
        /// The identifier of a client certificate for an API stage.
        /// </summary>
        [Input("clientCertificateId")]
        public Input<string>? ClientCertificateId { get; set; }

        /// <summary>
        /// The identifier of the Deployment that the stage points to.
        /// </summary>
        [Input("deploymentId")]
        public Input<string>? DeploymentId { get; set; }

        /// <summary>
        /// The stage's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The version of the associated API documentation.
        /// </summary>
        [Input("documentationVersion")]
        public Input<string>? DocumentationVersion { get; set; }

        [Input("methodSettings")]
        private InputList<Inputs.StageMethodSettingArgs>? _methodSettings;

        /// <summary>
        /// A map that defines the method settings for a Stage resource. Keys (designated as ``/{method_setting_key`` below) are method paths defined as ``{resource_path}/{http_method}`` for an individual method override, or ``/\*/\*`` for overriding all methods in the stage.
        /// </summary>
        public InputList<Inputs.StageMethodSettingArgs> MethodSettings
        {
            get => _methodSettings ?? (_methodSettings = new InputList<Inputs.StageMethodSettingArgs>());
            set => _methodSettings = value;
        }

        /// <summary>
        /// The string identifier of the associated RestApi.
        /// </summary>
        [Input("restApiId", required: true)]
        public Input<string> RestApiId { get; set; } = null!;

        /// <summary>
        /// The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway. Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.
        /// </summary>
        [Input("stageName")]
        public Input<string>? StageName { get; set; }

        [Input("tags")]
        private InputList<Inputs.StageTagArgs>? _tags;

        /// <summary>
        /// The collection of tags. Each tag element is associated with a given resource.
        /// </summary>
        public InputList<Inputs.StageTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.StageTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies whether active tracing with X-ray is enabled for the Stage.
        /// </summary>
        [Input("tracingEnabled")]
        public Input<bool>? TracingEnabled { get; set; }

        [Input("variables")]
        private InputMap<string>? _variables;

        /// <summary>
        /// A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value. Variable names are limited to alphanumeric characters. Values must match the following regular expression: ``[A-Za-z0-9-._~:/?#&amp;=,]+``.
        /// </summary>
        public InputMap<string> Variables
        {
            get => _variables ?? (_variables = new InputMap<string>());
            set => _variables = value;
        }

        public StageArgs()
        {
        }
        public static new StageArgs Empty => new StageArgs();
    }
}
