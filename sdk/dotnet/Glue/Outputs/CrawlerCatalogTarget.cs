// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    [OutputType]
    public sealed class CrawlerCatalogTarget
    {
        public readonly string? ConnectionName;
        public readonly string? DatabaseName;
        public readonly string? DlqEventQueueArn;
        public readonly string? EventQueueArn;
        public readonly ImmutableArray<string> Tables;

        [OutputConstructor]
        private CrawlerCatalogTarget(
            string? connectionName,

            string? databaseName,

            string? dlqEventQueueArn,

            string? eventQueueArn,

            ImmutableArray<string> tables)
        {
            ConnectionName = connectionName;
            DatabaseName = databaseName;
            DlqEventQueueArn = dlqEventQueueArn;
            EventQueueArn = eventQueueArn;
            Tables = tables;
        }
    }
}
