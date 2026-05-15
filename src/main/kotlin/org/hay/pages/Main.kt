package org.hay.pages

import com.intellij.openapi.project.Project
import com.intellij.openapi.wm.ToolWindow
import com.intellij.ui.JBColor
import com.intellij.ui.dsl.builder.panel
import org.hay.service.AgentService
import org.hay.service.PromptService
import java.util.concurrent.atomic.AtomicBoolean
import javax.swing.JButton
import javax.swing.JPanel

class Main(val project: Project, val toolwindow: ToolWindow, val agentService: AgentService = AgentService(project)) {

    /**
     * create panel
     */
    fun createPanel(): JPanel {
        return panel {
            separator(JBColor.BLUE)
            indent {
                row {
                    button("generate") {
                        agentService.start(
                            PromptService.getUnittestSystem(),
                            PromptService.getUnittestUser())
                    }
                }
            }
        }.withBackground(JBColor.ORANGE)
    }

}