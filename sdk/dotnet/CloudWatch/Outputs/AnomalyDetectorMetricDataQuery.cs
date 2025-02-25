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
    public sealed class AnomalyDetectorMetricDataQuery
    {
        public readonly string? AccountId;
        public readonly string? Expression;
        public readonly string Id;
        public readonly string? Label;
        public readonly Outputs.AnomalyDetectorMetricStat? MetricStat;
        public readonly int? Period;
        public readonly bool? ReturnData;

        [OutputConstructor]
        private AnomalyDetectorMetricDataQuery(
            string? accountId,

            string? expression,

            string id,

            string? label,

            Outputs.AnomalyDetectorMetricStat? metricStat,

            int? period,

            bool? returnData)
        {
            AccountId = accountId;
            Expression = expression;
            Id = id;
            Label = label;
            MetricStat = metricStat;
            Period = period;
            ReturnData = returnData;
        }
    }
}
