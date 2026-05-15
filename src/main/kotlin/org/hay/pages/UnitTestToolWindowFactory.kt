package org.hay.pages

import com.intellij.openapi.diagnostic.Logger
import com.intellij.openapi.project.Project
import com.intellij.openapi.wm.ToolWindow
import com.intellij.openapi.wm.ToolWindowFactory
import com.intellij.ui.content.ContentFactory
import org.hay.service.PromptService

class UnitTestToolWindowFactory: ToolWindowFactory {

    val log = Logger.getInstance(UnitTestToolWindowFactory::class.java)

    override fun createToolWindowContent(
        project: Project,
        toolWindow: ToolWindow
    ) {
        if (project.isDisposed) {
            log.warn("project: ${project.basePath} has disposed, can't create toolwindow")
            return
        }
        PromptService.registerCodebaseVariable(project)
        val panel = Main(project, toolWindow).createPanel()
        val content = ContentFactory.getInstance().createContent(panel, "", false)
        toolWindow.contentManager.addContent(content)
    }

}