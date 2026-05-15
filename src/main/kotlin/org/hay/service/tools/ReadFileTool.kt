package org.hay.service.tools

import org.hay.service.dto.Function
import org.hay.service.dto.Parameter
import org.hay.service.dto.Tool
import org.hay.service.dto.Type

class ReadFileTool : ToolBehavior {

    val name = "read_file"
    val description = "读取文件内容工具"

    override fun buildTool(): Tool {
        return Tool(
            function = Function(
                name = name,
                description = description,
                parameters = Parameter(type = Type.OBJ.value, properties = mutableMapOf(
                    Pair("path", Parameter(type = Type.STR.value, description = "文件的绝对路径"))
                ))
            )
        )
    }

}

data class ReadFileArgument(
    val path: String
)