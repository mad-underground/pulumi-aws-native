// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Inputs
{

    /// <summary>
    /// The process of matching duplicate profiles. If Matching = true, Amazon Connect Customer Profiles starts a weekly batch process called Identity Resolution Job. If you do not specify a date and time for Identity Resolution Job to run, by default it runs every Saturday at 12AM UTC to detect duplicate profiles in your domains. After the Identity Resolution Job completes, use the GetMatches API to return and review the results. Or, if you have configured ExportingConfig in the MatchingRequest, you can download the results from S3.
    /// </summary>
    public sealed class DomainMatchingArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoMerging")]
        public Input<Inputs.DomainAutoMergingArgs>? AutoMerging { get; set; }

        /// <summary>
        /// The flag that enables the matching process of duplicate profiles.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("exportingConfig")]
        public Input<Inputs.DomainExportingConfigArgs>? ExportingConfig { get; set; }

        [Input("jobSchedule")]
        public Input<Inputs.DomainJobScheduleArgs>? JobSchedule { get; set; }

        public DomainMatchingArgs()
        {
        }
        public static new DomainMatchingArgs Empty => new DomainMatchingArgs();
    }
}
