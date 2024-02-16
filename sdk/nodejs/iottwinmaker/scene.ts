// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::IoTTwinMaker::Scene
 */
export class Scene extends pulumi.CustomResource {
    /**
     * Get an existing Scene resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Scene {
        return new Scene(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:iottwinmaker:Scene';

    /**
     * Returns true if the given object is an instance of Scene.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Scene {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Scene.__pulumiType;
    }

    /**
     * The ARN of the scene.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A list of capabilities that the scene uses to render.
     */
    public readonly capabilities!: pulumi.Output<string[] | undefined>;
    /**
     * The relative path that specifies the location of the content definition file.
     */
    public readonly contentLocation!: pulumi.Output<string>;
    /**
     * The date and time when the scene was created.
     */
    public /*out*/ readonly creationDateTime!: pulumi.Output<string>;
    /**
     * The description of the scene.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A key-value pair of generated scene metadata for the scene.
     */
    public /*out*/ readonly generatedSceneMetadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * The ID of the scene.
     */
    public readonly sceneId!: pulumi.Output<string>;
    /**
     * A key-value pair of scene metadata for the scene.
     */
    public readonly sceneMetadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A key-value pair to associate with a resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The date and time of the current update.
     */
    public /*out*/ readonly updateDateTime!: pulumi.Output<string>;
    /**
     * The ID of the scene.
     */
    public readonly workspaceId!: pulumi.Output<string>;

    /**
     * Create a Scene resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SceneArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.contentLocation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contentLocation'");
            }
            if ((!args || args.sceneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sceneId'");
            }
            if ((!args || args.workspaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceId'");
            }
            resourceInputs["capabilities"] = args ? args.capabilities : undefined;
            resourceInputs["contentLocation"] = args ? args.contentLocation : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["sceneId"] = args ? args.sceneId : undefined;
            resourceInputs["sceneMetadata"] = args ? args.sceneMetadata : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["workspaceId"] = args ? args.workspaceId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["creationDateTime"] = undefined /*out*/;
            resourceInputs["generatedSceneMetadata"] = undefined /*out*/;
            resourceInputs["updateDateTime"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["capabilities"] = undefined /*out*/;
            resourceInputs["contentLocation"] = undefined /*out*/;
            resourceInputs["creationDateTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["generatedSceneMetadata"] = undefined /*out*/;
            resourceInputs["sceneId"] = undefined /*out*/;
            resourceInputs["sceneMetadata"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["updateDateTime"] = undefined /*out*/;
            resourceInputs["workspaceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["sceneId", "workspaceId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Scene.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Scene resource.
 */
export interface SceneArgs {
    /**
     * A list of capabilities that the scene uses to render.
     */
    capabilities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The relative path that specifies the location of the content definition file.
     */
    contentLocation: pulumi.Input<string>;
    /**
     * The description of the scene.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the scene.
     */
    sceneId: pulumi.Input<string>;
    /**
     * A key-value pair of scene metadata for the scene.
     */
    sceneMetadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A key-value pair to associate with a resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the scene.
     */
    workspaceId: pulumi.Input<string>;
}
