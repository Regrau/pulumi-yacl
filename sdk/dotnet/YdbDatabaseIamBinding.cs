// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Yandex
{
    /// <summary>
    /// ## yandex\_ydb\_database\_iam\_binding
    /// 
    /// Allows creation and management of a single binding within IAM policy for
    /// an existing Managed YDB Database instance.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Yandex = Pulumi.Yandex;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var database1 = new Yandex.YdbDatabaseServerless("database1", new()
    ///     {
    ///         FolderId = data.Yandex_resourcemanager_folder.Test_folder.Id,
    ///     });
    /// 
    ///     var viewer = new Yandex.YdbDatabaseIamBinding("viewer", new()
    ///     {
    ///         DatabaseId = database1.Id,
    ///         Role = "ydb.viewer",
    ///         Members = new[]
    ///         {
    ///             "userAccount:foo_user_id",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the `database_id` and role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import yandex:index/ydbDatabaseIamBinding:ydbDatabaseIamBinding viewer "database_id ydb.viewer"
    /// ```
    /// </summary>
    [YandexResourceType("yandex:index/ydbDatabaseIamBinding:ydbDatabaseIamBinding")]
    public partial class YdbDatabaseIamBinding : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The [Managed Service YDB instance](https://cloud.yandex.com/docs/ydb/) Database ID to apply a binding to.
        /// </summary>
        [Output("databaseId")]
        public Output<string> DatabaseId { get; private set; } = null!;

        /// <summary>
        /// Identities that will be granted the privilege in `role`.
        /// Each entry can have one of the following values:
        /// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
        /// * **serviceAccount:{service_account_id}**: A unique service account ID.
        /// * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. See [roles](https://cloud.yandex.com/docs/ydb/security/).
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        [Output("sleepAfter")]
        public Output<int?> SleepAfter { get; private set; } = null!;


        /// <summary>
        /// Create a YdbDatabaseIamBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public YdbDatabaseIamBinding(string name, YdbDatabaseIamBindingArgs args, CustomResourceOptions? options = null)
            : base("yandex:index/ydbDatabaseIamBinding:ydbDatabaseIamBinding", name, args ?? new YdbDatabaseIamBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private YdbDatabaseIamBinding(string name, Input<string> id, YdbDatabaseIamBindingState? state = null, CustomResourceOptions? options = null)
            : base("yandex:index/ydbDatabaseIamBinding:ydbDatabaseIamBinding", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/regrau/pulumi-yandex/releases",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing YdbDatabaseIamBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static YdbDatabaseIamBinding Get(string name, Input<string> id, YdbDatabaseIamBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new YdbDatabaseIamBinding(name, id, state, options);
        }
    }

    public sealed class YdbDatabaseIamBindingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [Managed Service YDB instance](https://cloud.yandex.com/docs/ydb/) Database ID to apply a binding to.
        /// </summary>
        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        [Input("members", required: true)]
        private InputList<string>? _members;

        /// <summary>
        /// Identities that will be granted the privilege in `role`.
        /// Each entry can have one of the following values:
        /// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
        /// * **serviceAccount:{service_account_id}**: A unique service account ID.
        /// * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be applied. See [roles](https://cloud.yandex.com/docs/ydb/security/).
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        [Input("sleepAfter")]
        public Input<int>? SleepAfter { get; set; }

        public YdbDatabaseIamBindingArgs()
        {
        }
        public static new YdbDatabaseIamBindingArgs Empty => new YdbDatabaseIamBindingArgs();
    }

    public sealed class YdbDatabaseIamBindingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [Managed Service YDB instance](https://cloud.yandex.com/docs/ydb/) Database ID to apply a binding to.
        /// </summary>
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// Identities that will be granted the privilege in `role`.
        /// Each entry can have one of the following values:
        /// * **userAccount:{user_id}**: A unique user ID that represents a specific Yandex account.
        /// * **serviceAccount:{service_account_id}**: A unique service account ID.
        /// * **system:{allUsers|allAuthenticatedUsers}**: see [system groups](https://cloud.yandex.com/docs/iam/concepts/access-control/system-group)
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be applied. See [roles](https://cloud.yandex.com/docs/ydb/security/).
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("sleepAfter")]
        public Input<int>? SleepAfter { get; set; }

        public YdbDatabaseIamBindingState()
        {
        }
        public static new YdbDatabaseIamBindingState Empty => new YdbDatabaseIamBindingState();
    }
}