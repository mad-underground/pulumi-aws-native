// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class SpaceSharingSettings
    {
        public readonly Pulumi.AwsNative.SageMaker.SpaceSharingSettingsSharingType SharingType;

        [OutputConstructor]
        private SpaceSharingSettings(Pulumi.AwsNative.SageMaker.SpaceSharingSettingsSharingType sharingType)
        {
            SharingType = sharingType;
        }
    }
}
