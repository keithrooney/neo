package com.neo.core.models;

import org.immutables.value.Value;

@Value.Immutable
public interface Address {

	String getLine1();

	String getLine2();

	County getCounty();

}
