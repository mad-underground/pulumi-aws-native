// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    public sealed class DataSourceTemplateConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("template", required: true)]
        public Input<string> Template { get; set; } = null!;

        public DataSourceTemplateConfigurationArgs()
        {
        }
        public static new DataSourceTemplateConfigurationArgs Empty => new DataSourceTemplateConfigurationArgs();
    }
}