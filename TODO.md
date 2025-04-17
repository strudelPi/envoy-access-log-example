# TODOs
- predelat backend na nginx
- na connection manager jde nastavit i tcp i http logy
    - oba maji `access_log_type` = 6 (`DownstreamEnd`) - asi kvuli nastaveni kdy se pisou logy, slo by zmenit? chceme?
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