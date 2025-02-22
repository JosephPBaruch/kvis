kubectl run test-pod --rm -it --image=busybox -- /bin/sh

wget -O- http://kvis-client.default.svc.cluster.local:8081
