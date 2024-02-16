// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::ImageBuilder::ContainerRecipe
 */
export class ContainerRecipe extends pulumi.CustomResource {
    /**
     * Get an existing ContainerRecipe resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ContainerRecipe {
        return new ContainerRecipe(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:imagebuilder:ContainerRecipe';

    /**
     * Returns true if the given object is an instance of ContainerRecipe.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerRecipe {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerRecipe.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the container recipe.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Components for build and test that are included in the container recipe.
     */
    public readonly components!: pulumi.Output<outputs.imagebuilder.ContainerRecipeComponentConfiguration[] | undefined>;
    /**
     * Specifies the type of container, such as Docker.
     */
    public readonly containerType!: pulumi.Output<enums.imagebuilder.ContainerRecipeContainerType | undefined>;
    /**
     * The description of the container recipe.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside. The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.
     */
    public readonly dockerfileTemplateData!: pulumi.Output<string | undefined>;
    /**
     * The S3 URI for the Dockerfile that will be used to build your container image.
     */
    public readonly dockerfileTemplateUri!: pulumi.Output<string | undefined>;
    /**
     * Specifies the operating system version for the source image.
     */
    public readonly imageOsVersionOverride!: pulumi.Output<string | undefined>;
    /**
     * A group of options that can be used to configure an instance for building and testing container images.
     */
    public readonly instanceConfiguration!: pulumi.Output<outputs.imagebuilder.ContainerRecipeInstanceConfiguration | undefined>;
    /**
     * Identifies which KMS key is used to encrypt the container image.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * The name of the container recipe.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The source image for the container recipe.
     */
    public readonly parentImage!: pulumi.Output<string | undefined>;
    /**
     * Specifies the operating system platform when you use a custom source image.
     */
    public readonly platformOverride!: pulumi.Output<enums.imagebuilder.ContainerRecipePlatformOverride | undefined>;
    /**
     * Tags that are attached to the container recipe.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The destination repository for the container image.
     */
    public readonly targetRepository!: pulumi.Output<outputs.imagebuilder.ContainerRecipeTargetContainerRepository | undefined>;
    /**
     * The semantic version of the container recipe (<major>.<minor>.<patch>).
     */
    public readonly version!: pulumi.Output<string | undefined>;
    /**
     * The working directory to be used during build and test workflows.
     */
    public readonly workingDirectory!: pulumi.Output<string | undefined>;

    /**
     * Create a ContainerRecipe resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ContainerRecipeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["components"] = args ? args.components : undefined;
            resourceInputs["containerType"] = args ? args.containerType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dockerfileTemplateData"] = args ? args.dockerfileTemplateData : undefined;
            resourceInputs["dockerfileTemplateUri"] = args ? args.dockerfileTemplateUri : undefined;
            resourceInputs["imageOsVersionOverride"] = args ? args.imageOsVersionOverride : undefined;
            resourceInputs["instanceConfiguration"] = args ? args.instanceConfiguration : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentImage"] = args ? args.parentImage : undefined;
            resourceInputs["platformOverride"] = args ? args.platformOverride : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targetRepository"] = args ? args.targetRepository : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["workingDirectory"] = args ? args.workingDirectory : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["components"] = undefined /*out*/;
            resourceInputs["containerType"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["dockerfileTemplateData"] = undefined /*out*/;
            resourceInputs["dockerfileTemplateUri"] = undefined /*out*/;
            resourceInputs["imageOsVersionOverride"] = undefined /*out*/;
            resourceInputs["instanceConfiguration"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parentImage"] = undefined /*out*/;
            resourceInputs["platformOverride"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["targetRepository"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
            resourceInputs["workingDirectory"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["components[*]", "containerType", "description", "dockerfileTemplateData", "dockerfileTemplateUri", "imageOsVersionOverride", "instanceConfiguration", "kmsKeyId", "name", "parentImage", "platformOverride", "tags.*", "targetRepository", "version", "workingDirectory"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ContainerRecipe.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ContainerRecipe resource.
 */
export interface ContainerRecipeArgs {
    /**
     * Components for build and test that are included in the container recipe.
     */
    components?: pulumi.Input<pulumi.Input<inputs.imagebuilder.ContainerRecipeComponentConfigurationArgs>[]>;
    /**
     * Specifies the type of container, such as Docker.
     */
    containerType?: pulumi.Input<enums.imagebuilder.ContainerRecipeContainerType>;
    /**
     * The description of the container recipe.
     */
    description?: pulumi.Input<string>;
    /**
     * Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside. The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.
     */
    dockerfileTemplateData?: pulumi.Input<string>;
    /**
     * The S3 URI for the Dockerfile that will be used to build your container image.
     */
    dockerfileTemplateUri?: pulumi.Input<string>;
    /**
     * Specifies the operating system version for the source image.
     */
    imageOsVersionOverride?: pulumi.Input<string>;
    /**
     * A group of options that can be used to configure an instance for building and testing container images.
     */
    instanceConfiguration?: pulumi.Input<inputs.imagebuilder.ContainerRecipeInstanceConfigurationArgs>;
    /**
     * Identifies which KMS key is used to encrypt the container image.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The name of the container recipe.
     */
    name?: pulumi.Input<string>;
    /**
     * The source image for the container recipe.
     */
    parentImage?: pulumi.Input<string>;
    /**
     * Specifies the operating system platform when you use a custom source image.
     */
    platformOverride?: pulumi.Input<enums.imagebuilder.ContainerRecipePlatformOverride>;
    /**
     * Tags that are attached to the container recipe.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The destination repository for the container image.
     */
    targetRepository?: pulumi.Input<inputs.imagebuilder.ContainerRecipeTargetContainerRepositoryArgs>;
    /**
     * The semantic version of the container recipe (<major>.<minor>.<patch>).
     */
    version?: pulumi.Input<string>;
    /**
     * The working directory to be used during build and test workflows.
     */
    workingDirectory?: pulumi.Input<string>;
}
