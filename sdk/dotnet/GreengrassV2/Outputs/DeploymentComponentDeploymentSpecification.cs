// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2.Outputs
{

    [OutputType]
    public sealed class DeploymentComponentDeploymentSpecification
    {
        public readonly string? ComponentVersion;
        public readonly Outputs.DeploymentComponentConfigurationUpdate? ConfigurationUpdate;
        public readonly Outputs.DeploymentComponentRunWith? RunWith;

        [OutputConstructor]
        private DeploymentComponentDeploymentSpecification(
            string? componentVersion,

            Outputs.DeploymentComponentConfigurationUpdate? configurationUpdate,

            Outputs.DeploymentComponentRunWith? runWith)
        {
            ComponentVersion = componentVersion;
            ConfigurationUpdate = configurationUpdate;
            RunWith = runWith;
        }
    }
}
