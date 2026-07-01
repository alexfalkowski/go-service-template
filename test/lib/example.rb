# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

module Example
  class << self
    def config
      @config ||= Nonnative::ConfigurationFile.load('.config/server.yml')
    end

    def health_grpc
      @health_grpc ||= Nonnative.grpc_health(host: 'localhost', port: 12_000, service: 'example.v1.Service')
    end

    def user_agent
      @user_agent ||= Nonnative::Header.grpc_user_agent('Example-ruby-client/1.0 gRPC/1.0')
    end
  end

  module V1
  end
end
