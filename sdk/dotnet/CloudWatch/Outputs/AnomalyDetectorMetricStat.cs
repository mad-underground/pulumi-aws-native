// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudWatch.Outputs
{

    [OutputType]
    public sealed class AnomalyDetectorMetricStat
    {
        public readonly Outputs.AnomalyDetectorMetric Metric;
        public readonly int Period;
        public readonly string Stat;
        public readonly string? Unit;

        [OutputConstructor]
        private AnomalyDetectorMetricStat(
            Outputs.AnomalyDetectorMetric metric,

            int period,

            string stat,

            string? unit)
        {
            Metric = metric;
            Period = period;
            Stat = stat;
            Unit = unit;
        }
    }
}
