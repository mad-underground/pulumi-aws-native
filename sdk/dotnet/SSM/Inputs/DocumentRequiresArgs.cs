// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSM.Inputs
{

    public sealed class DocumentRequiresArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the required SSM document. The name can be an Amazon Resource Name (ARN).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The document version required by the current document.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public DocumentRequiresArgs()
        {
        }
        public static new DocumentRequiresArgs Empty => new DocumentRequiresArgs();
    }
}
