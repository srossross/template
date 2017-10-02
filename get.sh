set -e
VERSTION=${TEMPLATE_VERSION:-v0.0.3-3}

case "$(uname)" in
  Linux)
    OS="linux"
    ;;
  Darwin)
    OS="darwin"
    ;;
  *)
    echo "Could not determine OS (must be one of darwin or linux)"
    exit 1;
esac

case "$(uname -m)" in
  x86_64)
    ARCH="amd64"
    ;;
  *)
    echo "Could not determine ARCH (must be one of amd64)"
    exit 1;
esac

URL_BASE="https://github.com/srossross/template/releases/download"
echo "Fetching tar: ${URL_BASE}/${VERSTION}/template-$OS-$ARCH.tgz"
curl --fail ${URL_BASE}/${VERSTION}/template-$OS-$ARCH.tgz -o template-$OS-$ARCH.tgz

tar -xvf template-$OS-$ARCH.tgz
chmod +x ./template-$OS-$ARCH
mv template-$OS-$ARCH /usr/local/bin/template
