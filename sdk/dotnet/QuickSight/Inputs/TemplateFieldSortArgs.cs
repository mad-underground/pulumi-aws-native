// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateFieldSortArgs : global::Pulumi.ResourceArgs
    {
        [Input("direction", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.TemplateSortDirection> Direction { get; set; } = null!;

        [Input("fieldId", required: true)]
        public Input<string> FieldId { get; set; } = null!;

        public TemplateFieldSortArgs()
        {
        }
        public static new TemplateFieldSortArgs Empty => new TemplateFieldSortArgs();
    }
}
