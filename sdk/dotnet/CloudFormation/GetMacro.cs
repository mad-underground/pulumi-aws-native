// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    public static class GetMacro
    {
        /// <summary>
        /// Resource Type definition for AWS::CloudFormation::Macro
        /// </summary>
        public static Task<GetMacroResult> InvokeAsync(GetMacroArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMacroResult>("aws-native:cloudformation:getMacro", args ?? new GetMacroArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::CloudFormation::Macro
        /// </summary>
        public static Output<GetMacroResult> Invoke(GetMacroInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMacroResult>("aws-native:cloudformation:getMacro", args ?? new GetMacroInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMacroArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetMacroArgs()
        {
        }
        public static new GetMacroArgs Empty => new GetMacroArgs();
    }

    public sealed class GetMacroInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetMacroInvokeArgs()
        {
        }
        public static new GetMacroInvokeArgs Empty => new GetMacroInvokeArgs();
    }


    [OutputType]
    public sealed class GetMacroResult
    {
        public readonly string? Description;
        public readonly string? FunctionName;
        public readonly string? Id;
        public readonly string? LogGroupName;
        public readonly string? LogRoleARN;

        [OutputConstructor]
        private GetMacroResult(
            string? description,

            string? functionName,

            string? id,

            string? logGroupName,

            string? logRoleARN)
        {
            Description = description;
            FunctionName = functionName;
            Id = id;
            LogGroupName = logGroupName;
            LogRoleARN = logRoleARN;
        }
    }
}
