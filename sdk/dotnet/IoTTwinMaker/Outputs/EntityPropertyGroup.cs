// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTTwinMaker.Outputs
{

    /// <summary>
    /// An object that specifies information about a property group.
    /// </summary>
    [OutputType]
    public sealed class EntityPropertyGroup
    {
        /// <summary>
        /// The type of property group.
        /// </summary>
        public readonly Pulumi.AwsNative.IoTTwinMaker.EntityPropertyGroupGroupType? GroupType;
        /// <summary>
        /// The list of property names in the property group.
        /// </summary>
        public readonly ImmutableArray<string> PropertyNames;

        [OutputConstructor]
        private EntityPropertyGroup(
            Pulumi.AwsNative.IoTTwinMaker.EntityPropertyGroupGroupType? groupType,

            ImmutableArray<string> propertyNames)
        {
            GroupType = groupType;
            PropertyNames = propertyNames;
        }
    }
}
