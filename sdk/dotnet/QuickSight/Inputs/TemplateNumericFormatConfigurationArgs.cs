// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateNumericFormatConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("currencyDisplayFormatConfiguration")]
        public Input<Inputs.TemplateCurrencyDisplayFormatConfigurationArgs>? CurrencyDisplayFormatConfiguration { get; set; }

        [Input("numberDisplayFormatConfiguration")]
        public Input<Inputs.TemplateNumberDisplayFormatConfigurationArgs>? NumberDisplayFormatConfiguration { get; set; }

        [Input("percentageDisplayFormatConfiguration")]
        public Input<Inputs.TemplatePercentageDisplayFormatConfigurationArgs>? PercentageDisplayFormatConfiguration { get; set; }

        public TemplateNumericFormatConfigurationArgs()
        {
        }
        public static new TemplateNumericFormatConfigurationArgs Empty => new TemplateNumericFormatConfigurationArgs();
    }
}
