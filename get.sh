set -euo pipefail

VERSTION=${TEMPLATE_VERSION:-v0.0.3-3}
PREFIX=${TEMPLATE_INSTALL_PREFIX:-"/usr/local/bin"}

if [ -z "$TEMPLATE_OS" ]; then
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
else
  OS="$TEMPLATE_OS"
fi

if [ -z "$TEMPLATE_ARCH" ]; then
  case "$(uname -m)" in
    x86_64)
      ARCH="amd64"
    x86)
      ARCH="386"
      ;;
    *)
      echo "Could not determine ARCH (must be one of amd64 or i386)"
      exit 1;
  esac
else
  ARCH="$TEMPLATE_ARCH"
fi

URL_BASE="https://github.com/srossross/template/releases/download"
echo "Fetching tar: ${URL_BASE}/${VERSTION}/template-$OS-$ARCH.tgz"
curl --fail ${URL_BASE}/${VERSTION}/template-$OS-$ARCH.tgz -o template-$OS-$ARCH.tgz

tar -xvf template-$OS-$ARCH.tgz
chmod +x ./template-$OS-$ARCH
mv template-$OS-$ARCH ${PREFIX}/template

${PREFIX}/template version
