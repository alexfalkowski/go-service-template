lint:
	buf lint || true

format:
	buf format -w

generate:
	buf generate

update-all:
	buf mod update

breaking:
	buf breaking --against 'https://github.com/alexfalkowski/go-service-template.git#branch=master,subdir=api' || true
