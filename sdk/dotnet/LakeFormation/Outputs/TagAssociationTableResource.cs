// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LakeFormation.Outputs
{

    [OutputType]
    public sealed class TagAssociationTableResource
    {
        public readonly string CatalogId;
        public readonly string DatabaseName;
        public readonly string? Name;
        public readonly Outputs.TagAssociationTableWildcard? TableWildcard;

        [OutputConstructor]
        private TagAssociationTableResource(
            string catalogId,

            string databaseName,

            string? name,

            Outputs.TagAssociationTableWildcard? tableWildcard)
        {
            CatalogId = catalogId;
            DatabaseName = databaseName;
            Name = name;
            TableWildcard = tableWildcard;
        }
    }
}
