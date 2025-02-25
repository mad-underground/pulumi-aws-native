// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Outputs
{

    [OutputType]
    public sealed class PipeLogConfiguration
    {
        public readonly Outputs.PipeCloudwatchLogsLogDestination? CloudwatchLogsLogDestination;
        public readonly Outputs.PipeFirehoseLogDestination? FirehoseLogDestination;
        public readonly ImmutableArray<Pulumi.AwsNative.Pipes.PipeIncludeExecutionDataOption> IncludeExecutionData;
        public readonly Pulumi.AwsNative.Pipes.PipeLogLevel? Level;
        public readonly Outputs.PipeS3LogDestination? S3LogDestination;

        [OutputConstructor]
        private PipeLogConfiguration(
            Outputs.PipeCloudwatchLogsLogDestination? cloudwatchLogsLogDestination,

            Outputs.PipeFirehoseLogDestination? firehoseLogDestination,

            ImmutableArray<Pulumi.AwsNative.Pipes.PipeIncludeExecutionDataOption> includeExecutionData,

            Pulumi.AwsNative.Pipes.PipeLogLevel? level,

            Outputs.PipeS3LogDestination? s3LogDestination)
        {
            CloudwatchLogsLogDestination = cloudwatchLogsLogDestination;
            FirehoseLogDestination = firehoseLogDestination;
            IncludeExecutionData = includeExecutionData;
            Level = level;
            S3LogDestination = s3LogDestination;
        }
    }
}
