// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ivs
{
    public static class GetStage
    {
        /// <summary>
        /// Resource Definition for type AWS::IVS::Stage.
        /// </summary>
        public static Task<GetStageResult> InvokeAsync(GetStageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStageResult>("aws-native:ivs:getStage", args ?? new GetStageArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Definition for type AWS::IVS::Stage.
        /// </summary>
        public static Output<GetStageResult> Invoke(GetStageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStageResult>("aws-native:ivs:getStage", args ?? new GetStageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Stage ARN is automatically generated on creation and assigned as the unique identifier.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetStageArgs()
        {
        }
        public static new GetStageArgs Empty => new GetStageArgs();
    }

    public sealed class GetStageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Stage ARN is automatically generated on creation and assigned as the unique identifier.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetStageInvokeArgs()
        {
        }
        public static new GetStageInvokeArgs Empty => new GetStageInvokeArgs();
    }


    [OutputType]
    public sealed class GetStageResult
    {
        /// <summary>
        /// ID of the active session within the stage.
        /// </summary>
        public readonly string? ActiveSessionId;
        /// <summary>
        /// Stage ARN is automatically generated on creation and assigned as the unique identifier.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Stage name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.StageTag> Tags;

        [OutputConstructor]
        private GetStageResult(
            string? activeSessionId,

            string? arn,

            string? name,

            ImmutableArray<Outputs.StageTag> tags)
        {
            ActiveSessionId = activeSessionId;
            Arn = arn;
            Name = name;
            Tags = tags;
        }
    }
}
