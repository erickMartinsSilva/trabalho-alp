echo "Iniciando back-end (load balancer + APIs)..."

cd loadbalancer
docker compose up -d --build
if [ $? -ne 0 ]; then
  echo "Erro ao iniciar back-end. Verifique os logs do Docker."
  exit 1
fi

echo "Iniciando interface..."
cd ../cli/src
go run .
if [ $? -ne 0 ]; then
  echo "Erro ao iniciar interface. Verifique os logs do Go."
  exit 1
fi