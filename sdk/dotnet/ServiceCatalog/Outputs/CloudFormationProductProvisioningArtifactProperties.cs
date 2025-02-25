// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalog.Outputs
{

    [OutputType]
    public sealed class CloudFormationProductProvisioningArtifactProperties
    {
        public readonly string? Description;
        public readonly bool? DisableTemplateValidation;
        public readonly object Info;
        public readonly string? Name;
        public readonly string? Type;

        [OutputConstructor]
        private CloudFormationProductProvisioningArtifactProperties(
            string? description,

            bool? disableTemplateValidation,

            object info,

            string? name,

            string? type)
        {
            Description = description;
            DisableTemplateValidation = disableTemplateValidation;
            Info = info;
            Name = name;
            Type = type;
        }
    }
}
