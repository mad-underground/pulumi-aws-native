// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateDataSetReferenceArgs : global::Pulumi.ResourceArgs
    {
        [Input("dataSetArn", required: true)]
        public Input<string> DataSetArn { get; set; } = null!;

        [Input("dataSetPlaceholder", required: true)]
        public Input<string> DataSetPlaceholder { get; set; } = null!;

        public TemplateDataSetReferenceArgs()
        {
        }
        public static new TemplateDataSetReferenceArgs Empty => new TemplateDataSetReferenceArgs();
    }
}
