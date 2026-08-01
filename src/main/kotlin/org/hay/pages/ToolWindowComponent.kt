package org.hay.pages

import com.intellij.openapi.project.Project
import com.intellij.openapi.wm.ToolWindow
import org.hay.pages.common.Colors
import java.awt.BorderLayout
import javax.swing.JPanel

class ToolWindowComponent(project: Project, toolWindow: ToolWindow) {

    var askInputComponent: AskInputComponent = AskInputComponent(project)
    var contentComponent: ContentComponent = ContentComponent(project)

     fun createPanel(): JPanel {
        return JPanel(BorderLayout()).apply {
            background = Colors.PLUGIN_COLOR

            add(contentComponent.createPanel(), BorderLayout.CENTER)
            add(askInputComponent.createPanel(), BorderLayout.SOUTH)
        }
    }

}