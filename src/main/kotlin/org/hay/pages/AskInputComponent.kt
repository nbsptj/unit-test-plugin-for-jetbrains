package org.hay.pages

import com.intellij.openapi.project.Project
import com.intellij.openapi.ui.VerticalFlowLayout
import com.intellij.ui.components.JBTextArea
import com.intellij.ui.dsl.builder.RightGap
import com.intellij.ui.dsl.builder.columns
import com.intellij.ui.dsl.builder.panel
import com.intellij.ui.dsl.gridLayout.HorizontalAlign
import org.hay.Bundle
import java.awt.event.ActionListener
import javax.swing.JButton
import javax.swing.JPanel

class AskInputComponent(project: Project) {

    fun createPanel(): JPanel = panel {
        row {
            textArea()
                .resizableColumn()
                .horizontalAlign(HorizontalAlign.FILL)
        }
        row {
            button(Bundle.message("send_button.text"), actionListener = {
                println("clicked!")
            })
                .horizontalAlign(HorizontalAlign.RIGHT)
        }
    }

}