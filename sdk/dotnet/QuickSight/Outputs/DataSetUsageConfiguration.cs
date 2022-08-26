// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;The dataset usage configuration for the dataset.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSetUsageConfiguration
    {
        public readonly bool? DisableUseAsDirectQuerySource;
        public readonly bool? DisableUseAsImportedSource;

        [OutputConstructor]
        private DataSetUsageConfiguration(
            bool? disableUseAsDirectQuerySource,

            bool? disableUseAsImportedSource)
        {
            DisableUseAsDirectQuerySource = disableUseAsDirectQuerySource;
            DisableUseAsImportedSource = disableUseAsImportedSource;
        }
    }
}