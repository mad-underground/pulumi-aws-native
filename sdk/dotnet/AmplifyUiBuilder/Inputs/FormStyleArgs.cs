// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class FormStyleArgs : global::Pulumi.ResourceArgs
    {
        [Input("horizontalGap")]
        public InputUnion<Inputs.FormStyleConfig0PropertiesArgs, Inputs.FormStyleConfig1PropertiesArgs>? HorizontalGap { get; set; }

        [Input("outerPadding")]
        public InputUnion<Inputs.FormStyleConfig0PropertiesArgs, Inputs.FormStyleConfig1PropertiesArgs>? OuterPadding { get; set; }

        [Input("verticalGap")]
        public InputUnion<Inputs.FormStyleConfig0PropertiesArgs, Inputs.FormStyleConfig1PropertiesArgs>? VerticalGap { get; set; }

        public FormStyleArgs()
        {
        }
        public static new FormStyleArgs Empty => new FormStyleArgs();
    }
}
