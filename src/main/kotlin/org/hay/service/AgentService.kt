package org.hay.service

import com.alibaba.fastjson.JSON
import com.alibaba.fastjson.JSONException
import com.intellij.openapi.diagnostic.Logger
import com.intellij.openapi.project.Project
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaTypeOrNull
import okhttp3.RequestBody.Companion.toRequestBody
import okio.IOException
import org.apache.commons.lang3.StringUtils
import org.hay.Constant
import org.hay.service.dto.AgentRequest
import org.hay.service.dto.AgentResponse
import org.hay.service.dto.AssistantMessage
import org.hay.service.dto.SystemMessage
import org.hay.service.dto.UserMessage
import org.hay.service.tools.ToolBehavior
import java.time.Duration
import java.util.concurrent.TimeUnit

class AgentService(project: Project) : Callback {

    private val logger = Logger.getInstance(AgentService::class.java)
    enum class Status {
        INIT, ING, EXCEPTION_FINISH, FINISH;
    }
    private var status = Status.INIT
    private lateinit var agentRequest: AgentRequest

    fun start(systemMessage: String, userMessage: String) {
        if (status == Status.ING) {
            logger.warn("还在生成中")
            return
        }
        // request dto
        agentRequest = AgentRequest(ArrayList(), ToolBehavior.allTools())
        agentRequest.messages.add(SystemMessage(systemMessage))
        agentRequest.messages.add(UserMessage(userMessage))

        // http client
        val client = OkHttpClient.Builder()
            .connectTimeout(5, TimeUnit.MINUTES)
            .readTimeout(5, TimeUnit.MINUTES)
            .build()

        val requestBody = JSON.toJSONString(agentRequest)
            .toRequestBody("application/json;charset=UTF-8".toMediaTypeOrNull())
        val httpRequest = Request.Builder().url(Constant.LLM_SERVER_URL)
            .post(requestBody)
            .addHeader("Content-Type", "application/json")
            .build()
        status = Status.ING
        client.newCall(httpRequest).enqueue(this)
    }

    override fun onFailure(call: Call, e: IOException) {
        logger.error("onFailure", e)
        status = Status.EXCEPTION_FINISH
    }

    override fun onResponse(call: Call, response: Response) {
        try {
            response.body.use { body ->
                body.source().use {
                    while (true) {
                        val line = it.readUtf8Line() ?: break
                        println(line)
                    }
                }
            }
        } catch (e: Exception) {
            logger.error("onResponse exception", e)
            status = Status.EXCEPTION_FINISH
        } finally {
            status = Status.FINISH
        }
    }

    fun processLine(line: String) {
        val agentResponse: AgentResponse
        try {
            agentResponse = JSON.parseObject(line, AgentResponse::class.java)
        } catch (e: JSONException) {
            logger.error("json: $line parse exception", e)
            return
        }
        val contents = StringBuilder()
        for (c in agentResponse.choices) {
            contents.append(c.delta.content ?: "")
            if (c.delta.toolCalls.isNotEmpty()) {
                agentRequest.messages.add(AssistantMessage(contents.toString()))
                contents.clear()


            }
        }
    }

}