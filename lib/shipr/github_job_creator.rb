module Shipr
  # Public: This class is responsible for taking the params provided by
  # GitHub's deployment event and adapting it to the interface required by
  # Shipr::JobCreator.
  class GitHubJobCreator
    def initialize(params)
      @params = Hashie::Mash.new(params)
    end

    def self.create(*args)
      new(*args).create
    end

    def create
      JobCreator.create name, attributes
    end

    private

    attr_reader :params

    def name
      params.name
    end

    def attributes
      { sha: params.sha,
        guid: params.id,
        force: payload.force,
        environment: payload.environment,
        config: payload.config,
        description: params.description }.delete_if { |k,v| v.nil? }
    end

    def payload
      params.payload || Hashie::Mash.new
    end
  end
end
