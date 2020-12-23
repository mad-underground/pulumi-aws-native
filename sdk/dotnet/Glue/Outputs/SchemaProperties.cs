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
    public sealed class SchemaProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-checkpointversion
        /// </summary>
        public readonly Outputs.SchemaSchemaVersion? CheckpointVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-compatibility
        /// </summary>
        public readonly string Compatibility;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-dataformat
        /// </summary>
        public readonly string DataFormat;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-registry
        /// </summary>
        public readonly Outputs.SchemaRegistry? Registry;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-schemadefinition
        /// </summary>
        public readonly string SchemaDefinition;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schema.html#cfn-glue-schema-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;

        [OutputConstructor]
        private SchemaProperties(
            Outputs.SchemaSchemaVersion? CheckpointVersion,

            string Compatibility,

            string DataFormat,

            string? Description,

            string Name,

            Outputs.SchemaRegistry? Registry,

            string SchemaDefinition,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags)
        {
            this.CheckpointVersion = CheckpointVersion;
            this.Compatibility = Compatibility;
            this.DataFormat = DataFormat;
            this.Description = Description;
            this.Name = Name;
            this.Registry = Registry;
            this.SchemaDefinition = SchemaDefinition;
            this.Tags = Tags;
        }
    }
}
