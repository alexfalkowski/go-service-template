.PHONY: features

setup-bundler:
	gem install bundler
	bundle config set path 'vendor/bundle'

setup: setup-bundler

lint:
	bundle exec rubocop

fix-lint:
	bundle exec rubocop -A

dep:
	bundle check || bundle install
	bundle clean --force

outdated:
	bundle outdated --only-explicit

update-all:
	bundle update

features: clean-reports
	bundle exec cucumber --profile report --fail-fast $(feature)

clean-reports:
	rm -rf reports/*.*

clean: clean-reports

update-dep:
	bundle update $(gem)
	bundle clean --force
