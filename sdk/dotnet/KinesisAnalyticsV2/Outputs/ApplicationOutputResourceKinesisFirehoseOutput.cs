// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Outputs
{

    [OutputType]
    public sealed class ApplicationOutputResourceKinesisFirehoseOutput
    {
        public readonly string ResourceARN;

        [OutputConstructor]
        private ApplicationOutputResourceKinesisFirehoseOutput(string resourceARN)
        {
            ResourceARN = resourceARN;
        }
    }
}