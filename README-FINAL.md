# Microsserviços gRPC — Parte Final

## Implementado
- Shipping (gRPC), cálculo do prazo.
- Order valida itens de estoque (tabela `inventory_items`).
- Order chama Shipping **após** pagamento ok (adapter incluso).
- Dockerfiles + docker-compose.
- Kubernetes

## Rodar
```
cd microservices-proto
./run.sh

cd ../microservices
docker compose up --build
```
