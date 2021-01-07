// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Lambda.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-imageconfig.html
    /// </summary>
    public sealed class FunctionImageConfigArgs : Pulumi.ResourceArgs
    {
        [Input("Command")]
        private InputList<string>? _Command;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-imageconfig.html#cfn-lambda-function-imageconfig-command
        /// </summary>
        public InputList<string> Command
        {
            get => _Command ?? (_Command = new InputList<string>());
            set => _Command = value;
        }

        [Input("EntryPoint")]
        private InputList<string>? _EntryPoint;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-imageconfig.html#cfn-lambda-function-imageconfig-entrypoint
        /// </summary>
        public InputList<string> EntryPoint
        {
            get => _EntryPoint ?? (_EntryPoint = new InputList<string>());
            set => _EntryPoint = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-imageconfig.html#cfn-lambda-function-imageconfig-workingdirectory
        /// </summary>
        [Input("WorkingDirectory")]
        public Input<string>? WorkingDirectory { get; set; }

        public FunctionImageConfigArgs()
        {
        }
    }
}