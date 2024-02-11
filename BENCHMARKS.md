# Go Expert Challenge - Rate Limiter Benchmarks

Resultados dos testes de carga realizados com o [Grafana k6](https://k6.io/) para avaliar o comportamento da solução desenvolvida sob pressão. Em todos os casos, a aplicação estava disponível a partir de um container Docker respondendo na porta `8080` e o k6 foi executado a partir de outro container.

## Hardware

- **CPU**: AMD Ryzen 7 3800X (8C/16T)
- **RAM**: 32GB DDR4
- **OS**: Arch Linux
- **Runtime**: Docker (alpine)

## Smoke test

Teste com fim de validar se o serviço está respondendo corretamente.

- **Target**: 5 usuários
- **Duração**: 1 minuto

Comando para execução:

```sh
make test_k6_smoke
```

### Resultado

```plaintext
K6_WEB_DASHBOARD=true K6_WEB_DASHBOARD_EXPORT=scripts/k6/smoke-test-report.html k6 run scripts/k6/smoke.test.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

     execution: local
        script: scripts/k6/smoke.test.js
 web dashboard: http://127.0.0.1:5665
        output: -

     scenarios: (100.00%) 1 scenario, 5 max VUs, 1m30s max duration (incl. graceful stop):
              * default: 5 looping VUs for 1m0s (gracefulStop: 30s)


     data_received..................: 265 MB  4.4 MB/s
     data_sent......................: 92 MB   1.5 MB/s
     http_req_blocked...............: min=520ns    med=1.13µs   avg=1.24µs   p(90)=1.59µs   p(95)=1.8µs    max=585.5µs  count=1005218
     http_req_connecting............: min=0s       med=0s       avg=0ns      p(90)=0s       p(95)=0s       max=121.24µs count=1005218
     http_req_duration..............: min=131.3µs  med=231.47µs avg=264.46µs p(90)=332.67µs p(95)=409.97µs max=114.24ms count=1005218
       { expected_response:true }...: min=208.95µs med=325.41µs avg=389.75µs p(90)=499.68µs p(95)=602.65µs max=7.75ms   count=6600   
     ✓ { status:200 }...............: min=208.95µs med=325.41µs avg=389.75µs p(90)=499.68µs p(95)=602.65µs max=7.75ms   count=6600   
     ✓ { status:429 }...............: min=131.3µs  med=231.18µs avg=263.63µs p(90)=330.29µs p(95)=407.52µs max=114.24ms count=998618 
     ✓ { status:500 }...............: min=0s       med=0s       avg=0s       p(90)=0s       p(95)=0s       max=0s       count=0      
     http_req_failed................: 99.34%  ✓ 998618       ✗ 6600
     http_req_receiving.............: min=6.48µs   med=17.53µs  avg=18.19µs  p(90)=24.15µs  p(95)=26.77µs  max=3.44ms   count=1005218
     http_req_sending...............: min=2.5µs    med=5.34µs   avg=5.6µs    p(90)=7.2µs    p(95)=8.04µs   max=5.93ms   count=1005218
     http_req_tls_handshaking.......: min=0s       med=0s       avg=0s       p(90)=0s       p(95)=0s       max=0s       count=1005218
     http_req_waiting...............: min=113.25µs med=207.94µs avg=240.66µs p(90)=306.54µs p(95)=381.84µs max=114.2ms  count=1005218
     http_reqs......................: 1005218 16753.447204/s
     iteration_duration.............: min=342.87µs med=529.66µs avg=590.6µs  p(90)=728.73µs p(95)=864.94µs max=114.84ms count=502609 
     iterations.....................: 502609  8376.723602/s
     vus............................: 5       min=5          max=5 
     vus_max........................: 5       min=5          max=5 


running (1m00.0s), 0/5 VUs, 502609 complete and 0 interrupted iterations
default ✓ [======================================] 5 VUs  1m0s
```

## Stress test

Teste com fim de validar se o serviço está respondendo corretamente sob pressão.

- Stage 1:
    - **Target**: 200 usuários
    - **Duração**: 10 minutos
- Stage 2:
    - **Target**: 200 usuários
    - **Duração**: 30 minutos
- Stage 3:
    - **Target**: 0 usuários
    - **Duração**: 5 minutos

Comando para execução:

```sh
make test_k6_stress
```

### Resultado

```plaintext
K6_WEB_DASHBOARD=true K6_WEB_DASHBOARD_EXPORT=scripts/k6/stress/stress-test-report.html k6 run scripts/k6/stress/stress.test.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

     execution: local
        script: scripts/k6/stress/stress.test.js
 web dashboard: http://127.0.0.1:5665
        output: -

     scenarios: (100.00%) 1 scenario, 200 max VUs, 45m30s max duration (incl. graceful stop):
              * default: Up to 200 looping VUs for 45m0s over 3 stages (gracefulRampDown: 30s, gracefulStop: 30s)


     data_received..................: 231 MB 86 kB/s
     data_sent......................: 82 MB  30 kB/s
     http_req_blocked...............: min=470ns    med=1.57µs   avg=1.89µs   p(90)=2.71µs   p(95)=3.22µs   max=30.24ms count=899092
     http_req_connecting............: min=0s       med=0s       avg=53ns     p(90)=0s       p(95)=0s       max=30.2ms  count=899092
     http_req_duration..............: min=156.46µs med=449.96µs avg=518.1µs  p(90)=870.27µs p(95)=1.04ms   max=38.66ms count=899092
       { expected_response:true }...: min=195.84µs med=509.64µs avg=572.93µs p(90)=959.18µs p(95)=1.13ms   max=38.66ms count=273695
     ✓ { status:200 }...............: min=195.84µs med=509.64µs avg=572.93µs p(90)=959.18µs p(95)=1.13ms   max=38.66ms count=273695
     ✓ { status:429 }...............: min=156.46µs med=430.37µs avg=494.1µs  p(90)=819.84µs p(95)=995.28µs max=37.89ms count=625397
     ✓ { status:500 }...............: min=0s       med=0s       avg=0s       p(90)=0s       p(95)=0s       max=0s      count=0     
     http_req_failed................: 69.55% ✓ 625397     ✗ 273695
     http_req_receiving.............: min=6.31µs   med=20.53µs  avg=22.51µs  p(90)=29.42µs  p(95)=36.21µs  max=3.2ms   count=899092
     http_req_sending...............: min=2.82µs   med=6.94µs   avg=7.99µs   p(90)=11.02µs  p(95)=14.04µs  max=2.07ms  count=899092
     http_req_tls_handshaking.......: min=0s       med=0s       avg=0s       p(90)=0s       p(95)=0s       max=0s      count=899092
     http_req_waiting...............: min=137.2µs  med=418.01µs avg=487.59µs p(90)=837.05µs p(95)=1.01ms   max=38.62ms count=899092
     http_reqs......................: 899092 332.971322/s
     iteration_duration.............: min=1s       med=1s       avg=1s       p(90)=1s       p(95)=1s       max=1.03s   count=449546
     iterations.....................: 449546 166.485661/s
     vus............................: 1      min=1        max=200 
     vus_max........................: 200    min=200      max=200 


running (45m00.2s), 000/200 VUs, 449546 complete and 0 interrupted iterations
default ✓ [======================================] 000/200 VUs  45m0s
```
