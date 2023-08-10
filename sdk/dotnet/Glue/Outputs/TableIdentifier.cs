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
    public sealed class TableIdentifier
    {
        public readonly string? CatalogId;
        public readonly string? DatabaseName;
        public readonly string? Name;
        public readonly string? Region;

        [OutputConstructor]
        private TableIdentifier(
            string? catalogId,

            string? databaseName,

            string? name,

            string? region)
        {
            CatalogId = catalogId;
            DatabaseName = databaseName;
            Name = name;
            Region = region;
        }
    }
}
