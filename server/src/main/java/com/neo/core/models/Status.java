package com.neo.core.models;

/**
 * PENDING -> ACTIVE<br>
 * PENDING -> DECLINED<br>
 * PENDING -> ACTIVE -> SUSPENDED -> LOCKED<br>
 * PENDING -> ACTIVE -> CLOSED
 */
public enum Status {

	PENDING,

	ACTIVE,

	DECLINED,

	SUSPENDED,

	LOCKED,

	CLOSED

}
