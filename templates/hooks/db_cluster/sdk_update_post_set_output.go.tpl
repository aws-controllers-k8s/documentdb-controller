// ModifyDBCluster returns stale status for MasterUserSecret when unsetting ManageMasterUserPassword.
if delta.DifferentAt("Spec.ManageMasterUserPassword") && desired.ko.Spec.ManageMasterUserPassword != nil && *desired.ko.Spec.ManageMasterUserPassword == false {
    ackcondition.SetSynced(&resource{ko}, corev1.ConditionFalse, aws.String("DBCluster updated, requeue for updated status"), nil)
}