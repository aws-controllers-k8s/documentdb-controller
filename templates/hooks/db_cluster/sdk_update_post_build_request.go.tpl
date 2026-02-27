// When ManageMasterUserPassword is in the delta and set to false need to ensure that
// MasterUserPassword is included in the payload.
if delta.DifferentAt("Spec.ManageMasterUserPassword") && desired.ko.Spec.ManageMasterUserPassword != nil && *desired.ko.Spec.ManageMasterUserPassword == false {
    if desired.ko.Spec.MasterUserPassword != nil {
        tmpSecret, err := rm.rr.SecretValueFromReference(ctx, desired.ko.Spec.MasterUserPassword)
        if err != nil {
            return nil, ackrequeue.Needed(err)
        }
        if tmpSecret != "" {
            input.MasterUserPassword = aws.String(tmpSecret)
        }
    }
}