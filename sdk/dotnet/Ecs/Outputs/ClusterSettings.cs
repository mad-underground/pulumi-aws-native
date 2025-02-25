// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Outputs
{

    /// <summary>
    /// The setting to use when creating a cluster. This parameter is used to enable CloudWatch Container Insights for a cluster. If this value is specified, it will override the containerInsights value set with PutAccountSetting or PutAccountSettingDefault.
    /// </summary>
    [OutputType]
    public sealed class ClusterSettings
    {
        public readonly string? Name;
        public readonly string? Value;

        [OutputConstructor]
        private ClusterSettings(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
