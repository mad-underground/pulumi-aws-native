// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Inputs
{

    public sealed class DatabaseDatabaseIdentifierArgs : Pulumi.ResourceArgs
    {
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        public DatabaseDatabaseIdentifierArgs()
        {
        }
    }
}