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

"""Integration tests for the DocumentDB API DBCluster resource
"""

import time

import pytest

from acktest.k8s import resource as k8s
from acktest.resources import random_suffix_name
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_documentdb_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e import condition
from e2e import db_cluster
from e2e.fixtures import k8s_secret
from e2e import tag

RESOURCE_PLURAL = 'dbclusters'

DELETE_WAIT_AFTER_SECONDS = 120
CHECK_STATUS_WAIT_SECONDS = 60*2
MODIFY_WAIT_AFTER_SECONDS = 20

# MUP == Master user password...
MUP_NS = "default"
MUP_SEC_NAME = "dbclustersecrets"
MUP_SEC_KEY = "master_user_password"
MUP_SEC_VAL = "secretpass123456"

@pytest.fixture
def docdb_cluster(k8s_secret):
    db_cluster_id = random_suffix_name("ack-docdb", 20)
    secret = k8s_secret(
        MUP_NS,
        f"{MUP_SEC_NAME}-docdb",
        MUP_SEC_KEY,
        MUP_SEC_VAL,
    )

    replacements = REPLACEMENT_VALUES.copy()
    replacements["DB_CLUSTER_ID"] = db_cluster_id
    replacements["MASTER_USER_NAME"] = "ackuser"
    replacements["MASTER_USER_PASS_SECRET_NAMESPACE"] = secret.ns
    replacements["MASTER_USER_PASS_SECRET_NAME"] = secret.name
    replacements["MASTER_USER_PASS_SECRET_KEY"] = secret.key

    resource_data = load_documentdb_resource(
        "db_cluster",
        additional_replacements=replacements,
    )

    ref = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        db_cluster_id, namespace="default",
    )
    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    assert cr is not None
    assert 'status' in cr
    assert 'status' in cr['status']
    assert cr['status']['status'] == 'creating'
    condition.assert_not_synced(ref)

    yield (ref, cr, db_cluster_id)

    try:
        _, deleted = k8s.delete_custom_resource(ref, 3, 10)
        assert deleted
        time.sleep(DELETE_WAIT_AFTER_SECONDS)
    except:
        pass

    db_cluster.wait_until_deleted(db_cluster_id)

@service_marker
@pytest.mark.canary
class TestDBCluster:
    def test_crud(
            self, docdb_cluster,
    ):
        ref, cr, db_cluster_id = docdb_cluster

        db_cluster.wait_until(
            db_cluster_id,
            db_cluster.status_matches('available'),
        )

        time.sleep(CHECK_STATUS_WAIT_SECONDS)

        # Before we update the DBCluster CR below, let's check to see that the
        # Status field in the CR has been updated to something other than
        # 'creating', which is what is set after the initial creation.  The
        # CR's `Status.Status` should be updated because the CR is requeued on
        # successful reconciliation loops and subsequent reconciliation loops
        # call ReadOne and should update the CR's Status with the latest
        # observed information.
        cr = k8s.get_resource(ref)
        assert cr is not None
        assert 'status' in cr
        assert 'status' in cr['status']
        assert cr['status']['status'] != 'creating'
        condition.assert_synced(ref)
