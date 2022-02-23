// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evidently
{
    public static class GetLaunch
    {
        /// <summary>
        /// Resource Type definition for AWS::Evidently::Launch.
        /// </summary>
        public static Task<GetLaunchResult> InvokeAsync(GetLaunchArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLaunchResult>("aws-native:evidently:getLaunch", args ?? new GetLaunchArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Evidently::Launch.
        /// </summary>
        public static Output<GetLaunchResult> Invoke(GetLaunchInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLaunchResult>("aws-native:evidently:getLaunch", args ?? new GetLaunchInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLaunchArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetLaunchArgs()
        {
        }
    }

    public sealed class GetLaunchInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetLaunchInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLaunchResult
    {
        public readonly string? Arn;
        public readonly string? Description;
        public readonly ImmutableArray<Outputs.LaunchGroupObject> Groups;
        public readonly ImmutableArray<Outputs.LaunchMetricDefinitionObject> MetricMonitors;
        public readonly string? RandomizationSalt;
        public readonly ImmutableArray<Outputs.LaunchStepConfig> ScheduledSplitsConfig;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.LaunchTag> Tags;

        [OutputConstructor]
        private GetLaunchResult(
            string? arn,

            string? description,

            ImmutableArray<Outputs.LaunchGroupObject> groups,

            ImmutableArray<Outputs.LaunchMetricDefinitionObject> metricMonitors,

            string? randomizationSalt,

            ImmutableArray<Outputs.LaunchStepConfig> scheduledSplitsConfig,

            ImmutableArray<Outputs.LaunchTag> tags)
        {
            Arn = arn;
            Description = description;
            Groups = groups;
            MetricMonitors = metricMonitors;
            RandomizationSalt = randomizationSalt;
            ScheduledSplitsConfig = scheduledSplitsConfig;
            Tags = tags;
        }
    }
}