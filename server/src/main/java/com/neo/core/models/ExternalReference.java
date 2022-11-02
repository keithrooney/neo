package com.neo.core.models;

import org.immutables.value.Value;

@Value.Immutable
public interface ExternalReference {
	String getId();
}
