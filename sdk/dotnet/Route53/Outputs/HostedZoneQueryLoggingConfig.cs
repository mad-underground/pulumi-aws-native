// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Route53.Outputs
{

    [OutputType]
    public sealed class HostedZoneQueryLoggingConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-queryloggingconfig.html#cfn-route53-hostedzone-queryloggingconfig-cloudwatchlogsloggrouparn
        /// </summary>
        public readonly string CloudWatchLogsLogGroupArn;

        [OutputConstructor]
        private HostedZoneQueryLoggingConfig(string CloudWatchLogsLogGroupArn)
        {
            this.CloudWatchLogsLogGroupArn = CloudWatchLogsLogGroupArn;
        }
    }
}