# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Integration tests for the DocDB API DBInstance resource
"""

import time

import pytest

from acktest.k8s import resource as k8s
from acktest.resources import random_suffix_name
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_documentdb_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e import condition
from e2e import db_instance
from e2e import tag
from e2e.fixtures import k8s_secret
from .test_db_cluster import docdb_cluster

RESOURCE_PLURAL = 'dbinstances'

DELETE_WAIT_AFTER_SECONDS = 60

MAX_WAIT_FOR_SYNCED_MINUTES = 20

MODIFY_WAIT_AFTER_SECONDS = 60

MUP_NS = "default"
MUP_SEC_NAME_PREFIX = "dbinstancesecrets"
MUP_SEC_KEY = "master_user_password"
MUP_SEC_VAL = "secretpass123456"

@pytest.fixture
def simple_instance(k8s_secret, docdb_cluster):
    ref, cr, db_cluster_id = docdb_cluster
    db_instance_id = random_suffix_name("test-docdb", 20)

    replacements = REPLACEMENT_VALUES.copy()
    replacements['COPY_TAGS_TO_SNAPSHOT'] = "False"
    replacements["DB_INSTANCE_ID"] = db_instance_id
    replacements["DB_CLUSTER_ID"] = db_cluster_id

    resource_data = load_documentdb_resource(
        "db_instance",
        additional_replacements=replacements,
    )

    # Create the k8s resource
    ref = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        db_instance_id, namespace="default",
    )
    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    assert cr is not None
    assert k8s.get_resource_exists(ref)

    yield (ref, cr)

    # Try to delete, if doesn't already exist
    try:
        _, deleted = k8s.delete_custom_resource(ref, 3, 10)
    except:
        pass
    db_instance.wait_until_deleted(db_instance_id)


@service_marker
@pytest.mark.canary
class TestDBInstance:
    def test_crud(
            self,
            simple_instance,
    ):
        (ref, cr) = simple_instance
        db_instance_id = cr["spec"]["dbInstanceIdentifier"]

        assert 'status' in cr
        assert 'dbInstanceStatus' in cr['status']
        assert cr['status']['dbInstanceStatus'] == 'creating'
        condition.assert_not_synced(ref)

        # Wait for the resource to get synced
        assert k8s.wait_on_condition(ref, "ACK.ResourceSynced", "True", wait_periods=MAX_WAIT_FOR_SYNCED_MINUTES)

        # After the resource is synced, assert that DBInstanceStatus is available
        latest = db_instance.get(db_instance_id)
        assert latest is not None
        assert latest['DBInstanceStatus'] == 'available'

        cr = k8s.get_resource(ref)
        assert cr is not None
        assert 'status' in cr
        assert 'dbInstanceStatus' in cr['status']
        assert cr['status']['dbInstanceStatus'] != 'creating'
        condition.assert_synced(ref)

        latest = db_instance.get(db_instance_id)
        assert latest is not None
        assert latest['CopyTagsToSnapshot'] is False
        updates = {
            "spec": {"copyTagsToSnapshot": True},
        }
        k8s.patch_custom_resource(ref, updates)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)

        # wait for the resource to get synced after the patch
        assert k8s.wait_on_condition(ref, "ACK.ResourceSynced", "True", wait_periods=MAX_WAIT_FOR_SYNCED_MINUTES)

        # After resource is synced again, assert that patches are reflected in the AWS resource
        latest = db_instance.get(db_instance_id)
        assert latest is not None
        assert latest['CopyTagsToSnapshot'] is True

        updates = {
            "spec": {"copyTagsToSnapshot": False},
        }
        k8s.patch_custom_resource(ref, updates)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)

        # wait for the resource to get synced after the patch
        assert k8s.wait_on_condition(ref, "ACK.ResourceSynced", "True", wait_periods=MAX_WAIT_FOR_SYNCED_MINUTES)

        # After resource is synced again, assert that patches are reflected in the AWS resource
        latest = db_instance.get(db_instance_id)
        assert latest is not None
        assert latest['CopyTagsToSnapshot'] is False

        arn = latest['DBInstanceArn']
        expect_tags = [
            {"Key": "environment", "Value": "dev"}
        ]
        latest_tags = tag.clean(db_instance.get_tags(arn))
        assert expect_tags == latest_tags

        # OK, now let's update the tag set and check that the tags are
        # updated accordingly.
        new_tags = [
            {
                "key": "environment",
                "value": "prod",
            }
        ]
        updates = {
            "spec": {"tags": new_tags},
        }
        k8s.patch_custom_resource(ref, updates)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)

        latest_tags = tag.clean(db_instance.get_tags(arn))
        after_update_expected_tags = [
            {
                "Key": "environment",
                "Value": "prod",
            }
        ]
        assert latest_tags == after_update_expected_tags
