// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Backup.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html
    /// </summary>
    [OutputType]
    public sealed class BackupPlanBackupRuleResourceType
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-completionwindowminutes
        /// </summary>
        public readonly double? CompletionWindowMinutes;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-copyactions
        /// </summary>
        public readonly ImmutableArray<Outputs.BackupPlanCopyActionResourceType> CopyActions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-lifecycle
        /// </summary>
        public readonly Outputs.BackupPlanLifecycleResourceType? Lifecycle;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-recoverypointtags
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? RecoveryPointTags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-rulename
        /// </summary>
        public readonly string RuleName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-scheduleexpression
        /// </summary>
        public readonly string? ScheduleExpression;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-startwindowminutes
        /// </summary>
        public readonly double? StartWindowMinutes;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-targetbackupvault
        /// </summary>
        public readonly string TargetBackupVault;

        [OutputConstructor]
        private BackupPlanBackupRuleResourceType(
            double? completionWindowMinutes,

            ImmutableArray<Outputs.BackupPlanCopyActionResourceType> copyActions,

            Outputs.BackupPlanLifecycleResourceType? lifecycle,

            Union<System.Text.Json.JsonElement, string>? recoveryPointTags,

            string ruleName,

            string? scheduleExpression,

            double? startWindowMinutes,

            string targetBackupVault)
        {
            CompletionWindowMinutes = completionWindowMinutes;
            CopyActions = copyActions;
            Lifecycle = lifecycle;
            RecoveryPointTags = recoveryPointTags;
            RuleName = ruleName;
            ScheduleExpression = scheduleExpression;
            StartWindowMinutes = startWindowMinutes;
            TargetBackupVault = targetBackupVault;
        }
    }
}
