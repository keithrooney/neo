package com.neo.core.models;

import java.sql.Timestamp;

public interface Model {
	String getId();
	Timestamp getCreatedAt();
	Timestamp getUpdateAt();
	Timestamp getDeletedAt();
}
