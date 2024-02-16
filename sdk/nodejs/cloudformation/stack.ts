// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::CloudFormation::Stack resource nests a stack as a resource in a top-level template.
 */
export class Stack extends pulumi.CustomResource {
    /**
     * Get an existing Stack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Stack {
        return new Stack(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudformation:Stack';

    /**
     * Returns true if the given object is an instance of Stack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Stack {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Stack.__pulumiType;
    }

    public readonly capabilities!: pulumi.Output<enums.cloudformation.StackCapabilitiesItem[] | undefined>;
    public /*out*/ readonly changeSetId!: pulumi.Output<string>;
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly disableRollback!: pulumi.Output<boolean | undefined>;
    public readonly enableTerminationProtection!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly lastUpdateTime!: pulumi.Output<string>;
    public readonly notificationArns!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly outputs!: pulumi.Output<outputs.cloudformation.StackOutput[]>;
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly parentId!: pulumi.Output<string>;
    public readonly roleArn!: pulumi.Output<string | undefined>;
    public /*out*/ readonly rootId!: pulumi.Output<string>;
    public /*out*/ readonly stackId!: pulumi.Output<string>;
    public readonly stackName!: pulumi.Output<string>;
    public readonly stackPolicyBody!: pulumi.Output<any | undefined>;
    public readonly stackPolicyUrl!: pulumi.Output<string | undefined>;
    public /*out*/ readonly stackStatus!: pulumi.Output<enums.cloudformation.StackStatus>;
    public readonly stackStatusReason!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.cloudformation.StackTag[] | undefined>;
    public readonly templateBody!: pulumi.Output<any | undefined>;
    public readonly templateUrl!: pulumi.Output<string | undefined>;
    public readonly timeoutInMinutes!: pulumi.Output<number | undefined>;

    /**
     * Create a Stack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StackArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["capabilities"] = args ? args.capabilities : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableRollback"] = args ? args.disableRollback : undefined;
            resourceInputs["enableTerminationProtection"] = args ? args.enableTerminationProtection : undefined;
            resourceInputs["notificationArns"] = args ? args.notificationArns : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["stackName"] = args ? args.stackName : undefined;
            resourceInputs["stackPolicyBody"] = args ? args.stackPolicyBody : undefined;
            resourceInputs["stackPolicyUrl"] = args ? args.stackPolicyUrl : undefined;
            resourceInputs["stackStatusReason"] = args ? args.stackStatusReason : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["templateBody"] = args ? args.templateBody : undefined;
            resourceInputs["templateUrl"] = args ? args.templateUrl : undefined;
            resourceInputs["timeoutInMinutes"] = args ? args.timeoutInMinutes : undefined;
            resourceInputs["changeSetId"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["lastUpdateTime"] = undefined /*out*/;
            resourceInputs["outputs"] = undefined /*out*/;
            resourceInputs["parentId"] = undefined /*out*/;
            resourceInputs["rootId"] = undefined /*out*/;
            resourceInputs["stackId"] = undefined /*out*/;
            resourceInputs["stackStatus"] = undefined /*out*/;
        } else {
            resourceInputs["capabilities"] = undefined /*out*/;
            resourceInputs["changeSetId"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["disableRollback"] = undefined /*out*/;
            resourceInputs["enableTerminationProtection"] = undefined /*out*/;
            resourceInputs["lastUpdateTime"] = undefined /*out*/;
            resourceInputs["notificationArns"] = undefined /*out*/;
            resourceInputs["outputs"] = undefined /*out*/;
            resourceInputs["parameters"] = undefined /*out*/;
            resourceInputs["parentId"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["rootId"] = undefined /*out*/;
            resourceInputs["stackId"] = undefined /*out*/;
            resourceInputs["stackName"] = undefined /*out*/;
            resourceInputs["stackPolicyBody"] = undefined /*out*/;
            resourceInputs["stackPolicyUrl"] = undefined /*out*/;
            resourceInputs["stackStatus"] = undefined /*out*/;
            resourceInputs["stackStatusReason"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["templateBody"] = undefined /*out*/;
            resourceInputs["templateUrl"] = undefined /*out*/;
            resourceInputs["timeoutInMinutes"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["stackName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Stack.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Stack resource.
 */
export interface StackArgs {
    capabilities?: pulumi.Input<pulumi.Input<enums.cloudformation.StackCapabilitiesItem>[]>;
    description?: pulumi.Input<string>;
    disableRollback?: pulumi.Input<boolean>;
    enableTerminationProtection?: pulumi.Input<boolean>;
    notificationArns?: pulumi.Input<pulumi.Input<string>[]>;
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    roleArn?: pulumi.Input<string>;
    stackName?: pulumi.Input<string>;
    stackPolicyBody?: any;
    stackPolicyUrl?: pulumi.Input<string>;
    stackStatusReason?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.cloudformation.StackTagArgs>[]>;
    templateBody?: any;
    templateUrl?: pulumi.Input<string>;
    timeoutInMinutes?: pulumi.Input<number>;
}
