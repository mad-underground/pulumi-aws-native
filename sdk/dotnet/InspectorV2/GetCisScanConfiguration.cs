// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.InspectorV2
{
    public static class GetCisScanConfiguration
    {
        /// <summary>
        /// CIS Scan Configuration resource schema
        /// </summary>
        public static Task<GetCisScanConfigurationResult> InvokeAsync(GetCisScanConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCisScanConfigurationResult>("aws-native:inspectorv2:getCisScanConfiguration", args ?? new GetCisScanConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// CIS Scan Configuration resource schema
        /// </summary>
        public static Output<GetCisScanConfigurationResult> Invoke(GetCisScanConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCisScanConfigurationResult>("aws-native:inspectorv2:getCisScanConfiguration", args ?? new GetCisScanConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCisScanConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// CIS Scan configuration unique identifier
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetCisScanConfigurationArgs()
        {
        }
        public static new GetCisScanConfigurationArgs Empty => new GetCisScanConfigurationArgs();
    }

    public sealed class GetCisScanConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// CIS Scan configuration unique identifier
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetCisScanConfigurationInvokeArgs()
        {
        }
        public static new GetCisScanConfigurationInvokeArgs Empty => new GetCisScanConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetCisScanConfigurationResult
    {
        /// <summary>
        /// CIS Scan configuration unique identifier
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Name of the scan
        /// </summary>
        public readonly string? ScanName;
        public readonly Outputs.CisScanConfigurationSchedule? Schedule;
        public readonly Pulumi.AwsNative.InspectorV2.CisScanConfigurationCisSecurityLevel? SecurityLevel;
        public readonly ImmutableDictionary<string, string>? Tags;
        public readonly Outputs.CisScanConfigurationCisTargets? Targets;

        [OutputConstructor]
        private GetCisScanConfigurationResult(
            string? arn,

            string? scanName,

            Outputs.CisScanConfigurationSchedule? schedule,

            Pulumi.AwsNative.InspectorV2.CisScanConfigurationCisSecurityLevel? securityLevel,

            ImmutableDictionary<string, string>? tags,

            Outputs.CisScanConfigurationCisTargets? targets)
        {
            Arn = arn;
            ScanName = scanName;
            Schedule = schedule;
            SecurityLevel = securityLevel;
            Tags = tags;
            Targets = targets;
        }
    }
}
