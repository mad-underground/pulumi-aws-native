// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Glue.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html
    /// </summary>
    public sealed class ConnectionConnectionInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-connectionproperties
        /// </summary>
        [Input("ConnectionProperties")]
        public InputUnion<System.Text.Json.JsonElement, string>? ConnectionProperties { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-connectiontype
        /// </summary>
        [Input("ConnectionType", required: true)]
        public Input<string> ConnectionType { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-description
        /// </summary>
        [Input("Description")]
        public Input<string>? Description { get; set; }

        [Input("MatchCriteria")]
        private InputList<string>? _MatchCriteria;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-matchcriteria
        /// </summary>
        public InputList<string> MatchCriteria
        {
            get => _MatchCriteria ?? (_MatchCriteria = new InputList<string>());
            set => _MatchCriteria = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-name
        /// </summary>
        [Input("Name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-connection-connectioninput.html#cfn-glue-connection-connectioninput-physicalconnectionrequirements
        /// </summary>
        [Input("PhysicalConnectionRequirements")]
        public Input<Inputs.ConnectionPhysicalConnectionRequirementsArgs>? PhysicalConnectionRequirements { get; set; }

        public ConnectionConnectionInputArgs()
        {
        }
    }
}
