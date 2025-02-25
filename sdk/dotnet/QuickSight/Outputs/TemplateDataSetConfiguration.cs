// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateDataSetConfiguration
    {
        public readonly ImmutableArray<Outputs.TemplateColumnGroupSchema> ColumnGroupSchemaList;
        public readonly Outputs.TemplateDataSetSchema? DataSetSchema;
        public readonly string? Placeholder;

        [OutputConstructor]
        private TemplateDataSetConfiguration(
            ImmutableArray<Outputs.TemplateColumnGroupSchema> columnGroupSchemaList,

            Outputs.TemplateDataSetSchema? dataSetSchema,

            string? placeholder)
        {
            ColumnGroupSchemaList = columnGroupSchemaList;
            DataSetSchema = dataSetSchema;
            Placeholder = placeholder;
        }
    }
}
