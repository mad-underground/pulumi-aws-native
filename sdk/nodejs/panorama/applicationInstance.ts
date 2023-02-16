// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Schema for ApplicationInstance CloudFormation Resource
 */
export class ApplicationInstance extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApplicationInstance {
        return new ApplicationInstance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:panorama:ApplicationInstance';

    /**
     * Returns true if the given object is an instance of ApplicationInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationInstance.__pulumiType;
    }

    public /*out*/ readonly applicationInstanceId!: pulumi.Output<string>;
    public readonly applicationInstanceIdToReplace!: pulumi.Output<string | undefined>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly createdTime!: pulumi.Output<number>;
    public readonly defaultRuntimeContextDevice!: pulumi.Output<string>;
    public /*out*/ readonly defaultRuntimeContextDeviceName!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly deviceId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly healthStatus!: pulumi.Output<enums.panorama.ApplicationInstanceHealthStatus>;
    public /*out*/ readonly lastUpdatedTime!: pulumi.Output<number>;
    public readonly manifestOverridesPayload!: pulumi.Output<outputs.panorama.ApplicationInstanceManifestOverridesPayload | undefined>;
    public readonly manifestPayload!: pulumi.Output<outputs.panorama.ApplicationInstanceManifestPayload>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly runtimeRoleArn!: pulumi.Output<string | undefined>;
    public /*out*/ readonly status!: pulumi.Output<enums.panorama.ApplicationInstanceStatus>;
    public /*out*/ readonly statusDescription!: pulumi.Output<string>;
    public readonly statusFilter!: pulumi.Output<enums.panorama.ApplicationInstanceStatusFilter | undefined>;
    public readonly tags!: pulumi.Output<outputs.panorama.ApplicationInstanceTag[] | undefined>;

    /**
     * Create a ApplicationInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationInstanceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.defaultRuntimeContextDevice === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultRuntimeContextDevice'");
            }
            if ((!args || args.manifestPayload === undefined) && !opts.urn) {
                throw new Error("Missing required property 'manifestPayload'");
            }
            resourceInputs["applicationInstanceIdToReplace"] = args ? args.applicationInstanceIdToReplace : undefined;
            resourceInputs["defaultRuntimeContextDevice"] = args ? args.defaultRuntimeContextDevice : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceId"] = args ? args.deviceId : undefined;
            resourceInputs["manifestOverridesPayload"] = args ? args.manifestOverridesPayload : undefined;
            resourceInputs["manifestPayload"] = args ? args.manifestPayload : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["runtimeRoleArn"] = args ? args.runtimeRoleArn : undefined;
            resourceInputs["statusFilter"] = args ? args.statusFilter : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["applicationInstanceId"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["defaultRuntimeContextDeviceName"] = undefined /*out*/;
            resourceInputs["healthStatus"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusDescription"] = undefined /*out*/;
        } else {
            resourceInputs["applicationInstanceId"] = undefined /*out*/;
            resourceInputs["applicationInstanceIdToReplace"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["defaultRuntimeContextDevice"] = undefined /*out*/;
            resourceInputs["defaultRuntimeContextDeviceName"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["deviceId"] = undefined /*out*/;
            resourceInputs["healthStatus"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
            resourceInputs["manifestOverridesPayload"] = undefined /*out*/;
            resourceInputs["manifestPayload"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["runtimeRoleArn"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusDescription"] = undefined /*out*/;
            resourceInputs["statusFilter"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApplicationInstance resource.
 */
export interface ApplicationInstanceArgs {
    applicationInstanceIdToReplace?: pulumi.Input<string>;
    defaultRuntimeContextDevice: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    deviceId?: pulumi.Input<string>;
    manifestOverridesPayload?: pulumi.Input<inputs.panorama.ApplicationInstanceManifestOverridesPayloadArgs>;
    manifestPayload: pulumi.Input<inputs.panorama.ApplicationInstanceManifestPayloadArgs>;
    name?: pulumi.Input<string>;
    runtimeRoleArn?: pulumi.Input<string>;
    statusFilter?: pulumi.Input<enums.panorama.ApplicationInstanceStatusFilter>;
    tags?: pulumi.Input<pulumi.Input<inputs.panorama.ApplicationInstanceTagArgs>[]>;
}
