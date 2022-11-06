package com.neo.core.models;

import org.immutables.value.Value;

@Value.Immutable
public interface Account extends Model {

	ExternalReference getReference();

}
