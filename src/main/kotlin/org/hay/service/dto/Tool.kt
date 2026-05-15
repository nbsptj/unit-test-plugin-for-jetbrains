package org.hay.service.dto

data class Tool (
    val type: String = "function",
    val function: Function
)

data class Function (
    val name: String,
    val description: String,
    val parameters: Parameter?,
)

data class Parameter (
    val type: String, // object, string, number, integer, boolean, array, enum, anyOf
    val description: String? = null,
    val items: MutableList<Parameter>? = null, // type is array
    val properties: MutableMap<String, Parameter>? = null, // type is object
    val additionalProperties: MutableList<Parameter>? = null, // type is object
    val required: MutableList<String>? = null,
    val anyOf: MutableList<Parameter>? = null,
)

enum class Type(val value: String) {
    OBJ("object"),
    STR("string"),
    NUM("number"),
    INT("integer"),
    BOOL("boolean"),
    ARR("array"),
    ENUM("enum"),
    ANY_OF("anyOf"),
    ;
}
