// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::ImageBuilder::Image
 */
export class Image extends pulumi.CustomResource {
    /**
     * Get an existing Image resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Image {
        return new Image(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:imagebuilder:Image';

    /**
     * Returns true if the given object is an instance of Image.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Image {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Image.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the image.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.
     */
    public readonly containerRecipeArn!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the distribution configuration.
     */
    public readonly distributionConfigurationArn!: pulumi.Output<string | undefined>;
    /**
     * Collects additional information about the image being created, including the operating system (OS) version and package list.
     */
    public readonly enhancedImageMetadataEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The execution role name/ARN for the image build, if provided
     */
    public readonly executionRole!: pulumi.Output<string | undefined>;
    /**
     * The AMI ID of the EC2 AMI in current region.
     */
    public /*out*/ readonly imageId!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.
     */
    public readonly imageRecipeArn!: pulumi.Output<string | undefined>;
    /**
     * Contains settings for vulnerability scans.
     */
    public readonly imageScanningConfiguration!: pulumi.Output<outputs.imagebuilder.ImageScanningConfiguration | undefined>;
    /**
     * The image tests configuration used when creating this image.
     */
    public readonly imageTestsConfiguration!: pulumi.Output<outputs.imagebuilder.ImageTestsConfiguration | undefined>;
    /**
     * URI for containers created in current Region with default ECR image tag
     */
    public /*out*/ readonly imageUri!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the infrastructure configuration.
     */
    public readonly infrastructureConfigurationArn!: pulumi.Output<string | undefined>;
    /**
     * The name of the image.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The tags associated with the image.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Workflows to define the image build process
     */
    public readonly workflows!: pulumi.Output<outputs.imagebuilder.ImageWorkflowConfiguration[] | undefined>;

    /**
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ImageArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["containerRecipeArn"] = args ? args.containerRecipeArn : undefined;
            resourceInputs["distributionConfigurationArn"] = args ? args.distributionConfigurationArn : undefined;
            resourceInputs["enhancedImageMetadataEnabled"] = args ? args.enhancedImageMetadataEnabled : undefined;
            resourceInputs["executionRole"] = args ? args.executionRole : undefined;
            resourceInputs["imageRecipeArn"] = args ? args.imageRecipeArn : undefined;
            resourceInputs["imageScanningConfiguration"] = args ? args.imageScanningConfiguration : undefined;
            resourceInputs["imageTestsConfiguration"] = args ? args.imageTestsConfiguration : undefined;
            resourceInputs["infrastructureConfigurationArn"] = args ? args.infrastructureConfigurationArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["workflows"] = args ? args.workflows : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["imageId"] = undefined /*out*/;
            resourceInputs["imageUri"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["containerRecipeArn"] = undefined /*out*/;
            resourceInputs["distributionConfigurationArn"] = undefined /*out*/;
            resourceInputs["enhancedImageMetadataEnabled"] = undefined /*out*/;
            resourceInputs["executionRole"] = undefined /*out*/;
            resourceInputs["imageId"] = undefined /*out*/;
            resourceInputs["imageRecipeArn"] = undefined /*out*/;
            resourceInputs["imageScanningConfiguration"] = undefined /*out*/;
            resourceInputs["imageTestsConfiguration"] = undefined /*out*/;
            resourceInputs["imageUri"] = undefined /*out*/;
            resourceInputs["infrastructureConfigurationArn"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["workflows"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["containerRecipeArn", "distributionConfigurationArn", "enhancedImageMetadataEnabled", "imageRecipeArn", "imageScanningConfiguration", "imageTestsConfiguration", "infrastructureConfigurationArn", "tags.*", "workflows[*]"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Image.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    /**
     * The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.
     */
    containerRecipeArn?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the distribution configuration.
     */
    distributionConfigurationArn?: pulumi.Input<string>;
    /**
     * Collects additional information about the image being created, including the operating system (OS) version and package list.
     */
    enhancedImageMetadataEnabled?: pulumi.Input<boolean>;
    /**
     * The execution role name/ARN for the image build, if provided
     */
    executionRole?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.
     */
    imageRecipeArn?: pulumi.Input<string>;
    /**
     * Contains settings for vulnerability scans.
     */
    imageScanningConfiguration?: pulumi.Input<inputs.imagebuilder.ImageScanningConfigurationArgs>;
    /**
     * The image tests configuration used when creating this image.
     */
    imageTestsConfiguration?: pulumi.Input<inputs.imagebuilder.ImageTestsConfigurationArgs>;
    /**
     * The Amazon Resource Name (ARN) of the infrastructure configuration.
     */
    infrastructureConfigurationArn?: pulumi.Input<string>;
    /**
     * The tags associated with the image.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Workflows to define the image build process
     */
    workflows?: pulumi.Input<pulumi.Input<inputs.imagebuilder.ImageWorkflowConfigurationArgs>[]>;
}
