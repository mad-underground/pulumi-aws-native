// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone
{
    public static class GetProject
    {
        /// <summary>
        /// Amazon DataZone projects are business use case–based groupings of people, assets (data), and tools used to simplify access to the AWS analytics.
        /// </summary>
        public static Task<GetProjectResult> InvokeAsync(GetProjectArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("aws-native:datazone:getProject", args ?? new GetProjectArgs(), options.WithDefaults());

        /// <summary>
        /// Amazon DataZone projects are business use case–based groupings of people, assets (data), and tools used to simplify access to the AWS analytics.
        /// </summary>
        public static Output<GetProjectResult> Invoke(GetProjectInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectResult>("aws-native:datazone:getProject", args ?? new GetProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the Amazon DataZone domain in which the project was created.
        /// </summary>
        [Input("domainId", required: true)]
        public string DomainId { get; set; } = null!;

        /// <summary>
        /// The ID of the Amazon DataZone project.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetProjectArgs()
        {
        }
        public static new GetProjectArgs Empty => new GetProjectArgs();
    }

    public sealed class GetProjectInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier of the Amazon DataZone domain in which the project was created.
        /// </summary>
        [Input("domainId", required: true)]
        public Input<string> DomainId { get; set; } = null!;

        /// <summary>
        /// The ID of the Amazon DataZone project.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetProjectInvokeArgs()
        {
        }
        public static new GetProjectInvokeArgs Empty => new GetProjectInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectResult
    {
        /// <summary>
        /// The timestamp of when the project was created.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// The Amazon DataZone user who created the project.
        /// </summary>
        public readonly string? CreatedBy;
        /// <summary>
        /// The description of the Amazon DataZone project.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The identifier of the Amazon DataZone domain in which the project was created.
        /// </summary>
        public readonly string? DomainId;
        /// <summary>
        /// The glossary terms that can be used in this Amazon DataZone project.
        /// </summary>
        public readonly ImmutableArray<string> GlossaryTerms;
        /// <summary>
        /// The ID of the Amazon DataZone project.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The timestamp of when the project was last updated.
        /// </summary>
        public readonly string? LastUpdatedAt;
        /// <summary>
        /// The name of the Amazon DataZone project.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private GetProjectResult(
            string? createdAt,

            string? createdBy,

            string? description,

            string? domainId,

            ImmutableArray<string> glossaryTerms,

            string? id,

            string? lastUpdatedAt,

            string? name)
        {
            CreatedAt = createdAt;
            CreatedBy = createdBy;
            Description = description;
            DomainId = domainId;
            GlossaryTerms = glossaryTerms;
            Id = id;
            LastUpdatedAt = lastUpdatedAt;
            Name = name;
        }
    }
}
