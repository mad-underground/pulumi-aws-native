// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Greengrass.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-functiondefinitionversion.html
    /// </summary>
    public sealed class FunctionDefinitionFunctionDefinitionVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-functiondefinitionversion.html#cfn-greengrass-functiondefinition-functiondefinitionversion-defaultconfig
        /// </summary>
        [Input("DefaultConfig")]
        public Input<Inputs.FunctionDefinitionDefaultConfigArgs>? DefaultConfig { get; set; }

        [Input("Functions", required: true)]
        private InputList<Inputs.FunctionDefinitionFunctionArgs>? _Functions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-functiondefinitionversion.html#cfn-greengrass-functiondefinition-functiondefinitionversion-functions
        /// </summary>
        public InputList<Inputs.FunctionDefinitionFunctionArgs> Functions
        {
            get => _Functions ?? (_Functions = new InputList<Inputs.FunctionDefinitionFunctionArgs>());
            set => _Functions = value;
        }

        public FunctionDefinitionFunctionDefinitionVersionArgs()
        {
        }
    }
}