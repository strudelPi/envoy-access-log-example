# TODOs
- na connection manager jde nastavit i tcp i http logy
    - oba maji `access_log_type` = 6 (`DownstreamEnd`) - asi kvuli nastaveni kdy se pisou logy, slo by zmenit? chceme?
        -> TODO
## Testovani
- pro 1 curl req (1 req, 1 connection)
    - 2 logy (1http s http logerem, 1 tcp s tcp logerem)
        - oba maji stejny `duration` a `stream_id` (pro http stejny jako `request_id`)
- 1 connection - vice req
> `curl --http1.1 http://localhost:8000 --next http://localhost:8000`
    - i s envoyem nastaveny na reuse connection(curl hlasi `Re-using existing connection`), je `stream_id` ruzny, stejne tak jsou tam 2 tcp logy namisto 1
        - stejny chovani i s `--header "Connection: keep-alive"`
- websockets
    - 1 spojeni 1 req (x zprav skrz ws) - oba type 6, 1tcp 1 http (stejny `stream_id`)
        - TODO proc?
- http2
    - parralel streams
        - stejny chovani jako http1.1 keep-alive: 1 req = 1tcp+1http log, neni nijak poznat ze vice req sdili jedno spojeni (test pomoci `sudo tcpdump -i lo -nn port 8010` -> vidim jediny client port a jediny set syn->ack)
            - testovano i s forcovanim http2 na backendu (pro jistotu?)
        -> todle mi proste nesedi na docs `unique id of stream (TCP connection, long-live HTTP2 stream, HTTP request)`
### Listener logy
- http2 mx(2req): s listener logovanim k 4 logum (2 per req) pribyvaji 2  logy
    - 1 tcp OK, 1http -> proc 1??