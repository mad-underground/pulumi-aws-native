// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync
{
    public static class GetFunctionConfiguration
    {
        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Task<GetFunctionConfigurationResult> InvokeAsync(GetFunctionConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFunctionConfigurationResult>("aws-native:appsync:getFunctionConfiguration", args ?? new GetFunctionConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Output<GetFunctionConfigurationResult> Invoke(GetFunctionConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFunctionConfigurationResult>("aws-native:appsync:getFunctionConfiguration", args ?? new GetFunctionConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN for the function generated by the service
        /// </summary>
        [Input("functionArn", required: true)]
        public string FunctionArn { get; set; } = null!;

        public GetFunctionConfigurationArgs()
        {
        }
        public static new GetFunctionConfigurationArgs Empty => new GetFunctionConfigurationArgs();
    }

    public sealed class GetFunctionConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN for the function generated by the service
        /// </summary>
        [Input("functionArn", required: true)]
        public Input<string> FunctionArn { get; set; } = null!;

        public GetFunctionConfigurationInvokeArgs()
        {
        }
        public static new GetFunctionConfigurationInvokeArgs Empty => new GetFunctionConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetFunctionConfigurationResult
    {
        /// <summary>
        /// The resolver code that contains the request and response functions. When code is used, the runtime is required. The runtime value must be APPSYNC_JS.
        /// </summary>
        public readonly string? Code;
        /// <summary>
        /// The name of data source this function will attach.
        /// </summary>
        public readonly string? DataSourceName;
        /// <summary>
        /// The function description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The ARN for the function generated by the service
        /// </summary>
        public readonly string? FunctionArn;
        /// <summary>
        /// The unique identifier for the function generated by the service
        /// </summary>
        public readonly string? FunctionId;
        /// <summary>
        /// The version of the request mapping template. Currently, only the 2018-05-29 version of the template is supported.
        /// </summary>
        public readonly string? FunctionVersion;
        /// <summary>
        /// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a BatchInvoke operation.
        /// </summary>
        public readonly int? MaxBatchSize;
        /// <summary>
        /// The name of the function.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The Function request mapping template. Functions support only the 2018-05-29 version of the request mapping template.
        /// </summary>
        public readonly string? RequestMappingTemplate;
        /// <summary>
        /// The Function response mapping template.
        /// </summary>
        public readonly string? ResponseMappingTemplate;
        /// <summary>
        /// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.
        /// </summary>
        public readonly Outputs.FunctionConfigurationAppSyncRuntime? Runtime;
        /// <summary>
        /// Describes a Sync configuration for a resolver. Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is invoked.
        /// </summary>
        public readonly Outputs.FunctionConfigurationSyncConfig? SyncConfig;

        [OutputConstructor]
        private GetFunctionConfigurationResult(
            string? code,

            string? dataSourceName,

            string? description,

            string? functionArn,

            string? functionId,

            string? functionVersion,

            int? maxBatchSize,

            string? name,

            string? requestMappingTemplate,

            string? responseMappingTemplate,

            Outputs.FunctionConfigurationAppSyncRuntime? runtime,

            Outputs.FunctionConfigurationSyncConfig? syncConfig)
        {
            Code = code;
            DataSourceName = dataSourceName;
            Description = description;
            FunctionArn = functionArn;
            FunctionId = functionId;
            FunctionVersion = functionVersion;
            MaxBatchSize = maxBatchSize;
            Name = name;
            RequestMappingTemplate = requestMappingTemplate;
            ResponseMappingTemplate = responseMappingTemplate;
            Runtime = runtime;
            SyncConfig = syncConfig;
        }
    }
}
