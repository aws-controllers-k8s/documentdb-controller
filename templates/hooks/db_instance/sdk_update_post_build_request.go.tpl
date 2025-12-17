        if !delta.DifferentAt("Spec.CACertificateIdentifier") {
			input.CACertificateIdentifier = nil
		}

		if !delta.DifferentAt("Spec.AutoMinorVersionUpgrade") {
			input.AutoMinorVersionUpgrade = nil
		}