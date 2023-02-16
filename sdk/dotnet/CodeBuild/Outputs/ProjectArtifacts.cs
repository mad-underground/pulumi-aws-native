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
    public sealed class ProjectArtifacts
    {
        public readonly string? ArtifactIdentifier;
        public readonly bool? EncryptionDisabled;
        public readonly string? Location;
        public readonly string? Name;
        public readonly string? NamespaceType;
        public readonly bool? OverrideArtifactName;
        public readonly string? Packaging;
        public readonly string? Path;
        public readonly string Type;

        [OutputConstructor]
        private ProjectArtifacts(
            string? artifactIdentifier,

            bool? encryptionDisabled,

            string? location,

            string? name,

            string? namespaceType,

            bool? overrideArtifactName,

            string? packaging,

            string? path,

            string type)
        {
            ArtifactIdentifier = artifactIdentifier;
            EncryptionDisabled = encryptionDisabled;
            Location = location;
            Name = name;
            NamespaceType = namespaceType;
            OverrideArtifactName = overrideArtifactName;
            Packaging = packaging;
            Path = path;
            Type = type;
        }
    }
}
