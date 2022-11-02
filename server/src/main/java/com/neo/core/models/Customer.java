package com.neo.core.models;

import java.sql.Date;

import org.immutables.value.Value;

@Value.Immutable
public interface Customer extends Model {

	String getEmail();

	Address getAddress();

	String getForename();

	String getSurname();

	Date getDob();

	ExternalReference getReference();

}
