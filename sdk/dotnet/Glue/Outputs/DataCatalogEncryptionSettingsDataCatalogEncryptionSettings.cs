// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Glue.Outputs
{

    [OutputType]
    public sealed class DataCatalogEncryptionSettingsDataCatalogEncryptionSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-datacatalogencryptionsettings.html#cfn-glue-datacatalogencryptionsettings-datacatalogencryptionsettings-connectionpasswordencryption
        /// </summary>
        public readonly Outputs.DataCatalogEncryptionSettingsConnectionPasswordEncryption? ConnectionPasswordEncryption;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-datacatalogencryptionsettings.html#cfn-glue-datacatalogencryptionsettings-datacatalogencryptionsettings-encryptionatrest
        /// </summary>
        public readonly Outputs.DataCatalogEncryptionSettingsEncryptionAtRest? EncryptionAtRest;

        [OutputConstructor]
        private DataCatalogEncryptionSettingsDataCatalogEncryptionSettings(
            Outputs.DataCatalogEncryptionSettingsConnectionPasswordEncryption? ConnectionPasswordEncryption,

            Outputs.DataCatalogEncryptionSettingsEncryptionAtRest? EncryptionAtRest)
        {
            this.ConnectionPasswordEncryption = ConnectionPasswordEncryption;
            this.EncryptionAtRest = EncryptionAtRest;
        }
    }
}