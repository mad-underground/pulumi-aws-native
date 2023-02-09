// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KendraRanking
{
    public static class GetExecutionPlan
    {
        /// <summary>
        /// A KendraRanking Rescore execution plan
        /// </summary>
        public static Task<GetExecutionPlanResult> InvokeAsync(GetExecutionPlanArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExecutionPlanResult>("aws-native:kendraranking:getExecutionPlan", args ?? new GetExecutionPlanArgs(), options.WithDefaults());

        /// <summary>
        /// A KendraRanking Rescore execution plan
        /// </summary>
        public static Output<GetExecutionPlanResult> Invoke(GetExecutionPlanInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetExecutionPlanResult>("aws-native:kendraranking:getExecutionPlan", args ?? new GetExecutionPlanInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExecutionPlanArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetExecutionPlanArgs()
        {
        }
        public static new GetExecutionPlanArgs Empty => new GetExecutionPlanArgs();
    }

    public sealed class GetExecutionPlanInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetExecutionPlanInvokeArgs()
        {
        }
        public static new GetExecutionPlanInvokeArgs Empty => new GetExecutionPlanInvokeArgs();
    }


    [OutputType]
    public sealed class GetExecutionPlanResult
    {
        public readonly string? Arn;
        /// <summary>
        /// Capacity units
        /// </summary>
        public readonly Outputs.ExecutionPlanCapacityUnitsConfiguration? CapacityUnits;
        /// <summary>
        /// A description for the execution plan
        /// </summary>
        public readonly string? Description;
        public readonly string? Id;
        public readonly string? Name;
        /// <summary>
        /// Tags for labeling the execution plan
        /// </summary>
        public readonly ImmutableArray<Outputs.ExecutionPlanTag> Tags;

        [OutputConstructor]
        private GetExecutionPlanResult(
            string? arn,

            Outputs.ExecutionPlanCapacityUnitsConfiguration? capacityUnits,

            string? description,

            string? id,

            string? name,

            ImmutableArray<Outputs.ExecutionPlanTag> tags)
        {
            Arn = arn;
            CapacityUnits = capacityUnits;
            Description = description;
            Id = id;
            Name = name;
            Tags = tags;
        }
    }
}