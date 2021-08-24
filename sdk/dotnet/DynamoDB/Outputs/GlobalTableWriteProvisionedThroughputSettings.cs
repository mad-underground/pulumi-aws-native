// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDB.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-writeprovisionedthroughputsettings.html
    /// </summary>
    [OutputType]
    public sealed class GlobalTableWriteProvisionedThroughputSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-writeprovisionedthroughputsettings.html#cfn-dynamodb-globaltable-writeprovisionedthroughputsettings-writecapacityautoscalingsettings
        /// </summary>
        public readonly Outputs.GlobalTableCapacityAutoScalingSettings? WriteCapacityAutoScalingSettings;

        [OutputConstructor]
        private GlobalTableWriteProvisionedThroughputSettings(Outputs.GlobalTableCapacityAutoScalingSettings? writeCapacityAutoScalingSettings)
        {
            WriteCapacityAutoScalingSettings = writeCapacityAutoScalingSettings;
        }
    }
}