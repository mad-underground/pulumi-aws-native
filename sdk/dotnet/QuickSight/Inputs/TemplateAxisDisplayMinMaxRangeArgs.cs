// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateAxisDisplayMinMaxRangeArgs : global::Pulumi.ResourceArgs
    {
        [Input("maximum")]
        public Input<double>? Maximum { get; set; }

        [Input("minimum")]
        public Input<double>? Minimum { get; set; }

        public TemplateAxisDisplayMinMaxRangeArgs()
        {
        }
        public static new TemplateAxisDisplayMinMaxRangeArgs Empty => new TemplateAxisDisplayMinMaxRangeArgs();
    }
}
