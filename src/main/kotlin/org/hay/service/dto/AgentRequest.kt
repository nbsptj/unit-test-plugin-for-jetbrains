package org.hay.service.dto

data class AgentRequest (
    val messages: MutableList<Message>,
    val tools: List<Tool>,
)