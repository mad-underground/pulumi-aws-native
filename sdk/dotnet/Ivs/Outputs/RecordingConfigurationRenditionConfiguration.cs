// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ivs.Outputs
{

    /// <summary>
    /// Rendition Configuration describes which renditions should be recorded for a stream.
    /// </summary>
    [OutputType]
    public sealed class RecordingConfigurationRenditionConfiguration
    {
        /// <summary>
        /// Resolution Selection indicates which set of renditions are recorded for a stream.
        /// </summary>
        public readonly Pulumi.AwsNative.Ivs.RecordingConfigurationRenditionConfigurationRenditionSelection? RenditionSelection;
        /// <summary>
        /// Renditions indicates which renditions are recorded for a stream.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Ivs.RecordingConfigurationRenditionConfigurationRenditionsItem> Renditions;

        [OutputConstructor]
        private RecordingConfigurationRenditionConfiguration(
            Pulumi.AwsNative.Ivs.RecordingConfigurationRenditionConfigurationRenditionSelection? renditionSelection,

            ImmutableArray<Pulumi.AwsNative.Ivs.RecordingConfigurationRenditionConfigurationRenditionsItem> renditions)
        {
            RenditionSelection = renditionSelection;
            Renditions = renditions;
        }
    }
}
