// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateFieldLabelTypeArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldId")]
        public Input<string>? FieldId { get; set; }

        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateVisibility>? Visibility { get; set; }

        public TemplateFieldLabelTypeArgs()
        {
        }
        public static new TemplateFieldLabelTypeArgs Empty => new TemplateFieldLabelTypeArgs();
    }
}