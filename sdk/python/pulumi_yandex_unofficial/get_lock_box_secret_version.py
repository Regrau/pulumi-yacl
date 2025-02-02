# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetLockBoxSecretVersionResult',
    'AwaitableGetLockBoxSecretVersionResult',
    'get_lock_box_secret_version',
    'get_lock_box_secret_version_output',
]

@pulumi.output_type
class GetLockBoxSecretVersionResult:
    """
    A collection of values returned by getLockBoxSecretVersion.
    """
    def __init__(__self__, entries=None, id=None, secret_id=None, version_id=None):
        if entries and not isinstance(entries, list):
            raise TypeError("Expected argument 'entries' to be a list")
        pulumi.set(__self__, "entries", entries)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if secret_id and not isinstance(secret_id, str):
            raise TypeError("Expected argument 'secret_id' to be a str")
        pulumi.set(__self__, "secret_id", secret_id)
        if version_id and not isinstance(version_id, str):
            raise TypeError("Expected argument 'version_id' to be a str")
        pulumi.set(__self__, "version_id", version_id)

    @property
    @pulumi.getter
    def entries(self) -> Sequence['outputs.GetLockBoxSecretVersionEntryResult']:
        """
        List of entries in the Yandex Cloud Lockbox secret version.
        """
        return pulumi.get(self, "entries")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> str:
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> str:
        return pulumi.get(self, "version_id")


class AwaitableGetLockBoxSecretVersionResult(GetLockBoxSecretVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLockBoxSecretVersionResult(
            entries=self.entries,
            id=self.id,
            secret_id=self.secret_id,
            version_id=self.version_id)


def get_lock_box_secret_version(secret_id: Optional[str] = None,
                                version_id: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLockBoxSecretVersionResult:
    """
    Get information about Yandex Cloud Lockbox secret version. For more information,
    see [the official documentation](https://cloud.yandex.com/en/docs/lockbox/).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_yandex as yandex

    my_secret_version = yandex.get_lock_box_secret_version(secret_id="some-secret-id",
        version_id="some-version-id")
    pulumi.export("mySecretEntries", my_secret_version.entries)
    ```

    If you're creating the secret in the same project, then you should indicate `version_id`,
    since otherwise you may refer to a wrong version of the secret
    (e.g. the first version, when it is still empty).

    ```python
    import pulumi
    import pulumi_yandex as yandex
    import pulumi_yandex_unofficial as yandex

    my_secret = yandex.LockboxSecret("mySecret")
    # ...
    my_versionlockbox_secret_version = yandex.LockboxSecretVersion("myVersionlockboxSecretVersion", secret_id=my_secret.id)
    # ...
    my_version_lock_box_secret_version = yandex.get_lock_box_secret_version_output(secret_id=my_secret.id,
        version_id=my_versionlockbox_secret_version.id)
    pulumi.export("mySecretEntries", my_versionlockbox_secret_version.entries)
    ```


    :param str secret_id: The Yandex Cloud Lockbox secret ID.
    :param str version_id: The Yandex Cloud Lockbox secret version ID.
    """
    __args__ = dict()
    __args__['secretId'] = secret_id
    __args__['versionId'] = version_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('yandex:index/getLockBoxSecretVersion:getLockBoxSecretVersion', __args__, opts=opts, typ=GetLockBoxSecretVersionResult).value

    return AwaitableGetLockBoxSecretVersionResult(
        entries=__ret__.entries,
        id=__ret__.id,
        secret_id=__ret__.secret_id,
        version_id=__ret__.version_id)


@_utilities.lift_output_func(get_lock_box_secret_version)
def get_lock_box_secret_version_output(secret_id: Optional[pulumi.Input[str]] = None,
                                       version_id: Optional[pulumi.Input[Optional[str]]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLockBoxSecretVersionResult]:
    """
    Get information about Yandex Cloud Lockbox secret version. For more information,
    see [the official documentation](https://cloud.yandex.com/en/docs/lockbox/).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_yandex as yandex

    my_secret_version = yandex.get_lock_box_secret_version(secret_id="some-secret-id",
        version_id="some-version-id")
    pulumi.export("mySecretEntries", my_secret_version.entries)
    ```

    If you're creating the secret in the same project, then you should indicate `version_id`,
    since otherwise you may refer to a wrong version of the secret
    (e.g. the first version, when it is still empty).

    ```python
    import pulumi
    import pulumi_yandex as yandex
    import pulumi_yandex_unofficial as yandex

    my_secret = yandex.LockboxSecret("mySecret")
    # ...
    my_versionlockbox_secret_version = yandex.LockboxSecretVersion("myVersionlockboxSecretVersion", secret_id=my_secret.id)
    # ...
    my_version_lock_box_secret_version = yandex.get_lock_box_secret_version_output(secret_id=my_secret.id,
        version_id=my_versionlockbox_secret_version.id)
    pulumi.export("mySecretEntries", my_versionlockbox_secret_version.entries)
    ```


    :param str secret_id: The Yandex Cloud Lockbox secret ID.
    :param str version_id: The Yandex Cloud Lockbox secret version ID.
    """
    ...
