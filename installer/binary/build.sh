echo "==> gathering dependencies..."

echo "  cloning firefox-boxes/boxes-ipc"
git clone https://github.com/firefox-boxes/boxes-ipc.git
echo "  cloning firefox-boxes/boxes-shell"
git clone https://github.com/firefox-boxes/boxes-shell.git

chmod +x boxes-ipc/build.sh
chmod +x boxes-shell/build.sh

echo ""
echo "==> building boxes-ipc"
cd boxes-ipc && ./build.sh && cd ..

echo ""
echo "==> building boxes-shell"
cd boxes-shell && ./build.sh && cd ..

echo ""
echo "==> gathering executables"
mkdir dist
cp boxes-ipc/dist/* dist/
cp boxes-shell/dist/* dist/

echo ""
echo "==> clean up"
rm -rf boxes-ipc
rm -rf boxes-shell