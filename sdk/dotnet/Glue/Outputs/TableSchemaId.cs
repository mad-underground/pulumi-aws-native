// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Glue.Outputs
{

    [OutputType]
    public sealed class TableSchemaId
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-schemaid.html#cfn-glue-table-schemaid-registryname
        /// </summary>
        public readonly string? RegistryName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-schemaid.html#cfn-glue-table-schemaid-schemaarn
        /// </summary>
        public readonly string? SchemaArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-schemaid.html#cfn-glue-table-schemaid-schemaname
        /// </summary>
        public readonly string? SchemaName;

        [OutputConstructor]
        private TableSchemaId(
            string? RegistryName,

            string? SchemaArn,

            string? SchemaName)
        {
            this.RegistryName = RegistryName;
            this.SchemaArn = SchemaArn;
            this.SchemaName = SchemaName;
        }
    }
}
