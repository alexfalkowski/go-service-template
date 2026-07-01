# frozen_string_literal: true

When('the system requests the health status with gRPC') do
  @response = Example.health_grpc.check
end

Then('the system should respond with a healthy status with gRPC') do
  expect(@response.status).to eq(:SERVING)
end
