package org.hay.pages

import com.intellij.openapi.project.Project
import com.intellij.openapi.wm.ToolWindow
import com.intellij.openapi.wm.ToolWindowFactory
import com.intellij.ui.content.Content
import com.intellij.ui.content.ContentFactory

class ToolWindowFactory: ToolWindowFactory {

    var toolWindowComponent: ToolWindowComponent? = null
    var content: Content? = null

    override fun createToolWindowContent(
        project: Project,
        toolWindow: ToolWindow
    ) {
        toolWindowComponent = ToolWindowComponent(project, toolWindow)
        content = ContentFactory.getInstance().createContent(toolWindowComponent!!.createPanel(), "", false)
        toolWindow.contentManager.addContent(content!!)
    }

}