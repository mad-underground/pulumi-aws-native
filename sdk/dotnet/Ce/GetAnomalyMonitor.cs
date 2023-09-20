// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ce
{
    public static class GetAnomalyMonitor
    {
        /// <summary>
        /// AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. You can use Cost Anomaly Detection by creating monitor.
        /// </summary>
        public static Task<GetAnomalyMonitorResult> InvokeAsync(GetAnomalyMonitorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAnomalyMonitorResult>("aws-native:ce:getAnomalyMonitor", args ?? new GetAnomalyMonitorArgs(), options.WithDefaults());

        /// <summary>
        /// AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. You can use Cost Anomaly Detection by creating monitor.
        /// </summary>
        public static Output<GetAnomalyMonitorResult> Invoke(GetAnomalyMonitorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAnomalyMonitorResult>("aws-native:ce:getAnomalyMonitor", args ?? new GetAnomalyMonitorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAnomalyMonitorArgs : global::Pulumi.InvokeArgs
    {
        [Input("monitorArn", required: true)]
        public string MonitorArn { get; set; } = null!;

        public GetAnomalyMonitorArgs()
        {
        }
        public static new GetAnomalyMonitorArgs Empty => new GetAnomalyMonitorArgs();
    }

    public sealed class GetAnomalyMonitorInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("monitorArn", required: true)]
        public Input<string> MonitorArn { get; set; } = null!;

        public GetAnomalyMonitorInvokeArgs()
        {
        }
        public static new GetAnomalyMonitorInvokeArgs Empty => new GetAnomalyMonitorInvokeArgs();
    }


    [OutputType]
    public sealed class GetAnomalyMonitorResult
    {
        /// <summary>
        /// The date when the monitor was created. 
        /// </summary>
        public readonly string? CreationDate;
        /// <summary>
        /// The value for evaluated dimensions.
        /// </summary>
        public readonly int? DimensionalValueCount;
        /// <summary>
        /// The date when the monitor last evaluated for anomalies.
        /// </summary>
        public readonly string? LastEvaluatedDate;
        /// <summary>
        /// The date when the monitor was last updated.
        /// </summary>
        public readonly string? LastUpdatedDate;
        public readonly string? MonitorArn;
        /// <summary>
        /// The name of the monitor.
        /// </summary>
        public readonly string? MonitorName;

        [OutputConstructor]
        private GetAnomalyMonitorResult(
            string? creationDate,

            int? dimensionalValueCount,

            string? lastEvaluatedDate,

            string? lastUpdatedDate,

            string? monitorArn,

            string? monitorName)
        {
            CreationDate = creationDate;
            DimensionalValueCount = dimensionalValueCount;
            LastEvaluatedDate = lastEvaluatedDate;
            LastUpdatedDate = lastUpdatedDate;
            MonitorArn = monitorArn;
            MonitorName = monitorName;
        }
    }
}