// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.CloudWatch.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-configuration.html
    /// </summary>
    public sealed class AnomalyDetectorConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("ExcludedTimeRanges")]
        private InputList<Inputs.AnomalyDetectorRangeArgs>? _ExcludedTimeRanges;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-configuration.html#cfn-cloudwatch-anomalydetector-configuration-excludedtimeranges
        /// </summary>
        public InputList<Inputs.AnomalyDetectorRangeArgs> ExcludedTimeRanges
        {
            get => _ExcludedTimeRanges ?? (_ExcludedTimeRanges = new InputList<Inputs.AnomalyDetectorRangeArgs>());
            set => _ExcludedTimeRanges = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-configuration.html#cfn-cloudwatch-anomalydetector-configuration-metrictimezone
        /// </summary>
        [Input("MetricTimeZone")]
        public Input<string>? MetricTimeZone { get; set; }

        public AnomalyDetectorConfigurationArgs()
        {
        }
    }
}