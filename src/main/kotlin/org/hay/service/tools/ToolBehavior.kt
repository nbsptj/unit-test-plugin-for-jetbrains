package org.hay.service.tools

import org.hay.service.dto.Tool

// tool common behavior
interface ToolBehavior {

    companion object {

        @JvmStatic
        fun allTools(): List<Tool> {
            val tools = listOf(
                ReadFileTool(),
            )
            return tools.map { it.buildTool() }
        }

    }

    fun buildTool(): Tool

}