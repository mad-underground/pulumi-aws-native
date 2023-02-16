// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUIBuilder.Outputs
{

    [OutputType]
    public sealed class FormStyle
    {
        public readonly Outputs.FormStyleConfig? HorizontalGap;
        public readonly Outputs.FormStyleConfig? OuterPadding;
        public readonly Outputs.FormStyleConfig? VerticalGap;

        [OutputConstructor]
        private FormStyle(
            Outputs.FormStyleConfig? horizontalGap,

            Outputs.FormStyleConfig? outerPadding,

            Outputs.FormStyleConfig? verticalGap)
        {
            HorizontalGap = horizontalGap;
            OuterPadding = outerPadding;
            VerticalGap = verticalGap;
        }
    }
}
