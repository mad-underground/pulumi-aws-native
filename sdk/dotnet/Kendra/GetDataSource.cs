// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra
{
    public static class GetDataSource
    {
        /// <summary>
        /// Kendra DataSource
        /// </summary>
        public static Task<GetDataSourceResult> InvokeAsync(GetDataSourceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataSourceResult>("aws-native:kendra:getDataSource", args ?? new GetDataSourceArgs(), options.WithDefaults());

        /// <summary>
        /// Kendra DataSource
        /// </summary>
        public static Output<GetDataSourceResult> Invoke(GetDataSourceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDataSourceResult>("aws-native:kendra:getDataSource", args ?? new GetDataSourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataSourceArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("indexId", required: true)]
        public string IndexId { get; set; } = null!;

        public GetDataSourceArgs()
        {
        }
    }

    public sealed class GetDataSourceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("indexId", required: true)]
        public Input<string> IndexId { get; set; } = null!;

        public GetDataSourceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDataSourceResult
    {
        public readonly string? Arn;
        public readonly Outputs.DataSourceConfiguration? DataSourceConfiguration;
        public readonly string? Description;
        public readonly string? Id;
        public readonly string? IndexId;
        public readonly string? Name;
        public readonly string? RoleArn;
        public readonly string? Schedule;
        /// <summary>
        /// Tags for labeling the data source
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceTag> Tags;

        [OutputConstructor]
        private GetDataSourceResult(
            string? arn,

            Outputs.DataSourceConfiguration? dataSourceConfiguration,

            string? description,

            string? id,

            string? indexId,

            string? name,

            string? roleArn,

            string? schedule,

            ImmutableArray<Outputs.DataSourceTag> tags)
        {
            Arn = arn;
            DataSourceConfiguration = dataSourceConfiguration;
            Description = description;
            Id = id;
            IndexId = indexId;
            Name = name;
            RoleArn = roleArn;
            Schedule = schedule;
            Tags = tags;
        }
    }
}
