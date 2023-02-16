// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// &lt;p&gt;The source analysis of the template.&lt;/p&gt;
    /// </summary>
    public sealed class TemplateSourceAnalysisArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the resource.&lt;/p&gt;
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        [Input("dataSetReferences", required: true)]
        private InputList<Inputs.TemplateDataSetReferenceArgs>? _dataSetReferences;

        /// <summary>
        /// &lt;p&gt;A structure containing information about the dataset references used as placeholders
        ///             in the template.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.TemplateDataSetReferenceArgs> DataSetReferences
        {
            get => _dataSetReferences ?? (_dataSetReferences = new InputList<Inputs.TemplateDataSetReferenceArgs>());
            set => _dataSetReferences = value;
        }

        public TemplateSourceAnalysisArgs()
        {
        }
        public static new TemplateSourceAnalysisArgs Empty => new TemplateSourceAnalysisArgs();
    }
}
