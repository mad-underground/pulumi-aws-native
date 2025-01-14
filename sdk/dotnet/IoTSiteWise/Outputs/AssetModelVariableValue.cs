// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Outputs
{

    [OutputType]
    public sealed class AssetModelVariableValue
    {
        /// <summary>
        /// The External ID of the hierarchy that is trying to be referenced
        /// </summary>
        public readonly string? HierarchyExternalId;
        /// <summary>
        /// The ID of the hierarchy that is trying to be referenced
        /// </summary>
        public readonly string? HierarchyId;
        public readonly string? HierarchyLogicalId;
        /// <summary>
        /// The External ID of the property that is trying to be referenced
        /// </summary>
        public readonly string? PropertyExternalId;
        /// <summary>
        /// The ID of the property that is trying to be referenced
        /// </summary>
        public readonly string? PropertyId;
        public readonly string? PropertyLogicalId;
        /// <summary>
        /// The path of the property that is trying to be referenced
        /// </summary>
        public readonly ImmutableArray<Outputs.AssetModelPropertyPathDefinition> PropertyPath;

        [OutputConstructor]
        private AssetModelVariableValue(
            string? hierarchyExternalId,

            string? hierarchyId,

            string? hierarchyLogicalId,

            string? propertyExternalId,

            string? propertyId,

            string? propertyLogicalId,

            ImmutableArray<Outputs.AssetModelPropertyPathDefinition> propertyPath)
        {
            HierarchyExternalId = hierarchyExternalId;
            HierarchyId = hierarchyId;
            HierarchyLogicalId = hierarchyLogicalId;
            PropertyExternalId = propertyExternalId;
            PropertyId = propertyId;
            PropertyLogicalId = propertyLogicalId;
            PropertyPath = propertyPath;
        }
    }
}
