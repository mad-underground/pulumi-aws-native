// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class ComponentPredicateArgs : global::Pulumi.ResourceArgs
    {
        [Input("and")]
        private InputList<Inputs.ComponentPredicateArgs>? _and;
        public InputList<Inputs.ComponentPredicateArgs> And
        {
            get => _and ?? (_and = new InputList<Inputs.ComponentPredicateArgs>());
            set => _and = value;
        }

        [Input("field")]
        public Input<string>? Field { get; set; }

        [Input("operand")]
        public Input<string>? Operand { get; set; }

        [Input("operandType")]
        public Input<string>? OperandType { get; set; }

        [Input("operator")]
        public Input<string>? Operator { get; set; }

        [Input("or")]
        private InputList<Inputs.ComponentPredicateArgs>? _or;
        public InputList<Inputs.ComponentPredicateArgs> Or
        {
            get => _or ?? (_or = new InputList<Inputs.ComponentPredicateArgs>());
            set => _or = value;
        }

        public ComponentPredicateArgs()
        {
        }
        public static new ComponentPredicateArgs Empty => new ComponentPredicateArgs();
    }
}
