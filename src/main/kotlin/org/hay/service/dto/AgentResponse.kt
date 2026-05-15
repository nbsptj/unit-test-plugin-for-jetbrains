package org.hay.service.dto

import com.alibaba.fastjson.annotation.JSONField

data class AgentResponse(
    val id: String,
    @JSONField(name = "object") val obj: String,
    val created: Long,
    val model: String,
    val choices: MutableList<Choice>,
    val usage: MutableMap<String, Int>,
)

data class Choice(
    val index: Int,
    val message: Message,
    @JSONField(name = "finish_reason") val finishReason: String,
    val delta: Delta,
)

data class Delta(
    val role: String,
    val content: String?,
    @JSONField(name = "tool_calls") val toolCalls: MutableList<ToolCall>,
)

data class ToolCall(
    val id: String,
    val type: String,
    val function: ExecuteFunction,
)

data class ExecuteFunction (
    val name: String,
    val arguments: String,
)
