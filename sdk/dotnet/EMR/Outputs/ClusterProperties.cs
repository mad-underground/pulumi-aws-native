// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.EMR.Outputs
{

    [OutputType]
    public sealed class ClusterProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-additionalinfo
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? AdditionalInfo;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-applications
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterApplication> Applications;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-autoscalingrole
        /// </summary>
        public readonly string? AutoScalingRole;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-bootstrapactions
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterBootstrapActionConfig> BootstrapActions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-configurations
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterConfiguration> Configurations;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-customamiid
        /// </summary>
        public readonly string? CustomAmiId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-ebsrootvolumesize
        /// </summary>
        public readonly int? EbsRootVolumeSize;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-instances
        /// </summary>
        public readonly Outputs.ClusterJobFlowInstancesConfig Instances;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-jobflowrole
        /// </summary>
        public readonly string JobFlowRole;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-kerberosattributes
        /// </summary>
        public readonly Outputs.ClusterKerberosAttributes? KerberosAttributes;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-logencryptionkmskeyid
        /// </summary>
        public readonly string? LogEncryptionKmsKeyId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-loguri
        /// </summary>
        public readonly string? LogUri;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-managedscalingpolicy
        /// </summary>
        public readonly Outputs.ClusterManagedScalingPolicy? ManagedScalingPolicy;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-releaselabel
        /// </summary>
        public readonly string? ReleaseLabel;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-scaledownbehavior
        /// </summary>
        public readonly string? ScaleDownBehavior;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-securityconfiguration
        /// </summary>
        public readonly string? SecurityConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-servicerole
        /// </summary>
        public readonly string ServiceRole;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-stepconcurrencylevel
        /// </summary>
        public readonly int? StepConcurrencyLevel;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-steps
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterStepConfig> Steps;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html#cfn-elasticmapreduce-cluster-visibletoallusers
        /// </summary>
        public readonly bool? VisibleToAllUsers;

        [OutputConstructor]
        private ClusterProperties(
            Union<System.Text.Json.JsonElement, string>? AdditionalInfo,

            ImmutableArray<Outputs.ClusterApplication> Applications,

            string? AutoScalingRole,

            ImmutableArray<Outputs.ClusterBootstrapActionConfig> BootstrapActions,

            ImmutableArray<Outputs.ClusterConfiguration> Configurations,

            string? CustomAmiId,

            int? EbsRootVolumeSize,

            Outputs.ClusterJobFlowInstancesConfig Instances,

            string JobFlowRole,

            Outputs.ClusterKerberosAttributes? KerberosAttributes,

            string? LogEncryptionKmsKeyId,

            string? LogUri,

            Outputs.ClusterManagedScalingPolicy? ManagedScalingPolicy,

            string Name,

            string? ReleaseLabel,

            string? ScaleDownBehavior,

            string? SecurityConfiguration,

            string ServiceRole,

            int? StepConcurrencyLevel,

            ImmutableArray<Outputs.ClusterStepConfig> Steps,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags,

            bool? VisibleToAllUsers)
        {
            this.AdditionalInfo = AdditionalInfo;
            this.Applications = Applications;
            this.AutoScalingRole = AutoScalingRole;
            this.BootstrapActions = BootstrapActions;
            this.Configurations = Configurations;
            this.CustomAmiId = CustomAmiId;
            this.EbsRootVolumeSize = EbsRootVolumeSize;
            this.Instances = Instances;
            this.JobFlowRole = JobFlowRole;
            this.KerberosAttributes = KerberosAttributes;
            this.LogEncryptionKmsKeyId = LogEncryptionKmsKeyId;
            this.LogUri = LogUri;
            this.ManagedScalingPolicy = ManagedScalingPolicy;
            this.Name = Name;
            this.ReleaseLabel = ReleaseLabel;
            this.ScaleDownBehavior = ScaleDownBehavior;
            this.SecurityConfiguration = SecurityConfiguration;
            this.ServiceRole = ServiceRole;
            this.StepConcurrencyLevel = StepConcurrencyLevel;
            this.Steps = Steps;
            this.Tags = Tags;
            this.VisibleToAllUsers = VisibleToAllUsers;
        }
    }
}