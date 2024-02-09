// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// The configuration for the file system and kernels in a SageMaker image running as a JupyterLab app.
    /// </summary>
    [OutputType]
    public sealed class AppImageConfigJupyterLabAppImageConfig
    {
        /// <summary>
        /// The container configuration for a SageMaker image.
        /// </summary>
        public readonly Outputs.AppImageConfigContainerConfig? ContainerConfig;

        [OutputConstructor]
        private AppImageConfigJupyterLabAppImageConfig(Outputs.AppImageConfigContainerConfig? containerConfig)
        {
            ContainerConfig = containerConfig;
        }
    }
}
