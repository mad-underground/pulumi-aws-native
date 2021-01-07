// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Backup.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html
    /// </summary>
    public sealed class BackupVaultPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-accesspolicy
        /// </summary>
        [Input("AccessPolicy")]
        public InputUnion<System.Text.Json.JsonElement, string>? AccessPolicy { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-backupvaultname
        /// </summary>
        [Input("BackupVaultName", required: true)]
        public Input<string> BackupVaultName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-backupvaulttags
        /// </summary>
        [Input("BackupVaultTags")]
        public InputUnion<System.Text.Json.JsonElement, string>? BackupVaultTags { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-encryptionkeyarn
        /// </summary>
        [Input("EncryptionKeyArn")]
        public Input<string>? EncryptionKeyArn { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-notifications
        /// </summary>
        [Input("Notifications")]
        public Input<Inputs.BackupVaultNotificationObjectTypeArgs>? Notifications { get; set; }

        public BackupVaultPropertiesArgs()
        {
        }
    }
}