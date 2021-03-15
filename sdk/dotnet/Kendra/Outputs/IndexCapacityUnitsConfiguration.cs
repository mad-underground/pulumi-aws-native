// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    [OutputType]
    public sealed class IndexCapacityUnitsConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-capacityunitsconfiguration.html#cfn-kendra-index-capacityunitsconfiguration-querycapacityunits
        /// </summary>
        public readonly int QueryCapacityUnits;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-capacityunitsconfiguration.html#cfn-kendra-index-capacityunitsconfiguration-storagecapacityunits
        /// </summary>
        public readonly int StorageCapacityUnits;

        [OutputConstructor]
        private IndexCapacityUnitsConfiguration(
            int queryCapacityUnits,

            int storageCapacityUnits)
        {
            QueryCapacityUnits = queryCapacityUnits;
            StorageCapacityUnits = storageCapacityUnits;
        }
    }
}
