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
    /// A collection of settings that apply to spaces of Amazon SageMaker Studio. These settings are specified when the CreateSpace API is called.
    /// </summary>
    [OutputType]
    public sealed class SpaceSettings
    {
        public readonly Pulumi.AwsNative.SageMaker.SpaceAppType? AppType;
        /// <summary>
        /// The CodeEditor app settings.
        /// </summary>
        public readonly Outputs.SpaceCodeEditorAppSettings? CodeEditorAppSettings;
        public readonly ImmutableArray<Outputs.SpaceCustomFileSystem> CustomFileSystems;
        /// <summary>
        /// The JupyterLab app settings.
        /// </summary>
        public readonly Outputs.SpaceJupyterLabAppSettings? JupyterLabAppSettings;
        /// <summary>
        /// The Jupyter server's app settings.
        /// </summary>
        public readonly Outputs.SpaceJupyterServerAppSettings? JupyterServerAppSettings;
        /// <summary>
        /// The kernel gateway app settings.
        /// </summary>
        public readonly Outputs.SpaceKernelGatewayAppSettings? KernelGatewayAppSettings;
        /// <summary>
        /// Default storage settings for a space.
        /// </summary>
        public readonly Outputs.SpaceStorageSettings? SpaceStorageSettings;

        [OutputConstructor]
        private SpaceSettings(
            Pulumi.AwsNative.SageMaker.SpaceAppType? appType,

            Outputs.SpaceCodeEditorAppSettings? codeEditorAppSettings,

            ImmutableArray<Outputs.SpaceCustomFileSystem> customFileSystems,

            Outputs.SpaceJupyterLabAppSettings? jupyterLabAppSettings,

            Outputs.SpaceJupyterServerAppSettings? jupyterServerAppSettings,

            Outputs.SpaceKernelGatewayAppSettings? kernelGatewayAppSettings,

            Outputs.SpaceStorageSettings? spaceStorageSettings)
        {
            AppType = appType;
            CodeEditorAppSettings = codeEditorAppSettings;
            CustomFileSystems = customFileSystems;
            JupyterLabAppSettings = jupyterLabAppSettings;
            JupyterServerAppSettings = jupyterServerAppSettings;
            KernelGatewayAppSettings = kernelGatewayAppSettings;
            SpaceStorageSettings = spaceStorageSettings;
        }
    }
}
