ARCH := $(shell go env GOARCH)
PHONY: install
install:
	go build -o terraform-provider-tftest
	mkdir -p ~/.terraform.d/plugins/local/prashantv/tftest/0.1.0/darwin_$(ARCH)
	cp terraform-provider-tftest ~/.terraform.d/plugins/local/prashantv/tftest/0.1.0/darwin_$(ARCH)
	make clean-tf

.PHONY: clean
clean: clean-tf
	rm -rf ~/.terraform.d/plugins/local/prashantv/tftest

.PHONY: clean-tf
clean-tf:
	cd repro && rm -rf db .terraform .terraform.lock.hcl terraform.tfstate*

repro/.terraform.lock.hcl:
	cd repro && terraform init

.PHONY: apply
apply: repro/.terraform.lock.hcl
	cd repro && terraform apply -auto-approve

.PHONY: destroy
destroy: repro/.terraform.lock.hcl
	cd repro && CUSTOMIZE_DIFF_PANIC=1 terraform destroy

.PHONY: destroy-no-refresh
destroy-no-refresh: repro/.terraform.lock.hcl
	cd repro && CUSTOMIZE_DIFF_PANIC=1 terraform destroy -refresh=false

