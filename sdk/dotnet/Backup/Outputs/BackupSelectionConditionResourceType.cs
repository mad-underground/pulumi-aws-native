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
    public sealed class BackupSelectionConditionResourceType
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditionresourcetype.html#cfn-backup-backupselection-conditionresourcetype-conditionkey
        /// </summary>
        public readonly string ConditionKey;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditionresourcetype.html#cfn-backup-backupselection-conditionresourcetype-conditiontype
        /// </summary>
        public readonly string ConditionType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupselection-conditionresourcetype.html#cfn-backup-backupselection-conditionresourcetype-conditionvalue
        /// </summary>
        public readonly string ConditionValue;

        [OutputConstructor]
        private BackupSelectionConditionResourceType(
            string ConditionKey,

            string ConditionType,

            string ConditionValue)
        {
            this.ConditionKey = ConditionKey;
            this.ConditionType = ConditionType;
            this.ConditionValue = ConditionValue;
        }
    }
}