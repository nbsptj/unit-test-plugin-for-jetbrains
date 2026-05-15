package org.hay.service.dto

import com.alibaba.fastjson.annotation.JSONField

open class Message (
    val role: String, // system, user, assistant, tool
    val content: String?, // 消息内容
)

open class NameMessage (
    role: String,
    content: String?,
    val name: String?, // 可以选填的参与者的名称，为模型提供信息以区分相同角色的参与者。
) : Message(role, content)

class SystemMessage (
    content: String,
    name: String? = null,
) : NameMessage("system", content, name)

class UserMessage (
    content: String?,
    name: String? = null,
) : NameMessage("user", content, name)

class AssistantMessage (
    content: String?,
    name: String? = null,
) : NameMessage("assistant", content, name)

class ToolMessage (
    content: String?,
    @JSONField(name = "tool_call_id") val toolCallId: String,
) : Message("tool", content)