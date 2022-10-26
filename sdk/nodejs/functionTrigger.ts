// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class FunctionTrigger extends pulumi.CustomResource {
    /**
     * Get an existing FunctionTrigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionTriggerState, opts?: pulumi.CustomResourceOptions): FunctionTrigger {
        return new FunctionTrigger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/functionTrigger:FunctionTrigger';

    /**
     * Returns true if the given object is an instance of FunctionTrigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FunctionTrigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FunctionTrigger.__pulumiType;
    }

    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly dlq!: pulumi.Output<outputs.FunctionTriggerDlq | undefined>;
    public readonly folderId!: pulumi.Output<string>;
    public readonly function!: pulumi.Output<outputs.FunctionTriggerFunction>;
    public readonly iot!: pulumi.Output<outputs.FunctionTriggerIot | undefined>;
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly logGroup!: pulumi.Output<outputs.FunctionTriggerLogGroup | undefined>;
    public readonly logging!: pulumi.Output<outputs.FunctionTriggerLogging | undefined>;
    public readonly messageQueue!: pulumi.Output<outputs.FunctionTriggerMessageQueue | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly objectStorage!: pulumi.Output<outputs.FunctionTriggerObjectStorage | undefined>;
    public readonly timer!: pulumi.Output<outputs.FunctionTriggerTimer | undefined>;

    /**
     * Create a FunctionTrigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionTriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionTriggerArgs | FunctionTriggerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionTriggerState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dlq"] = state ? state.dlq : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["function"] = state ? state.function : undefined;
            resourceInputs["iot"] = state ? state.iot : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["logGroup"] = state ? state.logGroup : undefined;
            resourceInputs["logging"] = state ? state.logging : undefined;
            resourceInputs["messageQueue"] = state ? state.messageQueue : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["objectStorage"] = state ? state.objectStorage : undefined;
            resourceInputs["timer"] = state ? state.timer : undefined;
        } else {
            const args = argsOrState as FunctionTriggerArgs | undefined;
            if ((!args || args.function === undefined) && !opts.urn) {
                throw new Error("Missing required property 'function'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dlq"] = args ? args.dlq : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["function"] = args ? args.function : undefined;
            resourceInputs["iot"] = args ? args.iot : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["logGroup"] = args ? args.logGroup : undefined;
            resourceInputs["logging"] = args ? args.logging : undefined;
            resourceInputs["messageQueue"] = args ? args.messageQueue : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["objectStorage"] = args ? args.objectStorage : undefined;
            resourceInputs["timer"] = args ? args.timer : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FunctionTrigger.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FunctionTrigger resources.
 */
export interface FunctionTriggerState {
    createdAt?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    dlq?: pulumi.Input<inputs.FunctionTriggerDlq>;
    folderId?: pulumi.Input<string>;
    function?: pulumi.Input<inputs.FunctionTriggerFunction>;
    iot?: pulumi.Input<inputs.FunctionTriggerIot>;
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    logGroup?: pulumi.Input<inputs.FunctionTriggerLogGroup>;
    logging?: pulumi.Input<inputs.FunctionTriggerLogging>;
    messageQueue?: pulumi.Input<inputs.FunctionTriggerMessageQueue>;
    name?: pulumi.Input<string>;
    objectStorage?: pulumi.Input<inputs.FunctionTriggerObjectStorage>;
    timer?: pulumi.Input<inputs.FunctionTriggerTimer>;
}

/**
 * The set of arguments for constructing a FunctionTrigger resource.
 */
export interface FunctionTriggerArgs {
    description?: pulumi.Input<string>;
    dlq?: pulumi.Input<inputs.FunctionTriggerDlq>;
    folderId?: pulumi.Input<string>;
    function: pulumi.Input<inputs.FunctionTriggerFunction>;
    iot?: pulumi.Input<inputs.FunctionTriggerIot>;
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    logGroup?: pulumi.Input<inputs.FunctionTriggerLogGroup>;
    logging?: pulumi.Input<inputs.FunctionTriggerLogging>;
    messageQueue?: pulumi.Input<inputs.FunctionTriggerMessageQueue>;
    name?: pulumi.Input<string>;
    objectStorage?: pulumi.Input<inputs.FunctionTriggerObjectStorage>;
    timer?: pulumi.Input<inputs.FunctionTriggerTimer>;
}