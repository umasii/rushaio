package task

var FTL_ORDERCHK_CHLS = `[{"status":"COMPLETE","method":"POST","protocolVersion":"HTTP/2.0","scheme":"https","host":"www.footlocker.com","actualPort":-1,"path":"/api/users/orders/status","query":"timestamp=1607898588263","remoteAddress":"/151.101.194.132","clientAddress":null,"clientPort":0,"times":{"start":"2020-12-13T14:29:48.264-08:00","requestBegin":"2020-12-13T14:29:48.266-08:00","requestComplete":"2020-12-13T14:29:48.266-08:00","responseBegin":"2020-12-13T14:29:48.625-08:00","end":"2020-12-13T14:29:48.626-08:00"},"durations":{"total":360,"dns":null,"connect":null,"ssl":null,"request":0,"response":1,"latency":359},"speeds":{"overall":194,"request":0,"response":0},"totalSize":70,"request":{"sizes":{"headers":0,"body":70},"mimeType":"application/json","charset":null,"contentEncoding":null,"header":{"headers":[{"name":":method","value":"POST"},{"name":":authority","value":"www.footlocker.com"},{"name":":scheme","value":"https"},{"name":":path","value":"/api/users/orders/status?timestamp=1607898588263"},{"name":"content-length","value":"70"},{"name":"accept","value":"application/json"},{"name":"x-csrf-token","value":"79e6d93c-1d4c-4088-8732-d9c9e82aefb2"},{"name":"user-agent","value":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.101 Safari/537.36"},{"name":"x-fl-request-id","value":"b5633770-3d92-11eb-9880-f918ed803562"},{"name":"content-type","value":"application/json"},{"name":"sec-gpc","value":"1"},{"name":"origin","value":"https://www.footlocker.com"},{"name":"sec-fetch-site","value":"same-origin"},{"name":"sec-fetch-mode","value":"cors"},{"name":"sec-fetch-dest","value":"empty"},{"name":"referer","value":"https://www.footlocker.com/order/search.html"},{"name":"accept-encoding","value":"gzip, deflate, br"},{"name":"accept-language","value":"en-US,en;q=0.9"},{"name":"cookie","value":"at_check=true; AMCVS_40A3741F578E26BA7F000101%40AdobeOrg=1; se=spell_on; ku1-sid=S0kJNkYV-BIJNkHKI_D_2; ku1-vid=7fc88a42-f0e1-35c5-5bd6-3d3e14a3cb86; BVImplmain_site=8001; AMCV_40A3741F578E26BA7F000101%40AdobeOrg=-432600572%7CMCIDTS%7C18610%7CMCMID%7C74459837118264485021102531848811151607%7CMCOPTOUT-1607905759s%7CNONE%7CvVersion%7C4.5.2; JSESSIONID=1xuzsn523bh2g1ueq0jau0nyu0.fzcexflapipdb648882; mbox=PC#82a8a85dffd04e9494cdf81b40a6ad35.35_0#1671143360|session#5d8fa37bb1c840d5b60d2c762b75f042#1607900420"}]},"body":{"text":"{\"code\":\"V2001624425\",\"customerEmail\":\"HeatherJackson1@recsemail.com\"}","charset":null}},"response":{"status":-1,"sizes":{"headers":0,"body":0},"mimeType":"application/json","charset":null,"contentEncoding":"gzip","header":{"headers":[{"name":"server","value":"nginx"},{"name":"content-type","value":"application/json"},{"name":"access-control-allow-origin","value":"*"},{"name":"breadcrumbid","value":"ID-cb729394a048-1606936063319-1-44267152"},{"name":"cdn-loop","value":"Fastly"},{"name":"fastly-cachetype","value":"PASS"},{"name":"fastly-client-ip","value":"75.39.160.46"},{"name":"fastly-ff","value":"5MQcJXFgiskeUxWO+tAd2TIjT+BGVd2NnourcOKFgqw=!PAO!cache-pao17450-PAO"},{"name":"fastly-force-shield","value":"1"},{"name":"fastly-orig-accept-encoding","value":"gzip, deflate, br"},{"name":"fastly-soc-x-request-id","value":"90eca04fc277d88fafe9080f827087a11eff1933ca665e994c3bed2fbf3454ed"},{"name":"fastly-ssl","value":"1"},{"name":"locid","value":"e70153a8-aeae-4e5f-80f8-15cb01df78d6-oms"},{"name":"locid","value":"55293fe9-4ee1-4f08-9f1b-630abd3ae116-api"},{"name":"origin","value":"https://www.footlocker.com"},{"name":"request-context","value":"appId=cid-v1:d9f7dc17-cae7-46be-bccb-4ec17fd21f96"},{"name":"response_token","value":"b5633770-3d92-11eb-9880-f918ed803562"},{"name":"sec-fetch-dest","value":"empty"},{"name":"sec-fetch-mode","value":"cors"},{"name":"sec-fetch-site","value":"same-origin"},{"name":"sec-gpc","value":"1"},{"name":"signal","value":"spell_on"},{"name":"srvid","value":"-136"},{"name":"srvid","value":"-132"},{"name":"true-client-ip","value":"75.39.160.46"},{"name":"x-akamai-device-characteristics","value":"is_mobile=false;is_tablet=false"},{"name":"x-api-client-id","value":"FLAPI"},{"name":"x-api-correlation-id","value":"b5633770-3d92-11eb-9880-f918ed803562"},{"name":"x-api-gateway-version","value":"1.0.2097"},{"name":"x-aspnet-version","value":"4.0.30319"},{"name":"x-csrf-token","value":"79e6d93c-1d4c-4088-8732-d9c9e82aefb2"},{"name":"x-datadome-headers-pairs","value":"x-datadome-headers=X-DataDome%20Set-Cookie,x-datadome=protected,set-cookie=datadome%3D5zISnNk27yu7Nj8gImu1RIyjV8_r~OlwIgNT8yLHGIF-I6EA3PGWWeksJL6lt7~vwnD_b.95p-_7B_uKD~6tTlvJhrdnjY92qzurdUUcez%3B%20Max-Age%3D31536000%3B%20Domain%3D.footlocker.com%3B%20Path%3D%2F%3B%20Secure%3B%20SameSite%3DLax"},{"name":"x-fl-asnum","value":"7018"},{"name":"x-fl-request-id","value":"b5633770-3d92-11eb-9880-f918ed803562"},{"name":"x-flapi-inventory-system","value":"obf"},{"name":"x-footlocker-newrelic-transaction-name","value":"/{siteId}/users/orders/status"},{"name":"x-forwarded-host","value":"www.footlocker.com"},{"name":"x-newrelic-app-data","value":"PxQFV15XCQEFR1RUAQgCVlMGAhFORDQHUjZKA1ZLVVFHDFYPbU5gEhZfWQYlDFZHQgsNDlJDa0kSAmocARMQFl8PXRBLZBtBVgRLAllBGzIrZWNIRE8IHQBIUUwHBgVRUAMABU5SUk4SWg5VXghRBgYAW1EEBQBbBxQbBwcPS1Zt"},{"name":"x-powered-by","value":"ASP.NET"},{"name":"content-encoding","value":"gzip"},{"name":"accept-ranges","value":"bytes"},{"name":"date","value":"Sun, 13 Dec 2020 22:29:48 GMT"},{"name":"via","value":"1.1 varnish"},{"name":"x-datadome","value":"protected"},{"name":"set-cookie","value":"datadome=5zISnNk27yu7Nj8gImu1RIyjV8_r~OlwIgNT8yLHGIF-I6EA3PGWWeksJL6lt7~vwnD_b.95p-_7B_uKD~6tTlvJhrdnjY92qzurdUUcez; Max-Age=31536000; Domain=.footlocker.com; Path=/; Secure; SameSite=Lax"},{"name":"x-served-by","value":"cache-pao17450-PAO"},{"name":"x-cache","value":"MISS"},{"name":"x-cache-hits","value":"0"},{"name":"x-timer","value":"S1607898588.275059,VS0,VE350"},{"name":"strict-transport-security","value":"max-age=31557600"},{"name":"x-fl-qit-debug","value":"start"},{"name":"x-akamai-edgescape","value":"country_code=US,zip=95648"},{"name":"vary","value":"Accept-Encoding"},{"name":"x-fl-edge","value":"Fastly"}]},"body":{"text":"{\"orderNumber\":\"V2001624425\",\"customerNumber\":\"HSZPZ0HVIQS\",\"orderDate\":\"Dec. 13th, 2020\",\"controllerOrderDate\":\"12/13/2020\",\"orderStatus\":\"Cancelled\",\"shipAddresses\":[{\"firstName\":\"Matthew\",\"lastName\":\"J Valverde\",\"street1\":\"1539 Scriven Avenue AB12\",\"street2\":\"\",\"city\":\"N. Bellmore\",\"state\":\"NY\",\"zip\":\"11710-2805\",\"countryName\":\"US\",\"countryCode\":\"\",\"phone\":\"7178666120\",\"contactID\":0,\"index\":1}],\"billContact\":{\"firstName\":\"Matthew\",\"lastName\":\"J Valverde\",\"street1\":\"1539 Scriven Avenue AB12\",\"street2\":\"\",\"city\":\"N. Bellmore\",\"state\":\"NY\",\"zip\":\"11710-2805\",\"countryName\":\"US\",\"countryCode\":\"\",\"phone\":\"7188106333\",\"contactID\":0,\"index\":1},\"lineItems\":[{\"productNumber\":\"392734\",\"productDescription\":\"Jordan AJ 1 Mid - Boys' Grade School\",\"unitPrice\":\"90.0\",\"orderQuantity\":\"2\",\"itemStatus\":\"Cancelled\",\"shipChargeType\":\"\",\"shipToIndex\":1,\"size\":\"04.5\",\"sku\":\"54725073\",\"shipAmount\":0.0},{\"productNumber\":\"392737\",\"productDescription\":\"Jordan AJ 1 Mid - Boys' Grade School\",\"unitPrice\":\"90.0\",\"orderQuantity\":\"1\",\"itemStatus\":\"Cancelled\",\"shipChargeType\":\"\",\"shipToIndex\":1,\"size\":\"06.0\",\"sku\":\"54725073\",\"shipAmount\":0.0}],\"shipments\":[],\"netAmount\":\"270.00\",\"discountAmount\":\"0.00\",\"couponAmount\":\"0.00\",\"taxAmount\":\"12.50\",\"shipAmount\":\"0.00\",\"orderTotal\":\"282.50\"}","charset":null}}}]`