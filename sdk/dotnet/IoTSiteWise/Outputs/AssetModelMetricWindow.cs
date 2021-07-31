// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-metricwindow.html
    /// </summary>
    [OutputType]
    public sealed class AssetModelMetricWindow
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-metricwindow.html#cfn-iotsitewise-assetmodel-metricwindow-tumbling
        /// </summary>
        public readonly Outputs.AssetModelTumblingWindow? Tumbling;

        [OutputConstructor]
        private AssetModelMetricWindow(Outputs.AssetModelTumblingWindow? tumbling)
        {
            Tumbling = tumbling;
        }
    }
}
