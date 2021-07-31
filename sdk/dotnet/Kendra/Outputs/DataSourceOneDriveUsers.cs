// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveusers.html
    /// </summary>
    [OutputType]
    public sealed class DataSourceOneDriveUsers
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveusers.html#cfn-kendra-datasource-onedriveusers-onedriveuserlist
        /// </summary>
        public readonly Outputs.DataSourceOneDriveUserList? OneDriveUserList;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveusers.html#cfn-kendra-datasource-onedriveusers-onedriveusers3path
        /// </summary>
        public readonly Outputs.DataSourceS3Path? OneDriveUserS3Path;

        [OutputConstructor]
        private DataSourceOneDriveUsers(
            Outputs.DataSourceOneDriveUserList? oneDriveUserList,

            Outputs.DataSourceS3Path? oneDriveUserS3Path)
        {
            OneDriveUserList = oneDriveUserList;
            OneDriveUserS3Path = oneDriveUserS3Path;
        }
    }
}
