// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.MSK.Outputs
{

    [OutputType]
    public sealed class ClusterStorageInfo
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-storageinfo.html#cfn-msk-cluster-storageinfo-ebsstorageinfo
        /// </summary>
        public readonly Outputs.ClusterEBSStorageInfo? EBSStorageInfo;

        [OutputConstructor]
        private ClusterStorageInfo(Outputs.ClusterEBSStorageInfo? EBSStorageInfo)
        {
            this.EBSStorageInfo = EBSStorageInfo;
        }
    }
}