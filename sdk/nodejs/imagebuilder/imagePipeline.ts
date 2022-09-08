// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::ImageBuilder::ImagePipeline
 */
export class ImagePipeline extends pulumi.CustomResource {
    /**
     * Get an existing ImagePipeline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ImagePipeline {
        return new ImagePipeline(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:imagebuilder:ImagePipeline';

    /**
     * Returns true if the given object is an instance of ImagePipeline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ImagePipeline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ImagePipeline.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the image pipeline.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.
     */
    public readonly containerRecipeArn!: pulumi.Output<string | undefined>;
    /**
     * The description of the image pipeline.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.
     */
    public readonly distributionConfigurationArn!: pulumi.Output<string | undefined>;
    /**
     * Collects additional information about the image being created, including the operating system (OS) version and package list.
     */
    public readonly enhancedImageMetadataEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.
     */
    public readonly imageRecipeArn!: pulumi.Output<string | undefined>;
    /**
     * The image tests configuration of the image pipeline.
     */
    public readonly imageTestsConfiguration!: pulumi.Output<outputs.imagebuilder.ImagePipelineImageTestsConfiguration | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.
     */
    public readonly infrastructureConfigurationArn!: pulumi.Output<string | undefined>;
    /**
     * The name of the image pipeline.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The schedule of the image pipeline.
     */
    public readonly schedule!: pulumi.Output<outputs.imagebuilder.ImagePipelineSchedule | undefined>;
    /**
     * The status of the image pipeline.
     */
    public readonly status!: pulumi.Output<enums.imagebuilder.ImagePipelineStatus | undefined>;
    /**
     * The tags of this image pipeline.
     */
    public readonly tags!: pulumi.Output<any | undefined>;

    /**
     * Create a ImagePipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ImagePipelineArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["containerRecipeArn"] = args ? args.containerRecipeArn : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["distributionConfigurationArn"] = args ? args.distributionConfigurationArn : undefined;
            resourceInputs["enhancedImageMetadataEnabled"] = args ? args.enhancedImageMetadataEnabled : undefined;
            resourceInputs["imageRecipeArn"] = args ? args.imageRecipeArn : undefined;
            resourceInputs["imageTestsConfiguration"] = args ? args.imageTestsConfiguration : undefined;
            resourceInputs["infrastructureConfigurationArn"] = args ? args.infrastructureConfigurationArn : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["containerRecipeArn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["distributionConfigurationArn"] = undefined /*out*/;
            resourceInputs["enhancedImageMetadataEnabled"] = undefined /*out*/;
            resourceInputs["imageRecipeArn"] = undefined /*out*/;
            resourceInputs["imageTestsConfiguration"] = undefined /*out*/;
            resourceInputs["infrastructureConfigurationArn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["schedule"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ImagePipeline.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ImagePipeline resource.
 */
export interface ImagePipelineArgs {
    /**
     * The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.
     */
    containerRecipeArn?: pulumi.Input<string>;
    /**
     * The description of the image pipeline.
     */
    description?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.
     */
    distributionConfigurationArn?: pulumi.Input<string>;
    /**
     * Collects additional information about the image being created, including the operating system (OS) version and package list.
     */
    enhancedImageMetadataEnabled?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.
     */
    imageRecipeArn?: pulumi.Input<string>;
    /**
     * The image tests configuration of the image pipeline.
     */
    imageTestsConfiguration?: pulumi.Input<inputs.imagebuilder.ImagePipelineImageTestsConfigurationArgs>;
    /**
     * The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.
     */
    infrastructureConfigurationArn?: pulumi.Input<string>;
    /**
     * The name of the image pipeline.
     */
    name?: pulumi.Input<string>;
    /**
     * The schedule of the image pipeline.
     */
    schedule?: pulumi.Input<inputs.imagebuilder.ImagePipelineScheduleArgs>;
    /**
     * The status of the image pipeline.
     */
    status?: pulumi.Input<enums.imagebuilder.ImagePipelineStatus>;
    /**
     * The tags of this image pipeline.
     */
    tags?: any;
}
