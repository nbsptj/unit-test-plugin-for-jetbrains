package org.hay.service

import com.intellij.openapi.diagnostic.Logger
import com.intellij.openapi.project.Project
import org.apache.commons.io.IOUtils
import org.apache.commons.lang3.StringUtils
import org.apache.commons.lang3.Strings
import java.io.FileNotFoundException
import java.nio.charset.StandardCharsets
import java.util.concurrent.ConcurrentHashMap

class PromptService {

    companion object {

        private val logger = Logger.getInstance(PromptService::class.java)

        private val promptCache = ConcurrentHashMap<String, String>()

        private val codebaseVariableCache = ConcurrentHashMap<String, String>()

        fun getPrompt(path: String) : String {
            try {
                var content = promptCache[path]
                if (StringUtils.isBlank(content)) {
                    // load from resources
                    val inputStream = PromptService::class.java.classLoader.getResourceAsStream(path) ?: throw FileNotFoundException("$path not found")
                    content = IOUtils.toString(inputStream, StandardCharsets.UTF_8)
                    for (key in VariableKey.values()) {
                        val value = codebaseVariableCache[key.name] ?: continue
                        content = Strings.CS.replace(content, "@${key.name}@", value)
                    }
                }
                if (StringUtils.isNotBlank(content)) {
                    promptCache[path] = content as String
                }
                if (StringUtils.isBlank(content)) {
                    throw FileNotFoundException("$path not found")
                }
                return content as String
            } catch (e: Exception) {
                logger.error("cannot load prompt for $path", e)
            }
            throw FileNotFoundException("$path not found")
        }

        fun getUnittestSystem() : String {
            return getPrompt("prompt/unittest_system.md")
        }

        fun getUnittestUser() : String {
            return getPrompt("prompt/unittest_user.md")
        }

        fun registerCodebaseVariable(project: Project) {
            promptCache.clear()
            val cache = codebaseVariableCache
            cache[VariableKey.PROJECT_PATH.name] = project.basePath ?: "."
        }

    }

    enum class VariableKey {
        PROJECT_PATH,
        ;
    }

}