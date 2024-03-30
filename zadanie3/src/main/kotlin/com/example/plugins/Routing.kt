package com.example.plugins

import com.example.Message
import com.example.tryToSendMessage
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import kotlinx.serialization.json.Json

fun Application.configureRouting() {
    routing {
        post("/sendMessage") {
            val receivedMessage = Json.decodeFromString<Message>(call.receiveText())
            val messageStatus = tryToSendMessage(receivedMessage.content)
            call.respondText(messageStatus)
        }
    }
}
