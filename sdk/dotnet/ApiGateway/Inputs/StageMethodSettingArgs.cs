// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway.Inputs
{

    /// <summary>
    /// The ``MethodSetting`` property type configures settings for all methods in a stage.
    ///  The ``MethodSettings`` property of the ``AWS::ApiGateway::Stage`` resource contains a list of ``MethodSetting`` property types.
    /// </summary>
    public sealed class StageMethodSettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether the cached responses are encrypted.
        /// </summary>
        [Input("cacheDataEncrypted")]
        public Input<bool>? CacheDataEncrypted { get; set; }

        /// <summary>
        /// Specifies the time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.
        /// </summary>
        [Input("cacheTtlInSeconds")]
        public Input<int>? CacheTtlInSeconds { get; set; }

        /// <summary>
        /// Specifies whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached.
        /// </summary>
        [Input("cachingEnabled")]
        public Input<bool>? CachingEnabled { get; set; }

        /// <summary>
        /// Specifies whether data trace logging is enabled for this method, which affects the log entries pushed to Amazon CloudWatch Logs. This can be useful to troubleshoot APIs, but can result in logging sensitive data. We recommend that you don't enable this option for production APIs.
        /// </summary>
        [Input("dataTraceEnabled")]
        public Input<bool>? DataTraceEnabled { get; set; }

        /// <summary>
        /// The HTTP method. To apply settings to multiple resources and methods, specify an asterisk (``*``) for the ``HttpMethod`` and ``/*`` for the ``ResourcePath``. This parameter is required when you specify a ``MethodSetting``.
        /// </summary>
        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        /// <summary>
        /// Specifies the logging level for this method, which affects the log entries pushed to Amazon CloudWatch Logs. Valid values are ``OFF``, ``ERROR``, and ``INFO``. Choose ``ERROR`` to write only error-level entries to CloudWatch Logs, or choose ``INFO`` to include all ``ERROR`` events as well as extra informational events.
        /// </summary>
        [Input("loggingLevel")]
        public Input<string>? LoggingLevel { get; set; }

        /// <summary>
        /// Specifies whether Amazon CloudWatch metrics are enabled for this method.
        /// </summary>
        [Input("metricsEnabled")]
        public Input<bool>? MetricsEnabled { get; set; }

        /// <summary>
        /// The resource path for this method. Forward slashes (``/``) are encoded as ``~1`` and the initial slash must include a forward slash. For example, the path value ``/resource/subresource`` must be encoded as ``/~1resource~1subresource``. To specify the root path, use only a slash (``/``). To apply settings to multiple resources and methods, specify an asterisk (``*``) for the ``HttpMethod`` and ``/*`` for the ``ResourcePath``. This parameter is required when you specify a ``MethodSetting``.
        /// </summary>
        [Input("resourcePath")]
        public Input<string>? ResourcePath { get; set; }

        /// <summary>
        /// Specifies the throttling burst limit.
        /// </summary>
        [Input("throttlingBurstLimit")]
        public Input<int>? ThrottlingBurstLimit { get; set; }

        /// <summary>
        /// Specifies the throttling rate limit.
        /// </summary>
        [Input("throttlingRateLimit")]
        public Input<double>? ThrottlingRateLimit { get; set; }

        public StageMethodSettingArgs()
        {
        }
        public static new StageMethodSettingArgs Empty => new StageMethodSettingArgs();
    }
}
