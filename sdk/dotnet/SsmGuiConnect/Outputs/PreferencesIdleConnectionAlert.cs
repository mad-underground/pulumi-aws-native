// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SsmGuiConnect.Outputs
{

    [OutputType]
    public sealed class PreferencesIdleConnectionAlert
    {
        public readonly Pulumi.AwsNative.SsmGuiConnect.PreferencesIdleConnectionAlertType? Type;
        public readonly int Value;

        [OutputConstructor]
        private PreferencesIdleConnectionAlert(
            Pulumi.AwsNative.SsmGuiConnect.PreferencesIdleConnectionAlertType? type,

            int value)
        {
            Type = type;
            Value = value;
        }
    }
}
