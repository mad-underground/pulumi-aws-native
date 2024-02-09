// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// The CodeEditor app settings.
    /// </summary>
    public sealed class SpaceCodeEditorAppSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("defaultResourceSpec")]
        public Input<Inputs.SpaceResourceSpecArgs>? DefaultResourceSpec { get; set; }

        public SpaceCodeEditorAppSettingsArgs()
        {
        }
        public static new SpaceCodeEditorAppSettingsArgs Empty => new SpaceCodeEditorAppSettingsArgs();
    }
}
