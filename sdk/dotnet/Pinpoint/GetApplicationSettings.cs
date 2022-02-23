// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint
{
    public static class GetApplicationSettings
    {
        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::ApplicationSettings
        /// </summary>
        public static Task<GetApplicationSettingsResult> InvokeAsync(GetApplicationSettingsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationSettingsResult>("aws-native:pinpoint:getApplicationSettings", args ?? new GetApplicationSettingsArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::ApplicationSettings
        /// </summary>
        public static Output<GetApplicationSettingsResult> Invoke(GetApplicationSettingsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetApplicationSettingsResult>("aws-native:pinpoint:getApplicationSettings", args ?? new GetApplicationSettingsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationSettingsArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetApplicationSettingsArgs()
        {
        }
    }

    public sealed class GetApplicationSettingsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetApplicationSettingsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApplicationSettingsResult
    {
        public readonly Outputs.ApplicationSettingsCampaignHook? CampaignHook;
        public readonly bool? CloudWatchMetricsEnabled;
        public readonly string? Id;
        public readonly Outputs.ApplicationSettingsLimits? Limits;
        public readonly Outputs.ApplicationSettingsQuietTime? QuietTime;

        [OutputConstructor]
        private GetApplicationSettingsResult(
            Outputs.ApplicationSettingsCampaignHook? campaignHook,

            bool? cloudWatchMetricsEnabled,

            string? id,

            Outputs.ApplicationSettingsLimits? limits,

            Outputs.ApplicationSettingsQuietTime? quietTime)
        {
            CampaignHook = campaignHook;
            CloudWatchMetricsEnabled = cloudWatchMetricsEnabled;
            Id = id;
            Limits = limits;
            QuietTime = quietTime;
        }
    }
}