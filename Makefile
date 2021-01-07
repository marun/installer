clean:
	./hack/virt-delete-sno.sh || true
	rm -rf mydir

generate:
	mkdir -p mydir
	cp ./install-config.yaml mydir/
	OPENSHIFT_INSTALL_RELEASE_IMAGE_OVERRIDE="quay.io/openshift-release-dev/ocp-release:4.6.12-x86_64" ./bin/openshift-install create manifests --dir=mydir
	sed -i 's/test1.//' ./mydir/manifests/cluster-ingress-02-config.yml
	cp ./sno_manifest.yaml mydir/openshift/
	OPENSHIFT_INSTALL_RELEASE_IMAGE_OVERRIDE="quay.io/openshift-release-dev/ocp-release:4.6.12-x86_64" ./bin/openshift-install create ignition-configs --dir=mydir

embed: download_iso
	cp installer-image.iso.bak installer-image.iso
	sudo docker run --privileged --rm -v /dev:/dev -v /run/udev:/run/udev -v `pwd`:/data -w /data quay.io/coreos/coreos-installer:release iso ignition embed /data/installer-image.iso -f --ignition-file /data/mydir/bootstrap.ign -o /data/installer-SNO-image.iso
	mkdir -p /tmp/images
	mv -f installer-SNO-image.iso /tmp/images/installer-SNO-image.iso

download_iso:
	./hack/download_live_iso.sh

start-iso:
	./hack/virt-install-sno-iso-ign.sh
start:
	./hack/virt-install-sno-ign.sh ./mydir/bootstrap.ign

network:
	./hack/virt-create-net.sh

ssh:
	ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no core@192.168.126.10

image:
	curl -O -L https://releases-art-rhcos.svc.ci.openshift.org/art/storage/releases/rhcos-4.6/46.82.202008181646-0/x86_64/rhcos-46.82.202008181646-0-qemu.x86_64.qcow2.gz
	mv rhcos-46.82.202008181646-0-qemu.x86_64.qcow2.gz /tmp
	sudo gunzip /tmp/rhcos-46.82.202008181646-0-qemu.x86_64.qcow2.gz
