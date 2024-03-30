package com.example

import com.example.plugins.configureRouting
import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.client.request.*
import io.ktor.http.*
import io.ktor.http.content.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.serialization.kotlinx.json.*
import io.ktor.util.*
import kotlinx.coroutines.withContext
import kotlinx.coroutines.Dispatchers
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.Json
import kotlinx.serialization.encodeToString


const val DISCORD_WEBHOOK_URL = "https://discord.com/api/webhooks/1221547563878977547/TTvNdzA92QA3VOnwN3qCXFvB8H4VNGeXPmUpmqx9BIQEizSbbaP8XS0Sev0ob2Oadd5q"


@Serializable
data class Message(val content: String)


fun main() {
    val server = createServer()
    server.start(wait = true)
}


fun createServer(): NettyApplicationEngine = embeddedServer(Netty, port = 8080) {
    configureRouting()
}


fun createHttpClient(): HttpClient = HttpClient(CIO) {
    install(ContentNegotiation) {
        json(Json{
            isLenient = true
            ignoreUnknownKeys = true
        })
    }
}


suspend fun tryToSendMessage(message: String): String = withContext(Dispatchers.IO) {
    val httpClient = createHttpClient()
    try {
        sendMessage(message, httpClient)
    } catch (e: Exception) {
        "An error occurred while sending the message: ${e.message}"
    } finally {
        httpClient.close()
    }
}


@OptIn(InternalAPI::class)
private suspend fun sendMessage(message: String, httpClient: HttpClient): String {
    val payload = Json.encodeToString(Message(message))
    val response = httpClient.post(DISCORD_WEBHOOK_URL) {
        contentType(ContentType.Application.Json)
        body = TextContent(payload, ContentType.Application.Json)
    }
    return "Message sent with status: ${response.status}"
}
