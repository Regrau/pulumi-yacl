// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Yandex Cloud Logging group resource. For more information, see
 * [the official documentation](https://cloud.yandex.com/en/docs/logging/concepts/log-group).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const group1 = new yandex.LoggingGroup("group1", {
 *     folderId: yandex_resourcemanager_folder_test_folder.id,
 * });
 * ```
 */
export class LoggingGroup extends pulumi.CustomResource {
    /**
     * Get an existing LoggingGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoggingGroupState, opts?: pulumi.CustomResourceOptions): LoggingGroup {
        return new LoggingGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/loggingGroup:LoggingGroup';

    /**
     * Returns true if the given object is an instance of LoggingGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoggingGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoggingGroup.__pulumiType;
    }

    /**
     * ID of the cloud that the Yandex Cloud Logging group belong to.
     */
    public /*out*/ readonly cloudId!: pulumi.Output<string>;
    /**
     * The Yandex Cloud Logging group creation timestamp.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly dataStream!: pulumi.Output<string | undefined>;
    /**
     * A description for the Yandex Cloud Logging group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ID of the folder that the Yandex Cloud Logging group belongs to.
     * It will be deduced from provider configuration if not set explicitly.
     */
    public readonly folderId!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs to assign to the Yandex Cloud Logging group.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name for the Yandex Cloud Logging group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Log entries retention period for the Yandex Cloud Logging group.
     */
    public readonly retentionPeriod!: pulumi.Output<string>;
    /**
     * The Yandex Cloud Logging group status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a LoggingGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LoggingGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoggingGroupArgs | LoggingGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoggingGroupState | undefined;
            resourceInputs["cloudId"] = state ? state.cloudId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["dataStream"] = state ? state.dataStream : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["retentionPeriod"] = state ? state.retentionPeriod : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as LoggingGroupArgs | undefined;
            resourceInputs["dataStream"] = args ? args.dataStream : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["retentionPeriod"] = args ? args.retentionPeriod : undefined;
            resourceInputs["cloudId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoggingGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoggingGroup resources.
 */
export interface LoggingGroupState {
    /**
     * ID of the cloud that the Yandex Cloud Logging group belong to.
     */
    cloudId?: pulumi.Input<string>;
    /**
     * The Yandex Cloud Logging group creation timestamp.
     */
    createdAt?: pulumi.Input<string>;
    dataStream?: pulumi.Input<string>;
    /**
     * A description for the Yandex Cloud Logging group.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of the folder that the Yandex Cloud Logging group belongs to.
     * It will be deduced from provider configuration if not set explicitly.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the Yandex Cloud Logging group.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name for the Yandex Cloud Logging group.
     */
    name?: pulumi.Input<string>;
    /**
     * Log entries retention period for the Yandex Cloud Logging group.
     */
    retentionPeriod?: pulumi.Input<string>;
    /**
     * The Yandex Cloud Logging group status.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoggingGroup resource.
 */
export interface LoggingGroupArgs {
    dataStream?: pulumi.Input<string>;
    /**
     * A description for the Yandex Cloud Logging group.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of the folder that the Yandex Cloud Logging group belongs to.
     * It will be deduced from provider configuration if not set explicitly.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the Yandex Cloud Logging group.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name for the Yandex Cloud Logging group.
     */
    name?: pulumi.Input<string>;
    /**
     * Log entries retention period for the Yandex Cloud Logging group.
     */
    retentionPeriod?: pulumi.Input<string>;
}
