// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecretsManager
{
    public static class GetRotationSchedule
    {
        /// <summary>
        /// Resource Type definition for AWS::SecretsManager::RotationSchedule
        /// </summary>
        public static Task<GetRotationScheduleResult> InvokeAsync(GetRotationScheduleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRotationScheduleResult>("aws-native:secretsmanager:getRotationSchedule", args ?? new GetRotationScheduleArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SecretsManager::RotationSchedule
        /// </summary>
        public static Output<GetRotationScheduleResult> Invoke(GetRotationScheduleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRotationScheduleResult>("aws-native:secretsmanager:getRotationSchedule", args ?? new GetRotationScheduleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRotationScheduleArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetRotationScheduleArgs()
        {
        }
    }

    public sealed class GetRotationScheduleInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetRotationScheduleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRotationScheduleResult
    {
        public readonly Outputs.RotationScheduleHostedRotationLambda? HostedRotationLambda;
        public readonly string? Id;
        public readonly bool? RotateImmediatelyOnUpdate;
        public readonly string? RotationLambdaARN;
        public readonly Outputs.RotationScheduleRotationRules? RotationRules;

        [OutputConstructor]
        private GetRotationScheduleResult(
            Outputs.RotationScheduleHostedRotationLambda? hostedRotationLambda,

            string? id,

            bool? rotateImmediatelyOnUpdate,

            string? rotationLambdaARN,

            Outputs.RotationScheduleRotationRules? rotationRules)
        {
            HostedRotationLambda = hostedRotationLambda;
            Id = id;
            RotateImmediatelyOnUpdate = rotateImmediatelyOnUpdate;
            RotationLambdaARN = rotationLambdaARN;
            RotationRules = rotationRules;
        }
    }
}