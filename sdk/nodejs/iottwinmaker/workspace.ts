// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::IoTTwinMaker::Workspace
 */
export class Workspace extends pulumi.CustomResource {
    /**
     * Get an existing Workspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Workspace {
        return new Workspace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iottwinmaker:Workspace';

    /**
     * Returns true if the given object is an instance of Workspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workspace.__pulumiType;
    }

    /**
     * The ARN of the workspace.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The date and time when the workspace was created.
     */
    public /*out*/ readonly creationDateTime!: pulumi.Output<string>;
    /**
     * The description of the workspace.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the execution role associated with the workspace.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The ARN of the S3 bucket where resources associated with the workspace are stored.
     */
    public readonly s3Location!: pulumi.Output<string>;
    /**
     * A map of key-value pairs to associate with a resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The date and time of the current update.
     */
    public /*out*/ readonly updateDateTime!: pulumi.Output<string>;
    /**
     * The ID of the workspace.
     */
    public readonly workspaceId!: pulumi.Output<string>;

    /**
     * Create a Workspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkspaceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.s3Location === undefined) && !opts.urn) {
                throw new Error("Missing required property 's3Location'");
            }
            if ((!args || args.workspaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["s3Location"] = args ? args.s3Location : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["workspaceId"] = args ? args.workspaceId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["creationDateTime"] = undefined /*out*/;
            resourceInputs["updateDateTime"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["creationDateTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["role"] = undefined /*out*/;
            resourceInputs["s3Location"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["updateDateTime"] = undefined /*out*/;
            resourceInputs["workspaceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["workspaceId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Workspace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Workspace resource.
 */
export interface WorkspaceArgs {
    /**
     * The description of the workspace.
     */
    description?: pulumi.Input<string>;
    /**
     * The ARN of the execution role associated with the workspace.
     */
    role: pulumi.Input<string>;
    /**
     * The ARN of the S3 bucket where resources associated with the workspace are stored.
     */
    s3Location: pulumi.Input<string>;
    /**
     * A map of key-value pairs to associate with a resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the workspace.
     */
    workspaceId: pulumi.Input<string>;
}
