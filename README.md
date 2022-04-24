# market-project



hey -n 100000 -c 4 -q 200 http://localhost:3500/api/product/list/1

hey -n 100000 -c 4 -q 200 -m POST -T "application/json"  -d '{"name":"Ram DDR4 P2666 Gigabyte","quantity":10,"unit":"peace","price":20,"price_unit":"dollar","user_id":10,"company_id":1}' http://localhost:3500/api/product/upload