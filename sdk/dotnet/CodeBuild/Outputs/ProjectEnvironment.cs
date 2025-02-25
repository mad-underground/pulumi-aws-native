// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeBuild.Outputs
{

    [OutputType]
    public sealed class ProjectEnvironment
    {
        public readonly string? Certificate;
        public readonly string ComputeType;
        public readonly ImmutableArray<Outputs.ProjectEnvironmentVariable> EnvironmentVariables;
        public readonly Outputs.ProjectFleet? Fleet;
        public readonly string Image;
        public readonly string? ImagePullCredentialsType;
        public readonly bool? PrivilegedMode;
        public readonly Outputs.ProjectRegistryCredential? RegistryCredential;
        public readonly string Type;

        [OutputConstructor]
        private ProjectEnvironment(
            string? certificate,

            string computeType,

            ImmutableArray<Outputs.ProjectEnvironmentVariable> environmentVariables,

            Outputs.ProjectFleet? fleet,

            string image,

            string? imagePullCredentialsType,

            bool? privilegedMode,

            Outputs.ProjectRegistryCredential? registryCredential,

            string type)
        {
            Certificate = certificate;
            ComputeType = computeType;
            EnvironmentVariables = environmentVariables;
            Fleet = fleet;
            Image = image;
            ImagePullCredentialsType = imagePullCredentialsType;
            PrivilegedMode = privilegedMode;
            RegistryCredential = registryCredential;
            Type = type;
        }
    }
}
