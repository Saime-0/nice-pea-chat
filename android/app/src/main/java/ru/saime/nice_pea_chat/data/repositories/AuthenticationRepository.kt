package ru.saime.nice_pea_chat.data.repositories

import ru.saime.nice_pea_chat.data.api.AuthenticationApi
import ru.saime.nice_pea_chat.data.api.AuthnResult
import ru.saime.nice_pea_chat.data.api.LoginResult
import ru.saime.nice_pea_chat.data.store.NpcClientStore
import ru.saime.nice_pea_chat.network.authzHeaderValue

class AuthenticationRepository(
    private val api: AuthenticationApi,
    private val npcStore: NpcClientStore,
) {
    suspend fun authn(token: String, server: String = npcStore.baseUrl): Result<AuthnResult> {
        return api.authn(
            server = server,
            token = authzHeaderValue(token),
        )
    }

    suspend fun login(key: String, server: String = npcStore.baseUrl): Result<LoginResult> {
        return api.login(
            server = server,
            key = key,
        )
    }
}