// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudWatch.Outputs
{

    [OutputType]
    public sealed class MetricStreamMetricStreamFilter
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-metricstream-metricstreamfilter.html#cfn-cloudwatch-metricstream-metricstreamfilter-namespace
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private MetricStreamMetricStreamFilter(string @namespace)
        {
            Namespace = @namespace;
        }
    }
}
