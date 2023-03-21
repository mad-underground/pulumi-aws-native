// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateShortFormatText
    {
        public readonly string? PlainText;
        public readonly string? RichText;

        [OutputConstructor]
        private TemplateShortFormatText(
            string? plainText,

            string? richText)
        {
            PlainText = plainText;
            RichText = richText;
        }
    }
}
