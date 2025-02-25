// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Outputs
{

    [OutputType]
    public sealed class CampaignHook
    {
        public readonly string? LambdaFunctionName;
        public readonly string? Mode;
        public readonly string? WebUrl;

        [OutputConstructor]
        private CampaignHook(
            string? lambdaFunctionName,

            string? mode,

            string? webUrl)
        {
            LambdaFunctionName = lambdaFunctionName;
            Mode = mode;
            WebUrl = webUrl;
        }
    }
}
