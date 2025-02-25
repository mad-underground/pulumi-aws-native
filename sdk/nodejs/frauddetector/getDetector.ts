// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A resource schema for a Detector in Amazon Fraud Detector.
 */
export function getDetector(args: GetDetectorArgs, opts?: pulumi.InvokeOptions): Promise<GetDetectorResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:frauddetector:getDetector", {
        "arn": args.arn,
    }, opts);
}

export interface GetDetectorArgs {
    /**
     * The ARN of the detector.
     */
    arn: string;
}

export interface GetDetectorResult {
    /**
     * The ARN of the detector.
     */
    readonly arn?: string;
    /**
     * The models to associate with this detector.
     */
    readonly associatedModels?: outputs.frauddetector.DetectorModel[];
    /**
     * The time when the detector was created.
     */
    readonly createdTime?: string;
    /**
     * The description of the detector.
     */
    readonly description?: string;
    /**
     * The active version ID of the detector
     */
    readonly detectorVersionId?: string;
    /**
     * The desired detector version status for the detector
     */
    readonly detectorVersionStatus?: enums.frauddetector.DetectorVersionStatus;
    /**
     * The event type to associate this detector with.
     */
    readonly eventType?: outputs.frauddetector.DetectorEventType;
    /**
     * The time when the detector was last updated.
     */
    readonly lastUpdatedTime?: string;
    readonly ruleExecutionMode?: enums.frauddetector.DetectorRuleExecutionMode;
    readonly rules?: outputs.frauddetector.DetectorRule[];
    /**
     * Tags associated with this detector.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * A resource schema for a Detector in Amazon Fraud Detector.
 */
export function getDetectorOutput(args: GetDetectorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDetectorResult> {
    return pulumi.output(args).apply((a: any) => getDetector(a, opts))
}

export interface GetDetectorOutputArgs {
    /**
     * The ARN of the detector.
     */
    arn: pulumi.Input<string>;
}
