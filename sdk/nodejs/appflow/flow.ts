// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html
 */
export class Flow extends pulumi.CustomResource {
    /**
     * Get an existing Flow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Flow {
        return new Flow(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:AppFlow:Flow';

    /**
     * Returns true if the given object is an instance of Flow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Flow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Flow.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-destinationflowconfiglist
     */
    public readonly destinationFlowConfigList!: pulumi.Output<outputs.AppFlow.FlowDestinationFlowConfig[]>;
    public /*out*/ readonly flowArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-flowname
     */
    public readonly flowName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-kmsarn
     */
    public readonly kMSArn!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-sourceflowconfig
     */
    public readonly sourceFlowConfig!: pulumi.Output<outputs.AppFlow.FlowSourceFlowConfig>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-tasks
     */
    public readonly tasks!: pulumi.Output<outputs.AppFlow.FlowTask[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-triggerconfig
     */
    public readonly triggerConfig!: pulumi.Output<outputs.AppFlow.FlowTriggerConfig>;

    /**
     * Create a Flow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlowArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.destinationFlowConfigList === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationFlowConfigList'");
            }
            if ((!args || args.flowName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flowName'");
            }
            if ((!args || args.sourceFlowConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceFlowConfig'");
            }
            if ((!args || args.tasks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tasks'");
            }
            if ((!args || args.triggerConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'triggerConfig'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["destinationFlowConfigList"] = args ? args.destinationFlowConfigList : undefined;
            inputs["flowName"] = args ? args.flowName : undefined;
            inputs["kMSArn"] = args ? args.kMSArn : undefined;
            inputs["sourceFlowConfig"] = args ? args.sourceFlowConfig : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tasks"] = args ? args.tasks : undefined;
            inputs["triggerConfig"] = args ? args.triggerConfig : undefined;
            inputs["flowArn"] = undefined /*out*/;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["destinationFlowConfigList"] = undefined /*out*/;
            inputs["flowArn"] = undefined /*out*/;
            inputs["flowName"] = undefined /*out*/;
            inputs["kMSArn"] = undefined /*out*/;
            inputs["sourceFlowConfig"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["tasks"] = undefined /*out*/;
            inputs["triggerConfig"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Flow.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Flow resource.
 */
export interface FlowArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-description
     */
    description?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-destinationflowconfiglist
     */
    destinationFlowConfigList: pulumi.Input<pulumi.Input<inputs.AppFlow.FlowDestinationFlowConfigArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-flowname
     */
    flowName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-kmsarn
     */
    kMSArn?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-sourceflowconfig
     */
    sourceFlowConfig: pulumi.Input<inputs.AppFlow.FlowSourceFlowConfigArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-tags
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-tasks
     */
    tasks: pulumi.Input<pulumi.Input<inputs.AppFlow.FlowTaskArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appflow-flow.html#cfn-appflow-flow-triggerconfig
     */
    triggerConfig: pulumi.Input<inputs.AppFlow.FlowTriggerConfigArgs>;
}
