// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// &lt;p&gt;Information about the format for a source file or files.&lt;/p&gt;
    /// </summary>
    public sealed class DataSetUploadSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;Whether the file has a header row, or the files each have a header row.&lt;/p&gt;
        /// </summary>
        [Input("containsHeader")]
        public Input<bool>? ContainsHeader { get; set; }

        /// <summary>
        /// &lt;p&gt;The delimiter between values in the file.&lt;/p&gt;
        /// </summary>
        [Input("delimiter")]
        public Input<string>? Delimiter { get; set; }

        [Input("format")]
        public Input<Pulumi.AwsNative.QuickSight.DataSetFileFormat>? Format { get; set; }

        /// <summary>
        /// &lt;p&gt;A row number to start reading data from.&lt;/p&gt;
        /// </summary>
        [Input("startFromRow")]
        public Input<double>? StartFromRow { get; set; }

        [Input("textQualifier")]
        public Input<Pulumi.AwsNative.QuickSight.DataSetTextQualifier>? TextQualifier { get; set; }

        public DataSetUploadSettingsArgs()
        {
        }
        public static new DataSetUploadSettingsArgs Empty => new DataSetUploadSettingsArgs();
    }
}
