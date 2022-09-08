// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::GlobalAccelerator::Listener
 */
export class Listener extends pulumi.CustomResource {
    /**
     * Get an existing Listener resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Listener {
        return new Listener(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:globalaccelerator:Listener';

    /**
     * Returns true if the given object is an instance of Listener.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Listener {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Listener.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the accelerator.
     */
    public readonly acceleratorArn!: pulumi.Output<string>;
    /**
     * Client affinity lets you direct all requests from a user to the same endpoint.
     */
    public readonly clientAffinity!: pulumi.Output<enums.globalaccelerator.ListenerClientAffinity | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the listener.
     */
    public /*out*/ readonly listenerArn!: pulumi.Output<string>;
    public readonly portRanges!: pulumi.Output<outputs.globalaccelerator.ListenerPortRange[]>;
    /**
     * The protocol for the listener.
     */
    public readonly protocol!: pulumi.Output<enums.globalaccelerator.ListenerProtocol>;

    /**
     * Create a Listener resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.acceleratorArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceleratorArn'");
            }
            if ((!args || args.portRanges === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portRanges'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            resourceInputs["acceleratorArn"] = args ? args.acceleratorArn : undefined;
            resourceInputs["clientAffinity"] = args ? args.clientAffinity : undefined;
            resourceInputs["portRanges"] = args ? args.portRanges : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["listenerArn"] = undefined /*out*/;
        } else {
            resourceInputs["acceleratorArn"] = undefined /*out*/;
            resourceInputs["clientAffinity"] = undefined /*out*/;
            resourceInputs["listenerArn"] = undefined /*out*/;
            resourceInputs["portRanges"] = undefined /*out*/;
            resourceInputs["protocol"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Listener.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Listener resource.
 */
export interface ListenerArgs {
    /**
     * The Amazon Resource Name (ARN) of the accelerator.
     */
    acceleratorArn: pulumi.Input<string>;
    /**
     * Client affinity lets you direct all requests from a user to the same endpoint.
     */
    clientAffinity?: pulumi.Input<enums.globalaccelerator.ListenerClientAffinity>;
    portRanges: pulumi.Input<pulumi.Input<inputs.globalaccelerator.ListenerPortRangeArgs>[]>;
    /**
     * The protocol for the listener.
     */
    protocol: pulumi.Input<enums.globalaccelerator.ListenerProtocol>;
}
