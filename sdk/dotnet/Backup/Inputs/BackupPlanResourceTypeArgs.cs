// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Backup.Inputs
{

    public sealed class BackupPlanResourceTypeArgs : global::Pulumi.ResourceArgs
    {
        [Input("advancedBackupSettings")]
        private InputList<Inputs.BackupPlanAdvancedBackupSettingResourceTypeArgs>? _advancedBackupSettings;
        public InputList<Inputs.BackupPlanAdvancedBackupSettingResourceTypeArgs> AdvancedBackupSettings
        {
            get => _advancedBackupSettings ?? (_advancedBackupSettings = new InputList<Inputs.BackupPlanAdvancedBackupSettingResourceTypeArgs>());
            set => _advancedBackupSettings = value;
        }

        [Input("backupPlanName", required: true)]
        public Input<string> BackupPlanName { get; set; } = null!;

        [Input("backupPlanRule", required: true)]
        private InputList<Inputs.BackupPlanBackupRuleResourceTypeArgs>? _backupPlanRule;
        public InputList<Inputs.BackupPlanBackupRuleResourceTypeArgs> BackupPlanRule
        {
            get => _backupPlanRule ?? (_backupPlanRule = new InputList<Inputs.BackupPlanBackupRuleResourceTypeArgs>());
            set => _backupPlanRule = value;
        }

        public BackupPlanResourceTypeArgs()
        {
        }
        public static new BackupPlanResourceTypeArgs Empty => new BackupPlanResourceTypeArgs();
    }
}
