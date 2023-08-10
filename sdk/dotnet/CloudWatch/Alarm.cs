// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudWatch
{
    /// <summary>
    /// Resource Type definition for AWS::CloudWatch::Alarm
    /// </summary>
    [AwsNativeResourceType("aws-native:cloudwatch:Alarm")]
    public partial class Alarm : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
        /// </summary>
        [Output("actionsEnabled")]
        public Output<bool?> ActionsEnabled { get; private set; } = null!;

        /// <summary>
        /// The list of actions to execute when this alarm transitions into an ALARM state from any other state.
        /// </summary>
        [Output("alarmActions")]
        public Output<ImmutableArray<string>> AlarmActions { get; private set; } = null!;

        /// <summary>
        /// The description of the alarm.
        /// </summary>
        [Output("alarmDescription")]
        public Output<string?> AlarmDescription { get; private set; } = null!;

        /// <summary>
        /// The name of the alarm.
        /// </summary>
        [Output("alarmName")]
        public Output<string?> AlarmName { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name is a unique name for each resource.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The arithmetic operation to use when comparing the specified statistic and threshold.
        /// </summary>
        [Output("comparisonOperator")]
        public Output<string> ComparisonOperator { get; private set; } = null!;

        /// <summary>
        /// The number of datapoints that must be breaching to trigger the alarm.
        /// </summary>
        [Output("datapointsToAlarm")]
        public Output<int?> DatapointsToAlarm { get; private set; } = null!;

        /// <summary>
        /// The dimensions for the metric associated with the alarm. For an alarm based on a math expression, you can't specify Dimensions. Instead, you use Metrics.
        /// </summary>
        [Output("dimensions")]
        public Output<ImmutableArray<Outputs.AlarmDimension>> Dimensions { get; private set; } = null!;

        /// <summary>
        /// Used only for alarms based on percentiles.
        /// </summary>
        [Output("evaluateLowSampleCountPercentile")]
        public Output<string?> EvaluateLowSampleCountPercentile { get; private set; } = null!;

        /// <summary>
        /// The number of periods over which data is compared to the specified threshold.
        /// </summary>
        [Output("evaluationPeriods")]
        public Output<int> EvaluationPeriods { get; private set; } = null!;

        /// <summary>
        /// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
        /// </summary>
        [Output("extendedStatistic")]
        public Output<string?> ExtendedStatistic { get; private set; } = null!;

        /// <summary>
        /// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.
        /// </summary>
        [Output("insufficientDataActions")]
        public Output<ImmutableArray<string>> InsufficientDataActions { get; private set; } = null!;

        /// <summary>
        /// The name of the metric associated with the alarm.
        /// </summary>
        [Output("metricName")]
        public Output<string?> MetricName { get; private set; } = null!;

        /// <summary>
        /// An array that enables you to create an alarm based on the result of a metric math expression.
        /// </summary>
        [Output("metrics")]
        public Output<ImmutableArray<Outputs.AlarmMetricDataQuery>> Metrics { get; private set; } = null!;

        /// <summary>
        /// The namespace of the metric associated with the alarm.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The actions to execute when this alarm transitions to the OK state from any other state.
        /// </summary>
        [Output("okActions")]
        public Output<ImmutableArray<string>> OkActions { get; private set; } = null!;

        /// <summary>
        /// The period in seconds, over which the statistic is applied.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// The statistic for the metric associated with the alarm, other than percentile.
        /// </summary>
        [Output("statistic")]
        public Output<string?> Statistic { get; private set; } = null!;

        /// <summary>
        /// In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.
        /// </summary>
        [Output("threshold")]
        public Output<double?> Threshold { get; private set; } = null!;

        /// <summary>
        /// In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.
        /// </summary>
        [Output("thresholdMetricId")]
        public Output<string?> ThresholdMetricId { get; private set; } = null!;

        /// <summary>
        /// Sets how this alarm is to handle missing data points. Valid values are breaching, notBreaching, ignore, and missing.
        /// </summary>
        [Output("treatMissingData")]
        public Output<string?> TreatMissingData { get; private set; } = null!;

        /// <summary>
        /// The unit of the metric associated with the alarm.
        /// </summary>
        [Output("unit")]
        public Output<string?> Unit { get; private set; } = null!;


        /// <summary>
        /// Create a Alarm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Alarm(string name, AlarmArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cloudwatch:Alarm", name, args ?? new AlarmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Alarm(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudwatch:Alarm", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Alarm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Alarm Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Alarm(name, id, options);
        }
    }

    public sealed class AlarmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
        /// </summary>
        [Input("actionsEnabled")]
        public Input<bool>? ActionsEnabled { get; set; }

        [Input("alarmActions")]
        private InputList<string>? _alarmActions;

        /// <summary>
        /// The list of actions to execute when this alarm transitions into an ALARM state from any other state.
        /// </summary>
        public InputList<string> AlarmActions
        {
            get => _alarmActions ?? (_alarmActions = new InputList<string>());
            set => _alarmActions = value;
        }

        /// <summary>
        /// The description of the alarm.
        /// </summary>
        [Input("alarmDescription")]
        public Input<string>? AlarmDescription { get; set; }

        /// <summary>
        /// The name of the alarm.
        /// </summary>
        [Input("alarmName")]
        public Input<string>? AlarmName { get; set; }

        /// <summary>
        /// The arithmetic operation to use when comparing the specified statistic and threshold.
        /// </summary>
        [Input("comparisonOperator", required: true)]
        public Input<string> ComparisonOperator { get; set; } = null!;

        /// <summary>
        /// The number of datapoints that must be breaching to trigger the alarm.
        /// </summary>
        [Input("datapointsToAlarm")]
        public Input<int>? DatapointsToAlarm { get; set; }

        [Input("dimensions")]
        private InputList<Inputs.AlarmDimensionArgs>? _dimensions;

        /// <summary>
        /// The dimensions for the metric associated with the alarm. For an alarm based on a math expression, you can't specify Dimensions. Instead, you use Metrics.
        /// </summary>
        public InputList<Inputs.AlarmDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.AlarmDimensionArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// Used only for alarms based on percentiles.
        /// </summary>
        [Input("evaluateLowSampleCountPercentile")]
        public Input<string>? EvaluateLowSampleCountPercentile { get; set; }

        /// <summary>
        /// The number of periods over which data is compared to the specified threshold.
        /// </summary>
        [Input("evaluationPeriods", required: true)]
        public Input<int> EvaluationPeriods { get; set; } = null!;

        /// <summary>
        /// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
        /// </summary>
        [Input("extendedStatistic")]
        public Input<string>? ExtendedStatistic { get; set; }

        [Input("insufficientDataActions")]
        private InputList<string>? _insufficientDataActions;

        /// <summary>
        /// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.
        /// </summary>
        public InputList<string> InsufficientDataActions
        {
            get => _insufficientDataActions ?? (_insufficientDataActions = new InputList<string>());
            set => _insufficientDataActions = value;
        }

        /// <summary>
        /// The name of the metric associated with the alarm.
        /// </summary>
        [Input("metricName")]
        public Input<string>? MetricName { get; set; }

        [Input("metrics")]
        private InputList<Inputs.AlarmMetricDataQueryArgs>? _metrics;

        /// <summary>
        /// An array that enables you to create an alarm based on the result of a metric math expression.
        /// </summary>
        public InputList<Inputs.AlarmMetricDataQueryArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<Inputs.AlarmMetricDataQueryArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// The namespace of the metric associated with the alarm.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("okActions")]
        private InputList<string>? _okActions;

        /// <summary>
        /// The actions to execute when this alarm transitions to the OK state from any other state.
        /// </summary>
        public InputList<string> OkActions
        {
            get => _okActions ?? (_okActions = new InputList<string>());
            set => _okActions = value;
        }

        /// <summary>
        /// The period in seconds, over which the statistic is applied.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The statistic for the metric associated with the alarm, other than percentile.
        /// </summary>
        [Input("statistic")]
        public Input<string>? Statistic { get; set; }

        /// <summary>
        /// In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.
        /// </summary>
        [Input("threshold")]
        public Input<double>? Threshold { get; set; }

        /// <summary>
        /// In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.
        /// </summary>
        [Input("thresholdMetricId")]
        public Input<string>? ThresholdMetricId { get; set; }

        /// <summary>
        /// Sets how this alarm is to handle missing data points. Valid values are breaching, notBreaching, ignore, and missing.
        /// </summary>
        [Input("treatMissingData")]
        public Input<string>? TreatMissingData { get; set; }

        /// <summary>
        /// The unit of the metric associated with the alarm.
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        public AlarmArgs()
        {
        }
        public static new AlarmArgs Empty => new AlarmArgs();
    }
}
