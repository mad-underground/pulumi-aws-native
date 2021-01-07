// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.DataBrew.Outputs
{

    [OutputType]
    public sealed class JobProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-datasetname
        /// </summary>
        public readonly string? DatasetName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-encryptionkeyarn
        /// </summary>
        public readonly string? EncryptionKeyArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-encryptionmode
        /// </summary>
        public readonly string? EncryptionMode;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-logsubscription
        /// </summary>
        public readonly string? LogSubscription;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-maxcapacity
        /// </summary>
        public readonly int? MaxCapacity;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-maxretries
        /// </summary>
        public readonly int? MaxRetries;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-outputlocation
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? OutputLocation;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-outputs
        /// </summary>
        public readonly ImmutableArray<Outputs.JobOutput> Outputs;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-projectname
        /// </summary>
        public readonly string? ProjectName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-recipe
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? Recipe;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-rolearn
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-timeout
        /// </summary>
        public readonly int? Timeout;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private JobProperties(
            string? DatasetName,

            string? EncryptionKeyArn,

            string? EncryptionMode,

            string? LogSubscription,

            int? MaxCapacity,

            int? MaxRetries,

            string Name,

            Union<System.Text.Json.JsonElement, string>? OutputLocation,

            ImmutableArray<Outputs.JobOutput> Outputs,

            string? ProjectName,

            Union<System.Text.Json.JsonElement, string>? Recipe,

            string RoleArn,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags,

            int? Timeout,

            string Type)
        {
            this.DatasetName = DatasetName;
            this.EncryptionKeyArn = EncryptionKeyArn;
            this.EncryptionMode = EncryptionMode;
            this.LogSubscription = LogSubscription;
            this.MaxCapacity = MaxCapacity;
            this.MaxRetries = MaxRetries;
            this.Name = Name;
            this.OutputLocation = OutputLocation;
            this.Outputs = Outputs;
            this.ProjectName = ProjectName;
            this.Recipe = Recipe;
            this.RoleArn = RoleArn;
            this.Tags = Tags;
            this.Timeout = Timeout;
            this.Type = Type;
        }
    }
}