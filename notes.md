# Notes
## Lua scripts
- `httpCall()` from lua script will only be logged on the receiving (listener) side; see [logs](./data/lua_access_logs_loopback_service.json)
    - in case the call resolves to a an envoy listener
        - `downstream_direct_remote` is IP of the original listener + ephemeral port (new socket), `downstream_remote` got set to another IP I had on the interface, not sure why
