// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateCalculatedFieldArgs : global::Pulumi.ResourceArgs
    {
        [Input("dataSetIdentifier", required: true)]
        public Input<string> DataSetIdentifier { get; set; } = null!;

        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public TemplateCalculatedFieldArgs()
        {
        }
        public static new TemplateCalculatedFieldArgs Empty => new TemplateCalculatedFieldArgs();
    }
}
