// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Fis.Inputs
{

    public sealed class ExperimentTemplateLogConfigurationS3ConfigurationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public ExperimentTemplateLogConfigurationS3ConfigurationPropertiesArgs()
        {
        }
        public static new ExperimentTemplateLogConfigurationS3ConfigurationPropertiesArgs Empty => new ExperimentTemplateLogConfigurationS3ConfigurationPropertiesArgs();
    }
}
