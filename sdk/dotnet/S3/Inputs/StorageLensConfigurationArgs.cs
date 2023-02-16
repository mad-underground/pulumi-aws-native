// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// Specifies the details of Amazon S3 Storage Lens configuration.
    /// </summary>
    public sealed class StorageLensConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountLevel", required: true)]
        public Input<Inputs.StorageLensAccountLevelArgs> AccountLevel { get; set; } = null!;

        [Input("awsOrg")]
        public Input<Inputs.StorageLensAwsOrgArgs>? AwsOrg { get; set; }

        [Input("dataExport")]
        public Input<Inputs.StorageLensDataExportArgs>? DataExport { get; set; }

        [Input("exclude")]
        public Input<Inputs.StorageLensBucketsAndRegionsArgs>? Exclude { get; set; }

        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("include")]
        public Input<Inputs.StorageLensBucketsAndRegionsArgs>? Include { get; set; }

        /// <summary>
        /// Specifies whether the Amazon S3 Storage Lens configuration is enabled or disabled.
        /// </summary>
        [Input("isEnabled", required: true)]
        public Input<bool> IsEnabled { get; set; } = null!;

        /// <summary>
        /// The ARN for the Amazon S3 Storage Lens configuration.
        /// </summary>
        [Input("storageLensArn")]
        public Input<string>? StorageLensArn { get; set; }

        public StorageLensConfigurationArgs()
        {
        }
        public static new StorageLensConfigurationArgs Empty => new StorageLensConfigurationArgs();
    }
}
