// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rum.Outputs
{

    /// <summary>
    /// A single metric definition
    /// </summary>
    [OutputType]
    public sealed class AppMonitorMetricDefinition
    {
        /// <summary>
        /// Use this field only if you are sending the metric to CloudWatch.
        /// 
        /// This field is a map of field paths to dimension names. It defines the dimensions to associate with this metric in CloudWatch. For extended metrics, valid values for the entries in this field are the following:
        /// 
        /// "metadata.pageId": "PageId"
        /// 
        /// "metadata.browserName": "BrowserName"
        /// 
        /// "metadata.deviceType": "DeviceType"
        /// 
        /// "metadata.osName": "OSName"
        /// 
        /// "metadata.countryCode": "CountryCode"
        /// 
        /// "event_details.fileType": "FileType"
        /// 
        /// All dimensions listed in this field must also be included in EventPattern.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? DimensionKeys;
        /// <summary>
        /// The pattern that defines the metric, specified as a JSON object. RUM checks events that happen in a user's session against the pattern, and events that match the pattern are sent to the metric destination.
        /// 
        /// When you define extended metrics, the metric definition is not valid if EventPattern is omitted.
        /// 
        /// Example event patterns:
        /// 
        /// '{ "event_type": ["com.amazon.rum.js_error_event"], "metadata": { "browserName": [ "Chrome", "Safari" ], } }'
        /// 
        /// '{ "event_type": ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [ "Chrome", "Firefox" ] }, "event_details": { "duration": [{ "numeric": [ "&lt;", 2000 ] }] } }'
        /// 
        /// '{ "event_type": ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [ "Chrome", "Safari" ], "countryCode": [ "US" ] }, "event_details": { "duration": [{ "numeric": [ "&gt;=", 2000, "&lt;", 8000 ] }] } }'
        /// 
        /// If the metrics destination' is CloudWatch and the event also matches a value in DimensionKeys, then the metric is published with the specified dimensions.
        /// </summary>
        public readonly string? EventPattern;
        /// <summary>
        /// The name for the metric that is defined in this structure. For extended metrics, valid values are the following:
        /// 
        /// PerformanceNavigationDuration
        /// 
        /// PerformanceResourceDuration
        /// 
        /// NavigationSatisfiedTransaction
        /// 
        /// NavigationToleratedTransaction
        /// 
        /// NavigationFrustratedTransaction
        /// 
        /// WebVitalsCumulativeLayoutShift
        /// 
        /// WebVitalsFirstInputDelay
        /// 
        /// WebVitalsLargestContentfulPaint
        /// 
        /// JsErrorCount
        /// 
        /// HttpErrorCount
        /// 
        /// SessionCount
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The namespace used by CloudWatch Metrics for the metric that is defined in this structure
        /// </summary>
        public readonly string? Namespace;
        /// <summary>
        /// The CloudWatch metric unit to use for this metric. If you omit this field, the metric is recorded with no unit.
        /// </summary>
        public readonly string? UnitLabel;
        /// <summary>
        /// The field within the event object that the metric value is sourced from.
        /// 
        /// If you omit this field, a hardcoded value of 1 is pushed as the metric value. This is useful if you just want to count the number of events that the filter catches.
        /// 
        /// If this metric is sent to Evidently, this field will be passed to Evidently raw and Evidently will handle data extraction from the event.
        /// </summary>
        public readonly string? ValueKey;

        [OutputConstructor]
        private AppMonitorMetricDefinition(
            ImmutableDictionary<string, string>? dimensionKeys,

            string? eventPattern,

            string name,

            string? @namespace,

            string? unitLabel,

            string? valueKey)
        {
            DimensionKeys = dimensionKeys;
            EventPattern = eventPattern;
            Name = name;
            Namespace = @namespace;
            UnitLabel = unitLabel;
            ValueKey = valueKey;
        }
    }
}
