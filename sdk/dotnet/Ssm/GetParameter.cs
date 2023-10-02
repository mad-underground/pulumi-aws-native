// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ssm
{
    public static class GetParameter
    {
        /// <summary>
        /// Resource Type definition for AWS::SSM::Parameter
        /// </summary>
        public static Task<GetParameterResult> InvokeAsync(GetParameterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetParameterResult>("aws-native:ssm:getParameter", args ?? new GetParameterArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SSM::Parameter
        /// </summary>
        public static Output<GetParameterResult> Invoke(GetParameterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetParameterResult>("aws-native:ssm:getParameter", args ?? new GetParameterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetParameterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the parameter.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetParameterArgs()
        {
        }
        public static new GetParameterArgs Empty => new GetParameterArgs();
    }

    public sealed class GetParameterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the parameter.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetParameterInvokeArgs()
        {
        }
        public static new GetParameterInvokeArgs Empty => new GetParameterInvokeArgs();
    }


    [OutputType]
    public sealed class GetParameterResult
    {
        /// <summary>
        /// The corresponding DataType of the parameter.
        /// </summary>
        public readonly Pulumi.AwsNative.Ssm.ParameterDataType? DataType;
        /// <summary>
        /// The type of the parameter.
        /// </summary>
        public readonly Pulumi.AwsNative.Ssm.ParameterType? Type;
        /// <summary>
        /// The value associated with the parameter.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private GetParameterResult(
            Pulumi.AwsNative.Ssm.ParameterDataType? dataType,

            Pulumi.AwsNative.Ssm.ParameterType? type,

            string? value)
        {
            DataType = dataType;
            Type = type;
            Value = value;
        }
    }
}
