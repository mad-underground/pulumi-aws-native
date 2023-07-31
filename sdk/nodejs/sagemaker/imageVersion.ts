// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::ImageVersion
 */
export class ImageVersion extends pulumi.CustomResource {
    /**
     * Get an existing ImageVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ImageVersion {
        return new ImageVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sagemaker:ImageVersion';

    /**
     * Returns true if the given object is an instance of ImageVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ImageVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ImageVersion.__pulumiType;
    }

    public readonly alias!: pulumi.Output<string | undefined>;
    public readonly aliases!: pulumi.Output<string[] | undefined>;
    public readonly baseImage!: pulumi.Output<string>;
    public /*out*/ readonly containerImage!: pulumi.Output<string>;
    public readonly horovod!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly imageArn!: pulumi.Output<string>;
    public readonly imageName!: pulumi.Output<string>;
    public /*out*/ readonly imageVersionArn!: pulumi.Output<string>;
    public readonly jobType!: pulumi.Output<enums.sagemaker.ImageVersionJobType | undefined>;
    public readonly mLFramework!: pulumi.Output<string | undefined>;
    public readonly processor!: pulumi.Output<enums.sagemaker.ImageVersionProcessor | undefined>;
    public readonly programmingLang!: pulumi.Output<string | undefined>;
    public readonly releaseNotes!: pulumi.Output<string | undefined>;
    public readonly vendorGuidance!: pulumi.Output<enums.sagemaker.ImageVersionVendorGuidance | undefined>;
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a ImageVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.baseImage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'baseImage'");
            }
            if ((!args || args.imageName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageName'");
            }
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["aliases"] = args ? args.aliases : undefined;
            resourceInputs["baseImage"] = args ? args.baseImage : undefined;
            resourceInputs["horovod"] = args ? args.horovod : undefined;
            resourceInputs["imageName"] = args ? args.imageName : undefined;
            resourceInputs["jobType"] = args ? args.jobType : undefined;
            resourceInputs["mLFramework"] = args ? args.mLFramework : undefined;
            resourceInputs["processor"] = args ? args.processor : undefined;
            resourceInputs["programmingLang"] = args ? args.programmingLang : undefined;
            resourceInputs["releaseNotes"] = args ? args.releaseNotes : undefined;
            resourceInputs["vendorGuidance"] = args ? args.vendorGuidance : undefined;
            resourceInputs["containerImage"] = undefined /*out*/;
            resourceInputs["imageArn"] = undefined /*out*/;
            resourceInputs["imageVersionArn"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        } else {
            resourceInputs["alias"] = undefined /*out*/;
            resourceInputs["aliases"] = undefined /*out*/;
            resourceInputs["baseImage"] = undefined /*out*/;
            resourceInputs["containerImage"] = undefined /*out*/;
            resourceInputs["horovod"] = undefined /*out*/;
            resourceInputs["imageArn"] = undefined /*out*/;
            resourceInputs["imageName"] = undefined /*out*/;
            resourceInputs["imageVersionArn"] = undefined /*out*/;
            resourceInputs["jobType"] = undefined /*out*/;
            resourceInputs["mLFramework"] = undefined /*out*/;
            resourceInputs["processor"] = undefined /*out*/;
            resourceInputs["programmingLang"] = undefined /*out*/;
            resourceInputs["releaseNotes"] = undefined /*out*/;
            resourceInputs["vendorGuidance"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ImageVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ImageVersion resource.
 */
export interface ImageVersionArgs {
    alias?: pulumi.Input<string>;
    aliases?: pulumi.Input<pulumi.Input<string>[]>;
    baseImage: pulumi.Input<string>;
    horovod?: pulumi.Input<boolean>;
    imageName: pulumi.Input<string>;
    jobType?: pulumi.Input<enums.sagemaker.ImageVersionJobType>;
    mLFramework?: pulumi.Input<string>;
    processor?: pulumi.Input<enums.sagemaker.ImageVersionProcessor>;
    programmingLang?: pulumi.Input<string>;
    releaseNotes?: pulumi.Input<string>;
    vendorGuidance?: pulumi.Input<enums.sagemaker.ImageVersionVendorGuidance>;
}
