// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.IoTAnalytics.Outputs
{

    [OutputType]
    public sealed class DatastoreProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-datastore.html#cfn-iotanalytics-datastore-datastorename
        /// </summary>
        public readonly string? DatastoreName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-datastore.html#cfn-iotanalytics-datastore-datastorestorage
        /// </summary>
        public readonly Outputs.DatastoreDatastoreStorage? DatastoreStorage;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-datastore.html#cfn-iotanalytics-datastore-retentionperiod
        /// </summary>
        public readonly Outputs.DatastoreRetentionPeriod? RetentionPeriod;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-datastore.html#cfn-iotanalytics-datastore-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;

        [OutputConstructor]
        private DatastoreProperties(
            string? DatastoreName,

            Outputs.DatastoreDatastoreStorage? DatastoreStorage,

            Outputs.DatastoreRetentionPeriod? RetentionPeriod,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags)
        {
            this.DatastoreName = DatastoreName;
            this.DatastoreStorage = DatastoreStorage;
            this.RetentionPeriod = RetentionPeriod;
            this.Tags = Tags;
        }
    }
}