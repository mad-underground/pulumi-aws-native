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
    public sealed class CrawlerTargets
    {
        public readonly ImmutableArray<Outputs.CrawlerCatalogTarget> CatalogTargets;
        public readonly ImmutableArray<Outputs.CrawlerDeltaTarget> DeltaTargets;
        public readonly ImmutableArray<Outputs.CrawlerDynamoDbTarget> DynamoDbTargets;
        public readonly ImmutableArray<Outputs.CrawlerIcebergTarget> IcebergTargets;
        public readonly ImmutableArray<Outputs.CrawlerJdbcTarget> JdbcTargets;
        public readonly ImmutableArray<Outputs.CrawlerMongoDbTarget> MongoDbTargets;
        public readonly ImmutableArray<Outputs.CrawlerS3Target> S3Targets;

        [OutputConstructor]
        private CrawlerTargets(
            ImmutableArray<Outputs.CrawlerCatalogTarget> catalogTargets,

            ImmutableArray<Outputs.CrawlerDeltaTarget> deltaTargets,

            ImmutableArray<Outputs.CrawlerDynamoDbTarget> dynamoDbTargets,

            ImmutableArray<Outputs.CrawlerIcebergTarget> icebergTargets,

            ImmutableArray<Outputs.CrawlerJdbcTarget> jdbcTargets,

            ImmutableArray<Outputs.CrawlerMongoDbTarget> mongoDbTargets,

            ImmutableArray<Outputs.CrawlerS3Target> s3Targets)
        {
            CatalogTargets = catalogTargets;
            DeltaTargets = deltaTargets;
            DynamoDbTargets = dynamoDbTargets;
            IcebergTargets = icebergTargets;
            JdbcTargets = jdbcTargets;
            MongoDbTargets = mongoDbTargets;
            S3Targets = s3Targets;
        }
    }
}
