// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EntityResolution
{
    public static class GetIdMappingWorkflow
    {
        /// <summary>
        /// IdMappingWorkflow defined in AWS Entity Resolution service
        /// </summary>
        public static Task<GetIdMappingWorkflowResult> InvokeAsync(GetIdMappingWorkflowArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdMappingWorkflowResult>("aws-native:entityresolution:getIdMappingWorkflow", args ?? new GetIdMappingWorkflowArgs(), options.WithDefaults());

        /// <summary>
        /// IdMappingWorkflow defined in AWS Entity Resolution service
        /// </summary>
        public static Output<GetIdMappingWorkflowResult> Invoke(GetIdMappingWorkflowInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdMappingWorkflowResult>("aws-native:entityresolution:getIdMappingWorkflow", args ?? new GetIdMappingWorkflowInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdMappingWorkflowArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the IdMappingWorkflow
        /// </summary>
        [Input("workflowName", required: true)]
        public string WorkflowName { get; set; } = null!;

        public GetIdMappingWorkflowArgs()
        {
        }
        public static new GetIdMappingWorkflowArgs Empty => new GetIdMappingWorkflowArgs();
    }

    public sealed class GetIdMappingWorkflowInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the IdMappingWorkflow
        /// </summary>
        [Input("workflowName", required: true)]
        public Input<string> WorkflowName { get; set; } = null!;

        public GetIdMappingWorkflowInvokeArgs()
        {
        }
        public static new GetIdMappingWorkflowInvokeArgs Empty => new GetIdMappingWorkflowInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdMappingWorkflowResult
    {
        public readonly string? CreatedAt;
        /// <summary>
        /// The description of the IdMappingWorkflow
        /// </summary>
        public readonly string? Description;
        public readonly Outputs.IdMappingWorkflowIdMappingTechniques? IdMappingTechniques;
        public readonly ImmutableArray<Outputs.IdMappingWorkflowInputSource> InputSourceConfig;
        public readonly ImmutableArray<Outputs.IdMappingWorkflowOutputSource> OutputSourceConfig;
        public readonly string? RoleArn;
        public readonly ImmutableArray<Outputs.IdMappingWorkflowTag> Tags;
        public readonly string? UpdatedAt;
        public readonly string? WorkflowArn;

        [OutputConstructor]
        private GetIdMappingWorkflowResult(
            string? createdAt,

            string? description,

            Outputs.IdMappingWorkflowIdMappingTechniques? idMappingTechniques,

            ImmutableArray<Outputs.IdMappingWorkflowInputSource> inputSourceConfig,

            ImmutableArray<Outputs.IdMappingWorkflowOutputSource> outputSourceConfig,

            string? roleArn,

            ImmutableArray<Outputs.IdMappingWorkflowTag> tags,

            string? updatedAt,

            string? workflowArn)
        {
            CreatedAt = createdAt;
            Description = description;
            IdMappingTechniques = idMappingTechniques;
            InputSourceConfig = inputSourceConfig;
            OutputSourceConfig = outputSourceConfig;
            RoleArn = roleArn;
            Tags = tags;
            UpdatedAt = updatedAt;
            WorkflowArn = workflowArn;
        }
    }
}