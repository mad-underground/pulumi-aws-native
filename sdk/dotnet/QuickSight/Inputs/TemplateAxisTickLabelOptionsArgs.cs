// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateAxisTickLabelOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("labelOptions")]
        public Input<Inputs.TemplateLabelOptionsArgs>? LabelOptions { get; set; }

        [Input("rotationAngle")]
        public Input<double>? RotationAngle { get; set; }

        public TemplateAxisTickLabelOptionsArgs()
        {
        }
        public static new TemplateAxisTickLabelOptionsArgs Empty => new TemplateAxisTickLabelOptionsArgs();
    }
}
