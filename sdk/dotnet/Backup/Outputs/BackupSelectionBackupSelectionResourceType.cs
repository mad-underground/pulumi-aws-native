// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Backup.Outputs
{

    [OutputType]
    public sealed class BackupSelectionBackupSelectionResourceType
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-backupselectionresourcetype.html#cfn-backup-backupselection-backupselectionresourcetype-iamrolearn
        /// </summary>
        public readonly string IamRoleArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-backupselectionresourcetype.html#cfn-backup-backupselection-backupselectionresourcetype-listoftags
        /// </summary>
        public readonly ImmutableArray<Outputs.BackupSelectionConditionResourceType> ListOfTags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-backupselectionresourcetype.html#cfn-backup-backupselection-backupselectionresourcetype-resources
        /// </summary>
        public readonly ImmutableArray<string> Resources;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-backupselectionresourcetype.html#cfn-backup-backupselection-backupselectionresourcetype-selectionname
        /// </summary>
        public readonly string SelectionName;

        [OutputConstructor]
        private BackupSelectionBackupSelectionResourceType(
            string IamRoleArn,

            ImmutableArray<Outputs.BackupSelectionConditionResourceType> ListOfTags,

            ImmutableArray<string> Resources,

            string SelectionName)
        {
            this.IamRoleArn = IamRoleArn;
            this.ListOfTags = ListOfTags;
            this.Resources = Resources;
            this.SelectionName = SelectionName;
        }
    }
}