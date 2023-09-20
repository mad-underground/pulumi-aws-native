// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Omics.Inputs
{

    public sealed class AnnotationStoreTsvStoreOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotationType")]
        public Input<Pulumi.AwsNative.Omics.AnnotationStoreAnnotationType>? AnnotationType { get; set; }

        [Input("formatToHeader")]
        public Input<Inputs.AnnotationStoreFormatToHeaderArgs>? FormatToHeader { get; set; }

        [Input("schema")]
        private InputList<Inputs.AnnotationStoreSchemaItemArgs>? _schema;
        public InputList<Inputs.AnnotationStoreSchemaItemArgs> Schema
        {
            get => _schema ?? (_schema = new InputList<Inputs.AnnotationStoreSchemaItemArgs>());
            set => _schema = value;
        }

        public AnnotationStoreTsvStoreOptionsArgs()
        {
        }
        public static new AnnotationStoreTsvStoreOptionsArgs Empty => new AnnotationStoreTsvStoreOptionsArgs();
    }
}