package org.hay.pages

import com.intellij.openapi.project.Project
import com.intellij.ui.dsl.builder.panel
import javax.swing.JPanel

class ContentComponent(project: Project) {

    fun createPanel(): JPanel = panel {
        row {
            label("Hello")
        }
    }

}