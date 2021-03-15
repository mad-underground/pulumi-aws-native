// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Athena.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html
    /// </summary>
    public sealed class WorkGroupWorkGroupConfigurationUpdatesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-bytesscannedcutoffperquery
        /// </summary>
        [Input("bytesScannedCutoffPerQuery")]
        public Input<int>? BytesScannedCutoffPerQuery { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-enforceworkgroupconfiguration
        /// </summary>
        [Input("enforceWorkGroupConfiguration")]
        public Input<bool>? EnforceWorkGroupConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-publishcloudwatchmetricsenabled
        /// </summary>
        [Input("publishCloudWatchMetricsEnabled")]
        public Input<bool>? PublishCloudWatchMetricsEnabled { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-removebytesscannedcutoffperquery
        /// </summary>
        [Input("removeBytesScannedCutoffPerQuery")]
        public Input<bool>? RemoveBytesScannedCutoffPerQuery { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-requesterpaysenabled
        /// </summary>
        [Input("requesterPaysEnabled")]
        public Input<bool>? RequesterPaysEnabled { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-resultconfigurationupdates
        /// </summary>
        [Input("resultConfigurationUpdates")]
        public Input<Inputs.WorkGroupResultConfigurationUpdatesArgs>? ResultConfigurationUpdates { get; set; }

        public WorkGroupWorkGroupConfigurationUpdatesArgs()
        {
        }
    }
}
