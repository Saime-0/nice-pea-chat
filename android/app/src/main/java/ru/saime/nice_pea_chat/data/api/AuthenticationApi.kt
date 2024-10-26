package ru.saime.nice_pea_chat.data.api

import retrofit2.http.GET
import retrofit2.http.Path
import retrofit2.http.Query
import java.time.LocalDateTime

interface AuthenticationApi {
    @GET("{server}/authn")
    suspend fun authn(
        @Path("server", encoded = true) server: String,
        @Query("token") token: String,
    ): Result<AuthnResult>

    @GET("{server}/authn/login")
    suspend fun login(
        @Path("server", encoded = true) server: String,
        @Query("key") key: String,
    ): Result<LoginResult>
}



data class User(
    val id: Int,
    val username: String,
    val createdAt: LocalDateTime // Используем LocalDateTime для представления времени
)

data class Session(
    val id: Int,
    val userId: Int,
    val token: String,
    val createdAt: LocalDateTime,
    val expiresAt: LocalDateTime
)

data class AuthnResult(
    val user: User,
    val session: Session
)

data class LoginResult(
    val user: User,
    val session: Session
)
