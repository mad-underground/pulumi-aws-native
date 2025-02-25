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
    /// A collection of settings that apply to users of Amazon SageMaker Studio. These settings are specified when the CreateUserProfile API is called, and as DefaultUserSettings when the CreateDomain API is called.
    /// </summary>
    [OutputType]
    public sealed class UserProfileUserSettings
    {
        public readonly Outputs.UserProfileCodeEditorAppSettings? CodeEditorAppSettings;
        public readonly ImmutableArray<Outputs.UserProfileCustomFileSystemConfig> CustomFileSystemConfigs;
        public readonly Outputs.UserProfileCustomPosixUserConfig? CustomPosixUserConfig;
        /// <summary>
        /// Defines which Amazon SageMaker application users are directed to by default.
        /// </summary>
        public readonly string? DefaultLandingUri;
        /// <summary>
        /// The user profile Amazon Resource Name (ARN).
        /// </summary>
        public readonly string? ExecutionRole;
        public readonly Outputs.UserProfileJupyterLabAppSettings? JupyterLabAppSettings;
        /// <summary>
        /// The Jupyter server's app settings.
        /// </summary>
        public readonly Outputs.UserProfileJupyterServerAppSettings? JupyterServerAppSettings;
        /// <summary>
        /// The kernel gateway app settings.
        /// </summary>
        public readonly Outputs.UserProfileKernelGatewayAppSettings? KernelGatewayAppSettings;
        public readonly Outputs.UserProfileRStudioServerProAppSettings? RStudioServerProAppSettings;
        /// <summary>
        /// The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;
        /// <summary>
        /// The sharing settings.
        /// </summary>
        public readonly Outputs.UserProfileSharingSettings? SharingSettings;
        public readonly Outputs.UserProfileDefaultSpaceStorageSettings? SpaceStorageSettings;
        /// <summary>
        /// Indicates whether the Studio experience is available to users. If not, users cannot access Studio.
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.UserProfileUserSettingsStudioWebPortal? StudioWebPortal;

        [OutputConstructor]
        private UserProfileUserSettings(
            Outputs.UserProfileCodeEditorAppSettings? codeEditorAppSettings,

            ImmutableArray<Outputs.UserProfileCustomFileSystemConfig> customFileSystemConfigs,

            Outputs.UserProfileCustomPosixUserConfig? customPosixUserConfig,

            string? defaultLandingUri,

            string? executionRole,

            Outputs.UserProfileJupyterLabAppSettings? jupyterLabAppSettings,

            Outputs.UserProfileJupyterServerAppSettings? jupyterServerAppSettings,

            Outputs.UserProfileKernelGatewayAppSettings? kernelGatewayAppSettings,

            Outputs.UserProfileRStudioServerProAppSettings? rStudioServerProAppSettings,

            ImmutableArray<string> securityGroups,

            Outputs.UserProfileSharingSettings? sharingSettings,

            Outputs.UserProfileDefaultSpaceStorageSettings? spaceStorageSettings,

            Pulumi.AwsNative.SageMaker.UserProfileUserSettingsStudioWebPortal? studioWebPortal)
        {
            CodeEditorAppSettings = codeEditorAppSettings;
            CustomFileSystemConfigs = customFileSystemConfigs;
            CustomPosixUserConfig = customPosixUserConfig;
            DefaultLandingUri = defaultLandingUri;
            ExecutionRole = executionRole;
            JupyterLabAppSettings = jupyterLabAppSettings;
            JupyterServerAppSettings = jupyterServerAppSettings;
            KernelGatewayAppSettings = kernelGatewayAppSettings;
            RStudioServerProAppSettings = rStudioServerProAppSettings;
            SecurityGroups = securityGroups;
            SharingSettings = sharingSettings;
            SpaceStorageSettings = spaceStorageSettings;
            StudioWebPortal = studioWebPortal;
        }
    }
}
