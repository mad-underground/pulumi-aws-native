// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor.Outputs
{

    /// <summary>
    /// The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
    /// </summary>
    [OutputType]
    public sealed class PlaybackConfigurationAvailSuppression
    {
        /// <summary>
        /// Sets the ad suppression mode. By default, ad suppression is set to OFF and all ad breaks are filled with ads or slate. When Mode is set to BEHIND_LIVE_EDGE, ad suppression is active and MediaTailor won't fill ad breaks on or behind the ad suppression Value time in the manifest lookback window.
        /// </summary>
        public readonly Pulumi.AwsNative.MediaTailor.PlaybackConfigurationAvailSuppressionMode? Mode;
        /// <summary>
        /// A live edge offset time in HH:MM:SS. MediaTailor won't fill ad breaks on or behind this time in the manifest lookback window. If Value is set to 00:00:00, it is in sync with the live edge, and MediaTailor won't fill any ad breaks on or behind the live edge. If you set a Value time, MediaTailor won't fill any ad breaks on or behind this time in the manifest lookback window. For example, if you set 00:45:00, then MediaTailor will fill ad breaks that occur within 45 minutes behind the live edge, but won't fill ad breaks on or behind 45 minutes behind the live edge.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private PlaybackConfigurationAvailSuppression(
            Pulumi.AwsNative.MediaTailor.PlaybackConfigurationAvailSuppressionMode? mode,

            string? value)
        {
            Mode = mode;
            Value = value;
        }
    }
}
