// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Glue::Classifier
 *
 * @deprecated Classifier is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Classifier extends pulumi.CustomResource {
    /**
     * Get an existing Classifier resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Classifier {
        pulumi.log.warn("Classifier is deprecated: Classifier is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Classifier(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:glue:Classifier';

    /**
     * Returns true if the given object is an instance of Classifier.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Classifier {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Classifier.__pulumiType;
    }

    public readonly csvClassifier!: pulumi.Output<outputs.glue.ClassifierCsvClassifier | undefined>;
    public readonly grokClassifier!: pulumi.Output<outputs.glue.ClassifierGrokClassifier | undefined>;
    public readonly jsonClassifier!: pulumi.Output<outputs.glue.ClassifierJsonClassifier | undefined>;
    public readonly xMLClassifier!: pulumi.Output<outputs.glue.ClassifierXMLClassifier | undefined>;

    /**
     * Create a Classifier resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Classifier is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: ClassifierArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Classifier is deprecated: Classifier is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["csvClassifier"] = args ? args.csvClassifier : undefined;
            resourceInputs["grokClassifier"] = args ? args.grokClassifier : undefined;
            resourceInputs["jsonClassifier"] = args ? args.jsonClassifier : undefined;
            resourceInputs["xMLClassifier"] = args ? args.xMLClassifier : undefined;
        } else {
            resourceInputs["csvClassifier"] = undefined /*out*/;
            resourceInputs["grokClassifier"] = undefined /*out*/;
            resourceInputs["jsonClassifier"] = undefined /*out*/;
            resourceInputs["xMLClassifier"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Classifier.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Classifier resource.
 */
export interface ClassifierArgs {
    csvClassifier?: pulumi.Input<inputs.glue.ClassifierCsvClassifierArgs>;
    grokClassifier?: pulumi.Input<inputs.glue.ClassifierGrokClassifierArgs>;
    jsonClassifier?: pulumi.Input<inputs.glue.ClassifierJsonClassifierArgs>;
    xMLClassifier?: pulumi.Input<inputs.glue.ClassifierXMLClassifierArgs>;
}
