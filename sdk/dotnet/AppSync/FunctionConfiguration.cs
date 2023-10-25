// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync
{
    /// <summary>
    /// An example resource schema demonstrating some basic constructs and validation rules.
    /// </summary>
    [AwsNativeResourceType("aws-native:appsync:FunctionConfiguration")]
    public partial class FunctionConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS AppSync GraphQL API that you want to attach using this function.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// The resolver code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
        /// </summary>
        [Output("code")]
        public Output<string?> Code { get; private set; } = null!;

        /// <summary>
        /// The Amazon S3 endpoint (where the code is located??).
        /// </summary>
        [Output("codeS3Location")]
        public Output<string?> CodeS3Location { get; private set; } = null!;

        /// <summary>
        /// The name of data source this function will attach.
        /// </summary>
        [Output("dataSourceName")]
        public Output<string> DataSourceName { get; private set; } = null!;

        /// <summary>
        /// The function description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ARN for the function generated by the service
        /// </summary>
        [Output("functionArn")]
        public Output<string> FunctionArn { get; private set; } = null!;

        /// <summary>
        /// The unique identifier for the function generated by the service
        /// </summary>
        [Output("functionId")]
        public Output<string> FunctionId { get; private set; } = null!;

        /// <summary>
        /// The version of the request mapping template. Currently, only the 2018-05-29 version of the template is supported.
        /// </summary>
        [Output("functionVersion")]
        public Output<string?> FunctionVersion { get; private set; } = null!;

        /// <summary>
        /// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a BatchInvoke operation.
        /// </summary>
        [Output("maxBatchSize")]
        public Output<int?> MaxBatchSize { get; private set; } = null!;

        /// <summary>
        /// The name of the function.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        /// </summary>
        [Output("requestMappingTemplate")]
        public Output<string?> RequestMappingTemplate { get; private set; } = null!;

        /// <summary>
        /// Describes a Sync configuration for a resolver. Contains information on which Conflict Detection, as well as Resolution strategy, should be performed when the resolver is invoked.
        /// </summary>
        [Output("requestMappingTemplateS3Location")]
        public Output<string?> RequestMappingTemplateS3Location { get; private set; } = null!;

        /// <summary>
        /// The Function response mapping template.
        /// </summary>
        [Output("responseMappingTemplate")]
        public Output<string?> ResponseMappingTemplate { get; private set; } = null!;

        /// <summary>
        /// The location of a response mapping template in an Amazon S3 bucket. Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
        /// </summary>
        [Output("responseMappingTemplateS3Location")]
        public Output<string?> ResponseMappingTemplateS3Location { get; private set; } = null!;

        /// <summary>
        /// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.
        /// </summary>
        [Output("runtime")]
        public Output<Outputs.FunctionConfigurationAppSyncRuntime?> Runtime { get; private set; } = null!;

        /// <summary>
        /// Describes a Sync configuration for a resolver. Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is invoked.
        /// </summary>
        [Output("syncConfig")]
        public Output<Outputs.FunctionConfigurationSyncConfig?> SyncConfig { get; private set; } = null!;


        /// <summary>
        /// Create a FunctionConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FunctionConfiguration(string name, FunctionConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appsync:FunctionConfiguration", name, args ?? new FunctionConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FunctionConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appsync:FunctionConfiguration", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "apiId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FunctionConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FunctionConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FunctionConfiguration(name, id, options);
        }
    }

    public sealed class FunctionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS AppSync GraphQL API that you want to attach using this function.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// The resolver code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
        /// </summary>
        [Input("code")]
        public Input<string>? Code { get; set; }

        /// <summary>
        /// The Amazon S3 endpoint (where the code is located??).
        /// </summary>
        [Input("codeS3Location")]
        public Input<string>? CodeS3Location { get; set; }

        /// <summary>
        /// The name of data source this function will attach.
        /// </summary>
        [Input("dataSourceName", required: true)]
        public Input<string> DataSourceName { get; set; } = null!;

        /// <summary>
        /// The function description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The version of the request mapping template. Currently, only the 2018-05-29 version of the template is supported.
        /// </summary>
        [Input("functionVersion")]
        public Input<string>? FunctionVersion { get; set; }

        /// <summary>
        /// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a BatchInvoke operation.
        /// </summary>
        [Input("maxBatchSize")]
        public Input<int>? MaxBatchSize { get; set; }

        /// <summary>
        /// The name of the function.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        /// </summary>
        [Input("requestMappingTemplate")]
        public Input<string>? RequestMappingTemplate { get; set; }

        /// <summary>
        /// Describes a Sync configuration for a resolver. Contains information on which Conflict Detection, as well as Resolution strategy, should be performed when the resolver is invoked.
        /// </summary>
        [Input("requestMappingTemplateS3Location")]
        public Input<string>? RequestMappingTemplateS3Location { get; set; }

        /// <summary>
        /// The Function response mapping template.
        /// </summary>
        [Input("responseMappingTemplate")]
        public Input<string>? ResponseMappingTemplate { get; set; }

        /// <summary>
        /// The location of a response mapping template in an Amazon S3 bucket. Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.
        /// </summary>
        [Input("responseMappingTemplateS3Location")]
        public Input<string>? ResponseMappingTemplateS3Location { get; set; }

        /// <summary>
        /// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.
        /// </summary>
        [Input("runtime")]
        public Input<Inputs.FunctionConfigurationAppSyncRuntimeArgs>? Runtime { get; set; }

        /// <summary>
        /// Describes a Sync configuration for a resolver. Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is invoked.
        /// </summary>
        [Input("syncConfig")]
        public Input<Inputs.FunctionConfigurationSyncConfigArgs>? SyncConfig { get; set; }

        public FunctionConfigurationArgs()
        {
        }
        public static new FunctionConfigurationArgs Empty => new FunctionConfigurationArgs();
    }
}
